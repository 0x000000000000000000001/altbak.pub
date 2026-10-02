// gopurs-aff's native-fiber / two-callback ABI, backed by the qualified Purust
// Aff runtime. The benchmark's PureScript and TAST inputs are shared verbatim.
fn gopurs_aff_util() -> AffUtil {
    use Purs_Data_Either::Either;
    AffUtil {
        is_left: aff_fn(|value| crate::Value::Bool(matches!(value.unwrap_class::<Arc<Either>>().as_ref(), Either::Left(_)))),
        from_left: aff_fn(|value| match value.unwrap_class::<Arc<Either>>().as_ref() {
            Either::Left(value) => value.clone(), _ => panic!("Expected Left"),
        }),
        from_right: aff_fn(|value| match value.unwrap_class::<Arc<Either>>().as_ref() {
            Either::Right(value) => value.clone(), _ => panic!("Expected Right"),
        }),
        left: aff_fn(|value| crate::Value::Class(Arc::new(Arc::new(Either::Left(value))))),
        right: aff_fn(|value| crate::Value::Class(Arc::new(Arc::new(Either::Right(value))))),
    }
}

fn gopurs_aff_field(fiber: &AffValue, name: &str) -> AffValue {
    let crate::Value::Record_a(record) = fiber.resolve() else { panic!("Expected native fiber record") };
    record.get_field(name).expect("native fiber method").clone()
}

pub fn Effect_Aff__makeFiberNative(aff: AffValue) -> AffValue {
    Effect_Aff__makeFiber().unwrap_func2()(crate::Value::Unit, aff)
}

pub fn Effect_Aff__makeSupervisedFiber(aff: AffValue) -> AffValue {
    gopurs_base_makeSupervisedFiber().unwrap_func2()(crate::Value::Unit, aff)
}

pub fn Effect_Aff__forkAffNative(aff: AffValue) -> AffValue {
    // The PS forkAff wrapper starts the suspended child with _runFiber.
    Effect_Aff__fork(false, aff)
}

pub fn Effect_Aff__runFiber(fiber: AffValue) -> AffValue {
    gopurs_aff_field(&fiber, "run")
}

pub fn Effect_Aff__isSuspendedFiber(fiber: AffValue) -> AffValue {
    gopurs_aff_field(&fiber, "isSuspended")
}

pub fn Effect_Aff__onCompleteFiber(fiber: AffValue, options: AffValue) -> AffValue {
    aff_call(&gopurs_aff_field(&fiber, "onComplete"), options)
}

pub fn Effect_Aff__joinFiber(
    fiber: AffValue,
    on_error: purust_core::Func1<AffValue, AffValue>,
    on_success: purust_core::Func1<AffValue, AffValue>,
) -> AffValue {
    let callback = aff_fn(move |result| match gopurs_aff_util().decode(result) {
        Ok(value) => on_success(value), Err(error) => on_error(error),
    });
    aff_call(&gopurs_aff_field(&fiber, "join"), callback)
}

pub fn Effect_Aff__killFiber(
    fiber: AffValue, error: AffValue,
    on_error: purust_core::Func1<AffValue, AffValue>,
    on_success: purust_core::Func1<(), AffValue>,
) -> AffValue {
    let callback = aff_fn(move |result| match gopurs_aff_util().decode(result) {
        Ok(_) => on_success(()), Err(error) => on_error(error),
    });
    gopurs_aff_field(&fiber, "kill").unwrap_func2()(error, callback)
}

pub fn Effect_Aff__makeAffImpl(
    build: purust_core::Func2<
        purust_core::Func1<AffValue, AffValue>,
        purust_core::Func1<AffValue, AffValue>, AffValue>,
) -> AffValue {
    gopurs_base_makeAff(purust_core::Func1::Shared(Arc::new(move |callback| {
        let failure = callback.clone();
        build(
            purust_core::Func1::Shared(Arc::new(move |error| failure(Arc::new(Purs_Data_Either::Either::Left(error))))),
            purust_core::Func1::Shared(Arc::new(move |value| callback(Arc::new(Purs_Data_Either::Either::Right(value))))),
        )
    })))
}
