pub mod PureScript_Control_Monad_Writer_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_c6e4dcf3::PureScript_Control_Monad_Cont_Class;
    use crate::module_7a9a81dd::PureScript_Control_Monad_Error_Class;
    use crate::module_6641b520::PureScript_Control_Monad_Reader_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_e3c9db92::PureScript_Control_Monad_State_Class;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_c8a91fca::PureScript_Control_Monad_Writer_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Writer_Trans_WriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_WriterT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_WriterT.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Control_Monad_Writer_Trans_runWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_runWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_runWriterT.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_Writer_Trans_newtypeWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_newtypeWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_newtypeWriterT.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                   &&&add(string("Coercible0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Control_Monad_Writer_Trans_monadTransWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadTransWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadTransWriterT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonoid|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                                      &&&add(string("lift"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let dictMonoid
                                                                                                                                                  =
                                                                                                                                                  dictMonoid.clone();
                                                                                                                                              move
                                                                                                                                                  |dictMonad|
                                                                                                                                                  {
                                                                                                                                                      let Bind1 =
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                      let pure_var =
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let Bind1
                                                                                                                                                                          =
                                                                                                                                                                          Bind1.clone();
                                                                                                                                                                      let pure_var
                                                                                                                                                                          =
                                                                                                                                                                          pure_var.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |m|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                    &&&Bind1),
                                                                                                                                                                                                                                                                                 m),
                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                |a|
                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                    &&&pure_var),
                                                                                                                                                                                                                                                                                                 &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictMonoid)))))))
                                                                                                                                                                  })
                                                                                                                                                  }
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Monad_Writer_Trans_mapWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_mapWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_mapWriterT.get_or_init(||
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
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                          &&&matchValue_1))
                                                                                                  }
                                                                                          })))
    }
    pub fn Control_Monad_Writer_Trans_functorWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_functorWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_functorWriterT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictFunctor|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                   &&&add(string("map"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictFunctor
                                                                                                                                               =
                                                                                                                                               dictFunctor.clone();
                                                                                                                                           move
                                                                                                                                               |f|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                   &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_mapWriterT()),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                      &&&dictFunctor),
                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                     let f
                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                         f.clone();
                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                         |v|
                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                             let matchValue:
                                                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                     &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                                                                         }
                                                                                                                                                                                                                                 })))
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_Writer_Trans_execWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_execWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_execWriterT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFunctor|
                                                                               &Func1::new({
                                                                                               let dictFunctor
                                                                                                   =
                                                                                                   dictFunctor.clone();
                                                                                               move
                                                                                                   |v|
                                                                                                   {
                                                                                                       let m =
                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                              &&&dictFunctor),
                                                                                                                                                                           &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                        &&&m)
                                                                                                   }
                                                                                           })))
    }
    pub fn Control_Monad_Writer_Trans_applyWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_applyWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_applyWriterT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictSemigroup|
                                                                                &Func1::new({
                                                                                                let dictSemigroup
                                                                                                    =
                                                                                                    dictSemigroup.clone();
                                                                                                move
                                                                                                    |dictApply|
                                                                                                    {
                                                                                                        let Functor0 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                    Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                        let functorWriterT1 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_functorWriterT(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                                                         &&&add(string("apply"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let Functor0
                                                                                                                                                                     =
                                                                                                                                                                     Functor0.clone();
                                                                                                                                                                 let dictApply
                                                                                                                                                                     =
                                                                                                                                                                     dictApply.clone();
                                                                                                                                                                 move
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
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                                                                                              &&{
                                                                                                                                                                                                                                    let k =
                                                                                                                                                                                                                                        &Func1::new(move
                                                                                                                                                                                                                                                        |v3|
                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                        let v3
                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                            v3.clone();
                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                            |v4|
                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                let matchValue_3:
                                                                                                                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v3);
                                                                                                                                                                                                                                                                                let matchValue_4:
                                                                                                                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(v4);
                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                            },
                                                                                                                                                                                                                                                                                                                                                                         &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                                                                                                         &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                            })))
                                                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                                                    }));
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                           &&&dictApply),
                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                              &&&k),
                                                                                                                                                                                                                                                                                                                                           &&&matchValue)),
                                                                                                                                                                                                                                                                     &&&matchValue_1)
                                                                                                                                                                                                                                })
                                                                                                                                                                                         }
                                                                                                                                                                                 })
                                                                                                                                                             }),
                                                                                                                                                add(string("Functor0"),
                                                                                                                                                    &&Func1::new({
                                                                                                                                                                     let functorWriterT1
                                                                                                                                                                         =
                                                                                                                                                                         functorWriterT1.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |usd__unused|
                                                                                                                                                                         &functorWriterT1
                                                                                                                                                                 }),
                                                                                                                                                    empty::<string,
                                                                                                                                                            &dyn Any>())))
                                                                                                    }
                                                                                            })))
    }
    pub fn Control_Monad_Writer_Trans_bindWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_bindWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_bindWriterT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictSemigroup|
                                                                               {
                                                                                   let applyWriterT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applyWriterT(),
                                                                                                                        dictSemigroup);
                                                                                   &Func1::new({
                                                                                                   let applyWriterT1
                                                                                                       =
                                                                                                       applyWriterT1.clone();
                                                                                                   let dictSemigroup
                                                                                                       =
                                                                                                       dictSemigroup.clone();
                                                                                                   move
                                                                                                       |dictBind|
                                                                                                       {
                                                                                                           let Apply0 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                           let Functor0 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(&&Apply0)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                           let applyWriterT2 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&applyWriterT1,
                                                                                                                                                &&&Apply0);
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                                                            &&&add(string("bind"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let Functor0
                                                                                                                                                                        =
                                                                                                                                                                        Functor0.clone();
                                                                                                                                                                    let dictBind
                                                                                                                                                                        =
                                                                                                                                                                        dictBind.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v|
                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                        let v
                                                                                                                                                                                            =
                                                                                                                                                                                            v.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |k|
                                                                                                                                                                                            {
                                                                                                                                                                                                let matchValue =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(k);
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                          &&&dictBind),
                                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                                                                                      let matchValue_1
                                                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                                                          matchValue_1.clone();
                                                                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                              let matchValue_3:
                                                                                                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                                              let wt =
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                             &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                }));
                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                    let matchValue_3
                                                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                                                        matchValue_3.clone();
                                                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                                                        |v3|
                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                            let matchValue_4:
                                                                                                                                                                                                                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(v3);
                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                     },
                                                                                                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        })))
                                                                                                                                                                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                                                               &&&wt)
                                                                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                                                                  })))
                                                                                                                                                                                            }
                                                                                                                                                                                    })
                                                                                                                                                                }),
                                                                                                                                                   add(string("Apply0"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let applyWriterT2
                                                                                                                                                                            =
                                                                                                                                                                            applyWriterT2.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |usd__unused|
                                                                                                                                                                            &applyWriterT2
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                       }
                                                                                               })
                                                                               }))
    }
    pub fn Control_Monad_Writer_Trans_semigroupWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_semigroupWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_semigroupWriterT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictApply|
                                                                                    &Func1::new({
                                                                                                    let dictApply
                                                                                                        =
                                                                                                        dictApply.clone();
                                                                                                    move
                                                                                                        |dictSemigroup|
                                                                                                        {
                                                                                                            let applyWriterT1 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applyWriterT(),
                                                                                                                                                                                    dictSemigroup),
                                                                                                                                                 &&&dictApply);
                                                                                                            &Func1::new({
                                                                                                                            let applyWriterT1
                                                                                                                                =
                                                                                                                                applyWriterT1.clone();
                                                                                                                            move
                                                                                                                                |dictSemigroup1|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                                                 &&&add(string("append"),
                                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                                             &&&applyWriterT1),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                             dictSemigroup1)),
                                                                                                                                                                        empty::<string,
                                                                                                                                                                                &dyn Any>()))
                                                                                                                        })
                                                                                                        }
                                                                                                })))
    }
    pub fn Control_Monad_Writer_Trans_applicativeWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_applicativeWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_applicativeWriterT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonoid|
                                                                                      {
                                                                                          let applyWriterT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applyWriterT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          &Func1::new({
                                                                                                          let applyWriterT1
                                                                                                              =
                                                                                                              applyWriterT1.clone();
                                                                                                          let dictMonoid
                                                                                                              =
                                                                                                              dictMonoid.clone();
                                                                                                          move
                                                                                                              |dictApplicative|
                                                                                                              {
                                                                                                                  let pure_var =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       dictApplicative);
                                                                                                                  let applyWriterT2 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&applyWriterT1,
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                                                   &&&add(string("pure"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let pure_var
                                                                                                                                                                               =
                                                                                                                                                                               pure_var.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |a|
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                      &&&pure_var),
                                                                                                                                                                                                                                                   &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                                                                                                              &&&dictMonoid)))))
                                                                                                                                                                       }),
                                                                                                                                                          add(string("Apply0"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let applyWriterT2
                                                                                                                                                                                   =
                                                                                                                                                                                   applyWriterT2.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   &applyWriterT2
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Control_Monad_Writer_Trans_monadWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadWriterT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonoid|
                                                                                {
                                                                                    let applicativeWriterT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applicativeWriterT(),
                                                                                                                         dictMonoid);
                                                                                    let bindWriterT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_bindWriterT(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    &Func1::new({
                                                                                                    let applicativeWriterT1
                                                                                                        =
                                                                                                        applicativeWriterT1.clone();
                                                                                                    let bindWriterT1
                                                                                                        =
                                                                                                        bindWriterT1.clone();
                                                                                                    move
                                                                                                        |dictMonad|
                                                                                                        {
                                                                                                            let applicativeWriterT2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&applicativeWriterT1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            let bindWriterT2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&bindWriterT1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                                                             &&&add(string("Applicative0"),
                                                                                                                                                    &&Func1::new({
                                                                                                                                                                     let applicativeWriterT2
                                                                                                                                                                         =
                                                                                                                                                                         applicativeWriterT2.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |usd__unused|
                                                                                                                                                                         &applicativeWriterT2
                                                                                                                                                                 }),
                                                                                                                                                    add(string("Bind1"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let bindWriterT2
                                                                                                                                                                             =
                                                                                                                                                                             bindWriterT2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused_1|
                                                                                                                                                                             &bindWriterT2
                                                                                                                                                                     }),
                                                                                                                                                        empty::<string,
                                                                                                                                                                &dyn Any>())))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Control_Monad_Writer_Trans_monadAskWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadAskWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadAskWriterT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonoid|
                                                                                   {
                                                                                       let monadTransWriterT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                            dictMonoid);
                                                                                       let monadWriterT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                            dictMonoid);
                                                                                       &Func1::new({
                                                                                                       let monadTransWriterT1
                                                                                                           =
                                                                                                           monadTransWriterT1.clone();
                                                                                                       let monadWriterT1
                                                                                                           =
                                                                                                           monadWriterT1.clone();
                                                                                                       move
                                                                                                           |dictMonadAsk|
                                                                                                           {
                                                                                                               let monadWriterT2 =
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                                                &&&add(string("ask"),
                                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                               &&&monadTransWriterT1),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                                                                                                                            dictMonadAsk)),
                                                                                                                                                       add(string("Monad0"),
                                                                                                                                                           &&Func1::new({
                                                                                                                                                                            let monadWriterT2
                                                                                                                                                                                =
                                                                                                                                                                                monadWriterT2.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |usd__unused|
                                                                                                                                                                                &monadWriterT2
                                                                                                                                                                        }),
                                                                                                                                                           empty::<string,
                                                                                                                                                                   &dyn Any>())))
                                                                                                           }
                                                                                                   })
                                                                                   }))
    }
    pub fn Control_Monad_Writer_Trans_monadReaderWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadReaderWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadReaderWriterT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonoid|
                                                                                      {
                                                                                          let monadAskWriterT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadAskWriterT(),
                                                                                                                               dictMonoid);
                                                                                          &Func1::new({
                                                                                                          let monadAskWriterT1
                                                                                                              =
                                                                                                              monadAskWriterT1.clone();
                                                                                                          move
                                                                                                              |dictMonadReader|
                                                                                                              {
                                                                                                                  let monadAskWriterT2 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&monadAskWriterT1,
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadAsk0"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictMonadReader)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                                                                   &&&add(string("local"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let dictMonadReader
                                                                                                                                                                               =
                                                                                                                                                                               dictMonadReader.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |f|
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_mapWriterT(),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                                                                                                      &&&dictMonadReader),
                                                                                                                                                                                                                                                   f))
                                                                                                                                                                       }),
                                                                                                                                                          add(string("MonadAsk0"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let monadAskWriterT2
                                                                                                                                                                                   =
                                                                                                                                                                                   monadAskWriterT2.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   &monadAskWriterT2
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Control_Monad_Writer_Trans_monadContWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadContWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadContWriterT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonoid|
                                                                                    {
                                                                                        let monadWriterT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                             dictMonoid);
                                                                                        &Func1::new({
                                                                                                        let dictMonoid
                                                                                                            =
                                                                                                            dictMonoid.clone();
                                                                                                        let monadWriterT1
                                                                                                            =
                                                                                                            monadWriterT1.clone();
                                                                                                        move
                                                                                                            |dictMonadCont|
                                                                                                            {
                                                                                                                let monadWriterT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonadCont)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_MonadContusd_Dict(),
                                                                                                                                                 &&&add(string("callCC"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let dictMonadCont
                                                                                                                                                                             =
                                                                                                                                                                             dictMonadCont.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |f|
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_callCC(),
                                                                                                                                                                                                                                                                                    &&&dictMonadCont),
                                                                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                                                                   let f
                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                       f.clone();
                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                       |c|
                                                                                                                                                                                                                                                                       &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                     let c
                                                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                                                         c.clone();
                                                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                                                         |a|
                                                                                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&c,
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&dictMonoid)))))
                                                                                                                                                                                                                                                                                                                                                 })))
                                                                                                                                                                                                                                                               })))
                                                                                                                                                                     }),
                                                                                                                                                        add(string("Monad0"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let monadWriterT2
                                                                                                                                                                                 =
                                                                                                                                                                                 monadWriterT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                 &monadWriterT2
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
    pub fn Control_Monad_Writer_Trans_monadEffectWriter() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadEffectWriter:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadEffectWriter.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonoid|
                                                                                     {
                                                                                         let lift =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                                                                 dictMonoid));
                                                                                         let monadWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                              dictMonoid);
                                                                                         &Func1::new({
                                                                                                         let lift
                                                                                                             =
                                                                                                             lift.clone();
                                                                                                         let monadWriterT1
                                                                                                             =
                                                                                                             monadWriterT1.clone();
                                                                                                         move
                                                                                                             |dictMonadEffect|
                                                                                                             {
                                                                                                                 let Monad0 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                 let monadWriterT2 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                      &&&Monad0);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                                                  &&&add(string("liftEffect"),
                                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&lift,
                                                                                                                                                                                                                                                                 &&&Monad0)),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                                              dictMonadEffect)),
                                                                                                                                                         add(string("Monad0"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let monadWriterT2
                                                                                                                                                                                  =
                                                                                                                                                                                  monadWriterT2.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |usd__unused|
                                                                                                                                                                                  &monadWriterT2
                                                                                                                                                                          }),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))
                                                                                                             }
                                                                                                     })
                                                                                     }))
    }
    pub fn Control_Monad_Writer_Trans_monadRecWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadRecWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadRecWriterT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonoid|
                                                                                   {
                                                                                       let Semigroup0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadWriterT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                            dictMonoid);
                                                                                       &Func1::new({
                                                                                                       let Semigroup0
                                                                                                           =
                                                                                                           Semigroup0.clone();
                                                                                                       let dictMonoid
                                                                                                           =
                                                                                                           dictMonoid.clone();
                                                                                                       let monadWriterT1
                                                                                                           =
                                                                                                           monadWriterT1.clone();
                                                                                                       move
                                                                                                           |dictMonadRec|
                                                                                                           {
                                                                                                               let Monad0 =
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                               let Bind1 =
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                           Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                               let Applicative0 =
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                               let monadWriterT2 =
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                    &&&Monad0);
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                                                                &&&add(string("tailRecM"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let Applicative0
                                                                                                                                                                            =
                                                                                                                                                                            Applicative0.clone();
                                                                                                                                                                        let Bind1
                                                                                                                                                                            =
                                                                                                                                                                            Bind1.clone();
                                                                                                                                                                        let dictMonadRec
                                                                                                                                                                            =
                                                                                                                                                                            dictMonadRec.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |f|
                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                            let f
                                                                                                                                                                                                =
                                                                                                                                                                                                f.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |a|
                                                                                                                                                                                                {
                                                                                                                                                                                                    let f_prime =
                                                                                                                                                                                                        &Func1::new(move
                                                                                                                                                                                                                        |v|
                                                                                                                                                                                                                        {
                                                                                                                                                                                                                            let matchValue:
                                                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                            let w =
                                                                                                                                                                                                                                match matchValue.as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                };
                                                                                                                                                                                                                            let wt =
                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                              }));
                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                                                                                &&&wt),
                                                                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                                                                               let w
                                                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                                                   w.clone();
                                                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                                                   |v2|
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                       let matchValue_1:
                                                                                                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                                                                                                                                       let w1 =
                                                                                                                                                                                                                                                                                           match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                           };
                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                           &&&Applicative0),
                                                                                                                                                                                                                                                                                                                        &&{
                                                                                                                                                                                                                                                                                                                              let matchValue_2:
                                                                                                                                                                                                                                                                                                                                      LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                            });
                                                                                                                                                                                                                                                                                                                              match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                  Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_2_1_0)
                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                  &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&w),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&w1))))),
                                                                                                                                                                                                                                                                                                                                  Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                  &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(matchValue_2_0_0,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&w),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&w1))))),
                                                                                                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                                                                                                          })
                                                                                                                                                                                                                                                                                   }
                                                                                                                                                                                                                                                                           }))
                                                                                                                                                                                                                        });
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                                                                                                              &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                           &&&f_prime),
                                                                                                                                                                                                                                                                        &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                                                                                                                                   &&&dictMonoid)))))
                                                                                                                                                                                                }
                                                                                                                                                                                        })
                                                                                                                                                                    }),
                                                                                                                                                       add(string("Monad0"),
                                                                                                                                                           &&Func1::new({
                                                                                                                                                                            let monadWriterT2
                                                                                                                                                                                =
                                                                                                                                                                                monadWriterT2.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |usd__unused|
                                                                                                                                                                                &monadWriterT2
                                                                                                                                                                        }),
                                                                                                                                                           empty::<string,
                                                                                                                                                                   &dyn Any>())))
                                                                                                           }
                                                                                                   })
                                                                                   }))
    }
    pub fn Control_Monad_Writer_Trans_monadStateWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadStateWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadStateWriterT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonoid|
                                                                                     {
                                                                                         let monadTransWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                              dictMonoid);
                                                                                         let monadWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                              dictMonoid);
                                                                                         &Func1::new({
                                                                                                         let monadTransWriterT1
                                                                                                             =
                                                                                                             monadTransWriterT1.clone();
                                                                                                         let monadWriterT1
                                                                                                             =
                                                                                                             monadWriterT1.clone();
                                                                                                         move
                                                                                                             |dictMonadState|
                                                                                                             {
                                                                                                                 let Monad0 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                 let monadWriterT2 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_MonadStateusd_Dict(),
                                                                                                                                                  &&&add(string("state"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let Monad0
                                                                                                                                                                              =
                                                                                                                                                                              Monad0.clone();
                                                                                                                                                                          let dictMonadState
                                                                                                                                                                              =
                                                                                                                                                                              dictMonadState.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |f|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                                                     &&&monadTransWriterT1),
                                                                                                                                                                                                                                                  &&&Monad0),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                                                                                                                                     &&&dictMonadState),
                                                                                                                                                                                                                                                  f))
                                                                                                                                                                      }),
                                                                                                                                                         add(string("Monad0"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let monadWriterT2
                                                                                                                                                                                  =
                                                                                                                                                                                  monadWriterT2.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |usd__unused|
                                                                                                                                                                                  &monadWriterT2
                                                                                                                                                                          }),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))
                                                                                                             }
                                                                                                     })
                                                                                     }))
    }
    pub fn Control_Monad_Writer_Trans_monadTellWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadTellWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadTellWriterT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonoid|
                                                                                    {
                                                                                        let Semigroup0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let monadWriterT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                             dictMonoid);
                                                                                        &Func1::new({
                                                                                                        let Semigroup0
                                                                                                            =
                                                                                                            Semigroup0.clone();
                                                                                                        let monadWriterT1
                                                                                                            =
                                                                                                            monadWriterT1.clone();
                                                                                                        move
                                                                                                            |dictMonad|
                                                                                                            {
                                                                                                                let monadWriterT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                     dictMonad);
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadTellusd_Dict(),
                                                                                                                                                 &&&add(string("tell"),
                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                             &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                  |usd__arg1|
                                                                                                                                                                                                                                                                                  Func1::new({
                                                                                                                                                                                                                                                                                                 let usd__arg1
                                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                                     usd__arg1.clone();
                                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                                     |usd__arg2|
                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                             usd__arg2.clone()))
                                                                                                                                                                                                                                                                                             })),
                                                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                                                        add(string("Semigroup0"),
                                                                                                                                                            &&Func1::new(move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &Semigroup0),
                                                                                                                                                            add(string("Monad1"),
                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                 let monadWriterT2
                                                                                                                                                                                     =
                                                                                                                                                                                     monadWriterT2.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |usd__unused_1|
                                                                                                                                                                                     &monadWriterT2
                                                                                                                                                                             }),
                                                                                                                                                                empty::<string,
                                                                                                                                                                        &dyn Any>()))))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
    pub fn Control_Monad_Writer_Trans_monadWriterWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadWriterWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadWriterWriterT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonoid|
                                                                                      {
                                                                                          let monadTellWriterT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTellWriterT(),
                                                                                                                               dictMonoid);
                                                                                          &Func1::new({
                                                                                                          let dictMonoid
                                                                                                              =
                                                                                                              dictMonoid.clone();
                                                                                                          let monadTellWriterT1
                                                                                                              =
                                                                                                              monadTellWriterT1.clone();
                                                                                                          move
                                                                                                              |dictMonad|
                                                                                                              {
                                                                                                                  let Bind1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                  let Applicative0 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                  let pure_var =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       &&&Applicative0);
                                                                                                                  let pure1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       &&&Applicative0);
                                                                                                                  let monadTellWriterT2 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&monadTellWriterT1,
                                                                                                                                                       dictMonad);
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadWriterusd_Dict(),
                                                                                                                                                   &&&add(string("listen"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let Bind1
                                                                                                                                                                               =
                                                                                                                                                                               Bind1.clone();
                                                                                                                                                                           let pure_var
                                                                                                                                                                               =
                                                                                                                                                                               pure_var.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v|
                                                                                                                                                                               {
                                                                                                                                                                                   let m =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                             &&&Bind1),
                                                                                                                                                                                                                                                                                          &&&m),
                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                         |v1|
                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                             let matchValue:
                                                                                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                             let w =
                                                                                                                                                                                                                                                                                 match matchValue.as_ref()
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                 };
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                 &&&pure_var),
                                                                                                                                                                                                                                                                                                              &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                                                                &w)),
                                                                                                                                                                                                                                                                                                                                                                        &w)))
                                                                                                                                                                                                                                                                         })))
                                                                                                                                                                               }
                                                                                                                                                                       }),
                                                                                                                                                          add(string("pass"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let Bind1
                                                                                                                                                                                   =
                                                                                                                                                                                   Bind1.clone();
                                                                                                                                                                               let pure1
                                                                                                                                                                                   =
                                                                                                                                                                                   pure1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |v_1|
                                                                                                                                                                                   {
                                                                                                                                                                                       let m_1 =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                 &&&Bind1),
                                                                                                                                                                                                                                                                                              &&&m_1),
                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                             |v1_1|
                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                 let matchValue_1:
                                                                                                                                                                                                                                                                                         LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                                                                                                 let activePatternResult:
                                                                                                                                                                                                                                                                                         LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                        });
                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                     &&&pure1),
                                                                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                                             },
                                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                },
                                                                                                                                                                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                }))))
                                                                                                                                                                                                                                                                             })))
                                                                                                                                                                                   }
                                                                                                                                                                           }),
                                                                                                                                                              add(string("Monoid0"),
                                                                                                                                                                  &&Func1::new(move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   &dictMonoid),
                                                                                                                                                                  add(string("MonadTell1"),
                                                                                                                                                                      &&Func1::new({
                                                                                                                                                                                       let monadTellWriterT2
                                                                                                                                                                                           =
                                                                                                                                                                                           monadTellWriterT2.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |usd__unused_1|
                                                                                                                                                                                           &monadTellWriterT2
                                                                                                                                                                                   }),
                                                                                                                                                                      empty::<string,
                                                                                                                                                                              &dyn Any>())))))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Control_Monad_Writer_Trans_monadThrowWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadThrowWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadThrowWriterT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonoid|
                                                                                     {
                                                                                         let monadTransWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                              dictMonoid);
                                                                                         let monadWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                              dictMonoid);
                                                                                         &Func1::new({
                                                                                                         let monadTransWriterT1
                                                                                                             =
                                                                                                             monadTransWriterT1.clone();
                                                                                                         let monadWriterT1
                                                                                                             =
                                                                                                             monadWriterT1.clone();
                                                                                                         move
                                                                                                             |dictMonadThrow|
                                                                                                             {
                                                                                                                 let Monad0 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                 let monadWriterT2 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                                                                  &&&add(string("throwError"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let Monad0
                                                                                                                                                                              =
                                                                                                                                                                              Monad0.clone();
                                                                                                                                                                          let dictMonadThrow
                                                                                                                                                                              =
                                                                                                                                                                              dictMonadThrow.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |e|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                                                     &&&monadTransWriterT1),
                                                                                                                                                                                                                                                  &&&Monad0),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                                                                                     &&&dictMonadThrow),
                                                                                                                                                                                                                                                  e))
                                                                                                                                                                      }),
                                                                                                                                                         add(string("Monad0"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let monadWriterT2
                                                                                                                                                                                  =
                                                                                                                                                                                  monadWriterT2.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |usd__unused|
                                                                                                                                                                                  &monadWriterT2
                                                                                                                                                                          }),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))
                                                                                                             }
                                                                                                     })
                                                                                     }))
    }
    pub fn Control_Monad_Writer_Trans_monadErrorWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadErrorWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadErrorWriterT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonoid|
                                                                                     {
                                                                                         let monadThrowWriterT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadThrowWriterT(),
                                                                                                                              dictMonoid);
                                                                                         &Func1::new({
                                                                                                         let monadThrowWriterT1
                                                                                                             =
                                                                                                             monadThrowWriterT1.clone();
                                                                                                         move
                                                                                                             |dictMonadError|
                                                                                                             {
                                                                                                                 let monadThrowWriterT2 =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&monadThrowWriterT1,
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadThrow0"),
                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictMonadError)),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                                                                  &&&add(string("catchError"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let dictMonadError
                                                                                                                                                                              =
                                                                                                                                                                              dictMonadError.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v|
                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                              let v
                                                                                                                                                                                                  =
                                                                                                                                                                                                  v.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |h|
                                                                                                                                                                                                  {
                                                                                                                                                                                                      let matchValue =
                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                      let matchValue_1 =
                                                                                                                                                                                                          Sharpurs_Prelude::unbox(h);
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                          &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                                                                                                                                                &&&dictMonadError),
                                                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                                                                                            let matchValue_1
                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                matchValue_1.clone();
                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                |e|
                                                                                                                                                                                                                                                                                                &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                            e))
                                                                                                                                                                                                                                                                                        })))
                                                                                                                                                                                                  }
                                                                                                                                                                                          })
                                                                                                                                                                      }),
                                                                                                                                                         add(string("MonadThrow0"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let monadThrowWriterT2
                                                                                                                                                                                  =
                                                                                                                                                                                  monadThrowWriterT2.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |usd__unused|
                                                                                                                                                                                  &monadThrowWriterT2
                                                                                                                                                                          }),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))
                                                                                                             }
                                                                                                     })
                                                                                     }))
    }
    pub fn Control_Monad_Writer_Trans_monadSTWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadSTWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadSTWriterT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonoid|
                                                                                  {
                                                                                      let lift =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                                                              dictMonoid));
                                                                                      let monadWriterT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                           dictMonoid);
                                                                                      &Func1::new({
                                                                                                      let lift
                                                                                                          =
                                                                                                          lift.clone();
                                                                                                      let monadWriterT1
                                                                                                          =
                                                                                                          monadWriterT1.clone();
                                                                                                      move
                                                                                                          |dictMonadST|
                                                                                                          {
                                                                                                              let Monad0 =
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                              let monadWriterT2 =
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                   &&&Monad0);
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                                               &&&add(string("liftST"),
                                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&lift,
                                                                                                                                                                                                                                                              &&&Monad0)),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                                                           dictMonadST)),
                                                                                                                                                      add(string("Monad0"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let monadWriterT2
                                                                                                                                                                               =
                                                                                                                                                                               monadWriterT2.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |usd__unused|
                                                                                                                                                                               &monadWriterT2
                                                                                                                                                                       }),
                                                                                                                                                          empty::<string,
                                                                                                                                                                  &dyn Any>())))
                                                                                                          }
                                                                                                  })
                                                                                  }))
    }
    pub fn Control_Monad_Writer_Trans_monoidWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monoidWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monoidWriterT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictApplicative|
                                                                                 {
                                                                                     let semigroupWriterT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_semigroupWriterT(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     &Func1::new({
                                                                                                     let dictApplicative
                                                                                                         =
                                                                                                         dictApplicative.clone();
                                                                                                     let semigroupWriterT1
                                                                                                         =
                                                                                                         semigroupWriterT1.clone();
                                                                                                     move
                                                                                                         |dictMonoid|
                                                                                                         {
                                                                                                             let applicativeWriterT1 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applicativeWriterT(),
                                                                                                                                                                                     dictMonoid),
                                                                                                                                                  &&&dictApplicative);
                                                                                                             let semigroupWriterT2 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&semigroupWriterT1,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             &Func1::new({
                                                                                                                             let applicativeWriterT1
                                                                                                                                 =
                                                                                                                                 applicativeWriterT1.clone();
                                                                                                                             let semigroupWriterT2
                                                                                                                                 =
                                                                                                                                 semigroupWriterT2.clone();
                                                                                                                             move
                                                                                                                                 |dictMonoid1|
                                                                                                                                 {
                                                                                                                                     let semigroupWriterT3 =
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&semigroupWriterT2,
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictMonoid1)),
                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                                                      &&&add(string("mempty"),
                                                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                  &&&applicativeWriterT1),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                  dictMonoid1)),
                                                                                                                                                                             add(string("Semigroup0"),
                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                  let semigroupWriterT3
                                                                                                                                                                                                      =
                                                                                                                                                                                                      semigroupWriterT3.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |usd__unused|
                                                                                                                                                                                                      &semigroupWriterT3
                                                                                                                                                                                              }),
                                                                                                                                                                                 empty::<string,
                                                                                                                                                                                         &dyn Any>())))
                                                                                                                                 }
                                                                                                                         })
                                                                                                         }
                                                                                                 })
                                                                                 }))
    }
    pub fn Control_Monad_Writer_Trans_altWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_altWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_altWriterT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictAlt|
                                                                              {
                                                                                  let functorWriterT1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_functorWriterT(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictAlt)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                                                   &&&add(string("alt"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictAlt
                                                                                                                                               =
                                                                                                                                               dictAlt.clone();
                                                                                                                                           move
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
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                 &&&dictAlt),
                                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          add(string("Functor0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let functorWriterT1
                                                                                                                                                   =
                                                                                                                                                   functorWriterT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &functorWriterT1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Control_Monad_Writer_Trans_plusWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_plusWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_plusWriterT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictPlus|
                                                                               {
                                                                                   let altWriterT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_altWriterT(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                    &&&add(string("empty"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                                                                dictPlus)),
                                                                                                                           add(string("Alt0"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let altWriterT1
                                                                                                                                                    =
                                                                                                                                                    altWriterT1.clone();
                                                                                                                                                move
                                                                                                                                                    |usd__unused|
                                                                                                                                                    &altWriterT1
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>())))
                                                                               }))
    }
    pub fn Control_Monad_Writer_Trans_alternativeWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_alternativeWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_alternativeWriterT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonoid|
                                                                                      {
                                                                                          let applicativeWriterT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_applicativeWriterT(),
                                                                                                                               dictMonoid);
                                                                                          &Func1::new({
                                                                                                          let applicativeWriterT1
                                                                                                              =
                                                                                                              applicativeWriterT1.clone();
                                                                                                          move
                                                                                                              |dictAlternative|
                                                                                                              {
                                                                                                                  let applicativeWriterT2 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&applicativeWriterT1,
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  let plusWriterT1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_plusWriterT(),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                                                   &&&add(string("Applicative0"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let applicativeWriterT2
                                                                                                                                                                               =
                                                                                                                                                                               applicativeWriterT2.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |usd__unused|
                                                                                                                                                                               &applicativeWriterT2
                                                                                                                                                                       }),
                                                                                                                                                          add(string("Plus1"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let plusWriterT1
                                                                                                                                                                                   =
                                                                                                                                                                                   plusWriterT1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                                   &plusWriterT1
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Control_Monad_Writer_Trans_monadPlusWriterT() -> &dyn Any {
        static Control_Monad_Writer_Trans_monadPlusWriterT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Trans_monadPlusWriterT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonoid|
                                                                                    {
                                                                                        let monadWriterT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadWriterT(),
                                                                                                                             dictMonoid);
                                                                                        let alternativeWriterT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_alternativeWriterT(),
                                                                                                                             dictMonoid);
                                                                                        &Func1::new({
                                                                                                        let alternativeWriterT1
                                                                                                            =
                                                                                                            alternativeWriterT1.clone();
                                                                                                        let monadWriterT1
                                                                                                            =
                                                                                                            monadWriterT1.clone();
                                                                                                        move
                                                                                                            |dictMonadPlus|
                                                                                                            {
                                                                                                                let monadWriterT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&monadWriterT1,
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                let alternativeWriterT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&alternativeWriterT1,
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alternative1"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                                                 &&&add(string("Monad0"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let monadWriterT2
                                                                                                                                                                             =
                                                                                                                                                                             monadWriterT2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &monadWriterT2
                                                                                                                                                                     }),
                                                                                                                                                        add(string("Alternative1"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let alternativeWriterT2
                                                                                                                                                                                 =
                                                                                                                                                                                 alternativeWriterT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                 &alternativeWriterT2
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
}
