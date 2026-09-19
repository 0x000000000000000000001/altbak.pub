pub mod System {
    use super::*;
    use crate::module_3bd9ae6a::Native_::LrcPtr;
    use crate::module_eae1ac5e::String_::append;
    use crate::module_eae1ac5e::String_::isEmpty;
    use crate::module_eae1ac5e::String_::string;
    #[derive(Clone, Debug, Default)]
    pub struct Array {}
    impl System::Array {
        pub fn _ctor() -> LrcPtr<System::Array> {
            ();
            ();
            LrcPtr::new(System::Array {})
        }
    }
    impl core::fmt::Display for System::Array {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, Default)]
    pub struct Enum {}
    impl System::Enum {
        pub fn _ctor() -> LrcPtr<System::Enum> {
            ();
            ();
            LrcPtr::new(System::Enum {})
        }
    }
    impl core::fmt::Display for System::Enum {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, Default)]
    pub struct Exception {
        message: string,
    }
    impl System::Exception {
        pub fn _ctor__Z721C83C5(message: string) -> LrcPtr<System::Exception> {
            let message_1: string;
            ();
            message_1 = message;
            ();
            LrcPtr::new(System::Exception { message: message_1 })
        }
        pub fn _ctor() -> LrcPtr<System::Exception> {
            System::Exception::_ctor__Z721C83C5(string(""))
        }
        pub fn get_Message(&self) -> string {
            if isEmpty(self.message.clone()) {
                string("Specified argument was out of the range of valid values.")
            } else {
                self.message.clone()
            }
        }
    }
    impl core::fmt::Display for System::Exception {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, Default)]
    pub struct InvalidOperationException {
        message: string,
    }
    impl System::InvalidOperationException {
        pub fn _ctor__Z721C83C5(message: string) -> LrcPtr<System::InvalidOperationException> {
            let message_1: string;
            ();
            message_1 = message;
            ();
            LrcPtr::new(System::InvalidOperationException { message: message_1 })
        }
        pub fn _ctor() -> LrcPtr<System::InvalidOperationException> {
            System::InvalidOperationException::_ctor__Z721C83C5(string(""))
        }
        pub fn get_Message(&self) -> string {
            if isEmpty(self.message.clone()) {
                string("Operation is not valid due to the current state of the object.")
            } else {
                self.message.clone()
            }
        }
    }
    impl core::fmt::Display for System::InvalidOperationException {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, Default)]
    pub struct ArgumentException {
        paramName: string,
        message: string,
    }
    impl System::ArgumentException {
        pub fn _ctor__Z384F8060(
            message: string,
            paramName: string,
        ) -> LrcPtr<System::ArgumentException> {
            let paramName_1: string;
            let message_1: string;
            ();
            message_1 = message;
            paramName_1 = paramName;
            ();
            LrcPtr::new(System::ArgumentException {
                paramName: paramName_1,
                message: message_1,
            })
        }
        pub fn _ctor() -> LrcPtr<System::ArgumentException> {
            System::ArgumentException::_ctor__Z384F8060(string(""), string(""))
        }
        pub fn _ctor__Z721C83C5(message: string) -> LrcPtr<System::ArgumentException> {
            System::ArgumentException::_ctor__Z384F8060(message, string(""))
        }
        pub fn get_Message(&self) -> string {
            let message: string = if isEmpty(self.message.clone()) {
                string("Value does not fall within the expected range.")
            } else {
                self.message.clone()
            };
            if isEmpty(self.paramName.clone()) {
                message.clone()
            } else {
                append(
                    append(
                        append(message, string(" (Parameter \'")),
                        self.paramName.clone(),
                    ),
                    string("\')"),
                )
            }
        }
        pub fn get_ParamName(&self) -> string {
            self.paramName.clone()
        }
    }
    impl core::fmt::Display for System::ArgumentException {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, Default)]
    pub struct ArgumentOutOfRangeException {
        paramName: string,
        message: string,
    }
    impl System::ArgumentOutOfRangeException {
        pub fn _ctor__Z384F8060(
            paramName: string,
            message: string,
        ) -> LrcPtr<System::ArgumentOutOfRangeException> {
            let paramName_1: string;
            let message_1: string;
            ();
            paramName_1 = paramName;
            message_1 = message;
            ();
            LrcPtr::new(System::ArgumentOutOfRangeException {
                paramName: paramName_1,
                message: message_1,
            })
        }
        pub fn _ctor() -> LrcPtr<System::ArgumentOutOfRangeException> {
            System::ArgumentOutOfRangeException::_ctor__Z384F8060(string(""), string(""))
        }
        pub fn _ctor__Z721C83C5(paramName: string) -> LrcPtr<System::ArgumentOutOfRangeException> {
            System::ArgumentOutOfRangeException::_ctor__Z384F8060(paramName, string(""))
        }
        pub fn get_Message(&self) -> string {
            let message: string = if isEmpty(self.message.clone()) {
                string("Specified argument was out of the range of valid values.")
            } else {
                self.message.clone()
            };
            if isEmpty(self.paramName.clone()) {
                message.clone()
            } else {
                append(
                    append(
                        append(message, string(" (Parameter \'")),
                        self.paramName.clone(),
                    ),
                    string("\')"),
                )
            }
        }
        pub fn get_ParamName(&self) -> string {
            self.paramName.clone()
        }
    }
    impl core::fmt::Display for System::ArgumentOutOfRangeException {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub mod Collections {
        use super::*;
        pub mod Generic {
            use super::*;
            use crate::module_3350cf54::Range_::rangeNumeric;
            use crate::module_3bd9ae6a::Native_::defaultOf;
            use crate::module_3bd9ae6a::Native_::Func0;
            use crate::module_3bd9ae6a::Native_::Func1;
            use crate::module_3bd9ae6a::Native_::MutCell;
            use crate::module_52af85ec::Seq_::delay;
            use crate::module_52af85ec::Seq_::map;
            use crate::module_52af85ec::Seq_::toArray;
            use crate::module_6dcf8b93::NativeArray_::add;
            use crate::module_6dcf8b93::NativeArray_::count as count_1;
            use crate::module_6dcf8b93::NativeArray_::new_init;
            use crate::module_6dcf8b93::NativeArray_::new_with_capacity;
            use crate::module_6dcf8b93::NativeArray_::Array as Array_1;
            use crate::module_971078fd::Interfaces_::System::Collections::Generic::IEnumerable_1;
            use crate::module_971078fd::Interfaces_::System::Collections::Generic::IEnumerator_1;
            use crate::module_c6216f2::Array_::copyTo;
            use crate::module_c6216f2::Array_::fill;
            #[derive(Clone, Debug, Default)]
            pub struct Stack_1<T: Eq + core::hash::Hash + Clone + 'static> {
                contents: MutCell<Array_1<T>>,
                count: MutCell<i32>,
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> System::Collections::Generic::Stack_1<T> {
                fn _ctor__Z3B4C077E(
                    initialContents: Array_1<T>,
                    initialCount: i32,
                ) -> LrcPtr<System::Collections::Generic::Stack_1<T>> {
                    let contents: Array_1<T>;
                    let count: i32;
                    ();
                    contents = initialContents;
                    count = initialCount;
                    ();
                    LrcPtr::new(System::Collections::Generic::Stack_1::<T> {
                        contents: MutCell::new(contents),
                        count: MutCell::new(count),
                    })
                }
                pub fn _ctor__Z524259A4(
                    initialCapacity: i32,
                ) -> LrcPtr<System::Collections::Generic::Stack_1<T>> {
                    System::Collections::Generic::Stack_1::_ctor__Z3B4C077E(
                        new_init(&defaultOf(), initialCapacity),
                        0_i32,
                    )
                }
                pub fn _ctor() -> LrcPtr<System::Collections::Generic::Stack_1<T>> {
                    System::Collections::Generic::Stack_1::_ctor__Z524259A4(4_i32)
                }
                pub fn _ctor__BB573A(
                    xs: LrcPtr<dyn IEnumerable_1<T>>,
                ) -> LrcPtr<System::Collections::Generic::Stack_1<T>> {
                    let arr: Array_1<T> = toArray(xs);
                    System::Collections::Generic::Stack_1::_ctor__Z3B4C077E(
                        arr.clone(),
                        count_1(arr),
                    )
                }
                pub fn Ensure_Z524259A4(&self, newSize: i32) {
                    let oldSize: i32 = count_1(self.contents.get().clone());
                    if newSize > oldSize {
                        let old: Array_1<T> = self.contents.get().clone();
                        self.contents
                            .set(new_init(&defaultOf(), newSize.max(oldSize * 2_i32)));
                        copyTo(
                            old,
                            0_i32,
                            self.contents.get().clone(),
                            0_i32,
                            self.count.get().clone(),
                        )
                    }
                }
                pub fn get_Count(&self) -> i32 {
                    self.count.get().clone()
                }
                pub fn Pop(&self) -> T {
                    self.count.set(self.count.get().clone() - 1_i32);
                    (self.contents.get())[self.count.get()].clone()
                }
                pub fn Peek(&self) -> T {
                    (self.contents.get())[self.count.get().clone() - 1_i32].clone()
                }
                pub fn Contains_2B595(&self, x: T) -> bool {
                    let found: MutCell<bool> = MutCell::new(false);
                    let i: MutCell<i32> = MutCell::new(0_i32);
                    while if i.get() < self.count.get().clone() {
                        !found.get()
                    } else {
                        false
                    } {
                        if x.clone() == (self.contents.get())[i.get()].clone() {
                            found.set(true)
                        } else {
                            i.set(i.get() + 1_i32)
                        };
                    }
                    found.get()
                }
                pub fn TryPeek_1F3DB691(&self, result: &MutCell<T>) -> bool {
                    if self.count.get().clone() > 0_i32 {
                        result.set(self.Peek());
                        true
                    } else {
                        false
                    }
                }
                pub fn TryPop_1F3DB691(&self, result: &MutCell<T>) -> bool {
                    if self.count.get().clone() > 0_i32 {
                        result.set(self.Pop());
                        true
                    } else {
                        false
                    }
                }
                pub fn Push_2B595(&self, x: T) {
                    self.Ensure_Z524259A4(self.count.get().clone() + 1_i32);
                    (self.contents.get()).get_mut()[self.count.get() as usize] = x;
                    self.count.set(self.count.get().clone() + 1_i32)
                }
                pub fn Clear(&self) {
                    self.count.set(0_i32);
                    fill(
                        self.contents.get().clone(),
                        0_i32,
                        count_1(self.contents.get().clone()),
                        defaultOf(),
                    )
                }
                pub fn TrimExcess(&self) {
                    if self.count.get().clone() as f64 / count_1(self.contents.get().clone()) as f64
                        > 0.9_f64
                    {
                        self.Ensure_Z524259A4(self.count.get().clone());
                    };
                }
                pub fn ToArray(&self) -> Array_1<T> {
                    let res: Array_1<T> = new_with_capacity::<T>(self.count.get().clone());
                    for i in 0_i32..=self.count.get().clone() - 1_i32 {
                        add(
                            res.clone(),
                            (self.contents.get())[self.count.get().clone() - 1_i32 - i].clone(),
                        );
                    }
                    res.clone()
                }
                pub fn toSeq(&self) -> LrcPtr<dyn IEnumerable_1<T>> {
                    let count: i32 = self.count.get().clone();
                    let contents: Array_1<T> = self.contents.get().clone();
                    delay(Func0::new({
                        let contents = contents.clone();
                        let count = count.clone();
                        move || {
                            map(
                                Func1::new({
                                    let contents = contents.clone();
                                    move |i: i32| contents[i].clone()
                                }),
                                rangeNumeric(count - 1_i32, -1_i32, 0_i32),
                            )
                        }
                    }))
                }
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> core::fmt::Display
                for System::Collections::Generic::Stack_1<T>
            {
                fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                    write!(f, "{}", core::any::type_name::<Self>())
                }
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> IEnumerable_1<T> for Stack_1<T> {
                fn GetEnumerator(&self) -> LrcPtr<dyn IEnumerator_1<T>> {
                    IEnumerable_1::GetEnumerator(self.toSeq().as_ref())
                }
            }
            #[derive(Clone, Debug, Default)]
            pub struct Queue_1<T: Eq + core::hash::Hash + Clone + 'static> {
                contents: MutCell<Array_1<T>>,
                count: MutCell<i32>,
                head: MutCell<i32>,
                tail: MutCell<i32>,
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> System::Collections::Generic::Queue_1<T> {
                fn _ctor__Z3B4C077E(
                    initialContents: Array_1<T>,
                    initialCount: i32,
                ) -> LrcPtr<System::Collections::Generic::Queue_1<T>> {
                    let contents: Array_1<T>;
                    let count: i32;
                    let head: i32;
                    let tail: i32;
                    ();
                    contents = initialContents;
                    count = initialCount;
                    head = 0_i32;
                    tail = initialCount;
                    ();
                    LrcPtr::new(System::Collections::Generic::Queue_1::<T> {
                        contents: MutCell::new(contents),
                        count: MutCell::new(count),
                        head: MutCell::new(head),
                        tail: MutCell::new(tail),
                    })
                }
                pub fn _ctor__Z524259A4(
                    initialCapacity: i32,
                ) -> LrcPtr<System::Collections::Generic::Queue_1<T>> {
                    if initialCapacity < 0_i32 {
                        panic!("{}", string("capacity is less than 0"),);
                    }
                    System::Collections::Generic::Queue_1::_ctor__Z3B4C077E(
                        new_init(&defaultOf(), initialCapacity),
                        0_i32,
                    )
                }
                pub fn _ctor() -> LrcPtr<System::Collections::Generic::Queue_1<T>> {
                    System::Collections::Generic::Queue_1::_ctor__Z524259A4(4_i32)
                }
                pub fn _ctor__BB573A(
                    xs: LrcPtr<dyn IEnumerable_1<T>>,
                ) -> LrcPtr<System::Collections::Generic::Queue_1<T>> {
                    let arr: Array_1<T> = toArray(xs);
                    System::Collections::Generic::Queue_1::_ctor__Z3B4C077E(
                        arr.clone(),
                        count_1(arr),
                    )
                }
                pub fn get_Count(&self) -> i32 {
                    self.count.get().clone()
                }
                pub fn Enqueue_2B595(&self, value: T) {
                    if self.count.get().clone() == self.size() {
                        self.ensure_Z524259A4(self.count.get().clone() + 1_i32);
                    }
                    (self.contents.get()).get_mut()[self.tail.get() as usize] = value;
                    self.tail
                        .set((self.tail.get().clone() + 1_i32) % self.size());
                    self.count.set(self.count.get().clone() + 1_i32)
                }
                pub fn Dequeue(&self) -> T {
                    if self.count.get().clone() == 0_i32 {
                        panic!("{}", string("Queue is empty"),);
                    }
                    {
                        let value: T = (self.contents.get())[self.head.get()].clone();
                        self.head
                            .set((self.head.get().clone() + 1_i32) % self.size());
                        self.count.set(self.count.get().clone() - 1_i32);
                        value
                    }
                }
                pub fn Peek(&self) -> T {
                    if self.count.get().clone() == 0_i32 {
                        panic!("{}", string("Queue is empty"),);
                    }
                    (self.contents.get())[self.head.get()].clone()
                }
                pub fn TryDequeue_1F3DB691(&self, result: &MutCell<T>) -> bool {
                    if self.count.get().clone() == 0_i32 {
                        false
                    } else {
                        result.set(self.Dequeue());
                        true
                    }
                }
                pub fn TryPeek_1F3DB691(&self, result: &MutCell<T>) -> bool {
                    if self.count.get().clone() == 0_i32 {
                        false
                    } else {
                        result.set(self.Peek());
                        true
                    }
                }
                pub fn Contains_2B595(&self, x: T) -> bool {
                    let found: MutCell<bool> = MutCell::new(false);
                    let i: MutCell<i32> = MutCell::new(0_i32);
                    while if i.get() < self.count.get().clone() {
                        !found.get()
                    } else {
                        false
                    } {
                        if x.clone()
                            == (self.contents.get())[self.toIndex_Z524259A4(i.get())].clone()
                        {
                            found.set(true)
                        } else {
                            i.set(i.get() + 1_i32)
                        };
                    }
                    found.get()
                }
                pub fn Clear(&self) {
                    self.count.set(0_i32);
                    self.head.set(0_i32);
                    self.tail.set(0_i32);
                    fill(self.contents.get().clone(), 0_i32, self.size(), defaultOf())
                }
                pub fn TrimExcess(&self) {
                    if self.count.get().clone() as f64 / count_1(self.contents.get().clone()) as f64
                        > 0.9_f64
                    {
                        self.ensure_Z524259A4(self.count.get().clone());
                    };
                }
                pub fn ToArray(&self) -> Array_1<T> {
                    let res: Array_1<T> = new_with_capacity::<T>(self.count.get().clone());
                    for i in 0_i32..=self.count.get().clone() - 1_i32 {
                        add(
                            res.clone(),
                            (self.contents.get())[self.toIndex_Z524259A4(i)].clone(),
                        );
                    }
                    res.clone()
                }
                pub fn CopyTo_Z3B4C077E(&self, target: Array_1<T>, start: i32) {
                    let i: MutCell<i32> = MutCell::new(start);
                    for i_1 in 0_i32..=self.count.get().clone() - 1_i32 {
                        target.get_mut()[(start + i_1) as usize] =
                            (self.contents.get())[self.toIndex_Z524259A4(i_1)].clone();
                    }
                }
                pub fn size(&self) -> i32 {
                    count_1(self.contents.get().clone())
                }
                pub fn toIndex_Z524259A4(&self, i: i32) -> i32 {
                    (self.head.get().clone() + i) % self.size()
                }
                pub fn ensure_Z524259A4(&self, requiredSize: i32) {
                    let newBuffer: Array_1<T> = new_init(&defaultOf(), requiredSize);
                    if self.head.get().clone() < self.tail.get().clone() {
                        copyTo(
                            self.contents.get().clone(),
                            self.head.get().clone(),
                            newBuffer.clone(),
                            0_i32,
                            self.count.get().clone(),
                        )
                    } else {
                        copyTo(
                            self.contents.get().clone(),
                            self.head.get().clone(),
                            newBuffer.clone(),
                            0_i32,
                            self.size() - self.head.get().clone(),
                        );
                        copyTo(
                            self.contents.get().clone(),
                            0_i32,
                            newBuffer.clone(),
                            self.size() - self.head.get().clone(),
                            self.tail.get().clone(),
                        )
                    }
                    self.head.set(0_i32);
                    self.tail.set(self.count.get().clone());
                    self.contents.set(newBuffer)
                }
                pub fn toSeq(&self) -> LrcPtr<dyn IEnumerable_1<T>> {
                    let head: i32 = self.head.get().clone();
                    let count: i32 = self.count.get().clone();
                    let contents: Array_1<T> = self.contents.get().clone();
                    delay(Func0::new({
                        let contents = contents.clone();
                        let count = count.clone();
                        let head = head.clone();
                        move || {
                            map(
                                Func1::new({
                                    let contents = contents.clone();
                                    move |i: i32| {
                                        contents[(head + i) % count_1(contents.clone())].clone()
                                    }
                                }),
                                rangeNumeric(0_i32, 1_i32, count - 1_i32),
                            )
                        }
                    }))
                }
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> core::fmt::Display
                for System::Collections::Generic::Queue_1<T>
            {
                fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                    write!(f, "{}", core::any::type_name::<Self>())
                }
            }
            impl<T: Eq + core::hash::Hash + Clone + 'static> IEnumerable_1<T> for Queue_1<T> {
                fn GetEnumerator(&self) -> LrcPtr<dyn IEnumerator_1<T>> {
                    IEnumerable_1::GetEnumerator(self.toSeq().as_ref())
                }
            }
        }
    }
    pub mod Text {
        use super::*;
        use crate::module_3bd9ae6a::Native_::fromFluent;
        use crate::module_3bd9ae6a::Native_::Func1;
        use crate::module_3bd9ae6a::Native_::Lrc;
        use crate::module_6dcf8b93::NativeArray_::add;
        use crate::module_6dcf8b93::NativeArray_::new_with_capacity;
        use crate::module_6dcf8b93::NativeArray_::Array as Array_1;
        use crate::module_c6216f2::Array_::sumBy;
        use crate::module_eae1ac5e::String_::concat;
        use crate::module_eae1ac5e::String_::fromChars;
        use crate::module_eae1ac5e::String_::length;
        use crate::module_eae1ac5e::String_::ofBoolean;
        use crate::module_eae1ac5e::String_::ofChar;
        use crate::module_eae1ac5e::String_::substring2;
        use crate::module_eae1ac5e::String_::toString;
        #[derive(Clone, Debug, Default)]
        pub struct StringBuilder {
            buf: Array_1<string>,
        }
        impl System::Text::StringBuilder {
            pub fn _ctor__Z18115A39(
                value: string,
                capacity: i32,
            ) -> LrcPtr<System::Text::StringBuilder> {
                let buf: Array_1<string>;
                ();
                buf = new_with_capacity::<string>(capacity);
                if !isEmpty(value.clone()) {
                    add(buf.clone(), value);
                }
                ();
                LrcPtr::new(System::Text::StringBuilder { buf: buf })
            }
            pub fn _ctor__Z524259A4(capacity: i32) -> LrcPtr<System::Text::StringBuilder> {
                System::Text::StringBuilder::_ctor__Z18115A39(string(""), capacity)
            }
            pub fn _ctor__Z721C83C5(value: string) -> LrcPtr<System::Text::StringBuilder> {
                System::Text::StringBuilder::_ctor__Z18115A39(value, 16_i32)
            }
            pub fn _ctor() -> LrcPtr<System::Text::StringBuilder> {
                System::Text::StringBuilder::_ctor__Z18115A39(string(""), 16_i32)
            }
            pub fn Append_Z721C83C5(
                self: &Lrc<Self>,
                s: string,
            ) -> LrcPtr<System::Text::StringBuilder> {
                fromFluent({
                    add(self.buf.clone(), s);
                    self.clone()
                })
            }
            pub fn Append_Z1FBCCD16(
                self: &Lrc<Self>,
                o: bool,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(ofBoolean(o))
            }
            pub fn Append_244C7CD6(
                self: &Lrc<Self>,
                c: char,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(ofChar(c))
            }
            pub fn Append_Z510FF069(
                self: &Lrc<Self>,
                o: i8,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_244D3E44(self: &Lrc<Self>, o: u8) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z524259E6(
                self: &Lrc<Self>,
                o: i16,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z6EF82811(
                self: &Lrc<Self>,
                o: u16,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z524259A4(
                self: &Lrc<Self>,
                o: i32,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z6EF827D7(
                self: &Lrc<Self>,
                o: u32,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z524259C1(
                self: &Lrc<Self>,
                o: i64,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z6EF827B6(
                self: &Lrc<Self>,
                o: u64,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_Z7138B98C(
                self: &Lrc<Self>,
                o: f32,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_5E38073B(
                self: &Lrc<Self>,
                o: f64,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(o))
            }
            pub fn Append_487EF8FB(
                self: &Lrc<Self>,
                s: string,
                index: i32,
                count: i32,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(substring2(s, index, count))
            }
            pub fn Append_Z372E4D23(
                self: &Lrc<Self>,
                cs: Array_1<char>,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(fromChars(cs))
            }
            pub fn Append_43A65C09(
                self: &Lrc<Self>,
                sb: LrcPtr<System::Text::StringBuilder>,
            ) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(toString(sb))
            }
            pub fn AppendLine(self: &Lrc<Self>) -> LrcPtr<System::Text::StringBuilder> {
                self.Append_Z721C83C5(string("\n"))
            }
            pub fn AppendLine_Z721C83C5(
                self: &Lrc<Self>,
                s: string,
            ) -> LrcPtr<System::Text::StringBuilder> {
                (self.Append_Z721C83C5(s)).AppendLine()
            }
            pub fn Clear(self: &Lrc<Self>) -> LrcPtr<System::Text::StringBuilder> {
                fromFluent({
                    self.buf.get_mut().clear();
                    self.clone()
                })
            }
            pub fn get_Length(&self) -> i32 {
                sumBy(Func1::new(move |s: string| length(s)), self.buf.clone())
            }
            pub fn ToString_(&self) -> string {
                concat(self.buf.clone())
            }
            pub fn ToString_Z37302880(&self, index: i32, count: i32) -> string {
                substring2(toString(self.clone()), index, count)
            }
        }
        impl core::fmt::Display for System::Text::StringBuilder {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", self.ToString_())
            }
        }
    }
}
