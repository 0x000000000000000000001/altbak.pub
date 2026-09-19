pub mod PureScript_Data_Array {
    use super::*;
    use fable_library_rust::Array_::equals;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::NativeArray_::count as count_1;
    use fable_library_rust::NativeArray_::new_array;
    use fable_library_rust::NativeArray_::new_empty;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_fcf3066b::PureScript_Control_Monad_ST_Internal;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_ba715d91::PureScript_Data_Array_NonEmpty_Internal;
    use crate::module_e9f7eca9::PureScript_Data_Array_ST_Iterator;
    use crate::module_acc76ca5::PureScript_Data_Array_ST;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_Array_FFI {
        use super::*;
        use fable_library_rust::Array_::copyTo;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IComparer_1;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::Lrc;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::interface_cast;
        use fable_library_rust::Native_::referenceEquals;
        use fable_library_rust::NativeArray_::add as add_1;
        use fable_library_rust::NativeArray_::new_copy;
        use fable_library_rust::NativeArray_::new_init;
        use fable_library_rust::System::Array;
        pub fn rangeImpl(startVal: &dyn Any, endVal: &dyn Any) -> &dyn Any {
            let start: i32 = startVal.clone();
            let endV: i32 = endVal.clone();
            let step: i32 = if start > endV { -1_i32 } else { 1_i32 };
            let size: i32 = (endV - start) * step + 1_i32;
            let result = new_init(&defaultOf(), size);
            let i: MutCell<i32> = MutCell::new(start);
            let n: MutCell<i32> = MutCell::new(0_i32);
            while i.get() != endV {
                result.get_mut()[n.get() as usize] = &i.get();
                n.set(n.get() + 1_i32);
                i.set(i.get() + step)
            }
            result.get_mut()[n.get() as usize] = &i.get();
            &result
        }
        pub fn replicateImpl(countVal: &dyn Any, value: &dyn Any)
         -> &dyn Any {
            let count: i32 = countVal.clone();
            if count < 1_i32 {
                &new_init(&defaultOf(), 0_i32)
            } else {
                let result = new_init(&defaultOf(), count);
                for i in 0_i32..=count - 1_i32 {
                    result.get_mut()[i as usize] = value.clone();
                }
                &result
            }
        }
        pub fn length(xs: &dyn Any) -> &dyn Any { &count_1(xs.clone()) }
        pub fn unconsImpl(empty: &dyn Any, next: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            let arr = xs.clone();
            if count_1(arr.clone()) == 0_i32 {
                Sharpurs_Prelude::sharpurs_apply(empty, &defaultOf())
            } else {
                let head = arr[0_i32].clone();
                let tail =
                    new_init(&defaultOf(), count_1(arr.clone()) - 1_i32);
                copyTo(arr.clone(), 1_i32, tail.clone(), 0_i32,
                       count_1(arr) - 1_i32);
                Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(next,
                                                                                   &head),
                                                 &&tail)
            }
        }
        pub fn indexImpl(just: &dyn Any, nothing: &dyn Any, xs: &dyn Any,
                         iVal: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let i: i32 = iVal.clone();
            if if i < 0_i32 { true } else { i >= count_1(arr.clone()) } {
                nothing.clone()
            } else { Sharpurs_Prelude::sharpurs_apply(just, &arr[i].clone()) }
        }
        pub fn _updateAt(just: &dyn Any, nothing: &dyn Any, iVal: &dyn Any,
                         a: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let i: i32 = iVal.clone();
            if if i < 0_i32 { true } else { i >= count_1(arr.clone()) } {
                nothing.clone()
            } else {
                let l1 = new_init(&defaultOf(), count_1(arr.clone()));
                copyTo(arr.clone(), 0_i32, l1.clone(), 0_i32, count_1(arr));
                l1.get_mut()[i as usize] = a.clone();
                Sharpurs_Prelude::sharpurs_apply(just, &&l1)
            }
        }
        pub fn _insertAt(just: &dyn Any, nothing: &dyn Any, iVal: &dyn Any,
                         a: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let i: i32 = iVal.clone();
            if if i < 0_i32 { true } else { i > count_1(arr.clone()) } {
                nothing.clone()
            } else {
                let l1 = new_init(&defaultOf(), count_1(arr.clone()) + 1_i32);
                copyTo(arr.clone(), 0_i32, l1.clone(), 0_i32, i);
                l1.get_mut()[i as usize] = a.clone();
                copyTo(arr.clone(), i, l1.clone(), i + 1_i32,
                       count_1(arr) - i);
                Sharpurs_Prelude::sharpurs_apply(just, &&l1)
            }
        }
        pub fn _deleteAt(just: &dyn Any, nothing: &dyn Any, iVal: &dyn Any,
                         xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let i: i32 = iVal.clone();
            if if i < 0_i32 { true } else { i >= count_1(arr.clone()) } {
                nothing.clone()
            } else {
                let l1 = new_init(&defaultOf(), count_1(arr.clone()) - 1_i32);
                copyTo(arr.clone(), 0_i32, l1.clone(), 0_i32, i);
                copyTo(arr.clone(), i + 1_i32, l1.clone(), i,
                       count_1(arr) - i - 1_i32);
                Sharpurs_Prelude::sharpurs_apply(just, &&l1)
            }
        }
        pub fn reverse(xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let l1 = new_init(&defaultOf(), count_1(arr.clone()));
            for i in 0_i32..=count_1(arr.clone()) - 1_i32 {
                l1.get_mut()[i as usize] =
                    arr[count_1(arr.clone()) - 1_i32 - i].clone();
            }
            &l1
        }
        pub fn concat(xss: &dyn Any) -> &dyn Any {
            let arrs = xss.clone();
            let totalLength: MutCell<i32> = MutCell::new(0_i32);
            for idx in 0_i32..=count_1(arrs.clone()) - 1_i32 {
                let xs = arrs[idx].clone();
                totalLength.set(totalLength.get() + count_1(xs))
            }
            {
                let result = new_init(&defaultOf(), totalLength.get());
                let current: MutCell<i32> = MutCell::new(0_i32);
                for idx_1 in 0_i32..=count_1(arrs.clone()) - 1_i32 {
                    let xsArr = arrs[idx_1].clone();
                    copyTo(xsArr.clone(), 0_i32, result.clone(),
                           current.get(), count_1(xsArr.clone()));
                    current.set(current.get() + count_1(xsArr))
                }
                &result
            }
        }
        pub fn filterImpl(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let res = new_empty::<&dyn Any>();
            for idx in 0_i32..=count_1(arr.clone()) - 1_i32 {
                let x = arr[idx].clone();
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                             &x))
                   {
                    add_1(res.clone(), x);
                }
            }
            &new_copy(res.clone())
        }
        pub fn sliceImpl(sVal: &dyn Any, eVal: &dyn Any, lVal: &dyn Any)
         -> &dyn Any {
            let s: MutCell<i32> = MutCell::new(sVal.clone());
            let e: MutCell<i32> = MutCell::new(eVal.clone());
            let l = lVal.clone();
            if s.get() < 0_i32 { s.set(count_1(l.clone()) + s.get()); }
            if e.get() < 0_i32 { e.set(count_1(l.clone()) + e.get()); }
            if s.get() < 0_i32 { s.set(0_i32); }
            if e.get() > count_1(l.clone()) { e.set(count_1(l.clone())); }
            if s.get() > e.get() { s.set(e.get()); }
            {
                let res = new_init(&defaultOf(), e.get() - s.get());
                copyTo(l, s.get(), res.clone(), 0_i32, e.get() - s.get());
                &res
            }
        }
        pub fn zipWithImpl(f: &dyn Any, xs: &dyn Any, ys: &dyn Any)
         -> &dyn Any {
            let arrX = xs.clone();
            let arrY = ys.clone();
            let length_1: i32 =
                count_1(arrX.clone()).min(count_1(arrY.clone()));
            let result = new_init(&defaultOf(), length_1);
            for i in 0_i32..=length_1 - 1_i32 {
                let step1 =
                    Sharpurs_Prelude::sharpurs_apply(f, &arrX[i].clone());
                result.get_mut()[i as usize] =
                    Sharpurs_Prelude::sharpurs_apply(&step1, &arrY[i].clone())
            }
            &result
        }
        pub fn unsafeIndexImpl(xs: &dyn Any, n: &dyn Any) -> &dyn Any {
            xs[n].clone()
        }
        pub fn sortByImpl(compare: &dyn Any, fromOrdering: &dyn Any,
                          xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            if count_1(arr.clone()) < 2_i32 {
                &arr
            } else {
                let comparer =
                    {
                        struct ObjectExpr {
                        }
                        impl <T: Clone + 'static> core::fmt::Display for
                         dyn IComparer_1<T> {
                            fn fmt(&self, f: &mut core::fmt::Formatter)
                             -> core::fmt::Result {
                                write!(f, "{}", core::any::type_name::<Self>())
                            }
                        }
                        impl <T: Clone + 'static> IComparer_1<T> for
                         ObjectExpr<T> {
                            fn Compare(&self, a: &dyn Any, b: &dyn Any)
                             -> i32 {
                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(fromOrdering,
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(compare,
                                                                                                                                                              a),
                                                                                                                            b)))
                            }
                        }
                        interface_cast!(LrcPtr::new(ObjectExpr::<&dyn Any>{}),
                                        Lrc<dyn IComparer_1<&dyn Any>>,)
                    };
                panic!("{}", 1_i32.get_Message(),)
            }
        }
        pub fn scanrImpl(f: &dyn Any, b: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let outArr = new_init(&defaultOf(), count_1(arr.clone()));
            let acc = b.clone();
            for i in (0_i32..=count_1(arr.clone()) - 1_i32).rev() {
                let step1 =
                    Sharpurs_Prelude::sharpurs_apply(f, &arr[i].clone());
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1, &acc));
                outArr.get_mut()[i as usize] = acc
            }
            &outArr
        }
        pub fn scanlImpl(f: &dyn Any, b: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let outArr = new_init(&defaultOf(), count_1(arr.clone()));
            let acc = b.clone();
            for i in 0_i32..=count_1(arr.clone()) - 1_i32 {
                let step1 = Sharpurs_Prelude::sharpurs_apply(f, &acc);
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1,
                                                         &arr[i].clone()));
                outArr.get_mut()[i as usize] = acc
            }
            &outArr
        }
        pub fn partitionImpl(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let yes = new_empty::<&dyn Any>();
            let no = new_empty::<&dyn Any>();
            for idx in 0_i32..=count_1(arr.clone()) - 1_i32 {
                let x = arr[idx].clone();
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                             &x))
                   {
                    add_1(yes.clone(), x)
                } else { add_1(no.clone(), x) }
            }
            {
                let res = empty_1::<string, &dyn Any>();
                let res_1 = add(string("yes"), &new_copy(yes.clone()), res);
                &add(string("no"), &new_copy(no.clone()), res_1)
            }
        }
        #[derive(Clone, Debug,)]
        enum ConsList {
            Cons(&dyn Any,
                 LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList>),
            EmptyList,
        }
        impl core::fmt::Display for
         PureScript_Data_Array::Data_Array_FFI::ConsList {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn fromFoldableImpl(foldr: &dyn Any, xsVal: &dyn Any)
         -> &dyn Any {
            let list:
                    LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList> =
                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(foldr,
                                                                                                                                              &&Func1::new(move
                                                                                                                                                               |head|
                                                                                                                                                               &Func1::new({
                                                                                                                                                                               let head
                                                                                                                                                                                   =
                                                                                                                                                                                   head.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |tail|
                                                                                                                                                                                   &LrcPtr::new(PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(head,
                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(tail)))
                                                                                                                                                                           }))),
                                                                                                            &&LrcPtr::new(PureScript_Data_Array::Data_Array_FFI::ConsList::EmptyList)),
                                                                          xsVal));
            fn countElements(l:
                                 LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList>,
                             acc: i32) -> i32 {
                let l:
                        MutCell<LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList>> =
                    MutCell::new(l.clone());
                let acc: MutCell<i32> = MutCell::new(acc);
                '_countElements:
                    loop  {
                        break '_countElements
                            (match l.get().as_ref() {
                                 PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(l_0_0,
                                                                                       l_0_1)
                                 => {
                                     let l_temp:
                                             LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList> =
                                         match l.get().as_ref() {
                                             PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(_,
                                                                                                   x)
                                             => x.clone(),
                                             _ => unreachable!(),
                                         };
                                     let acc_temp: i32 = acc.get() + 1_i32;
                                     l.set(l_temp);
                                     acc.set(acc_temp);
                                     continue '_countElements
                                 }
                                 _ => acc.get(),
                             }) ;
                    }
            }
            let size: i32 = countElements(list.clone(), 0_i32);
            let result = new_init(&defaultOf(), size);
            let fillArray =
                Func2::new({
                               let result = result.clone();
                               move
                                   |l_1:
                                        LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList>,
                                    i: i32|
                                   {
                                       let l_1:
                                               MutCell<LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList>> =
                                           MutCell::new(l_1.clone());
                                       let i: MutCell<i32> = MutCell::new(i);
                                       '_fillArray:
                                           loop  {
                                               break '_fillArray
                                                   (match l_1.get().as_ref() {
                                                        PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(l_1_0_0,
                                                                                                              l_1_0_1)
                                                        => {
                                                            result.get_mut()[i.get()
                                                                                 as
                                                                                 usize]
                                                                =
                                                                match l_1.get().as_ref()
                                                                    {
                                                                    PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(x,
                                                                                                                          _)
                                                                    =>
                                                                    x.clone(),
                                                                    _ =>
                                                                    unreachable!(),
                                                                };
                                                            {
                                                                let l_1_temp:
                                                                        LrcPtr<PureScript_Data_Array::Data_Array_FFI::ConsList> =
                                                                    match l_1.get().as_ref()
                                                                        {
                                                                        PureScript_Data_Array::Data_Array_FFI::ConsList::Cons(_,
                                                                                                                              x)
                                                                        =>
                                                                        x.clone(),
                                                                        _ =>
                                                                        unreachable!(),
                                                                    };
                                                                let i_temp:
                                                                        i32 =
                                                                    i.get() +
                                                                        1_i32;
                                                                l_1.set(l_1_temp);
                                                                i.set(i_temp);
                                                                continue
                                                                    '_fillArray

                                                            }
                                                        }
                                                        _ => (),
                                                    }) ;
                                           }
                                   }
                           });
            fillArray(list, 0_i32);
            &result
        }
        pub fn findMapImpl(nothing: &dyn Any, isJust: &dyn Any, f: &dyn Any,
                           xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let result = nothing.clone();
            let i: MutCell<i32> = MutCell::new(0_i32);
            while if i.get() < count_1(arr.clone()) {
                      referenceEquals(&result, &nothing)
                  } else { false } {
                let res =
                    Sharpurs_Prelude::sharpurs_apply(f,
                                                     &arr[i.get()].clone());
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(isJust,
                                                                             &res))
                   {
                    result.set(res);
                }
                i.set(i.get() + 1_i32)
            }
            result
        }
        pub fn findLastIndexImpl(just: &dyn Any, nothing: &dyn Any,
                                 f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let result = nothing.clone();
            let i: MutCell<i32> = MutCell::new(count_1(arr.clone()) - 1_i32);
            while if i.get() >= 0_i32 {
                      referenceEquals(&result, &nothing)
                  } else { false } {
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                             &arr[i.get()].clone()))
                   {
                    result.set(Sharpurs_Prelude::sharpurs_apply(just,
                                                                &&i.get()));
                }
                i.set(i.get() - 1_i32)
            }
            result
        }
        pub fn findIndexImpl(just: &dyn Any, nothing: &dyn Any, f: &dyn Any,
                             xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let result = nothing.clone();
            let i: MutCell<i32> = MutCell::new(0_i32);
            while if i.get() < count_1(arr.clone()) {
                      referenceEquals(&result, &nothing)
                  } else { false } {
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                             &arr[i.get()].clone()))
                   {
                    result.set(Sharpurs_Prelude::sharpurs_apply(just,
                                                                &&i.get()));
                }
                i.set(i.get() + 1_i32)
            }
            result
        }
        pub fn anyImpl(p: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let result: MutCell<bool> = MutCell::new(false);
            let i: MutCell<i32> = MutCell::new(0_i32);
            while if i.get() < count_1(arr.clone()) {
                      !result.get()
                  } else { false } {
                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(p,
                                                                             &arr[i.get()].clone()))
                   {
                    result.set(true);
                }
                i.set(i.get() + 1_i32)
            }
            &result.get()
        }
        pub fn allImpl(p: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let result: MutCell<bool> = MutCell::new(true);
            let i: MutCell<i32> = MutCell::new(0_i32);
            while if i.get() < count_1(arr.clone()) {
                      result.get()
                  } else { false } {
                if !Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(p,
                                                                              &arr[i.get()].clone()))
                   {
                    result.set(false);
                }
                i.set(i.get() + 1_i32)
            }
            &result.get()
        }
    }
    pub fn Data_Array__deleteAt() -> &dyn Any {
        static Data_Array__deleteAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array__deleteAt.get_or_init(||
                                             &Func1::new(move |just|
                                                             Func1::new({
                                                                            let just
                                                                                =
                                                                                just.clone();
                                                                            move
                                                                                |nothing|
                                                                                Func1::new({
                                                                                               let nothing
                                                                                                   =
                                                                                                   nothing.clone();
                                                                                               move
                                                                                                   |iVal|
                                                                                                   Func1::new({
                                                                                                                  let iVal
                                                                                                                      =
                                                                                                                      iVal.clone();
                                                                                                                  move
                                                                                                                      |xs|
                                                                                                                      PureScript_Data_Array::Data_Array_FFI::_deleteAt(&just,
                                                                                                                                                                       &nothing,
                                                                                                                                                                       &iVal,
                                                                                                                                                                       xs)
                                                                                                              })
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array__insertAt() -> &dyn Any {
        static Data_Array__insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array__insertAt.get_or_init(||
                                             &Func1::new(move |just|
                                                             Func1::new({
                                                                            let just
                                                                                =
                                                                                just.clone();
                                                                            move
                                                                                |nothing|
                                                                                Func1::new({
                                                                                               let nothing
                                                                                                   =
                                                                                                   nothing.clone();
                                                                                               move
                                                                                                   |iVal|
                                                                                                   Func1::new({
                                                                                                                  let iVal
                                                                                                                      =
                                                                                                                      iVal.clone();
                                                                                                                  move
                                                                                                                      |a|
                                                                                                                      Func1::new({
                                                                                                                                     let a
                                                                                                                                         =
                                                                                                                                         a.clone();
                                                                                                                                     move
                                                                                                                                         |xs|
                                                                                                                                         PureScript_Data_Array::Data_Array_FFI::_insertAt(&just,
                                                                                                                                                                                          &nothing,
                                                                                                                                                                                          &iVal,
                                                                                                                                                                                          &a,
                                                                                                                                                                                          xs)
                                                                                                                                 })
                                                                                                              })
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array__updateAt() -> &dyn Any {
        static Data_Array__updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array__updateAt.get_or_init(||
                                             &Func1::new(move |just|
                                                             Func1::new({
                                                                            let just
                                                                                =
                                                                                just.clone();
                                                                            move
                                                                                |nothing|
                                                                                Func1::new({
                                                                                               let nothing
                                                                                                   =
                                                                                                   nothing.clone();
                                                                                               move
                                                                                                   |iVal|
                                                                                                   Func1::new({
                                                                                                                  let iVal
                                                                                                                      =
                                                                                                                      iVal.clone();
                                                                                                                  move
                                                                                                                      |a|
                                                                                                                      Func1::new({
                                                                                                                                     let a
                                                                                                                                         =
                                                                                                                                         a.clone();
                                                                                                                                     move
                                                                                                                                         |xs|
                                                                                                                                         PureScript_Data_Array::Data_Array_FFI::_updateAt(&just,
                                                                                                                                                                                          &nothing,
                                                                                                                                                                                          &iVal,
                                                                                                                                                                                          &a,
                                                                                                                                                                                          xs)
                                                                                                                                 })
                                                                                                              })
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array_allImpl() -> &dyn Any {
        static Data_Array_allImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_allImpl.get_or_init(||
                                           &Func1::new(move |p|
                                                           Func1::new({
                                                                          let p
                                                                              =
                                                                              p.clone();
                                                                          move
                                                                              |xs|
                                                                              PureScript_Data_Array::Data_Array_FFI::allImpl(&p,
                                                                                                                             xs)
                                                                      })))
    }
    pub fn Data_Array_anyImpl() -> &dyn Any {
        static Data_Array_anyImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_anyImpl.get_or_init(||
                                           &Func1::new(move |p|
                                                           Func1::new({
                                                                          let p
                                                                              =
                                                                              p.clone();
                                                                          move
                                                                              |xs|
                                                                              PureScript_Data_Array::Data_Array_FFI::anyImpl(&p,
                                                                                                                             xs)
                                                                      })))
    }
    pub fn Data_Array_concat() -> &dyn Any {
        static Data_Array_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_concat.get_or_init(||
                                          &Func1::new(move |xss|
                                                          PureScript_Data_Array::Data_Array_FFI::concat(xss)))
    }
    pub fn Data_Array_filterImpl() -> &dyn Any {
        static Data_Array_filterImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_filterImpl.get_or_init(||
                                              &Func1::new(move |f|
                                                              Func1::new({
                                                                             let f
                                                                                 =
                                                                                 f.clone();
                                                                             move
                                                                                 |xs|
                                                                                 PureScript_Data_Array::Data_Array_FFI::filterImpl(&f,
                                                                                                                                   xs)
                                                                         })))
    }
    pub fn Data_Array_findIndexImpl() -> &dyn Any {
        static Data_Array_findIndexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findIndexImpl.get_or_init(||
                                                 &Func1::new(move |just|
                                                                 Func1::new({
                                                                                let just
                                                                                    =
                                                                                    just.clone();
                                                                                move
                                                                                    |nothing|
                                                                                    Func1::new({
                                                                                                   let nothing
                                                                                                       =
                                                                                                       nothing.clone();
                                                                                                   move
                                                                                                       |f|
                                                                                                       Func1::new({
                                                                                                                      let f
                                                                                                                          =
                                                                                                                          f.clone();
                                                                                                                      move
                                                                                                                          |xs|
                                                                                                                          PureScript_Data_Array::Data_Array_FFI::findIndexImpl(&just,
                                                                                                                                                                               &nothing,
                                                                                                                                                                               &f,
                                                                                                                                                                               xs)
                                                                                                                  })
                                                                                               })
                                                                            })))
    }
    pub fn Data_Array_findLastIndexImpl() -> &dyn Any {
        static Data_Array_findLastIndexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findLastIndexImpl.get_or_init(||
                                                     &Func1::new(move |just|
                                                                     Func1::new({
                                                                                    let just
                                                                                        =
                                                                                        just.clone();
                                                                                    move
                                                                                        |nothing|
                                                                                        Func1::new({
                                                                                                       let nothing
                                                                                                           =
                                                                                                           nothing.clone();
                                                                                                       move
                                                                                                           |f|
                                                                                                           Func1::new({
                                                                                                                          let f
                                                                                                                              =
                                                                                                                              f.clone();
                                                                                                                          move
                                                                                                                              |xs|
                                                                                                                              PureScript_Data_Array::Data_Array_FFI::findLastIndexImpl(&just,
                                                                                                                                                                                       &nothing,
                                                                                                                                                                                       &f,
                                                                                                                                                                                       xs)
                                                                                                                      })
                                                                                                   })
                                                                                })))
    }
    pub fn Data_Array_findMapImpl() -> &dyn Any {
        static Data_Array_findMapImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findMapImpl.get_or_init(||
                                               &Func1::new(move |nothing|
                                                               Func1::new({
                                                                              let nothing
                                                                                  =
                                                                                  nothing.clone();
                                                                              move
                                                                                  |isJust|
                                                                                  Func1::new({
                                                                                                 let isJust
                                                                                                     =
                                                                                                     isJust.clone();
                                                                                                 move
                                                                                                     |f|
                                                                                                     Func1::new({
                                                                                                                    let f
                                                                                                                        =
                                                                                                                        f.clone();
                                                                                                                    move
                                                                                                                        |xs|
                                                                                                                        PureScript_Data_Array::Data_Array_FFI::findMapImpl(&nothing,
                                                                                                                                                                           &isJust,
                                                                                                                                                                           &f,
                                                                                                                                                                           xs)
                                                                                                                })
                                                                                             })
                                                                          })))
    }
    pub fn Data_Array_fromFoldableImpl() -> &dyn Any {
        static Data_Array_fromFoldableImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_fromFoldableImpl.get_or_init(||
                                                    &Func1::new(move |foldr|
                                                                    Func1::new({
                                                                                   let foldr
                                                                                       =
                                                                                       foldr.clone();
                                                                                   move
                                                                                       |xsVal|
                                                                                       PureScript_Data_Array::Data_Array_FFI::fromFoldableImpl(&foldr,
                                                                                                                                               xsVal)
                                                                               })))
    }
    pub fn Data_Array_indexImpl() -> &dyn Any {
        static Data_Array_indexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_indexImpl.get_or_init(||
                                             &Func1::new(move |just|
                                                             Func1::new({
                                                                            let just
                                                                                =
                                                                                just.clone();
                                                                            move
                                                                                |nothing|
                                                                                Func1::new({
                                                                                               let nothing
                                                                                                   =
                                                                                                   nothing.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   Func1::new({
                                                                                                                  let xs
                                                                                                                      =
                                                                                                                      xs.clone();
                                                                                                                  move
                                                                                                                      |iVal|
                                                                                                                      PureScript_Data_Array::Data_Array_FFI::indexImpl(&just,
                                                                                                                                                                       &nothing,
                                                                                                                                                                       &xs,
                                                                                                                                                                       iVal)
                                                                                                              })
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array_length() -> &dyn Any {
        static Data_Array_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_length.get_or_init(||
                                          &Func1::new(move |xs|
                                                          PureScript_Data_Array::Data_Array_FFI::length(xs)))
    }
    pub fn Data_Array_partitionImpl() -> &dyn Any {
        static Data_Array_partitionImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_partitionImpl.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 Func1::new({
                                                                                let f
                                                                                    =
                                                                                    f.clone();
                                                                                move
                                                                                    |xs|
                                                                                    PureScript_Data_Array::Data_Array_FFI::partitionImpl(&f,
                                                                                                                                         xs)
                                                                            })))
    }
    pub fn Data_Array_rangeImpl() -> &dyn Any {
        static Data_Array_rangeImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_rangeImpl.get_or_init(||
                                             &Func1::new(move |startVal|
                                                             Func1::new({
                                                                            let startVal
                                                                                =
                                                                                startVal.clone();
                                                                            move
                                                                                |endVal|
                                                                                PureScript_Data_Array::Data_Array_FFI::rangeImpl(&startVal,
                                                                                                                                 endVal)
                                                                        })))
    }
    pub fn Data_Array_replicateImpl() -> &dyn Any {
        static Data_Array_replicateImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_replicateImpl.get_or_init(||
                                                 &Func1::new(move |countVal|
                                                                 Func1::new({
                                                                                let countVal
                                                                                    =
                                                                                    countVal.clone();
                                                                                move
                                                                                    |value|
                                                                                    PureScript_Data_Array::Data_Array_FFI::replicateImpl(&countVal,
                                                                                                                                         value)
                                                                            })))
    }
    pub fn Data_Array_reverse() -> &dyn Any {
        static Data_Array_reverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_reverse.get_or_init(||
                                           &Func1::new(move |xs|
                                                           PureScript_Data_Array::Data_Array_FFI::reverse(xs)))
    }
    pub fn Data_Array_scanlImpl() -> &dyn Any {
        static Data_Array_scanlImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_scanlImpl.get_or_init(||
                                             &Func1::new(move |f|
                                                             Func1::new({
                                                                            let f
                                                                                =
                                                                                f.clone();
                                                                            move
                                                                                |b|
                                                                                Func1::new({
                                                                                               let b
                                                                                                   =
                                                                                                   b.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   PureScript_Data_Array::Data_Array_FFI::scanlImpl(&f,
                                                                                                                                                    &b,
                                                                                                                                                    xs)
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array_scanrImpl() -> &dyn Any {
        static Data_Array_scanrImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_scanrImpl.get_or_init(||
                                             &Func1::new(move |f|
                                                             Func1::new({
                                                                            let f
                                                                                =
                                                                                f.clone();
                                                                            move
                                                                                |b|
                                                                                Func1::new({
                                                                                               let b
                                                                                                   =
                                                                                                   b.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   PureScript_Data_Array::Data_Array_FFI::scanrImpl(&f,
                                                                                                                                                    &b,
                                                                                                                                                    xs)
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array_sliceImpl() -> &dyn Any {
        static Data_Array_sliceImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_sliceImpl.get_or_init(||
                                             &Func1::new(move |sVal|
                                                             Func1::new({
                                                                            let sVal
                                                                                =
                                                                                sVal.clone();
                                                                            move
                                                                                |eVal|
                                                                                Func1::new({
                                                                                               let eVal
                                                                                                   =
                                                                                                   eVal.clone();
                                                                                               move
                                                                                                   |lVal|
                                                                                                   PureScript_Data_Array::Data_Array_FFI::sliceImpl(&sVal,
                                                                                                                                                    &eVal,
                                                                                                                                                    lVal)
                                                                                           })
                                                                        })))
    }
    pub fn Data_Array_sortByImpl() -> &dyn Any {
        static Data_Array_sortByImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_sortByImpl.get_or_init(||
                                              &Func1::new(move |compare|
                                                              Func1::new({
                                                                             let compare
                                                                                 =
                                                                                 compare.clone();
                                                                             move
                                                                                 |fromOrdering|
                                                                                 Func1::new({
                                                                                                let fromOrdering
                                                                                                    =
                                                                                                    fromOrdering.clone();
                                                                                                move
                                                                                                    |xs|
                                                                                                    PureScript_Data_Array::Data_Array_FFI::sortByImpl(&compare,
                                                                                                                                                      &fromOrdering,
                                                                                                                                                      xs)
                                                                                            })
                                                                         })))
    }
    pub fn Data_Array_unconsImpl() -> &dyn Any {
        static Data_Array_unconsImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unconsImpl.get_or_init(||
                                              &Func1::new(move |empty|
                                                              Func1::new({
                                                                             let empty
                                                                                 =
                                                                                 empty.clone();
                                                                             move
                                                                                 |next|
                                                                                 Func1::new({
                                                                                                let next
                                                                                                    =
                                                                                                    next.clone();
                                                                                                move
                                                                                                    |xs|
                                                                                                    PureScript_Data_Array::Data_Array_FFI::unconsImpl(&empty,
                                                                                                                                                      &next,
                                                                                                                                                      xs)
                                                                                            })
                                                                         })))
    }
    pub fn Data_Array_unsafeIndexImpl() -> &dyn Any {
        static Data_Array_unsafeIndexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unsafeIndexImpl.get_or_init(||
                                                   &Func1::new(move |xs|
                                                                   Func1::new({
                                                                                  let xs
                                                                                      =
                                                                                      xs.clone();
                                                                                  move
                                                                                      |n|
                                                                                      PureScript_Data_Array::Data_Array_FFI::unsafeIndexImpl(&xs,
                                                                                                                                             n)
                                                                              })))
    }
    pub fn Data_Array_zipWithImpl() -> &dyn Any {
        static Data_Array_zipWithImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_zipWithImpl.get_or_init(||
                                               &Func1::new(move |f|
                                                               Func1::new({
                                                                              let f
                                                                                  =
                                                                                  f.clone();
                                                                              move
                                                                                  |xs|
                                                                                  Func1::new({
                                                                                                 let xs
                                                                                                     =
                                                                                                     xs.clone();
                                                                                                 move
                                                                                                     |ys|
                                                                                                     PureScript_Data_Array::Data_Array_FFI::zipWithImpl(&f,
                                                                                                                                                        &xs,
                                                                                                                                                        ys)
                                                                                             })
                                                                          })))
    }
    pub fn Data_Array_intercalate1() -> &dyn Any {
        static Data_Array_intercalate1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_intercalate1.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_intercalate(),
                                                                                 &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Array_zero() -> &dyn Any {
        static Data_Array_zero: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_zero.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                         &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Data_Array_one() -> &dyn Any {
        static Data_Array_one: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_one.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                        &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Data_Array_void() -> &dyn Any {
        static Data_Array_void: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_void.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                         &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()))
    }
    pub fn Data_Array_pure() -> &dyn Any {
        static Data_Array_pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_pure.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                         &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()))
    }
    pub fn Data_Array_fromJust() -> &dyn Any {
        static Data_Array_fromJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_fromJust.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                             &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Array_foldMap1() -> &dyn Any {
        static Data_Array_foldMap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldMap1.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                             &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Array_fold1() -> &dyn Any {
        static Data_Array_fold1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_fold1.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_fold(),
                                                                          &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Array_not() -> &dyn Any {
        static Data_Array_not: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_not.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                        &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()))
    }
    pub fn Data_Array_zipWith() -> &dyn Any {
        static Data_Array_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_zipWith.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                            &&&PureScript_Data_Array::Data_Array_zipWithImpl()))
    }
    pub fn Data_Array_zipWithA() -> &dyn Any {
        static Data_Array_zipWithA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_zipWithA.get_or_init(||
                                            &Func1::new(move |dictApplicative|
                                                            &Func1::new({
                                                                            let dictApplicative
                                                                                =
                                                                                dictApplicative.clone();
                                                                            move
                                                                                |f|
                                                                                &Func1::new({
                                                                                                let f
                                                                                                    =
                                                                                                    f.clone();
                                                                                                move
                                                                                                    |xs|
                                                                                                    &Func1::new({
                                                                                                                    let xs
                                                                                                                        =
                                                                                                                        xs.clone();
                                                                                                                    move
                                                                                                                        |ys|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                               &&&PureScript_Data_Traversable::Data_Traversable_traversableArray()),
                                                                                                                                                                                            &&&dictApplicative),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_zipWith(),
                                                                                                                                                                                                                                                                  &&&f),
                                                                                                                                                                                                                               &&&xs),
                                                                                                                                                                                            ys))
                                                                                                                })
                                                                                            })
                                                                        })))
    }
    pub fn Data_Array_zip() -> &dyn Any {
        static Data_Array_zip: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_zip.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_zipWith(),
                                                                        &&&Func1::new(move
                                                                                          |usd__arg1|
                                                                                          Func1::new({
                                                                                                         let usd__arg1
                                                                                                             =
                                                                                                             usd__arg1.clone();
                                                                                                         move
                                                                                                             |usd__arg2|
                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                     usd__arg2.clone()))
                                                                                                     }))))
    }
    pub fn Data_Array_updateAtIndices() -> &dyn Any {
        static Data_Array_updateAtIndices: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_updateAtIndices.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictFoldable|
                                                                   &Func1::new({
                                                                                   let dictFoldable
                                                                                       =
                                                                                       dictFoldable.clone();
                                                                                   move
                                                                                       |us|
                                                                                       &Func1::new({
                                                                                                       let us
                                                                                                           =
                                                                                                           us.clone();
                                                                                                       move
                                                                                                           |xs|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_withArray(),
                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                    |res|
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_traverse_(),
                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                           &&&dictFoldable),
                                                                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                                                                          let res
                                                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                                                              res.clone();
                                                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                                                              |v|
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                  let matchValue:
                                                                                                                                                                                                                                                                                                                                          LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_poke(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                                                                   &&&res)
                                                                                                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                                                     &&&us))),
                                                                                                                                                                               xs))
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Array_updateAt() -> &dyn Any {
        static Data_Array_updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_updateAt.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn5(),
                                                                                                                                                   &&&PureScript_Data_Array::Data_Array__updateAt()),
                                                                                                                &&&Func1::new(move
                                                                                                                                  |usd__arg1|
                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_unsafeIndex() -> &dyn Any {
        static Data_Array_unsafeIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unsafeIndex.get_or_init(||
                                               &Func1::new(move |usd__unused|
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                                &&&PureScript_Data_Array::Data_Array_unsafeIndexImpl())))
    }
    pub fn Data_Array_uncons() -> &dyn Any {
        static Data_Array_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_uncons.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                 &&&PureScript_Data_Array::Data_Array_unconsImpl()),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                 &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                           &&&Func1::new(move
                                                                                             |x|
                                                                                             &Func1::new({
                                                                                                             let x
                                                                                                                 =
                                                                                                                 x.clone();
                                                                                                             move
                                                                                                                 |xs|
                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("head"),
                                                                                                                                                                             &&x,
                                                                                                                                                                             add(string("tail"),
                                                                                                                                                                                 xs.clone(),
                                                                                                                                                                                 empty_1::<string,
                                                                                                                                                                                           &dyn Any>()))))
                                                                                                         }))))
    }
    pub fn Data_Array_toUnfoldable() -> &dyn Any {
        static Data_Array_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_toUnfoldable.get_or_init(||
                                                &Func1::new(move
                                                                |dictUnfoldable|
                                                                &Func1::new({
                                                                                let dictUnfoldable
                                                                                    =
                                                                                    dictUnfoldable.clone();
                                                                                move
                                                                                    |xs|
                                                                                    {
                                                                                        let len =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                             xs);
                                                                                        let f =
                                                                                            &Func1::new({
                                                                                                            let len
                                                                                                                =
                                                                                                                len.clone();
                                                                                                            let xs
                                                                                                                =
                                                                                                                xs.clone();
                                                                                                            move
                                                                                                                |i|
                                                                                                                {
                                                                                                                    let matchValue =
                                                                                                                        Sharpurs_Prelude::unbox(i);
                                                                                                                    if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                       &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                 &&&len))
                                                                                                                       {
                                                                                                                        let i1_2 =
                                                                                                                            matchValue;
                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                                          let i1_2
                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                              i1_2.clone();
                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                              |usd__unused|
                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                  &&&xs),
                                                                                                                                                                                                                                                                                                                               &&&i1_2)
                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                           &&&i1_2),
                                                                                                                                                                                                                                                                        &&&1_i32)))))
                                                                                                                    } else {
                                                                                                                        if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                           {
                                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                        } else {
                                                                                                                            panic!("{}",
                                                                                                                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.fs"),
                                  Data1: 347_i32,
                                  Data2: 206_i32,}).get_Message(),)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                        });
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                               &&&dictUnfoldable),
                                                                                                                                                            &&&f),
                                                                                                                         &&&0_i32)
                                                                                    }
                                                                            })))
    }
    pub fn Data_Array_tail() -> &dyn Any {
        static Data_Array_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_tail.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                               &&&PureScript_Data_Array::Data_Array_unconsImpl()),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                               &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                         &&&Func1::new(move
                                                                                           |v|
                                                                                           &Func1::new(move
                                                                                                           |xs|
                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(xs.clone()))))))
    }
    pub fn Data_Array_sortBy() -> &dyn Any {
        static Data_Array_sortBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_sortBy.get_or_init(||
                                          &Func1::new(move |comp|
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                                 &&&PureScript_Data_Array::Data_Array_sortByImpl()),
                                                                                                                              comp),
                                                                                           &&&Func1::new(move
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let matchValue:
                                                                                                                         LrcPtr<Data_Ordering_Ordering> =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 match matchValue.as_ref()
                                                                                                                     {
                                                                                                                     Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                     =>
                                                                                                                     &0_i32,
                                                                                                                     Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                         &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                      &&&1_i32),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     &1_i32,
                                                                                                                 }
                                                                                                             }))))
    }
    pub fn Data_Array_sortWith() -> &dyn Any {
        static Data_Array_sortWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_sortWith.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            &Func1::new({
                                                                            let dictOrd
                                                                                =
                                                                                dictOrd.clone();
                                                                            move
                                                                                |f|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortBy(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_comparing(),
                                                                                                                                                                                       &&&dictOrd),
                                                                                                                                                    f))
                                                                        })))
    }
    pub fn Data_Array_sort() -> &dyn Any {
        static Data_Array_sort: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_sort.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        {
                                                            let compare =
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                 dictOrd);
                                                            &Func1::new({
                                                                            let compare
                                                                                =
                                                                                compare.clone();
                                                                            move
                                                                                |xs|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortBy(),
                                                                                                                                                    &&&compare),
                                                                                                                 xs)
                                                                        })
                                                        }))
    }
    pub fn Data_Array_snoc() -> &dyn Any {
        static Data_Array_snoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_snoc.get_or_init(||
                                        &Func1::new(move |xs|
                                                        &Func1::new({
                                                                        let xs
                                                                            =
                                                                            xs.clone();
                                                                        move
                                                                            |x|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_withArray(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                      x)),
                                                                                                                                                &&&xs))
                                                                    })))
    }
    pub fn Data_Array_slice() -> &dyn Any {
        static Data_Array_slice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_slice.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                          &&&PureScript_Data_Array::Data_Array_sliceImpl()))
    }
    pub fn Data_Array_splitAt() -> &dyn Any {
        static Data_Array_splitAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_splitAt.get_or_init(||
                                           &Func1::new(move |v|
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
                                                                                   if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                &&&0_i32))
                                                                                      {
                                                                                       &add(string("before"),
                                                                                            &&new_empty::<&dyn Any>(),
                                                                                            add(string("after"),
                                                                                                &&matchValue_1,
                                                                                                empty_1::<string,
                                                                                                          &dyn Any>()))
                                                                                   } else {
                                                                                       let xs_2 =
                                                                                           matchValue_1;
                                                                                       let i_2 =
                                                                                           matchValue;
                                                                                       &add(string("before"),
                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                    &&&0_i32),
                                                                                                                                                                 &&&i_2),
                                                                                                                              &&&xs_2),
                                                                                            add(string("after"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                        &&&i_2),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                        &&&xs_2)),
                                                                                                                                  &&&xs_2),
                                                                                                empty_1::<string,
                                                                                                          &dyn Any>()))
                                                                                   }
                                                                               }
                                                                       })))
    }
    pub fn Data_Array_take() -> &dyn Any {
        static Data_Array_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_take.get_or_init(||
                                        &Func1::new(move |n|
                                                        &Func1::new({
                                                                        let n
                                                                            =
                                                                            n.clone();
                                                                        move
                                                                            |xs|
                                                                            {
                                                                                let matchValue =
                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                    &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                 &&&n),
                                                                                                                                              &&&1_i32));
                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                 &matchValue)
                                                                                    {
                                                                                    0_i32
                                                                                    =>
                                                                                    &new_empty::<&dyn Any>(),
                                                                                    _
                                                                                    =>
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                           &&&0_i32),
                                                                                                                                                        &&&n),
                                                                                                                     xs),
                                                                                }
                                                                            }
                                                                    })))
    }
    pub fn Data_Array_singleton() -> &dyn Any {
        static Data_Array_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_singleton.get_or_init(||
                                             &Func1::new(move |a|
                                                             &new_array(&[a.clone()])))
    }
    pub fn Data_Array_scanr() -> &dyn Any {
        static Data_Array_scanr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_scanr.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                          &&&PureScript_Data_Array::Data_Array_scanrImpl()))
    }
    pub fn Data_Array_scanl() -> &dyn Any {
        static Data_Array_scanl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_scanl.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                          &&&PureScript_Data_Array::Data_Array_scanlImpl()))
    }
    pub fn Data_Array_replicate() -> &dyn Any {
        static Data_Array_replicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_replicate.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                              &&&PureScript_Data_Array::Data_Array_replicateImpl()))
    }
    pub fn Data_Array_range() -> &dyn Any {
        static Data_Array_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_range.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                          &&&PureScript_Data_Array::Data_Array_rangeImpl()))
    }
    pub fn Data_Array_partition() -> &dyn Any {
        static Data_Array_partition: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_partition.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                              &&&PureScript_Data_Array::Data_Array_partitionImpl()))
    }
    pub fn Data_Array_null() -> &dyn Any {
        static Data_Array_null: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_null.get_or_init(||
                                        &Func1::new(move |xs|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                               &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                               xs)),
                                                                                         &&&0_i32)))
    }
    pub fn Data_Array_modifyAtIndices() -> &dyn Any {
        static Data_Array_modifyAtIndices: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_modifyAtIndices.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictFoldable|
                                                                   &Func1::new({
                                                                                   let dictFoldable
                                                                                       =
                                                                                       dictFoldable.clone();
                                                                                   move
                                                                                       |is|
                                                                                       &Func1::new({
                                                                                                       let is
                                                                                                           =
                                                                                                           is.clone();
                                                                                                       move
                                                                                                           |f|
                                                                                                           &Func1::new({
                                                                                                                           let f
                                                                                                                               =
                                                                                                                               f.clone();
                                                                                                                           move
                                                                                                                               |xs|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_withArray(),
                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                        |res|
                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_traverse_(),
                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                              let res
                                                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                                                  res.clone();
                                                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                                                  |i|
                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_modify(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                         i),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                                                                                                                                                                                   &&&res)
                                                                                                                                                                                                                                                                                                                                          })),
                                                                                                                                                                                                                                                                                         &&&is))),
                                                                                                                                                                                                   xs))
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Array_mapWithIndex() -> &dyn Any {
        static Data_Array_mapWithIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_mapWithIndex.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                 &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexArray()))
    }
    pub fn Data_Array_intersperse() -> &dyn Any {
        static Data_Array_intersperse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_intersperse.get_or_init(||
                                               &Func1::new(move |a|
                                                               &Func1::new({
                                                                               let a
                                                                                   =
                                                                                   a.clone();
                                                                               move
                                                                                   |arr|
                                                                                   {
                                                                                       let matchValue =
                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                      arr));
                                                                                       if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                          &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                    &&&2_i32))
                                                                                          {
                                                                                           arr.clone()
                                                                                       } else {
                                                                                           if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                              {
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_run(),
                                                                                                                                &&{
                                                                                                                                      let unsafeGetElem =
                                                                                                                                          &Func1::new({
                                                                                                                                                          let arr
                                                                                                                                                              =
                                                                                                                                                              arr.clone();
                                                                                                                                                          move
                                                                                                                                                              |idx|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                 let idx
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     idx.clone();
                                                                                                                                                                                                                 move
                                                                                                                                                                                                                     |usd__unused|
                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                         &&&arr),
                                                                                                                                                                                                                                                      &&&idx)
                                                                                                                                                                                                             }))
                                                                                                                                                      });
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                             &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                          &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                         let unsafeGetElem
                                                                                                                                                                                             =
                                                                                                                                                                                             unsafeGetElem.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |out|
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&unsafeGetElem,
                                                                                                                                                                                                                                                                                                                                                                          &&&0_i32)),
                                                                                                                                                                                                                                                                                                    out)),
                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                let out
                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                    out.clone();
                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_for(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&1_i32),
                                                                                                                                                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                             |idx_1|
                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&out)),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                let idx_1
                                                                                                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                                                                                                    idx_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                                                                                                    |usd__unused_2|
                                                                                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&unsafeGetElem,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&idx_1)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&out))
                                                                                                                                                                                                                                                                                                                                                                                                                            }))))),
                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                       |usd__unused_3|
                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                        &&&out)))
                                                                                                                                                                                                                                            }))
                                                                                                                                                                                     }))
                                                                                                                                  })
                                                                                           } else {
                                                                                               panic!("{}",
                                                                                                      LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.fs"),
                                  Data1: 383_i32,
                                  Data2: 162_i32,}).get_Message(),)
                                                                                           }
                                                                                       }
                                                                                   }
                                                                           })))
    }
    pub fn Data_Array_intercalate() -> &dyn Any {
        static Data_Array_intercalate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_intercalate.get_or_init(||
                                               &Func1::new(move |dictMonoid|
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_intercalate1(),
                                                                                                dictMonoid)))
    }
    pub fn Data_Array_insertAt() -> &dyn Any {
        static Data_Array_insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_insertAt.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn5(),
                                                                                                                                                   &&&PureScript_Data_Array::Data_Array__insertAt()),
                                                                                                                &&&Func1::new(move
                                                                                                                                  |usd__arg1|
                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_init() -> &dyn Any {
        static Data_Array_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_init.get_or_init(||
                                        &Func1::new(move |xs|
                                                        {
                                                            let matchValue =
                                                                Sharpurs_Prelude::unbox(xs);
                                                            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_null(),
                                                                                                                         &&&matchValue))
                                                               {
                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                            } else {
                                                                if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                   {
                                                                    let xs1_3 =
                                                                        matchValue;
                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                                                  &&&PureScript_Data_Array::Data_Array_zero()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                                                                                                        &&&xs1_3)),
                                                                                                                                                                                                                                  &&&PureScript_Data_Array::Data_Array_one())),
                                                                                                                                                            &&&xs1_3)))
                                                                } else {
                                                                    panic!("{}",
                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.fs"),
                                  Data1: 389_i32,
                                  Data2: 53_i32,}).get_Message(),)
                                                                }
                                                            }
                                                        }))
    }
    pub fn Data_Array_index() -> &dyn Any {
        static Data_Array_index: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_index.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                &&&PureScript_Data_Array::Data_Array_indexImpl()),
                                                                                                             &&&Func1::new(move
                                                                                                                               |usd__arg1|
                                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                          &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_last() -> &dyn Any {
        static Data_Array_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_last.get_or_init(||
                                        &Func1::new(move |xs|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                            xs),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                  &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                  xs)),
                                                                                                                            &&&1_i32))))
    }
    pub fn Data_Array_unsnoc() -> &dyn Any {
        static Data_Array_unsnoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unsnoc.get_or_init(||
                                          &Func1::new(move |xs|
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                 &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                      |v|
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let v
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          v.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                                                          &add(string("init"),
                                                                                                                                                                                                                                               &&v,
                                                                                                                                                                                                                                               add(string("last"),
                                                                                                                                                                                                                                                   v1.clone(),
                                                                                                                                                                                                                                                   empty_1::<string,
                                                                                                                                                                                                                                                             &dyn Any>()))
                                                                                                                                                                                                                                  }))),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_init(),
                                                                                                                                                                                                    xs))),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_last(),
                                                                                                                              xs))))
    }
    pub fn Data_Array_modifyAt() -> &dyn Any {
        static Data_Array_modifyAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_modifyAt.get_or_init(||
                                            &Func1::new(move |i|
                                                            &Func1::new({
                                                                            let i
                                                                                =
                                                                                i.clone();
                                                                            move
                                                                                |f|
                                                                                &Func1::new({
                                                                                                let f
                                                                                                    =
                                                                                                    f.clone();
                                                                                                move
                                                                                                    |xs|
                                                                                                    {
                                                                                                        let go =
                                                                                                            &Func1::new({
                                                                                                                            let xs
                                                                                                                                =
                                                                                                                                xs.clone();
                                                                                                                            move
                                                                                                                                |x|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_updateAt(),
                                                                                                                                                                                                                                       &&&i),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                       x)),
                                                                                                                                                                 &&&xs)
                                                                                                                        });
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                               &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                            &&&go),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                                                                                                               xs),
                                                                                                                                                                            &&&i))
                                                                                                    }
                                                                                            })
                                                                        })))
    }
    pub fn Data_Array_unzip() -> &dyn Any {
        static Data_Array_unzip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unzip.get_or_init(||
                                         &Func1::new(move |xs|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                   &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                             &&&Func1::new({
                                                                                                                                               let xs
                                                                                                                                                   =
                                                                                                                                                   xs.clone();
                                                                                                                                               move
                                                                                                                                                   |fsts|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                          &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                       &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                      let fsts
                                                                                                                                                                                                          =
                                                                                                                                                                                                          fsts.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |snds|
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_iterator(),
                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                   |v|
                                                                                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&xs),
                                                                                                                                                                                                                                                                                                                                                                    v)))),
                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                             let snds
                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                 snds.clone();
                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                 |iter|
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_iterate(),
                                                                                                                                                                                                                                                                                                                                                                                                           iter),
                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                          |v_1|
                                                                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                                                                              let matchValue:
                                                                                                                                                                                                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Array::Data_Array_void()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&fsts))),
                                                                                                                                                                                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Array::Data_Array_void()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&snds))))
                                                                                                                                                                                                                                                                                                                                                                                          }))),
                                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&fsts)),
                                                                                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                       |fsts_prime|
                                                                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&snds)),
                                                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                          let fsts_prime
                                                                                                                                                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                                                                                                                                                              fsts_prime.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                                                                                                                                                              |snds_prime|
                                                                                                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Array::Data_Array_pure()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&fsts_prime,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         snds_prime.clone())))
                                                                                                                                                                                                                                                                                                                                                                                                                      }))))))
                                                                                                                                                                                                                                                         }))
                                                                                                                                                                                                  }))
                                                                                                                                           })))))
    }
    pub fn Data_Array_head() -> &dyn Any {
        static Data_Array_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_head.get_or_init(||
                                        &Func1::new(move |xs|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                            xs),
                                                                                         &&&0_i32)))
    }
    pub fn Data_Array_nubBy() -> &dyn Any {
        static Data_Array_nubBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_nubBy.get_or_init(||
                                         &Func1::new(move |comp|
                                                         &Func1::new({
                                                                         let comp
                                                                             =
                                                                             comp.clone();
                                                                         move
                                                                             |xs|
                                                                             {
                                                                                 let indexedAndSorted =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortBy(),
                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                           |x|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let x
                                                                                                                                                                                               =
                                                                                                                                                                                               x.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |y|
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&comp,
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                                                                                                                                                                                                                                      &&&x)),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                                                                                                                                                                                                   y))
                                                                                                                                                                                       }))),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_mapWithIndex(),
                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                              |usd__arg1|
                                                                                                                                                                                                              Func1::new({
                                                                                                                                                                                                                             let usd__arg1
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 usd__arg1.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |usd__arg2|
                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                         usd__arg2.clone()))
                                                                                                                                                                                                                         }))),
                                                                                                                                                         xs));
                                                                                 let matchValue:
                                                                                         LrcPtr<Data_Maybe_Maybe> =
                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_head(),
                                                                                                                                                &&&indexedAndSorted));
                                                                                 match matchValue.as_ref()
                                                                                     {
                                                                                     Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                     =>
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               &&&PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                                                                                                                            &&&PureScript_Data_Tuple::Data_Tuple_snd())),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortWith(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                               &&&PureScript_Data_Tuple::Data_Tuple_fst())),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Array_ST::Data_Array_ST_unsafeThaw()),
                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_singleton(),
                                                                                                                                                                                                                                                                                                                                        &&matchValue_1_0))),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let indexedAndSorted
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     indexedAndSorted.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |result|
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_foreach(),
                                                                                                                                                                                                                                                                                                                                                                                               &&&indexedAndSorted),
                                                                                                                                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                              let result
                                                                                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                                                                                  result.clone();
                                                                                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                                                                                  |v1|
                                                                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                                                                      let pair:
                                                                                                                                                                                                                                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |usd__unused|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Array::Data_Array_fromJust()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Array::Data_Array_last()))))),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&result))),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                         let pair
                                                                                                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                                                                                                             pair.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                                                                                                             |lst|
                                                                                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_when(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&comp,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                lst),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&match pair.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)))),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Array::Data_Array_void()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&pair),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&result)))
                                                                                                                                                                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                                                                                                                                                          }))),
                                                                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                                                                        let result
                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                            result.clone();
                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                            |usd__unused_1|
                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                             &&&result)
                                                                                                                                                                                                                                                                                                    }))
                                                                                                                                                                                                                                             }))))),
                                                                                     _
                                                                                     =>
                                                                                     &new_empty::<&dyn Any>(),
                                                                                 }
                                                                             }
                                                                     })))
    }
    pub fn Data_Array_nub() -> &dyn Any {
        static Data_Array_nub: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_nub.get_or_init(||
                                       &Func1::new(move |dictOrd|
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubBy(),
                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                           dictOrd))))
    }
    pub fn Data_Array_groupBy() -> &dyn Any {
        static Data_Array_groupBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_groupBy.get_or_init(||
                                           &Func1::new(move |op|
                                                           &Func1::new({
                                                                           let op
                                                                               =
                                                                               op.clone();
                                                                           move
                                                                               |xs|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                         &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                      &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                     let xs
                                                                                                                                                                         =
                                                                                                                                                                         xs.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |result|
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_iterator(),
                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                  |v|
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                                                                                                                                                                                                                                                                      &&&xs),
                                                                                                                                                                                                                                                                                                                                   v)))),
                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                            let result
                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                result.clone();
                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                |iter|
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_iterate(),
                                                                                                                                                                                                                                                                                                                                                                          iter),
                                                                                                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                         let iter
                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                             iter.clone();
                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                             |x|
                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                   let x
                                                                                                                                                                                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                                                                                                                                                                                       x.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                                                                                                                                                                                                       |sub|
                                                                                                                                                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&x),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              sub)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          let sub
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              sub.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |usd__unused|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_pushWhile(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&op,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&x)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&iter),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&sub)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&sub)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |grp|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_NonEmptyArray(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           grp)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&result)))))
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                                                                                                                                                                                                                                                                               })))
                                                                                                                                                                                                                                                                                                                                                     }))),
                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                   |usd__unused_2|
                                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                    &&&result)))
                                                                                                                                                                                                                        }))
                                                                                                                                                                 })))
                                                                       })))
    }
    pub fn Data_Array_groupAllBy() -> &dyn Any {
        static Data_Array_groupAllBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_groupAllBy.get_or_init(||
                                              &Func1::new(move |cmp|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_groupBy(),
                                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                                       let cmp
                                                                                                                                                                                           =
                                                                                                                                                                                           cmp.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |x|
                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                           let x
                                                                                                                                                                                                               =
                                                                                                                                                                                                               x.clone();
                                                                                                                                                                                                           move
                                                                                                                                                                                                               |y|
                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                         &&&x),
                                                                                                                                                                                                                                                                                                                      y)),
                                                                                                                                                                                                                                                &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))
                                                                                                                                                                                                       })
                                                                                                                                                                                   }))),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortBy(),
                                                                                                                                  cmp))))
    }
    pub fn Data_Array_groupAll() -> &dyn Any {
        static Data_Array_groupAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_groupAll.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_groupAllBy(),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                dictOrd))))
    }
    pub fn Data_Array_group() -> &dyn Any {
        static Data_Array_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_group.get_or_init(||
                                         &Func1::new(move |dictEq|
                                                         {
                                                             let eq =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                  dictEq);
                                                             &Func1::new({
                                                                             let eq
                                                                                 =
                                                                                 eq.clone();
                                                                             move
                                                                                 |xs|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_groupBy(),
                                                                                                                                                     &&&eq),
                                                                                                                  xs)
                                                                         })
                                                         }))
    }
    pub fn Data_Array_fromFoldable() -> &dyn Any {
        static Data_Array_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_fromFoldable.get_or_init(||
                                                &Func1::new(move
                                                                |dictFoldable|
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                                                                    &&&PureScript_Data_Array::Data_Array_fromFoldableImpl()),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                    dictFoldable))))
    }
    pub fn Data_Array_foldr() -> &dyn Any {
        static Data_Array_foldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldr.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                          &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Array_foldl() -> &dyn Any {
        static Data_Array_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldl.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                          &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Array_transpose() -> &dyn Any {
        static Data_Array_transpose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_transpose.get_or_init(||
                                             &Func1::new(move |xs|
                                                             {
                                                                 let buildNext =
                                                                     &Func1::new({
                                                                                     let xs
                                                                                         =
                                                                                         xs.clone();
                                                                                     move
                                                                                         |idx|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_applyFlipped(),
                                                                                                                                                             &&&xs),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                   &&&PureScript_Data_Array::Data_Array_foldl()),
                                                                                                                                                                                                &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                               let idx
                                                                                                                                                                                   =
                                                                                                                                                                                   idx.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |acc|
                                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                                   let acc
                                                                                                                                                                                                       =
                                                                                                                                                                                                       acc.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |nextArr|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                 &&&acc),
                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                |el|
                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&new_array(&[el.clone()])),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Array::Data_Array_snoc()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          el)),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&acc))))),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                                                                                                                                                                                                              nextArr),
                                                                                                                                                                                                                                                                           &&&idx))
                                                                                                                                                                                               })
                                                                                                                                                                           })))
                                                                                 });
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
                                                                                                                |idx_1|
                                                                                                                Func1::new({
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               let idx_1
                                                                                                                                   =
                                                                                                                                   idx_1.clone();
                                                                                                                               move
                                                                                                                                   |allArrays|
                                                                                                                                   go_tco(idx_1)(allArrays.clone())
                                                                                                                           })
                                                                                                        })
                                                                                    });
                                                                     let go_1 =
                                                                         Lazy(go_2);
                                                                     let go_tco =
                                                                         Func1::new({
                                                                                        let buildNext
                                                                                            =
                                                                                            buildNext.clone();
                                                                                        move
                                                                                            |idx_2|
                                                                                            fix1(&(move
                                                                                                       |go_tco,
                                                                                                        idx_2|
                                                                                                       Func1::new({
                                                                                                                      let go_tco
                                                                                                                          =
                                                                                                                          go_tco.clone();
                                                                                                                      let idx_2
                                                                                                                          =
                                                                                                                          idx_2.clone();
                                                                                                                      move
                                                                                                                          |allArrays_1|
                                                                                                                          {
                                                                                                                              let matchValue:
                                                                                                                                      LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&buildNext,
                                                                                                                                                                                             &&&idx_2));
                                                                                                                              match matchValue.as_ref()
                                                                                                                                  {
                                                                                                                                  Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                  =>
                                                                                                                                  go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                             &&&idx_2),
                                                                                                                                                                          &&&1_i32))(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_snoc(),
                                                                                                                                                                                                                                                         allArrays_1),
                                                                                                                                                                                                                      &&matchValue_1_0)),
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  allArrays_1.clone(),
                                                                                                                              }
                                                                                                                          }
                                                                                                                  })),
                                                                                                 idx_2.clone())
                                                                                    });
                                                                     let go =
                                                                         go_1.Value;
                                                                     go_tco(&0_i32)(&new_empty::<&dyn Any>())
                                                                 }
                                                             }))
    }
    pub fn Data_Array_foldRecM() -> &dyn Any {
        static Data_Array_foldRecM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldRecM.get_or_init(||
                                            &Func1::new(move |dictMonadRec|
                                                            {
                                                                let Monad0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                            Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                let Applicative0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                            Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                let Bind1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                            Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
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
                                                                                                        |b|
                                                                                                        &Func1::new({
                                                                                                                        let b
                                                                                                                            =
                                                                                                                            b.clone();
                                                                                                                        move
                                                                                                                            |array|
                                                                                                                            {
                                                                                                                                let go =
                                                                                                                                    &Func1::new({
                                                                                                                                                    let array
                                                                                                                                                        =
                                                                                                                                                        array.clone();
                                                                                                                                                    move
                                                                                                                                                        |res|
                                                                                                                                                        &Func1::new({
                                                                                                                                                                        let res
                                                                                                                                                                            =
                                                                                                                                                                            res.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |i|
                                                                                                                                                                            {
                                                                                                                                                                                let matchValue =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&res);
                                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(i);
                                                                                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                                &&&matchValue_1),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                                                                                &&&array)))
                                                                                                                                                                                   {
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                        &&&Applicative0),
                                                                                                                                                                                                                     &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&matchValue)))
                                                                                                                                                                                } else {
                                                                                                                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                                                       {
                                                                                                                                                                                        let i1_3 =
                                                                                                                                                                                            matchValue_1;
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                    let i1_3
                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                        i1_3.clone();
                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                        |usd__unused|
                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                                                                            &&&array),
                                                                                                                                                                                                                                                                                                                                                                                         &&&i1_3)
                                                                                                                                                                                                                                                                                                                                                })))),
                                                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                                                           let i1_3
                                                                                                                                                                                                                                               =
                                                                                                                                                                                                                                               i1_3.clone();
                                                                                                                                                                                                                                           move
                                                                                                                                                                                                                                               |res_prime|
                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                   &&&Applicative0),
                                                                                                                                                                                                                                                                                &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                                                                                                                                                       res_prime.clone(),
                                                                                                                                                                                                                                                                                                                                                                       add(string("b"),
                                                                                                                                                                                                                                                                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&i1_3),
                                                                                                                                                                                                                                                                                                                                                                                                             &&&1_i32),
                                                                                                                                                                                                                                                                                                                                                                           empty_1::<string,
                                                                                                                                                                                                                                                                                                                                                                                     &dyn Any>())))))
                                                                                                                                                                                                                                       }))
                                                                                                                                                                                    } else {
                                                                                                                                                                                        panic!("{}",
                                                                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.fs"),
                                  Data1: 428_i32,
                                  Data2: 603_i32,}).get_Message(),)
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                });
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM2(),
                                                                                                                                                                                                                                                                          &&&dictMonadRec),
                                                                                                                                                                                                                                       &&&go),
                                                                                                                                                                                                    &&&b),
                                                                                                                                                                 &&&0_i32)
                                                                                                                            }
                                                                                                                    })
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Data_Array_foldMap() -> &dyn Any {
        static Data_Array_foldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldMap.get_or_init(||
                                           &Func1::new(move |dictMonoid|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldMap1(),
                                                                                            dictMonoid)))
    }
    pub fn Data_Array_foldM_0040433() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_Array::Data_Array_foldM_tco(dictMonad))
    }
    pub fn Data_Array_foldM_0040433_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Array_foldM_0040433_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Array_foldM_0040433_002d1.get_or_init(||
                                                       Lazy(Data_Array_foldM_0040433.clone()))
    }
    pub fn Data_Array_foldM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_Array_foldM_0040433_002d1 =
                            Data_Array_foldM_0040433_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |f|
                            &Func1::new({
                                            let Data_Array_foldM_0040433_002d1
                                                =
                                                Data_Array_foldM_0040433_002d1.clone();
                                            let f = f.clone();
                                            move |b|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                       &&&PureScript_Data_Array::Data_Array_unconsImpl()),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let b
                                                                                                                                          =
                                                                                                                                          b.clone();
                                                                                                                                      move
                                                                                                                                          |v|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                              &&&Applicative0),
                                                                                                                                                                           &&&b)
                                                                                                                                  })),
                                                                                 &&&Func1::new({
                                                                                                   let Data_Array_foldM_0040433_002d1
                                                                                                       =
                                                                                                       Data_Array_foldM_0040433_002d1.clone();
                                                                                                   let b
                                                                                                       =
                                                                                                       b.clone();
                                                                                                   move
                                                                                                       |a|
                                                                                                       &Func1::new({
                                                                                                                       let Data_Array_foldM_0040433_002d1
                                                                                                                           =
                                                                                                                           Data_Array_foldM_0040433_002d1.clone();
                                                                                                                       let a
                                                                                                                           =
                                                                                                                           a.clone();
                                                                                                                       move
                                                                                                                           |as_var|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                  &&&Bind1),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                     &&&b),
                                                                                                                                                                                                                                  &&&a)),
                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                              let Data_Array_foldM_0040433_002d1
                                                                                                                                                                                  =
                                                                                                                                                                                  Data_Array_foldM_0040433_002d1.clone();
                                                                                                                                                                              let as_var
                                                                                                                                                                                  =
                                                                                                                                                                                  as_var.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |b_prime|
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_Array_foldM_0040433_002d1.Value,
                                                                                                                                                                                                                                                                                                                            &&&dictMonad),
                                                                                                                                                                                                                                                                                         &&&f),
                                                                                                                                                                                                                                                      b_prime),
                                                                                                                                                                                                                   &&&as_var)
                                                                                                                                                                          }))
                                                                                                                   })
                                                                                               }))
                                        })
                    })
    }
    pub fn Data_Array_foldM() -> &dyn Any {
        static Data_Array_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_foldM.get_or_init(|| Data_Array_foldM_0040433_002d1.Value)
    }
    pub fn Data_Array_fold() -> &dyn Any {
        static Data_Array_fold: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_fold.get_or_init(||
                                        &Func1::new(move |dictMonoid|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_fold1(),
                                                                                         dictMonoid)))
    }
    pub fn Data_Array_findMap() -> &dyn Any {
        static Data_Array_findMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findMap.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                  &&&PureScript_Data_Array::Data_Array_findMapImpl()),
                                                                                                               &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                            &&&PureScript_Data_Maybe::Data_Maybe_isJust()))
    }
    pub fn Data_Array_findLastIndex() -> &dyn Any {
        static Data_Array_findLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findLastIndex.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                        &&&PureScript_Data_Array::Data_Array_findLastIndexImpl()),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |usd__arg1|
                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                  &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_insertBy() -> &dyn Any {
        static Data_Array_insertBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_insertBy.get_or_init(||
                                            &Func1::new(move |cmp|
                                                            &Func1::new({
                                                                            let cmp
                                                                                =
                                                                                cmp.clone();
                                                                            move
                                                                                |x|
                                                                                &Func1::new({
                                                                                                let x
                                                                                                    =
                                                                                                    x.clone();
                                                                                                move
                                                                                                    |ys|
                                                                                                    {
                                                                                                        let i =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                   &&&0_i32),
                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                  |v|
                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                      v),
                                                                                                                                                                                                                                   &&&1_i32))),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findLastIndex(),
                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                     |y|
                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                                               &&&x),
                                                                                                                                                                                                                                                                                                                                            y)),
                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)))),
                                                                                                                                                                                ys));
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                         &&&Func1::new({
                                                                                                                                                           let i
                                                                                                                                                               =
                                                                                                                                                               i.clone();
                                                                                                                                                           let ys
                                                                                                                                                               =
                                                                                                                                                               ys.clone();
                                                                                                                                                           move
                                                                                                                                                               |usd__unused|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_insertAt(),
                                                                                                                                                                                                                                                                                                         &&&i),
                                                                                                                                                                                                                                                                      &&&x),
                                                                                                                                                                                                                                   &&&ys))
                                                                                                                                                       }))
                                                                                                    }
                                                                                            })
                                                                        })))
    }
    pub fn Data_Array_insert() -> &dyn Any {
        static Data_Array_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_insert.get_or_init(||
                                          &Func1::new(move |dictOrd|
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_insertBy(),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                              dictOrd))))
    }
    pub fn Data_Array_findIndex() -> &dyn Any {
        static Data_Array_findIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_findIndex.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                    &&&PureScript_Data_Array::Data_Array_findIndexImpl()),
                                                                                                                 &&&Func1::new(move
                                                                                                                                   |usd__arg1|
                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_span() -> &dyn Any {
        static Data_Array_span: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_span.get_or_init(||
                                        &Func1::new(move |p|
                                                        &Func1::new({
                                                                        let p
                                                                            =
                                                                            p.clone();
                                                                        move
                                                                            |arr|
                                                                            {
                                                                                let matchValue:
                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                    |x|
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                        x)))),
                                                                                                                                               arr));
                                                                                if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                       =
                                                                                       matchValue.as_ref()
                                                                                   {
                                                                                    if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                &Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                        =>
                                                                                                                                                                        x.clone(),
                                                                                                                                                                        _
                                                                                                                                                                        =>
                                                                                                                                                                        unreachable!(),
                                                                                                                                                                    })).is_some()
                                                                                       {
                                                                                        &add(string("init"),
                                                                                             &&new_empty::<&dyn Any>(),
                                                                                             add(string("rest"),
                                                                                                 arr.clone(),
                                                                                                 empty_1::<string,
                                                                                                           &dyn Any>()))
                                                                                    } else {
                                                                                        match matchValue.as_ref()
                                                                                            {
                                                                                            Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                            =>
                                                                                            &add(string("init"),
                                                                                                 arr.clone(),
                                                                                                 add(string("rest"),
                                                                                                     &&new_empty::<&dyn Any>(),
                                                                                                     empty_1::<string,
                                                                                                               &dyn Any>())),
                                                                                            Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                            =>
                                                                                            {
                                                                                                let i =
                                                                                                    match matchValue.as_ref()
                                                                                                        {
                                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                        =>
                                                                                                        x.clone(),
                                                                                                        _
                                                                                                        =>
                                                                                                        unreachable!(),
                                                                                                    };
                                                                                                &add(string("init"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                             &&&0_i32),
                                                                                                                                                                          &&&i),
                                                                                                                                       arr),
                                                                                                     add(string("rest"),
                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                                 &&&i),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                 arr)),
                                                                                                                                           arr),
                                                                                                         empty_1::<string,
                                                                                                                   &dyn Any>()))
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    match matchValue.as_ref()
                                                                                        {
                                                                                        Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                        =>
                                                                                        &add(string("init"),
                                                                                             arr.clone(),
                                                                                             add(string("rest"),
                                                                                                 &&new_empty::<&dyn Any>(),
                                                                                                 empty_1::<string,
                                                                                                           &dyn Any>())),
                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                        =>
                                                                                        {
                                                                                            let i =
                                                                                                match matchValue.as_ref()
                                                                                                    {
                                                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                    =>
                                                                                                    x.clone(),
                                                                                                    _
                                                                                                    =>
                                                                                                    unreachable!(),
                                                                                                };
                                                                                            &add(string("init"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                         &&&0_i32),
                                                                                                                                                                      &&&i),
                                                                                                                                   arr),
                                                                                                 add(string("rest"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                                             &&&i),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                             arr)),
                                                                                                                                       arr),
                                                                                                     empty_1::<string,
                                                                                                               &dyn Any>()))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                    })))
    }
    pub fn Data_Array_takeWhile() -> &dyn Any {
        static Data_Array_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_takeWhile.get_or_init(||
                                             &Func1::new(move |p|
                                                             &Func1::new({
                                                                             let p
                                                                                 =
                                                                                 p.clone();
                                                                             move
                                                                                 |xs|
                                                                                 find(string("init"),
                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_span(),
                                                                                                                                                                                   &&&p),
                                                                                                                                                xs)))
                                                                         })))
    }
    pub fn Data_Array_find() -> &dyn Any {
        static Data_Array_find: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_find.get_or_init(||
                                        &Func1::new(move |f|
                                                        &Func1::new({
                                                                        let f
                                                                            =
                                                                            f.clone();
                                                                        move
                                                                            |xs|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                   &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let xs
                                                                                                                                                                                                         =
                                                                                                                                                                                                         xs.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |usd__unused|
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                          &&&xs)
                                                                                                                                                                                                 }))),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                                                                                   &&&f),
                                                                                                                                                xs))
                                                                    })))
    }
    pub fn Data_Array_filter() -> &dyn Any {
        static Data_Array_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_filter.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                           &&&PureScript_Data_Array::Data_Array_filterImpl()))
    }
    pub fn Data_Array_intersectBy() -> &dyn Any {
        static Data_Array_intersectBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_intersectBy.get_or_init(||
                                               &Func1::new(move |eq|
                                                               &Func1::new({
                                                                               let eq
                                                                                   =
                                                                                   eq.clone();
                                                                               move
                                                                                   |xs|
                                                                                   &Func1::new({
                                                                                                   let xs
                                                                                                       =
                                                                                                       xs.clone();
                                                                                                   move
                                                                                                       |ys|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_filter(),
                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                             let ys
                                                                                                                                                                                                 =
                                                                                                                                                                                                 ys.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |x|
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isJust(),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                                                                                                                                                                                                           x)),
                                                                                                                                                                                                                                                                     &&&ys))
                                                                                                                                                                                         })),
                                                                                                                                        &&&xs)
                                                                                               })
                                                                           })))
    }
    pub fn Data_Array_intersect() -> &dyn Any {
        static Data_Array_intersect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_intersect.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_intersectBy(),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                 dictEq))))
    }
    pub fn Data_Array_elemLastIndex() -> &dyn Any {
        static Data_Array_elemLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_elemLastIndex.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 &Func1::new({
                                                                                 let dictEq
                                                                                     =
                                                                                     dictEq.clone();
                                                                                 move
                                                                                     |x|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findLastIndex(),
                                                                                                                      &&&Func1::new({
                                                                                                                                        let x
                                                                                                                                            =
                                                                                                                                            x.clone();
                                                                                                                                        move
                                                                                                                                            |v|
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                   &&&dictEq),
                                                                                                                                                                                                                v),
                                                                                                                                                                             &&&x)
                                                                                                                                    }))
                                                                             })))
    }
    pub fn Data_Array_elemIndex() -> &dyn Any {
        static Data_Array_elemIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_elemIndex.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             &Func1::new({
                                                                             let dictEq
                                                                                 =
                                                                                 dictEq.clone();
                                                                             move
                                                                                 |x|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                  &&&Func1::new({
                                                                                                                                    let x
                                                                                                                                        =
                                                                                                                                        x.clone();
                                                                                                                                    move
                                                                                                                                        |v|
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                                                                            v),
                                                                                                                                                                         &&&x)
                                                                                                                                }))
                                                                         })))
    }
    pub fn Data_Array_notElem() -> &dyn Any {
        static Data_Array_notElem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_notElem.get_or_init(||
                                           &Func1::new(move |dictEq|
                                                           &Func1::new({
                                                                           let dictEq
                                                                               =
                                                                               dictEq.clone();
                                                                           move
                                                                               |a|
                                                                               &Func1::new({
                                                                                               let a
                                                                                                   =
                                                                                                   a.clone();
                                                                                               move
                                                                                                   |arr|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_isNothing()),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_elemIndex(),
                                                                                                                                                                                                                                             &&&dictEq),
                                                                                                                                                                                                          &&&a),
                                                                                                                                                                       arr))
                                                                                           })
                                                                       })))
    }
    pub fn Data_Array_elem() -> &dyn Any {
        static Data_Array_elem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_elem.get_or_init(||
                                        &Func1::new(move |dictEq|
                                                        &Func1::new({
                                                                        let dictEq
                                                                            =
                                                                            dictEq.clone();
                                                                        move
                                                                            |a|
                                                                            &Func1::new({
                                                                                            let a
                                                                                                =
                                                                                                a.clone();
                                                                                            move
                                                                                                |arr|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_isJust()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_elemIndex(),
                                                                                                                                                                                                                                          &&&dictEq),
                                                                                                                                                                                                       &&&a),
                                                                                                                                                                    arr))
                                                                                        })
                                                                    })))
    }
    pub fn Data_Array_dropWhile() -> &dyn Any {
        static Data_Array_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_dropWhile.get_or_init(||
                                             &Func1::new(move |p|
                                                             &Func1::new({
                                                                             let p
                                                                                 =
                                                                                 p.clone();
                                                                             move
                                                                                 |xs|
                                                                                 find(string("rest"),
                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_span(),
                                                                                                                                                                                   &&&p),
                                                                                                                                                xs)))
                                                                         })))
    }
    pub fn Data_Array_dropEnd() -> &dyn Any {
        static Data_Array_dropEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_dropEnd.get_or_init(||
                                           &Func1::new(move |n|
                                                           &Func1::new({
                                                                           let n
                                                                               =
                                                                               n.clone();
                                                                           move
                                                                               |xs|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_take(),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                                                            xs)),
                                                                                                                                                                                      &&&n)),
                                                                                                                xs)
                                                                       })))
    }
    pub fn Data_Array_drop() -> &dyn Any {
        static Data_Array_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_drop.get_or_init(||
                                        &Func1::new(move |n|
                                                        &Func1::new({
                                                                        let n
                                                                            =
                                                                            n.clone();
                                                                        move
                                                                            |xs|
                                                                            {
                                                                                let matchValue =
                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                    &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                 &&&n),
                                                                                                                                              &&&1_i32));
                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                 &matchValue)
                                                                                    {
                                                                                    0_i32
                                                                                    =>
                                                                                    xs.clone(),
                                                                                    _
                                                                                    =>
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                           &&&n),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                           xs)),
                                                                                                                     xs),
                                                                                }
                                                                            }
                                                                    })))
    }
    pub fn Data_Array_takeEnd() -> &dyn Any {
        static Data_Array_takeEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_takeEnd.get_or_init(||
                                           &Func1::new(move |n|
                                                           &Func1::new({
                                                                           let n
                                                                               =
                                                                               n.clone();
                                                                           move
                                                                               |xs|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_drop(),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                                                            xs)),
                                                                                                                                                                                      &&&n)),
                                                                                                                xs)
                                                                       })))
    }
    pub fn Data_Array_deleteAt() -> &dyn Any {
        static Data_Array_deleteAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_deleteAt.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                   &&&PureScript_Data_Array::Data_Array__deleteAt()),
                                                                                                                &&&Func1::new(move
                                                                                                                                  |usd__arg1|
                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_deleteBy() -> &dyn Any {
        static Data_Array_deleteBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_deleteBy.get_or_init(||
                                            &Func1::new(move |v|
                                                            &Func1::new({
                                                                            let v
                                                                                =
                                                                                v.clone();
                                                                            move
                                                                                |v1|
                                                                                &Func1::new({
                                                                                                let v1
                                                                                                    =
                                                                                                    v1.clone();
                                                                                                move
                                                                                                    |v2|
                                                                                                    {
                                                                                                        let matchValue =
                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                        let matchValue_1 =
                                                                                                            Sharpurs_Prelude::unbox(&&v1);
                                                                                                        let matchValue_2 =
                                                                                                            Sharpurs_Prelude::unbox(v2);
                                                                                                        if {
                                                                                                               let testExpr =
                                                                                                                   matchValue_2.clone();
                                                                                                               if !equals(testExpr.clone(),
                                                                                                                          new_empty::<&dyn Any>())
                                                                                                                  {
                                                                                                                   count_1(testExpr)
                                                                                                                       ==
                                                                                                                       0_i32
                                                                                                               } else {
                                                                                                                   false
                                                                                                               }
                                                                                                           }
                                                                                                           {
                                                                                                            &new_empty::<&dyn Any>()
                                                                                                        } else {
                                                                                                            let ys =
                                                                                                                matchValue_2;
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                   &&&ys),
                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                  let ys
                                                                                                                                                                                                      =
                                                                                                                                                                                                      ys.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |i|
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                          &&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial()),
                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                         let i
                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                             i.clone();
                                                                                                                                                                                                                                                         let ys
                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                             ys.clone();
                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteAt(),
                                                                                                                                                                                                                                                                                                                                                                    &&&i),
                                                                                                                                                                                                                                                                                                                                 &&&ys))
                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                              })),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                      &&&matchValue_1)),
                                                                                                                                                                                &&&ys))
                                                                                                        }
                                                                                                    }
                                                                                            })
                                                                        })))
    }
    pub fn Data_Array_delete() -> &dyn Any {
        static Data_Array_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_delete.get_or_init(||
                                          &Func1::new(move |dictEq|
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteBy(),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                              dictEq))))
    }
    pub fn Data_Array_difference() -> &dyn Any {
        static Data_Array_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_difference.get_or_init(||
                                              &Func1::new(move |dictEq|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldr(),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_delete(),
                                                                                                                                  dictEq))))
    }
    pub fn Data_Array_cons() -> &dyn Any {
        static Data_Array_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_cons.get_or_init(||
                                        &Func1::new(move |x|
                                                        &Func1::new({
                                                                        let x
                                                                            =
                                                                            x.clone();
                                                                        move
                                                                            |xs|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                &&&new_array(&[&x])),
                                                                                                             xs)
                                                                    })))
    }
    pub fn Data_Array_some_0040487() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_Array::Data_Array_some_tco(dictAlternative))
    }
    pub fn Data_Array_some_0040487_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Array_some_0040487_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Array_some_0040487_002d1.get_or_init(||
                                                      Lazy(Data_Array_some_0040487.clone()))
    }
    pub fn Data_Array_many_0040489() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_Array::Data_Array_many_tco(dictAlternative))
    }
    pub fn Data_Array_many_0040489_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Array_many_0040489_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Array_many_0040489_002d1.get_or_init(||
                                                      Lazy(Data_Array_many_0040489.clone()))
    }
    pub fn Data_Array_some_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Apply0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Apply0 = Apply0.clone();
                        let Data_Array_many_0040489_002d1 =
                            Data_Array_many_0040489_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_Array_many_0040489_002d1
                                                =
                                                Data_Array_many_0040489_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                       &&&Apply0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                             &&&Functor0),
                                                                                                                                                                                          &&&PureScript_Data_Array::Data_Array_cons()),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                       &&&dictLazy),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let Data_Array_many_0040489_002d1
                                                                                                                                          =
                                                                                                                                          Data_Array_many_0040489_002d1.clone();
                                                                                                                                      let v
                                                                                                                                          =
                                                                                                                                          v.clone();
                                                                                                                                      move
                                                                                                                                          |v1|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_Array_many_0040489_002d1.Value,
                                                                                                                                                                                                                                                 &&&dictAlternative),
                                                                                                                                                                                                              &&&dictLazy),
                                                                                                                                                                           &&&v)
                                                                                                                                  })))
                                        })
                    })
    }
    pub fn Data_Array_some() -> &dyn Any {
        static Data_Array_some: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_some.get_or_init(|| Data_Array_some_0040487_002d1.Value)
    }
    pub fn Data_Array_many_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Alt0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictAlternative)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Alt0 = Alt0.clone();
                        let Applicative0 = Applicative0.clone();
                        let Data_Array_some_0040487_002d1 =
                            Data_Array_some_0040487_002d1.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_Array_some_0040487_002d1
                                                =
                                                Data_Array_some_0040487_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                       &&&Alt0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_Array_some_0040487_002d1.Value,
                                                                                                                                                                                                                             &&&dictAlternative),
                                                                                                                                                                                          &&&dictLazy),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       &&&Applicative0),
                                                                                                                    &&&new_empty::<&dyn Any>()))
                                        })
                    })
    }
    pub fn Data_Array_many() -> &dyn Any {
        static Data_Array_many: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_many.get_or_init(|| Data_Array_many_0040489_002d1.Value)
    }
    pub fn Data_Array_concatMap() -> &dyn Any {
        static Data_Array_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_concatMap.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                 &&&PureScript_Control_Bind::Control_Bind_bindArray())))
    }
    pub fn Data_Array_mapMaybe() -> &dyn Any {
        static Data_Array_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_mapMaybe.get_or_init(||
                                            &Func1::new(move |f|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_concatMap(),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                         &&&new_empty::<&dyn Any>()),
                                                                                                                                                                                                      &&&PureScript_Data_Array::Data_Array_singleton())),
                                                                                                                                f))))
    }
    pub fn Data_Array_filterA() -> &dyn Any {
        static Data_Array_filterA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_filterA.get_or_init(||
                                           &Func1::new(move |dictApplicative|
                                                           {
                                                               let Functor0 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                            Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                               &Func1::new({
                                                                               let Functor0
                                                                                   =
                                                                                   Functor0.clone();
                                                                               let dictApplicative
                                                                                   =
                                                                                   dictApplicative.clone();
                                                                               move
                                                                                   |p|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_Traversable::Data_Traversable_traversableArray()),
                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                            let p
                                                                                                                                                                                                                =
                                                                                                                                                                                                                p.clone();
                                                                                                                                                                                                            move
                                                                                                                                                                                                                |x|
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                       &&&Functor0),
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
                                                                                                                                                                                                                                                                                                                       x)),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                                    x))
                                                                                                                                                                                                        }))),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_mapMaybe(),
                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                            |v|
                                                                                                                                                                                                            {
                                                                                                                                                                                                                let matchValue:
                                                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&match matchValue.as_ref()
                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                              });
                                                                                                                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                 &matchValue_1)
                                                                                                                                                                                                                    {
                                                                                                                                                                                                                    0_i32
                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                    _
                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                                }
                                                                                                                                                                                                            }))))
                                                                           })
                                                           }))
    }
    pub fn Data_Array_catMaybes() -> &dyn Any {
        static Data_Array_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_catMaybes.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_mapMaybe(),
                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn())))
    }
    pub fn Data_Array_any() -> &dyn Any {
        static Data_Array_any: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_any.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                        &&&PureScript_Data_Array::Data_Array_anyImpl()))
    }
    pub fn Data_Array_nubByEq() -> &dyn Any {
        static Data_Array_nubByEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_nubByEq.get_or_init(||
                                           &Func1::new(move |eq|
                                                           &Func1::new({
                                                                           let eq
                                                                               =
                                                                               eq.clone();
                                                                           move
                                                                               |xs|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                         &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                      &&&PureScript_Data_Array_ST::Data_Array_ST_new()),
                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                     let xs
                                                                                                                                                                         =
                                                                                                                                                                         xs.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |arr|
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_foreach(),
                                                                                                                                                                                                                                                                                                                   &&&xs),
                                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                                  let arr
                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                      arr.clone();
                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                      |x|
                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Array::Data_Array_not()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_any(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        let x
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            x.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |v|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                v),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    })))),
                                                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&arr))),
                                                                                                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                         let x
                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                             x.clone();
                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                             |e|
                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_when(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    e)),
                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Array::Data_Array_void()),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&x),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&arr)))
                                                                                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                            let arr
                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                arr.clone();
                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                |usd__unused|
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                                                                 &&&arr)
                                                                                                                                                                                                                        }))
                                                                                                                                                                 })))
                                                                       })))
    }
    pub fn Data_Array_nubEq() -> &dyn Any {
        static Data_Array_nubEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_nubEq.get_or_init(||
                                         &Func1::new(move |dictEq|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubByEq(),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                             dictEq))))
    }
    pub fn Data_Array_unionBy() -> &dyn Any {
        static Data_Array_unionBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_unionBy.get_or_init(||
                                           &Func1::new(move |eq|
                                                           &Func1::new({
                                                                           let eq
                                                                               =
                                                                               eq.clone();
                                                                           move
                                                                               |xs|
                                                                               &Func1::new({
                                                                                               let xs
                                                                                                   =
                                                                                                   xs.clone();
                                                                                               move
                                                                                                   |ys|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                                       &&&xs),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldl(),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteBy(),
                                                                                                                                                                                                                                                                                                                   &&&eq))),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubByEq(),
                                                                                                                                                                                                                                                                                &&&eq),
                                                                                                                                                                                                                                             ys)),
                                                                                                                                                                       &&&xs))
                                                                                           })
                                                                       })))
    }
    pub fn Data_Array_union() -> &dyn Any {
        static Data_Array_union: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_union.get_or_init(||
                                         &Func1::new(move |dictEq|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unionBy(),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                             dictEq))))
    }
    pub fn Data_Array_alterAt() -> &dyn Any {
        static Data_Array_alterAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_alterAt.get_or_init(||
                                           &Func1::new(move |i|
                                                           &Func1::new({
                                                                           let i
                                                                               =
                                                                               i.clone();
                                                                           move
                                                                               |f|
                                                                               &Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   {
                                                                                                       let go =
                                                                                                           &Func1::new({
                                                                                                                           let xs
                                                                                                                               =
                                                                                                                               xs.clone();
                                                                                                                           move
                                                                                                                               |x|
                                                                                                                               {
                                                                                                                                   let matchValue:
                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                  x));
                                                                                                                                   match matchValue.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                       =>
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_updateAt(),
                                                                                                                                                                                                                                              &&&i),
                                                                                                                                                                                                           &&matchValue_1_0),
                                                                                                                                                                        &&&xs),
                                                                                                                                       _
                                                                                                                                       =>
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteAt(),
                                                                                                                                                                                                           &&&i),
                                                                                                                                                                        &&&xs),
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       });
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                           &&&go),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_index(),
                                                                                                                                                                                                              xs),
                                                                                                                                                                           &&&i))
                                                                                                   }
                                                                                           })
                                                                       })))
    }
    pub fn Data_Array_all() -> &dyn Any {
        static Data_Array_all: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Array_all.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                        &&&PureScript_Data_Array::Data_Array_allImpl()))
    }
}
