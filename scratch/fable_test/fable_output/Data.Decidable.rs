pub mod PureScript_Data_Decidable {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_18b8108c::PureScript_Data_Comparison;
    use crate::module_abb3d73f::PureScript_Data_Decide;
    use crate::module_747333f6::PureScript_Data_Divisible;
    use crate::module_59dcd94b::PureScript_Data_Equivalence;
    use crate::module_851c93ca::PureScript_Data_Op;
    use crate::module_d23c04ec::PureScript_Data_Predicate;
    use crate::module_38c1e8e1::PureScript_Data_Void;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Decidable_identity() -> &dyn Any {
        static Data_Decidable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_identity.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Decidable_Decidableusd_Dict() -> &dyn Any {
        static Data_Decidable_Decidableusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_Decidableusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Decidable_lose() -> &dyn Any {
        static Data_Decidable_lose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_lose.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("lose"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Decidable_lost() -> &dyn Any {
        static Data_Decidable_lost: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_lost.get_or_init(||
                                            &Func1::new(move |dictDecidable|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decidable::Data_Decidable_lose(),
                                                                                                                                dictDecidable),
                                                                                             &&&PureScript_Data_Decidable::Data_Decidable_identity())))
    }
    pub fn Data_Decidable_decidablePredicate() -> &dyn Any {
        static Data_Decidable_decidablePredicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_decidablePredicate.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decidable::Data_Decidable_Decidableusd_Dict(),
                                                                                           &&&add(string("lose"),
                                                                                                  &&Func1::new(move
                                                                                                                   |f|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Predicate::Data_Predicate_Predicate(),
                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                      let f
                                                                                                                                                                          =
                                                                                                                                                                          f.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |a|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Void::Data_Void_absurd(),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                              a))
                                                                                                                                                                  }))),
                                                                                                  add(string("Decide0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Decide::Data_Decide_choosePredicate()),
                                                                                                      add(string("Divisible1"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused_1|
                                                                                                                           &PureScript_Data_Divisible::Data_Divisible_divisiblePredicate()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Decidable_decidableOp() -> &dyn Any {
        static Data_Decidable_decidableOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_decidableOp.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonoid|
                                                                   {
                                                                       let chooseOp =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_chooseOp(),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                       let divisibleOp =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divisible::Data_Divisible_divisibleOp(),
                                                                                                            dictMonoid);
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decidable::Data_Decidable_Decidableusd_Dict(),
                                                                                                        &&&add(string("lose"),
                                                                                                               &&Func1::new(move
                                                                                                                                |f|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Op::Data_Op_Op(),
                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                   let f
                                                                                                                                                                                       =
                                                                                                                                                                                       f.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |a|
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Void::Data_Void_absurd(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                           a))
                                                                                                                                                                               }))),
                                                                                                               add(string("Decide0"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let chooseOp
                                                                                                                                        =
                                                                                                                                        chooseOp.clone();
                                                                                                                                    move
                                                                                                                                        |usd__unused|
                                                                                                                                        &chooseOp
                                                                                                                                }),
                                                                                                                   add(string("Divisible1"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let divisibleOp
                                                                                                                                            =
                                                                                                                                            divisibleOp.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &divisibleOp
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))
                                                                   }))
    }
    pub fn Data_Decidable_decidableEquivalence() -> &dyn Any {
        static Data_Decidable_decidableEquivalence: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Decidable_decidableEquivalence.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decidable::Data_Decidable_Decidableusd_Dict(),
                                                                                             &&&add(string("lose"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                        let f
                                                                                                                                                                            =
                                                                                                                                                                            f.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |a|
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Void::Data_Void_absurd(),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                a))
                                                                                                                                                                    }))),
                                                                                                    add(string("Decide0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Decide::Data_Decide_chooseEquivalence()),
                                                                                                        add(string("Divisible1"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused_1|
                                                                                                                             &PureScript_Data_Divisible::Data_Divisible_divisibleEquivalence()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))))
    }
    pub fn Data_Decidable_decidableComparison() -> &dyn Any {
        static Data_Decidable_decidableComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decidable_decidableComparison.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decidable::Data_Decidable_Decidableusd_Dict(),
                                                                                            &&&add(string("lose"),
                                                                                                   &&Func1::new(move
                                                                                                                    |f|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                       let f
                                                                                                                                                                           =
                                                                                                                                                                           f.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |a|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let a
                                                                                                                                                                                               =
                                                                                                                                                                                               a.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |v|
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Void::Data_Void_absurd(),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                   &&&a))
                                                                                                                                                                                       })
                                                                                                                                                                   }))),
                                                                                                   add(string("Decide0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Decide::Data_Decide_chooseComparison()),
                                                                                                       add(string("Divisible1"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused_1|
                                                                                                                            &PureScript_Data_Divisible::Data_Divisible_divisibleComparison()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))))
    }
}
