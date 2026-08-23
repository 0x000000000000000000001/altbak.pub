import fs from 'fs';
let code = fs.readFileSync('output/purust_output/Purs_Record_Unsafe/src/lib.rs', 'utf8');

code = code.replace(/pub fn Record_Unsafe_unsafeSet\(mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType\) -> UnknownType \{ crate::Value::Thunk\(perceus_ptr::PerceusPtr::new\(crate::Thunk \{ ..Default::default\(\) \}\)\) \}/,
`pub fn Record_Unsafe_unsafeSet(mut a0: String, mut a1: UnknownType, mut a2: UnknownType) -> UnknownType { crate::Value::Thunk(perceus_ptr::PerceusPtr::new(crate::Thunk { ..Default::default() })) }`);

code = code.replace(/pub fn Record_Unsafe_unsafeGet\(mut a0: UnknownType, mut a1: UnknownType\) -> UnknownType \{ crate::Value::Thunk\(perceus_ptr::PerceusPtr::new\(crate::Thunk \{ ..Default::default\(\) \}\)\) \}/,
`pub fn Record_Unsafe_unsafeGet(mut a0: String, mut a1: UnknownType) -> UnknownType { crate::Value::Thunk(perceus_ptr::PerceusPtr::new(crate::Thunk { ..Default::default() })) }`);

code = code.replace(/pub fn Record_Unsafe_unsafeHas\(mut a0: UnknownType, mut a1: UnknownType\) -> UnknownType \{ crate::Value::Thunk\(perceus_ptr::PerceusPtr::new\(crate::Thunk \{ ..Default::default\(\) \}\)\) \}/,
`pub fn Record_Unsafe_unsafeHas(mut a0: String, mut a1: UnknownType) -> UnknownType { crate::Value::Thunk(perceus_ptr::PerceusPtr::new(crate::Thunk { ..Default::default() })) }`);

code = code.replace(/pub fn Record_Unsafe_unsafeDelete\(mut a0: UnknownType, mut a1: UnknownType\) -> UnknownType \{ crate::Value::Thunk\(perceus_ptr::PerceusPtr::new\(crate::Thunk \{ ..Default::default\(\) \}\)\) \}/,
`pub fn Record_Unsafe_unsafeDelete(mut a0: String, mut a1: UnknownType) -> UnknownType { crate::Value::Thunk(perceus_ptr::PerceusPtr::new(crate::Thunk { ..Default::default() })) }`);

fs.writeFileSync('output/purust_output/Purs_Record_Unsafe/src/lib.rs', code);
