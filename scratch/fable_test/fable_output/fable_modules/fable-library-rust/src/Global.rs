pub mod Global_ {
    use super::*;
    pub mod SR {
        use super::*;
        use crate::module_3bd9ae6a::Native_::MutCell;
        use crate::module_eae1ac5e::String_::string;
        pub fn arrayWasEmpty() -> string {
            static arrayWasEmpty: MutCell<Option<string>> = MutCell::new(None);
            arrayWasEmpty.get_or_init(|| string("The input array was empty."))
        }
        pub fn arrayIndexOutOfBounds() -> string {
            static arrayIndexOutOfBounds: MutCell<Option<string>> = MutCell::new(None);
            arrayIndexOutOfBounds
                .get_or_init(|| string("The index was outside the range of elements in the array."))
        }
        pub fn arraysHadDifferentLengths() -> string {
            static arraysHadDifferentLengths: MutCell<Option<string>> = MutCell::new(None);
            arraysHadDifferentLengths.get_or_init(|| string("The arrays have different lengths."))
        }
        pub fn enumerationAlreadyFinished() -> string {
            static enumerationAlreadyFinished: MutCell<Option<string>> = MutCell::new(None);
            enumerationAlreadyFinished.get_or_init(|| string("Enumeration already finished."))
        }
        pub fn enumerationNotStarted() -> string {
            static enumerationNotStarted: MutCell<Option<string>> = MutCell::new(None);
            enumerationNotStarted
                .get_or_init(|| string("Enumeration has not started. Call MoveNext."))
        }
        pub fn indexOutOfBounds() -> string {
            static indexOutOfBounds: MutCell<Option<string>> = MutCell::new(None);
            indexOutOfBounds
                .get_or_init(|| string("The index was outside the range of elements in the list."))
        }
        pub fn inputListWasEmpty() -> string {
            static inputListWasEmpty: MutCell<Option<string>> = MutCell::new(None);
            inputListWasEmpty.get_or_init(|| string("The input list was empty."))
        }
        pub fn inputMustBeNonNegative() -> string {
            static inputMustBeNonNegative: MutCell<Option<string>> = MutCell::new(None);
            inputMustBeNonNegative.get_or_init(|| string("The input must be non-negative."))
        }
        pub fn inputMustBePositive() -> string {
            static inputMustBePositive: MutCell<Option<string>> = MutCell::new(None);
            inputMustBePositive.get_or_init(|| string("The input must be positive."))
        }
        pub fn inputSequenceEmpty() -> string {
            static inputSequenceEmpty: MutCell<Option<string>> = MutCell::new(None);
            inputSequenceEmpty.get_or_init(|| string("The input sequence was empty."))
        }
        pub fn inputSequenceTooLong() -> string {
            static inputSequenceTooLong: MutCell<Option<string>> = MutCell::new(None);
            inputSequenceTooLong
                .get_or_init(|| string("The input sequence contains more than one element."))
        }
        pub fn keyNotFound() -> string {
            static keyNotFound: MutCell<Option<string>> = MutCell::new(None);
            keyNotFound
                .get_or_init(|| string("The item, key, or index was not found in the collection."))
        }
        pub fn keyNotFoundAlt() -> string {
            static keyNotFoundAlt: MutCell<Option<string>> = MutCell::new(None);
            keyNotFoundAlt.get_or_init(|| {
                string("An index satisfying the predicate was not found in the collection.")
            })
        }
        pub fn listsHadDifferentLengths() -> string {
            static listsHadDifferentLengths: MutCell<Option<string>> = MutCell::new(None);
            listsHadDifferentLengths.get_or_init(|| string("The lists had different lengths."))
        }
        pub fn mapCannotBeMutated() -> string {
            static mapCannotBeMutated: MutCell<Option<string>> = MutCell::new(None);
            mapCannotBeMutated.get_or_init(|| string("Map values cannot be mutated."))
        }
        pub fn notAPermutation() -> string {
            static notAPermutation: MutCell<Option<string>> = MutCell::new(None);
            notAPermutation.get_or_init(|| string("The function did not compute a permutation."))
        }
        pub fn notEnoughElements() -> string {
            static notEnoughElements: MutCell<Option<string>> = MutCell::new(None);
            notEnoughElements.get_or_init(|| {
                string("The input sequence has an insufficient number of elements.")
            })
        }
        pub fn outOfRange() -> string {
            static outOfRange: MutCell<Option<string>> = MutCell::new(None);
            outOfRange.get_or_init(|| string("The index is outside the legal range."))
        }
        pub fn resetNotSupported() -> string {
            static resetNotSupported: MutCell<Option<string>> = MutCell::new(None);
            resetNotSupported.get_or_init(|| string("Reset is not supported on this enumerator."))
        }
        pub fn setContainsNoElements() -> string {
            static setContainsNoElements: MutCell<Option<string>> = MutCell::new(None);
            setContainsNoElements.get_or_init(|| string("Set contains no elements."))
        }
    }
}
