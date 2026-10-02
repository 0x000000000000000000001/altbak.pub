// These three modules belong to the compiled dependency closure but are not
// exercised by Test.Main. Trap explicitly instead of accepting numeric/boolean
// default stubs as evidence of a working Aff program.
export const unusedFfi = {
  'Control.Extend': `pub fn Control_Extend_arrayExtend(_f: purust_core::Func1<UnknownType, UnknownType>, _xs: UnknownType) -> UnknownType {
    panic!("Unused benchmark FFI reached: Control.Extend.arrayExtend")
}
`,
  'Data.Int.Bits': ['and', 'or', 'xor', 'shl', 'shr', 'zshr', 'complement'].map(name =>
    `pub fn Data_Int_Bits_${name}(_a: i64${name === 'complement' ? '' : ', _b: i64'}) -> i64 {
    panic!("Unused benchmark FFI reached: Data.Int.Bits.${name}")
}`).join('\n') + '\n',
  'Performance.Minibench': `pub fn Performance_Minibench_gc() -> UnknownType {
    panic!("Unused benchmark FFI reached: Performance.Minibench.gc")
}
pub fn Performance_Minibench_timeNs() -> UnknownType {
    panic!("Unused benchmark FFI reached: Performance.Minibench.timeNs")
}
pub fn Performance_Minibench_toFixed(_value: f64) -> String {
    panic!("Unused benchmark FFI reached: Performance.Minibench.toFixed")
}
`,
};
