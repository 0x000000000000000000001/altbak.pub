pub mod PureScript_Effect_Aff {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::NativeArray_::new_array;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_7a9a81dd::PureScript_Control_Monad_Error_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_255d4e69::PureScript_Control_Parallel_Class;
    use crate::module_cd8e8e29::PureScript_Control_Parallel;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_ceb943a5::PureScript_Effect_Exception;
    use crate::module_90c22cd8::PureScript_Effect_Unsafe;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    use fable_library_rust::System::Lazy_1;
    pub mod Effect_Aff_FFI {
        use super::*;
        use fable_library_rust::Array_::find as find_1;
        use fable_library_rust::Async_::Async;
        use fable_library_rust::Async_::awaitTask;
        use fable_library_rust::Async_::cancel as cancel_1;
        use fable_library_rust::Async_::createCancellationToken;
        use fable_library_rust::Async_::fromContinuations;
        use fable_library_rust::Async_::isCancellationRequested;
        use fable_library_rust::Async_::start;
        use fable_library_rust::Async_::startAsTask;
        use fable_library_rust::Async_::startChild;
        use fable_library_rust::Async_::startWithContinuations;
        use fable_library_rust::AsyncBuilder_::bind;
        use fable_library_rust::AsyncBuilder_::delay;
        use fable_library_rust::AsyncBuilder_::r_return;
        use fable_library_rust::AsyncBuilder_::singleton;
        use fable_library_rust::AsyncBuilder_::zero;
        use fable_library_rust::Map_::item;
        use fable_library_rust::Monitor_::lock;
        use fable_library_rust::Native_::Arc;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::refCell;
        use fable_library_rust::Native_::referenceEquals;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::Option_::defaultArg;
        use fable_library_rust::Reflection_::name;
        use fable_library_rust::Task_::Task;
        use fable_library_rust::Task_::new;
        use crate::module_aa21d1e7::Sharpurs_Prelude::SharpursRuntime;
        use fable_library_rust::System::AggregateException;
        use fable_library_rust::System::Exception;
        use fable_library_rust::System::OperationCanceledException;
        use fable_library_rust::System::Reflection::MethodInfo;
        use fable_library_rust::System::Threading::CancellationToken;
        use fable_library_rust::System::Threading::CancellationTokenRegistration;
        use fable_library_rust::System::Threading::CancellationTokenSource;
        use fable_library_rust::System::Threading::Tasks::TaskCanceledException;
        use fable_library_rust::System::Threading::Tasks::TaskCompletionSource_1;
        pub fn _unwrapException(e: LrcPtr<Exception>) -> LrcPtr<Exception> {
            let e = MutCell::new(e.clone());
            '__unwrapException:
                loop  {
                    break '__unwrapException
                        (if let Some(e) =
                                (e as
                                     &dyn Any).downcast_ref::<LrcPtr<AggregateException>>()
                            {
                             if referenceEquals(&defaultOf::<&dyn Any>(),
                                                &1_i32) {
                                 let ae_1: LrcPtr<AggregateException> =
                                     e.get();
                                 {
                                     let e_temp;
                                     e.set(e_temp);
                                     continue '__unwrapException
                                 }
                             } else { e.get() }
                         } else { e.get() }) ;
                }
        }
        #[derive(Clone, Debug, Default,)]
        pub struct AffState {
            pub Token: CancellationToken,
            pub KillError: LrcPtr<MutCell<Option<LrcPtr<Exception>>>>,
            pub Supervisor: Option<LrcPtr<CancellationTokenSource>>,
        }
        impl core::fmt::Display for
         PureScript_Effect_Aff::Effect_Aff_FFI::AffState {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug,)]
        pub struct NativeFiber {
            aff: Func1<LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>,
                       Arc<Async<&dyn Any>>>,
            tcs: LrcPtr<TaskCompletionSource_1<Result<&dyn Any,
                                                      LrcPtr<Exception>>>>,
            cts: LrcPtr<CancellationTokenSource>,
            killError: LrcPtr<MutCell<Option<LrcPtr<Exception>>>>,
            started: MutCell<bool>,
        }
        impl PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber {
            pub fn _ctor__480546E6(aff:
                                       Func1<LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>,
                                             Arc<Async<&dyn Any>>>)
             -> LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber> {
                let aff_1;
                let tcs;
                let cts: LrcPtr<CancellationTokenSource>;
                let killError: LrcPtr<MutCell<Option<LrcPtr<Exception>>>>;
                let started: bool;
                ();
                aff_1 = aff;
                tcs = new();
                cts = createCancellationToken();
                killError.set(refCell(None::<LrcPtr<Exception>>));
                started = false;
                ();
                LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber{aff:
                                                                                   aff_1,
                                                                               tcs:
                                                                                   tcs,
                                                                               cts:
                                                                                   cts,
                                                                               killError:
                                                                                   killError,
                                                                               started:
                                                                                   MutCell::new(started),})
            }
            pub fn get_Cts(&self) -> LrcPtr<CancellationTokenSource> {
                self.cts.clone()
            }
            pub fn get_Aff(&self)
             ->
                 Func1<LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>,
                       Arc<Async<&dyn Any>>> {
                self.aff.clone()
            }
            pub fn get_KillError(&self) -> Option<LrcPtr<Exception>> {
                self.killError.get()
            }
            pub fn Cancel_229D3F39(&self, ex: LrcPtr<Exception>) {
                (self.killError).set(Some(ex));
                cancel_1(self.cts.clone())
            }
            pub fn Start(&self) {
                lock(self.clone(),
                     Func0::new(move ||
                                    if !self.started.get() {
                                        self.started.set(true);
                                        SharpursRuntime::EventLoopAdd(1_i32);
                                        startWithContinuations((self.aff)(LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::AffState{Token:
                                                                                                                                          self.cts.clone(),
                                                                                                                                      KillError:
                                                                                                                                          self.killError.clone(),
                                                                                                                                      Supervisor:
                                                                                                                                          None::<LrcPtr<CancellationTokenSource>>,})),
                                                               Func1::new(move
                                                                              |res|
                                                                              {
                                                                                  ();
                                                                                  SharpursRuntime::EventLoopDone()
                                                                              }),
                                                               Func1::new(move
                                                                              |ex:
                                                                                   LrcPtr<Exception>|
                                                                              {
                                                                                  ();
                                                                                  SharpursRuntime::EventLoopDone()
                                                                              }),
                                                               Func1::new(move
                                                                              |ex_1:
                                                                                   LrcPtr<OperationCanceledException>|
                                                                              {
                                                                                  ();
                                                                                  SharpursRuntime::EventLoopDone()
                                                                              }),
                                                               Some(createCancellationToken()))
                                    }));
            }
            pub fn WaitAsync(&self)
             -> Arc<Async<Result<&dyn Any, LrcPtr<Exception>>>> {
                awaitTask(defaultOf::<&dyn Any>())
            }
            pub fn get_IsSuspended(&self) -> bool { !self.started.get() }
        }
        impl core::fmt::Display for
         PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn _applyFn(func: &dyn Any, arg: &dyn Any) -> &dyn Any {
            let method: LrcPtr<dyn MethodInfo> =
                find_1(Func1::new(move |m: LrcPtr<dyn MethodInfo>|
                                      if name(m.clone()) == string("Invoke") {
                                          count(m[string("parameters")]) ==
                                              1_i32
                                      } else { false }),
                       defaultOf::<&dyn Any>());
            defaultOf::<&dyn Any>()
        }
        pub fn _pure(a: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let a = a.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                r_return(a)
                        })
        }
        pub fn _throwError(eVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let eVal = eVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new(move ||
                                                     r_return(panic!("{}",
                                                                     eVal.get_Message(),))))
                        })
        }
        pub fn _catchError(affVal: &dyn Any, kVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            let kVal = kVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         singleton.TryWith(delay(Func0::new({
                                                                                                let ctx
                                                                                                    =
                                                                                                    ctx.clone();
                                                                                                move
                                                                                                    ||
                                                                                                    singleton.ReturnFrom(affVal(ctx.clone()))
                                                                                            })),
                                                                           Func1::new({
                                                                                          let ctx
                                                                                              =
                                                                                              ctx.clone();
                                                                                          move
                                                                                              |_arg:
                                                                                                   LrcPtr<Exception>|
                                                                                              if let Some(_arg)
                                                                                                     =
                                                                                                     (_arg
                                                                                                          as
                                                                                                          &dyn Any).downcast_ref::<LrcPtr<Exception>>()
                                                                                                 {
                                                                                                  let handler =
                                                                                                      Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&kVal,
                                                                                                                                                                               &&_arg));
                                                                                                  singleton.ReturnFrom(handler(ctx.clone()))
                                                                                              } else {
                                                                                                  panic!("{}",
                                                                                                         _arg.get_Message(),);
                                                                                                  defaultOf::<&dyn Any>()
                                                                                              }
                                                                                      }))
                                                 }))
                        })
        }
        pub fn _map(fVal: &dyn Any, affVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            let fVal = fVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         bind(affVal(ctx.clone()),
                                                              Func1::new(move
                                                                             |_arg|
                                                                             r_return(PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&fVal,
                                                                                                                                      _arg))))
                                                 }))
                        })
        }
        pub fn _bind(affVal: &dyn Any, kVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            let kVal = kVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         if isCancellationRequested(ctx.Token.clone())
                                                            {
                                                             r_return(panic!("{}",
                                                                             defaultArg(ctx.KillError.get(),
           Exception::_ctor__Z721C83C5(string("Cancelled"))).get_Message(),))
                                                         } else {
                                                             bind(affVal(ctx.clone()),
                                                                  Func1::new({
                                                                                 let ctx
                                                                                     =
                                                                                     ctx.clone();
                                                                                 move
                                                                                     |_arg|
                                                                                     if isCancellationRequested(ctx.Token.clone())
                                                                                        {
                                                                                         r_return(panic!("{}",
                                                                                                         defaultArg(ctx.KillError.get(),
           Exception::_ctor__Z721C83C5(string("Cancelled"))).get_Message(),))
                                                                                     } else {
                                                                                         let nextAff =
                                                                                             Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&kVal,
                                                                                                                                                                      _arg));
                                                                                         singleton.ReturnFrom(nextAff(ctx.clone()))
                                                                                     }
                                                                             }))
                                                         }
                                                 }))
                        })
        }
        pub fn _delay(rightVal: &dyn Any, msVal: &dyn Any) -> &dyn Any {
            &Func1::new(move
                            |ctx:
                                 LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                            delay(Func0::new({
                                                 let ctx = ctx.clone();
                                                 move ||
                                                     singleton.TryWith(delay(Func0::new(move
                                                                                            ||
                                                                                            bind(awaitTask(defaultOf::<&dyn Any>()),
                                                                                                 Func1::new(move
                                                                                                                |_arg:
                                                                                                                     ()|
                                                                                                                r_return(&defaultOf()))))),
                                                                       Func1::new({
                                                                                      let ctx
                                                                                          =
                                                                                          ctx.clone();
                                                                                      move
                                                                                          |_arg_1:
                                                                                               LrcPtr<Exception>|
                                                                                          if let Some(_arg_1)
                                                                                                 =
                                                                                                 (_arg_1
                                                                                                      as
                                                                                                      &dyn Any).downcast_ref::<LrcPtr<TaskCanceledException>>()
                                                                                             {
                                                                                              r_return(panic!("{}",
                                                                                                              defaultArg(ctx.KillError.get(),
           Exception::_ctor__Z721C83C5(string("Cancelled"))).get_Message(),))
                                                                                          } else {
                                                                                              panic!("{}",
                                                                                                     _arg_1.get_Message(),);
                                                                                              defaultOf::<&dyn Any>()
                                                                                          }
                                                                                  }))
                                             })))
        }
        pub fn _liftEffect(effVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let effVal = effVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new(move ||
                                                     r_return(PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&effVal,
                                                                                                              &defaultOf()))))
                        })
        }
        pub fn _makeFiberNative(affVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            move |_dummy|
                                &PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber::_ctor__480546E6(affVal)
                        })
        }
        pub fn _runFiber(nfVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let nfVal = nfVal.clone();
                            move |_dummy| { nfVal.Start(); &defaultOf() }
                        })
        }
        pub fn _killFiber(nfVal: &dyn Any, errVal: &dyn Any,
                          onErrorVal: &dyn Any, onSuccessVal: &dyn Any)
         -> &dyn Any {
            &Func1::new({
                            let errVal = errVal.clone();
                            let nfVal = nfVal.clone();
                            let onSuccessVal = onSuccessVal.clone();
                            move |_dummy|
                                {
                                    let nf:
                                            LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber> =
                                        nfVal;
                                    nf.Cancel_229D3F39(errVal);
                                    nf.Start();
                                    start(delay(Func0::new({
                                                               let nf =
                                                                   nf.clone();
                                                               move ||
                                                                   bind(nf.WaitAsync(),
                                                                        Func1::new(move
                                                                                       |_arg:
                                                                                            Result<&dyn Any,
                                                                                                   LrcPtr<Exception>>|
                                                                                       {
                                                                                           {
                                                                                               let value =
                                                                                                   PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&onSuccessVal,
                                                                                                                                                                                                    &&defaultOf()),
                                                                                                                                                   &defaultOf());
                                                                                               ()
                                                                                           }
                                                                                           zero()
                                                                                       }))
                                                           })));
                                    &Func1::new(move |_dummy_1| &defaultOf())
                                }
                        })
        }
        pub fn _joinFiber(nfVal: &dyn Any, onErrorVal: &dyn Any,
                          onSuccessVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let nfVal = nfVal.clone();
                            let onErrorVal = onErrorVal.clone();
                            let onSuccessVal = onSuccessVal.clone();
                            move |_dummy|
                                {
                                    let nf:
                                            LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber> =
                                        nfVal;
                                    nf.Start();
                                    start(delay(Func0::new({
                                                               let nf =
                                                                   nf.clone();
                                                               move ||
                                                                   bind(nf.WaitAsync(),
                                                                        Func1::new(move
                                                                                       |_arg:
                                                                                            Result<&dyn Any,
                                                                                                   LrcPtr<Exception>>|
                                                                                       {
                                                                                           let res =
                                                                                               _arg;
                                                                                           match &res
                                                                                               {
                                                                                               Err(res_1_0)
                                                                                               =>
                                                                                               {
                                                                                                   {
                                                                                                       let value_1 =
                                                                                                           PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&onErrorVal,
                                                                                                                                                                                                            &res_1_0),
                                                                                                                                                           &defaultOf());
                                                                                                       ()
                                                                                                   }
                                                                                                   zero()
                                                                                               }
                                                                                               Ok(res_0_0)
                                                                                               =>
                                                                                               {
                                                                                                   {
                                                                                                       let value =
                                                                                                           PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&onSuccessVal,
                                                                                                                                                                                                            res_0_0),
                                                                                                                                                           &defaultOf());
                                                                                                       ()
                                                                                                   }
                                                                                                   zero()
                                                                                               }
                                                                                           }
                                                                                       }))
                                                           })));
                                    &Func1::new(move |_dummy_1| &defaultOf())
                                }
                        })
        }
        pub fn _onCompleteFiber(nfVal: &dyn Any, onCompleteVal: &dyn Any)
         -> &dyn Any {
            &Func1::new(move |_dummy|
                            &Func1::new(move |_dummy_1| &defaultOf()))
        }
        pub fn _isSuspendedFiber(nfVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let nfVal = nfVal.clone();
                            move |_dummy| &nfVal.get_IsSuspended()
                        })
        }
        pub fn _makeAffImpl(buildVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let buildVal = buildVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         {
                                                             let tcs = new();
                                                             let completed:
                                                                     LrcPtr<MutCell<bool>> =
                                                                 LrcPtr::new(MutCell::new(false));
                                                             let lockObj = ();
                                                             let successCb =
                                                                 Func1::new({
                                                                                let completed
                                                                                    =
                                                                                    completed.clone();
                                                                                let lockObj
                                                                                    =
                                                                                    lockObj.clone();
                                                                                move
                                                                                    |res|
                                                                                    {
                                                                                        if lock(lockObj,
                                                                                                Func0::new({
                                                                                                               let completed
                                                                                                                   =
                                                                                                                   completed.clone();
                                                                                                               move
                                                                                                                   ||
                                                                                                                   if completed.get()
                                                                                                                      {
                                                                                                                       false
                                                                                                                   } else {
                                                                                                                       completed.set(true);
                                                                                                                       true
                                                                                                                   }
                                                                                                           }))
                                                                                           {
                                                                                            ();
                                                                                        }
                                                                                        &Func1::new(move
                                                                                                        |_arg|
                                                                                                        &defaultOf())
                                                                                    }
                                                                            });
                                                             let errorCb =
                                                                 Func1::new({
                                                                                let completed
                                                                                    =
                                                                                    completed.clone();
                                                                                let lockObj
                                                                                    =
                                                                                    lockObj.clone();
                                                                                move
                                                                                    |errVal|
                                                                                    {
                                                                                        if lock(lockObj,
                                                                                                Func0::new({
                                                                                                               let completed
                                                                                                                   =
                                                                                                                   completed.clone();
                                                                                                               move
                                                                                                                   ||
                                                                                                                   if completed.get()
                                                                                                                      {
                                                                                                                       false
                                                                                                                   } else {
                                                                                                                       completed.set(true);
                                                                                                                       true
                                                                                                                   }
                                                                                                           }))
                                                                                           {
                                                                                            ();
                                                                                        }
                                                                                        &Func1::new(move
                                                                                                        |_arg_1|
                                                                                                        &defaultOf())
                                                                                    }
                                                                            });
                                                             let canceler =
                                                                 PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&buildVal,
                                                                                                                                                                                                                   &&errorCb),
                                                                                                                                                                  &&successCb),
                                                                                                                 &defaultOf());
                                                             let reg:
                                                                     CancellationTokenRegistration =
                                                                 ctx.Token.register(Func0::new({
                                                                                                   let canceler
                                                                                                       =
                                                                                                       canceler.clone();
                                                                                                   let ctx
                                                                                                       =
                                                                                                       ctx.clone();
                                                                                                   move
                                                                                                       ||
                                                                                                       {
                                                                                                           let cancelAffObj =
                                                                                                               PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&canceler,
                                                                                                                                                               &&defaultArg(ctx.KillError.get(),
                                                                                                                                                                            Exception::_ctor__Z721C83C5(string("Cancelled"))));
                                                                                                           println!("cancelAffObj type: {}",
                                                                                                                    string("System.Object"));
                                                                                                           {
                                                                                                               let t =
                                                                                                                   startAsTask((Sharpurs_Prelude::unbox(&cancelAffObj))(LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::AffState{Token:
                                                                                                                                                                                                                                        createCancellationToken(),
                                                                                                                                                                                                                                    KillError:
                                                                                                                                                                                                                                        refCell(None::<LrcPtr<Exception>>),
                                                                                                                                                                                                                                    Supervisor:
                                                                                                                                                                                                                                        None::<LrcPtr<CancellationTokenSource>>,})),
                                                                                                                               None::<i32>,
                                                                                                                               Some(createCancellationToken()));
                                                                                                               ()
                                                                                                           }
                                                                                                       }
                                                                                               }));
                                                             singleton.TryFinally(delay(Func0::new(move
                                                                                                       ||
                                                                                                       singleton.TryWith(delay(Func0::new(move
                                                                                                                                              ||
                                                                                                                                              bind(awaitTask(defaultOf::<&dyn Any>()),
                                                                                                                                                   Func1::new(move
                                                                                                                                                                  |_arg_2|
                                                                                                                                                                  r_return(_arg_2.clone()))))),
                                                                                                                         Func1::new(move
                                                                                                                                        |_arg_3:
                                                                                                                                             LrcPtr<Exception>|
                                                                                                                                        r_return(panic!("{}",
                                                                                                                                                        PureScript_Effect_Aff::Effect_Aff_FFI::_unwrapException(_arg_3).get_Message(),)))))),
                                                                                  Func0::new(move
                                                                                                 ||
                                                                                                 defaultOf::<&dyn Any>()))
                                                         }
                                                 }))
                        })
        }
        pub fn _forkAffNative(affVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         {
                                                             let nf:
                                                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber> =
                                                                 PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber::_ctor__480546E6(affVal);
                                                             singleton.Combine({
                                                                                   let matchValue:
                                                                                           Option<LrcPtr<CancellationTokenSource>> =
                                                                                       ctx.Supervisor.clone();
                                                                                   match &matchValue
                                                                                       {
                                                                                       None
                                                                                       =>
                                                                                       {
                                                                                           ();
                                                                                           zero()
                                                                                       }
                                                                                       Some(matchValue_0_0)
                                                                                       =>
                                                                                       {
                                                                                           let supCts:
                                                                                                   LrcPtr<CancellationTokenSource> =
                                                                                               matchValue_0_0.clone();
                                                                                           {
                                                                                               let value:
                                                                                                       CancellationTokenRegistration =
                                                                                                   supCts.register(Func0::new({
                                                                                                                                  let nf
                                                                                                                                      =
                                                                                                                                      nf.clone();
                                                                                                                                  move
                                                                                                                                      ||
                                                                                                                                      cancel_1(nf.get_Cts())
                                                                                                                              }));
                                                                                               ()
                                                                                           }
                                                                                           zero()
                                                                                       }
                                                                                   }
                                                                               },
                                                                               delay(Func0::new({
                                                                                                    let nf
                                                                                                        =
                                                                                                        nf.clone();
                                                                                                    move
                                                                                                        ||
                                                                                                        r_return(&nf)
                                                                                                })))
                                                         }
                                                 }))
                        })
        }
        pub fn _makeSupervisedFiber(affVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            move |_dummy|
                                {
                                    let supCts:
                                            LrcPtr<CancellationTokenSource> =
                                        createCancellationToken();
                                    let nf:
                                            LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber> =
                                        PureScript_Effect_Aff::Effect_Aff_FFI::NativeFiber::_ctor__480546E6(Func1::new({
                                                                                                                           let supCts
                                                                                                                               =
                                                                                                                               supCts.clone();
                                                                                                                           move
                                                                                                                               |ctx:
                                                                                                                                    LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                                                                                                               delay(Func0::new({
                                                                                                                                                    let ctx
                                                                                                                                                        =
                                                                                                                                                        ctx.clone();
                                                                                                                                                    let supCts
                                                                                                                                                        =
                                                                                                                                                        supCts.clone();
                                                                                                                                                    move
                                                                                                                                                        ||
                                                                                                                                                        singleton.ReturnFrom(affVal(LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::AffState{Token:
                                                                                                                                                                                                                                                    ctx.Token.clone(),
                                                                                                                                                                                                                                                KillError:
                                                                                                                                                                                                                                                    ctx.KillError.clone(),
                                                                                                                                                                                                                                                Supervisor:
                                                                                                                                                                                                                                                    Some(supCts.clone()),})))
                                                                                                                                                }))
                                                                                                                       }));
                                    {
                                        let value:
                                                CancellationTokenRegistration =
                                            supCts.register(Func0::new({
                                                                           let nf
                                                                               =
                                                                               nf.clone();
                                                                           move
                                                                               ||
                                                                               cancel_1(nf.get_Cts())
                                                                       }));
                                        ()
                                    }
                                    &add(string("fiber"), &nf,
                                         add(string("supervisor"), &supCts,
                                             empty::<string, &dyn Any>()))
                                }
                        })
        }
        pub fn _killAll(errVal: &dyn Any, supVal: &dyn Any, cbVal: &dyn Any)
         -> &dyn Any {
            &Func1::new({
                            let cbVal = cbVal.clone();
                            let supVal = supVal.clone();
                            move |_dummy|
                                {
                                    cancel_1(supVal);
                                    {
                                        let value =
                                            PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&cbVal,
                                                                                            &defaultOf());
                                        ()
                                    }
                                    &Func1::new(move |_dummy2| &defaultOf())
                                }
                        })
        }
        pub fn _sequential(affVal: &dyn Any) -> &dyn Any { affVal.clone() }
        pub fn awaitUninterruptibly<T: Clone + 'static>(t: Arc<Task<T>>)
         -> Arc<Async<T>> {
            fromContinuations(Func1::new(move
                                             |tupledArg:
                                                  LrcPtr<(Func1<T, ()>,
                                                          Func1<LrcPtr<Exception>,
                                                                ()>,
                                                          Func1<LrcPtr<OperationCanceledException>,
                                                                ()>)>|
                                             {
                                                 let completed:
                                                         MutCell<bool> =
                                                     MutCell::new(false);
                                                 let lockObj = ();
                                                 ()
                                             }))
        }
        pub fn generalBracket(acquireVal: &dyn Any, optionsVal: &dyn Any,
                              useVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let acquireVal = acquireVal.clone();
                            let optionsVal = optionsVal.clone();
                            let useVal = useVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         {
                                                             let emptyCtx:
                                                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState> =
                                                                 LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::AffState{Token:
                                                                                                                                 createCancellationToken(),
                                                                                                                             KillError:
                                                                                                                                 refCell(None::<LrcPtr<Exception>>),
                                                                                                                             Supervisor:
                                                                                                                                 None::<LrcPtr<CancellationTokenSource>>,});
                                                             bind(delay(Func0::new({
                                                                                       let emptyCtx
                                                                                           =
                                                                                           emptyCtx.clone();
                                                                                       move
                                                                                           ||
                                                                                           singleton.TryWith(delay(Func0::new({
                                                                                                                                  let emptyCtx
                                                                                                                                      =
                                                                                                                                      emptyCtx.clone();
                                                                                                                                  move
                                                                                                                                      ||
                                                                                                                                      bind(PureScript_Effect_Aff::Effect_Aff_FFI::awaitUninterruptibly(startAsTask(acquireVal(emptyCtx.clone()),
                                                                                                                                                                                                                   None::<i32>,
                                                                                                                                                                                                                   Some(createCancellationToken()))),
                                                                                                                                           Func1::new(move
                                                                                                                                                          |_arg|
                                                                                                                                                          r_return(Ok(_arg.clone()))))
                                                                                                                              })),
                                                                                                             Func1::new(move
                                                                                                                            |_arg_1:
                                                                                                                                 LrcPtr<Exception>|
                                                                                                                            if let Some(_arg_1)
                                                                                                                                   =
                                                                                                                                   (_arg_1
                                                                                                                                        as
                                                                                                                                        &dyn Any).downcast_ref::<LrcPtr<Exception>>()
                                                                                                                               {
                                                                                                                                r_return(Err(_arg_1.clone()))
                                                                                                                            } else {
                                                                                                                                panic!("{}",
                                                                                                                                       _arg_1.get_Message(),);
                                                                                                                                defaultOf::<Result<&dyn Any,
                                                                                                                                                   LrcPtr<Exception>>>()
                                                                                                                            }))
                                                                                   })),
                                                                  Func1::new({
                                                                                 let ctx
                                                                                     =
                                                                                     ctx.clone();
                                                                                 let emptyCtx
                                                                                     =
                                                                                     emptyCtx.clone();
                                                                                 move
                                                                                     |_arg_2:
                                                                                          Result<&dyn Any,
                                                                                                 LrcPtr<Exception>>|
                                                                                     {
                                                                                         let resourceResult =
                                                                                             _arg_2;
                                                                                         match &resourceResult
                                                                                             {
                                                                                             Ok(resourceResult_0_0)
                                                                                             =>
                                                                                             {
                                                                                                 let resource =
                                                                                                     resourceResult_0_0.clone();
                                                                                                 let options =
                                                                                                     optionsVal;
                                                                                                 if isCancellationRequested(ctx.Token.clone())
                                                                                                    {
                                                                                                     let ex =
                                                                                                         defaultArg(ctx.KillError.get(),
                                                                                                                    Exception::_ctor__Z721C83C5(string("Cancelled")));
                                                                                                     bind(PureScript_Effect_Aff::Effect_Aff_FFI::awaitUninterruptibly(startAsTask((Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&item(string("killed"),
                                                                                                                                                                                                                                                                                                                   options.clone()),
                                                                                                                                                                                                                                                                                                             &&ex),
                                                                                                                                                                                                                                                            &resource)))(emptyCtx.clone()),
                                                                                                                                                                                  None::<i32>,
                                                                                                                                                                                  Some(createCancellationToken()))),
                                                                                                          Func1::new({
                                                                                                                         let ex
                                                                                                                             =
                                                                                                                             ex.clone();
                                                                                                                         move
                                                                                                                             |_arg_3|
                                                                                                                             r_return(panic!("{}",
                                                                                                                                             ex.get_Message(),))
                                                                                                                     }))
                                                                                                 } else {
                                                                                                     let useFn =
                                                                                                         Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&useVal,
                                                                                                                                                                                  &resource));
                                                                                                     bind(delay(Func0::new({
                                                                                                                               let ctx
                                                                                                                                   =
                                                                                                                                   ctx.clone();
                                                                                                                               let useFn
                                                                                                                                   =
                                                                                                                                   useFn.clone();
                                                                                                                               move
                                                                                                                                   ||
                                                                                                                                   singleton.TryWith(delay(Func0::new({
                                                                                                                                                                          let ctx
                                                                                                                                                                              =
                                                                                                                                                                              ctx.clone();
                                                                                                                                                                          let useFn
                                                                                                                                                                              =
                                                                                                                                                                              useFn.clone();
                                                                                                                                                                          move
                                                                                                                                                                              ||
                                                                                                                                                                              bind(useFn(ctx.clone()),
                                                                                                                                                                                   Func1::new(move
                                                                                                                                                                                                  |_arg_4|
                                                                                                                                                                                                  r_return(Ok(_arg_4.clone()))))
                                                                                                                                                                      })),
                                                                                                                                                     Func1::new(move
                                                                                                                                                                    |_arg_5:
                                                                                                                                                                         LrcPtr<Exception>|
                                                                                                                                                                    if let Some(_arg_5)
                                                                                                                                                                           =
                                                                                                                                                                           (_arg_5
                                                                                                                                                                                as
                                                                                                                                                                                &dyn Any).downcast_ref::<LrcPtr<Exception>>()
                                                                                                                                                                       {
                                                                                                                                                                        r_return(Err(_arg_5.clone()))
                                                                                                                                                                    } else {
                                                                                                                                                                        panic!("{}",
                                                                                                                                                                               _arg_5.get_Message(),);
                                                                                                                                                                        defaultOf::<Result<&dyn Any,
                                                                                                                                                                                           LrcPtr<Exception>>>()
                                                                                                                                                                    }))
                                                                                                                           })),
                                                                                                          Func1::new({
                                                                                                                         let ctx
                                                                                                                             =
                                                                                                                             ctx.clone();
                                                                                                                         let emptyCtx
                                                                                                                             =
                                                                                                                             emptyCtx.clone();
                                                                                                                         let options
                                                                                                                             =
                                                                                                                             options.clone();
                                                                                                                         let resource
                                                                                                                             =
                                                                                                                             resource.clone();
                                                                                                                         move
                                                                                                                             |_arg_6:
                                                                                                                                  Result<&dyn Any,
                                                                                                                                         LrcPtr<Exception>>|
                                                                                                                             {
                                                                                                                                 let useResult =
                                                                                                                                     _arg_6;
                                                                                                                                 bind(PureScript_Effect_Aff::Effect_Aff_FFI::awaitUninterruptibly(startAsTask((match &useResult
                                                                                                                                                                                                                   {
                                                                                                                                                                                                                   Err(useResult_1_0)
                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                   {
                                                                                                                                                                                                                       let e_2 =
                                                                                                                                                                                                                           useResult_1_0.clone();
                                                                                                                                                                                                                       if if e_2.get_Message()
                                                                                                                                                                                                                                 ==
                                                                                                                                                                                                                                 string("Cancelled")
                                                                                                                                                                                                                             {
                                                                                                                                                                                                                              true
                                                                                                                                                                                                                          } else {
                                                                                                                                                                                                                              (e_2
                                                                                                                                                                                                                                   as
                                                                                                                                                                                                                                   &dyn Any).is::<LrcPtr<OperationCanceledException>>()
                                                                                                                                                                                                                          }
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                           let actualEx =
                                                                                                                                                                                                                               defaultArg(ctx.KillError.get(),
                                                                                                                                                                                                                                          e_2.clone());
                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&item(string("killed"),
                                                                                                                                                                                                                                                                                                                                                           options.clone()),
                                                                                                                                                                                                                                                                                                                                                     &&actualEx),
                                                                                                                                                                                                                                                                                                    &resource))
                                                                                                                                                                                                                       } else {
                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&item(string("failed"),
                                                                                                                                                                                                                                                                                                                                                           options.clone()),
                                                                                                                                                                                                                                                                                                                                                     &&e_2),
                                                                                                                                                                                                                                                                                                    &resource))
                                                                                                                                                                                                                       }
                                                                                                                                                                                                                   }
                                                                                                                                                                                                                   Ok(useResult_0_0)
                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&item(string("completed"),
                                                                                                                                                                                                                                                                                                                                                   options.clone()),
                                                                                                                                                                                                                                                                                                                                             useResult_0_0),
                                                                                                                                                                                                                                                                                            &resource)),
                                                                                                                                                                                                               })(emptyCtx.clone()),
                                                                                                                                                                                                              None::<i32>,
                                                                                                                                                                                                              Some(createCancellationToken()))),
                                                                                                                                      Func1::new({
                                                                                                                                                     let useResult
                                                                                                                                                         =
                                                                                                                                                         useResult.clone();
                                                                                                                                                     move
                                                                                                                                                         |_arg_7|
                                                                                                                                                         match &useResult
                                                                                                                                                             {
                                                                                                                                                             Err(useResult_1_0)
                                                                                                                                                             =>
                                                                                                                                                             r_return(panic!("{}",
                                                                                                                                                                             useResult_1_0.get_Message(),)),
                                                                                                                                                             Ok(useResult_0_0)
                                                                                                                                                             =>
                                                                                                                                                             r_return(useResult_0_0.clone()),
                                                                                                                                                         }
                                                                                                                                                 }))
                                                                                                                             }
                                                                                                                     }))
                                                                                                 }
                                                                                             }
                                                                                             Err(resourceResult_1_0)
                                                                                             =>
                                                                                             r_return(panic!("{}",
                                                                                                             resourceResult_1_0.get_Message(),)),
                                                                                         }
                                                                                     }
                                                                             }))
                                                         }
                                                 }))
                        })
        }
        pub fn _parAffMap(fVal: &dyn Any, affVal: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let affVal = affVal.clone();
                            let fVal = fVal.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         bind(affVal(ctx.clone()),
                                                              Func1::new(move
                                                                             |_arg|
                                                                             r_return(PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&fVal,
                                                                                                                                      _arg))))
                                                 }))
                        })
        }
        pub fn _parAffApply(aff1Val: &dyn Any, aff2Val: &dyn Any)
         -> &dyn Any {
            &Func1::new({
                            let aff1Val = aff1Val.clone();
                            let aff2Val = aff2Val.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         bind(startChild(aff1Val(ctx.clone())),
                                                              Func1::new({
                                                                             let ctx
                                                                                 =
                                                                                 ctx.clone();
                                                                             move
                                                                                 |_arg:
                                                                                      Arc<Async<&dyn Any>>|
                                                                                 bind(startChild(aff2Val(ctx.clone())),
                                                                                      Func1::new({
                                                                                                     let _arg
                                                                                                         =
                                                                                                         _arg.clone();
                                                                                                     move
                                                                                                         |_arg_1:
                                                                                                              Arc<Async<&dyn Any>>|
                                                                                                         bind(_arg.clone(),
                                                                                                              Func1::new({
                                                                                                                             let _arg_1
                                                                                                                                 =
                                                                                                                                 _arg_1.clone();
                                                                                                                             move
                                                                                                                                 |_arg_2|
                                                                                                                                 bind(_arg_1.clone(),
                                                                                                                                      Func1::new({
                                                                                                                                                     let _arg_2
                                                                                                                                                         =
                                                                                                                                                         _arg_2.clone();
                                                                                                                                                     move
                                                                                                                                                         |_arg_3|
                                                                                                                                                         r_return(PureScript_Effect_Aff::Effect_Aff_FFI::_applyFn(&_arg_2,
                                                                                                                                                                                                                  _arg_3))
                                                                                                                                                 }))
                                                                                                                         }))
                                                                                                 }))
                                                                         }))
                                                 }))
                        })
        }
        pub fn _parAffAlt(aff1Val: &dyn Any, aff2Val: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let aff1Val = aff1Val.clone();
                            let aff2Val = aff2Val.clone();
                            move
                                |ctx:
                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState>|
                                delay(Func0::new({
                                                     let ctx = ctx.clone();
                                                     move ||
                                                         {
                                                             let token:
                                                                     LrcPtr<PureScript_Effect_Aff::Effect_Aff_FFI::AffState> =
                                                                 LrcPtr::new(PureScript_Effect_Aff::Effect_Aff_FFI::AffState{Token:
                                                                                                                                 defaultOf::<&dyn Any>(),
                                                                                                                             KillError:
                                                                                                                                 ctx.KillError.clone(),
                                                                                                                             Supervisor:
                                                                                                                                 None::<LrcPtr<CancellationTokenSource>>,});
                                                             let t1 =
                                                                 startAsTask(aff1Val(token.clone()),
                                                                             None::<i32>,
                                                                             Some(createCancellationToken()));
                                                             let t2 =
                                                                 startAsTask(aff2Val(token),
                                                                             None::<i32>,
                                                                             Some(createCancellationToken()));
                                                             bind(awaitTask(defaultOf::<&dyn Any>()),
                                                                  Func1::new({
                                                                                 let t1
                                                                                     =
                                                                                     t1.clone();
                                                                                 let t2
                                                                                     =
                                                                                     t2.clone();
                                                                                 move
                                                                                     |_arg:
                                                                                          Arc<Task<&dyn Any>>|
                                                                                     {
                                                                                         let firstCompleted =
                                                                                             _arg;
                                                                                         let otherTask =
                                                                                             if referenceEquals(&firstCompleted,
                                                                                                                &t1)
                                                                                                {
                                                                                                 t2.clone()
                                                                                             } else {
                                                                                                 t1.clone()
                                                                                             };
                                                                                         if referenceEquals(&defaultOf::<&dyn Any>(),
                                                                                                            &5_i32)
                                                                                            {
                                                                                             singleton.Combine(if !defaultOf::<&dyn Any>()
                                                                                                                  {
                                                                                                                   cancel_1(defaultOf::<&dyn Any>());
                                                                                                                   zero()
                                                                                                               } else {
                                                                                                                   zero()
                                                                                                               },
                                                                                                               delay(Func0::new({
                                                                                                                                    let firstCompleted
                                                                                                                                        =
                                                                                                                                        firstCompleted.clone();
                                                                                                                                    let otherTask
                                                                                                                                        =
                                                                                                                                        otherTask.clone();
                                                                                                                                    move
                                                                                                                                        ||
                                                                                                                                        singleton.Combine(singleton.TryWith(delay(Func0::new({
                                                                                                                                                                                                 let otherTask
                                                                                                                                                                                                     =
                                                                                                                                                                                                     otherTask.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     ||
                                                                                                                                                                                                     bind(awaitTask(otherTask.clone()),
                                                                                                                                                                                                          Func1::new(move
                                                                                                                                                                                                                         |_arg_1|
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                             ();
                                                                                                                                                                                                                             zero()
                                                                                                                                                                                                                         }))
                                                                                                                                                                                             })),
                                                                                                                                                                            Func1::new(move
                                                                                                                                                                                           |_arg_2:
                                                                                                                                                                                                LrcPtr<Exception>|
                                                                                                                                                                                           {
                                                                                                                                                                                               ();
                                                                                                                                                                                               zero()
                                                                                                                                                                                           })),
                                                                                                                                                          delay(Func0::new({
                                                                                                                                                                               let firstCompleted
                                                                                                                                                                                   =
                                                                                                                                                                                   firstCompleted.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   ||
                                                                                                                                                                                   r_return(firstCompleted.get_result())
                                                                                                                                                                           })))
                                                                                                                                })))
                                                                                         } else {
                                                                                             singleton.TryWith(delay(Func0::new({
                                                                                                                                    let otherTask
                                                                                                                                        =
                                                                                                                                        otherTask.clone();
                                                                                                                                    move
                                                                                                                                        ||
                                                                                                                                        bind(awaitTask(otherTask.clone()),
                                                                                                                                             Func1::new(move
                                                                                                                                                            |_arg_3|
                                                                                                                                                            r_return(_arg_3.clone())))
                                                                                                                                })),
                                                                                                               Func1::new(move
                                                                                                                              |_arg_4:
                                                                                                                                   LrcPtr<Exception>|
                                                                                                                              r_return(panic!("{}",
                                                                                                                                              PureScript_Effect_Aff::Effect_Aff_FFI::_unwrapException(defaultOf::<&dyn Any>()).get_Message(),))))
                                                                                         }
                                                                                     }
                                                                             }))
                                                         }
                                                 }))
                        })
        }
    }
    pub fn Effect_Aff__bind() -> &dyn Any {
        static Effect_Aff__bind: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__bind.get_or_init(||
                                         &Func1::new(move |affVal|
                                                         Func1::new({
                                                                        let affVal
                                                                            =
                                                                            affVal.clone();
                                                                        move
                                                                            |kVal|
                                                                            PureScript_Effect_Aff::Effect_Aff_FFI::_bind(&affVal,
                                                                                                                         kVal)
                                                                    })))
    }
    pub fn Effect_Aff__catchError() -> &dyn Any {
        static Effect_Aff__catchError: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__catchError.get_or_init(||
                                               &Func1::new(move |affVal|
                                                               Func1::new({
                                                                              let affVal
                                                                                  =
                                                                                  affVal.clone();
                                                                              move
                                                                                  |kVal|
                                                                                  PureScript_Effect_Aff::Effect_Aff_FFI::_catchError(&affVal,
                                                                                                                                     kVal)
                                                                          })))
    }
    pub fn Effect_Aff__delay() -> &dyn Any {
        static Effect_Aff__delay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__delay.get_or_init(||
                                          &Func1::new(move |rightVal|
                                                          Func1::new({
                                                                         let rightVal
                                                                             =
                                                                             rightVal.clone();
                                                                         move
                                                                             |msVal|
                                                                             PureScript_Effect_Aff::Effect_Aff_FFI::_delay(&rightVal,
                                                                                                                           msVal)
                                                                     })))
    }
    pub fn Effect_Aff__forkAffNative() -> &dyn Any {
        static Effect_Aff__forkAffNative: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__forkAffNative.get_or_init(||
                                                  &Func1::new(move |affVal|
                                                                  PureScript_Effect_Aff::Effect_Aff_FFI::_forkAffNative(affVal)))
    }
    pub fn Effect_Aff__isSuspendedFiber() -> &dyn Any {
        static Effect_Aff__isSuspendedFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__isSuspendedFiber.get_or_init(||
                                                     &Func1::new(move |nfVal|
                                                                     PureScript_Effect_Aff::Effect_Aff_FFI::_isSuspendedFiber(nfVal)))
    }
    pub fn Effect_Aff__joinFiber() -> &dyn Any {
        static Effect_Aff__joinFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__joinFiber.get_or_init(||
                                              &Func1::new(move |nfVal|
                                                              Func1::new({
                                                                             let nfVal
                                                                                 =
                                                                                 nfVal.clone();
                                                                             move
                                                                                 |onErrorVal|
                                                                                 Func1::new({
                                                                                                let onErrorVal
                                                                                                    =
                                                                                                    onErrorVal.clone();
                                                                                                move
                                                                                                    |onSuccessVal|
                                                                                                    PureScript_Effect_Aff::Effect_Aff_FFI::_joinFiber(&nfVal,
                                                                                                                                                      &onErrorVal,
                                                                                                                                                      onSuccessVal)
                                                                                            })
                                                                         })))
    }
    pub fn Effect_Aff__killAll() -> &dyn Any {
        static Effect_Aff__killAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__killAll.get_or_init(||
                                            &Func1::new(move |errVal|
                                                            Func1::new({
                                                                           let errVal
                                                                               =
                                                                               errVal.clone();
                                                                           move
                                                                               |supVal|
                                                                               Func1::new({
                                                                                              let supVal
                                                                                                  =
                                                                                                  supVal.clone();
                                                                                              move
                                                                                                  |cbVal|
                                                                                                  PureScript_Effect_Aff::Effect_Aff_FFI::_killAll(&errVal,
                                                                                                                                                  &supVal,
                                                                                                                                                  cbVal)
                                                                                          })
                                                                       })))
    }
    pub fn Effect_Aff__killFiber() -> &dyn Any {
        static Effect_Aff__killFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__killFiber.get_or_init(||
                                              &Func1::new(move |nfVal|
                                                              Func1::new({
                                                                             let nfVal
                                                                                 =
                                                                                 nfVal.clone();
                                                                             move
                                                                                 |errVal|
                                                                                 Func1::new({
                                                                                                let errVal
                                                                                                    =
                                                                                                    errVal.clone();
                                                                                                move
                                                                                                    |onErrorVal|
                                                                                                    Func1::new({
                                                                                                                   let onErrorVal
                                                                                                                       =
                                                                                                                       onErrorVal.clone();
                                                                                                                   move
                                                                                                                       |onSuccessVal|
                                                                                                                       PureScript_Effect_Aff::Effect_Aff_FFI::_killFiber(&nfVal,
                                                                                                                                                                         &errVal,
                                                                                                                                                                         &onErrorVal,
                                                                                                                                                                         onSuccessVal)
                                                                                                               })
                                                                                            })
                                                                         })))
    }
    pub fn Effect_Aff__liftEffect() -> &dyn Any {
        static Effect_Aff__liftEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__liftEffect.get_or_init(||
                                               &Func1::new(move |effVal|
                                                               PureScript_Effect_Aff::Effect_Aff_FFI::_liftEffect(effVal)))
    }
    pub fn Effect_Aff__makeAffImpl() -> &dyn Any {
        static Effect_Aff__makeAffImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__makeAffImpl.get_or_init(||
                                                &Func1::new(move |buildVal|
                                                                PureScript_Effect_Aff::Effect_Aff_FFI::_makeAffImpl(buildVal)))
    }
    pub fn Effect_Aff__makeFiberNative() -> &dyn Any {
        static Effect_Aff__makeFiberNative: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__makeFiberNative.get_or_init(||
                                                    &Func1::new(move |affVal|
                                                                    PureScript_Effect_Aff::Effect_Aff_FFI::_makeFiberNative(affVal)))
    }
    pub fn Effect_Aff__makeSupervisedFiber() -> &dyn Any {
        static Effect_Aff__makeSupervisedFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__makeSupervisedFiber.get_or_init(||
                                                        &Func1::new(move
                                                                        |affVal|
                                                                        PureScript_Effect_Aff::Effect_Aff_FFI::_makeSupervisedFiber(affVal)))
    }
    pub fn Effect_Aff__map() -> &dyn Any {
        static Effect_Aff__map: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__map.get_or_init(||
                                        &Func1::new(move |fVal|
                                                        Func1::new({
                                                                       let fVal
                                                                           =
                                                                           fVal.clone();
                                                                       move
                                                                           |affVal|
                                                                           PureScript_Effect_Aff::Effect_Aff_FFI::_map(&fVal,
                                                                                                                       affVal)
                                                                   })))
    }
    pub fn Effect_Aff__onCompleteFiber() -> &dyn Any {
        static Effect_Aff__onCompleteFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__onCompleteFiber.get_or_init(||
                                                    &Func1::new(move |nfVal|
                                                                    Func1::new({
                                                                                   let nfVal
                                                                                       =
                                                                                       nfVal.clone();
                                                                                   move
                                                                                       |onCompleteVal|
                                                                                       PureScript_Effect_Aff::Effect_Aff_FFI::_onCompleteFiber(&nfVal,
                                                                                                                                               onCompleteVal)
                                                                               })))
    }
    pub fn Effect_Aff__parAffAlt() -> &dyn Any {
        static Effect_Aff__parAffAlt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__parAffAlt.get_or_init(||
                                              &Func1::new(move |aff1Val|
                                                              Func1::new({
                                                                             let aff1Val
                                                                                 =
                                                                                 aff1Val.clone();
                                                                             move
                                                                                 |aff2Val|
                                                                                 PureScript_Effect_Aff::Effect_Aff_FFI::_parAffAlt(&aff1Val,
                                                                                                                                   aff2Val)
                                                                         })))
    }
    pub fn Effect_Aff__parAffApply() -> &dyn Any {
        static Effect_Aff__parAffApply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__parAffApply.get_or_init(||
                                                &Func1::new(move |aff1Val|
                                                                Func1::new({
                                                                               let aff1Val
                                                                                   =
                                                                                   aff1Val.clone();
                                                                               move
                                                                                   |aff2Val|
                                                                                   PureScript_Effect_Aff::Effect_Aff_FFI::_parAffApply(&aff1Val,
                                                                                                                                       aff2Val)
                                                                           })))
    }
    pub fn Effect_Aff__parAffMap() -> &dyn Any {
        static Effect_Aff__parAffMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__parAffMap.get_or_init(||
                                              &Func1::new(move |fVal|
                                                              Func1::new({
                                                                             let fVal
                                                                                 =
                                                                                 fVal.clone();
                                                                             move
                                                                                 |affVal|
                                                                                 PureScript_Effect_Aff::Effect_Aff_FFI::_parAffMap(&fVal,
                                                                                                                                   affVal)
                                                                         })))
    }
    pub fn Effect_Aff__pure() -> &dyn Any {
        static Effect_Aff__pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__pure.get_or_init(||
                                         &Func1::new(move |a|
                                                         PureScript_Effect_Aff::Effect_Aff_FFI::_pure(a)))
    }
    pub fn Effect_Aff__runFiber() -> &dyn Any {
        static Effect_Aff__runFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__runFiber.get_or_init(||
                                             &Func1::new(move |nfVal|
                                                             PureScript_Effect_Aff::Effect_Aff_FFI::_runFiber(nfVal)))
    }
    pub fn Effect_Aff__sequential() -> &dyn Any {
        static Effect_Aff__sequential: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__sequential.get_or_init(||
                                               &Func1::new(move |affVal|
                                                               PureScript_Effect_Aff::Effect_Aff_FFI::_sequential(affVal)))
    }
    pub fn Effect_Aff__throwError() -> &dyn Any {
        static Effect_Aff__throwError: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff__throwError.get_or_init(||
                                               &Func1::new(move |eVal|
                                                               PureScript_Effect_Aff::Effect_Aff_FFI::_throwError(eVal)))
    }
    pub fn Effect_Aff_generalBracket() -> &dyn Any {
        static Effect_Aff_generalBracket: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_generalBracket.get_or_init(||
                                                  &Func1::new(move
                                                                  |acquireVal|
                                                                  Func1::new({
                                                                                 let acquireVal
                                                                                     =
                                                                                     acquireVal.clone();
                                                                                 move
                                                                                     |optionsVal|
                                                                                     Func1::new({
                                                                                                    let optionsVal
                                                                                                        =
                                                                                                        optionsVal.clone();
                                                                                                    move
                                                                                                        |useVal|
                                                                                                        PureScript_Effect_Aff::Effect_Aff_FFI::generalBracket(&acquireVal,
                                                                                                                                                              &optionsVal,
                                                                                                                                                              useVal)
                                                                                                })
                                                                             })))
    }
    pub fn Effect_Aff_pure() -> &dyn Any {
        static Effect_Aff_pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_pure.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                         &&&PureScript_Effect::Effect_applicativeEffect()))
    }
    pub fn Effect_Aff_void() -> &dyn Any {
        static Effect_Aff_void: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_void.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                         &&&PureScript_Effect::Effect_functorEffect()))
    }
    pub fn Effect_Aff_void1() -> &dyn Any {
        static Effect_Aff_void1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_void1.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                          &&&PureScript_Effect::Effect_functorEffect()))
    }
    pub fn Effect_Aff_Fiber() -> &dyn Any {
        static Effect_Aff_Fiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Fiber.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Effect_Aff_Canceler() -> &dyn Any {
        static Effect_Aff_Canceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Canceler.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Effect_Aff_newtypeCanceler() -> &dyn Any {
        static Effect_Aff_newtypeCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_newtypeCanceler.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                    &&&add(string("Coercible0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &Sharpurs_Prelude::Prim_undefined()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))
    }
    pub fn Effect_Aff_makeFiber() -> &dyn Any {
        static Effect_Aff_makeFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_makeFiber.get_or_init(||
                                             &Func1::new(move |aff|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                    &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__makeFiberNative(),
                                                                                                                                                                    aff)),
                                                                                              &&&Func1::new(move
                                                                                                                |nf|
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                    &&&PureScript_Effect_Aff::Effect_Aff_pure()),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_Fiber(),
                                                                                                                                                                                    &&&add(string("run"),
                                                                                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__runFiber(),
                                                                                                                                                                                                                             nf),
                                                                                                                                                                                           add(string("kill"),
                                                                                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__killFiber(),
                                                                                                                                                                                                                                 nf),
                                                                                                                                                                                               add(string("join"),
                                                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__joinFiber(),
                                                                                                                                                                                                                                     nf),
                                                                                                                                                                                                   add(string("onComplete"),
                                                                                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__onCompleteFiber(),
                                                                                                                                                                                                                                         nf),
                                                                                                                                                                                                       add(string("isSuspended"),
                                                                                                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__isSuspendedFiber(),
                                                                                                                                                                                                                                             nf),
                                                                                                                                                                                                           empty::<string,
                                                                                                                                                                                                                   &dyn Any>())))))))))))
    }
    pub fn Effect_Aff_makeAff() -> &dyn Any {
        static Effect_Aff_makeAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_makeAff.get_or_init(||
                                           &Func1::new(move |build|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__makeAffImpl(),
                                                                                            &&&Func1::new({
                                                                                                              let build
                                                                                                                  =
                                                                                                                  build.clone();
                                                                                                              move
                                                                                                                  |onError|
                                                                                                                  &Func1::new({
                                                                                                                                  let onError
                                                                                                                                      =
                                                                                                                                      onError.clone();
                                                                                                                                  move
                                                                                                                                      |onSuccess|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&build,
                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                         let onSuccess
                                                                                                                                                                                             =
                                                                                                                                                                                             onSuccess.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |either|
                                                                                                                                                                                             {
                                                                                                                                                                                                 let matchValue:
                                                                                                                                                                                                         LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(either);
                                                                                                                                                                                                 match matchValue.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                     Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                     =>
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&onSuccess,
                                                                                                                                                                                                                                      &&matchValue_1_0),
                                                                                                                                                                                                     Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                                                                                     =>
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&onError,
                                                                                                                                                                                                                                      &&matchValue_0_0),
                                                                                                                                                                                                 }
                                                                                                                                                                                             }
                                                                                                                                                                                     }))
                                                                                                                              })
                                                                                                          }))))
    }
    pub fn Effect_Aff_launchSuspendedAff() -> &dyn Any {
        static Effect_Aff_launchSuspendedAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_launchSuspendedAff.get_or_init(||
                                                      &Func1::new(move |aff|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                       aff)))
    }
    pub fn Effect_Aff_launchAff() -> &dyn Any {
        static Effect_Aff_launchAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_launchAff.get_or_init(||
                                             &Func1::new(move |aff|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                    &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                                                                                    aff)),
                                                                                              &&&Func1::new(move
                                                                                                                |fiber|
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                          &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                       &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                    &&find(string("run"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(fiber)))),
                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                   let fiber
                                                                                                                                                                       =
                                                                                                                                                                       fiber.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused|
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                           &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                        &&&fiber)
                                                                                                                                                               }))))))
    }
    pub fn Effect_Aff_launchAff_() -> &dyn Any {
        static Effect_Aff_launchAff_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_launchAff_.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                  &&&PureScript_Effect_Aff::Effect_Aff_void()),
                                                                               &&&PureScript_Effect_Aff::Effect_Aff_launchAff()))
    }
    pub fn Effect_Aff_functorParAff() -> &dyn Any {
        static Effect_Aff_functorParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_functorParAff.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                  &&&add(string("map"),
                                                                                         &&PureScript_Effect_Aff::Effect_Aff__parAffMap(),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Effect_Aff_functorAff() -> &dyn Any {
        static Effect_Aff_functorAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_functorAff.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                               &&&add(string("map"),
                                                                                      &&PureScript_Effect_Aff::Effect_Aff__map(),
                                                                                      empty::<string,
                                                                                              &dyn Any>())))
    }
    pub fn Effect_Aff_delay() -> &dyn Any {
        static Effect_Aff_delay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_delay.get_or_init(||
                                         &Func1::new(move |v|
                                                         {
                                                             let n =
                                                                 Sharpurs_Prelude::unbox(v);
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                                                                                                    &&&PureScript_Effect_Aff::Effect_Aff__delay()),
                                                                                                                                 &&&Func1::new(move
                                                                                                                                                   |usd__arg1|
                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone())))),
                                                                                              &&&n)
                                                         }))
    }
    pub fn Effect_Aff_bracket() -> &dyn Any {
        static Effect_Aff_bracket: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_bracket.get_or_init(||
                                           &Func1::new(move |acquire|
                                                           &Func1::new({
                                                                           let acquire
                                                                               =
                                                                               acquire.clone();
                                                                           move
                                                                               |completed|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_generalBracket(),
                                                                                                                                                   &&&acquire),
                                                                                                                &&&add(string("killed"),
                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                         completed),
                                                                                                                       add(string("failed"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                             completed),
                                                                                                                           add(string("completed"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                 completed),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))
                                                                       })))
    }
    pub fn Effect_Aff_applyParAff() -> &dyn Any {
        static Effect_Aff_applyParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applyParAff.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                &&&add(string("apply"),
                                                                                       &&PureScript_Effect_Aff::Effect_Aff__parAffApply(),
                                                                                       add(string("Functor0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Effect_Aff::Effect_Aff_functorParAff()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Effect_Aff_semigroupParAff() -> &dyn Any {
        static Effect_Aff_semigroupParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_semigroupParAff.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictSemigroup|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                    &&&add(string("append"),
                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                &&&PureScript_Effect_Aff::Effect_Aff_applyParAff()),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                dictSemigroup)),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Effect_Aff_monadAff_0040440() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let Effect_Aff_applicativeAff_0040443_002d1
                                                                     =
                                                                     Effect_Aff_applicativeAff_0040443_002d1.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     &Effect_Aff_applicativeAff_0040443_002d1.Value
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let Effect_Aff_bindAff_0040441_002d1
                                                                         =
                                                                         Effect_Aff_bindAff_0040441_002d1.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         &Effect_Aff_bindAff_0040441_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_Aff_monadAff_0040440_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_Aff_monadAff_0040440_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_Aff_monadAff_0040440_002d1.get_or_init(||
                                                          Lazy(Effect_Aff_monadAff_0040440.clone()))
    }
    pub fn Effect_Aff_bindAff_0040441() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&PureScript_Effect_Aff::Effect_Aff__bind(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Effect_Aff_applyAff_0040442_002d1
                                                                         =
                                                                         Effect_Aff_applyAff_0040442_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Effect_Aff_applyAff_0040442_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_Aff_bindAff_0040441_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_Aff_bindAff_0040441_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_Aff_bindAff_0040441_002d1.get_or_init(||
                                                         Lazy(Effect_Aff_bindAff_0040441.clone()))
    }
    pub fn Effect_Aff_applyAff_0040442() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&&Effect_Aff_monadAff_0040440_002d1.Value),
                                                add(string("Functor0"),
                                                    &&Func1::new(move
                                                                     |usd__unused|
                                                                     &PureScript_Effect_Aff::Effect_Aff_functorAff()),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_Aff_applyAff_0040442_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_Aff_applyAff_0040442_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_Aff_applyAff_0040442_002d1.get_or_init(||
                                                          Lazy(Effect_Aff_applyAff_0040442.clone()))
    }
    pub fn Effect_Aff_applicativeAff_0040443() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &&PureScript_Effect_Aff::Effect_Aff__pure(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Effect_Aff_applyAff_0040442_002d1
                                                                         =
                                                                         Effect_Aff_applyAff_0040442_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Effect_Aff_applyAff_0040442_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_Aff_applicativeAff_0040443_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_Aff_applicativeAff_0040443_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_Aff_applicativeAff_0040443_002d1.get_or_init(||
                                                                Lazy(Effect_Aff_applicativeAff_0040443.clone()))
    }
    pub fn Effect_Aff_monadAff() -> &dyn Any {
        static Effect_Aff_monadAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadAff.get_or_init(||
                                            Effect_Aff_monadAff_0040440_002d1.Value)
    }
    pub fn Effect_Aff_bindAff() -> &dyn Any {
        static Effect_Aff_bindAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_bindAff.get_or_init(||
                                           Effect_Aff_bindAff_0040441_002d1.Value)
    }
    pub fn Effect_Aff_applyAff() -> &dyn Any {
        static Effect_Aff_applyAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applyAff.get_or_init(||
                                            Effect_Aff_applyAff_0040442_002d1.Value)
    }
    pub fn Effect_Aff_applicativeAff() -> &dyn Any {
        static Effect_Aff_applicativeAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applicativeAff.get_or_init(||
                                                  Effect_Aff_applicativeAff_0040443_002d1.Value)
    }
    pub fn Effect_Aff_pure1() -> &dyn Any {
        static Effect_Aff_pure1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_pure1.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                          &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()))
    }
    pub fn Effect_Aff_cancelWith() -> &dyn Any {
        static Effect_Aff_cancelWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_cancelWith.get_or_init(||
                                              &Func1::new(move |aff|
                                                              &Func1::new({
                                                                              let aff
                                                                                  =
                                                                                  aff.clone();
                                                                              move
                                                                                  |v|
                                                                                  {
                                                                                      let matchValue =
                                                                                          Sharpurs_Prelude::unbox(&&aff);
                                                                                      let matchValue_1 =
                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_generalBracket(),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                   &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                                          &&&add(string("killed"),
                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                  let matchValue_1
                                                                                                                                                                                      =
                                                                                                                                                                                      matchValue_1.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |e|
                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                      let e
                                                                                                                                                                                                          =
                                                                                                                                                                                                          e.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                           &&&e)
                                                                                                                                                                                                  })
                                                                                                                                                                              }),
                                                                                                                                                                 add(string("failed"),
                                                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                          &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff())),
                                                                                                                                                                     add(string("completed"),
                                                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                              &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff())),
                                                                                                                                                                         empty::<string,
                                                                                                                                                                                 &dyn Any>())))),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                          &&&matchValue))
                                                                                  }
                                                                          })))
    }
    pub fn Effect_Aff_finally() -> &dyn Any {
        static Effect_Aff_finally: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_finally.get_or_init(||
                                           &Func1::new(move |fin|
                                                           &Func1::new({
                                                                           let fin
                                                                               =
                                                                               fin.clone();
                                                                           move
                                                                               |a|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_bracket(),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                            &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                                         &&&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                      &&&fin)),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                   a))
                                                                       })))
    }
    pub fn Effect_Aff_invincible() -> &dyn Any {
        static Effect_Aff_invincible: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_invincible.get_or_init(||
                                              &Func1::new(move |a|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_bracket(),
                                                                                                                                                                     a),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                        &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                               &&&PureScript_Effect_Aff::Effect_Aff_pure1())))
    }
    pub fn Effect_Aff_lazyAff() -> &dyn Any {
        static Effect_Aff_lazyAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_lazyAff.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                            &&&add(string("defer"),
                                                                                   &&Func1::new(move
                                                                                                    |f|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                              &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                           &&&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                     f)),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Effect_Aff_parallelAff() -> &dyn Any {
        static Effect_Aff_parallelAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_parallelAff.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_Parallelusd_Dict(),
                                                                                &&&add(string("parallel"),
                                                                                       &&PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce(),
                                                                                       add(string("sequential"),
                                                                                           &&PureScript_Effect_Aff::Effect_Aff__sequential(),
                                                                                           add(string("Apply0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Effect_Aff::Effect_Aff_applyAff()),
                                                                                               add(string("Apply1"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused_1|
                                                                                                                    &PureScript_Effect_Aff::Effect_Aff_applyParAff()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))))
    }
    pub fn Effect_Aff_applicativeParAff() -> &dyn Any {
        static Effect_Aff_applicativeParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applicativeParAff.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                      &&&add(string("pure"),
                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                                                                                                                     &&&PureScript_Effect_Aff::Effect_Aff_parallelAff())),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                  &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff())),
                                                                                             add(string("Apply0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Effect_Aff::Effect_Aff_applyParAff()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Effect_Aff_monoidParAff() -> &dyn Any {
        static Effect_Aff_monoidParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monoidParAff.get_or_init(||
                                                &Func1::new(move |dictMonoid|
                                                                {
                                                                    let semigroupParAff1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_semigroupParAff(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                   Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                     &&&add(string("mempty"),
                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                 &&&PureScript_Effect_Aff::Effect_Aff_applicativeParAff()),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                 dictMonoid)),
                                                                                                            add(string("Semigroup0"),
                                                                                                                &&Func1::new({
                                                                                                                                 let semigroupParAff1
                                                                                                                                     =
                                                                                                                                     semigroupParAff1.clone();
                                                                                                                                 move
                                                                                                                                     |usd__unused|
                                                                                                                                     &semigroupParAff1
                                                                                                                             }),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>())))
                                                                }))
    }
    pub fn Effect_Aff_semigroupCanceler() -> &dyn Any {
        static Effect_Aff_semigroupCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_semigroupCanceler.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                      &&&add(string("append"),
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
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_Canceler(),
                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                         let matchValue_1
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |err|
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel::Control_Parallel_parSequence_(),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Effect_Aff::Effect_Aff_parallelAff()),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Effect_Aff::Effect_Aff_applicativeParAff()),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                                                                                                                                                                                              &&&new_array(&[Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                              err),
                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                              err)]))
                                                                                                                                                                                     }))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Effect_Aff_semigroupAff() -> &dyn Any {
        static Effect_Aff_semigroupAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_semigroupAff.get_or_init(||
                                                &Func1::new(move
                                                                |dictSemigroup|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                 &&&add(string("append"),
                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                             &&&PureScript_Effect_Aff::Effect_Aff_applyAff()),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                             dictSemigroup)),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Effect_Aff_monadEffectAff() -> &dyn Any {
        static Effect_Aff_monadEffectAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadEffectAff.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                   &&&add(string("liftEffect"),
                                                                                          &&PureScript_Effect_Aff::Effect_Aff__liftEffect(),
                                                                                          add(string("Monad0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Effect_Aff::Effect_Aff_monadAff()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Effect_Aff_liftEffect() -> &dyn Any {
        static Effect_Aff_liftEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_liftEffect.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                               &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff()))
    }
    pub fn Effect_Aff_effectCanceler() -> &dyn Any {
        static Effect_Aff_effectCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_effectCanceler.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                      &&&PureScript_Effect_Aff::Effect_Aff_Canceler()),
                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                      &&&PureScript_Effect_Aff::Effect_Aff_liftEffect())))
    }
    pub fn Effect_Aff_joinFiber() -> &dyn Any {
        static Effect_Aff_joinFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_joinFiber.get_or_init(||
                                             &Func1::new(move |v|
                                                             {
                                                                 let t =
                                                                     Sharpurs_Prelude::unbox(v);
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeAff(),
                                                                                                  &&&Func1::new({
                                                                                                                    let t
                                                                                                                        =
                                                                                                                        t.clone();
                                                                                                                    move
                                                                                                                        |k|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                                                                                                            &&&PureScript_Effect_Aff::Effect_Aff_effectCanceler()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&find(string("join"),
                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&t)),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let k
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     k.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |err|
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(err.clone())))
                                                                                                                                                                                                                                             })),
                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                              let k
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  k.clone();
                                                                                                                                                                                                              move
                                                                                                                                                                                                                  |a|
                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                   &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(a.clone())))
                                                                                                                                                                                                          })))
                                                                                                                }))
                                                             }))
    }
    pub fn Effect_Aff_functorFiber() -> &dyn Any {
        static Effect_Aff_functorFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_functorFiber.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                 &&&add(string("map"),
                                                                                        &&Func1::new(move
                                                                                                         |f|
                                                                                                         &Func1::new({
                                                                                                                         let f
                                                                                                                             =
                                                                                                                             f.clone();
                                                                                                                         move
                                                                                                                             |t|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Unsafe::Effect_Unsafe_unsafePerformEffect(),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Effect_Aff::Effect_Aff_functorAff()),
                                                                                                                                                                                                                                                                       &&&f),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_joinFiber(),
                                                                                                                                                                                                                                                                       t))))
                                                                                                                     })),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Effect_Aff_applyFiber() -> &dyn Any {
        static Effect_Aff_applyFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applyFiber.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                               &&&add(string("apply"),
                                                                                      &&Func1::new(move
                                                                                                       |t1|
                                                                                                       &Func1::new({
                                                                                                                       let t1
                                                                                                                           =
                                                                                                                           t1.clone();
                                                                                                                       move
                                                                                                                           |t2|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Unsafe::Effect_Unsafe_unsafePerformEffect(),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_applyAff()),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_joinFiber(),
                                                                                                                                                                                                                                                                                                        &&&t1)),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_joinFiber(),
                                                                                                                                                                                                                                                                     t2))))
                                                                                                                   })),
                                                                                      add(string("Functor0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Effect_Aff::Effect_Aff_functorFiber()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Effect_Aff_applicativeFiber() -> &dyn Any {
        static Effect_Aff_applicativeFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_applicativeFiber.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                     &&&add(string("pure"),
                                                                                            &&Func1::new(move
                                                                                                             |a|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Unsafe::Effect_Unsafe_unsafePerformEffect(),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                       &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                                    a)))),
                                                                                            add(string("Apply0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Effect_Aff::Effect_Aff_applyFiber()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Effect_Aff_forkAff() -> &dyn Any {
        static Effect_Aff_forkAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_forkAff.get_or_init(||
                                           &Func1::new(move |aff|
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                  &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__forkAffNative(),
                                                                                                                                                                  aff)),
                                                                                            &&&Func1::new(move
                                                                                                              |nf|
                                                                                                              {
                                                                                                                  let fiber =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_Fiber(),
                                                                                                                                                       &&&add(string("run"),
                                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__runFiber(),
                                                                                                                                                                                                nf),
                                                                                                                                                              add(string("kill"),
                                                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__killFiber(),
                                                                                                                                                                                                    nf),
                                                                                                                                                                  add(string("join"),
                                                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__joinFiber(),
                                                                                                                                                                                                        nf),
                                                                                                                                                                      add(string("onComplete"),
                                                                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__onCompleteFiber(),
                                                                                                                                                                                                            nf),
                                                                                                                                                                          add(string("isSuspended"),
                                                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__isSuspendedFiber(),
                                                                                                                                                                                                                nf),
                                                                                                                                                                              empty::<string,
                                                                                                                                                                                      &dyn Any>()))))));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                            &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                         &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                                                                            &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__runFiber(),
                                                                                                                                                                                                                                                            nf))),
                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                     let fiber
                                                                                                                                                                         =
                                                                                                                                                                         fiber.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |usd__unused|
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                             &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                          &&&fiber)
                                                                                                                                                                 }))
                                                                                                              }))))
    }
    pub fn Effect_Aff_killFiber() -> &dyn Any {
        static Effect_Aff_killFiber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_killFiber.get_or_init(||
                                             &Func1::new(move |e|
                                                             &Func1::new({
                                                                             let e
                                                                                 =
                                                                                 e.clone();
                                                                             move
                                                                                 |v|
                                                                                 {
                                                                                     let matchValue =
                                                                                         Sharpurs_Prelude::unbox(&&e);
                                                                                     let t =
                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                     let e1 =
                                                                                         matchValue;
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                            &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                                               &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff()),
                                                                                                                                                                                            &&find(string("isSuspended"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&t)))),
                                                                                                                      &&&Func1::new({
                                                                                                                                        let e1
                                                                                                                                            =
                                                                                                                                            e1.clone();
                                                                                                                                        let t
                                                                                                                                            =
                                                                                                                                            t.clone();
                                                                                                                                        move
                                                                                                                                            |suspended|
                                                                                                                                            {
                                                                                                                                                let matchValue_3 =
                                                                                                                                                    Sharpurs_Prelude::unbox(suspended);
                                                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                 &matchValue_3)
                                                                                                                                                    {
                                                                                                                                                    0_i32
                                                                                                                                                    =>
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_liftEffect()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_void1()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&find(string("kill"),
                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&t)),
                                                                                                                                                                                                                                                                                                                                 &&&e1),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))))),
                                                                                                                                                    _
                                                                                                                                                    =>
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeAff(),
                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                       |k|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                              &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                                                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_effectCanceler()),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&find(string("kill"),
                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&t)),
                                                                                                                                                                                                                                                                                                                                                 &&&e1),
                                                                                                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                let k
                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                    k.clone();
                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                    |err|
                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(err.clone())))
                                                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                                                             let k
                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                 k.clone();
                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit())))
                                                                                                                                                                                                                                                                                         }))))),
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                    }))
                                                                                 }
                                                                         })))
    }
    pub fn Effect_Aff_fiberCanceler() -> &dyn Any {
        static Effect_Aff_fiberCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_fiberCanceler.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&&PureScript_Effect_Aff::Effect_Aff_Canceler()),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                     &&&PureScript_Effect_Aff::Effect_Aff_killFiber())))
    }
    pub fn Effect_Aff_supervise() -> &dyn Any {
        static Effect_Aff_supervise: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_supervise.get_or_init(||
                                             &Func1::new(move |aff|
                                                             {
                                                                 let killError =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Exception::Effect_Exception_error(),
                                                                                                      &&&string("[Aff] Child fiber outlived parent"));
                                                                 let killAll =
                                                                     &Func1::new(move
                                                                                     |err|
                                                                                     &Func1::new({
                                                                                                     let err
                                                                                                         =
                                                                                                         err.clone();
                                                                                                     move
                                                                                                         |sup|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeAff(),
                                                                                                                                          &&&Func1::new({
                                                                                                                                                            let sup
                                                                                                                                                                =
                                                                                                                                                                sup.clone();
                                                                                                                                                            move
                                                                                                                                                                |k|
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Effect_Aff::Effect_Aff__killAll()),
                                                                                                                                                                                                                                                                       &&&err),
                                                                                                                                                                                                                                    &&find(string("supervisor"),
                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&sup))),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(k,
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Either::Data_Either_applicativeEither()),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_Unit::Data_Unit_unit())))
                                                                                                                                                        }))
                                                                                                 }));
                                                                 let acquire =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                            &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__makeSupervisedFiber(),
                                                                                                                                                                            aff)),
                                                                                                      &&&Func1::new(move
                                                                                                                        |sup_1|
                                                                                                                        {
                                                                                                                            let fiber =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_Fiber(),
                                                                                                                                                                 &&&add(string("run"),
                                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__runFiber(),
                                                                                                                                                                                                          &&find(string("fiber"),
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(sup_1))),
                                                                                                                                                                        add(string("kill"),
                                                                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__killFiber(),
                                                                                                                                                                                                              &&find(string("fiber"),
                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(sup_1))),
                                                                                                                                                                            add(string("join"),
                                                                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__joinFiber(),
                                                                                                                                                                                                                  &&find(string("fiber"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(sup_1))),
                                                                                                                                                                                add(string("onComplete"),
                                                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__onCompleteFiber(),
                                                                                                                                                                                                                      &&find(string("fiber"),
                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(sup_1))),
                                                                                                                                                                                    add(string("isSuspended"),
                                                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff__isSuspendedFiber(),
                                                                                                                                                                                                                          &&find(string("fiber"),
                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(sup_1))),
                                                                                                                                                                                        empty::<string,
                                                                                                                                                                                                &dyn Any>()))))));
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                      &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                   &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                &&find(string("run"),
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(&&fiber)))),
                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                               let fiber
                                                                                                                                                                                   =
                                                                                                                                                                                   fiber.clone();
                                                                                                                                                                               let sup_1
                                                                                                                                                                                   =
                                                                                                                                                                                   sup_1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                       &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                    &&&add(string("supervisor"),
                                                                                                                                                                                                                           &find(string("supervisor"),
                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&sup_1)),
                                                                                                                                                                                                                           add(string("fiber"),
                                                                                                                                                                                                                               &&fiber,
                                                                                                                                                                                                                               empty::<string,
                                                                                                                                                                                                                                       &dyn Any>())))
                                                                                                                                                                           }))
                                                                                                                        }));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_generalBracket(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                                                              &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff()),
                                                                                                                                                                                                           &&&acquire)),
                                                                                                                                     &&&add(string("killed"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let killAll
                                                                                                                                                                 =
                                                                                                                                                                 killAll.clone();
                                                                                                                                                             move
                                                                                                                                                                 |err_1|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let err_1
                                                                                                                                                                                     =
                                                                                                                                                                                     err_1.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |sup_2|
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel::Control_Parallel_parSequence_(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Effect_Aff::Effect_Aff_parallelAff()),
                                                                                                                                                                                                                                                                                            &&&PureScript_Effect_Aff::Effect_Aff_applicativeParAff()),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                                                                                                                                                                                      &&&new_array(&[Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_killFiber(),
                                                                                                                                                                                                                                                                                                         &&&err_1),
                                                                                                                                                                                                                                                                      &&find(string("fiber"),
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(sup_2))),
                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&killAll,
                                                                                                                                                                                                                                                                                                         &&&err_1),
                                                                                                                                                                                                                                                                      sup_2)]))
                                                                                                                                                                             })
                                                                                                                                                         }),
                                                                                                                                            add(string("failed"),
                                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&killAll,
                                                                                                                                                                                                                     &&&killError)),
                                                                                                                                                add(string("completed"),
                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&killAll,
                                                                                                                                                                                                                         &&&killError)),
                                                                                                                                                    empty::<string,
                                                                                                                                                            &dyn Any>())))),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_joinFiber()),
                                                                                                                                     &&&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       find(string("fiber"),
                                                                                                                                                            Sharpurs_Prelude::unbox(v)))))
                                                             }))
    }
    pub fn Effect_Aff_suspendAff() -> &dyn Any {
        static Effect_Aff_suspendAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_suspendAff.get_or_init(||
                                              &Func1::new(move |aff|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                  &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff()),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeFiber(),
                                                                                                                                  aff))))
    }
    pub fn Effect_Aff_monadSTAff() -> &dyn Any {
        static Effect_Aff_monadSTAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadSTAff.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                               &&&add(string("liftST"),
                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                              &&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_monadSTEffect())),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_monadEffectAff())),
                                                                                      add(string("Monad0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Effect_Aff::Effect_Aff_monadAff()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Effect_Aff_monadThrowAff() -> &dyn Any {
        static Effect_Aff_monadThrowAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadThrowAff.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                  &&&add(string("throwError"),
                                                                                         &&PureScript_Effect_Aff::Effect_Aff__throwError(),
                                                                                         add(string("Monad0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Effect_Aff::Effect_Aff_monadAff()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Effect_Aff_monadErrorAff() -> &dyn Any {
        static Effect_Aff_monadErrorAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadErrorAff.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                  &&&add(string("catchError"),
                                                                                         &&PureScript_Effect_Aff::Effect_Aff__catchError(),
                                                                                         add(string("MonadThrow0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Effect_Aff::Effect_Aff_monadThrowAff()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Effect_Aff_attempt() -> &dyn Any {
        static Effect_Aff_attempt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_attempt.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_try(),
                                                                            &&&PureScript_Effect_Aff::Effect_Aff_monadErrorAff()))
    }
    pub fn Effect_Aff_runAff() -> &dyn Any {
        static Effect_Aff_runAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_runAff.get_or_init(||
                                          &Func1::new(move |k|
                                                          &Func1::new({
                                                                          let k
                                                                              =
                                                                              k.clone();
                                                                          move
                                                                              |aff|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                  &&&PureScript_Effect_Aff::Effect_Aff_launchAff()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_liftEffect()),
                                                                                                                                                                                                                        &&&k)),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_try(),
                                                                                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_monadErrorAff()),
                                                                                                                                                                                     aff)))
                                                                      })))
    }
    pub fn Effect_Aff_runAff_() -> &dyn Any {
        static Effect_Aff_runAff_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_runAff_.get_or_init(||
                                           &Func1::new(move |k|
                                                           &Func1::new({
                                                                           let k
                                                                               =
                                                                               k.clone();
                                                                           move
                                                                               |aff|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                   &&&PureScript_Effect_Aff::Effect_Aff_void()),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_runAff(),
                                                                                                                                                                                      &&&k),
                                                                                                                                                   aff))
                                                                       })))
    }
    pub fn Effect_Aff_runSuspendedAff() -> &dyn Any {
        static Effect_Aff_runSuspendedAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_runSuspendedAff.get_or_init(||
                                                   &Func1::new(move |k|
                                                                   &Func1::new({
                                                                                   let k
                                                                                       =
                                                                                       k.clone();
                                                                                   move
                                                                                       |aff|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_launchSuspendedAff()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                 &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                    &&&PureScript_Effect_Aff::Effect_Aff_liftEffect()),
                                                                                                                                                                                                                                 &&&k)),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_try(),
                                                                                                                                                                                                                                 &&&PureScript_Effect_Aff::Effect_Aff_monadErrorAff()),
                                                                                                                                                                                              aff)))
                                                                               })))
    }
    pub fn Effect_Aff_monadRecAff() -> &dyn Any {
        static Effect_Aff_monadRecAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monadRecAff.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                &&&add(string("tailRecM"),
                                                                                       &&Func1::new(move
                                                                                                        |k|
                                                                                                        {
                                                                                                            let go_2 =
                                                                                                                Func0::new({
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               move
                                                                                                                                   ||
                                                                                                                                   &Func1::new({
                                                                                                                                                   let go_tco
                                                                                                                                                       =
                                                                                                                                                       go_tco.clone();
                                                                                                                                                   move
                                                                                                                                                       |a|
                                                                                                                                                       go_tco(a.clone())
                                                                                                                                               })
                                                                                                                           });
                                                                                                            let go_1 =
                                                                                                                Lazy(go_2);
                                                                                                            let go_tco =
                                                                                                                Func1::new({
                                                                                                                               let k
                                                                                                                                   =
                                                                                                                                   k.clone();
                                                                                                                               move
                                                                                                                                   |a_1|
                                                                                                                                   fix1(&(move
                                                                                                                                              |go_tco,
                                                                                                                                               a_1|
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                     &&&PureScript_Effect_Aff::Effect_Aff_bindAff()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                     a_1)),
                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                 let go_tco
                                                                                                                                                                                                     =
                                                                                                                                                                                                     go_tco.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |res|
                                                                                                                                                                                                     {
                                                                                                                                                                                                         let matchValue:
                                                                                                                                                                                                                 LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(res);
                                                                                                                                                                                                         match matchValue.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_0_0)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             go_tco(matchValue_0_0),
                                                                                                                                                                                                             Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                                                                                                              &&matchValue_1_0),
                                                                                                                                                                                                         }
                                                                                                                                                                                                     }
                                                                                                                                                                                             }))),
                                                                                                                                        a_1.clone())
                                                                                                                           });
                                                                                                            let go =
                                                                                                                go_1.Value;
                                                                                                            &go
                                                                                                        }),
                                                                                       add(string("Monad0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Effect_Aff::Effect_Aff_monadAff()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Effect_Aff_monoidAff() -> &dyn Any {
        static Effect_Aff_monoidAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monoidAff.get_or_init(||
                                             &Func1::new(move |dictMonoid|
                                                             {
                                                                 let semigroupAff1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_semigroupAff(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                  &&&add(string("mempty"),
                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                              &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                              dictMonoid)),
                                                                                                         add(string("Semigroup0"),
                                                                                                             &&Func1::new({
                                                                                                                              let semigroupAff1
                                                                                                                                  =
                                                                                                                                  semigroupAff1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &semigroupAff1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Effect_Aff_nonCanceler() -> &dyn Any {
        static Effect_Aff_nonCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_nonCanceler.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_Canceler(),
                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                         &&&PureScript_Effect_Aff::Effect_Aff_applicativeAff()),
                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Effect_Aff_monoidCanceler() -> &dyn Any {
        static Effect_Aff_monoidCanceler: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_monoidCanceler.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                   &&&add(string("mempty"),
                                                                                          &&PureScript_Effect_Aff::Effect_Aff_nonCanceler(),
                                                                                          add(string("Semigroup0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Effect_Aff::Effect_Aff_semigroupCanceler()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Effect_Aff_never() -> &dyn Any {
        static Effect_Aff_never: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_never.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff::Effect_Aff_makeAff(),
                                                                          &&&Func1::new(move
                                                                                            |v|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                &&&PureScript_Effect_Aff::Effect_Aff_monoidCanceler())))))
    }
    pub fn Effect_Aff_apathize() -> &dyn Any {
        static Effect_Aff_apathize: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_apathize.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                &&&PureScript_Effect_Aff::Effect_Aff_attempt()),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                   &&&PureScript_Effect_Aff::Effect_Aff_functorAff()),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Effect_Aff_altParAff() -> &dyn Any {
        static Effect_Aff_altParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_altParAff.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                              &&&add(string("alt"),
                                                                                     &&PureScript_Effect_Aff::Effect_Aff__parAffAlt(),
                                                                                     add(string("Functor0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Effect_Aff::Effect_Aff_functorParAff()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Effect_Aff_altAff() -> &dyn Any {
        static Effect_Aff_altAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_altAff.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                           &&&add(string("alt"),
                                                                                  &&Func1::new(move
                                                                                                   |a1|
                                                                                                   &Func1::new({
                                                                                                                   let a1
                                                                                                                       =
                                                                                                                       a1.clone();
                                                                                                                   move
                                                                                                                       |a2|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                              &&&PureScript_Effect_Aff::Effect_Aff_monadErrorAff()),
                                                                                                                                                                                           &&&a1),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                           a2))
                                                                                                               })),
                                                                                  add(string("Functor0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Effect_Aff::Effect_Aff_functorAff()),
                                                                                      empty::<string,
                                                                                              &dyn Any>()))))
    }
    pub fn Effect_Aff_plusAff() -> &dyn Any {
        static Effect_Aff_plusAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_plusAff.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                            &&&add(string("empty"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                        &&&PureScript_Effect_Aff::Effect_Aff_monadThrowAff()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Exception::Effect_Exception_error(),
                                                                                                                                                        &&&string("Always fails"))),
                                                                                   add(string("Alt0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Effect_Aff::Effect_Aff_altAff()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Effect_Aff_plusParAff() -> &dyn Any {
        static Effect_Aff_plusParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_plusParAff.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                               &&&add(string("empty"),
                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_parallelAff()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                           &&&PureScript_Effect_Aff::Effect_Aff_plusAff())),
                                                                                      add(string("Alt0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Effect_Aff::Effect_Aff_altParAff()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Effect_Aff_alternativeParAff() -> &dyn Any {
        static Effect_Aff_alternativeParAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_alternativeParAff.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                      &&&add(string("Applicative0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Effect_Aff::Effect_Aff_applicativeParAff()),
                                                                                             add(string("Plus1"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused_1|
                                                                                                                  &PureScript_Effect_Aff::Effect_Aff_plusParAff()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
}
