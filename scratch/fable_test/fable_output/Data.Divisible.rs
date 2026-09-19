pub mod PureScript_Data_Divisible {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_18b8108c::PureScript_Data_Comparison;
    use crate::module_d8620aa6::PureScript_Data_Divide;
    use crate::module_59dcd94b::PureScript_Data_Equivalence;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_851c93ca::PureScript_Data_Op;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_d23c04ec::PureScript_Data_Predicate;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Divisible_Divisibleusd_Dict() -> &dyn Any {
        static Data_Divisible_Divisibleusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Divisible_Divisibleusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Divisible_divisiblePredicate() -> &dyn Any {
        static Data_Divisible_divisiblePredicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Divisible_divisiblePredicate.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divisible::Data_Divisible_Divisibleusd_Dict(),
                                                                                           &&&add(string("conquer"),
                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Predicate::Data_Predicate_Predicate(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                       &&&true)),
                                                                                                  add(string("Divide0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Divide::Data_Divide_dividePredicate()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Divisible_divisibleOp() -> &dyn Any {
        static Data_Divisible_divisibleOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Divisible_divisibleOp.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonoid|
                                                                   {
                                                                       let divideOp =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divide::Data_Divide_divideOp(),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divisible::Data_Divisible_Divisibleusd_Dict(),
                                                                                                        &&&add(string("conquer"),
                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                    &&&PureScript_Data_Op::Data_Op_Op()),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                       dictMonoid))),
                                                                                                               add(string("Divide0"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let divideOp
                                                                                                                                        =
                                                                                                                                        divideOp.clone();
                                                                                                                                    move
                                                                                                                                        |usd__unused|
                                                                                                                                        &divideOp
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>())))
                                                                   }))
    }
    pub fn Data_Divisible_divisibleEquivalence() -> &dyn Any {
        static Data_Divisible_divisibleEquivalence: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Divisible_divisibleEquivalence.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divisible::Data_Divisible_Divisibleusd_Dict(),
                                                                                             &&&add(string("conquer"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                         &&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence()),
                                                                                                                                      &&&Func1::new(move
                                                                                                                                                        |v|
                                                                                                                                                        &Func1::new(move
                                                                                                                                                                        |v1|
                                                                                                                                                                        &true))),
                                                                                                    add(string("Divide0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Divide::Data_Divide_divideEquivalence()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Divisible_divisibleComparison() -> &dyn Any {
        static Data_Divisible_divisibleComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Divisible_divisibleComparison.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divisible::Data_Divisible_Divisibleusd_Dict(),
                                                                                            &&&add(string("conquer"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                        &&&PureScript_Data_Comparison::Data_Comparison_Comparison()),
                                                                                                                                     &&&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       &Func1::new(move
                                                                                                                                                                       |v1|
                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)))),
                                                                                                   add(string("Divide0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Divide::Data_Divide_divideComparison()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Divisible_conquer() -> &dyn Any {
        static Data_Divisible_conquer: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Divisible_conquer.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("conquer"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
