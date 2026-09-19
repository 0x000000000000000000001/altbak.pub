pub mod Range_ {
    use super::*;
    use crate::module_3bd9ae6a::Native_::compare;
    use crate::module_3bd9ae6a::Native_::getZero;
    use crate::module_3bd9ae6a::Native_::Func1;
    use crate::module_3bd9ae6a::Native_::LrcPtr;
    use crate::module_52af85ec::Seq_::unfold;
    use crate::module_971078fd::Interfaces_::System::Collections::Generic::IEnumerable_1;
    use crate::module_eae1ac5e::String_::fromCharCode;
    use crate::module_eae1ac5e::String_::string;
    pub fn rangeNumeric<T: PartialOrd + Default + core::ops::Add<Output = T> + Clone + 'static>(
        start: T,
        step: T,
        stop: T,
    ) -> LrcPtr<dyn IEnumerable_1<T>> {
        let stepComparedWithZero: i32 = compare(step.clone(), getZero());
        if stepComparedWithZero == 0_i32 {
            panic!("{}", string("The step of a range cannot be zero"),);
        }
        unfold(
            Func1::new({
                let step = step.clone();
                let stepComparedWithZero = stepComparedWithZero.clone();
                let stop = stop.clone();
                move |x: T| {
                    let comparedWithLast: i32 = compare(x.clone(), stop.clone());
                    if if if stepComparedWithZero > 0_i32 {
                        comparedWithLast <= 0_i32
                    } else {
                        false
                    } {
                        true
                    } else {
                        if stepComparedWithZero < 0_i32 {
                            comparedWithLast >= 0_i32
                        } else {
                            false
                        }
                    } {
                        Some(LrcPtr::new((x.clone(), x + step.clone())))
                    } else {
                        None::<LrcPtr<(T, T)>>
                    }
                }
            }),
            start,
        )
    }
    pub fn rangeChar(start: char, stop: char) -> LrcPtr<dyn IEnumerable_1<char>> {
        let intStop: u32 = stop as u32;
        unfold(
            Func1::new({
                let intStop = intStop.clone();
                move |c: u32| {
                    if c <= intStop {
                        Some(LrcPtr::new((fromCharCode(c), c + 1_u32)))
                    } else {
                        None::<LrcPtr<(char, u32)>>
                    }
                }
            }),
            start as u32,
        )
    }
}
