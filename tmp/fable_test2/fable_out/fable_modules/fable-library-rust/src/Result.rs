pub mod Result_ {
    use super::*;
    use crate::module_3bd9ae6a::Native_::Func1;
    pub fn map<a: Clone + 'static, b: Clone + 'static, c: Clone + 'static>(
        mapping: Func1<a, b>,
        result: Result<a, c>,
    ) -> Result<b, c> {
        match &result {
            Err(result_1_0) => Err(result_1_0.clone()),
            Ok(result_0_0) => Ok(mapping(result_0_0.clone())),
        }
    }
    pub fn mapError<a: Clone + 'static, b: Clone + 'static, c: Clone + 'static>(
        mapping: Func1<a, b>,
        result: Result<c, a>,
    ) -> Result<c, b> {
        match &result {
            Err(result_1_0) => Err(mapping(result_1_0.clone())),
            Ok(result_0_0) => Ok(result_0_0.clone()),
        }
    }
    pub fn bind<a: Clone + 'static, b: Clone + 'static, c: Clone + 'static>(
        binder: Func1<a, Result<b, c>>,
        result: Result<a, c>,
    ) -> Result<b, c> {
        match &result {
            Err(result_1_0) => Err(result_1_0.clone()),
            Ok(result_0_0) => binder(result_0_0.clone()),
        }
    }
}
