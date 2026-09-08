use purust_core::*;
use std::rc::Rc;

pub fn Data_Array_length() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut xs: UnknownType| -> UnknownType {
            let arr = xs.unwrap_array();
            mk_int(arr.len() as i64)
        })))
}

pub fn Data_Array_unconsImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut empty: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut next: UnknownType| -> UnknownType {
                    let empty = empty.clone();
                    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut xs: UnknownType| -> UnknownType {
                            let mut empty = empty.clone();
                            let mut next = next.clone();
                            let arr = xs.unwrap_array();
                            if arr.is_empty() {
                                empty.unwrap_func1()(crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() })))
                            } else {
                                let head = arr[0].clone();
                                let tail = arr[1..].to_vec();
                                next.unwrap_func1()(head).unwrap_func1()(crate::mk_array(tail))
                            }
                        })))
                })))
        })))
}

pub fn Data_Array_indexImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut just: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut nothing: UnknownType| -> UnknownType {
                    let just = just.clone();
                    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut xs: UnknownType| -> UnknownType {
                            let just = just.clone();
                            let nothing = nothing.clone();
                            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut i: UnknownType| -> UnknownType {
                                    let mut just = just.clone();
                                    let nothing = nothing.clone();
                                    let arr = xs.unwrap_array();
                                    let idx_raw = i.unwrap_int();
                                    if idx_raw < 0 || idx_raw >= arr.len() as i64 {
                                        nothing
                                    } else {
                                        just.unwrap_func1()(arr[idx_raw as usize].clone())
                                    }
                                })))
                        })))
                })))
        })))
}

pub fn Data_Array_sliceImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut s: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut e: UnknownType| -> UnknownType {
                    let s = s.clone();
                    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut l: UnknownType| -> UnknownType {
                            let arr = l.unwrap_array();
                            let start = std::cmp::max(0, s.unwrap_int()) as usize;
                            let end = std::cmp::min(arr.len() as i64, e.unwrap_int()) as usize;
                            let start = std::cmp::min(start, arr.len());
                            let end = std::cmp::max(start, end);
                            let sliced = arr[start..end].to_vec();
                            crate::mk_array(sliced)
                        })))
                })))
        })))
}

pub fn Data_Array_findIndexImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut just: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut nothing: UnknownType| -> UnknownType {
                    let just = just.clone();
                    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut f: UnknownType| -> UnknownType {
                            let just = just.clone();
                            let nothing = nothing.clone();
                            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut xs: UnknownType| -> UnknownType {
                                    let mut just = just.clone();
                                    let nothing = nothing.clone();
                                    let mut f = f.clone();
                                    let arr = xs.unwrap_array();
                                    for (i, x) in arr.iter().enumerate() {
                                        let res = f.unwrap_func1()(x.clone());
                                        if res.unwrap_bool() {
                                            return just.unwrap_func1()(mk_int(i as i64));
                                        }
                                    }
                                    nothing
                                })))
                        })))
                })))
        })))
}

pub fn Data_Array_filterImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut f: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut xs: UnknownType| -> UnknownType {
                    let mut f = f.clone();
                    let arr = xs.unwrap_array();
                    let mut res = Vec::new();
                    for x in arr.iter() {
                        let ok = f.unwrap_func1()(x.clone());
                        if ok.unwrap_bool() {
                            res.push(x.clone());
                        }
                    }
                    crate::mk_array(res)
                })))
        })))
}

pub fn Data_Array_rangeImpl() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut start: UnknownType| -> UnknownType {
            crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut end: UnknownType| -> UnknownType {
                    let s = start.unwrap_int();
                    let e = end.unwrap_int();
                    let mut res = Vec::new();
                    if s <= e {
                        for i in s..=e {
                            res.push(crate::mk_int(i));
                        }
                    } else {
                        let mut i = s;
                        while i >= e {
                            res.push(crate::mk_int(i));
                            i -= 1;
                        }
                    }
                    crate::mk_array(res)
                })))
        })))
}
