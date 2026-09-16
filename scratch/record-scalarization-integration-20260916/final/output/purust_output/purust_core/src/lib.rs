#![allow(warnings)]

use perceus_ptr::PerceusPtr;

#[derive(Clone)]
pub enum Void {}

#[derive(Clone)]
pub enum Value {
    Unit,
    Null,
    Int(i64),
    Number(f64),
    Bool(bool),
    String(String),
    Char(char),
    Array(std::rc::Rc<Vec<UnknownType>>),
    Func1(Func1<UnknownType, UnknownType>),
    Func2(Func2<UnknownType, UnknownType, UnknownType>),
    Func3(Func3<UnknownType, UnknownType, UnknownType, UnknownType>),
    Func4(Func4<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func5(Func5<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func6(Func6<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func7(Func7<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func8(Func8<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func9(Func9<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func10(Func10<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func11(Func11<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func12(Func12<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Class(std::rc::Rc<dyn std::any::Any>),
    Thunk(perceus_ptr::PerceusPtr<Thunk>),
    Record_a(perceus_ptr::PerceusPtr<Record_a>),
    DynamicRecord(perceus_ptr::PerceusPtr<RecordFields>),
    Record_Alt0_empty(perceus_ptr::PerceusPtr<Record_Alt0_empty>),
    Record_Alternative1_Monad0(perceus_ptr::PerceusPtr<Record_Alternative1_Monad0>),
    Record_Applicative0_Bind1(perceus_ptr::PerceusPtr<Record_Applicative0_Bind1>),
    Record_Applicative0_Plus1(perceus_ptr::PerceusPtr<Record_Applicative0_Plus1>),
    Record_Apply0_Apply1_parallel_sequential(perceus_ptr::PerceusPtr<Record_Apply0_Apply1_parallel_sequential>),
    Record_Apply0_bind(perceus_ptr::PerceusPtr<Record_Apply0_bind>),
    Record_Apply0_pure(perceus_ptr::PerceusPtr<Record_Apply0_pure>),
    Record_Biapply0_bipure(perceus_ptr::PerceusPtr<Record_Biapply0_bipure>),
    Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(perceus_ptr::PerceusPtr<Record_Bifoldable1_Bifunctor0_bisequence_bitraverse>),
    Record_Bifunctor0_biapply(perceus_ptr::PerceusPtr<Record_Bifunctor0_biapply>),
    Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(perceus_ptr::PerceusPtr<Record_Bounded0_Enum1_cardinality_fromEnum_toEnum>),
    Record_Coercible0(perceus_ptr::PerceusPtr<Record_Coercible0>),
    Record_Coercible0_proof(perceus_ptr::PerceusPtr<Record_Coercible0_proof>),
    Record_CommutativeRing0_degree_div_mod_kw(perceus_ptr::PerceusPtr<Record_CommutativeRing0_degree_div_mod_kw>),
    Record_Comonad0_ask(perceus_ptr::PerceusPtr<Record_Comonad0_ask>),
    Record_Comonad0_peek_pos(perceus_ptr::PerceusPtr<Record_Comonad0_peek_pos>),
    Record_Comonad0_track(perceus_ptr::PerceusPtr<Record_Comonad0_track>),
    Record_ComonadAsk0_local(perceus_ptr::PerceusPtr<Record_ComonadAsk0_local>),
    Record_Contravariant0_divide(perceus_ptr::PerceusPtr<Record_Contravariant0_divide>),
    Record_Decide0_Divisible1_lose(perceus_ptr::PerceusPtr<Record_Decide0_Divisible1_lose>),
    Record_Divide0_choose(perceus_ptr::PerceusPtr<Record_Divide0_choose>),
    Record_Divide0_conquer(perceus_ptr::PerceusPtr<Record_Divide0_conquer>),
    Record_DivisionRing1_EuclideanRing0(perceus_ptr::PerceusPtr<Record_DivisionRing1_EuclideanRing0>),
    Record_Eq0_compare(perceus_ptr::PerceusPtr<Record_Eq0_compare>),
    Record_Eq10_compare1(perceus_ptr::PerceusPtr<Record_Eq10_compare1>),
    Record_EqRecord0_compareRecord(perceus_ptr::PerceusPtr<Record_EqRecord0_compareRecord>),
    Record_Extend0_extract(perceus_ptr::PerceusPtr<Record_Extend0_extract>),
    Record_Foldable0_foldMap1_foldl1_foldr1(perceus_ptr::PerceusPtr<Record_Foldable0_foldMap1_foldl1_foldr1>),
    Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(perceus_ptr::PerceusPtr<Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex>),
    Record_Foldable1_Functor0_sequence_traverse(perceus_ptr::PerceusPtr<Record_Foldable1_Functor0_sequence_traverse>),
    Record_Foldable10_Traversable1_sequence1_traverse1(perceus_ptr::PerceusPtr<Record_Foldable10_Traversable1_sequence1_traverse1>),
    Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(perceus_ptr::PerceusPtr<Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex>),
    Record_Functor0_alt(perceus_ptr::PerceusPtr<Record_Functor0_alt>),
    Record_Functor0_apply(perceus_ptr::PerceusPtr<Record_Functor0_apply>),
    Record_Functor0_collect_distribute(perceus_ptr::PerceusPtr<Record_Functor0_collect_distribute>),
    Record_Functor0_extend(perceus_ptr::PerceusPtr<Record_Functor0_extend>),
    Record_Functor0_mapWithIndex(perceus_ptr::PerceusPtr<Record_Functor0_mapWithIndex>),
    Record_HeytingAlgebra0(perceus_ptr::PerceusPtr<Record_HeytingAlgebra0>),
    Record_HeytingAlgebraRecord0(perceus_ptr::PerceusPtr<Record_HeytingAlgebraRecord0>),
    Record_Monad0_ask(perceus_ptr::PerceusPtr<Record_Monad0_ask>),
    Record_Monad0_callCC(perceus_ptr::PerceusPtr<Record_Monad0_callCC>),
    Record_Monad0_liftEffect(perceus_ptr::PerceusPtr<Record_Monad0_liftEffect>),
    Record_Monad0_liftST(perceus_ptr::PerceusPtr<Record_Monad0_liftST>),
    Record_Monad0_state(perceus_ptr::PerceusPtr<Record_Monad0_state>),
    Record_Monad0_tailRecM(perceus_ptr::PerceusPtr<Record_Monad0_tailRecM>),
    Record_Monad0_throwError(perceus_ptr::PerceusPtr<Record_Monad0_throwError>),
    Record_Monad1_Semigroup0_tell(perceus_ptr::PerceusPtr<Record_Monad1_Semigroup0_tell>),
    Record_MonadAsk0_local(perceus_ptr::PerceusPtr<Record_MonadAsk0_local>),
    Record_MonadEffect0_liftAff(perceus_ptr::PerceusPtr<Record_MonadEffect0_liftAff>),
    Record_MonadTell1_Monoid0_listen_pass(perceus_ptr::PerceusPtr<Record_MonadTell1_Monoid0_listen_pass>),
    Record_MonadThrow0_catchError(perceus_ptr::PerceusPtr<Record_MonadThrow0_catchError>),
    Record_Ord0_bottom_top(perceus_ptr::PerceusPtr<Record_Ord0_bottom_top>),
    Record_Ord0_pred_succ(perceus_ptr::PerceusPtr<Record_Ord0_pred_succ>),
    Record_OrdRecord0_bottomRecord_topRecord(perceus_ptr::PerceusPtr<Record_OrdRecord0_bottomRecord_topRecord>),
    Record_Profunctor0_closed(perceus_ptr::PerceusPtr<Record_Profunctor0_closed>),
    Record_Profunctor0_first_second(perceus_ptr::PerceusPtr<Record_Profunctor0_first_second>),
    Record_Profunctor0_left_right(perceus_ptr::PerceusPtr<Record_Profunctor0_left_right>),
    Record_Ring0(perceus_ptr::PerceusPtr<Record_Ring0>),
    Record_Ring0_recip(perceus_ptr::PerceusPtr<Record_Ring0_recip>),
    Record_RingRecord0(perceus_ptr::PerceusPtr<Record_RingRecord0>),
    Record_Semigroup0_mempty(perceus_ptr::PerceusPtr<Record_Semigroup0_mempty>),
    Record_SemigroupRecord0_memptyRecord(perceus_ptr::PerceusPtr<Record_SemigroupRecord0_memptyRecord>),
    Record_Semigroupoid0_identity(perceus_ptr::PerceusPtr<Record_Semigroupoid0_identity>),
    Record_Semiring0_sub(perceus_ptr::PerceusPtr<Record_Semiring0_sub>),
    Record_SemiringRecord0_subRecord(perceus_ptr::PerceusPtr<Record_SemiringRecord0_subRecord>),
    Record_Unfoldable10_unfoldr(perceus_ptr::PerceusPtr<Record_Unfoldable10_unfoldr>),
    Record_a_b(perceus_ptr::PerceusPtr<Record_a_b>),
    Record_a_b_c(perceus_ptr::PerceusPtr<Record_a_b_c>),
    Record_a_b_c_d_e(perceus_ptr::PerceusPtr<Record_a_b_c_d_e>),
    Record_acc_init(perceus_ptr::PerceusPtr<Record_acc_init>),
    Record_acc_val(perceus_ptr::PerceusPtr<Record_acc_val>),
    Record_accum_value(perceus_ptr::PerceusPtr<Record_accum_value>),
    Record_add_mul_one_zero(perceus_ptr::PerceusPtr<Record_add_mul_one_zero>),
    Record_addRecord_mulRecord_oneRecord_zeroRecord(perceus_ptr::PerceusPtr<Record_addRecord_mulRecord_oneRecord_zeroRecord>),
    Record_after_before(perceus_ptr::PerceusPtr<Record_after_before>),
    Record_append(perceus_ptr::PerceusPtr<Record_append>),
    Record_appendRecord(perceus_ptr::PerceusPtr<Record_appendRecord>),
    Record_asList_asMap(perceus_ptr::PerceusPtr<Record_asList_asMap>),
    Record_bifoldMap_bifoldl_bifoldr(perceus_ptr::PerceusPtr<Record_bifoldMap_bifoldl_bifoldr>),
    Record_bimap(perceus_ptr::PerceusPtr<Record_bimap>),
    Record_c_d(perceus_ptr::PerceusPtr<Record_c_d>),
    Record_cmap(perceus_ptr::PerceusPtr<Record_cmap>),
    Record_completed_failed_killed(perceus_ptr::PerceusPtr<Record_completed_failed_killed>),
    Record_compose(perceus_ptr::PerceusPtr<Record_compose>),
    Record_conj_disj_ff_implies_not_tt(perceus_ptr::PerceusPtr<Record_conj_disj_ff_implies_not_tt>),
    Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(perceus_ptr::PerceusPtr<Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord>),
    Record_day_hour_millisecond_minute_month_second_year(perceus_ptr::PerceusPtr<Record_day_hour_millisecond_minute_month_second_year>),
    Record_defer(perceus_ptr::PerceusPtr<Record_defer>),
    Record_dimap(perceus_ptr::PerceusPtr<Record_dimap>),
    Record_discard(perceus_ptr::PerceusPtr<Record_discard>),
    Record_dotAll_global_ignoreCase_multiline_sticky_unicode(perceus_ptr::PerceusPtr<Record_dotAll_global_ignoreCase_multiline_sticky_unicode>),
    Record_e_f(perceus_ptr::PerceusPtr<Record_e_f>),
    Record_elem_pos(perceus_ptr::PerceusPtr<Record_elem_pos>),
    Record_eq(perceus_ptr::PerceusPtr<Record_eq>),
    Record_eq1(perceus_ptr::PerceusPtr<Record_eq1>),
    Record_eqRecord(perceus_ptr::PerceusPtr<Record_eqRecord>),
    Record_fiber_supervisor(perceus_ptr::PerceusPtr<Record_fiber_supervisor>),
    Record_foldMap_foldl_foldr(perceus_ptr::PerceusPtr<Record_foldMap_foldl_foldr>),
    Record_found_result(perceus_ptr::PerceusPtr<Record_found_result>),
    Record_from_to(perceus_ptr::PerceusPtr<Record_from_to>),
    Record_fromDuration_toDuration(perceus_ptr::PerceusPtr<Record_fromDuration_toDuration>),
    Record_fromLeft_fromRight_isLeft_left_right(perceus_ptr::PerceusPtr<Record_fromLeft_fromRight_isLeft_left_right>),
    Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(perceus_ptr::PerceusPtr<Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime>),
    Record_genericAppend_prime(perceus_ptr::PerceusPtr<Record_genericAppend_prime>),
    Record_genericBottom_prime(perceus_ptr::PerceusPtr<Record_genericBottom_prime>),
    Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(perceus_ptr::PerceusPtr<Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime>),
    Record_genericCompare_prime(perceus_ptr::PerceusPtr<Record_genericCompare_prime>),
    Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(perceus_ptr::PerceusPtr<Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime>),
    Record_genericEq_prime(perceus_ptr::PerceusPtr<Record_genericEq_prime>),
    Record_genericMempty_prime(perceus_ptr::PerceusPtr<Record_genericMempty_prime>),
    Record_genericPred_prime_genericSucc_prime(perceus_ptr::PerceusPtr<Record_genericPred_prime_genericSucc_prime>),
    Record_genericShow_prime(perceus_ptr::PerceusPtr<Record_genericShow_prime>),
    Record_genericShowArgs(perceus_ptr::PerceusPtr<Record_genericShowArgs>),
    Record_genericSub_prime(perceus_ptr::PerceusPtr<Record_genericSub_prime>),
    Record_genericTop_prime(perceus_ptr::PerceusPtr<Record_genericTop_prime>),
    Record_handler_rethrow(perceus_ptr::PerceusPtr<Record_handler_rethrow>),
    Record_head_tail(perceus_ptr::PerceusPtr<Record_head_tail>),
    Record_imap(perceus_ptr::PerceusPtr<Record_imap>),
    Record_index_value(perceus_ptr::PerceusPtr<Record_index_value>),
    Record_init_last(perceus_ptr::PerceusPtr<Record_init_last>),
    Record_init_rest(perceus_ptr::PerceusPtr<Record_init_rest>),
    Record_inj_prj(perceus_ptr::PerceusPtr<Record_inj_prj>),
    Record_isSuspended_join_kill_onComplete_run(perceus_ptr::PerceusPtr<Record_isSuspended_join_kill_onComplete_run>),
    Record_key_value(perceus_ptr::PerceusPtr<Record_key_value>),
    Record_keysImpl(perceus_ptr::PerceusPtr<Record_keysImpl>),
    Record_last_revInit(perceus_ptr::PerceusPtr<Record_last_revInit>),
    Record_lift(perceus_ptr::PerceusPtr<Record_lift>),
    Record_lower(perceus_ptr::PerceusPtr<Record_lower>),
    Record_map(perceus_ptr::PerceusPtr<Record_map>),
    Record_mappend__mempty_(perceus_ptr::PerceusPtr<Record_mappend__mempty_>),
    Record_myMethod(perceus_ptr::PerceusPtr<Record_myMethod>),
    Record_nes(perceus_ptr::PerceusPtr<Record_nes>),
    Record_no_yes(perceus_ptr::PerceusPtr<Record_no_yes>),
    Record_ps_minus_rust_minus_test(perceus_ptr::PerceusPtr<Record_ps_minus_rust_minus_test>),
    Record_reflectSymbol(perceus_ptr::PerceusPtr<Record_reflectSymbol>),
    Record_reflectType(perceus_ptr::PerceusPtr<Record_reflectType>),
    Record_show(perceus_ptr::PerceusPtr<Record_show>),
    Record_showRecordFields(perceus_ptr::PerceusPtr<Record_showRecordFields>),
    Record_state_val(perceus_ptr::PerceusPtr<Record_state_val>),
    Record_state_value(perceus_ptr::PerceusPtr<Record_state_value>),
    Record_unfoldr1(perceus_ptr::PerceusPtr<Record_unfoldr1>),
}

impl Value {
    pub fn resolve(&self) -> &Self {
        let mut value = self;
        while let Value::Thunk(thunk) = value {
            value = thunk.value.get().expect("recursive value used before initialization");
        }
        value
    }
    pub fn unwrap_unit(&self) {
        if !matches!(self.resolve(), Value::Unit) { panic!("Expected Unit"); }
    }
    pub fn unwrap_int(&self) -> i64 {
        if let Value::Int(v) = self.resolve() { *v } else { panic!("Expected Int"); }
    }
    pub fn unwrap_number(&self) -> f64 {
        // Foreign numbers can originate from a native PureScript Int.
        match self.resolve() { Value::Number(v) => *v, Value::Int(v) => *v as f64, _ => panic!("Expected Number") }
    }
    pub fn unwrap_bool(&self) -> bool {
        if let Value::Bool(v) = self.resolve() { *v } else { panic!("Expected Bool"); }
    }
    pub fn unwrap_string(&self) -> String {
        if let Value::String(v) = self.resolve() { v.clone() } else { panic!("Expected String"); }
    }
    pub fn unwrap_char(&self) -> char {
        if let Value::Char(v) = self.resolve() { *v } else { panic!("Expected Char"); }
    }
    pub fn unwrap_array(&self) -> std::rc::Rc<Vec<UnknownType>> {
        if let Value::Array(v) = self.resolve() { v.clone() } else { panic!("Expected Array"); }
    }
    pub fn unwrap_func1(&self) -> Func1<UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func1(v) = value { v.clone() } else if let Value::Record_a(v) = value { v.call.clone().unwrap() } else if let Value::Func2(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func1(Func1::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType| -> UnknownType { f2(a0.clone(), a1) } }))) })) } else if let Value::Func3(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func2(Func2::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2) } }))) })) } else if let Value::Func4(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func3(Func3::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3) } }))) })) } else if let Value::Func5(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func4(Func4::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4) } }))) })) } else if let Value::Func6(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func5(Func5::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5) } }))) })) } else if let Value::Func7(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func6(Func6::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6) } }))) })) } else if let Value::Func8(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func7(Func7::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7) } }))) })) } else if let Value::Func9(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func8(Func8::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8) } }))) })) } else if let Value::Func10(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func9(Func9::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9) } }))) })) } else if let Value::Func11(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func10(Func10::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9, a10) } }))) })) } else if let Value::Func12(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func11(Func11::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType, mut a11: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11) } }))) })) } else { panic!("Expected Func1"); }
    }
    pub fn unwrap_func2(&self) -> Func2<UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func2(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func2::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1) })) } else { panic!("Expected Func2 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func3(&self) -> Func3<UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func3(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func3::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2) })) } else { panic!("Expected Func3 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func4(&self) -> Func4<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func4(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func4::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3) })) } else { panic!("Expected Func4 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func5(&self) -> Func5<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func5(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func5::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4) })) } else { panic!("Expected Func5 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func6(&self) -> Func6<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func6(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func6::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5) })) } else { panic!("Expected Func6 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func7(&self) -> Func7<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func7(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func7::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6) })) } else { panic!("Expected Func7 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func8(&self) -> Func8<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func8(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func8::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7) })) } else { panic!("Expected Func8 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func9(&self) -> Func9<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func9(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func9::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8) })) } else { panic!("Expected Func9 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func10(&self) -> Func10<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func10(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func10::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9) })) } else { panic!("Expected Func10 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func11(&self) -> Func11<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func11(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func11::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9).unwrap_func1()(a10) })) } else { panic!("Expected Func11 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func12(&self) -> Func12<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func12(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func12::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType, mut a11: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9).unwrap_func1()(a10).unwrap_func1()(a11) })) } else { panic!("Expected Func12 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_class<T: 'static>(&self) -> &T {
        if let Value::Class(v) = self.resolve() { v.downcast_ref::<T>().unwrap() } else { panic!("Expected Class"); }
    }
    pub fn drop_explicit(self) {
    }
    pub fn __purust_ctor_tag(&self) -> &'static str {
        if let Value::Record_a(r) = self.resolve() { r.tag } else { panic!("Expected Record_a for tag"); }
    }
    pub fn get_Alt0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Alt0_empty(r) => r.Alt0.clone().unwrap(),
            Value::Record_a(r) => r.Alt0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Alt0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Alt0"),
        }
    }
    pub fn get_Alternative1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Alternative1_Monad0(r) => r.Alternative1.clone().unwrap(),
            Value::Record_a(r) => r.Alternative1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Alternative1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Alternative1"),
        }
    }
    pub fn get_Applicative0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Bind1(r) => r.Applicative0.clone().unwrap(),
            Value::Record_Applicative0_Plus1(r) => r.Applicative0.clone().unwrap(),
            Value::Record_a(r) => r.Applicative0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Applicative0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Applicative0"),
        }
    }
    pub fn get_Apply0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.Apply0.clone().unwrap(),
            Value::Record_Apply0_bind(r) => r.Apply0.clone().unwrap(),
            Value::Record_Apply0_pure(r) => r.Apply0.clone().unwrap(),
            Value::Record_a(r) => r.Apply0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Apply0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Apply0"),
        }
    }
    pub fn get_Apply1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.Apply1.clone().unwrap(),
            Value::Record_a(r) => r.Apply1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Apply1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Apply1"),
        }
    }
    pub fn get_Biapply0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Biapply0_bipure(r) => r.Biapply0.clone().unwrap(),
            Value::Record_a(r) => r.Biapply0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Biapply0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Biapply0"),
        }
    }
    pub fn get_Bifoldable1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.Bifoldable1.clone().unwrap(),
            Value::Record_a(r) => r.Bifoldable1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Bifoldable1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Bifoldable1"),
        }
    }
    pub fn get_Bifunctor0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.Bifunctor0.clone().unwrap(),
            Value::Record_Bifunctor0_biapply(r) => r.Bifunctor0.clone().unwrap(),
            Value::Record_a(r) => r.Bifunctor0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Bifunctor0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Bifunctor0"),
        }
    }
    pub fn get_Bind1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Bind1(r) => r.Bind1.clone().unwrap(),
            Value::Record_a(r) => r.Bind1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Bind1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Bind1"),
        }
    }
    pub fn get_Bounded0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.Bounded0.clone().unwrap(),
            Value::Record_a(r) => r.Bounded0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Bounded0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Bounded0"),
        }
    }
    pub fn get_Coercible0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Coercible0(r) => r.Coercible0.clone().unwrap(),
            Value::Record_Coercible0_proof(r) => r.Coercible0.clone().unwrap(),
            Value::Record_a(r) => r.Coercible0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Coercible0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Coercible0"),
        }
    }
    pub fn get_CommutativeRing0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.CommutativeRing0.clone().unwrap(),
            Value::Record_a(r) => r.CommutativeRing0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("CommutativeRing0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field CommutativeRing0"),
        }
    }
    pub fn get_Comonad0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_ask(r) => r.Comonad0.clone().unwrap(),
            Value::Record_Comonad0_peek_pos(r) => r.Comonad0.clone().unwrap(),
            Value::Record_Comonad0_track(r) => r.Comonad0.clone().unwrap(),
            Value::Record_a(r) => r.Comonad0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Comonad0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Comonad0"),
        }
    }
    pub fn get_ComonadAsk0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_ComonadAsk0_local(r) => r.ComonadAsk0.clone().unwrap(),
            Value::Record_a(r) => r.ComonadAsk0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ComonadAsk0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ComonadAsk0"),
        }
    }
    pub fn get_Contravariant0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Contravariant0_divide(r) => r.Contravariant0.clone().unwrap(),
            Value::Record_a(r) => r.Contravariant0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Contravariant0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Contravariant0"),
        }
    }
    pub fn get_Decide0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.Decide0.clone().unwrap(),
            Value::Record_a(r) => r.Decide0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Decide0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Decide0"),
        }
    }
    pub fn get_Divide0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Divide0_choose(r) => r.Divide0.clone().unwrap(),
            Value::Record_Divide0_conquer(r) => r.Divide0.clone().unwrap(),
            Value::Record_a(r) => r.Divide0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Divide0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Divide0"),
        }
    }
    pub fn get_Divisible1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.Divisible1.clone().unwrap(),
            Value::Record_a(r) => r.Divisible1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Divisible1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Divisible1"),
        }
    }
    pub fn get_DivisionRing1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_DivisionRing1_EuclideanRing0(r) => r.DivisionRing1.clone().unwrap(),
            Value::Record_a(r) => r.DivisionRing1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("DivisionRing1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field DivisionRing1"),
        }
    }
    pub fn get_Enum1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.Enum1.clone().unwrap(),
            Value::Record_a(r) => r.Enum1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Enum1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Enum1"),
        }
    }
    pub fn get_Eq0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Eq0_compare(r) => r.Eq0.clone().unwrap(),
            Value::Record_a(r) => r.Eq0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Eq0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Eq0"),
        }
    }
    pub fn get_Eq10(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Eq10_compare1(r) => r.Eq10.clone().unwrap(),
            Value::Record_a(r) => r.Eq10.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Eq10").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Eq10"),
        }
    }
    pub fn get_EqRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_EqRecord0_compareRecord(r) => r.EqRecord0.clone().unwrap(),
            Value::Record_a(r) => r.EqRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("EqRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field EqRecord0"),
        }
    }
    pub fn get_EuclideanRing0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_DivisionRing1_EuclideanRing0(r) => r.EuclideanRing0.clone().unwrap(),
            Value::Record_a(r) => r.EuclideanRing0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("EuclideanRing0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field EuclideanRing0"),
        }
    }
    pub fn get_Extend0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Extend0_extract(r) => r.Extend0.clone().unwrap(),
            Value::Record_a(r) => r.Extend0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Extend0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Extend0"),
        }
    }
    pub fn get_Foldable0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.Foldable0.clone().unwrap(),
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.Foldable0.clone().unwrap(),
            Value::Record_a(r) => r.Foldable0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Foldable0"),
        }
    }
    pub fn get_Foldable1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.Foldable1.clone().unwrap(),
            Value::Record_a(r) => r.Foldable1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Foldable1"),
        }
    }
    pub fn get_Foldable10(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.Foldable10.clone().unwrap(),
            Value::Record_a(r) => r.Foldable10.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable10").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Foldable10"),
        }
    }
    pub fn get_FoldableWithIndex1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.FoldableWithIndex1.clone().unwrap(),
            Value::Record_a(r) => r.FoldableWithIndex1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("FoldableWithIndex1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field FoldableWithIndex1"),
        }
    }
    pub fn get_Functor0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.Functor0.clone().unwrap(),
            Value::Record_Functor0_alt(r) => r.Functor0.clone().unwrap(),
            Value::Record_Functor0_apply(r) => r.Functor0.clone().unwrap(),
            Value::Record_Functor0_collect_distribute(r) => r.Functor0.clone().unwrap(),
            Value::Record_Functor0_extend(r) => r.Functor0.clone().unwrap(),
            Value::Record_Functor0_mapWithIndex(r) => r.Functor0.clone().unwrap(),
            Value::Record_a(r) => r.Functor0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Functor0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Functor0"),
        }
    }
    pub fn get_FunctorWithIndex0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.FunctorWithIndex0.clone().unwrap(),
            Value::Record_a(r) => r.FunctorWithIndex0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("FunctorWithIndex0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field FunctorWithIndex0"),
        }
    }
    pub fn get_HeytingAlgebra0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_HeytingAlgebra0(r) => r.HeytingAlgebra0.clone().unwrap(),
            Value::Record_a(r) => r.HeytingAlgebra0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("HeytingAlgebra0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field HeytingAlgebra0"),
        }
    }
    pub fn get_HeytingAlgebraRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_HeytingAlgebraRecord0(r) => r.HeytingAlgebraRecord0.clone().unwrap(),
            Value::Record_a(r) => r.HeytingAlgebraRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("HeytingAlgebraRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field HeytingAlgebraRecord0"),
        }
    }
    pub fn get_Monad0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Alternative1_Monad0(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_ask(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_callCC(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_liftEffect(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_liftST(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_state(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_tailRecM(r) => r.Monad0.clone().unwrap(),
            Value::Record_Monad0_throwError(r) => r.Monad0.clone().unwrap(),
            Value::Record_a(r) => r.Monad0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Monad0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Monad0"),
        }
    }
    pub fn get_Monad1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.Monad1.clone().unwrap(),
            Value::Record_a(r) => r.Monad1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Monad1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Monad1"),
        }
    }
    pub fn get_MonadAsk0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadAsk0_local(r) => r.MonadAsk0.clone().unwrap(),
            Value::Record_a(r) => r.MonadAsk0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadAsk0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field MonadAsk0"),
        }
    }
    pub fn get_MonadEffect0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadEffect0_liftAff(r) => r.MonadEffect0.clone().unwrap(),
            Value::Record_a(r) => r.MonadEffect0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadEffect0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field MonadEffect0"),
        }
    }
    pub fn get_MonadTell1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.MonadTell1.clone().unwrap(),
            Value::Record_a(r) => r.MonadTell1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadTell1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field MonadTell1"),
        }
    }
    pub fn get_MonadThrow0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadThrow0_catchError(r) => r.MonadThrow0.clone().unwrap(),
            Value::Record_a(r) => r.MonadThrow0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadThrow0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field MonadThrow0"),
        }
    }
    pub fn get_Monoid0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.Monoid0.clone().unwrap(),
            Value::Record_a(r) => r.Monoid0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Monoid0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Monoid0"),
        }
    }
    pub fn get_Ord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.Ord0.clone().unwrap(),
            Value::Record_Ord0_pred_succ(r) => r.Ord0.clone().unwrap(),
            Value::Record_a(r) => r.Ord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Ord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Ord0"),
        }
    }
    pub fn get_OrdRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.OrdRecord0.clone().unwrap(),
            Value::Record_a(r) => r.OrdRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("OrdRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field OrdRecord0"),
        }
    }
    pub fn get_Plus1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Plus1(r) => r.Plus1.clone().unwrap(),
            Value::Record_a(r) => r.Plus1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Plus1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Plus1"),
        }
    }
    pub fn get_Profunctor0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_closed(r) => r.Profunctor0.clone().unwrap(),
            Value::Record_Profunctor0_first_second(r) => r.Profunctor0.clone().unwrap(),
            Value::Record_Profunctor0_left_right(r) => r.Profunctor0.clone().unwrap(),
            Value::Record_a(r) => r.Profunctor0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Profunctor0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Profunctor0"),
        }
    }
    pub fn get_Ring0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ring0(r) => r.Ring0.clone().unwrap(),
            Value::Record_Ring0_recip(r) => r.Ring0.clone().unwrap(),
            Value::Record_a(r) => r.Ring0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Ring0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Ring0"),
        }
    }
    pub fn get_RingRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_RingRecord0(r) => r.RingRecord0.clone().unwrap(),
            Value::Record_a(r) => r.RingRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("RingRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field RingRecord0"),
        }
    }
    pub fn get_Semigroup0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.Semigroup0.clone().unwrap(),
            Value::Record_Semigroup0_mempty(r) => r.Semigroup0.clone().unwrap(),
            Value::Record_a(r) => r.Semigroup0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Semigroup0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Semigroup0"),
        }
    }
    pub fn get_SemigroupRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_SemigroupRecord0_memptyRecord(r) => r.SemigroupRecord0.clone().unwrap(),
            Value::Record_a(r) => r.SemigroupRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("SemigroupRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field SemigroupRecord0"),
        }
    }
    pub fn get_Semigroupoid0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Semigroupoid0_identity(r) => r.Semigroupoid0.clone().unwrap(),
            Value::Record_a(r) => r.Semigroupoid0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Semigroupoid0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Semigroupoid0"),
        }
    }
    pub fn get_Semiring0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Semiring0_sub(r) => r.Semiring0.clone().unwrap(),
            Value::Record_a(r) => r.Semiring0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Semiring0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Semiring0"),
        }
    }
    pub fn get_SemiringRecord0(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_SemiringRecord0_subRecord(r) => r.SemiringRecord0.clone().unwrap(),
            Value::Record_a(r) => r.SemiringRecord0.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("SemiringRecord0").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field SemiringRecord0"),
        }
    }
    pub fn get_Traversable1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.Traversable1.clone().unwrap(),
            Value::Record_a(r) => r.Traversable1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Traversable1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Traversable1"),
        }
    }
    pub fn get_Traversable2(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.Traversable2.clone().unwrap(),
            Value::Record_a(r) => r.Traversable2.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Traversable2").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Traversable2"),
        }
    }
    pub fn get_Unfoldable10(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Unfoldable10_unfoldr(r) => r.Unfoldable10.clone().unwrap(),
            Value::Record_a(r) => r.Unfoldable10.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("Unfoldable10").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field Unfoldable10"),
        }
    }
    pub fn get_a(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b(r) => r.a.clone().unwrap(),
            Value::Record_a_b_c(r) => r.a.clone().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.a.clone().unwrap(),
            Value::Record_a(r) => r.a.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("a").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn get_acc(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_acc_init(r) => r.acc.clone().unwrap(),
            Value::Record_acc_val(r) => r.acc.clone().unwrap(),
            Value::Record_a(r) => r.acc.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("acc").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field acc"),
        }
    }
    pub fn get_accum(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_accum_value(r) => r.accum.clone().unwrap(),
            Value::Record_a(r) => r.accum.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("accum").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field accum"),
        }
    }
    pub fn get_add(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.add.clone().unwrap(),
            Value::Record_a(r) => r.add.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("add").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field add"),
        }
    }
    pub fn get_addRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.addRecord.clone().unwrap(),
            Value::Record_a(r) => r.addRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("addRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field addRecord"),
        }
    }
    pub fn get_after(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_after_before(r) => r.after.clone().unwrap(),
            Value::Record_a(r) => r.after.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("after").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field after"),
        }
    }
    pub fn get_alt(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_alt(r) => r.alt.clone().unwrap(),
            Value::Record_a(r) => r.alt.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("alt").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field alt"),
        }
    }
    pub fn get_append(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_append(r) => r.append.clone().unwrap(),
            Value::Record_a(r) => r.append.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("append").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field append"),
        }
    }
    pub fn get_appendRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_appendRecord(r) => r.appendRecord.clone().unwrap(),
            Value::Record_a(r) => r.appendRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("appendRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field appendRecord"),
        }
    }
    pub fn get_apply(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_apply(r) => r.apply.clone().unwrap(),
            Value::Record_a(r) => r.apply.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("apply").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field apply"),
        }
    }
    pub fn get_asList(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_asList_asMap(r) => r.asList.clone().unwrap(),
            Value::Record_a(r) => r.asList.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("asList").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field asList"),
        }
    }
    pub fn get_asMap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_asList_asMap(r) => r.asMap.clone().unwrap(),
            Value::Record_a(r) => r.asMap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("asMap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field asMap"),
        }
    }
    pub fn get_ask(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_ask(r) => r.ask.clone().unwrap(),
            Value::Record_Monad0_ask(r) => r.ask.clone().unwrap(),
            Value::Record_a(r) => r.ask.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ask").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ask"),
        }
    }
    pub fn get_b(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b(r) => r.b.clone().unwrap(),
            Value::Record_a_b_c(r) => r.b.clone().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.b.clone().unwrap(),
            Value::Record_a(r) => r.b.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("b").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn get_before(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_after_before(r) => r.before.clone().unwrap(),
            Value::Record_a(r) => r.before.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("before").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field before"),
        }
    }
    pub fn get_biapply(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bifunctor0_biapply(r) => r.biapply.clone().unwrap(),
            Value::Record_a(r) => r.biapply.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("biapply").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field biapply"),
        }
    }
    pub fn get_bifoldMap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldMap.clone().unwrap(),
            Value::Record_a(r) => r.bifoldMap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldMap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bifoldMap"),
        }
    }
    pub fn get_bifoldl(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldl.clone().unwrap(),
            Value::Record_a(r) => r.bifoldl.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldl").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bifoldl"),
        }
    }
    pub fn get_bifoldr(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldr.clone().unwrap(),
            Value::Record_a(r) => r.bifoldr.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldr").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bifoldr"),
        }
    }
    pub fn get_bimap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_bimap(r) => r.bimap.clone().unwrap(),
            Value::Record_a(r) => r.bimap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bimap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bimap"),
        }
    }
    pub fn get_bind(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_bind(r) => r.bind.clone().unwrap(),
            Value::Record_a(r) => r.bind.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bind").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bind"),
        }
    }
    pub fn get_bipure(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Biapply0_bipure(r) => r.bipure.clone().unwrap(),
            Value::Record_a(r) => r.bipure.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bipure").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bipure"),
        }
    }
    pub fn get_bisequence(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.bisequence.clone().unwrap(),
            Value::Record_a(r) => r.bisequence.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bisequence").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bisequence"),
        }
    }
    pub fn get_bitraverse(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.bitraverse.clone().unwrap(),
            Value::Record_a(r) => r.bitraverse.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bitraverse").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bitraverse"),
        }
    }
    pub fn get_bottom(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.bottom.clone().unwrap(),
            Value::Record_a(r) => r.bottom.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bottom").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bottom"),
        }
    }
    pub fn get_bottomRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.bottomRecord.clone().unwrap(),
            Value::Record_a(r) => r.bottomRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("bottomRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field bottomRecord"),
        }
    }
    pub fn get_c(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_c(r) => r.c.clone().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.c.clone().unwrap(),
            Value::Record_c_d(r) => r.c.clone().unwrap(),
            Value::Record_a(r) => r.c.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("c").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn get_callCC(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_callCC(r) => r.callCC.clone().unwrap(),
            Value::Record_a(r) => r.callCC.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("callCC").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field callCC"),
        }
    }
    pub fn get_cardinality(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.cardinality.clone().unwrap(),
            Value::Record_a(r) => r.cardinality.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("cardinality").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field cardinality"),
        }
    }
    pub fn get_catchError(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadThrow0_catchError(r) => r.catchError.clone().unwrap(),
            Value::Record_a(r) => r.catchError.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("catchError").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field catchError"),
        }
    }
    pub fn get_choose(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Divide0_choose(r) => r.choose.clone().unwrap(),
            Value::Record_a(r) => r.choose.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("choose").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field choose"),
        }
    }
    pub fn get_closed(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_closed(r) => r.closed.clone().unwrap(),
            Value::Record_a(r) => r.closed.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("closed").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field closed"),
        }
    }
    pub fn get_cmap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_cmap(r) => r.cmap.clone().unwrap(),
            Value::Record_a(r) => r.cmap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("cmap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field cmap"),
        }
    }
    pub fn get_collect(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_collect_distribute(r) => r.collect.clone().unwrap(),
            Value::Record_a(r) => r.collect.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("collect").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field collect"),
        }
    }
    pub fn get_compare(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Eq0_compare(r) => r.compare.clone().unwrap(),
            Value::Record_a(r) => r.compare.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("compare").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field compare"),
        }
    }
    pub fn get_compare1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Eq10_compare1(r) => r.compare1.clone().unwrap(),
            Value::Record_a(r) => r.compare1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("compare1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field compare1"),
        }
    }
    pub fn get_compareRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_EqRecord0_compareRecord(r) => r.compareRecord.clone().unwrap(),
            Value::Record_a(r) => r.compareRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("compareRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field compareRecord"),
        }
    }
    pub fn get_completed(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.completed.clone().unwrap(),
            Value::Record_a(r) => r.completed.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("completed").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field completed"),
        }
    }
    pub fn get_compose(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_compose(r) => r.compose.clone().unwrap(),
            Value::Record_a(r) => r.compose.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("compose").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field compose"),
        }
    }
    pub fn get_conj(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.conj.clone().unwrap(),
            Value::Record_a(r) => r.conj.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("conj").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field conj"),
        }
    }
    pub fn get_conjRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.conjRecord.clone().unwrap(),
            Value::Record_a(r) => r.conjRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("conjRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field conjRecord"),
        }
    }
    pub fn get_conquer(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Divide0_conquer(r) => r.conquer.clone().unwrap(),
            Value::Record_a(r) => r.conquer.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("conquer").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field conquer"),
        }
    }
    pub fn get_d(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_c_d_e(r) => r.d.clone().unwrap(),
            Value::Record_c_d(r) => r.d.clone().unwrap(),
            Value::Record_a(r) => r.d.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("d").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn get_day(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.day.clone().unwrap(),
            Value::Record_a(r) => r.day.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("day").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field day"),
        }
    }
    pub fn get_defer(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_defer(r) => r.defer.clone().unwrap(),
            Value::Record_a(r) => r.defer.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("defer").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field defer"),
        }
    }
    pub fn get_degree(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.degree.clone().unwrap(),
            Value::Record_a(r) => r.degree.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("degree").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field degree"),
        }
    }
    pub fn get_dimap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dimap(r) => r.dimap.clone().unwrap(),
            Value::Record_a(r) => r.dimap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("dimap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field dimap"),
        }
    }
    pub fn get_discard(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_discard(r) => r.discard.clone().unwrap(),
            Value::Record_a(r) => r.discard.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("discard").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field discard"),
        }
    }
    pub fn get_disj(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.disj.clone().unwrap(),
            Value::Record_a(r) => r.disj.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("disj").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field disj"),
        }
    }
    pub fn get_disjRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.disjRecord.clone().unwrap(),
            Value::Record_a(r) => r.disjRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("disjRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field disjRecord"),
        }
    }
    pub fn get_distribute(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_collect_distribute(r) => r.distribute.clone().unwrap(),
            Value::Record_a(r) => r.distribute.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("distribute").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field distribute"),
        }
    }
    pub fn get_div(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.div.clone().unwrap(),
            Value::Record_a(r) => r.div.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("div").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field div"),
        }
    }
    pub fn get_divide(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Contravariant0_divide(r) => r.divide.clone().unwrap(),
            Value::Record_a(r) => r.divide.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("divide").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field divide"),
        }
    }
    pub fn get_dotAll(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.dotAll.clone().unwrap(),
            Value::Record_a(r) => r.dotAll.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("dotAll").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field dotAll"),
        }
    }
    pub fn get_e(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_c_d_e(r) => r.e.clone().unwrap(),
            Value::Record_e_f(r) => r.e.clone().unwrap(),
            Value::Record_a(r) => r.e.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("e").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn get_elem(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_elem_pos(r) => r.elem.clone().unwrap(),
            Value::Record_a(r) => r.elem.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("elem").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field elem"),
        }
    }
    pub fn get_empty(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Alt0_empty(r) => r.empty.clone().unwrap(),
            Value::Record_a(r) => r.empty.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("empty").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field empty"),
        }
    }
    pub fn get_eq(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_eq(r) => r.eq.clone().unwrap(),
            Value::Record_a(r) => r.eq.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("eq").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field eq"),
        }
    }
    pub fn get_eq1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_eq1(r) => r.eq1.clone().unwrap(),
            Value::Record_a(r) => r.eq1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("eq1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field eq1"),
        }
    }
    pub fn get_eqRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_eqRecord(r) => r.eqRecord.clone().unwrap(),
            Value::Record_a(r) => r.eqRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("eqRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field eqRecord"),
        }
    }
    pub fn get_extend(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_extend(r) => r.extend.clone().unwrap(),
            Value::Record_a(r) => r.extend.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("extend").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field extend"),
        }
    }
    pub fn get_extract(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Extend0_extract(r) => r.extract.clone().unwrap(),
            Value::Record_a(r) => r.extract.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("extract").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field extract"),
        }
    }
    pub fn get_f(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.f.clone().unwrap(),
            Value::Record_a(r) => r.f.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("f").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn get_failed(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.failed.clone().unwrap(),
            Value::Record_a(r) => r.failed.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("failed").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field failed"),
        }
    }
    pub fn get_ff(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.ff.clone().unwrap(),
            Value::Record_a(r) => r.ff.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ff").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ff"),
        }
    }
    pub fn get_ffRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.ffRecord.clone().unwrap(),
            Value::Record_a(r) => r.ffRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ffRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ffRecord"),
        }
    }
    pub fn get_fiber(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fiber_supervisor(r) => r.fiber.clone().unwrap(),
            Value::Record_a(r) => r.fiber.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("fiber").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field fiber"),
        }
    }
    pub fn get_first(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_first_second(r) => r.first.clone().unwrap(),
            Value::Record_a(r) => r.first.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("first").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field first"),
        }
    }
    pub fn get_foldMap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldMap.clone().unwrap(),
            Value::Record_a(r) => r.foldMap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldMap"),
        }
    }
    pub fn get_foldMap1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldMap1.clone().unwrap(),
            Value::Record_a(r) => r.foldMap1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMap1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldMap1"),
        }
    }
    pub fn get_foldMapWithIndex(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldMapWithIndex.clone().unwrap(),
            Value::Record_a(r) => r.foldMapWithIndex.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMapWithIndex").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldMapWithIndex"),
        }
    }
    pub fn get_foldl(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldl.clone().unwrap(),
            Value::Record_a(r) => r.foldl.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldl").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldl"),
        }
    }
    pub fn get_foldl1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldl1.clone().unwrap(),
            Value::Record_a(r) => r.foldl1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldl1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldl1"),
        }
    }
    pub fn get_foldlWithIndex(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldlWithIndex.clone().unwrap(),
            Value::Record_a(r) => r.foldlWithIndex.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldlWithIndex").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldlWithIndex"),
        }
    }
    pub fn get_foldr(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldr.clone().unwrap(),
            Value::Record_a(r) => r.foldr.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldr").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldr"),
        }
    }
    pub fn get_foldr1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldr1.clone().unwrap(),
            Value::Record_a(r) => r.foldr1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldr1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldr1"),
        }
    }
    pub fn get_foldrWithIndex(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldrWithIndex.clone().unwrap(),
            Value::Record_a(r) => r.foldrWithIndex.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("foldrWithIndex").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field foldrWithIndex"),
        }
    }
    pub fn get_found(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_found_result(r) => r.found.clone().unwrap(),
            Value::Record_a(r) => r.found.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("found").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field found"),
        }
    }
    pub fn get_from(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_from_to(r) => r.from.clone().unwrap(),
            Value::Record_a(r) => r.from.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("from").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field from"),
        }
    }
    pub fn get_fromDuration(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fromDuration_toDuration(r) => r.fromDuration.clone().unwrap(),
            Value::Record_a(r) => r.fromDuration.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("fromDuration").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field fromDuration"),
        }
    }
    pub fn get_fromEnum(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.fromEnum.clone().unwrap(),
            Value::Record_a(r) => r.fromEnum.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("fromEnum").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field fromEnum"),
        }
    }
    pub fn get_fromLeft(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.fromLeft.clone().unwrap(),
            Value::Record_a(r) => r.fromLeft.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("fromLeft").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field fromLeft"),
        }
    }
    pub fn get_fromRight(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.fromRight.clone().unwrap(),
            Value::Record_a(r) => r.fromRight.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("fromRight").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field fromRight"),
        }
    }
    pub fn get_genericAdd_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericAdd_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericAdd_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericAdd'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericAdd_prime"),
        }
    }
    pub fn get_genericAppend_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericAppend_prime(r) => r.genericAppend_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericAppend_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericAppend'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericAppend_prime"),
        }
    }
    pub fn get_genericBottom_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericBottom_prime(r) => r.genericBottom_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericBottom_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericBottom'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericBottom_prime"),
        }
    }
    pub fn get_genericCardinality_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericCardinality_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericCardinality_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericCardinality'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericCardinality_prime"),
        }
    }
    pub fn get_genericCompare_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericCompare_prime(r) => r.genericCompare_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericCompare_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericCompare'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericCompare_prime"),
        }
    }
    pub fn get_genericConj_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericConj_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericConj_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericConj'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericConj_prime"),
        }
    }
    pub fn get_genericDisj_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericDisj_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericDisj_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericDisj'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericDisj_prime"),
        }
    }
    pub fn get_genericEq_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericEq_prime(r) => r.genericEq_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericEq_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericEq'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericEq_prime"),
        }
    }
    pub fn get_genericFF_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericFF_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericFF_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericFF'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericFF_prime"),
        }
    }
    pub fn get_genericFromEnum_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericFromEnum_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericFromEnum_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericFromEnum'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericFromEnum_prime"),
        }
    }
    pub fn get_genericImplies_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericImplies_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericImplies_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericImplies'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericImplies_prime"),
        }
    }
    pub fn get_genericMempty_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericMempty_prime(r) => r.genericMempty_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericMempty_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericMempty'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericMempty_prime"),
        }
    }
    pub fn get_genericMul_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericMul_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericMul_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericMul'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericMul_prime"),
        }
    }
    pub fn get_genericNot_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericNot_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericNot_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericNot'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericNot_prime"),
        }
    }
    pub fn get_genericOne_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericOne_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericOne_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericOne'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericOne_prime"),
        }
    }
    pub fn get_genericPred_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericPred_prime_genericSucc_prime(r) => r.genericPred_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericPred_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericPred'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericPred_prime"),
        }
    }
    pub fn get_genericShow_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericShow_prime(r) => r.genericShow_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericShow_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericShow'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericShow_prime"),
        }
    }
    pub fn get_genericShowArgs(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericShowArgs(r) => r.genericShowArgs.clone().unwrap(),
            Value::Record_a(r) => r.genericShowArgs.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericShowArgs").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericShowArgs"),
        }
    }
    pub fn get_genericSub_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericSub_prime(r) => r.genericSub_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericSub_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericSub'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericSub_prime"),
        }
    }
    pub fn get_genericSucc_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericPred_prime_genericSucc_prime(r) => r.genericSucc_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericSucc_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericSucc'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericSucc_prime"),
        }
    }
    pub fn get_genericTT_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericTT_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericTT_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericTT'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericTT_prime"),
        }
    }
    pub fn get_genericToEnum_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericToEnum_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericToEnum_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericToEnum'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericToEnum_prime"),
        }
    }
    pub fn get_genericTop_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericTop_prime(r) => r.genericTop_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericTop_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericTop'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericTop_prime"),
        }
    }
    pub fn get_genericZero_prime(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericZero_prime.clone().unwrap(),
            Value::Record_a(r) => r.genericZero_prime.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("genericZero'").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field genericZero_prime"),
        }
    }
    pub fn get_global(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.global.clone().unwrap(),
            Value::Record_a(r) => r.global.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("global").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field global"),
        }
    }
    pub fn get_handler(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_handler_rethrow(r) => r.handler.clone().unwrap(),
            Value::Record_a(r) => r.handler.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("handler").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field handler"),
        }
    }
    pub fn get_head(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_head_tail(r) => r.head.clone().unwrap(),
            Value::Record_a(r) => r.head.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("head").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field head"),
        }
    }
    pub fn get_hour(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.hour.clone().unwrap(),
            Value::Record_a(r) => r.hour.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("hour").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field hour"),
        }
    }
    pub fn get_identity(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Semigroupoid0_identity(r) => r.identity.clone().unwrap(),
            Value::Record_a(r) => r.identity.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("identity").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field identity"),
        }
    }
    pub fn get_ignoreCase(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.ignoreCase.clone().unwrap(),
            Value::Record_a(r) => r.ignoreCase.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ignoreCase").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ignoreCase"),
        }
    }
    pub fn get_imap(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_imap(r) => r.imap.clone().unwrap(),
            Value::Record_a(r) => r.imap.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("imap").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field imap"),
        }
    }
    pub fn get_implies(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.implies.clone().unwrap(),
            Value::Record_a(r) => r.implies.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("implies").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field implies"),
        }
    }
    pub fn get_impliesRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.impliesRecord.clone().unwrap(),
            Value::Record_a(r) => r.impliesRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("impliesRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field impliesRecord"),
        }
    }
    pub fn get_index(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_index_value(r) => r.index.clone().unwrap(),
            Value::Record_a(r) => r.index.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("index").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field index"),
        }
    }
    pub fn get_init(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_acc_init(r) => r.init.clone().unwrap(),
            Value::Record_init_last(r) => r.init.clone().unwrap(),
            Value::Record_init_rest(r) => r.init.clone().unwrap(),
            Value::Record_a(r) => r.init.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("init").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field init"),
        }
    }
    pub fn get_inj(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_inj_prj(r) => r.inj.clone().unwrap(),
            Value::Record_a(r) => r.inj.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("inj").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field inj"),
        }
    }
    pub fn get_isLeft(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.isLeft.clone().unwrap(),
            Value::Record_a(r) => r.isLeft.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("isLeft").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field isLeft"),
        }
    }
    pub fn get_isSuspended(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.isSuspended.clone().unwrap(),
            Value::Record_a(r) => r.isSuspended.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("isSuspended").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field isSuspended"),
        }
    }
    pub fn get_join(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.join.clone().unwrap(),
            Value::Record_a(r) => r.join.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("join").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field join"),
        }
    }
    pub fn get_key(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_key_value(r) => r.key.clone().unwrap(),
            Value::Record_a(r) => r.key.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("key").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field key"),
        }
    }
    pub fn get_keysImpl(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_keysImpl(r) => r.keysImpl.clone().unwrap(),
            Value::Record_a(r) => r.keysImpl.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("keysImpl").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field keysImpl"),
        }
    }
    pub fn get_kill(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.kill.clone().unwrap(),
            Value::Record_a(r) => r.kill.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("kill").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field kill"),
        }
    }
    pub fn get_killed(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.killed.clone().unwrap(),
            Value::Record_a(r) => r.killed.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("killed").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field killed"),
        }
    }
    pub fn get_last(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_init_last(r) => r.last.clone().unwrap(),
            Value::Record_last_revInit(r) => r.last.clone().unwrap(),
            Value::Record_a(r) => r.last.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("last").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field last"),
        }
    }
    pub fn get_left(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_left_right(r) => r.left.clone().unwrap(),
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.left.clone().unwrap(),
            Value::Record_a(r) => r.left.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("left").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field left"),
        }
    }
    pub fn get_lift(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_lift(r) => r.lift.clone().unwrap(),
            Value::Record_a(r) => r.lift.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("lift").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field lift"),
        }
    }
    pub fn get_liftAff(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadEffect0_liftAff(r) => r.liftAff.clone().unwrap(),
            Value::Record_a(r) => r.liftAff.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("liftAff").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field liftAff"),
        }
    }
    pub fn get_liftEffect(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_liftEffect(r) => r.liftEffect.clone().unwrap(),
            Value::Record_a(r) => r.liftEffect.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("liftEffect").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field liftEffect"),
        }
    }
    pub fn get_liftST(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_liftST(r) => r.liftST.clone().unwrap(),
            Value::Record_a(r) => r.liftST.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("liftST").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field liftST"),
        }
    }
    pub fn get_listen(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.listen.clone().unwrap(),
            Value::Record_a(r) => r.listen.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("listen").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field listen"),
        }
    }
    pub fn get_local(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_ComonadAsk0_local(r) => r.local.clone().unwrap(),
            Value::Record_MonadAsk0_local(r) => r.local.clone().unwrap(),
            Value::Record_a(r) => r.local.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("local").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field local"),
        }
    }
    pub fn get_lose(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.lose.clone().unwrap(),
            Value::Record_a(r) => r.lose.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("lose").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field lose"),
        }
    }
    pub fn get_lower(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_lower(r) => r.lower.clone().unwrap(),
            Value::Record_a(r) => r.lower.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("lower").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field lower"),
        }
    }
    pub fn get_map(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_map(r) => r.map.clone().unwrap(),
            Value::Record_a(r) => r.map.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("map").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field map"),
        }
    }
    pub fn get_mapWithIndex(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Functor0_mapWithIndex(r) => r.mapWithIndex.clone().unwrap(),
            Value::Record_a(r) => r.mapWithIndex.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mapWithIndex").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mapWithIndex"),
        }
    }
    pub fn get_mappend_(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_mappend__mempty_(r) => r.mappend_.clone().unwrap(),
            Value::Record_a(r) => r.mappend_.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mappend_").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mappend_"),
        }
    }
    pub fn get_mempty(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Semigroup0_mempty(r) => r.mempty.clone().unwrap(),
            Value::Record_a(r) => r.mempty.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mempty").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mempty"),
        }
    }
    pub fn get_memptyRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_SemigroupRecord0_memptyRecord(r) => r.memptyRecord.clone().unwrap(),
            Value::Record_a(r) => r.memptyRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("memptyRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field memptyRecord"),
        }
    }
    pub fn get_mempty_(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_mappend__mempty_(r) => r.mempty_.clone().unwrap(),
            Value::Record_a(r) => r.mempty_.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mempty_").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mempty_"),
        }
    }
    pub fn get_millisecond(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.millisecond.clone().unwrap(),
            Value::Record_a(r) => r.millisecond.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("millisecond").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field millisecond"),
        }
    }
    pub fn get_minute(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.minute.clone().unwrap(),
            Value::Record_a(r) => r.minute.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("minute").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field minute"),
        }
    }
    pub fn get_mod_kw(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.mod_kw.clone().unwrap(),
            Value::Record_a(r) => r.mod_kw.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mod").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mod_kw"),
        }
    }
    pub fn get_month(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.month.clone().unwrap(),
            Value::Record_a(r) => r.month.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("month").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field month"),
        }
    }
    pub fn get_mul(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.mul.clone().unwrap(),
            Value::Record_a(r) => r.mul.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mul").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mul"),
        }
    }
    pub fn get_mulRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.mulRecord.clone().unwrap(),
            Value::Record_a(r) => r.mulRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("mulRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field mulRecord"),
        }
    }
    pub fn get_multiline(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.multiline.clone().unwrap(),
            Value::Record_a(r) => r.multiline.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("multiline").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field multiline"),
        }
    }
    pub fn get_myMethod(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_myMethod(r) => r.myMethod.clone().unwrap(),
            Value::Record_a(r) => r.myMethod.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("myMethod").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field myMethod"),
        }
    }
    pub fn get_nes(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_nes(r) => r.nes.clone().unwrap(),
            Value::Record_a(r) => r.nes.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("nes").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field nes"),
        }
    }
    pub fn get_no(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_no_yes(r) => r.no.clone().unwrap(),
            Value::Record_a(r) => r.no.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("no").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field no"),
        }
    }
    pub fn get_not(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.not.clone().unwrap(),
            Value::Record_a(r) => r.not.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("not").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field not"),
        }
    }
    pub fn get_notRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.notRecord.clone().unwrap(),
            Value::Record_a(r) => r.notRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("notRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field notRecord"),
        }
    }
    pub fn get_onComplete(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.onComplete.clone().unwrap(),
            Value::Record_a(r) => r.onComplete.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("onComplete").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field onComplete"),
        }
    }
    pub fn get_one(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.one.clone().unwrap(),
            Value::Record_a(r) => r.one.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("one").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field one"),
        }
    }
    pub fn get_oneRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.oneRecord.clone().unwrap(),
            Value::Record_a(r) => r.oneRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("oneRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field oneRecord"),
        }
    }
    pub fn get_parallel(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.parallel.clone().unwrap(),
            Value::Record_a(r) => r.parallel.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("parallel").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field parallel"),
        }
    }
    pub fn get_pass(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.pass.clone().unwrap(),
            Value::Record_a(r) => r.pass.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("pass").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field pass"),
        }
    }
    pub fn get_peek(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_peek_pos(r) => r.peek.clone().unwrap(),
            Value::Record_a(r) => r.peek.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("peek").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field peek"),
        }
    }
    pub fn get_pos(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_peek_pos(r) => r.pos.clone().unwrap(),
            Value::Record_elem_pos(r) => r.pos.clone().unwrap(),
            Value::Record_a(r) => r.pos.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("pos").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field pos"),
        }
    }
    pub fn get_pred(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ord0_pred_succ(r) => r.pred.clone().unwrap(),
            Value::Record_a(r) => r.pred.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("pred").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field pred"),
        }
    }
    pub fn get_prj(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_inj_prj(r) => r.prj.clone().unwrap(),
            Value::Record_a(r) => r.prj.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("prj").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field prj"),
        }
    }
    pub fn get_proof(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Coercible0_proof(r) => r.proof.clone().unwrap(),
            Value::Record_a(r) => r.proof.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("proof").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field proof"),
        }
    }
    pub fn get_ps_minus_rust_minus_test(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_ps_minus_rust_minus_test(r) => r.ps_minus_rust_minus_test.clone().unwrap(),
            Value::Record_a(r) => r.ps_minus_rust_minus_test.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ps-rust-test").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ps_minus_rust_minus_test"),
        }
    }
    pub fn get_pure(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_pure(r) => r.pure.clone().unwrap(),
            Value::Record_a(r) => r.pure.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("pure").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field pure"),
        }
    }
    pub fn get_recip(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ring0_recip(r) => r.recip.clone().unwrap(),
            Value::Record_a(r) => r.recip.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("recip").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field recip"),
        }
    }
    pub fn get_reflectSymbol(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_reflectSymbol(r) => r.reflectSymbol.clone().unwrap(),
            Value::Record_a(r) => r.reflectSymbol.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("reflectSymbol").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field reflectSymbol"),
        }
    }
    pub fn get_reflectType(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_reflectType(r) => r.reflectType.clone().unwrap(),
            Value::Record_a(r) => r.reflectType.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("reflectType").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field reflectType"),
        }
    }
    pub fn get_rest(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_init_rest(r) => r.rest.clone().unwrap(),
            Value::Record_a(r) => r.rest.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("rest").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field rest"),
        }
    }
    pub fn get_result(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_found_result(r) => r.result.clone().unwrap(),
            Value::Record_a(r) => r.result.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("result").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field result"),
        }
    }
    pub fn get_rethrow(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_handler_rethrow(r) => r.rethrow.clone().unwrap(),
            Value::Record_a(r) => r.rethrow.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("rethrow").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field rethrow"),
        }
    }
    pub fn get_revInit(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_last_revInit(r) => r.revInit.clone().unwrap(),
            Value::Record_a(r) => r.revInit.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("revInit").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field revInit"),
        }
    }
    pub fn get_right(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_left_right(r) => r.right.clone().unwrap(),
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.right.clone().unwrap(),
            Value::Record_a(r) => r.right.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("right").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field right"),
        }
    }
    pub fn get_run(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.run.clone().unwrap(),
            Value::Record_a(r) => r.run.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("run").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field run"),
        }
    }
    pub fn get_second(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_first_second(r) => r.second.clone().unwrap(),
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.second.clone().unwrap(),
            Value::Record_a(r) => r.second.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("second").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field second"),
        }
    }
    pub fn get_sequence(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.sequence.clone().unwrap(),
            Value::Record_a(r) => r.sequence.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("sequence").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field sequence"),
        }
    }
    pub fn get_sequence1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.sequence1.clone().unwrap(),
            Value::Record_a(r) => r.sequence1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("sequence1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field sequence1"),
        }
    }
    pub fn get_sequential(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.sequential.clone().unwrap(),
            Value::Record_a(r) => r.sequential.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("sequential").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field sequential"),
        }
    }
    pub fn get_show(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_show(r) => r.show.clone().unwrap(),
            Value::Record_a(r) => r.show.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("show").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field show"),
        }
    }
    pub fn get_showRecordFields(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_showRecordFields(r) => r.showRecordFields.clone().unwrap(),
            Value::Record_a(r) => r.showRecordFields.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("showRecordFields").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field showRecordFields"),
        }
    }
    pub fn get_state(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_state(r) => r.state.clone().unwrap(),
            Value::Record_state_val(r) => r.state.clone().unwrap(),
            Value::Record_state_value(r) => r.state.clone().unwrap(),
            Value::Record_a(r) => r.state.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("state").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field state"),
        }
    }
    pub fn get_sticky(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.sticky.clone().unwrap(),
            Value::Record_a(r) => r.sticky.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("sticky").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field sticky"),
        }
    }
    pub fn get_sub(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Semiring0_sub(r) => r.sub.clone().unwrap(),
            Value::Record_a(r) => r.sub.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("sub").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field sub"),
        }
    }
    pub fn get_subRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_SemiringRecord0_subRecord(r) => r.subRecord.clone().unwrap(),
            Value::Record_a(r) => r.subRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("subRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field subRecord"),
        }
    }
    pub fn get_succ(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ord0_pred_succ(r) => r.succ.clone().unwrap(),
            Value::Record_a(r) => r.succ.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("succ").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field succ"),
        }
    }
    pub fn get_supervisor(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fiber_supervisor(r) => r.supervisor.clone().unwrap(),
            Value::Record_a(r) => r.supervisor.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("supervisor").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field supervisor"),
        }
    }
    pub fn get_tail(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_head_tail(r) => r.tail.clone().unwrap(),
            Value::Record_a(r) => r.tail.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("tail").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field tail"),
        }
    }
    pub fn get_tailRecM(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_tailRecM(r) => r.tailRecM.clone().unwrap(),
            Value::Record_a(r) => r.tailRecM.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("tailRecM").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field tailRecM"),
        }
    }
    pub fn get_tell(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.tell.clone().unwrap(),
            Value::Record_a(r) => r.tell.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("tell").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field tell"),
        }
    }
    pub fn get_throwError(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Monad0_throwError(r) => r.throwError.clone().unwrap(),
            Value::Record_a(r) => r.throwError.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("throwError").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field throwError"),
        }
    }
    pub fn get_to(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_from_to(r) => r.to.clone().unwrap(),
            Value::Record_a(r) => r.to.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("to").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field to"),
        }
    }
    pub fn get_toDuration(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_fromDuration_toDuration(r) => r.toDuration.clone().unwrap(),
            Value::Record_a(r) => r.toDuration.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("toDuration").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field toDuration"),
        }
    }
    pub fn get_toEnum(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.toEnum.clone().unwrap(),
            Value::Record_a(r) => r.toEnum.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("toEnum").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field toEnum"),
        }
    }
    pub fn get_top(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.top.clone().unwrap(),
            Value::Record_a(r) => r.top.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("top").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field top"),
        }
    }
    pub fn get_topRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.topRecord.clone().unwrap(),
            Value::Record_a(r) => r.topRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("topRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field topRecord"),
        }
    }
    pub fn get_track(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_track(r) => r.track.clone().unwrap(),
            Value::Record_a(r) => r.track.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("track").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field track"),
        }
    }
    pub fn get_traverse(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.traverse.clone().unwrap(),
            Value::Record_a(r) => r.traverse.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("traverse").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field traverse"),
        }
    }
    pub fn get_traverse1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.traverse1.clone().unwrap(),
            Value::Record_a(r) => r.traverse1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("traverse1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field traverse1"),
        }
    }
    pub fn get_traverseWithIndex(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.traverseWithIndex.clone().unwrap(),
            Value::Record_a(r) => r.traverseWithIndex.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("traverseWithIndex").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field traverseWithIndex"),
        }
    }
    pub fn get_tt(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.tt.clone().unwrap(),
            Value::Record_a(r) => r.tt.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("tt").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field tt"),
        }
    }
    pub fn get_ttRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.ttRecord.clone().unwrap(),
            Value::Record_a(r) => r.ttRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("ttRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field ttRecord"),
        }
    }
    pub fn get_unfoldr(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_Unfoldable10_unfoldr(r) => r.unfoldr.clone().unwrap(),
            Value::Record_a(r) => r.unfoldr.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("unfoldr").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field unfoldr"),
        }
    }
    pub fn get_unfoldr1(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_unfoldr1(r) => r.unfoldr1.clone().unwrap(),
            Value::Record_a(r) => r.unfoldr1.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("unfoldr1").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field unfoldr1"),
        }
    }
    pub fn get_unicode(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.unicode.clone().unwrap(),
            Value::Record_a(r) => r.unicode.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("unicode").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field unicode"),
        }
    }
    pub fn get_val(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_acc_val(r) => r.val.clone().unwrap(),
            Value::Record_state_val(r) => r.val.clone().unwrap(),
            Value::Record_a(r) => r.val.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("val").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field val"),
        }
    }
    pub fn get_value(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_accum_value(r) => r.value.clone().unwrap(),
            Value::Record_index_value(r) => r.value.clone().unwrap(),
            Value::Record_key_value(r) => r.value.clone().unwrap(),
            Value::Record_state_value(r) => r.value.clone().unwrap(),
            Value::Record_a(r) => r.value.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("value").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field value"),
        }
    }
    pub fn get_year(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.year.clone().unwrap(),
            Value::Record_a(r) => r.year.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("year").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field year"),
        }
    }
    pub fn get_yes(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_no_yes(r) => r.yes.clone().unwrap(),
            Value::Record_a(r) => r.yes.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("yes").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field yes"),
        }
    }
    pub fn get_zero(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.zero.clone().unwrap(),
            Value::Record_a(r) => r.zero.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("zero").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field zero"),
        }
    }
    pub fn get_zeroRecord(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.zeroRecord.clone().unwrap(),
            Value::Record_a(r) => r.zeroRecord.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("zeroRecord").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field zeroRecord"),
        }
    }
    pub fn __purust_get_field(&self, name: &str) -> Option<Value> {
        match self.resolve() {
            Value::Record_Alt0_empty(r) => match name {
                "Alt0" => r.Alt0.clone(),
                "empty" => r.empty.clone(),
                _ => None,
            },
            Value::Record_Alternative1_Monad0(r) => match name {
                "Alternative1" => r.Alternative1.clone(),
                "Monad0" => r.Monad0.clone(),
                _ => None,
            },
            Value::Record_Applicative0_Bind1(r) => match name {
                "Applicative0" => r.Applicative0.clone(),
                "Bind1" => r.Bind1.clone(),
                _ => None,
            },
            Value::Record_Applicative0_Plus1(r) => match name {
                "Applicative0" => r.Applicative0.clone(),
                "Plus1" => r.Plus1.clone(),
                _ => None,
            },
            Value::Record_Apply0_Apply1_parallel_sequential(r) => match name {
                "Apply0" => r.Apply0.clone(),
                "Apply1" => r.Apply1.clone(),
                "parallel" => r.parallel.clone(),
                "sequential" => r.sequential.clone(),
                _ => None,
            },
            Value::Record_Apply0_bind(r) => match name {
                "Apply0" => r.Apply0.clone(),
                "bind" => r.bind.clone(),
                _ => None,
            },
            Value::Record_Apply0_pure(r) => match name {
                "Apply0" => r.Apply0.clone(),
                "pure" => r.pure.clone(),
                _ => None,
            },
            Value::Record_Biapply0_bipure(r) => match name {
                "Biapply0" => r.Biapply0.clone(),
                "bipure" => r.bipure.clone(),
                _ => None,
            },
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => match name {
                "Bifoldable1" => r.Bifoldable1.clone(),
                "Bifunctor0" => r.Bifunctor0.clone(),
                "bisequence" => r.bisequence.clone(),
                "bitraverse" => r.bitraverse.clone(),
                _ => None,
            },
            Value::Record_Bifunctor0_biapply(r) => match name {
                "Bifunctor0" => r.Bifunctor0.clone(),
                "biapply" => r.biapply.clone(),
                _ => None,
            },
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => match name {
                "Bounded0" => r.Bounded0.clone(),
                "Enum1" => r.Enum1.clone(),
                "cardinality" => r.cardinality.clone(),
                "fromEnum" => r.fromEnum.clone(),
                "toEnum" => r.toEnum.clone(),
                _ => None,
            },
            Value::Record_Coercible0(r) => match name {
                "Coercible0" => r.Coercible0.clone(),
                _ => None,
            },
            Value::Record_Coercible0_proof(r) => match name {
                "Coercible0" => r.Coercible0.clone(),
                "proof" => r.proof.clone(),
                _ => None,
            },
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => match name {
                "CommutativeRing0" => r.CommutativeRing0.clone(),
                "degree" => r.degree.clone(),
                "div" => r.div.clone(),
                "mod" => r.mod_kw.clone(),
                _ => None,
            },
            Value::Record_Comonad0_ask(r) => match name {
                "Comonad0" => r.Comonad0.clone(),
                "ask" => r.ask.clone(),
                _ => None,
            },
            Value::Record_Comonad0_peek_pos(r) => match name {
                "Comonad0" => r.Comonad0.clone(),
                "peek" => r.peek.clone(),
                "pos" => r.pos.clone(),
                _ => None,
            },
            Value::Record_Comonad0_track(r) => match name {
                "Comonad0" => r.Comonad0.clone(),
                "track" => r.track.clone(),
                _ => None,
            },
            Value::Record_ComonadAsk0_local(r) => match name {
                "ComonadAsk0" => r.ComonadAsk0.clone(),
                "local" => r.local.clone(),
                _ => None,
            },
            Value::Record_Contravariant0_divide(r) => match name {
                "Contravariant0" => r.Contravariant0.clone(),
                "divide" => r.divide.clone(),
                _ => None,
            },
            Value::Record_Decide0_Divisible1_lose(r) => match name {
                "Decide0" => r.Decide0.clone(),
                "Divisible1" => r.Divisible1.clone(),
                "lose" => r.lose.clone(),
                _ => None,
            },
            Value::Record_Divide0_choose(r) => match name {
                "Divide0" => r.Divide0.clone(),
                "choose" => r.choose.clone(),
                _ => None,
            },
            Value::Record_Divide0_conquer(r) => match name {
                "Divide0" => r.Divide0.clone(),
                "conquer" => r.conquer.clone(),
                _ => None,
            },
            Value::Record_DivisionRing1_EuclideanRing0(r) => match name {
                "DivisionRing1" => r.DivisionRing1.clone(),
                "EuclideanRing0" => r.EuclideanRing0.clone(),
                _ => None,
            },
            Value::Record_Eq0_compare(r) => match name {
                "Eq0" => r.Eq0.clone(),
                "compare" => r.compare.clone(),
                _ => None,
            },
            Value::Record_Eq10_compare1(r) => match name {
                "Eq10" => r.Eq10.clone(),
                "compare1" => r.compare1.clone(),
                _ => None,
            },
            Value::Record_EqRecord0_compareRecord(r) => match name {
                "EqRecord0" => r.EqRecord0.clone(),
                "compareRecord" => r.compareRecord.clone(),
                _ => None,
            },
            Value::Record_Extend0_extract(r) => match name {
                "Extend0" => r.Extend0.clone(),
                "extract" => r.extract.clone(),
                _ => None,
            },
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => match name {
                "Foldable0" => r.Foldable0.clone(),
                "foldMap1" => r.foldMap1.clone(),
                "foldl1" => r.foldl1.clone(),
                "foldr1" => r.foldr1.clone(),
                _ => None,
            },
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => match name {
                "Foldable0" => r.Foldable0.clone(),
                "foldMapWithIndex" => r.foldMapWithIndex.clone(),
                "foldlWithIndex" => r.foldlWithIndex.clone(),
                "foldrWithIndex" => r.foldrWithIndex.clone(),
                _ => None,
            },
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => match name {
                "Foldable1" => r.Foldable1.clone(),
                "Functor0" => r.Functor0.clone(),
                "sequence" => r.sequence.clone(),
                "traverse" => r.traverse.clone(),
                _ => None,
            },
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => match name {
                "Foldable10" => r.Foldable10.clone(),
                "Traversable1" => r.Traversable1.clone(),
                "sequence1" => r.sequence1.clone(),
                "traverse1" => r.traverse1.clone(),
                _ => None,
            },
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => match name {
                "FoldableWithIndex1" => r.FoldableWithIndex1.clone(),
                "FunctorWithIndex0" => r.FunctorWithIndex0.clone(),
                "Traversable2" => r.Traversable2.clone(),
                "traverseWithIndex" => r.traverseWithIndex.clone(),
                _ => None,
            },
            Value::Record_Functor0_alt(r) => match name {
                "Functor0" => r.Functor0.clone(),
                "alt" => r.alt.clone(),
                _ => None,
            },
            Value::Record_Functor0_apply(r) => match name {
                "Functor0" => r.Functor0.clone(),
                "apply" => r.apply.clone(),
                _ => None,
            },
            Value::Record_Functor0_collect_distribute(r) => match name {
                "Functor0" => r.Functor0.clone(),
                "collect" => r.collect.clone(),
                "distribute" => r.distribute.clone(),
                _ => None,
            },
            Value::Record_Functor0_extend(r) => match name {
                "Functor0" => r.Functor0.clone(),
                "extend" => r.extend.clone(),
                _ => None,
            },
            Value::Record_Functor0_mapWithIndex(r) => match name {
                "Functor0" => r.Functor0.clone(),
                "mapWithIndex" => r.mapWithIndex.clone(),
                _ => None,
            },
            Value::Record_HeytingAlgebra0(r) => match name {
                "HeytingAlgebra0" => r.HeytingAlgebra0.clone(),
                _ => None,
            },
            Value::Record_HeytingAlgebraRecord0(r) => match name {
                "HeytingAlgebraRecord0" => r.HeytingAlgebraRecord0.clone(),
                _ => None,
            },
            Value::Record_Monad0_ask(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "ask" => r.ask.clone(),
                _ => None,
            },
            Value::Record_Monad0_callCC(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "callCC" => r.callCC.clone(),
                _ => None,
            },
            Value::Record_Monad0_liftEffect(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "liftEffect" => r.liftEffect.clone(),
                _ => None,
            },
            Value::Record_Monad0_liftST(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "liftST" => r.liftST.clone(),
                _ => None,
            },
            Value::Record_Monad0_state(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "state" => r.state.clone(),
                _ => None,
            },
            Value::Record_Monad0_tailRecM(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "tailRecM" => r.tailRecM.clone(),
                _ => None,
            },
            Value::Record_Monad0_throwError(r) => match name {
                "Monad0" => r.Monad0.clone(),
                "throwError" => r.throwError.clone(),
                _ => None,
            },
            Value::Record_Monad1_Semigroup0_tell(r) => match name {
                "Monad1" => r.Monad1.clone(),
                "Semigroup0" => r.Semigroup0.clone(),
                "tell" => r.tell.clone(),
                _ => None,
            },
            Value::Record_MonadAsk0_local(r) => match name {
                "MonadAsk0" => r.MonadAsk0.clone(),
                "local" => r.local.clone(),
                _ => None,
            },
            Value::Record_MonadEffect0_liftAff(r) => match name {
                "MonadEffect0" => r.MonadEffect0.clone(),
                "liftAff" => r.liftAff.clone(),
                _ => None,
            },
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => match name {
                "MonadTell1" => r.MonadTell1.clone(),
                "Monoid0" => r.Monoid0.clone(),
                "listen" => r.listen.clone(),
                "pass" => r.pass.clone(),
                _ => None,
            },
            Value::Record_MonadThrow0_catchError(r) => match name {
                "MonadThrow0" => r.MonadThrow0.clone(),
                "catchError" => r.catchError.clone(),
                _ => None,
            },
            Value::Record_Ord0_bottom_top(r) => match name {
                "Ord0" => r.Ord0.clone(),
                "bottom" => r.bottom.clone(),
                "top" => r.top.clone(),
                _ => None,
            },
            Value::Record_Ord0_pred_succ(r) => match name {
                "Ord0" => r.Ord0.clone(),
                "pred" => r.pred.clone(),
                "succ" => r.succ.clone(),
                _ => None,
            },
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => match name {
                "OrdRecord0" => r.OrdRecord0.clone(),
                "bottomRecord" => r.bottomRecord.clone(),
                "topRecord" => r.topRecord.clone(),
                _ => None,
            },
            Value::Record_Profunctor0_closed(r) => match name {
                "Profunctor0" => r.Profunctor0.clone(),
                "closed" => r.closed.clone(),
                _ => None,
            },
            Value::Record_Profunctor0_first_second(r) => match name {
                "Profunctor0" => r.Profunctor0.clone(),
                "first" => r.first.clone(),
                "second" => r.second.clone(),
                _ => None,
            },
            Value::Record_Profunctor0_left_right(r) => match name {
                "Profunctor0" => r.Profunctor0.clone(),
                "left" => r.left.clone(),
                "right" => r.right.clone(),
                _ => None,
            },
            Value::Record_Ring0(r) => match name {
                "Ring0" => r.Ring0.clone(),
                _ => None,
            },
            Value::Record_Ring0_recip(r) => match name {
                "Ring0" => r.Ring0.clone(),
                "recip" => r.recip.clone(),
                _ => None,
            },
            Value::Record_RingRecord0(r) => match name {
                "RingRecord0" => r.RingRecord0.clone(),
                _ => None,
            },
            Value::Record_Semigroup0_mempty(r) => match name {
                "Semigroup0" => r.Semigroup0.clone(),
                "mempty" => r.mempty.clone(),
                _ => None,
            },
            Value::Record_SemigroupRecord0_memptyRecord(r) => match name {
                "SemigroupRecord0" => r.SemigroupRecord0.clone(),
                "memptyRecord" => r.memptyRecord.clone(),
                _ => None,
            },
            Value::Record_Semigroupoid0_identity(r) => match name {
                "Semigroupoid0" => r.Semigroupoid0.clone(),
                "identity" => r.identity.clone(),
                _ => None,
            },
            Value::Record_Semiring0_sub(r) => match name {
                "Semiring0" => r.Semiring0.clone(),
                "sub" => r.sub.clone(),
                _ => None,
            },
            Value::Record_SemiringRecord0_subRecord(r) => match name {
                "SemiringRecord0" => r.SemiringRecord0.clone(),
                "subRecord" => r.subRecord.clone(),
                _ => None,
            },
            Value::Record_Unfoldable10_unfoldr(r) => match name {
                "Unfoldable10" => r.Unfoldable10.clone(),
                "unfoldr" => r.unfoldr.clone(),
                _ => None,
            },
            Value::Record_a_b(r) => match name {
                "a" => r.a.clone(),
                "b" => r.b.clone(),
                _ => None,
            },
            Value::Record_a_b_c(r) => match name {
                "a" => r.a.clone(),
                "b" => r.b.clone(),
                "c" => r.c.clone(),
                _ => None,
            },
            Value::Record_a_b_c_d_e(r) => match name {
                "a" => r.a.clone(),
                "b" => r.b.clone(),
                "c" => r.c.clone(),
                "d" => r.d.clone(),
                "e" => r.e.clone(),
                _ => None,
            },
            Value::Record_acc_init(r) => match name {
                "acc" => r.acc.clone(),
                "init" => r.init.clone(),
                _ => None,
            },
            Value::Record_acc_val(r) => match name {
                "acc" => r.acc.clone(),
                "val" => r.val.clone(),
                _ => None,
            },
            Value::Record_accum_value(r) => match name {
                "accum" => r.accum.clone(),
                "value" => r.value.clone(),
                _ => None,
            },
            Value::Record_add_mul_one_zero(r) => match name {
                "add" => r.add.clone(),
                "mul" => r.mul.clone(),
                "one" => r.one.clone(),
                "zero" => r.zero.clone(),
                _ => None,
            },
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => match name {
                "addRecord" => r.addRecord.clone(),
                "mulRecord" => r.mulRecord.clone(),
                "oneRecord" => r.oneRecord.clone(),
                "zeroRecord" => r.zeroRecord.clone(),
                _ => None,
            },
            Value::Record_after_before(r) => match name {
                "after" => r.after.clone(),
                "before" => r.before.clone(),
                _ => None,
            },
            Value::Record_append(r) => match name {
                "append" => r.append.clone(),
                _ => None,
            },
            Value::Record_appendRecord(r) => match name {
                "appendRecord" => r.appendRecord.clone(),
                _ => None,
            },
            Value::Record_asList_asMap(r) => match name {
                "asList" => r.asList.clone(),
                "asMap" => r.asMap.clone(),
                _ => None,
            },
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => match name {
                "bifoldMap" => r.bifoldMap.clone(),
                "bifoldl" => r.bifoldl.clone(),
                "bifoldr" => r.bifoldr.clone(),
                _ => None,
            },
            Value::Record_bimap(r) => match name {
                "bimap" => r.bimap.clone(),
                _ => None,
            },
            Value::Record_c_d(r) => match name {
                "c" => r.c.clone(),
                "d" => r.d.clone(),
                _ => None,
            },
            Value::Record_cmap(r) => match name {
                "cmap" => r.cmap.clone(),
                _ => None,
            },
            Value::Record_completed_failed_killed(r) => match name {
                "completed" => r.completed.clone(),
                "failed" => r.failed.clone(),
                "killed" => r.killed.clone(),
                _ => None,
            },
            Value::Record_compose(r) => match name {
                "compose" => r.compose.clone(),
                _ => None,
            },
            Value::Record_conj_disj_ff_implies_not_tt(r) => match name {
                "conj" => r.conj.clone(),
                "disj" => r.disj.clone(),
                "ff" => r.ff.clone(),
                "implies" => r.implies.clone(),
                "not" => r.not.clone(),
                "tt" => r.tt.clone(),
                _ => None,
            },
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => match name {
                "conjRecord" => r.conjRecord.clone(),
                "disjRecord" => r.disjRecord.clone(),
                "ffRecord" => r.ffRecord.clone(),
                "impliesRecord" => r.impliesRecord.clone(),
                "notRecord" => r.notRecord.clone(),
                "ttRecord" => r.ttRecord.clone(),
                _ => None,
            },
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => match name {
                "day" => r.day.clone(),
                "hour" => r.hour.clone(),
                "millisecond" => r.millisecond.clone(),
                "minute" => r.minute.clone(),
                "month" => r.month.clone(),
                "second" => r.second.clone(),
                "year" => r.year.clone(),
                _ => None,
            },
            Value::Record_defer(r) => match name {
                "defer" => r.defer.clone(),
                _ => None,
            },
            Value::Record_dimap(r) => match name {
                "dimap" => r.dimap.clone(),
                _ => None,
            },
            Value::Record_discard(r) => match name {
                "discard" => r.discard.clone(),
                _ => None,
            },
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => match name {
                "dotAll" => r.dotAll.clone(),
                "global" => r.global.clone(),
                "ignoreCase" => r.ignoreCase.clone(),
                "multiline" => r.multiline.clone(),
                "sticky" => r.sticky.clone(),
                "unicode" => r.unicode.clone(),
                _ => None,
            },
            Value::Record_e_f(r) => match name {
                "e" => r.e.clone(),
                "f" => r.f.clone(),
                _ => None,
            },
            Value::Record_elem_pos(r) => match name {
                "elem" => r.elem.clone(),
                "pos" => r.pos.clone(),
                _ => None,
            },
            Value::Record_eq(r) => match name {
                "eq" => r.eq.clone(),
                _ => None,
            },
            Value::Record_eq1(r) => match name {
                "eq1" => r.eq1.clone(),
                _ => None,
            },
            Value::Record_eqRecord(r) => match name {
                "eqRecord" => r.eqRecord.clone(),
                _ => None,
            },
            Value::Record_fiber_supervisor(r) => match name {
                "fiber" => r.fiber.clone(),
                "supervisor" => r.supervisor.clone(),
                _ => None,
            },
            Value::Record_foldMap_foldl_foldr(r) => match name {
                "foldMap" => r.foldMap.clone(),
                "foldl" => r.foldl.clone(),
                "foldr" => r.foldr.clone(),
                _ => None,
            },
            Value::Record_found_result(r) => match name {
                "found" => r.found.clone(),
                "result" => r.result.clone(),
                _ => None,
            },
            Value::Record_from_to(r) => match name {
                "from" => r.from.clone(),
                "to" => r.to.clone(),
                _ => None,
            },
            Value::Record_fromDuration_toDuration(r) => match name {
                "fromDuration" => r.fromDuration.clone(),
                "toDuration" => r.toDuration.clone(),
                _ => None,
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => match name {
                "fromLeft" => r.fromLeft.clone(),
                "fromRight" => r.fromRight.clone(),
                "isLeft" => r.isLeft.clone(),
                "left" => r.left.clone(),
                "right" => r.right.clone(),
                _ => None,
            },
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => match name {
                "genericAdd'" => r.genericAdd_prime.clone(),
                "genericMul'" => r.genericMul_prime.clone(),
                "genericOne'" => r.genericOne_prime.clone(),
                "genericZero'" => r.genericZero_prime.clone(),
                _ => None,
            },
            Value::Record_genericAppend_prime(r) => match name {
                "genericAppend'" => r.genericAppend_prime.clone(),
                _ => None,
            },
            Value::Record_genericBottom_prime(r) => match name {
                "genericBottom'" => r.genericBottom_prime.clone(),
                _ => None,
            },
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => match name {
                "genericCardinality'" => r.genericCardinality_prime.clone(),
                "genericFromEnum'" => r.genericFromEnum_prime.clone(),
                "genericToEnum'" => r.genericToEnum_prime.clone(),
                _ => None,
            },
            Value::Record_genericCompare_prime(r) => match name {
                "genericCompare'" => r.genericCompare_prime.clone(),
                _ => None,
            },
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => match name {
                "genericConj'" => r.genericConj_prime.clone(),
                "genericDisj'" => r.genericDisj_prime.clone(),
                "genericFF'" => r.genericFF_prime.clone(),
                "genericImplies'" => r.genericImplies_prime.clone(),
                "genericNot'" => r.genericNot_prime.clone(),
                "genericTT'" => r.genericTT_prime.clone(),
                _ => None,
            },
            Value::Record_genericEq_prime(r) => match name {
                "genericEq'" => r.genericEq_prime.clone(),
                _ => None,
            },
            Value::Record_genericMempty_prime(r) => match name {
                "genericMempty'" => r.genericMempty_prime.clone(),
                _ => None,
            },
            Value::Record_genericPred_prime_genericSucc_prime(r) => match name {
                "genericPred'" => r.genericPred_prime.clone(),
                "genericSucc'" => r.genericSucc_prime.clone(),
                _ => None,
            },
            Value::Record_genericShow_prime(r) => match name {
                "genericShow'" => r.genericShow_prime.clone(),
                _ => None,
            },
            Value::Record_genericShowArgs(r) => match name {
                "genericShowArgs" => r.genericShowArgs.clone(),
                _ => None,
            },
            Value::Record_genericSub_prime(r) => match name {
                "genericSub'" => r.genericSub_prime.clone(),
                _ => None,
            },
            Value::Record_genericTop_prime(r) => match name {
                "genericTop'" => r.genericTop_prime.clone(),
                _ => None,
            },
            Value::Record_handler_rethrow(r) => match name {
                "handler" => r.handler.clone(),
                "rethrow" => r.rethrow.clone(),
                _ => None,
            },
            Value::Record_head_tail(r) => match name {
                "head" => r.head.clone(),
                "tail" => r.tail.clone(),
                _ => None,
            },
            Value::Record_imap(r) => match name {
                "imap" => r.imap.clone(),
                _ => None,
            },
            Value::Record_index_value(r) => match name {
                "index" => r.index.clone(),
                "value" => r.value.clone(),
                _ => None,
            },
            Value::Record_init_last(r) => match name {
                "init" => r.init.clone(),
                "last" => r.last.clone(),
                _ => None,
            },
            Value::Record_init_rest(r) => match name {
                "init" => r.init.clone(),
                "rest" => r.rest.clone(),
                _ => None,
            },
            Value::Record_inj_prj(r) => match name {
                "inj" => r.inj.clone(),
                "prj" => r.prj.clone(),
                _ => None,
            },
            Value::Record_isSuspended_join_kill_onComplete_run(r) => match name {
                "isSuspended" => r.isSuspended.clone(),
                "join" => r.join.clone(),
                "kill" => r.kill.clone(),
                "onComplete" => r.onComplete.clone(),
                "run" => r.run.clone(),
                _ => None,
            },
            Value::Record_key_value(r) => match name {
                "key" => r.key.clone(),
                "value" => r.value.clone(),
                _ => None,
            },
            Value::Record_keysImpl(r) => match name {
                "keysImpl" => r.keysImpl.clone(),
                _ => None,
            },
            Value::Record_last_revInit(r) => match name {
                "last" => r.last.clone(),
                "revInit" => r.revInit.clone(),
                _ => None,
            },
            Value::Record_lift(r) => match name {
                "lift" => r.lift.clone(),
                _ => None,
            },
            Value::Record_lower(r) => match name {
                "lower" => r.lower.clone(),
                _ => None,
            },
            Value::Record_map(r) => match name {
                "map" => r.map.clone(),
                _ => None,
            },
            Value::Record_mappend__mempty_(r) => match name {
                "mappend_" => r.mappend_.clone(),
                "mempty_" => r.mempty_.clone(),
                _ => None,
            },
            Value::Record_myMethod(r) => match name {
                "myMethod" => r.myMethod.clone(),
                _ => None,
            },
            Value::Record_nes(r) => match name {
                "nes" => r.nes.clone(),
                _ => None,
            },
            Value::Record_no_yes(r) => match name {
                "no" => r.no.clone(),
                "yes" => r.yes.clone(),
                _ => None,
            },
            Value::Record_ps_minus_rust_minus_test(r) => match name {
                "ps-rust-test" => r.ps_minus_rust_minus_test.clone(),
                _ => None,
            },
            Value::Record_reflectSymbol(r) => match name {
                "reflectSymbol" => r.reflectSymbol.clone(),
                _ => None,
            },
            Value::Record_reflectType(r) => match name {
                "reflectType" => r.reflectType.clone(),
                _ => None,
            },
            Value::Record_show(r) => match name {
                "show" => r.show.clone(),
                _ => None,
            },
            Value::Record_showRecordFields(r) => match name {
                "showRecordFields" => r.showRecordFields.clone(),
                _ => None,
            },
            Value::Record_state_val(r) => match name {
                "state" => r.state.clone(),
                "val" => r.val.clone(),
                _ => None,
            },
            Value::Record_state_value(r) => match name {
                "state" => r.state.clone(),
                "value" => r.value.clone(),
                _ => None,
            },
            Value::Record_unfoldr1(r) => match name {
                "unfoldr1" => r.unfoldr1.clone(),
                _ => None,
            },
            Value::Record_a(r) => match name {
                "Alt0" => r.Alt0.clone(),
                "Alternative1" => r.Alternative1.clone(),
                "Applicative0" => r.Applicative0.clone(),
                "Apply0" => r.Apply0.clone(),
                "Apply1" => r.Apply1.clone(),
                "Biapply0" => r.Biapply0.clone(),
                "Bifoldable1" => r.Bifoldable1.clone(),
                "Bifunctor0" => r.Bifunctor0.clone(),
                "Bind1" => r.Bind1.clone(),
                "Bounded0" => r.Bounded0.clone(),
                "Coercible0" => r.Coercible0.clone(),
                "CommutativeRing0" => r.CommutativeRing0.clone(),
                "Comonad0" => r.Comonad0.clone(),
                "ComonadAsk0" => r.ComonadAsk0.clone(),
                "Contravariant0" => r.Contravariant0.clone(),
                "Decide0" => r.Decide0.clone(),
                "Divide0" => r.Divide0.clone(),
                "Divisible1" => r.Divisible1.clone(),
                "DivisionRing1" => r.DivisionRing1.clone(),
                "Enum1" => r.Enum1.clone(),
                "Eq0" => r.Eq0.clone(),
                "Eq10" => r.Eq10.clone(),
                "EqRecord0" => r.EqRecord0.clone(),
                "EuclideanRing0" => r.EuclideanRing0.clone(),
                "Extend0" => r.Extend0.clone(),
                "Foldable0" => r.Foldable0.clone(),
                "Foldable1" => r.Foldable1.clone(),
                "Foldable10" => r.Foldable10.clone(),
                "FoldableWithIndex1" => r.FoldableWithIndex1.clone(),
                "Functor0" => r.Functor0.clone(),
                "FunctorWithIndex0" => r.FunctorWithIndex0.clone(),
                "HeytingAlgebra0" => r.HeytingAlgebra0.clone(),
                "HeytingAlgebraRecord0" => r.HeytingAlgebraRecord0.clone(),
                "Monad0" => r.Monad0.clone(),
                "Monad1" => r.Monad1.clone(),
                "MonadAsk0" => r.MonadAsk0.clone(),
                "MonadEffect0" => r.MonadEffect0.clone(),
                "MonadTell1" => r.MonadTell1.clone(),
                "MonadThrow0" => r.MonadThrow0.clone(),
                "Monoid0" => r.Monoid0.clone(),
                "Ord0" => r.Ord0.clone(),
                "OrdRecord0" => r.OrdRecord0.clone(),
                "Plus1" => r.Plus1.clone(),
                "Profunctor0" => r.Profunctor0.clone(),
                "Ring0" => r.Ring0.clone(),
                "RingRecord0" => r.RingRecord0.clone(),
                "Semigroup0" => r.Semigroup0.clone(),
                "SemigroupRecord0" => r.SemigroupRecord0.clone(),
                "Semigroupoid0" => r.Semigroupoid0.clone(),
                "Semiring0" => r.Semiring0.clone(),
                "SemiringRecord0" => r.SemiringRecord0.clone(),
                "Traversable1" => r.Traversable1.clone(),
                "Traversable2" => r.Traversable2.clone(),
                "Unfoldable10" => r.Unfoldable10.clone(),
                "a" => r.a.clone(),
                "acc" => r.acc.clone(),
                "accum" => r.accum.clone(),
                "add" => r.add.clone(),
                "addRecord" => r.addRecord.clone(),
                "after" => r.after.clone(),
                "alt" => r.alt.clone(),
                "append" => r.append.clone(),
                "appendRecord" => r.appendRecord.clone(),
                "apply" => r.apply.clone(),
                "asList" => r.asList.clone(),
                "asMap" => r.asMap.clone(),
                "ask" => r.ask.clone(),
                "b" => r.b.clone(),
                "before" => r.before.clone(),
                "biapply" => r.biapply.clone(),
                "bifoldMap" => r.bifoldMap.clone(),
                "bifoldl" => r.bifoldl.clone(),
                "bifoldr" => r.bifoldr.clone(),
                "bimap" => r.bimap.clone(),
                "bind" => r.bind.clone(),
                "bipure" => r.bipure.clone(),
                "bisequence" => r.bisequence.clone(),
                "bitraverse" => r.bitraverse.clone(),
                "bottom" => r.bottom.clone(),
                "bottomRecord" => r.bottomRecord.clone(),
                "c" => r.c.clone(),
                "callCC" => r.callCC.clone(),
                "cardinality" => r.cardinality.clone(),
                "catchError" => r.catchError.clone(),
                "choose" => r.choose.clone(),
                "closed" => r.closed.clone(),
                "cmap" => r.cmap.clone(),
                "collect" => r.collect.clone(),
                "compare" => r.compare.clone(),
                "compare1" => r.compare1.clone(),
                "compareRecord" => r.compareRecord.clone(),
                "completed" => r.completed.clone(),
                "compose" => r.compose.clone(),
                "conj" => r.conj.clone(),
                "conjRecord" => r.conjRecord.clone(),
                "conquer" => r.conquer.clone(),
                "d" => r.d.clone(),
                "day" => r.day.clone(),
                "defer" => r.defer.clone(),
                "degree" => r.degree.clone(),
                "dimap" => r.dimap.clone(),
                "discard" => r.discard.clone(),
                "disj" => r.disj.clone(),
                "disjRecord" => r.disjRecord.clone(),
                "distribute" => r.distribute.clone(),
                "div" => r.div.clone(),
                "divide" => r.divide.clone(),
                "dotAll" => r.dotAll.clone(),
                "e" => r.e.clone(),
                "elem" => r.elem.clone(),
                "empty" => r.empty.clone(),
                "eq" => r.eq.clone(),
                "eq1" => r.eq1.clone(),
                "eqRecord" => r.eqRecord.clone(),
                "extend" => r.extend.clone(),
                "extract" => r.extract.clone(),
                "f" => r.f.clone(),
                "failed" => r.failed.clone(),
                "ff" => r.ff.clone(),
                "ffRecord" => r.ffRecord.clone(),
                "fiber" => r.fiber.clone(),
                "first" => r.first.clone(),
                "foldMap" => r.foldMap.clone(),
                "foldMap1" => r.foldMap1.clone(),
                "foldMapWithIndex" => r.foldMapWithIndex.clone(),
                "foldl" => r.foldl.clone(),
                "foldl1" => r.foldl1.clone(),
                "foldlWithIndex" => r.foldlWithIndex.clone(),
                "foldr" => r.foldr.clone(),
                "foldr1" => r.foldr1.clone(),
                "foldrWithIndex" => r.foldrWithIndex.clone(),
                "found" => r.found.clone(),
                "from" => r.from.clone(),
                "fromDuration" => r.fromDuration.clone(),
                "fromEnum" => r.fromEnum.clone(),
                "fromLeft" => r.fromLeft.clone(),
                "fromRight" => r.fromRight.clone(),
                "genericAdd'" => r.genericAdd_prime.clone(),
                "genericAppend'" => r.genericAppend_prime.clone(),
                "genericBottom'" => r.genericBottom_prime.clone(),
                "genericCardinality'" => r.genericCardinality_prime.clone(),
                "genericCompare'" => r.genericCompare_prime.clone(),
                "genericConj'" => r.genericConj_prime.clone(),
                "genericDisj'" => r.genericDisj_prime.clone(),
                "genericEq'" => r.genericEq_prime.clone(),
                "genericFF'" => r.genericFF_prime.clone(),
                "genericFromEnum'" => r.genericFromEnum_prime.clone(),
                "genericImplies'" => r.genericImplies_prime.clone(),
                "genericMempty'" => r.genericMempty_prime.clone(),
                "genericMul'" => r.genericMul_prime.clone(),
                "genericNot'" => r.genericNot_prime.clone(),
                "genericOne'" => r.genericOne_prime.clone(),
                "genericPred'" => r.genericPred_prime.clone(),
                "genericShow'" => r.genericShow_prime.clone(),
                "genericShowArgs" => r.genericShowArgs.clone(),
                "genericSub'" => r.genericSub_prime.clone(),
                "genericSucc'" => r.genericSucc_prime.clone(),
                "genericTT'" => r.genericTT_prime.clone(),
                "genericToEnum'" => r.genericToEnum_prime.clone(),
                "genericTop'" => r.genericTop_prime.clone(),
                "genericZero'" => r.genericZero_prime.clone(),
                "global" => r.global.clone(),
                "handler" => r.handler.clone(),
                "head" => r.head.clone(),
                "hour" => r.hour.clone(),
                "identity" => r.identity.clone(),
                "ignoreCase" => r.ignoreCase.clone(),
                "imap" => r.imap.clone(),
                "implies" => r.implies.clone(),
                "impliesRecord" => r.impliesRecord.clone(),
                "index" => r.index.clone(),
                "init" => r.init.clone(),
                "inj" => r.inj.clone(),
                "isLeft" => r.isLeft.clone(),
                "isSuspended" => r.isSuspended.clone(),
                "join" => r.join.clone(),
                "key" => r.key.clone(),
                "keysImpl" => r.keysImpl.clone(),
                "kill" => r.kill.clone(),
                "killed" => r.killed.clone(),
                "last" => r.last.clone(),
                "left" => r.left.clone(),
                "lift" => r.lift.clone(),
                "liftAff" => r.liftAff.clone(),
                "liftEffect" => r.liftEffect.clone(),
                "liftST" => r.liftST.clone(),
                "listen" => r.listen.clone(),
                "local" => r.local.clone(),
                "lose" => r.lose.clone(),
                "lower" => r.lower.clone(),
                "map" => r.map.clone(),
                "mapWithIndex" => r.mapWithIndex.clone(),
                "mappend_" => r.mappend_.clone(),
                "mempty" => r.mempty.clone(),
                "memptyRecord" => r.memptyRecord.clone(),
                "mempty_" => r.mempty_.clone(),
                "millisecond" => r.millisecond.clone(),
                "minute" => r.minute.clone(),
                "mod" => r.mod_kw.clone(),
                "month" => r.month.clone(),
                "mul" => r.mul.clone(),
                "mulRecord" => r.mulRecord.clone(),
                "multiline" => r.multiline.clone(),
                "myMethod" => r.myMethod.clone(),
                "nes" => r.nes.clone(),
                "no" => r.no.clone(),
                "not" => r.not.clone(),
                "notRecord" => r.notRecord.clone(),
                "onComplete" => r.onComplete.clone(),
                "one" => r.one.clone(),
                "oneRecord" => r.oneRecord.clone(),
                "parallel" => r.parallel.clone(),
                "pass" => r.pass.clone(),
                "peek" => r.peek.clone(),
                "pos" => r.pos.clone(),
                "pred" => r.pred.clone(),
                "prj" => r.prj.clone(),
                "proof" => r.proof.clone(),
                "ps-rust-test" => r.ps_minus_rust_minus_test.clone(),
                "pure" => r.pure.clone(),
                "recip" => r.recip.clone(),
                "reflectSymbol" => r.reflectSymbol.clone(),
                "reflectType" => r.reflectType.clone(),
                "rest" => r.rest.clone(),
                "result" => r.result.clone(),
                "rethrow" => r.rethrow.clone(),
                "revInit" => r.revInit.clone(),
                "right" => r.right.clone(),
                "run" => r.run.clone(),
                "second" => r.second.clone(),
                "sequence" => r.sequence.clone(),
                "sequence1" => r.sequence1.clone(),
                "sequential" => r.sequential.clone(),
                "show" => r.show.clone(),
                "showRecordFields" => r.showRecordFields.clone(),
                "state" => r.state.clone(),
                "sticky" => r.sticky.clone(),
                "sub" => r.sub.clone(),
                "subRecord" => r.subRecord.clone(),
                "succ" => r.succ.clone(),
                "supervisor" => r.supervisor.clone(),
                "tail" => r.tail.clone(),
                "tailRecM" => r.tailRecM.clone(),
                "tell" => r.tell.clone(),
                "throwError" => r.throwError.clone(),
                "to" => r.to.clone(),
                "toDuration" => r.toDuration.clone(),
                "toEnum" => r.toEnum.clone(),
                "top" => r.top.clone(),
                "topRecord" => r.topRecord.clone(),
                "track" => r.track.clone(),
                "traverse" => r.traverse.clone(),
                "traverse1" => r.traverse1.clone(),
                "traverseWithIndex" => r.traverseWithIndex.clone(),
                "tt" => r.tt.clone(),
                "ttRecord" => r.ttRecord.clone(),
                "unfoldr" => r.unfoldr.clone(),
                "unfoldr1" => r.unfoldr1.clone(),
                "unicode" => r.unicode.clone(),
                "val" => r.val.clone(),
                "value" => r.value.clone(),
                "year" => r.year.clone(),
                "yes" => r.yes.clone(),
                "zero" => r.zero.clone(),
                "zeroRecord" => r.zeroRecord.clone(),
                _ => None,
            },
            Value::DynamicRecord(r) => r.get(name).cloned(),
            _ => panic!("Expected record"),
        }
    }
    pub fn __purust_set_field(mut self, name: &str, value: Value) -> Value {
        if matches!(self, Value::Thunk(_)) { self = self.resolve().clone(); }
        match &mut self {
            Value::Record_Alt0_empty(r) => match name {
                "Alt0" => { perceus_ptr::PerceusPtr::make_mut(r).Alt0 = Some(value); return self; },
                "empty" => { perceus_ptr::PerceusPtr::make_mut(r).empty = Some(value); return self; },
                _ => {},
            },
            Value::Record_Alternative1_Monad0(r) => match name {
                "Alternative1" => { perceus_ptr::PerceusPtr::make_mut(r).Alternative1 = Some(value); return self; },
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Applicative0_Bind1(r) => match name {
                "Applicative0" => { perceus_ptr::PerceusPtr::make_mut(r).Applicative0 = Some(value); return self; },
                "Bind1" => { perceus_ptr::PerceusPtr::make_mut(r).Bind1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Applicative0_Plus1(r) => match name {
                "Applicative0" => { perceus_ptr::PerceusPtr::make_mut(r).Applicative0 = Some(value); return self; },
                "Plus1" => { perceus_ptr::PerceusPtr::make_mut(r).Plus1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Apply0_Apply1_parallel_sequential(r) => match name {
                "Apply0" => { perceus_ptr::PerceusPtr::make_mut(r).Apply0 = Some(value); return self; },
                "Apply1" => { perceus_ptr::PerceusPtr::make_mut(r).Apply1 = Some(value); return self; },
                "parallel" => { perceus_ptr::PerceusPtr::make_mut(r).parallel = Some(value); return self; },
                "sequential" => { perceus_ptr::PerceusPtr::make_mut(r).sequential = Some(value); return self; },
                _ => {},
            },
            Value::Record_Apply0_bind(r) => match name {
                "Apply0" => { perceus_ptr::PerceusPtr::make_mut(r).Apply0 = Some(value); return self; },
                "bind" => { perceus_ptr::PerceusPtr::make_mut(r).bind = Some(value); return self; },
                _ => {},
            },
            Value::Record_Apply0_pure(r) => match name {
                "Apply0" => { perceus_ptr::PerceusPtr::make_mut(r).Apply0 = Some(value); return self; },
                "pure" => { perceus_ptr::PerceusPtr::make_mut(r).pure = Some(value); return self; },
                _ => {},
            },
            Value::Record_Biapply0_bipure(r) => match name {
                "Biapply0" => { perceus_ptr::PerceusPtr::make_mut(r).Biapply0 = Some(value); return self; },
                "bipure" => { perceus_ptr::PerceusPtr::make_mut(r).bipure = Some(value); return self; },
                _ => {},
            },
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => match name {
                "Bifoldable1" => { perceus_ptr::PerceusPtr::make_mut(r).Bifoldable1 = Some(value); return self; },
                "Bifunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Bifunctor0 = Some(value); return self; },
                "bisequence" => { perceus_ptr::PerceusPtr::make_mut(r).bisequence = Some(value); return self; },
                "bitraverse" => { perceus_ptr::PerceusPtr::make_mut(r).bitraverse = Some(value); return self; },
                _ => {},
            },
            Value::Record_Bifunctor0_biapply(r) => match name {
                "Bifunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Bifunctor0 = Some(value); return self; },
                "biapply" => { perceus_ptr::PerceusPtr::make_mut(r).biapply = Some(value); return self; },
                _ => {},
            },
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => match name {
                "Bounded0" => { perceus_ptr::PerceusPtr::make_mut(r).Bounded0 = Some(value); return self; },
                "Enum1" => { perceus_ptr::PerceusPtr::make_mut(r).Enum1 = Some(value); return self; },
                "cardinality" => { perceus_ptr::PerceusPtr::make_mut(r).cardinality = Some(value); return self; },
                "fromEnum" => { perceus_ptr::PerceusPtr::make_mut(r).fromEnum = Some(value); return self; },
                "toEnum" => { perceus_ptr::PerceusPtr::make_mut(r).toEnum = Some(value); return self; },
                _ => {},
            },
            Value::Record_Coercible0(r) => match name {
                "Coercible0" => { perceus_ptr::PerceusPtr::make_mut(r).Coercible0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Coercible0_proof(r) => match name {
                "Coercible0" => { perceus_ptr::PerceusPtr::make_mut(r).Coercible0 = Some(value); return self; },
                "proof" => { perceus_ptr::PerceusPtr::make_mut(r).proof = Some(value); return self; },
                _ => {},
            },
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => match name {
                "CommutativeRing0" => { perceus_ptr::PerceusPtr::make_mut(r).CommutativeRing0 = Some(value); return self; },
                "degree" => { perceus_ptr::PerceusPtr::make_mut(r).degree = Some(value); return self; },
                "div" => { perceus_ptr::PerceusPtr::make_mut(r).div = Some(value); return self; },
                "mod" => { perceus_ptr::PerceusPtr::make_mut(r).mod_kw = Some(value); return self; },
                _ => {},
            },
            Value::Record_Comonad0_ask(r) => match name {
                "Comonad0" => { perceus_ptr::PerceusPtr::make_mut(r).Comonad0 = Some(value); return self; },
                "ask" => { perceus_ptr::PerceusPtr::make_mut(r).ask = Some(value); return self; },
                _ => {},
            },
            Value::Record_Comonad0_peek_pos(r) => match name {
                "Comonad0" => { perceus_ptr::PerceusPtr::make_mut(r).Comonad0 = Some(value); return self; },
                "peek" => { perceus_ptr::PerceusPtr::make_mut(r).peek = Some(value); return self; },
                "pos" => { perceus_ptr::PerceusPtr::make_mut(r).pos = Some(value); return self; },
                _ => {},
            },
            Value::Record_Comonad0_track(r) => match name {
                "Comonad0" => { perceus_ptr::PerceusPtr::make_mut(r).Comonad0 = Some(value); return self; },
                "track" => { perceus_ptr::PerceusPtr::make_mut(r).track = Some(value); return self; },
                _ => {},
            },
            Value::Record_ComonadAsk0_local(r) => match name {
                "ComonadAsk0" => { perceus_ptr::PerceusPtr::make_mut(r).ComonadAsk0 = Some(value); return self; },
                "local" => { perceus_ptr::PerceusPtr::make_mut(r).local = Some(value); return self; },
                _ => {},
            },
            Value::Record_Contravariant0_divide(r) => match name {
                "Contravariant0" => { perceus_ptr::PerceusPtr::make_mut(r).Contravariant0 = Some(value); return self; },
                "divide" => { perceus_ptr::PerceusPtr::make_mut(r).divide = Some(value); return self; },
                _ => {},
            },
            Value::Record_Decide0_Divisible1_lose(r) => match name {
                "Decide0" => { perceus_ptr::PerceusPtr::make_mut(r).Decide0 = Some(value); return self; },
                "Divisible1" => { perceus_ptr::PerceusPtr::make_mut(r).Divisible1 = Some(value); return self; },
                "lose" => { perceus_ptr::PerceusPtr::make_mut(r).lose = Some(value); return self; },
                _ => {},
            },
            Value::Record_Divide0_choose(r) => match name {
                "Divide0" => { perceus_ptr::PerceusPtr::make_mut(r).Divide0 = Some(value); return self; },
                "choose" => { perceus_ptr::PerceusPtr::make_mut(r).choose = Some(value); return self; },
                _ => {},
            },
            Value::Record_Divide0_conquer(r) => match name {
                "Divide0" => { perceus_ptr::PerceusPtr::make_mut(r).Divide0 = Some(value); return self; },
                "conquer" => { perceus_ptr::PerceusPtr::make_mut(r).conquer = Some(value); return self; },
                _ => {},
            },
            Value::Record_DivisionRing1_EuclideanRing0(r) => match name {
                "DivisionRing1" => { perceus_ptr::PerceusPtr::make_mut(r).DivisionRing1 = Some(value); return self; },
                "EuclideanRing0" => { perceus_ptr::PerceusPtr::make_mut(r).EuclideanRing0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Eq0_compare(r) => match name {
                "Eq0" => { perceus_ptr::PerceusPtr::make_mut(r).Eq0 = Some(value); return self; },
                "compare" => { perceus_ptr::PerceusPtr::make_mut(r).compare = Some(value); return self; },
                _ => {},
            },
            Value::Record_Eq10_compare1(r) => match name {
                "Eq10" => { perceus_ptr::PerceusPtr::make_mut(r).Eq10 = Some(value); return self; },
                "compare1" => { perceus_ptr::PerceusPtr::make_mut(r).compare1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_EqRecord0_compareRecord(r) => match name {
                "EqRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).EqRecord0 = Some(value); return self; },
                "compareRecord" => { perceus_ptr::PerceusPtr::make_mut(r).compareRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_Extend0_extract(r) => match name {
                "Extend0" => { perceus_ptr::PerceusPtr::make_mut(r).Extend0 = Some(value); return self; },
                "extract" => { perceus_ptr::PerceusPtr::make_mut(r).extract = Some(value); return self; },
                _ => {},
            },
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => match name {
                "Foldable0" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable0 = Some(value); return self; },
                "foldMap1" => { perceus_ptr::PerceusPtr::make_mut(r).foldMap1 = Some(value); return self; },
                "foldl1" => { perceus_ptr::PerceusPtr::make_mut(r).foldl1 = Some(value); return self; },
                "foldr1" => { perceus_ptr::PerceusPtr::make_mut(r).foldr1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => match name {
                "Foldable0" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable0 = Some(value); return self; },
                "foldMapWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldMapWithIndex = Some(value); return self; },
                "foldlWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldlWithIndex = Some(value); return self; },
                "foldrWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldrWithIndex = Some(value); return self; },
                _ => {},
            },
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => match name {
                "Foldable1" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable1 = Some(value); return self; },
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "sequence" => { perceus_ptr::PerceusPtr::make_mut(r).sequence = Some(value); return self; },
                "traverse" => { perceus_ptr::PerceusPtr::make_mut(r).traverse = Some(value); return self; },
                _ => {},
            },
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => match name {
                "Foldable10" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable10 = Some(value); return self; },
                "Traversable1" => { perceus_ptr::PerceusPtr::make_mut(r).Traversable1 = Some(value); return self; },
                "sequence1" => { perceus_ptr::PerceusPtr::make_mut(r).sequence1 = Some(value); return self; },
                "traverse1" => { perceus_ptr::PerceusPtr::make_mut(r).traverse1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => match name {
                "FoldableWithIndex1" => { perceus_ptr::PerceusPtr::make_mut(r).FoldableWithIndex1 = Some(value); return self; },
                "FunctorWithIndex0" => { perceus_ptr::PerceusPtr::make_mut(r).FunctorWithIndex0 = Some(value); return self; },
                "Traversable2" => { perceus_ptr::PerceusPtr::make_mut(r).Traversable2 = Some(value); return self; },
                "traverseWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).traverseWithIndex = Some(value); return self; },
                _ => {},
            },
            Value::Record_Functor0_alt(r) => match name {
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "alt" => { perceus_ptr::PerceusPtr::make_mut(r).alt = Some(value); return self; },
                _ => {},
            },
            Value::Record_Functor0_apply(r) => match name {
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "apply" => { perceus_ptr::PerceusPtr::make_mut(r).apply = Some(value); return self; },
                _ => {},
            },
            Value::Record_Functor0_collect_distribute(r) => match name {
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "collect" => { perceus_ptr::PerceusPtr::make_mut(r).collect = Some(value); return self; },
                "distribute" => { perceus_ptr::PerceusPtr::make_mut(r).distribute = Some(value); return self; },
                _ => {},
            },
            Value::Record_Functor0_extend(r) => match name {
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "extend" => { perceus_ptr::PerceusPtr::make_mut(r).extend = Some(value); return self; },
                _ => {},
            },
            Value::Record_Functor0_mapWithIndex(r) => match name {
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "mapWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).mapWithIndex = Some(value); return self; },
                _ => {},
            },
            Value::Record_HeytingAlgebra0(r) => match name {
                "HeytingAlgebra0" => { perceus_ptr::PerceusPtr::make_mut(r).HeytingAlgebra0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_HeytingAlgebraRecord0(r) => match name {
                "HeytingAlgebraRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).HeytingAlgebraRecord0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_ask(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "ask" => { perceus_ptr::PerceusPtr::make_mut(r).ask = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_callCC(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "callCC" => { perceus_ptr::PerceusPtr::make_mut(r).callCC = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_liftEffect(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "liftEffect" => { perceus_ptr::PerceusPtr::make_mut(r).liftEffect = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_liftST(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "liftST" => { perceus_ptr::PerceusPtr::make_mut(r).liftST = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_state(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "state" => { perceus_ptr::PerceusPtr::make_mut(r).state = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_tailRecM(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "tailRecM" => { perceus_ptr::PerceusPtr::make_mut(r).tailRecM = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad0_throwError(r) => match name {
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "throwError" => { perceus_ptr::PerceusPtr::make_mut(r).throwError = Some(value); return self; },
                _ => {},
            },
            Value::Record_Monad1_Semigroup0_tell(r) => match name {
                "Monad1" => { perceus_ptr::PerceusPtr::make_mut(r).Monad1 = Some(value); return self; },
                "Semigroup0" => { perceus_ptr::PerceusPtr::make_mut(r).Semigroup0 = Some(value); return self; },
                "tell" => { perceus_ptr::PerceusPtr::make_mut(r).tell = Some(value); return self; },
                _ => {},
            },
            Value::Record_MonadAsk0_local(r) => match name {
                "MonadAsk0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadAsk0 = Some(value); return self; },
                "local" => { perceus_ptr::PerceusPtr::make_mut(r).local = Some(value); return self; },
                _ => {},
            },
            Value::Record_MonadEffect0_liftAff(r) => match name {
                "MonadEffect0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadEffect0 = Some(value); return self; },
                "liftAff" => { perceus_ptr::PerceusPtr::make_mut(r).liftAff = Some(value); return self; },
                _ => {},
            },
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => match name {
                "MonadTell1" => { perceus_ptr::PerceusPtr::make_mut(r).MonadTell1 = Some(value); return self; },
                "Monoid0" => { perceus_ptr::PerceusPtr::make_mut(r).Monoid0 = Some(value); return self; },
                "listen" => { perceus_ptr::PerceusPtr::make_mut(r).listen = Some(value); return self; },
                "pass" => { perceus_ptr::PerceusPtr::make_mut(r).pass = Some(value); return self; },
                _ => {},
            },
            Value::Record_MonadThrow0_catchError(r) => match name {
                "MonadThrow0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadThrow0 = Some(value); return self; },
                "catchError" => { perceus_ptr::PerceusPtr::make_mut(r).catchError = Some(value); return self; },
                _ => {},
            },
            Value::Record_Ord0_bottom_top(r) => match name {
                "Ord0" => { perceus_ptr::PerceusPtr::make_mut(r).Ord0 = Some(value); return self; },
                "bottom" => { perceus_ptr::PerceusPtr::make_mut(r).bottom = Some(value); return self; },
                "top" => { perceus_ptr::PerceusPtr::make_mut(r).top = Some(value); return self; },
                _ => {},
            },
            Value::Record_Ord0_pred_succ(r) => match name {
                "Ord0" => { perceus_ptr::PerceusPtr::make_mut(r).Ord0 = Some(value); return self; },
                "pred" => { perceus_ptr::PerceusPtr::make_mut(r).pred = Some(value); return self; },
                "succ" => { perceus_ptr::PerceusPtr::make_mut(r).succ = Some(value); return self; },
                _ => {},
            },
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => match name {
                "OrdRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).OrdRecord0 = Some(value); return self; },
                "bottomRecord" => { perceus_ptr::PerceusPtr::make_mut(r).bottomRecord = Some(value); return self; },
                "topRecord" => { perceus_ptr::PerceusPtr::make_mut(r).topRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_Profunctor0_closed(r) => match name {
                "Profunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Profunctor0 = Some(value); return self; },
                "closed" => { perceus_ptr::PerceusPtr::make_mut(r).closed = Some(value); return self; },
                _ => {},
            },
            Value::Record_Profunctor0_first_second(r) => match name {
                "Profunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Profunctor0 = Some(value); return self; },
                "first" => { perceus_ptr::PerceusPtr::make_mut(r).first = Some(value); return self; },
                "second" => { perceus_ptr::PerceusPtr::make_mut(r).second = Some(value); return self; },
                _ => {},
            },
            Value::Record_Profunctor0_left_right(r) => match name {
                "Profunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Profunctor0 = Some(value); return self; },
                "left" => { perceus_ptr::PerceusPtr::make_mut(r).left = Some(value); return self; },
                "right" => { perceus_ptr::PerceusPtr::make_mut(r).right = Some(value); return self; },
                _ => {},
            },
            Value::Record_Ring0(r) => match name {
                "Ring0" => { perceus_ptr::PerceusPtr::make_mut(r).Ring0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Ring0_recip(r) => match name {
                "Ring0" => { perceus_ptr::PerceusPtr::make_mut(r).Ring0 = Some(value); return self; },
                "recip" => { perceus_ptr::PerceusPtr::make_mut(r).recip = Some(value); return self; },
                _ => {},
            },
            Value::Record_RingRecord0(r) => match name {
                "RingRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).RingRecord0 = Some(value); return self; },
                _ => {},
            },
            Value::Record_Semigroup0_mempty(r) => match name {
                "Semigroup0" => { perceus_ptr::PerceusPtr::make_mut(r).Semigroup0 = Some(value); return self; },
                "mempty" => { perceus_ptr::PerceusPtr::make_mut(r).mempty = Some(value); return self; },
                _ => {},
            },
            Value::Record_SemigroupRecord0_memptyRecord(r) => match name {
                "SemigroupRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).SemigroupRecord0 = Some(value); return self; },
                "memptyRecord" => { perceus_ptr::PerceusPtr::make_mut(r).memptyRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_Semigroupoid0_identity(r) => match name {
                "Semigroupoid0" => { perceus_ptr::PerceusPtr::make_mut(r).Semigroupoid0 = Some(value); return self; },
                "identity" => { perceus_ptr::PerceusPtr::make_mut(r).identity = Some(value); return self; },
                _ => {},
            },
            Value::Record_Semiring0_sub(r) => match name {
                "Semiring0" => { perceus_ptr::PerceusPtr::make_mut(r).Semiring0 = Some(value); return self; },
                "sub" => { perceus_ptr::PerceusPtr::make_mut(r).sub = Some(value); return self; },
                _ => {},
            },
            Value::Record_SemiringRecord0_subRecord(r) => match name {
                "SemiringRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).SemiringRecord0 = Some(value); return self; },
                "subRecord" => { perceus_ptr::PerceusPtr::make_mut(r).subRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_Unfoldable10_unfoldr(r) => match name {
                "Unfoldable10" => { perceus_ptr::PerceusPtr::make_mut(r).Unfoldable10 = Some(value); return self; },
                "unfoldr" => { perceus_ptr::PerceusPtr::make_mut(r).unfoldr = Some(value); return self; },
                _ => {},
            },
            Value::Record_a_b(r) => match name {
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                _ => {},
            },
            Value::Record_a_b_c(r) => match name {
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                _ => {},
            },
            Value::Record_a_b_c_d_e(r) => match name {
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                "d" => { perceus_ptr::PerceusPtr::make_mut(r).d = Some(value); return self; },
                "e" => { perceus_ptr::PerceusPtr::make_mut(r).e = Some(value); return self; },
                _ => {},
            },
            Value::Record_acc_init(r) => match name {
                "acc" => { perceus_ptr::PerceusPtr::make_mut(r).acc = Some(value); return self; },
                "init" => { perceus_ptr::PerceusPtr::make_mut(r).init = Some(value); return self; },
                _ => {},
            },
            Value::Record_acc_val(r) => match name {
                "acc" => { perceus_ptr::PerceusPtr::make_mut(r).acc = Some(value); return self; },
                "val" => { perceus_ptr::PerceusPtr::make_mut(r).val = Some(value); return self; },
                _ => {},
            },
            Value::Record_accum_value(r) => match name {
                "accum" => { perceus_ptr::PerceusPtr::make_mut(r).accum = Some(value); return self; },
                "value" => { perceus_ptr::PerceusPtr::make_mut(r).value = Some(value); return self; },
                _ => {},
            },
            Value::Record_add_mul_one_zero(r) => match name {
                "add" => { perceus_ptr::PerceusPtr::make_mut(r).add = Some(value); return self; },
                "mul" => { perceus_ptr::PerceusPtr::make_mut(r).mul = Some(value); return self; },
                "one" => { perceus_ptr::PerceusPtr::make_mut(r).one = Some(value); return self; },
                "zero" => { perceus_ptr::PerceusPtr::make_mut(r).zero = Some(value); return self; },
                _ => {},
            },
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => match name {
                "addRecord" => { perceus_ptr::PerceusPtr::make_mut(r).addRecord = Some(value); return self; },
                "mulRecord" => { perceus_ptr::PerceusPtr::make_mut(r).mulRecord = Some(value); return self; },
                "oneRecord" => { perceus_ptr::PerceusPtr::make_mut(r).oneRecord = Some(value); return self; },
                "zeroRecord" => { perceus_ptr::PerceusPtr::make_mut(r).zeroRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_after_before(r) => match name {
                "after" => { perceus_ptr::PerceusPtr::make_mut(r).after = Some(value); return self; },
                "before" => { perceus_ptr::PerceusPtr::make_mut(r).before = Some(value); return self; },
                _ => {},
            },
            Value::Record_append(r) => match name {
                "append" => { perceus_ptr::PerceusPtr::make_mut(r).append = Some(value); return self; },
                _ => {},
            },
            Value::Record_appendRecord(r) => match name {
                "appendRecord" => { perceus_ptr::PerceusPtr::make_mut(r).appendRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_asList_asMap(r) => match name {
                "asList" => { perceus_ptr::PerceusPtr::make_mut(r).asList = Some(value); return self; },
                "asMap" => { perceus_ptr::PerceusPtr::make_mut(r).asMap = Some(value); return self; },
                _ => {},
            },
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => match name {
                "bifoldMap" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldMap = Some(value); return self; },
                "bifoldl" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldl = Some(value); return self; },
                "bifoldr" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldr = Some(value); return self; },
                _ => {},
            },
            Value::Record_bimap(r) => match name {
                "bimap" => { perceus_ptr::PerceusPtr::make_mut(r).bimap = Some(value); return self; },
                _ => {},
            },
            Value::Record_c_d(r) => match name {
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                "d" => { perceus_ptr::PerceusPtr::make_mut(r).d = Some(value); return self; },
                _ => {},
            },
            Value::Record_cmap(r) => match name {
                "cmap" => { perceus_ptr::PerceusPtr::make_mut(r).cmap = Some(value); return self; },
                _ => {},
            },
            Value::Record_completed_failed_killed(r) => match name {
                "completed" => { perceus_ptr::PerceusPtr::make_mut(r).completed = Some(value); return self; },
                "failed" => { perceus_ptr::PerceusPtr::make_mut(r).failed = Some(value); return self; },
                "killed" => { perceus_ptr::PerceusPtr::make_mut(r).killed = Some(value); return self; },
                _ => {},
            },
            Value::Record_compose(r) => match name {
                "compose" => { perceus_ptr::PerceusPtr::make_mut(r).compose = Some(value); return self; },
                _ => {},
            },
            Value::Record_conj_disj_ff_implies_not_tt(r) => match name {
                "conj" => { perceus_ptr::PerceusPtr::make_mut(r).conj = Some(value); return self; },
                "disj" => { perceus_ptr::PerceusPtr::make_mut(r).disj = Some(value); return self; },
                "ff" => { perceus_ptr::PerceusPtr::make_mut(r).ff = Some(value); return self; },
                "implies" => { perceus_ptr::PerceusPtr::make_mut(r).implies = Some(value); return self; },
                "not" => { perceus_ptr::PerceusPtr::make_mut(r).not = Some(value); return self; },
                "tt" => { perceus_ptr::PerceusPtr::make_mut(r).tt = Some(value); return self; },
                _ => {},
            },
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => match name {
                "conjRecord" => { perceus_ptr::PerceusPtr::make_mut(r).conjRecord = Some(value); return self; },
                "disjRecord" => { perceus_ptr::PerceusPtr::make_mut(r).disjRecord = Some(value); return self; },
                "ffRecord" => { perceus_ptr::PerceusPtr::make_mut(r).ffRecord = Some(value); return self; },
                "impliesRecord" => { perceus_ptr::PerceusPtr::make_mut(r).impliesRecord = Some(value); return self; },
                "notRecord" => { perceus_ptr::PerceusPtr::make_mut(r).notRecord = Some(value); return self; },
                "ttRecord" => { perceus_ptr::PerceusPtr::make_mut(r).ttRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => match name {
                "day" => { perceus_ptr::PerceusPtr::make_mut(r).day = Some(value); return self; },
                "hour" => { perceus_ptr::PerceusPtr::make_mut(r).hour = Some(value); return self; },
                "millisecond" => { perceus_ptr::PerceusPtr::make_mut(r).millisecond = Some(value); return self; },
                "minute" => { perceus_ptr::PerceusPtr::make_mut(r).minute = Some(value); return self; },
                "month" => { perceus_ptr::PerceusPtr::make_mut(r).month = Some(value); return self; },
                "second" => { perceus_ptr::PerceusPtr::make_mut(r).second = Some(value); return self; },
                "year" => { perceus_ptr::PerceusPtr::make_mut(r).year = Some(value); return self; },
                _ => {},
            },
            Value::Record_defer(r) => match name {
                "defer" => { perceus_ptr::PerceusPtr::make_mut(r).defer = Some(value); return self; },
                _ => {},
            },
            Value::Record_dimap(r) => match name {
                "dimap" => { perceus_ptr::PerceusPtr::make_mut(r).dimap = Some(value); return self; },
                _ => {},
            },
            Value::Record_discard(r) => match name {
                "discard" => { perceus_ptr::PerceusPtr::make_mut(r).discard = Some(value); return self; },
                _ => {},
            },
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => match name {
                "dotAll" => { perceus_ptr::PerceusPtr::make_mut(r).dotAll = Some(value); return self; },
                "global" => { perceus_ptr::PerceusPtr::make_mut(r).global = Some(value); return self; },
                "ignoreCase" => { perceus_ptr::PerceusPtr::make_mut(r).ignoreCase = Some(value); return self; },
                "multiline" => { perceus_ptr::PerceusPtr::make_mut(r).multiline = Some(value); return self; },
                "sticky" => { perceus_ptr::PerceusPtr::make_mut(r).sticky = Some(value); return self; },
                "unicode" => { perceus_ptr::PerceusPtr::make_mut(r).unicode = Some(value); return self; },
                _ => {},
            },
            Value::Record_e_f(r) => match name {
                "e" => { perceus_ptr::PerceusPtr::make_mut(r).e = Some(value); return self; },
                "f" => { perceus_ptr::PerceusPtr::make_mut(r).f = Some(value); return self; },
                _ => {},
            },
            Value::Record_elem_pos(r) => match name {
                "elem" => { perceus_ptr::PerceusPtr::make_mut(r).elem = Some(value); return self; },
                "pos" => { perceus_ptr::PerceusPtr::make_mut(r).pos = Some(value); return self; },
                _ => {},
            },
            Value::Record_eq(r) => match name {
                "eq" => { perceus_ptr::PerceusPtr::make_mut(r).eq = Some(value); return self; },
                _ => {},
            },
            Value::Record_eq1(r) => match name {
                "eq1" => { perceus_ptr::PerceusPtr::make_mut(r).eq1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_eqRecord(r) => match name {
                "eqRecord" => { perceus_ptr::PerceusPtr::make_mut(r).eqRecord = Some(value); return self; },
                _ => {},
            },
            Value::Record_fiber_supervisor(r) => match name {
                "fiber" => { perceus_ptr::PerceusPtr::make_mut(r).fiber = Some(value); return self; },
                "supervisor" => { perceus_ptr::PerceusPtr::make_mut(r).supervisor = Some(value); return self; },
                _ => {},
            },
            Value::Record_foldMap_foldl_foldr(r) => match name {
                "foldMap" => { perceus_ptr::PerceusPtr::make_mut(r).foldMap = Some(value); return self; },
                "foldl" => { perceus_ptr::PerceusPtr::make_mut(r).foldl = Some(value); return self; },
                "foldr" => { perceus_ptr::PerceusPtr::make_mut(r).foldr = Some(value); return self; },
                _ => {},
            },
            Value::Record_found_result(r) => match name {
                "found" => { perceus_ptr::PerceusPtr::make_mut(r).found = Some(value); return self; },
                "result" => { perceus_ptr::PerceusPtr::make_mut(r).result = Some(value); return self; },
                _ => {},
            },
            Value::Record_from_to(r) => match name {
                "from" => { perceus_ptr::PerceusPtr::make_mut(r).from = Some(value); return self; },
                "to" => { perceus_ptr::PerceusPtr::make_mut(r).to = Some(value); return self; },
                _ => {},
            },
            Value::Record_fromDuration_toDuration(r) => match name {
                "fromDuration" => { perceus_ptr::PerceusPtr::make_mut(r).fromDuration = Some(value); return self; },
                "toDuration" => { perceus_ptr::PerceusPtr::make_mut(r).toDuration = Some(value); return self; },
                _ => {},
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => match name {
                "fromLeft" => { perceus_ptr::PerceusPtr::make_mut(r).fromLeft = Some(value); return self; },
                "fromRight" => { perceus_ptr::PerceusPtr::make_mut(r).fromRight = Some(value); return self; },
                "isLeft" => { perceus_ptr::PerceusPtr::make_mut(r).isLeft = Some(value); return self; },
                "left" => { perceus_ptr::PerceusPtr::make_mut(r).left = Some(value); return self; },
                "right" => { perceus_ptr::PerceusPtr::make_mut(r).right = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => match name {
                "genericAdd'" => { perceus_ptr::PerceusPtr::make_mut(r).genericAdd_prime = Some(value); return self; },
                "genericMul'" => { perceus_ptr::PerceusPtr::make_mut(r).genericMul_prime = Some(value); return self; },
                "genericOne'" => { perceus_ptr::PerceusPtr::make_mut(r).genericOne_prime = Some(value); return self; },
                "genericZero'" => { perceus_ptr::PerceusPtr::make_mut(r).genericZero_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericAppend_prime(r) => match name {
                "genericAppend'" => { perceus_ptr::PerceusPtr::make_mut(r).genericAppend_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericBottom_prime(r) => match name {
                "genericBottom'" => { perceus_ptr::PerceusPtr::make_mut(r).genericBottom_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => match name {
                "genericCardinality'" => { perceus_ptr::PerceusPtr::make_mut(r).genericCardinality_prime = Some(value); return self; },
                "genericFromEnum'" => { perceus_ptr::PerceusPtr::make_mut(r).genericFromEnum_prime = Some(value); return self; },
                "genericToEnum'" => { perceus_ptr::PerceusPtr::make_mut(r).genericToEnum_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericCompare_prime(r) => match name {
                "genericCompare'" => { perceus_ptr::PerceusPtr::make_mut(r).genericCompare_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => match name {
                "genericConj'" => { perceus_ptr::PerceusPtr::make_mut(r).genericConj_prime = Some(value); return self; },
                "genericDisj'" => { perceus_ptr::PerceusPtr::make_mut(r).genericDisj_prime = Some(value); return self; },
                "genericFF'" => { perceus_ptr::PerceusPtr::make_mut(r).genericFF_prime = Some(value); return self; },
                "genericImplies'" => { perceus_ptr::PerceusPtr::make_mut(r).genericImplies_prime = Some(value); return self; },
                "genericNot'" => { perceus_ptr::PerceusPtr::make_mut(r).genericNot_prime = Some(value); return self; },
                "genericTT'" => { perceus_ptr::PerceusPtr::make_mut(r).genericTT_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericEq_prime(r) => match name {
                "genericEq'" => { perceus_ptr::PerceusPtr::make_mut(r).genericEq_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericMempty_prime(r) => match name {
                "genericMempty'" => { perceus_ptr::PerceusPtr::make_mut(r).genericMempty_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericPred_prime_genericSucc_prime(r) => match name {
                "genericPred'" => { perceus_ptr::PerceusPtr::make_mut(r).genericPred_prime = Some(value); return self; },
                "genericSucc'" => { perceus_ptr::PerceusPtr::make_mut(r).genericSucc_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericShow_prime(r) => match name {
                "genericShow'" => { perceus_ptr::PerceusPtr::make_mut(r).genericShow_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericShowArgs(r) => match name {
                "genericShowArgs" => { perceus_ptr::PerceusPtr::make_mut(r).genericShowArgs = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericSub_prime(r) => match name {
                "genericSub'" => { perceus_ptr::PerceusPtr::make_mut(r).genericSub_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_genericTop_prime(r) => match name {
                "genericTop'" => { perceus_ptr::PerceusPtr::make_mut(r).genericTop_prime = Some(value); return self; },
                _ => {},
            },
            Value::Record_handler_rethrow(r) => match name {
                "handler" => { perceus_ptr::PerceusPtr::make_mut(r).handler = Some(value); return self; },
                "rethrow" => { perceus_ptr::PerceusPtr::make_mut(r).rethrow = Some(value); return self; },
                _ => {},
            },
            Value::Record_head_tail(r) => match name {
                "head" => { perceus_ptr::PerceusPtr::make_mut(r).head = Some(value); return self; },
                "tail" => { perceus_ptr::PerceusPtr::make_mut(r).tail = Some(value); return self; },
                _ => {},
            },
            Value::Record_imap(r) => match name {
                "imap" => { perceus_ptr::PerceusPtr::make_mut(r).imap = Some(value); return self; },
                _ => {},
            },
            Value::Record_index_value(r) => match name {
                "index" => { perceus_ptr::PerceusPtr::make_mut(r).index = Some(value); return self; },
                "value" => { perceus_ptr::PerceusPtr::make_mut(r).value = Some(value); return self; },
                _ => {},
            },
            Value::Record_init_last(r) => match name {
                "init" => { perceus_ptr::PerceusPtr::make_mut(r).init = Some(value); return self; },
                "last" => { perceus_ptr::PerceusPtr::make_mut(r).last = Some(value); return self; },
                _ => {},
            },
            Value::Record_init_rest(r) => match name {
                "init" => { perceus_ptr::PerceusPtr::make_mut(r).init = Some(value); return self; },
                "rest" => { perceus_ptr::PerceusPtr::make_mut(r).rest = Some(value); return self; },
                _ => {},
            },
            Value::Record_inj_prj(r) => match name {
                "inj" => { perceus_ptr::PerceusPtr::make_mut(r).inj = Some(value); return self; },
                "prj" => { perceus_ptr::PerceusPtr::make_mut(r).prj = Some(value); return self; },
                _ => {},
            },
            Value::Record_isSuspended_join_kill_onComplete_run(r) => match name {
                "isSuspended" => { perceus_ptr::PerceusPtr::make_mut(r).isSuspended = Some(value); return self; },
                "join" => { perceus_ptr::PerceusPtr::make_mut(r).join = Some(value); return self; },
                "kill" => { perceus_ptr::PerceusPtr::make_mut(r).kill = Some(value); return self; },
                "onComplete" => { perceus_ptr::PerceusPtr::make_mut(r).onComplete = Some(value); return self; },
                "run" => { perceus_ptr::PerceusPtr::make_mut(r).run = Some(value); return self; },
                _ => {},
            },
            Value::Record_key_value(r) => match name {
                "key" => { perceus_ptr::PerceusPtr::make_mut(r).key = Some(value); return self; },
                "value" => { perceus_ptr::PerceusPtr::make_mut(r).value = Some(value); return self; },
                _ => {},
            },
            Value::Record_keysImpl(r) => match name {
                "keysImpl" => { perceus_ptr::PerceusPtr::make_mut(r).keysImpl = Some(value); return self; },
                _ => {},
            },
            Value::Record_last_revInit(r) => match name {
                "last" => { perceus_ptr::PerceusPtr::make_mut(r).last = Some(value); return self; },
                "revInit" => { perceus_ptr::PerceusPtr::make_mut(r).revInit = Some(value); return self; },
                _ => {},
            },
            Value::Record_lift(r) => match name {
                "lift" => { perceus_ptr::PerceusPtr::make_mut(r).lift = Some(value); return self; },
                _ => {},
            },
            Value::Record_lower(r) => match name {
                "lower" => { perceus_ptr::PerceusPtr::make_mut(r).lower = Some(value); return self; },
                _ => {},
            },
            Value::Record_map(r) => match name {
                "map" => { perceus_ptr::PerceusPtr::make_mut(r).map = Some(value); return self; },
                _ => {},
            },
            Value::Record_mappend__mempty_(r) => match name {
                "mappend_" => { perceus_ptr::PerceusPtr::make_mut(r).mappend_ = Some(value); return self; },
                "mempty_" => { perceus_ptr::PerceusPtr::make_mut(r).mempty_ = Some(value); return self; },
                _ => {},
            },
            Value::Record_myMethod(r) => match name {
                "myMethod" => { perceus_ptr::PerceusPtr::make_mut(r).myMethod = Some(value); return self; },
                _ => {},
            },
            Value::Record_nes(r) => match name {
                "nes" => { perceus_ptr::PerceusPtr::make_mut(r).nes = Some(value); return self; },
                _ => {},
            },
            Value::Record_no_yes(r) => match name {
                "no" => { perceus_ptr::PerceusPtr::make_mut(r).no = Some(value); return self; },
                "yes" => { perceus_ptr::PerceusPtr::make_mut(r).yes = Some(value); return self; },
                _ => {},
            },
            Value::Record_ps_minus_rust_minus_test(r) => match name {
                "ps-rust-test" => { perceus_ptr::PerceusPtr::make_mut(r).ps_minus_rust_minus_test = Some(value); return self; },
                _ => {},
            },
            Value::Record_reflectSymbol(r) => match name {
                "reflectSymbol" => { perceus_ptr::PerceusPtr::make_mut(r).reflectSymbol = Some(value); return self; },
                _ => {},
            },
            Value::Record_reflectType(r) => match name {
                "reflectType" => { perceus_ptr::PerceusPtr::make_mut(r).reflectType = Some(value); return self; },
                _ => {},
            },
            Value::Record_show(r) => match name {
                "show" => { perceus_ptr::PerceusPtr::make_mut(r).show = Some(value); return self; },
                _ => {},
            },
            Value::Record_showRecordFields(r) => match name {
                "showRecordFields" => { perceus_ptr::PerceusPtr::make_mut(r).showRecordFields = Some(value); return self; },
                _ => {},
            },
            Value::Record_state_val(r) => match name {
                "state" => { perceus_ptr::PerceusPtr::make_mut(r).state = Some(value); return self; },
                "val" => { perceus_ptr::PerceusPtr::make_mut(r).val = Some(value); return self; },
                _ => {},
            },
            Value::Record_state_value(r) => match name {
                "state" => { perceus_ptr::PerceusPtr::make_mut(r).state = Some(value); return self; },
                "value" => { perceus_ptr::PerceusPtr::make_mut(r).value = Some(value); return self; },
                _ => {},
            },
            Value::Record_unfoldr1(r) => match name {
                "unfoldr1" => { perceus_ptr::PerceusPtr::make_mut(r).unfoldr1 = Some(value); return self; },
                _ => {},
            },
            Value::Record_a(r) => match name {
                "Alt0" => { perceus_ptr::PerceusPtr::make_mut(r).Alt0 = Some(value); return self; },
                "Alternative1" => { perceus_ptr::PerceusPtr::make_mut(r).Alternative1 = Some(value); return self; },
                "Applicative0" => { perceus_ptr::PerceusPtr::make_mut(r).Applicative0 = Some(value); return self; },
                "Apply0" => { perceus_ptr::PerceusPtr::make_mut(r).Apply0 = Some(value); return self; },
                "Apply1" => { perceus_ptr::PerceusPtr::make_mut(r).Apply1 = Some(value); return self; },
                "Biapply0" => { perceus_ptr::PerceusPtr::make_mut(r).Biapply0 = Some(value); return self; },
                "Bifoldable1" => { perceus_ptr::PerceusPtr::make_mut(r).Bifoldable1 = Some(value); return self; },
                "Bifunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Bifunctor0 = Some(value); return self; },
                "Bind1" => { perceus_ptr::PerceusPtr::make_mut(r).Bind1 = Some(value); return self; },
                "Bounded0" => { perceus_ptr::PerceusPtr::make_mut(r).Bounded0 = Some(value); return self; },
                "Coercible0" => { perceus_ptr::PerceusPtr::make_mut(r).Coercible0 = Some(value); return self; },
                "CommutativeRing0" => { perceus_ptr::PerceusPtr::make_mut(r).CommutativeRing0 = Some(value); return self; },
                "Comonad0" => { perceus_ptr::PerceusPtr::make_mut(r).Comonad0 = Some(value); return self; },
                "ComonadAsk0" => { perceus_ptr::PerceusPtr::make_mut(r).ComonadAsk0 = Some(value); return self; },
                "Contravariant0" => { perceus_ptr::PerceusPtr::make_mut(r).Contravariant0 = Some(value); return self; },
                "Decide0" => { perceus_ptr::PerceusPtr::make_mut(r).Decide0 = Some(value); return self; },
                "Divide0" => { perceus_ptr::PerceusPtr::make_mut(r).Divide0 = Some(value); return self; },
                "Divisible1" => { perceus_ptr::PerceusPtr::make_mut(r).Divisible1 = Some(value); return self; },
                "DivisionRing1" => { perceus_ptr::PerceusPtr::make_mut(r).DivisionRing1 = Some(value); return self; },
                "Enum1" => { perceus_ptr::PerceusPtr::make_mut(r).Enum1 = Some(value); return self; },
                "Eq0" => { perceus_ptr::PerceusPtr::make_mut(r).Eq0 = Some(value); return self; },
                "Eq10" => { perceus_ptr::PerceusPtr::make_mut(r).Eq10 = Some(value); return self; },
                "EqRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).EqRecord0 = Some(value); return self; },
                "EuclideanRing0" => { perceus_ptr::PerceusPtr::make_mut(r).EuclideanRing0 = Some(value); return self; },
                "Extend0" => { perceus_ptr::PerceusPtr::make_mut(r).Extend0 = Some(value); return self; },
                "Foldable0" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable0 = Some(value); return self; },
                "Foldable1" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable1 = Some(value); return self; },
                "Foldable10" => { perceus_ptr::PerceusPtr::make_mut(r).Foldable10 = Some(value); return self; },
                "FoldableWithIndex1" => { perceus_ptr::PerceusPtr::make_mut(r).FoldableWithIndex1 = Some(value); return self; },
                "Functor0" => { perceus_ptr::PerceusPtr::make_mut(r).Functor0 = Some(value); return self; },
                "FunctorWithIndex0" => { perceus_ptr::PerceusPtr::make_mut(r).FunctorWithIndex0 = Some(value); return self; },
                "HeytingAlgebra0" => { perceus_ptr::PerceusPtr::make_mut(r).HeytingAlgebra0 = Some(value); return self; },
                "HeytingAlgebraRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).HeytingAlgebraRecord0 = Some(value); return self; },
                "Monad0" => { perceus_ptr::PerceusPtr::make_mut(r).Monad0 = Some(value); return self; },
                "Monad1" => { perceus_ptr::PerceusPtr::make_mut(r).Monad1 = Some(value); return self; },
                "MonadAsk0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadAsk0 = Some(value); return self; },
                "MonadEffect0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadEffect0 = Some(value); return self; },
                "MonadTell1" => { perceus_ptr::PerceusPtr::make_mut(r).MonadTell1 = Some(value); return self; },
                "MonadThrow0" => { perceus_ptr::PerceusPtr::make_mut(r).MonadThrow0 = Some(value); return self; },
                "Monoid0" => { perceus_ptr::PerceusPtr::make_mut(r).Monoid0 = Some(value); return self; },
                "Ord0" => { perceus_ptr::PerceusPtr::make_mut(r).Ord0 = Some(value); return self; },
                "OrdRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).OrdRecord0 = Some(value); return self; },
                "Plus1" => { perceus_ptr::PerceusPtr::make_mut(r).Plus1 = Some(value); return self; },
                "Profunctor0" => { perceus_ptr::PerceusPtr::make_mut(r).Profunctor0 = Some(value); return self; },
                "Ring0" => { perceus_ptr::PerceusPtr::make_mut(r).Ring0 = Some(value); return self; },
                "RingRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).RingRecord0 = Some(value); return self; },
                "Semigroup0" => { perceus_ptr::PerceusPtr::make_mut(r).Semigroup0 = Some(value); return self; },
                "SemigroupRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).SemigroupRecord0 = Some(value); return self; },
                "Semigroupoid0" => { perceus_ptr::PerceusPtr::make_mut(r).Semigroupoid0 = Some(value); return self; },
                "Semiring0" => { perceus_ptr::PerceusPtr::make_mut(r).Semiring0 = Some(value); return self; },
                "SemiringRecord0" => { perceus_ptr::PerceusPtr::make_mut(r).SemiringRecord0 = Some(value); return self; },
                "Traversable1" => { perceus_ptr::PerceusPtr::make_mut(r).Traversable1 = Some(value); return self; },
                "Traversable2" => { perceus_ptr::PerceusPtr::make_mut(r).Traversable2 = Some(value); return self; },
                "Unfoldable10" => { perceus_ptr::PerceusPtr::make_mut(r).Unfoldable10 = Some(value); return self; },
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "acc" => { perceus_ptr::PerceusPtr::make_mut(r).acc = Some(value); return self; },
                "accum" => { perceus_ptr::PerceusPtr::make_mut(r).accum = Some(value); return self; },
                "add" => { perceus_ptr::PerceusPtr::make_mut(r).add = Some(value); return self; },
                "addRecord" => { perceus_ptr::PerceusPtr::make_mut(r).addRecord = Some(value); return self; },
                "after" => { perceus_ptr::PerceusPtr::make_mut(r).after = Some(value); return self; },
                "alt" => { perceus_ptr::PerceusPtr::make_mut(r).alt = Some(value); return self; },
                "append" => { perceus_ptr::PerceusPtr::make_mut(r).append = Some(value); return self; },
                "appendRecord" => { perceus_ptr::PerceusPtr::make_mut(r).appendRecord = Some(value); return self; },
                "apply" => { perceus_ptr::PerceusPtr::make_mut(r).apply = Some(value); return self; },
                "asList" => { perceus_ptr::PerceusPtr::make_mut(r).asList = Some(value); return self; },
                "asMap" => { perceus_ptr::PerceusPtr::make_mut(r).asMap = Some(value); return self; },
                "ask" => { perceus_ptr::PerceusPtr::make_mut(r).ask = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                "before" => { perceus_ptr::PerceusPtr::make_mut(r).before = Some(value); return self; },
                "biapply" => { perceus_ptr::PerceusPtr::make_mut(r).biapply = Some(value); return self; },
                "bifoldMap" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldMap = Some(value); return self; },
                "bifoldl" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldl = Some(value); return self; },
                "bifoldr" => { perceus_ptr::PerceusPtr::make_mut(r).bifoldr = Some(value); return self; },
                "bimap" => { perceus_ptr::PerceusPtr::make_mut(r).bimap = Some(value); return self; },
                "bind" => { perceus_ptr::PerceusPtr::make_mut(r).bind = Some(value); return self; },
                "bipure" => { perceus_ptr::PerceusPtr::make_mut(r).bipure = Some(value); return self; },
                "bisequence" => { perceus_ptr::PerceusPtr::make_mut(r).bisequence = Some(value); return self; },
                "bitraverse" => { perceus_ptr::PerceusPtr::make_mut(r).bitraverse = Some(value); return self; },
                "bottom" => { perceus_ptr::PerceusPtr::make_mut(r).bottom = Some(value); return self; },
                "bottomRecord" => { perceus_ptr::PerceusPtr::make_mut(r).bottomRecord = Some(value); return self; },
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                "callCC" => { perceus_ptr::PerceusPtr::make_mut(r).callCC = Some(value); return self; },
                "cardinality" => { perceus_ptr::PerceusPtr::make_mut(r).cardinality = Some(value); return self; },
                "catchError" => { perceus_ptr::PerceusPtr::make_mut(r).catchError = Some(value); return self; },
                "choose" => { perceus_ptr::PerceusPtr::make_mut(r).choose = Some(value); return self; },
                "closed" => { perceus_ptr::PerceusPtr::make_mut(r).closed = Some(value); return self; },
                "cmap" => { perceus_ptr::PerceusPtr::make_mut(r).cmap = Some(value); return self; },
                "collect" => { perceus_ptr::PerceusPtr::make_mut(r).collect = Some(value); return self; },
                "compare" => { perceus_ptr::PerceusPtr::make_mut(r).compare = Some(value); return self; },
                "compare1" => { perceus_ptr::PerceusPtr::make_mut(r).compare1 = Some(value); return self; },
                "compareRecord" => { perceus_ptr::PerceusPtr::make_mut(r).compareRecord = Some(value); return self; },
                "completed" => { perceus_ptr::PerceusPtr::make_mut(r).completed = Some(value); return self; },
                "compose" => { perceus_ptr::PerceusPtr::make_mut(r).compose = Some(value); return self; },
                "conj" => { perceus_ptr::PerceusPtr::make_mut(r).conj = Some(value); return self; },
                "conjRecord" => { perceus_ptr::PerceusPtr::make_mut(r).conjRecord = Some(value); return self; },
                "conquer" => { perceus_ptr::PerceusPtr::make_mut(r).conquer = Some(value); return self; },
                "d" => { perceus_ptr::PerceusPtr::make_mut(r).d = Some(value); return self; },
                "day" => { perceus_ptr::PerceusPtr::make_mut(r).day = Some(value); return self; },
                "defer" => { perceus_ptr::PerceusPtr::make_mut(r).defer = Some(value); return self; },
                "degree" => { perceus_ptr::PerceusPtr::make_mut(r).degree = Some(value); return self; },
                "dimap" => { perceus_ptr::PerceusPtr::make_mut(r).dimap = Some(value); return self; },
                "discard" => { perceus_ptr::PerceusPtr::make_mut(r).discard = Some(value); return self; },
                "disj" => { perceus_ptr::PerceusPtr::make_mut(r).disj = Some(value); return self; },
                "disjRecord" => { perceus_ptr::PerceusPtr::make_mut(r).disjRecord = Some(value); return self; },
                "distribute" => { perceus_ptr::PerceusPtr::make_mut(r).distribute = Some(value); return self; },
                "div" => { perceus_ptr::PerceusPtr::make_mut(r).div = Some(value); return self; },
                "divide" => { perceus_ptr::PerceusPtr::make_mut(r).divide = Some(value); return self; },
                "dotAll" => { perceus_ptr::PerceusPtr::make_mut(r).dotAll = Some(value); return self; },
                "e" => { perceus_ptr::PerceusPtr::make_mut(r).e = Some(value); return self; },
                "elem" => { perceus_ptr::PerceusPtr::make_mut(r).elem = Some(value); return self; },
                "empty" => { perceus_ptr::PerceusPtr::make_mut(r).empty = Some(value); return self; },
                "eq" => { perceus_ptr::PerceusPtr::make_mut(r).eq = Some(value); return self; },
                "eq1" => { perceus_ptr::PerceusPtr::make_mut(r).eq1 = Some(value); return self; },
                "eqRecord" => { perceus_ptr::PerceusPtr::make_mut(r).eqRecord = Some(value); return self; },
                "extend" => { perceus_ptr::PerceusPtr::make_mut(r).extend = Some(value); return self; },
                "extract" => { perceus_ptr::PerceusPtr::make_mut(r).extract = Some(value); return self; },
                "f" => { perceus_ptr::PerceusPtr::make_mut(r).f = Some(value); return self; },
                "failed" => { perceus_ptr::PerceusPtr::make_mut(r).failed = Some(value); return self; },
                "ff" => { perceus_ptr::PerceusPtr::make_mut(r).ff = Some(value); return self; },
                "ffRecord" => { perceus_ptr::PerceusPtr::make_mut(r).ffRecord = Some(value); return self; },
                "fiber" => { perceus_ptr::PerceusPtr::make_mut(r).fiber = Some(value); return self; },
                "first" => { perceus_ptr::PerceusPtr::make_mut(r).first = Some(value); return self; },
                "foldMap" => { perceus_ptr::PerceusPtr::make_mut(r).foldMap = Some(value); return self; },
                "foldMap1" => { perceus_ptr::PerceusPtr::make_mut(r).foldMap1 = Some(value); return self; },
                "foldMapWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldMapWithIndex = Some(value); return self; },
                "foldl" => { perceus_ptr::PerceusPtr::make_mut(r).foldl = Some(value); return self; },
                "foldl1" => { perceus_ptr::PerceusPtr::make_mut(r).foldl1 = Some(value); return self; },
                "foldlWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldlWithIndex = Some(value); return self; },
                "foldr" => { perceus_ptr::PerceusPtr::make_mut(r).foldr = Some(value); return self; },
                "foldr1" => { perceus_ptr::PerceusPtr::make_mut(r).foldr1 = Some(value); return self; },
                "foldrWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).foldrWithIndex = Some(value); return self; },
                "found" => { perceus_ptr::PerceusPtr::make_mut(r).found = Some(value); return self; },
                "from" => { perceus_ptr::PerceusPtr::make_mut(r).from = Some(value); return self; },
                "fromDuration" => { perceus_ptr::PerceusPtr::make_mut(r).fromDuration = Some(value); return self; },
                "fromEnum" => { perceus_ptr::PerceusPtr::make_mut(r).fromEnum = Some(value); return self; },
                "fromLeft" => { perceus_ptr::PerceusPtr::make_mut(r).fromLeft = Some(value); return self; },
                "fromRight" => { perceus_ptr::PerceusPtr::make_mut(r).fromRight = Some(value); return self; },
                "genericAdd'" => { perceus_ptr::PerceusPtr::make_mut(r).genericAdd_prime = Some(value); return self; },
                "genericAppend'" => { perceus_ptr::PerceusPtr::make_mut(r).genericAppend_prime = Some(value); return self; },
                "genericBottom'" => { perceus_ptr::PerceusPtr::make_mut(r).genericBottom_prime = Some(value); return self; },
                "genericCardinality'" => { perceus_ptr::PerceusPtr::make_mut(r).genericCardinality_prime = Some(value); return self; },
                "genericCompare'" => { perceus_ptr::PerceusPtr::make_mut(r).genericCompare_prime = Some(value); return self; },
                "genericConj'" => { perceus_ptr::PerceusPtr::make_mut(r).genericConj_prime = Some(value); return self; },
                "genericDisj'" => { perceus_ptr::PerceusPtr::make_mut(r).genericDisj_prime = Some(value); return self; },
                "genericEq'" => { perceus_ptr::PerceusPtr::make_mut(r).genericEq_prime = Some(value); return self; },
                "genericFF'" => { perceus_ptr::PerceusPtr::make_mut(r).genericFF_prime = Some(value); return self; },
                "genericFromEnum'" => { perceus_ptr::PerceusPtr::make_mut(r).genericFromEnum_prime = Some(value); return self; },
                "genericImplies'" => { perceus_ptr::PerceusPtr::make_mut(r).genericImplies_prime = Some(value); return self; },
                "genericMempty'" => { perceus_ptr::PerceusPtr::make_mut(r).genericMempty_prime = Some(value); return self; },
                "genericMul'" => { perceus_ptr::PerceusPtr::make_mut(r).genericMul_prime = Some(value); return self; },
                "genericNot'" => { perceus_ptr::PerceusPtr::make_mut(r).genericNot_prime = Some(value); return self; },
                "genericOne'" => { perceus_ptr::PerceusPtr::make_mut(r).genericOne_prime = Some(value); return self; },
                "genericPred'" => { perceus_ptr::PerceusPtr::make_mut(r).genericPred_prime = Some(value); return self; },
                "genericShow'" => { perceus_ptr::PerceusPtr::make_mut(r).genericShow_prime = Some(value); return self; },
                "genericShowArgs" => { perceus_ptr::PerceusPtr::make_mut(r).genericShowArgs = Some(value); return self; },
                "genericSub'" => { perceus_ptr::PerceusPtr::make_mut(r).genericSub_prime = Some(value); return self; },
                "genericSucc'" => { perceus_ptr::PerceusPtr::make_mut(r).genericSucc_prime = Some(value); return self; },
                "genericTT'" => { perceus_ptr::PerceusPtr::make_mut(r).genericTT_prime = Some(value); return self; },
                "genericToEnum'" => { perceus_ptr::PerceusPtr::make_mut(r).genericToEnum_prime = Some(value); return self; },
                "genericTop'" => { perceus_ptr::PerceusPtr::make_mut(r).genericTop_prime = Some(value); return self; },
                "genericZero'" => { perceus_ptr::PerceusPtr::make_mut(r).genericZero_prime = Some(value); return self; },
                "global" => { perceus_ptr::PerceusPtr::make_mut(r).global = Some(value); return self; },
                "handler" => { perceus_ptr::PerceusPtr::make_mut(r).handler = Some(value); return self; },
                "head" => { perceus_ptr::PerceusPtr::make_mut(r).head = Some(value); return self; },
                "hour" => { perceus_ptr::PerceusPtr::make_mut(r).hour = Some(value); return self; },
                "identity" => { perceus_ptr::PerceusPtr::make_mut(r).identity = Some(value); return self; },
                "ignoreCase" => { perceus_ptr::PerceusPtr::make_mut(r).ignoreCase = Some(value); return self; },
                "imap" => { perceus_ptr::PerceusPtr::make_mut(r).imap = Some(value); return self; },
                "implies" => { perceus_ptr::PerceusPtr::make_mut(r).implies = Some(value); return self; },
                "impliesRecord" => { perceus_ptr::PerceusPtr::make_mut(r).impliesRecord = Some(value); return self; },
                "index" => { perceus_ptr::PerceusPtr::make_mut(r).index = Some(value); return self; },
                "init" => { perceus_ptr::PerceusPtr::make_mut(r).init = Some(value); return self; },
                "inj" => { perceus_ptr::PerceusPtr::make_mut(r).inj = Some(value); return self; },
                "isLeft" => { perceus_ptr::PerceusPtr::make_mut(r).isLeft = Some(value); return self; },
                "isSuspended" => { perceus_ptr::PerceusPtr::make_mut(r).isSuspended = Some(value); return self; },
                "join" => { perceus_ptr::PerceusPtr::make_mut(r).join = Some(value); return self; },
                "key" => { perceus_ptr::PerceusPtr::make_mut(r).key = Some(value); return self; },
                "keysImpl" => { perceus_ptr::PerceusPtr::make_mut(r).keysImpl = Some(value); return self; },
                "kill" => { perceus_ptr::PerceusPtr::make_mut(r).kill = Some(value); return self; },
                "killed" => { perceus_ptr::PerceusPtr::make_mut(r).killed = Some(value); return self; },
                "last" => { perceus_ptr::PerceusPtr::make_mut(r).last = Some(value); return self; },
                "left" => { perceus_ptr::PerceusPtr::make_mut(r).left = Some(value); return self; },
                "lift" => { perceus_ptr::PerceusPtr::make_mut(r).lift = Some(value); return self; },
                "liftAff" => { perceus_ptr::PerceusPtr::make_mut(r).liftAff = Some(value); return self; },
                "liftEffect" => { perceus_ptr::PerceusPtr::make_mut(r).liftEffect = Some(value); return self; },
                "liftST" => { perceus_ptr::PerceusPtr::make_mut(r).liftST = Some(value); return self; },
                "listen" => { perceus_ptr::PerceusPtr::make_mut(r).listen = Some(value); return self; },
                "local" => { perceus_ptr::PerceusPtr::make_mut(r).local = Some(value); return self; },
                "lose" => { perceus_ptr::PerceusPtr::make_mut(r).lose = Some(value); return self; },
                "lower" => { perceus_ptr::PerceusPtr::make_mut(r).lower = Some(value); return self; },
                "map" => { perceus_ptr::PerceusPtr::make_mut(r).map = Some(value); return self; },
                "mapWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).mapWithIndex = Some(value); return self; },
                "mappend_" => { perceus_ptr::PerceusPtr::make_mut(r).mappend_ = Some(value); return self; },
                "mempty" => { perceus_ptr::PerceusPtr::make_mut(r).mempty = Some(value); return self; },
                "memptyRecord" => { perceus_ptr::PerceusPtr::make_mut(r).memptyRecord = Some(value); return self; },
                "mempty_" => { perceus_ptr::PerceusPtr::make_mut(r).mempty_ = Some(value); return self; },
                "millisecond" => { perceus_ptr::PerceusPtr::make_mut(r).millisecond = Some(value); return self; },
                "minute" => { perceus_ptr::PerceusPtr::make_mut(r).minute = Some(value); return self; },
                "mod" => { perceus_ptr::PerceusPtr::make_mut(r).mod_kw = Some(value); return self; },
                "month" => { perceus_ptr::PerceusPtr::make_mut(r).month = Some(value); return self; },
                "mul" => { perceus_ptr::PerceusPtr::make_mut(r).mul = Some(value); return self; },
                "mulRecord" => { perceus_ptr::PerceusPtr::make_mut(r).mulRecord = Some(value); return self; },
                "multiline" => { perceus_ptr::PerceusPtr::make_mut(r).multiline = Some(value); return self; },
                "myMethod" => { perceus_ptr::PerceusPtr::make_mut(r).myMethod = Some(value); return self; },
                "nes" => { perceus_ptr::PerceusPtr::make_mut(r).nes = Some(value); return self; },
                "no" => { perceus_ptr::PerceusPtr::make_mut(r).no = Some(value); return self; },
                "not" => { perceus_ptr::PerceusPtr::make_mut(r).not = Some(value); return self; },
                "notRecord" => { perceus_ptr::PerceusPtr::make_mut(r).notRecord = Some(value); return self; },
                "onComplete" => { perceus_ptr::PerceusPtr::make_mut(r).onComplete = Some(value); return self; },
                "one" => { perceus_ptr::PerceusPtr::make_mut(r).one = Some(value); return self; },
                "oneRecord" => { perceus_ptr::PerceusPtr::make_mut(r).oneRecord = Some(value); return self; },
                "parallel" => { perceus_ptr::PerceusPtr::make_mut(r).parallel = Some(value); return self; },
                "pass" => { perceus_ptr::PerceusPtr::make_mut(r).pass = Some(value); return self; },
                "peek" => { perceus_ptr::PerceusPtr::make_mut(r).peek = Some(value); return self; },
                "pos" => { perceus_ptr::PerceusPtr::make_mut(r).pos = Some(value); return self; },
                "pred" => { perceus_ptr::PerceusPtr::make_mut(r).pred = Some(value); return self; },
                "prj" => { perceus_ptr::PerceusPtr::make_mut(r).prj = Some(value); return self; },
                "proof" => { perceus_ptr::PerceusPtr::make_mut(r).proof = Some(value); return self; },
                "ps-rust-test" => { perceus_ptr::PerceusPtr::make_mut(r).ps_minus_rust_minus_test = Some(value); return self; },
                "pure" => { perceus_ptr::PerceusPtr::make_mut(r).pure = Some(value); return self; },
                "recip" => { perceus_ptr::PerceusPtr::make_mut(r).recip = Some(value); return self; },
                "reflectSymbol" => { perceus_ptr::PerceusPtr::make_mut(r).reflectSymbol = Some(value); return self; },
                "reflectType" => { perceus_ptr::PerceusPtr::make_mut(r).reflectType = Some(value); return self; },
                "rest" => { perceus_ptr::PerceusPtr::make_mut(r).rest = Some(value); return self; },
                "result" => { perceus_ptr::PerceusPtr::make_mut(r).result = Some(value); return self; },
                "rethrow" => { perceus_ptr::PerceusPtr::make_mut(r).rethrow = Some(value); return self; },
                "revInit" => { perceus_ptr::PerceusPtr::make_mut(r).revInit = Some(value); return self; },
                "right" => { perceus_ptr::PerceusPtr::make_mut(r).right = Some(value); return self; },
                "run" => { perceus_ptr::PerceusPtr::make_mut(r).run = Some(value); return self; },
                "second" => { perceus_ptr::PerceusPtr::make_mut(r).second = Some(value); return self; },
                "sequence" => { perceus_ptr::PerceusPtr::make_mut(r).sequence = Some(value); return self; },
                "sequence1" => { perceus_ptr::PerceusPtr::make_mut(r).sequence1 = Some(value); return self; },
                "sequential" => { perceus_ptr::PerceusPtr::make_mut(r).sequential = Some(value); return self; },
                "show" => { perceus_ptr::PerceusPtr::make_mut(r).show = Some(value); return self; },
                "showRecordFields" => { perceus_ptr::PerceusPtr::make_mut(r).showRecordFields = Some(value); return self; },
                "state" => { perceus_ptr::PerceusPtr::make_mut(r).state = Some(value); return self; },
                "sticky" => { perceus_ptr::PerceusPtr::make_mut(r).sticky = Some(value); return self; },
                "sub" => { perceus_ptr::PerceusPtr::make_mut(r).sub = Some(value); return self; },
                "subRecord" => { perceus_ptr::PerceusPtr::make_mut(r).subRecord = Some(value); return self; },
                "succ" => { perceus_ptr::PerceusPtr::make_mut(r).succ = Some(value); return self; },
                "supervisor" => { perceus_ptr::PerceusPtr::make_mut(r).supervisor = Some(value); return self; },
                "tail" => { perceus_ptr::PerceusPtr::make_mut(r).tail = Some(value); return self; },
                "tailRecM" => { perceus_ptr::PerceusPtr::make_mut(r).tailRecM = Some(value); return self; },
                "tell" => { perceus_ptr::PerceusPtr::make_mut(r).tell = Some(value); return self; },
                "throwError" => { perceus_ptr::PerceusPtr::make_mut(r).throwError = Some(value); return self; },
                "to" => { perceus_ptr::PerceusPtr::make_mut(r).to = Some(value); return self; },
                "toDuration" => { perceus_ptr::PerceusPtr::make_mut(r).toDuration = Some(value); return self; },
                "toEnum" => { perceus_ptr::PerceusPtr::make_mut(r).toEnum = Some(value); return self; },
                "top" => { perceus_ptr::PerceusPtr::make_mut(r).top = Some(value); return self; },
                "topRecord" => { perceus_ptr::PerceusPtr::make_mut(r).topRecord = Some(value); return self; },
                "track" => { perceus_ptr::PerceusPtr::make_mut(r).track = Some(value); return self; },
                "traverse" => { perceus_ptr::PerceusPtr::make_mut(r).traverse = Some(value); return self; },
                "traverse1" => { perceus_ptr::PerceusPtr::make_mut(r).traverse1 = Some(value); return self; },
                "traverseWithIndex" => { perceus_ptr::PerceusPtr::make_mut(r).traverseWithIndex = Some(value); return self; },
                "tt" => { perceus_ptr::PerceusPtr::make_mut(r).tt = Some(value); return self; },
                "ttRecord" => { perceus_ptr::PerceusPtr::make_mut(r).ttRecord = Some(value); return self; },
                "unfoldr" => { perceus_ptr::PerceusPtr::make_mut(r).unfoldr = Some(value); return self; },
                "unfoldr1" => { perceus_ptr::PerceusPtr::make_mut(r).unfoldr1 = Some(value); return self; },
                "unicode" => { perceus_ptr::PerceusPtr::make_mut(r).unicode = Some(value); return self; },
                "val" => { perceus_ptr::PerceusPtr::make_mut(r).val = Some(value); return self; },
                "value" => { perceus_ptr::PerceusPtr::make_mut(r).value = Some(value); return self; },
                "year" => { perceus_ptr::PerceusPtr::make_mut(r).year = Some(value); return self; },
                "yes" => { perceus_ptr::PerceusPtr::make_mut(r).yes = Some(value); return self; },
                "zero" => { perceus_ptr::PerceusPtr::make_mut(r).zero = Some(value); return self; },
                "zeroRecord" => { perceus_ptr::PerceusPtr::make_mut(r).zeroRecord = Some(value); return self; },
                _ => {},
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert(name.to_owned(), value); return self; },
            _ => panic!("Expected record"),
        }
        let mut fields = RecordFields::new();
        match &self {
            Value::Record_Alt0_empty(r) => {
                if let Some(value) = &r.Alt0 { fields.insert("Alt0".to_owned(), value.clone()); }
                if let Some(value) = &r.empty { fields.insert("empty".to_owned(), value.clone()); }
            },
            Value::Record_Alternative1_Monad0(r) => {
                if let Some(value) = &r.Alternative1 { fields.insert("Alternative1".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
            },
            Value::Record_Applicative0_Bind1(r) => {
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bind1 { fields.insert("Bind1".to_owned(), value.clone()); }
            },
            Value::Record_Applicative0_Plus1(r) => {
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Plus1 { fields.insert("Plus1".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply1 { fields.insert("Apply1".to_owned(), value.clone()); }
                if let Some(value) = &r.parallel { fields.insert("parallel".to_owned(), value.clone()); }
                if let Some(value) = &r.sequential { fields.insert("sequential".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_bind(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.bind { fields.insert("bind".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_pure(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.pure { fields.insert("pure".to_owned(), value.clone()); }
            },
            Value::Record_Biapply0_bipure(r) => {
                if let Some(value) = &r.Biapply0 { fields.insert("Biapply0".to_owned(), value.clone()); }
                if let Some(value) = &r.bipure { fields.insert("bipure".to_owned(), value.clone()); }
            },
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                if let Some(value) = &r.Bifoldable1 { fields.insert("Bifoldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.bisequence { fields.insert("bisequence".to_owned(), value.clone()); }
                if let Some(value) = &r.bitraverse { fields.insert("bitraverse".to_owned(), value.clone()); }
            },
            Value::Record_Bifunctor0_biapply(r) => {
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.biapply { fields.insert("biapply".to_owned(), value.clone()); }
            },
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                if let Some(value) = &r.Bounded0 { fields.insert("Bounded0".to_owned(), value.clone()); }
                if let Some(value) = &r.Enum1 { fields.insert("Enum1".to_owned(), value.clone()); }
                if let Some(value) = &r.cardinality { fields.insert("cardinality".to_owned(), value.clone()); }
                if let Some(value) = &r.fromEnum { fields.insert("fromEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.toEnum { fields.insert("toEnum".to_owned(), value.clone()); }
            },
            Value::Record_Coercible0(r) => {
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
            },
            Value::Record_Coercible0_proof(r) => {
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
                if let Some(value) = &r.proof { fields.insert("proof".to_owned(), value.clone()); }
            },
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                if let Some(value) = &r.CommutativeRing0 { fields.insert("CommutativeRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.degree { fields.insert("degree".to_owned(), value.clone()); }
                if let Some(value) = &r.div { fields.insert("div".to_owned(), value.clone()); }
                if let Some(value) = &r.mod_kw { fields.insert("mod".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_ask(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_peek_pos(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.peek { fields.insert("peek".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_track(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.track { fields.insert("track".to_owned(), value.clone()); }
            },
            Value::Record_ComonadAsk0_local(r) => {
                if let Some(value) = &r.ComonadAsk0 { fields.insert("ComonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
            },
            Value::Record_Contravariant0_divide(r) => {
                if let Some(value) = &r.Contravariant0 { fields.insert("Contravariant0".to_owned(), value.clone()); }
                if let Some(value) = &r.divide { fields.insert("divide".to_owned(), value.clone()); }
            },
            Value::Record_Decide0_Divisible1_lose(r) => {
                if let Some(value) = &r.Decide0 { fields.insert("Decide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divisible1 { fields.insert("Divisible1".to_owned(), value.clone()); }
                if let Some(value) = &r.lose { fields.insert("lose".to_owned(), value.clone()); }
            },
            Value::Record_Divide0_choose(r) => {
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.choose { fields.insert("choose".to_owned(), value.clone()); }
            },
            Value::Record_Divide0_conquer(r) => {
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.conquer { fields.insert("conquer".to_owned(), value.clone()); }
            },
            Value::Record_DivisionRing1_EuclideanRing0(r) => {
                if let Some(value) = &r.DivisionRing1 { fields.insert("DivisionRing1".to_owned(), value.clone()); }
                if let Some(value) = &r.EuclideanRing0 { fields.insert("EuclideanRing0".to_owned(), value.clone()); }
            },
            Value::Record_Eq0_compare(r) => {
                if let Some(value) = &r.Eq0 { fields.insert("Eq0".to_owned(), value.clone()); }
                if let Some(value) = &r.compare { fields.insert("compare".to_owned(), value.clone()); }
            },
            Value::Record_Eq10_compare1(r) => {
                if let Some(value) = &r.Eq10 { fields.insert("Eq10".to_owned(), value.clone()); }
                if let Some(value) = &r.compare1 { fields.insert("compare1".to_owned(), value.clone()); }
            },
            Value::Record_EqRecord0_compareRecord(r) => {
                if let Some(value) = &r.EqRecord0 { fields.insert("EqRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.compareRecord { fields.insert("compareRecord".to_owned(), value.clone()); }
            },
            Value::Record_Extend0_extract(r) => {
                if let Some(value) = &r.Extend0 { fields.insert("Extend0".to_owned(), value.clone()); }
                if let Some(value) = &r.extract { fields.insert("extract".to_owned(), value.clone()); }
            },
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap1 { fields.insert("foldMap1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl1 { fields.insert("foldl1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr1 { fields.insert("foldr1".to_owned(), value.clone()); }
            },
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMapWithIndex { fields.insert("foldMapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldlWithIndex { fields.insert("foldlWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldrWithIndex { fields.insert("foldrWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                if let Some(value) = &r.Foldable1 { fields.insert("Foldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence { fields.insert("sequence".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse { fields.insert("traverse".to_owned(), value.clone()); }
            },
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                if let Some(value) = &r.Foldable10 { fields.insert("Foldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable1 { fields.insert("Traversable1".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence1 { fields.insert("sequence1".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse1 { fields.insert("traverse1".to_owned(), value.clone()); }
            },
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                if let Some(value) = &r.FoldableWithIndex1 { fields.insert("FoldableWithIndex1".to_owned(), value.clone()); }
                if let Some(value) = &r.FunctorWithIndex0 { fields.insert("FunctorWithIndex0".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable2 { fields.insert("Traversable2".to_owned(), value.clone()); }
                if let Some(value) = &r.traverseWithIndex { fields.insert("traverseWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_alt(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.alt { fields.insert("alt".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_apply(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.apply { fields.insert("apply".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_collect_distribute(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.collect { fields.insert("collect".to_owned(), value.clone()); }
                if let Some(value) = &r.distribute { fields.insert("distribute".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_extend(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.extend { fields.insert("extend".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_mapWithIndex(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.mapWithIndex { fields.insert("mapWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_HeytingAlgebra0(r) => {
                if let Some(value) = &r.HeytingAlgebra0 { fields.insert("HeytingAlgebra0".to_owned(), value.clone()); }
            },
            Value::Record_HeytingAlgebraRecord0(r) => {
                if let Some(value) = &r.HeytingAlgebraRecord0 { fields.insert("HeytingAlgebraRecord0".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_ask(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_callCC(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.callCC { fields.insert("callCC".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_liftEffect(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftEffect { fields.insert("liftEffect".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_liftST(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftST { fields.insert("liftST".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_state(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_tailRecM(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.tailRecM { fields.insert("tailRecM".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_throwError(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.throwError { fields.insert("throwError".to_owned(), value.clone()); }
            },
            Value::Record_Monad1_Semigroup0_tell(r) => {
                if let Some(value) = &r.Monad1 { fields.insert("Monad1".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.tell { fields.insert("tell".to_owned(), value.clone()); }
            },
            Value::Record_MonadAsk0_local(r) => {
                if let Some(value) = &r.MonadAsk0 { fields.insert("MonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
            },
            Value::Record_MonadEffect0_liftAff(r) => {
                if let Some(value) = &r.MonadEffect0 { fields.insert("MonadEffect0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftAff { fields.insert("liftAff".to_owned(), value.clone()); }
            },
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                if let Some(value) = &r.MonadTell1 { fields.insert("MonadTell1".to_owned(), value.clone()); }
                if let Some(value) = &r.Monoid0 { fields.insert("Monoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.listen { fields.insert("listen".to_owned(), value.clone()); }
                if let Some(value) = &r.pass { fields.insert("pass".to_owned(), value.clone()); }
            },
            Value::Record_MonadThrow0_catchError(r) => {
                if let Some(value) = &r.MonadThrow0 { fields.insert("MonadThrow0".to_owned(), value.clone()); }
                if let Some(value) = &r.catchError { fields.insert("catchError".to_owned(), value.clone()); }
            },
            Value::Record_Ord0_bottom_top(r) => {
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.bottom { fields.insert("bottom".to_owned(), value.clone()); }
                if let Some(value) = &r.top { fields.insert("top".to_owned(), value.clone()); }
            },
            Value::Record_Ord0_pred_succ(r) => {
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.pred { fields.insert("pred".to_owned(), value.clone()); }
                if let Some(value) = &r.succ { fields.insert("succ".to_owned(), value.clone()); }
            },
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => {
                if let Some(value) = &r.OrdRecord0 { fields.insert("OrdRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.bottomRecord { fields.insert("bottomRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.topRecord { fields.insert("topRecord".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_closed(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.closed { fields.insert("closed".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_first_second(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.first { fields.insert("first".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_left_right(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
            },
            Value::Record_Ring0(r) => {
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
            },
            Value::Record_Ring0_recip(r) => {
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
                if let Some(value) = &r.recip { fields.insert("recip".to_owned(), value.clone()); }
            },
            Value::Record_RingRecord0(r) => {
                if let Some(value) = &r.RingRecord0 { fields.insert("RingRecord0".to_owned(), value.clone()); }
            },
            Value::Record_Semigroup0_mempty(r) => {
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty { fields.insert("mempty".to_owned(), value.clone()); }
            },
            Value::Record_SemigroupRecord0_memptyRecord(r) => {
                if let Some(value) = &r.SemigroupRecord0 { fields.insert("SemigroupRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.memptyRecord { fields.insert("memptyRecord".to_owned(), value.clone()); }
            },
            Value::Record_Semigroupoid0_identity(r) => {
                if let Some(value) = &r.Semigroupoid0 { fields.insert("Semigroupoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.identity { fields.insert("identity".to_owned(), value.clone()); }
            },
            Value::Record_Semiring0_sub(r) => {
                if let Some(value) = &r.Semiring0 { fields.insert("Semiring0".to_owned(), value.clone()); }
                if let Some(value) = &r.sub { fields.insert("sub".to_owned(), value.clone()); }
            },
            Value::Record_SemiringRecord0_subRecord(r) => {
                if let Some(value) = &r.SemiringRecord0 { fields.insert("SemiringRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.subRecord { fields.insert("subRecord".to_owned(), value.clone()); }
            },
            Value::Record_Unfoldable10_unfoldr(r) => {
                if let Some(value) = &r.Unfoldable10 { fields.insert("Unfoldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr { fields.insert("unfoldr".to_owned(), value.clone()); }
            },
            Value::Record_a_b(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
            },
            Value::Record_a_b_c(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
            },
            Value::Record_a_b_c_d_e(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
            },
            Value::Record_acc_init(r) => {
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
            },
            Value::Record_acc_val(r) => {
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
            },
            Value::Record_accum_value(r) => {
                if let Some(value) = &r.accum { fields.insert("accum".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_add_mul_one_zero(r) => {
                if let Some(value) = &r.add { fields.insert("add".to_owned(), value.clone()); }
                if let Some(value) = &r.mul { fields.insert("mul".to_owned(), value.clone()); }
                if let Some(value) = &r.one { fields.insert("one".to_owned(), value.clone()); }
                if let Some(value) = &r.zero { fields.insert("zero".to_owned(), value.clone()); }
            },
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                if let Some(value) = &r.addRecord { fields.insert("addRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.mulRecord { fields.insert("mulRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.oneRecord { fields.insert("oneRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.zeroRecord { fields.insert("zeroRecord".to_owned(), value.clone()); }
            },
            Value::Record_after_before(r) => {
                if let Some(value) = &r.after { fields.insert("after".to_owned(), value.clone()); }
                if let Some(value) = &r.before { fields.insert("before".to_owned(), value.clone()); }
            },
            Value::Record_append(r) => {
                if let Some(value) = &r.append { fields.insert("append".to_owned(), value.clone()); }
            },
            Value::Record_appendRecord(r) => {
                if let Some(value) = &r.appendRecord { fields.insert("appendRecord".to_owned(), value.clone()); }
            },
            Value::Record_asList_asMap(r) => {
                if let Some(value) = &r.asList { fields.insert("asList".to_owned(), value.clone()); }
                if let Some(value) = &r.asMap { fields.insert("asMap".to_owned(), value.clone()); }
            },
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => {
                if let Some(value) = &r.bifoldMap { fields.insert("bifoldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldl { fields.insert("bifoldl".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldr { fields.insert("bifoldr".to_owned(), value.clone()); }
            },
            Value::Record_bimap(r) => {
                if let Some(value) = &r.bimap { fields.insert("bimap".to_owned(), value.clone()); }
            },
            Value::Record_c_d(r) => {
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
            },
            Value::Record_cmap(r) => {
                if let Some(value) = &r.cmap { fields.insert("cmap".to_owned(), value.clone()); }
            },
            Value::Record_completed_failed_killed(r) => {
                if let Some(value) = &r.completed { fields.insert("completed".to_owned(), value.clone()); }
                if let Some(value) = &r.failed { fields.insert("failed".to_owned(), value.clone()); }
                if let Some(value) = &r.killed { fields.insert("killed".to_owned(), value.clone()); }
            },
            Value::Record_compose(r) => {
                if let Some(value) = &r.compose { fields.insert("compose".to_owned(), value.clone()); }
            },
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                if let Some(value) = &r.conj { fields.insert("conj".to_owned(), value.clone()); }
                if let Some(value) = &r.disj { fields.insert("disj".to_owned(), value.clone()); }
                if let Some(value) = &r.ff { fields.insert("ff".to_owned(), value.clone()); }
                if let Some(value) = &r.implies { fields.insert("implies".to_owned(), value.clone()); }
                if let Some(value) = &r.not { fields.insert("not".to_owned(), value.clone()); }
                if let Some(value) = &r.tt { fields.insert("tt".to_owned(), value.clone()); }
            },
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                if let Some(value) = &r.conjRecord { fields.insert("conjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.disjRecord { fields.insert("disjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.ffRecord { fields.insert("ffRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.impliesRecord { fields.insert("impliesRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.notRecord { fields.insert("notRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.ttRecord { fields.insert("ttRecord".to_owned(), value.clone()); }
            },
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                if let Some(value) = &r.day { fields.insert("day".to_owned(), value.clone()); }
                if let Some(value) = &r.hour { fields.insert("hour".to_owned(), value.clone()); }
                if let Some(value) = &r.millisecond { fields.insert("millisecond".to_owned(), value.clone()); }
                if let Some(value) = &r.minute { fields.insert("minute".to_owned(), value.clone()); }
                if let Some(value) = &r.month { fields.insert("month".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
                if let Some(value) = &r.year { fields.insert("year".to_owned(), value.clone()); }
            },
            Value::Record_defer(r) => {
                if let Some(value) = &r.defer { fields.insert("defer".to_owned(), value.clone()); }
            },
            Value::Record_dimap(r) => {
                if let Some(value) = &r.dimap { fields.insert("dimap".to_owned(), value.clone()); }
            },
            Value::Record_discard(r) => {
                if let Some(value) = &r.discard { fields.insert("discard".to_owned(), value.clone()); }
            },
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                if let Some(value) = &r.dotAll { fields.insert("dotAll".to_owned(), value.clone()); }
                if let Some(value) = &r.global { fields.insert("global".to_owned(), value.clone()); }
                if let Some(value) = &r.ignoreCase { fields.insert("ignoreCase".to_owned(), value.clone()); }
                if let Some(value) = &r.multiline { fields.insert("multiline".to_owned(), value.clone()); }
                if let Some(value) = &r.sticky { fields.insert("sticky".to_owned(), value.clone()); }
                if let Some(value) = &r.unicode { fields.insert("unicode".to_owned(), value.clone()); }
            },
            Value::Record_e_f(r) => {
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
            },
            Value::Record_elem_pos(r) => {
                if let Some(value) = &r.elem { fields.insert("elem".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
            },
            Value::Record_eq(r) => {
                if let Some(value) = &r.eq { fields.insert("eq".to_owned(), value.clone()); }
            },
            Value::Record_eq1(r) => {
                if let Some(value) = &r.eq1 { fields.insert("eq1".to_owned(), value.clone()); }
            },
            Value::Record_eqRecord(r) => {
                if let Some(value) = &r.eqRecord { fields.insert("eqRecord".to_owned(), value.clone()); }
            },
            Value::Record_fiber_supervisor(r) => {
                if let Some(value) = &r.fiber { fields.insert("fiber".to_owned(), value.clone()); }
                if let Some(value) = &r.supervisor { fields.insert("supervisor".to_owned(), value.clone()); }
            },
            Value::Record_foldMap_foldl_foldr(r) => {
                if let Some(value) = &r.foldMap { fields.insert("foldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl { fields.insert("foldl".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr { fields.insert("foldr".to_owned(), value.clone()); }
            },
            Value::Record_found_result(r) => {
                if let Some(value) = &r.found { fields.insert("found".to_owned(), value.clone()); }
                if let Some(value) = &r.result { fields.insert("result".to_owned(), value.clone()); }
            },
            Value::Record_from_to(r) => {
                if let Some(value) = &r.from { fields.insert("from".to_owned(), value.clone()); }
                if let Some(value) = &r.to { fields.insert("to".to_owned(), value.clone()); }
            },
            Value::Record_fromDuration_toDuration(r) => {
                if let Some(value) = &r.fromDuration { fields.insert("fromDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.toDuration { fields.insert("toDuration".to_owned(), value.clone()); }
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                if let Some(value) = &r.fromLeft { fields.insert("fromLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.fromRight { fields.insert("fromRight".to_owned(), value.clone()); }
                if let Some(value) = &r.isLeft { fields.insert("isLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
            },
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                if let Some(value) = &r.genericAdd_prime { fields.insert("genericAdd'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMul_prime { fields.insert("genericMul'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericOne_prime { fields.insert("genericOne'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericZero_prime { fields.insert("genericZero'".to_owned(), value.clone()); }
            },
            Value::Record_genericAppend_prime(r) => {
                if let Some(value) = &r.genericAppend_prime { fields.insert("genericAppend'".to_owned(), value.clone()); }
            },
            Value::Record_genericBottom_prime(r) => {
                if let Some(value) = &r.genericBottom_prime { fields.insert("genericBottom'".to_owned(), value.clone()); }
            },
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => {
                if let Some(value) = &r.genericCardinality_prime { fields.insert("genericCardinality'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFromEnum_prime { fields.insert("genericFromEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericToEnum_prime { fields.insert("genericToEnum'".to_owned(), value.clone()); }
            },
            Value::Record_genericCompare_prime(r) => {
                if let Some(value) = &r.genericCompare_prime { fields.insert("genericCompare'".to_owned(), value.clone()); }
            },
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                if let Some(value) = &r.genericConj_prime { fields.insert("genericConj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericDisj_prime { fields.insert("genericDisj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFF_prime { fields.insert("genericFF'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericImplies_prime { fields.insert("genericImplies'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericNot_prime { fields.insert("genericNot'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTT_prime { fields.insert("genericTT'".to_owned(), value.clone()); }
            },
            Value::Record_genericEq_prime(r) => {
                if let Some(value) = &r.genericEq_prime { fields.insert("genericEq'".to_owned(), value.clone()); }
            },
            Value::Record_genericMempty_prime(r) => {
                if let Some(value) = &r.genericMempty_prime { fields.insert("genericMempty'".to_owned(), value.clone()); }
            },
            Value::Record_genericPred_prime_genericSucc_prime(r) => {
                if let Some(value) = &r.genericPred_prime { fields.insert("genericPred'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSucc_prime { fields.insert("genericSucc'".to_owned(), value.clone()); }
            },
            Value::Record_genericShow_prime(r) => {
                if let Some(value) = &r.genericShow_prime { fields.insert("genericShow'".to_owned(), value.clone()); }
            },
            Value::Record_genericShowArgs(r) => {
                if let Some(value) = &r.genericShowArgs { fields.insert("genericShowArgs".to_owned(), value.clone()); }
            },
            Value::Record_genericSub_prime(r) => {
                if let Some(value) = &r.genericSub_prime { fields.insert("genericSub'".to_owned(), value.clone()); }
            },
            Value::Record_genericTop_prime(r) => {
                if let Some(value) = &r.genericTop_prime { fields.insert("genericTop'".to_owned(), value.clone()); }
            },
            Value::Record_handler_rethrow(r) => {
                if let Some(value) = &r.handler { fields.insert("handler".to_owned(), value.clone()); }
                if let Some(value) = &r.rethrow { fields.insert("rethrow".to_owned(), value.clone()); }
            },
            Value::Record_head_tail(r) => {
                if let Some(value) = &r.head { fields.insert("head".to_owned(), value.clone()); }
                if let Some(value) = &r.tail { fields.insert("tail".to_owned(), value.clone()); }
            },
            Value::Record_imap(r) => {
                if let Some(value) = &r.imap { fields.insert("imap".to_owned(), value.clone()); }
            },
            Value::Record_index_value(r) => {
                if let Some(value) = &r.index { fields.insert("index".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_init_last(r) => {
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
            },
            Value::Record_init_rest(r) => {
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.rest { fields.insert("rest".to_owned(), value.clone()); }
            },
            Value::Record_inj_prj(r) => {
                if let Some(value) = &r.inj { fields.insert("inj".to_owned(), value.clone()); }
                if let Some(value) = &r.prj { fields.insert("prj".to_owned(), value.clone()); }
            },
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                if let Some(value) = &r.isSuspended { fields.insert("isSuspended".to_owned(), value.clone()); }
                if let Some(value) = &r.join { fields.insert("join".to_owned(), value.clone()); }
                if let Some(value) = &r.kill { fields.insert("kill".to_owned(), value.clone()); }
                if let Some(value) = &r.onComplete { fields.insert("onComplete".to_owned(), value.clone()); }
                if let Some(value) = &r.run { fields.insert("run".to_owned(), value.clone()); }
            },
            Value::Record_key_value(r) => {
                if let Some(value) = &r.key { fields.insert("key".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_keysImpl(r) => {
                if let Some(value) = &r.keysImpl { fields.insert("keysImpl".to_owned(), value.clone()); }
            },
            Value::Record_last_revInit(r) => {
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
                if let Some(value) = &r.revInit { fields.insert("revInit".to_owned(), value.clone()); }
            },
            Value::Record_lift(r) => {
                if let Some(value) = &r.lift { fields.insert("lift".to_owned(), value.clone()); }
            },
            Value::Record_lower(r) => {
                if let Some(value) = &r.lower { fields.insert("lower".to_owned(), value.clone()); }
            },
            Value::Record_map(r) => {
                if let Some(value) = &r.map { fields.insert("map".to_owned(), value.clone()); }
            },
            Value::Record_mappend__mempty_(r) => {
                if let Some(value) = &r.mappend_ { fields.insert("mappend_".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty_ { fields.insert("mempty_".to_owned(), value.clone()); }
            },
            Value::Record_myMethod(r) => {
                if let Some(value) = &r.myMethod { fields.insert("myMethod".to_owned(), value.clone()); }
            },
            Value::Record_nes(r) => {
                if let Some(value) = &r.nes { fields.insert("nes".to_owned(), value.clone()); }
            },
            Value::Record_no_yes(r) => {
                if let Some(value) = &r.no { fields.insert("no".to_owned(), value.clone()); }
                if let Some(value) = &r.yes { fields.insert("yes".to_owned(), value.clone()); }
            },
            Value::Record_ps_minus_rust_minus_test(r) => {
                if let Some(value) = &r.ps_minus_rust_minus_test { fields.insert("ps-rust-test".to_owned(), value.clone()); }
            },
            Value::Record_reflectSymbol(r) => {
                if let Some(value) = &r.reflectSymbol { fields.insert("reflectSymbol".to_owned(), value.clone()); }
            },
            Value::Record_reflectType(r) => {
                if let Some(value) = &r.reflectType { fields.insert("reflectType".to_owned(), value.clone()); }
            },
            Value::Record_show(r) => {
                if let Some(value) = &r.show { fields.insert("show".to_owned(), value.clone()); }
            },
            Value::Record_showRecordFields(r) => {
                if let Some(value) = &r.showRecordFields { fields.insert("showRecordFields".to_owned(), value.clone()); }
            },
            Value::Record_state_val(r) => {
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
            },
            Value::Record_state_value(r) => {
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_unfoldr1(r) => {
                if let Some(value) = &r.unfoldr1 { fields.insert("unfoldr1".to_owned(), value.clone()); }
            },
            Value::Record_a(r) => {
                if let Some(value) = &r.Alt0 { fields.insert("Alt0".to_owned(), value.clone()); }
                if let Some(value) = &r.Alternative1 { fields.insert("Alternative1".to_owned(), value.clone()); }
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply1 { fields.insert("Apply1".to_owned(), value.clone()); }
                if let Some(value) = &r.Biapply0 { fields.insert("Biapply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifoldable1 { fields.insert("Bifoldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bind1 { fields.insert("Bind1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bounded0 { fields.insert("Bounded0".to_owned(), value.clone()); }
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
                if let Some(value) = &r.CommutativeRing0 { fields.insert("CommutativeRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ComonadAsk0 { fields.insert("ComonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.Contravariant0 { fields.insert("Contravariant0".to_owned(), value.clone()); }
                if let Some(value) = &r.Decide0 { fields.insert("Decide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divisible1 { fields.insert("Divisible1".to_owned(), value.clone()); }
                if let Some(value) = &r.DivisionRing1 { fields.insert("DivisionRing1".to_owned(), value.clone()); }
                if let Some(value) = &r.Enum1 { fields.insert("Enum1".to_owned(), value.clone()); }
                if let Some(value) = &r.Eq0 { fields.insert("Eq0".to_owned(), value.clone()); }
                if let Some(value) = &r.Eq10 { fields.insert("Eq10".to_owned(), value.clone()); }
                if let Some(value) = &r.EqRecord0 { fields.insert("EqRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.EuclideanRing0 { fields.insert("EuclideanRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.Extend0 { fields.insert("Extend0".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable1 { fields.insert("Foldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable10 { fields.insert("Foldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.FoldableWithIndex1 { fields.insert("FoldableWithIndex1".to_owned(), value.clone()); }
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.FunctorWithIndex0 { fields.insert("FunctorWithIndex0".to_owned(), value.clone()); }
                if let Some(value) = &r.HeytingAlgebra0 { fields.insert("HeytingAlgebra0".to_owned(), value.clone()); }
                if let Some(value) = &r.HeytingAlgebraRecord0 { fields.insert("HeytingAlgebraRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad1 { fields.insert("Monad1".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadAsk0 { fields.insert("MonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadEffect0 { fields.insert("MonadEffect0".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadTell1 { fields.insert("MonadTell1".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadThrow0 { fields.insert("MonadThrow0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monoid0 { fields.insert("Monoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.OrdRecord0 { fields.insert("OrdRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Plus1 { fields.insert("Plus1".to_owned(), value.clone()); }
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
                if let Some(value) = &r.RingRecord0 { fields.insert("RingRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.SemigroupRecord0 { fields.insert("SemigroupRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroupoid0 { fields.insert("Semigroupoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semiring0 { fields.insert("Semiring0".to_owned(), value.clone()); }
                if let Some(value) = &r.SemiringRecord0 { fields.insert("SemiringRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable1 { fields.insert("Traversable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable2 { fields.insert("Traversable2".to_owned(), value.clone()); }
                if let Some(value) = &r.Unfoldable10 { fields.insert("Unfoldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.accum { fields.insert("accum".to_owned(), value.clone()); }
                if let Some(value) = &r.add { fields.insert("add".to_owned(), value.clone()); }
                if let Some(value) = &r.addRecord { fields.insert("addRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.after { fields.insert("after".to_owned(), value.clone()); }
                if let Some(value) = &r.alt { fields.insert("alt".to_owned(), value.clone()); }
                if let Some(value) = &r.append { fields.insert("append".to_owned(), value.clone()); }
                if let Some(value) = &r.appendRecord { fields.insert("appendRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.apply { fields.insert("apply".to_owned(), value.clone()); }
                if let Some(value) = &r.asList { fields.insert("asList".to_owned(), value.clone()); }
                if let Some(value) = &r.asMap { fields.insert("asMap".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.before { fields.insert("before".to_owned(), value.clone()); }
                if let Some(value) = &r.biapply { fields.insert("biapply".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldMap { fields.insert("bifoldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldl { fields.insert("bifoldl".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldr { fields.insert("bifoldr".to_owned(), value.clone()); }
                if let Some(value) = &r.bimap { fields.insert("bimap".to_owned(), value.clone()); }
                if let Some(value) = &r.bind { fields.insert("bind".to_owned(), value.clone()); }
                if let Some(value) = &r.bipure { fields.insert("bipure".to_owned(), value.clone()); }
                if let Some(value) = &r.bisequence { fields.insert("bisequence".to_owned(), value.clone()); }
                if let Some(value) = &r.bitraverse { fields.insert("bitraverse".to_owned(), value.clone()); }
                if let Some(value) = &r.bottom { fields.insert("bottom".to_owned(), value.clone()); }
                if let Some(value) = &r.bottomRecord { fields.insert("bottomRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.callCC { fields.insert("callCC".to_owned(), value.clone()); }
                if let Some(value) = &r.cardinality { fields.insert("cardinality".to_owned(), value.clone()); }
                if let Some(value) = &r.catchError { fields.insert("catchError".to_owned(), value.clone()); }
                if let Some(value) = &r.choose { fields.insert("choose".to_owned(), value.clone()); }
                if let Some(value) = &r.closed { fields.insert("closed".to_owned(), value.clone()); }
                if let Some(value) = &r.cmap { fields.insert("cmap".to_owned(), value.clone()); }
                if let Some(value) = &r.collect { fields.insert("collect".to_owned(), value.clone()); }
                if let Some(value) = &r.compare { fields.insert("compare".to_owned(), value.clone()); }
                if let Some(value) = &r.compare1 { fields.insert("compare1".to_owned(), value.clone()); }
                if let Some(value) = &r.compareRecord { fields.insert("compareRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.completed { fields.insert("completed".to_owned(), value.clone()); }
                if let Some(value) = &r.compose { fields.insert("compose".to_owned(), value.clone()); }
                if let Some(value) = &r.conj { fields.insert("conj".to_owned(), value.clone()); }
                if let Some(value) = &r.conjRecord { fields.insert("conjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.conquer { fields.insert("conquer".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.day { fields.insert("day".to_owned(), value.clone()); }
                if let Some(value) = &r.defer { fields.insert("defer".to_owned(), value.clone()); }
                if let Some(value) = &r.degree { fields.insert("degree".to_owned(), value.clone()); }
                if let Some(value) = &r.dimap { fields.insert("dimap".to_owned(), value.clone()); }
                if let Some(value) = &r.discard { fields.insert("discard".to_owned(), value.clone()); }
                if let Some(value) = &r.disj { fields.insert("disj".to_owned(), value.clone()); }
                if let Some(value) = &r.disjRecord { fields.insert("disjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.distribute { fields.insert("distribute".to_owned(), value.clone()); }
                if let Some(value) = &r.div { fields.insert("div".to_owned(), value.clone()); }
                if let Some(value) = &r.divide { fields.insert("divide".to_owned(), value.clone()); }
                if let Some(value) = &r.dotAll { fields.insert("dotAll".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.elem { fields.insert("elem".to_owned(), value.clone()); }
                if let Some(value) = &r.empty { fields.insert("empty".to_owned(), value.clone()); }
                if let Some(value) = &r.eq { fields.insert("eq".to_owned(), value.clone()); }
                if let Some(value) = &r.eq1 { fields.insert("eq1".to_owned(), value.clone()); }
                if let Some(value) = &r.eqRecord { fields.insert("eqRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.extend { fields.insert("extend".to_owned(), value.clone()); }
                if let Some(value) = &r.extract { fields.insert("extract".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
                if let Some(value) = &r.failed { fields.insert("failed".to_owned(), value.clone()); }
                if let Some(value) = &r.ff { fields.insert("ff".to_owned(), value.clone()); }
                if let Some(value) = &r.ffRecord { fields.insert("ffRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.fiber { fields.insert("fiber".to_owned(), value.clone()); }
                if let Some(value) = &r.first { fields.insert("first".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap { fields.insert("foldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap1 { fields.insert("foldMap1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMapWithIndex { fields.insert("foldMapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl { fields.insert("foldl".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl1 { fields.insert("foldl1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldlWithIndex { fields.insert("foldlWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr { fields.insert("foldr".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr1 { fields.insert("foldr1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldrWithIndex { fields.insert("foldrWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.found { fields.insert("found".to_owned(), value.clone()); }
                if let Some(value) = &r.from { fields.insert("from".to_owned(), value.clone()); }
                if let Some(value) = &r.fromDuration { fields.insert("fromDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.fromEnum { fields.insert("fromEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.fromLeft { fields.insert("fromLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.fromRight { fields.insert("fromRight".to_owned(), value.clone()); }
                if let Some(value) = &r.genericAdd_prime { fields.insert("genericAdd'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericAppend_prime { fields.insert("genericAppend'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericBottom_prime { fields.insert("genericBottom'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericCardinality_prime { fields.insert("genericCardinality'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericCompare_prime { fields.insert("genericCompare'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericConj_prime { fields.insert("genericConj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericDisj_prime { fields.insert("genericDisj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericEq_prime { fields.insert("genericEq'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFF_prime { fields.insert("genericFF'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFromEnum_prime { fields.insert("genericFromEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericImplies_prime { fields.insert("genericImplies'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMempty_prime { fields.insert("genericMempty'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMul_prime { fields.insert("genericMul'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericNot_prime { fields.insert("genericNot'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericOne_prime { fields.insert("genericOne'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericPred_prime { fields.insert("genericPred'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericShow_prime { fields.insert("genericShow'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericShowArgs { fields.insert("genericShowArgs".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSub_prime { fields.insert("genericSub'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSucc_prime { fields.insert("genericSucc'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTT_prime { fields.insert("genericTT'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericToEnum_prime { fields.insert("genericToEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTop_prime { fields.insert("genericTop'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericZero_prime { fields.insert("genericZero'".to_owned(), value.clone()); }
                if let Some(value) = &r.global { fields.insert("global".to_owned(), value.clone()); }
                if let Some(value) = &r.handler { fields.insert("handler".to_owned(), value.clone()); }
                if let Some(value) = &r.head { fields.insert("head".to_owned(), value.clone()); }
                if let Some(value) = &r.hour { fields.insert("hour".to_owned(), value.clone()); }
                if let Some(value) = &r.identity { fields.insert("identity".to_owned(), value.clone()); }
                if let Some(value) = &r.ignoreCase { fields.insert("ignoreCase".to_owned(), value.clone()); }
                if let Some(value) = &r.imap { fields.insert("imap".to_owned(), value.clone()); }
                if let Some(value) = &r.implies { fields.insert("implies".to_owned(), value.clone()); }
                if let Some(value) = &r.impliesRecord { fields.insert("impliesRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.index { fields.insert("index".to_owned(), value.clone()); }
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.inj { fields.insert("inj".to_owned(), value.clone()); }
                if let Some(value) = &r.isLeft { fields.insert("isLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.isSuspended { fields.insert("isSuspended".to_owned(), value.clone()); }
                if let Some(value) = &r.join { fields.insert("join".to_owned(), value.clone()); }
                if let Some(value) = &r.key { fields.insert("key".to_owned(), value.clone()); }
                if let Some(value) = &r.keysImpl { fields.insert("keysImpl".to_owned(), value.clone()); }
                if let Some(value) = &r.kill { fields.insert("kill".to_owned(), value.clone()); }
                if let Some(value) = &r.killed { fields.insert("killed".to_owned(), value.clone()); }
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.lift { fields.insert("lift".to_owned(), value.clone()); }
                if let Some(value) = &r.liftAff { fields.insert("liftAff".to_owned(), value.clone()); }
                if let Some(value) = &r.liftEffect { fields.insert("liftEffect".to_owned(), value.clone()); }
                if let Some(value) = &r.liftST { fields.insert("liftST".to_owned(), value.clone()); }
                if let Some(value) = &r.listen { fields.insert("listen".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
                if let Some(value) = &r.lose { fields.insert("lose".to_owned(), value.clone()); }
                if let Some(value) = &r.lower { fields.insert("lower".to_owned(), value.clone()); }
                if let Some(value) = &r.map { fields.insert("map".to_owned(), value.clone()); }
                if let Some(value) = &r.mapWithIndex { fields.insert("mapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.mappend_ { fields.insert("mappend_".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty { fields.insert("mempty".to_owned(), value.clone()); }
                if let Some(value) = &r.memptyRecord { fields.insert("memptyRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty_ { fields.insert("mempty_".to_owned(), value.clone()); }
                if let Some(value) = &r.millisecond { fields.insert("millisecond".to_owned(), value.clone()); }
                if let Some(value) = &r.minute { fields.insert("minute".to_owned(), value.clone()); }
                if let Some(value) = &r.mod_kw { fields.insert("mod".to_owned(), value.clone()); }
                if let Some(value) = &r.month { fields.insert("month".to_owned(), value.clone()); }
                if let Some(value) = &r.mul { fields.insert("mul".to_owned(), value.clone()); }
                if let Some(value) = &r.mulRecord { fields.insert("mulRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.multiline { fields.insert("multiline".to_owned(), value.clone()); }
                if let Some(value) = &r.myMethod { fields.insert("myMethod".to_owned(), value.clone()); }
                if let Some(value) = &r.nes { fields.insert("nes".to_owned(), value.clone()); }
                if let Some(value) = &r.no { fields.insert("no".to_owned(), value.clone()); }
                if let Some(value) = &r.not { fields.insert("not".to_owned(), value.clone()); }
                if let Some(value) = &r.notRecord { fields.insert("notRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.onComplete { fields.insert("onComplete".to_owned(), value.clone()); }
                if let Some(value) = &r.one { fields.insert("one".to_owned(), value.clone()); }
                if let Some(value) = &r.oneRecord { fields.insert("oneRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.parallel { fields.insert("parallel".to_owned(), value.clone()); }
                if let Some(value) = &r.pass { fields.insert("pass".to_owned(), value.clone()); }
                if let Some(value) = &r.peek { fields.insert("peek".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
                if let Some(value) = &r.pred { fields.insert("pred".to_owned(), value.clone()); }
                if let Some(value) = &r.prj { fields.insert("prj".to_owned(), value.clone()); }
                if let Some(value) = &r.proof { fields.insert("proof".to_owned(), value.clone()); }
                if let Some(value) = &r.ps_minus_rust_minus_test { fields.insert("ps-rust-test".to_owned(), value.clone()); }
                if let Some(value) = &r.pure { fields.insert("pure".to_owned(), value.clone()); }
                if let Some(value) = &r.recip { fields.insert("recip".to_owned(), value.clone()); }
                if let Some(value) = &r.reflectSymbol { fields.insert("reflectSymbol".to_owned(), value.clone()); }
                if let Some(value) = &r.reflectType { fields.insert("reflectType".to_owned(), value.clone()); }
                if let Some(value) = &r.rest { fields.insert("rest".to_owned(), value.clone()); }
                if let Some(value) = &r.result { fields.insert("result".to_owned(), value.clone()); }
                if let Some(value) = &r.rethrow { fields.insert("rethrow".to_owned(), value.clone()); }
                if let Some(value) = &r.revInit { fields.insert("revInit".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
                if let Some(value) = &r.run { fields.insert("run".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence { fields.insert("sequence".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence1 { fields.insert("sequence1".to_owned(), value.clone()); }
                if let Some(value) = &r.sequential { fields.insert("sequential".to_owned(), value.clone()); }
                if let Some(value) = &r.show { fields.insert("show".to_owned(), value.clone()); }
                if let Some(value) = &r.showRecordFields { fields.insert("showRecordFields".to_owned(), value.clone()); }
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.sticky { fields.insert("sticky".to_owned(), value.clone()); }
                if let Some(value) = &r.sub { fields.insert("sub".to_owned(), value.clone()); }
                if let Some(value) = &r.subRecord { fields.insert("subRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.succ { fields.insert("succ".to_owned(), value.clone()); }
                if let Some(value) = &r.supervisor { fields.insert("supervisor".to_owned(), value.clone()); }
                if let Some(value) = &r.tail { fields.insert("tail".to_owned(), value.clone()); }
                if let Some(value) = &r.tailRecM { fields.insert("tailRecM".to_owned(), value.clone()); }
                if let Some(value) = &r.tell { fields.insert("tell".to_owned(), value.clone()); }
                if let Some(value) = &r.throwError { fields.insert("throwError".to_owned(), value.clone()); }
                if let Some(value) = &r.to { fields.insert("to".to_owned(), value.clone()); }
                if let Some(value) = &r.toDuration { fields.insert("toDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.toEnum { fields.insert("toEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.top { fields.insert("top".to_owned(), value.clone()); }
                if let Some(value) = &r.topRecord { fields.insert("topRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.track { fields.insert("track".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse { fields.insert("traverse".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse1 { fields.insert("traverse1".to_owned(), value.clone()); }
                if let Some(value) = &r.traverseWithIndex { fields.insert("traverseWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.tt { fields.insert("tt".to_owned(), value.clone()); }
                if let Some(value) = &r.ttRecord { fields.insert("ttRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr { fields.insert("unfoldr".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr1 { fields.insert("unfoldr1".to_owned(), value.clone()); }
                if let Some(value) = &r.unicode { fields.insert("unicode".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
                if let Some(value) = &r.year { fields.insert("year".to_owned(), value.clone()); }
                if let Some(value) = &r.yes { fields.insert("yes".to_owned(), value.clone()); }
                if let Some(value) = &r.zero { fields.insert("zero".to_owned(), value.clone()); }
                if let Some(value) = &r.zeroRecord { fields.insert("zeroRecord".to_owned(), value.clone()); }
            },
            _ => unreachable!(),
        }
        fields.insert(name.to_owned(), value);
        Value::DynamicRecord(perceus_ptr::PerceusPtr::new(fields))
    }
    pub fn __purust_record_fields(&self) -> Option<RecordFields> {
        let mut fields = RecordFields::new();
        match self.resolve() {
            Value::Record_Alt0_empty(r) => {
                if let Some(value) = &r.Alt0 { fields.insert("Alt0".to_owned(), value.clone()); }
                if let Some(value) = &r.empty { fields.insert("empty".to_owned(), value.clone()); }
            },
            Value::Record_Alternative1_Monad0(r) => {
                if let Some(value) = &r.Alternative1 { fields.insert("Alternative1".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
            },
            Value::Record_Applicative0_Bind1(r) => {
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bind1 { fields.insert("Bind1".to_owned(), value.clone()); }
            },
            Value::Record_Applicative0_Plus1(r) => {
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Plus1 { fields.insert("Plus1".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply1 { fields.insert("Apply1".to_owned(), value.clone()); }
                if let Some(value) = &r.parallel { fields.insert("parallel".to_owned(), value.clone()); }
                if let Some(value) = &r.sequential { fields.insert("sequential".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_bind(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.bind { fields.insert("bind".to_owned(), value.clone()); }
            },
            Value::Record_Apply0_pure(r) => {
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.pure { fields.insert("pure".to_owned(), value.clone()); }
            },
            Value::Record_Biapply0_bipure(r) => {
                if let Some(value) = &r.Biapply0 { fields.insert("Biapply0".to_owned(), value.clone()); }
                if let Some(value) = &r.bipure { fields.insert("bipure".to_owned(), value.clone()); }
            },
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                if let Some(value) = &r.Bifoldable1 { fields.insert("Bifoldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.bisequence { fields.insert("bisequence".to_owned(), value.clone()); }
                if let Some(value) = &r.bitraverse { fields.insert("bitraverse".to_owned(), value.clone()); }
            },
            Value::Record_Bifunctor0_biapply(r) => {
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.biapply { fields.insert("biapply".to_owned(), value.clone()); }
            },
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                if let Some(value) = &r.Bounded0 { fields.insert("Bounded0".to_owned(), value.clone()); }
                if let Some(value) = &r.Enum1 { fields.insert("Enum1".to_owned(), value.clone()); }
                if let Some(value) = &r.cardinality { fields.insert("cardinality".to_owned(), value.clone()); }
                if let Some(value) = &r.fromEnum { fields.insert("fromEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.toEnum { fields.insert("toEnum".to_owned(), value.clone()); }
            },
            Value::Record_Coercible0(r) => {
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
            },
            Value::Record_Coercible0_proof(r) => {
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
                if let Some(value) = &r.proof { fields.insert("proof".to_owned(), value.clone()); }
            },
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                if let Some(value) = &r.CommutativeRing0 { fields.insert("CommutativeRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.degree { fields.insert("degree".to_owned(), value.clone()); }
                if let Some(value) = &r.div { fields.insert("div".to_owned(), value.clone()); }
                if let Some(value) = &r.mod_kw { fields.insert("mod".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_ask(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_peek_pos(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.peek { fields.insert("peek".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
            },
            Value::Record_Comonad0_track(r) => {
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.track { fields.insert("track".to_owned(), value.clone()); }
            },
            Value::Record_ComonadAsk0_local(r) => {
                if let Some(value) = &r.ComonadAsk0 { fields.insert("ComonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
            },
            Value::Record_Contravariant0_divide(r) => {
                if let Some(value) = &r.Contravariant0 { fields.insert("Contravariant0".to_owned(), value.clone()); }
                if let Some(value) = &r.divide { fields.insert("divide".to_owned(), value.clone()); }
            },
            Value::Record_Decide0_Divisible1_lose(r) => {
                if let Some(value) = &r.Decide0 { fields.insert("Decide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divisible1 { fields.insert("Divisible1".to_owned(), value.clone()); }
                if let Some(value) = &r.lose { fields.insert("lose".to_owned(), value.clone()); }
            },
            Value::Record_Divide0_choose(r) => {
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.choose { fields.insert("choose".to_owned(), value.clone()); }
            },
            Value::Record_Divide0_conquer(r) => {
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.conquer { fields.insert("conquer".to_owned(), value.clone()); }
            },
            Value::Record_DivisionRing1_EuclideanRing0(r) => {
                if let Some(value) = &r.DivisionRing1 { fields.insert("DivisionRing1".to_owned(), value.clone()); }
                if let Some(value) = &r.EuclideanRing0 { fields.insert("EuclideanRing0".to_owned(), value.clone()); }
            },
            Value::Record_Eq0_compare(r) => {
                if let Some(value) = &r.Eq0 { fields.insert("Eq0".to_owned(), value.clone()); }
                if let Some(value) = &r.compare { fields.insert("compare".to_owned(), value.clone()); }
            },
            Value::Record_Eq10_compare1(r) => {
                if let Some(value) = &r.Eq10 { fields.insert("Eq10".to_owned(), value.clone()); }
                if let Some(value) = &r.compare1 { fields.insert("compare1".to_owned(), value.clone()); }
            },
            Value::Record_EqRecord0_compareRecord(r) => {
                if let Some(value) = &r.EqRecord0 { fields.insert("EqRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.compareRecord { fields.insert("compareRecord".to_owned(), value.clone()); }
            },
            Value::Record_Extend0_extract(r) => {
                if let Some(value) = &r.Extend0 { fields.insert("Extend0".to_owned(), value.clone()); }
                if let Some(value) = &r.extract { fields.insert("extract".to_owned(), value.clone()); }
            },
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap1 { fields.insert("foldMap1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl1 { fields.insert("foldl1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr1 { fields.insert("foldr1".to_owned(), value.clone()); }
            },
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMapWithIndex { fields.insert("foldMapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldlWithIndex { fields.insert("foldlWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldrWithIndex { fields.insert("foldrWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                if let Some(value) = &r.Foldable1 { fields.insert("Foldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence { fields.insert("sequence".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse { fields.insert("traverse".to_owned(), value.clone()); }
            },
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                if let Some(value) = &r.Foldable10 { fields.insert("Foldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable1 { fields.insert("Traversable1".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence1 { fields.insert("sequence1".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse1 { fields.insert("traverse1".to_owned(), value.clone()); }
            },
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                if let Some(value) = &r.FoldableWithIndex1 { fields.insert("FoldableWithIndex1".to_owned(), value.clone()); }
                if let Some(value) = &r.FunctorWithIndex0 { fields.insert("FunctorWithIndex0".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable2 { fields.insert("Traversable2".to_owned(), value.clone()); }
                if let Some(value) = &r.traverseWithIndex { fields.insert("traverseWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_alt(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.alt { fields.insert("alt".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_apply(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.apply { fields.insert("apply".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_collect_distribute(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.collect { fields.insert("collect".to_owned(), value.clone()); }
                if let Some(value) = &r.distribute { fields.insert("distribute".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_extend(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.extend { fields.insert("extend".to_owned(), value.clone()); }
            },
            Value::Record_Functor0_mapWithIndex(r) => {
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.mapWithIndex { fields.insert("mapWithIndex".to_owned(), value.clone()); }
            },
            Value::Record_HeytingAlgebra0(r) => {
                if let Some(value) = &r.HeytingAlgebra0 { fields.insert("HeytingAlgebra0".to_owned(), value.clone()); }
            },
            Value::Record_HeytingAlgebraRecord0(r) => {
                if let Some(value) = &r.HeytingAlgebraRecord0 { fields.insert("HeytingAlgebraRecord0".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_ask(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_callCC(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.callCC { fields.insert("callCC".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_liftEffect(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftEffect { fields.insert("liftEffect".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_liftST(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftST { fields.insert("liftST".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_state(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_tailRecM(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.tailRecM { fields.insert("tailRecM".to_owned(), value.clone()); }
            },
            Value::Record_Monad0_throwError(r) => {
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.throwError { fields.insert("throwError".to_owned(), value.clone()); }
            },
            Value::Record_Monad1_Semigroup0_tell(r) => {
                if let Some(value) = &r.Monad1 { fields.insert("Monad1".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.tell { fields.insert("tell".to_owned(), value.clone()); }
            },
            Value::Record_MonadAsk0_local(r) => {
                if let Some(value) = &r.MonadAsk0 { fields.insert("MonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
            },
            Value::Record_MonadEffect0_liftAff(r) => {
                if let Some(value) = &r.MonadEffect0 { fields.insert("MonadEffect0".to_owned(), value.clone()); }
                if let Some(value) = &r.liftAff { fields.insert("liftAff".to_owned(), value.clone()); }
            },
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                if let Some(value) = &r.MonadTell1 { fields.insert("MonadTell1".to_owned(), value.clone()); }
                if let Some(value) = &r.Monoid0 { fields.insert("Monoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.listen { fields.insert("listen".to_owned(), value.clone()); }
                if let Some(value) = &r.pass { fields.insert("pass".to_owned(), value.clone()); }
            },
            Value::Record_MonadThrow0_catchError(r) => {
                if let Some(value) = &r.MonadThrow0 { fields.insert("MonadThrow0".to_owned(), value.clone()); }
                if let Some(value) = &r.catchError { fields.insert("catchError".to_owned(), value.clone()); }
            },
            Value::Record_Ord0_bottom_top(r) => {
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.bottom { fields.insert("bottom".to_owned(), value.clone()); }
                if let Some(value) = &r.top { fields.insert("top".to_owned(), value.clone()); }
            },
            Value::Record_Ord0_pred_succ(r) => {
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.pred { fields.insert("pred".to_owned(), value.clone()); }
                if let Some(value) = &r.succ { fields.insert("succ".to_owned(), value.clone()); }
            },
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => {
                if let Some(value) = &r.OrdRecord0 { fields.insert("OrdRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.bottomRecord { fields.insert("bottomRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.topRecord { fields.insert("topRecord".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_closed(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.closed { fields.insert("closed".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_first_second(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.first { fields.insert("first".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
            },
            Value::Record_Profunctor0_left_right(r) => {
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
            },
            Value::Record_Ring0(r) => {
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
            },
            Value::Record_Ring0_recip(r) => {
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
                if let Some(value) = &r.recip { fields.insert("recip".to_owned(), value.clone()); }
            },
            Value::Record_RingRecord0(r) => {
                if let Some(value) = &r.RingRecord0 { fields.insert("RingRecord0".to_owned(), value.clone()); }
            },
            Value::Record_Semigroup0_mempty(r) => {
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty { fields.insert("mempty".to_owned(), value.clone()); }
            },
            Value::Record_SemigroupRecord0_memptyRecord(r) => {
                if let Some(value) = &r.SemigroupRecord0 { fields.insert("SemigroupRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.memptyRecord { fields.insert("memptyRecord".to_owned(), value.clone()); }
            },
            Value::Record_Semigroupoid0_identity(r) => {
                if let Some(value) = &r.Semigroupoid0 { fields.insert("Semigroupoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.identity { fields.insert("identity".to_owned(), value.clone()); }
            },
            Value::Record_Semiring0_sub(r) => {
                if let Some(value) = &r.Semiring0 { fields.insert("Semiring0".to_owned(), value.clone()); }
                if let Some(value) = &r.sub { fields.insert("sub".to_owned(), value.clone()); }
            },
            Value::Record_SemiringRecord0_subRecord(r) => {
                if let Some(value) = &r.SemiringRecord0 { fields.insert("SemiringRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.subRecord { fields.insert("subRecord".to_owned(), value.clone()); }
            },
            Value::Record_Unfoldable10_unfoldr(r) => {
                if let Some(value) = &r.Unfoldable10 { fields.insert("Unfoldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr { fields.insert("unfoldr".to_owned(), value.clone()); }
            },
            Value::Record_a_b(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
            },
            Value::Record_a_b_c(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
            },
            Value::Record_a_b_c_d_e(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
            },
            Value::Record_acc_init(r) => {
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
            },
            Value::Record_acc_val(r) => {
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
            },
            Value::Record_accum_value(r) => {
                if let Some(value) = &r.accum { fields.insert("accum".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_add_mul_one_zero(r) => {
                if let Some(value) = &r.add { fields.insert("add".to_owned(), value.clone()); }
                if let Some(value) = &r.mul { fields.insert("mul".to_owned(), value.clone()); }
                if let Some(value) = &r.one { fields.insert("one".to_owned(), value.clone()); }
                if let Some(value) = &r.zero { fields.insert("zero".to_owned(), value.clone()); }
            },
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                if let Some(value) = &r.addRecord { fields.insert("addRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.mulRecord { fields.insert("mulRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.oneRecord { fields.insert("oneRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.zeroRecord { fields.insert("zeroRecord".to_owned(), value.clone()); }
            },
            Value::Record_after_before(r) => {
                if let Some(value) = &r.after { fields.insert("after".to_owned(), value.clone()); }
                if let Some(value) = &r.before { fields.insert("before".to_owned(), value.clone()); }
            },
            Value::Record_append(r) => {
                if let Some(value) = &r.append { fields.insert("append".to_owned(), value.clone()); }
            },
            Value::Record_appendRecord(r) => {
                if let Some(value) = &r.appendRecord { fields.insert("appendRecord".to_owned(), value.clone()); }
            },
            Value::Record_asList_asMap(r) => {
                if let Some(value) = &r.asList { fields.insert("asList".to_owned(), value.clone()); }
                if let Some(value) = &r.asMap { fields.insert("asMap".to_owned(), value.clone()); }
            },
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => {
                if let Some(value) = &r.bifoldMap { fields.insert("bifoldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldl { fields.insert("bifoldl".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldr { fields.insert("bifoldr".to_owned(), value.clone()); }
            },
            Value::Record_bimap(r) => {
                if let Some(value) = &r.bimap { fields.insert("bimap".to_owned(), value.clone()); }
            },
            Value::Record_c_d(r) => {
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
            },
            Value::Record_cmap(r) => {
                if let Some(value) = &r.cmap { fields.insert("cmap".to_owned(), value.clone()); }
            },
            Value::Record_completed_failed_killed(r) => {
                if let Some(value) = &r.completed { fields.insert("completed".to_owned(), value.clone()); }
                if let Some(value) = &r.failed { fields.insert("failed".to_owned(), value.clone()); }
                if let Some(value) = &r.killed { fields.insert("killed".to_owned(), value.clone()); }
            },
            Value::Record_compose(r) => {
                if let Some(value) = &r.compose { fields.insert("compose".to_owned(), value.clone()); }
            },
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                if let Some(value) = &r.conj { fields.insert("conj".to_owned(), value.clone()); }
                if let Some(value) = &r.disj { fields.insert("disj".to_owned(), value.clone()); }
                if let Some(value) = &r.ff { fields.insert("ff".to_owned(), value.clone()); }
                if let Some(value) = &r.implies { fields.insert("implies".to_owned(), value.clone()); }
                if let Some(value) = &r.not { fields.insert("not".to_owned(), value.clone()); }
                if let Some(value) = &r.tt { fields.insert("tt".to_owned(), value.clone()); }
            },
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                if let Some(value) = &r.conjRecord { fields.insert("conjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.disjRecord { fields.insert("disjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.ffRecord { fields.insert("ffRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.impliesRecord { fields.insert("impliesRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.notRecord { fields.insert("notRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.ttRecord { fields.insert("ttRecord".to_owned(), value.clone()); }
            },
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                if let Some(value) = &r.day { fields.insert("day".to_owned(), value.clone()); }
                if let Some(value) = &r.hour { fields.insert("hour".to_owned(), value.clone()); }
                if let Some(value) = &r.millisecond { fields.insert("millisecond".to_owned(), value.clone()); }
                if let Some(value) = &r.minute { fields.insert("minute".to_owned(), value.clone()); }
                if let Some(value) = &r.month { fields.insert("month".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
                if let Some(value) = &r.year { fields.insert("year".to_owned(), value.clone()); }
            },
            Value::Record_defer(r) => {
                if let Some(value) = &r.defer { fields.insert("defer".to_owned(), value.clone()); }
            },
            Value::Record_dimap(r) => {
                if let Some(value) = &r.dimap { fields.insert("dimap".to_owned(), value.clone()); }
            },
            Value::Record_discard(r) => {
                if let Some(value) = &r.discard { fields.insert("discard".to_owned(), value.clone()); }
            },
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                if let Some(value) = &r.dotAll { fields.insert("dotAll".to_owned(), value.clone()); }
                if let Some(value) = &r.global { fields.insert("global".to_owned(), value.clone()); }
                if let Some(value) = &r.ignoreCase { fields.insert("ignoreCase".to_owned(), value.clone()); }
                if let Some(value) = &r.multiline { fields.insert("multiline".to_owned(), value.clone()); }
                if let Some(value) = &r.sticky { fields.insert("sticky".to_owned(), value.clone()); }
                if let Some(value) = &r.unicode { fields.insert("unicode".to_owned(), value.clone()); }
            },
            Value::Record_e_f(r) => {
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
            },
            Value::Record_elem_pos(r) => {
                if let Some(value) = &r.elem { fields.insert("elem".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
            },
            Value::Record_eq(r) => {
                if let Some(value) = &r.eq { fields.insert("eq".to_owned(), value.clone()); }
            },
            Value::Record_eq1(r) => {
                if let Some(value) = &r.eq1 { fields.insert("eq1".to_owned(), value.clone()); }
            },
            Value::Record_eqRecord(r) => {
                if let Some(value) = &r.eqRecord { fields.insert("eqRecord".to_owned(), value.clone()); }
            },
            Value::Record_fiber_supervisor(r) => {
                if let Some(value) = &r.fiber { fields.insert("fiber".to_owned(), value.clone()); }
                if let Some(value) = &r.supervisor { fields.insert("supervisor".to_owned(), value.clone()); }
            },
            Value::Record_foldMap_foldl_foldr(r) => {
                if let Some(value) = &r.foldMap { fields.insert("foldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl { fields.insert("foldl".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr { fields.insert("foldr".to_owned(), value.clone()); }
            },
            Value::Record_found_result(r) => {
                if let Some(value) = &r.found { fields.insert("found".to_owned(), value.clone()); }
                if let Some(value) = &r.result { fields.insert("result".to_owned(), value.clone()); }
            },
            Value::Record_from_to(r) => {
                if let Some(value) = &r.from { fields.insert("from".to_owned(), value.clone()); }
                if let Some(value) = &r.to { fields.insert("to".to_owned(), value.clone()); }
            },
            Value::Record_fromDuration_toDuration(r) => {
                if let Some(value) = &r.fromDuration { fields.insert("fromDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.toDuration { fields.insert("toDuration".to_owned(), value.clone()); }
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                if let Some(value) = &r.fromLeft { fields.insert("fromLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.fromRight { fields.insert("fromRight".to_owned(), value.clone()); }
                if let Some(value) = &r.isLeft { fields.insert("isLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
            },
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                if let Some(value) = &r.genericAdd_prime { fields.insert("genericAdd'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMul_prime { fields.insert("genericMul'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericOne_prime { fields.insert("genericOne'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericZero_prime { fields.insert("genericZero'".to_owned(), value.clone()); }
            },
            Value::Record_genericAppend_prime(r) => {
                if let Some(value) = &r.genericAppend_prime { fields.insert("genericAppend'".to_owned(), value.clone()); }
            },
            Value::Record_genericBottom_prime(r) => {
                if let Some(value) = &r.genericBottom_prime { fields.insert("genericBottom'".to_owned(), value.clone()); }
            },
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => {
                if let Some(value) = &r.genericCardinality_prime { fields.insert("genericCardinality'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFromEnum_prime { fields.insert("genericFromEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericToEnum_prime { fields.insert("genericToEnum'".to_owned(), value.clone()); }
            },
            Value::Record_genericCompare_prime(r) => {
                if let Some(value) = &r.genericCompare_prime { fields.insert("genericCompare'".to_owned(), value.clone()); }
            },
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                if let Some(value) = &r.genericConj_prime { fields.insert("genericConj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericDisj_prime { fields.insert("genericDisj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFF_prime { fields.insert("genericFF'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericImplies_prime { fields.insert("genericImplies'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericNot_prime { fields.insert("genericNot'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTT_prime { fields.insert("genericTT'".to_owned(), value.clone()); }
            },
            Value::Record_genericEq_prime(r) => {
                if let Some(value) = &r.genericEq_prime { fields.insert("genericEq'".to_owned(), value.clone()); }
            },
            Value::Record_genericMempty_prime(r) => {
                if let Some(value) = &r.genericMempty_prime { fields.insert("genericMempty'".to_owned(), value.clone()); }
            },
            Value::Record_genericPred_prime_genericSucc_prime(r) => {
                if let Some(value) = &r.genericPred_prime { fields.insert("genericPred'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSucc_prime { fields.insert("genericSucc'".to_owned(), value.clone()); }
            },
            Value::Record_genericShow_prime(r) => {
                if let Some(value) = &r.genericShow_prime { fields.insert("genericShow'".to_owned(), value.clone()); }
            },
            Value::Record_genericShowArgs(r) => {
                if let Some(value) = &r.genericShowArgs { fields.insert("genericShowArgs".to_owned(), value.clone()); }
            },
            Value::Record_genericSub_prime(r) => {
                if let Some(value) = &r.genericSub_prime { fields.insert("genericSub'".to_owned(), value.clone()); }
            },
            Value::Record_genericTop_prime(r) => {
                if let Some(value) = &r.genericTop_prime { fields.insert("genericTop'".to_owned(), value.clone()); }
            },
            Value::Record_handler_rethrow(r) => {
                if let Some(value) = &r.handler { fields.insert("handler".to_owned(), value.clone()); }
                if let Some(value) = &r.rethrow { fields.insert("rethrow".to_owned(), value.clone()); }
            },
            Value::Record_head_tail(r) => {
                if let Some(value) = &r.head { fields.insert("head".to_owned(), value.clone()); }
                if let Some(value) = &r.tail { fields.insert("tail".to_owned(), value.clone()); }
            },
            Value::Record_imap(r) => {
                if let Some(value) = &r.imap { fields.insert("imap".to_owned(), value.clone()); }
            },
            Value::Record_index_value(r) => {
                if let Some(value) = &r.index { fields.insert("index".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_init_last(r) => {
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
            },
            Value::Record_init_rest(r) => {
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.rest { fields.insert("rest".to_owned(), value.clone()); }
            },
            Value::Record_inj_prj(r) => {
                if let Some(value) = &r.inj { fields.insert("inj".to_owned(), value.clone()); }
                if let Some(value) = &r.prj { fields.insert("prj".to_owned(), value.clone()); }
            },
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                if let Some(value) = &r.isSuspended { fields.insert("isSuspended".to_owned(), value.clone()); }
                if let Some(value) = &r.join { fields.insert("join".to_owned(), value.clone()); }
                if let Some(value) = &r.kill { fields.insert("kill".to_owned(), value.clone()); }
                if let Some(value) = &r.onComplete { fields.insert("onComplete".to_owned(), value.clone()); }
                if let Some(value) = &r.run { fields.insert("run".to_owned(), value.clone()); }
            },
            Value::Record_key_value(r) => {
                if let Some(value) = &r.key { fields.insert("key".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_keysImpl(r) => {
                if let Some(value) = &r.keysImpl { fields.insert("keysImpl".to_owned(), value.clone()); }
            },
            Value::Record_last_revInit(r) => {
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
                if let Some(value) = &r.revInit { fields.insert("revInit".to_owned(), value.clone()); }
            },
            Value::Record_lift(r) => {
                if let Some(value) = &r.lift { fields.insert("lift".to_owned(), value.clone()); }
            },
            Value::Record_lower(r) => {
                if let Some(value) = &r.lower { fields.insert("lower".to_owned(), value.clone()); }
            },
            Value::Record_map(r) => {
                if let Some(value) = &r.map { fields.insert("map".to_owned(), value.clone()); }
            },
            Value::Record_mappend__mempty_(r) => {
                if let Some(value) = &r.mappend_ { fields.insert("mappend_".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty_ { fields.insert("mempty_".to_owned(), value.clone()); }
            },
            Value::Record_myMethod(r) => {
                if let Some(value) = &r.myMethod { fields.insert("myMethod".to_owned(), value.clone()); }
            },
            Value::Record_nes(r) => {
                if let Some(value) = &r.nes { fields.insert("nes".to_owned(), value.clone()); }
            },
            Value::Record_no_yes(r) => {
                if let Some(value) = &r.no { fields.insert("no".to_owned(), value.clone()); }
                if let Some(value) = &r.yes { fields.insert("yes".to_owned(), value.clone()); }
            },
            Value::Record_ps_minus_rust_minus_test(r) => {
                if let Some(value) = &r.ps_minus_rust_minus_test { fields.insert("ps-rust-test".to_owned(), value.clone()); }
            },
            Value::Record_reflectSymbol(r) => {
                if let Some(value) = &r.reflectSymbol { fields.insert("reflectSymbol".to_owned(), value.clone()); }
            },
            Value::Record_reflectType(r) => {
                if let Some(value) = &r.reflectType { fields.insert("reflectType".to_owned(), value.clone()); }
            },
            Value::Record_show(r) => {
                if let Some(value) = &r.show { fields.insert("show".to_owned(), value.clone()); }
            },
            Value::Record_showRecordFields(r) => {
                if let Some(value) = &r.showRecordFields { fields.insert("showRecordFields".to_owned(), value.clone()); }
            },
            Value::Record_state_val(r) => {
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
            },
            Value::Record_state_value(r) => {
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
            },
            Value::Record_unfoldr1(r) => {
                if let Some(value) = &r.unfoldr1 { fields.insert("unfoldr1".to_owned(), value.clone()); }
            },
            Value::Record_a(r) => {
                if let Some(value) = &r.Alt0 { fields.insert("Alt0".to_owned(), value.clone()); }
                if let Some(value) = &r.Alternative1 { fields.insert("Alternative1".to_owned(), value.clone()); }
                if let Some(value) = &r.Applicative0 { fields.insert("Applicative0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply0 { fields.insert("Apply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Apply1 { fields.insert("Apply1".to_owned(), value.clone()); }
                if let Some(value) = &r.Biapply0 { fields.insert("Biapply0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifoldable1 { fields.insert("Bifoldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bifunctor0 { fields.insert("Bifunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.Bind1 { fields.insert("Bind1".to_owned(), value.clone()); }
                if let Some(value) = &r.Bounded0 { fields.insert("Bounded0".to_owned(), value.clone()); }
                if let Some(value) = &r.Coercible0 { fields.insert("Coercible0".to_owned(), value.clone()); }
                if let Some(value) = &r.CommutativeRing0 { fields.insert("CommutativeRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.Comonad0 { fields.insert("Comonad0".to_owned(), value.clone()); }
                if let Some(value) = &r.ComonadAsk0 { fields.insert("ComonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.Contravariant0 { fields.insert("Contravariant0".to_owned(), value.clone()); }
                if let Some(value) = &r.Decide0 { fields.insert("Decide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divide0 { fields.insert("Divide0".to_owned(), value.clone()); }
                if let Some(value) = &r.Divisible1 { fields.insert("Divisible1".to_owned(), value.clone()); }
                if let Some(value) = &r.DivisionRing1 { fields.insert("DivisionRing1".to_owned(), value.clone()); }
                if let Some(value) = &r.Enum1 { fields.insert("Enum1".to_owned(), value.clone()); }
                if let Some(value) = &r.Eq0 { fields.insert("Eq0".to_owned(), value.clone()); }
                if let Some(value) = &r.Eq10 { fields.insert("Eq10".to_owned(), value.clone()); }
                if let Some(value) = &r.EqRecord0 { fields.insert("EqRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.EuclideanRing0 { fields.insert("EuclideanRing0".to_owned(), value.clone()); }
                if let Some(value) = &r.Extend0 { fields.insert("Extend0".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable0 { fields.insert("Foldable0".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable1 { fields.insert("Foldable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Foldable10 { fields.insert("Foldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.FoldableWithIndex1 { fields.insert("FoldableWithIndex1".to_owned(), value.clone()); }
                if let Some(value) = &r.Functor0 { fields.insert("Functor0".to_owned(), value.clone()); }
                if let Some(value) = &r.FunctorWithIndex0 { fields.insert("FunctorWithIndex0".to_owned(), value.clone()); }
                if let Some(value) = &r.HeytingAlgebra0 { fields.insert("HeytingAlgebra0".to_owned(), value.clone()); }
                if let Some(value) = &r.HeytingAlgebraRecord0 { fields.insert("HeytingAlgebraRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad0 { fields.insert("Monad0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monad1 { fields.insert("Monad1".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadAsk0 { fields.insert("MonadAsk0".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadEffect0 { fields.insert("MonadEffect0".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadTell1 { fields.insert("MonadTell1".to_owned(), value.clone()); }
                if let Some(value) = &r.MonadThrow0 { fields.insert("MonadThrow0".to_owned(), value.clone()); }
                if let Some(value) = &r.Monoid0 { fields.insert("Monoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.Ord0 { fields.insert("Ord0".to_owned(), value.clone()); }
                if let Some(value) = &r.OrdRecord0 { fields.insert("OrdRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Plus1 { fields.insert("Plus1".to_owned(), value.clone()); }
                if let Some(value) = &r.Profunctor0 { fields.insert("Profunctor0".to_owned(), value.clone()); }
                if let Some(value) = &r.Ring0 { fields.insert("Ring0".to_owned(), value.clone()); }
                if let Some(value) = &r.RingRecord0 { fields.insert("RingRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroup0 { fields.insert("Semigroup0".to_owned(), value.clone()); }
                if let Some(value) = &r.SemigroupRecord0 { fields.insert("SemigroupRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semigroupoid0 { fields.insert("Semigroupoid0".to_owned(), value.clone()); }
                if let Some(value) = &r.Semiring0 { fields.insert("Semiring0".to_owned(), value.clone()); }
                if let Some(value) = &r.SemiringRecord0 { fields.insert("SemiringRecord0".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable1 { fields.insert("Traversable1".to_owned(), value.clone()); }
                if let Some(value) = &r.Traversable2 { fields.insert("Traversable2".to_owned(), value.clone()); }
                if let Some(value) = &r.Unfoldable10 { fields.insert("Unfoldable10".to_owned(), value.clone()); }
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.acc { fields.insert("acc".to_owned(), value.clone()); }
                if let Some(value) = &r.accum { fields.insert("accum".to_owned(), value.clone()); }
                if let Some(value) = &r.add { fields.insert("add".to_owned(), value.clone()); }
                if let Some(value) = &r.addRecord { fields.insert("addRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.after { fields.insert("after".to_owned(), value.clone()); }
                if let Some(value) = &r.alt { fields.insert("alt".to_owned(), value.clone()); }
                if let Some(value) = &r.append { fields.insert("append".to_owned(), value.clone()); }
                if let Some(value) = &r.appendRecord { fields.insert("appendRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.apply { fields.insert("apply".to_owned(), value.clone()); }
                if let Some(value) = &r.asList { fields.insert("asList".to_owned(), value.clone()); }
                if let Some(value) = &r.asMap { fields.insert("asMap".to_owned(), value.clone()); }
                if let Some(value) = &r.ask { fields.insert("ask".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.before { fields.insert("before".to_owned(), value.clone()); }
                if let Some(value) = &r.biapply { fields.insert("biapply".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldMap { fields.insert("bifoldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldl { fields.insert("bifoldl".to_owned(), value.clone()); }
                if let Some(value) = &r.bifoldr { fields.insert("bifoldr".to_owned(), value.clone()); }
                if let Some(value) = &r.bimap { fields.insert("bimap".to_owned(), value.clone()); }
                if let Some(value) = &r.bind { fields.insert("bind".to_owned(), value.clone()); }
                if let Some(value) = &r.bipure { fields.insert("bipure".to_owned(), value.clone()); }
                if let Some(value) = &r.bisequence { fields.insert("bisequence".to_owned(), value.clone()); }
                if let Some(value) = &r.bitraverse { fields.insert("bitraverse".to_owned(), value.clone()); }
                if let Some(value) = &r.bottom { fields.insert("bottom".to_owned(), value.clone()); }
                if let Some(value) = &r.bottomRecord { fields.insert("bottomRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.callCC { fields.insert("callCC".to_owned(), value.clone()); }
                if let Some(value) = &r.cardinality { fields.insert("cardinality".to_owned(), value.clone()); }
                if let Some(value) = &r.catchError { fields.insert("catchError".to_owned(), value.clone()); }
                if let Some(value) = &r.choose { fields.insert("choose".to_owned(), value.clone()); }
                if let Some(value) = &r.closed { fields.insert("closed".to_owned(), value.clone()); }
                if let Some(value) = &r.cmap { fields.insert("cmap".to_owned(), value.clone()); }
                if let Some(value) = &r.collect { fields.insert("collect".to_owned(), value.clone()); }
                if let Some(value) = &r.compare { fields.insert("compare".to_owned(), value.clone()); }
                if let Some(value) = &r.compare1 { fields.insert("compare1".to_owned(), value.clone()); }
                if let Some(value) = &r.compareRecord { fields.insert("compareRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.completed { fields.insert("completed".to_owned(), value.clone()); }
                if let Some(value) = &r.compose { fields.insert("compose".to_owned(), value.clone()); }
                if let Some(value) = &r.conj { fields.insert("conj".to_owned(), value.clone()); }
                if let Some(value) = &r.conjRecord { fields.insert("conjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.conquer { fields.insert("conquer".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.day { fields.insert("day".to_owned(), value.clone()); }
                if let Some(value) = &r.defer { fields.insert("defer".to_owned(), value.clone()); }
                if let Some(value) = &r.degree { fields.insert("degree".to_owned(), value.clone()); }
                if let Some(value) = &r.dimap { fields.insert("dimap".to_owned(), value.clone()); }
                if let Some(value) = &r.discard { fields.insert("discard".to_owned(), value.clone()); }
                if let Some(value) = &r.disj { fields.insert("disj".to_owned(), value.clone()); }
                if let Some(value) = &r.disjRecord { fields.insert("disjRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.distribute { fields.insert("distribute".to_owned(), value.clone()); }
                if let Some(value) = &r.div { fields.insert("div".to_owned(), value.clone()); }
                if let Some(value) = &r.divide { fields.insert("divide".to_owned(), value.clone()); }
                if let Some(value) = &r.dotAll { fields.insert("dotAll".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.elem { fields.insert("elem".to_owned(), value.clone()); }
                if let Some(value) = &r.empty { fields.insert("empty".to_owned(), value.clone()); }
                if let Some(value) = &r.eq { fields.insert("eq".to_owned(), value.clone()); }
                if let Some(value) = &r.eq1 { fields.insert("eq1".to_owned(), value.clone()); }
                if let Some(value) = &r.eqRecord { fields.insert("eqRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.extend { fields.insert("extend".to_owned(), value.clone()); }
                if let Some(value) = &r.extract { fields.insert("extract".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
                if let Some(value) = &r.failed { fields.insert("failed".to_owned(), value.clone()); }
                if let Some(value) = &r.ff { fields.insert("ff".to_owned(), value.clone()); }
                if let Some(value) = &r.ffRecord { fields.insert("ffRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.fiber { fields.insert("fiber".to_owned(), value.clone()); }
                if let Some(value) = &r.first { fields.insert("first".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap { fields.insert("foldMap".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMap1 { fields.insert("foldMap1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldMapWithIndex { fields.insert("foldMapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl { fields.insert("foldl".to_owned(), value.clone()); }
                if let Some(value) = &r.foldl1 { fields.insert("foldl1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldlWithIndex { fields.insert("foldlWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr { fields.insert("foldr".to_owned(), value.clone()); }
                if let Some(value) = &r.foldr1 { fields.insert("foldr1".to_owned(), value.clone()); }
                if let Some(value) = &r.foldrWithIndex { fields.insert("foldrWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.found { fields.insert("found".to_owned(), value.clone()); }
                if let Some(value) = &r.from { fields.insert("from".to_owned(), value.clone()); }
                if let Some(value) = &r.fromDuration { fields.insert("fromDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.fromEnum { fields.insert("fromEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.fromLeft { fields.insert("fromLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.fromRight { fields.insert("fromRight".to_owned(), value.clone()); }
                if let Some(value) = &r.genericAdd_prime { fields.insert("genericAdd'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericAppend_prime { fields.insert("genericAppend'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericBottom_prime { fields.insert("genericBottom'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericCardinality_prime { fields.insert("genericCardinality'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericCompare_prime { fields.insert("genericCompare'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericConj_prime { fields.insert("genericConj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericDisj_prime { fields.insert("genericDisj'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericEq_prime { fields.insert("genericEq'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFF_prime { fields.insert("genericFF'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericFromEnum_prime { fields.insert("genericFromEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericImplies_prime { fields.insert("genericImplies'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMempty_prime { fields.insert("genericMempty'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericMul_prime { fields.insert("genericMul'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericNot_prime { fields.insert("genericNot'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericOne_prime { fields.insert("genericOne'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericPred_prime { fields.insert("genericPred'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericShow_prime { fields.insert("genericShow'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericShowArgs { fields.insert("genericShowArgs".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSub_prime { fields.insert("genericSub'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericSucc_prime { fields.insert("genericSucc'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTT_prime { fields.insert("genericTT'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericToEnum_prime { fields.insert("genericToEnum'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericTop_prime { fields.insert("genericTop'".to_owned(), value.clone()); }
                if let Some(value) = &r.genericZero_prime { fields.insert("genericZero'".to_owned(), value.clone()); }
                if let Some(value) = &r.global { fields.insert("global".to_owned(), value.clone()); }
                if let Some(value) = &r.handler { fields.insert("handler".to_owned(), value.clone()); }
                if let Some(value) = &r.head { fields.insert("head".to_owned(), value.clone()); }
                if let Some(value) = &r.hour { fields.insert("hour".to_owned(), value.clone()); }
                if let Some(value) = &r.identity { fields.insert("identity".to_owned(), value.clone()); }
                if let Some(value) = &r.ignoreCase { fields.insert("ignoreCase".to_owned(), value.clone()); }
                if let Some(value) = &r.imap { fields.insert("imap".to_owned(), value.clone()); }
                if let Some(value) = &r.implies { fields.insert("implies".to_owned(), value.clone()); }
                if let Some(value) = &r.impliesRecord { fields.insert("impliesRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.index { fields.insert("index".to_owned(), value.clone()); }
                if let Some(value) = &r.init { fields.insert("init".to_owned(), value.clone()); }
                if let Some(value) = &r.inj { fields.insert("inj".to_owned(), value.clone()); }
                if let Some(value) = &r.isLeft { fields.insert("isLeft".to_owned(), value.clone()); }
                if let Some(value) = &r.isSuspended { fields.insert("isSuspended".to_owned(), value.clone()); }
                if let Some(value) = &r.join { fields.insert("join".to_owned(), value.clone()); }
                if let Some(value) = &r.key { fields.insert("key".to_owned(), value.clone()); }
                if let Some(value) = &r.keysImpl { fields.insert("keysImpl".to_owned(), value.clone()); }
                if let Some(value) = &r.kill { fields.insert("kill".to_owned(), value.clone()); }
                if let Some(value) = &r.killed { fields.insert("killed".to_owned(), value.clone()); }
                if let Some(value) = &r.last { fields.insert("last".to_owned(), value.clone()); }
                if let Some(value) = &r.left { fields.insert("left".to_owned(), value.clone()); }
                if let Some(value) = &r.lift { fields.insert("lift".to_owned(), value.clone()); }
                if let Some(value) = &r.liftAff { fields.insert("liftAff".to_owned(), value.clone()); }
                if let Some(value) = &r.liftEffect { fields.insert("liftEffect".to_owned(), value.clone()); }
                if let Some(value) = &r.liftST { fields.insert("liftST".to_owned(), value.clone()); }
                if let Some(value) = &r.listen { fields.insert("listen".to_owned(), value.clone()); }
                if let Some(value) = &r.local { fields.insert("local".to_owned(), value.clone()); }
                if let Some(value) = &r.lose { fields.insert("lose".to_owned(), value.clone()); }
                if let Some(value) = &r.lower { fields.insert("lower".to_owned(), value.clone()); }
                if let Some(value) = &r.map { fields.insert("map".to_owned(), value.clone()); }
                if let Some(value) = &r.mapWithIndex { fields.insert("mapWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.mappend_ { fields.insert("mappend_".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty { fields.insert("mempty".to_owned(), value.clone()); }
                if let Some(value) = &r.memptyRecord { fields.insert("memptyRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.mempty_ { fields.insert("mempty_".to_owned(), value.clone()); }
                if let Some(value) = &r.millisecond { fields.insert("millisecond".to_owned(), value.clone()); }
                if let Some(value) = &r.minute { fields.insert("minute".to_owned(), value.clone()); }
                if let Some(value) = &r.mod_kw { fields.insert("mod".to_owned(), value.clone()); }
                if let Some(value) = &r.month { fields.insert("month".to_owned(), value.clone()); }
                if let Some(value) = &r.mul { fields.insert("mul".to_owned(), value.clone()); }
                if let Some(value) = &r.mulRecord { fields.insert("mulRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.multiline { fields.insert("multiline".to_owned(), value.clone()); }
                if let Some(value) = &r.myMethod { fields.insert("myMethod".to_owned(), value.clone()); }
                if let Some(value) = &r.nes { fields.insert("nes".to_owned(), value.clone()); }
                if let Some(value) = &r.no { fields.insert("no".to_owned(), value.clone()); }
                if let Some(value) = &r.not { fields.insert("not".to_owned(), value.clone()); }
                if let Some(value) = &r.notRecord { fields.insert("notRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.onComplete { fields.insert("onComplete".to_owned(), value.clone()); }
                if let Some(value) = &r.one { fields.insert("one".to_owned(), value.clone()); }
                if let Some(value) = &r.oneRecord { fields.insert("oneRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.parallel { fields.insert("parallel".to_owned(), value.clone()); }
                if let Some(value) = &r.pass { fields.insert("pass".to_owned(), value.clone()); }
                if let Some(value) = &r.peek { fields.insert("peek".to_owned(), value.clone()); }
                if let Some(value) = &r.pos { fields.insert("pos".to_owned(), value.clone()); }
                if let Some(value) = &r.pred { fields.insert("pred".to_owned(), value.clone()); }
                if let Some(value) = &r.prj { fields.insert("prj".to_owned(), value.clone()); }
                if let Some(value) = &r.proof { fields.insert("proof".to_owned(), value.clone()); }
                if let Some(value) = &r.ps_minus_rust_minus_test { fields.insert("ps-rust-test".to_owned(), value.clone()); }
                if let Some(value) = &r.pure { fields.insert("pure".to_owned(), value.clone()); }
                if let Some(value) = &r.recip { fields.insert("recip".to_owned(), value.clone()); }
                if let Some(value) = &r.reflectSymbol { fields.insert("reflectSymbol".to_owned(), value.clone()); }
                if let Some(value) = &r.reflectType { fields.insert("reflectType".to_owned(), value.clone()); }
                if let Some(value) = &r.rest { fields.insert("rest".to_owned(), value.clone()); }
                if let Some(value) = &r.result { fields.insert("result".to_owned(), value.clone()); }
                if let Some(value) = &r.rethrow { fields.insert("rethrow".to_owned(), value.clone()); }
                if let Some(value) = &r.revInit { fields.insert("revInit".to_owned(), value.clone()); }
                if let Some(value) = &r.right { fields.insert("right".to_owned(), value.clone()); }
                if let Some(value) = &r.run { fields.insert("run".to_owned(), value.clone()); }
                if let Some(value) = &r.second { fields.insert("second".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence { fields.insert("sequence".to_owned(), value.clone()); }
                if let Some(value) = &r.sequence1 { fields.insert("sequence1".to_owned(), value.clone()); }
                if let Some(value) = &r.sequential { fields.insert("sequential".to_owned(), value.clone()); }
                if let Some(value) = &r.show { fields.insert("show".to_owned(), value.clone()); }
                if let Some(value) = &r.showRecordFields { fields.insert("showRecordFields".to_owned(), value.clone()); }
                if let Some(value) = &r.state { fields.insert("state".to_owned(), value.clone()); }
                if let Some(value) = &r.sticky { fields.insert("sticky".to_owned(), value.clone()); }
                if let Some(value) = &r.sub { fields.insert("sub".to_owned(), value.clone()); }
                if let Some(value) = &r.subRecord { fields.insert("subRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.succ { fields.insert("succ".to_owned(), value.clone()); }
                if let Some(value) = &r.supervisor { fields.insert("supervisor".to_owned(), value.clone()); }
                if let Some(value) = &r.tail { fields.insert("tail".to_owned(), value.clone()); }
                if let Some(value) = &r.tailRecM { fields.insert("tailRecM".to_owned(), value.clone()); }
                if let Some(value) = &r.tell { fields.insert("tell".to_owned(), value.clone()); }
                if let Some(value) = &r.throwError { fields.insert("throwError".to_owned(), value.clone()); }
                if let Some(value) = &r.to { fields.insert("to".to_owned(), value.clone()); }
                if let Some(value) = &r.toDuration { fields.insert("toDuration".to_owned(), value.clone()); }
                if let Some(value) = &r.toEnum { fields.insert("toEnum".to_owned(), value.clone()); }
                if let Some(value) = &r.top { fields.insert("top".to_owned(), value.clone()); }
                if let Some(value) = &r.topRecord { fields.insert("topRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.track { fields.insert("track".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse { fields.insert("traverse".to_owned(), value.clone()); }
                if let Some(value) = &r.traverse1 { fields.insert("traverse1".to_owned(), value.clone()); }
                if let Some(value) = &r.traverseWithIndex { fields.insert("traverseWithIndex".to_owned(), value.clone()); }
                if let Some(value) = &r.tt { fields.insert("tt".to_owned(), value.clone()); }
                if let Some(value) = &r.ttRecord { fields.insert("ttRecord".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr { fields.insert("unfoldr".to_owned(), value.clone()); }
                if let Some(value) = &r.unfoldr1 { fields.insert("unfoldr1".to_owned(), value.clone()); }
                if let Some(value) = &r.unicode { fields.insert("unicode".to_owned(), value.clone()); }
                if let Some(value) = &r.val { fields.insert("val".to_owned(), value.clone()); }
                if let Some(value) = &r.value { fields.insert("value".to_owned(), value.clone()); }
                if let Some(value) = &r.year { fields.insert("year".to_owned(), value.clone()); }
                if let Some(value) = &r.yes { fields.insert("yes".to_owned(), value.clone()); }
                if let Some(value) = &r.zero { fields.insert("zero".to_owned(), value.clone()); }
                if let Some(value) = &r.zeroRecord { fields.insert("zeroRecord".to_owned(), value.clone()); }
            },
            Value::DynamicRecord(r) => return Some((**r).clone()),
            _ => return None,
        }
        Some(fields)
    }
    pub fn __purust_borrow_Alt0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Alt0_empty(r) => r.Alt0.as_ref().unwrap(),
            Value::Record_a(r) => r.Alt0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Alt0").expect("Missing record field"),
            _ => panic!("Expected record with field Alt0"),
        }
    }
    pub fn __purust_borrow_Alternative1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Alternative1_Monad0(r) => r.Alternative1.as_ref().unwrap(),
            Value::Record_a(r) => r.Alternative1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Alternative1").expect("Missing record field"),
            _ => panic!("Expected record with field Alternative1"),
        }
    }
    pub fn __purust_borrow_Applicative0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Bind1(r) => r.Applicative0.as_ref().unwrap(),
            Value::Record_Applicative0_Plus1(r) => r.Applicative0.as_ref().unwrap(),
            Value::Record_a(r) => r.Applicative0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Applicative0").expect("Missing record field"),
            _ => panic!("Expected record with field Applicative0"),
        }
    }
    pub fn __purust_borrow_Apply0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.Apply0.as_ref().unwrap(),
            Value::Record_Apply0_bind(r) => r.Apply0.as_ref().unwrap(),
            Value::Record_Apply0_pure(r) => r.Apply0.as_ref().unwrap(),
            Value::Record_a(r) => r.Apply0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Apply0").expect("Missing record field"),
            _ => panic!("Expected record with field Apply0"),
        }
    }
    pub fn __purust_borrow_Apply1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.Apply1.as_ref().unwrap(),
            Value::Record_a(r) => r.Apply1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Apply1").expect("Missing record field"),
            _ => panic!("Expected record with field Apply1"),
        }
    }
    pub fn __purust_borrow_Biapply0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Biapply0_bipure(r) => r.Biapply0.as_ref().unwrap(),
            Value::Record_a(r) => r.Biapply0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Biapply0").expect("Missing record field"),
            _ => panic!("Expected record with field Biapply0"),
        }
    }
    pub fn __purust_borrow_Bifoldable1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.Bifoldable1.as_ref().unwrap(),
            Value::Record_a(r) => r.Bifoldable1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Bifoldable1").expect("Missing record field"),
            _ => panic!("Expected record with field Bifoldable1"),
        }
    }
    pub fn __purust_borrow_Bifunctor0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.Bifunctor0.as_ref().unwrap(),
            Value::Record_Bifunctor0_biapply(r) => r.Bifunctor0.as_ref().unwrap(),
            Value::Record_a(r) => r.Bifunctor0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Bifunctor0").expect("Missing record field"),
            _ => panic!("Expected record with field Bifunctor0"),
        }
    }
    pub fn __purust_borrow_Bind1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Bind1(r) => r.Bind1.as_ref().unwrap(),
            Value::Record_a(r) => r.Bind1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Bind1").expect("Missing record field"),
            _ => panic!("Expected record with field Bind1"),
        }
    }
    pub fn __purust_borrow_Bounded0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.Bounded0.as_ref().unwrap(),
            Value::Record_a(r) => r.Bounded0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Bounded0").expect("Missing record field"),
            _ => panic!("Expected record with field Bounded0"),
        }
    }
    pub fn __purust_borrow_Coercible0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Coercible0(r) => r.Coercible0.as_ref().unwrap(),
            Value::Record_Coercible0_proof(r) => r.Coercible0.as_ref().unwrap(),
            Value::Record_a(r) => r.Coercible0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Coercible0").expect("Missing record field"),
            _ => panic!("Expected record with field Coercible0"),
        }
    }
    pub fn __purust_borrow_CommutativeRing0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.CommutativeRing0.as_ref().unwrap(),
            Value::Record_a(r) => r.CommutativeRing0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("CommutativeRing0").expect("Missing record field"),
            _ => panic!("Expected record with field CommutativeRing0"),
        }
    }
    pub fn __purust_borrow_Comonad0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_ask(r) => r.Comonad0.as_ref().unwrap(),
            Value::Record_Comonad0_peek_pos(r) => r.Comonad0.as_ref().unwrap(),
            Value::Record_Comonad0_track(r) => r.Comonad0.as_ref().unwrap(),
            Value::Record_a(r) => r.Comonad0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Comonad0").expect("Missing record field"),
            _ => panic!("Expected record with field Comonad0"),
        }
    }
    pub fn __purust_borrow_ComonadAsk0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_ComonadAsk0_local(r) => r.ComonadAsk0.as_ref().unwrap(),
            Value::Record_a(r) => r.ComonadAsk0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ComonadAsk0").expect("Missing record field"),
            _ => panic!("Expected record with field ComonadAsk0"),
        }
    }
    pub fn __purust_borrow_Contravariant0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Contravariant0_divide(r) => r.Contravariant0.as_ref().unwrap(),
            Value::Record_a(r) => r.Contravariant0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Contravariant0").expect("Missing record field"),
            _ => panic!("Expected record with field Contravariant0"),
        }
    }
    pub fn __purust_borrow_Decide0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.Decide0.as_ref().unwrap(),
            Value::Record_a(r) => r.Decide0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Decide0").expect("Missing record field"),
            _ => panic!("Expected record with field Decide0"),
        }
    }
    pub fn __purust_borrow_Divide0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Divide0_choose(r) => r.Divide0.as_ref().unwrap(),
            Value::Record_Divide0_conquer(r) => r.Divide0.as_ref().unwrap(),
            Value::Record_a(r) => r.Divide0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Divide0").expect("Missing record field"),
            _ => panic!("Expected record with field Divide0"),
        }
    }
    pub fn __purust_borrow_Divisible1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.Divisible1.as_ref().unwrap(),
            Value::Record_a(r) => r.Divisible1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Divisible1").expect("Missing record field"),
            _ => panic!("Expected record with field Divisible1"),
        }
    }
    pub fn __purust_borrow_DivisionRing1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_DivisionRing1_EuclideanRing0(r) => r.DivisionRing1.as_ref().unwrap(),
            Value::Record_a(r) => r.DivisionRing1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("DivisionRing1").expect("Missing record field"),
            _ => panic!("Expected record with field DivisionRing1"),
        }
    }
    pub fn __purust_borrow_Enum1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.Enum1.as_ref().unwrap(),
            Value::Record_a(r) => r.Enum1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Enum1").expect("Missing record field"),
            _ => panic!("Expected record with field Enum1"),
        }
    }
    pub fn __purust_borrow_Eq0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Eq0_compare(r) => r.Eq0.as_ref().unwrap(),
            Value::Record_a(r) => r.Eq0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Eq0").expect("Missing record field"),
            _ => panic!("Expected record with field Eq0"),
        }
    }
    pub fn __purust_borrow_Eq10(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Eq10_compare1(r) => r.Eq10.as_ref().unwrap(),
            Value::Record_a(r) => r.Eq10.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Eq10").expect("Missing record field"),
            _ => panic!("Expected record with field Eq10"),
        }
    }
    pub fn __purust_borrow_EqRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_EqRecord0_compareRecord(r) => r.EqRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.EqRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("EqRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field EqRecord0"),
        }
    }
    pub fn __purust_borrow_EuclideanRing0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_DivisionRing1_EuclideanRing0(r) => r.EuclideanRing0.as_ref().unwrap(),
            Value::Record_a(r) => r.EuclideanRing0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("EuclideanRing0").expect("Missing record field"),
            _ => panic!("Expected record with field EuclideanRing0"),
        }
    }
    pub fn __purust_borrow_Extend0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Extend0_extract(r) => r.Extend0.as_ref().unwrap(),
            Value::Record_a(r) => r.Extend0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Extend0").expect("Missing record field"),
            _ => panic!("Expected record with field Extend0"),
        }
    }
    pub fn __purust_borrow_Foldable0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.Foldable0.as_ref().unwrap(),
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.Foldable0.as_ref().unwrap(),
            Value::Record_a(r) => r.Foldable0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable0").expect("Missing record field"),
            _ => panic!("Expected record with field Foldable0"),
        }
    }
    pub fn __purust_borrow_Foldable1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.Foldable1.as_ref().unwrap(),
            Value::Record_a(r) => r.Foldable1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable1").expect("Missing record field"),
            _ => panic!("Expected record with field Foldable1"),
        }
    }
    pub fn __purust_borrow_Foldable10(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.Foldable10.as_ref().unwrap(),
            Value::Record_a(r) => r.Foldable10.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Foldable10").expect("Missing record field"),
            _ => panic!("Expected record with field Foldable10"),
        }
    }
    pub fn __purust_borrow_FoldableWithIndex1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.FoldableWithIndex1.as_ref().unwrap(),
            Value::Record_a(r) => r.FoldableWithIndex1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("FoldableWithIndex1").expect("Missing record field"),
            _ => panic!("Expected record with field FoldableWithIndex1"),
        }
    }
    pub fn __purust_borrow_Functor0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_Functor0_alt(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_Functor0_apply(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_Functor0_collect_distribute(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_Functor0_extend(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_Functor0_mapWithIndex(r) => r.Functor0.as_ref().unwrap(),
            Value::Record_a(r) => r.Functor0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Functor0").expect("Missing record field"),
            _ => panic!("Expected record with field Functor0"),
        }
    }
    pub fn __purust_borrow_FunctorWithIndex0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.FunctorWithIndex0.as_ref().unwrap(),
            Value::Record_a(r) => r.FunctorWithIndex0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("FunctorWithIndex0").expect("Missing record field"),
            _ => panic!("Expected record with field FunctorWithIndex0"),
        }
    }
    pub fn __purust_borrow_HeytingAlgebra0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_HeytingAlgebra0(r) => r.HeytingAlgebra0.as_ref().unwrap(),
            Value::Record_a(r) => r.HeytingAlgebra0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("HeytingAlgebra0").expect("Missing record field"),
            _ => panic!("Expected record with field HeytingAlgebra0"),
        }
    }
    pub fn __purust_borrow_HeytingAlgebraRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_HeytingAlgebraRecord0(r) => r.HeytingAlgebraRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.HeytingAlgebraRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("HeytingAlgebraRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field HeytingAlgebraRecord0"),
        }
    }
    pub fn __purust_borrow_Monad0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Alternative1_Monad0(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_ask(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_callCC(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_liftEffect(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_liftST(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_state(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_tailRecM(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_Monad0_throwError(r) => r.Monad0.as_ref().unwrap(),
            Value::Record_a(r) => r.Monad0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Monad0").expect("Missing record field"),
            _ => panic!("Expected record with field Monad0"),
        }
    }
    pub fn __purust_borrow_Monad1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.Monad1.as_ref().unwrap(),
            Value::Record_a(r) => r.Monad1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Monad1").expect("Missing record field"),
            _ => panic!("Expected record with field Monad1"),
        }
    }
    pub fn __purust_borrow_MonadAsk0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadAsk0_local(r) => r.MonadAsk0.as_ref().unwrap(),
            Value::Record_a(r) => r.MonadAsk0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadAsk0").expect("Missing record field"),
            _ => panic!("Expected record with field MonadAsk0"),
        }
    }
    pub fn __purust_borrow_MonadEffect0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadEffect0_liftAff(r) => r.MonadEffect0.as_ref().unwrap(),
            Value::Record_a(r) => r.MonadEffect0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadEffect0").expect("Missing record field"),
            _ => panic!("Expected record with field MonadEffect0"),
        }
    }
    pub fn __purust_borrow_MonadTell1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.MonadTell1.as_ref().unwrap(),
            Value::Record_a(r) => r.MonadTell1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadTell1").expect("Missing record field"),
            _ => panic!("Expected record with field MonadTell1"),
        }
    }
    pub fn __purust_borrow_MonadThrow0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadThrow0_catchError(r) => r.MonadThrow0.as_ref().unwrap(),
            Value::Record_a(r) => r.MonadThrow0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("MonadThrow0").expect("Missing record field"),
            _ => panic!("Expected record with field MonadThrow0"),
        }
    }
    pub fn __purust_borrow_Monoid0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.Monoid0.as_ref().unwrap(),
            Value::Record_a(r) => r.Monoid0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Monoid0").expect("Missing record field"),
            _ => panic!("Expected record with field Monoid0"),
        }
    }
    pub fn __purust_borrow_Ord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.Ord0.as_ref().unwrap(),
            Value::Record_Ord0_pred_succ(r) => r.Ord0.as_ref().unwrap(),
            Value::Record_a(r) => r.Ord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Ord0").expect("Missing record field"),
            _ => panic!("Expected record with field Ord0"),
        }
    }
    pub fn __purust_borrow_OrdRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.OrdRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.OrdRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("OrdRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field OrdRecord0"),
        }
    }
    pub fn __purust_borrow_Plus1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Applicative0_Plus1(r) => r.Plus1.as_ref().unwrap(),
            Value::Record_a(r) => r.Plus1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Plus1").expect("Missing record field"),
            _ => panic!("Expected record with field Plus1"),
        }
    }
    pub fn __purust_borrow_Profunctor0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_closed(r) => r.Profunctor0.as_ref().unwrap(),
            Value::Record_Profunctor0_first_second(r) => r.Profunctor0.as_ref().unwrap(),
            Value::Record_Profunctor0_left_right(r) => r.Profunctor0.as_ref().unwrap(),
            Value::Record_a(r) => r.Profunctor0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Profunctor0").expect("Missing record field"),
            _ => panic!("Expected record with field Profunctor0"),
        }
    }
    pub fn __purust_borrow_Ring0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ring0(r) => r.Ring0.as_ref().unwrap(),
            Value::Record_Ring0_recip(r) => r.Ring0.as_ref().unwrap(),
            Value::Record_a(r) => r.Ring0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Ring0").expect("Missing record field"),
            _ => panic!("Expected record with field Ring0"),
        }
    }
    pub fn __purust_borrow_RingRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_RingRecord0(r) => r.RingRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.RingRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("RingRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field RingRecord0"),
        }
    }
    pub fn __purust_borrow_Semigroup0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.Semigroup0.as_ref().unwrap(),
            Value::Record_Semigroup0_mempty(r) => r.Semigroup0.as_ref().unwrap(),
            Value::Record_a(r) => r.Semigroup0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Semigroup0").expect("Missing record field"),
            _ => panic!("Expected record with field Semigroup0"),
        }
    }
    pub fn __purust_borrow_SemigroupRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_SemigroupRecord0_memptyRecord(r) => r.SemigroupRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.SemigroupRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("SemigroupRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field SemigroupRecord0"),
        }
    }
    pub fn __purust_borrow_Semigroupoid0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Semigroupoid0_identity(r) => r.Semigroupoid0.as_ref().unwrap(),
            Value::Record_a(r) => r.Semigroupoid0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Semigroupoid0").expect("Missing record field"),
            _ => panic!("Expected record with field Semigroupoid0"),
        }
    }
    pub fn __purust_borrow_Semiring0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Semiring0_sub(r) => r.Semiring0.as_ref().unwrap(),
            Value::Record_a(r) => r.Semiring0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Semiring0").expect("Missing record field"),
            _ => panic!("Expected record with field Semiring0"),
        }
    }
    pub fn __purust_borrow_SemiringRecord0(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_SemiringRecord0_subRecord(r) => r.SemiringRecord0.as_ref().unwrap(),
            Value::Record_a(r) => r.SemiringRecord0.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("SemiringRecord0").expect("Missing record field"),
            _ => panic!("Expected record with field SemiringRecord0"),
        }
    }
    pub fn __purust_borrow_Traversable1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.Traversable1.as_ref().unwrap(),
            Value::Record_a(r) => r.Traversable1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Traversable1").expect("Missing record field"),
            _ => panic!("Expected record with field Traversable1"),
        }
    }
    pub fn __purust_borrow_Traversable2(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.Traversable2.as_ref().unwrap(),
            Value::Record_a(r) => r.Traversable2.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Traversable2").expect("Missing record field"),
            _ => panic!("Expected record with field Traversable2"),
        }
    }
    pub fn __purust_borrow_Unfoldable10(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Unfoldable10_unfoldr(r) => r.Unfoldable10.as_ref().unwrap(),
            Value::Record_a(r) => r.Unfoldable10.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("Unfoldable10").expect("Missing record field"),
            _ => panic!("Expected record with field Unfoldable10"),
        }
    }
    pub fn __purust_borrow_a(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b(r) => r.a.as_ref().unwrap(),
            Value::Record_a_b_c(r) => r.a.as_ref().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.a.as_ref().unwrap(),
            Value::Record_a(r) => r.a.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("a").expect("Missing record field"),
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn __purust_borrow_acc(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_acc_init(r) => r.acc.as_ref().unwrap(),
            Value::Record_acc_val(r) => r.acc.as_ref().unwrap(),
            Value::Record_a(r) => r.acc.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("acc").expect("Missing record field"),
            _ => panic!("Expected record with field acc"),
        }
    }
    pub fn __purust_borrow_accum(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_accum_value(r) => r.accum.as_ref().unwrap(),
            Value::Record_a(r) => r.accum.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("accum").expect("Missing record field"),
            _ => panic!("Expected record with field accum"),
        }
    }
    pub fn __purust_borrow_add(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.add.as_ref().unwrap(),
            Value::Record_a(r) => r.add.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("add").expect("Missing record field"),
            _ => panic!("Expected record with field add"),
        }
    }
    pub fn __purust_borrow_addRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.addRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.addRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("addRecord").expect("Missing record field"),
            _ => panic!("Expected record with field addRecord"),
        }
    }
    pub fn __purust_borrow_after(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_after_before(r) => r.after.as_ref().unwrap(),
            Value::Record_a(r) => r.after.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("after").expect("Missing record field"),
            _ => panic!("Expected record with field after"),
        }
    }
    pub fn __purust_borrow_alt(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_alt(r) => r.alt.as_ref().unwrap(),
            Value::Record_a(r) => r.alt.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("alt").expect("Missing record field"),
            _ => panic!("Expected record with field alt"),
        }
    }
    pub fn __purust_borrow_append(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_append(r) => r.append.as_ref().unwrap(),
            Value::Record_a(r) => r.append.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("append").expect("Missing record field"),
            _ => panic!("Expected record with field append"),
        }
    }
    pub fn __purust_borrow_appendRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_appendRecord(r) => r.appendRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.appendRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("appendRecord").expect("Missing record field"),
            _ => panic!("Expected record with field appendRecord"),
        }
    }
    pub fn __purust_borrow_apply(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_apply(r) => r.apply.as_ref().unwrap(),
            Value::Record_a(r) => r.apply.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("apply").expect("Missing record field"),
            _ => panic!("Expected record with field apply"),
        }
    }
    pub fn __purust_borrow_asList(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_asList_asMap(r) => r.asList.as_ref().unwrap(),
            Value::Record_a(r) => r.asList.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("asList").expect("Missing record field"),
            _ => panic!("Expected record with field asList"),
        }
    }
    pub fn __purust_borrow_asMap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_asList_asMap(r) => r.asMap.as_ref().unwrap(),
            Value::Record_a(r) => r.asMap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("asMap").expect("Missing record field"),
            _ => panic!("Expected record with field asMap"),
        }
    }
    pub fn __purust_borrow_ask(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_ask(r) => r.ask.as_ref().unwrap(),
            Value::Record_Monad0_ask(r) => r.ask.as_ref().unwrap(),
            Value::Record_a(r) => r.ask.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ask").expect("Missing record field"),
            _ => panic!("Expected record with field ask"),
        }
    }
    pub fn __purust_borrow_b(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b(r) => r.b.as_ref().unwrap(),
            Value::Record_a_b_c(r) => r.b.as_ref().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.b.as_ref().unwrap(),
            Value::Record_a(r) => r.b.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("b").expect("Missing record field"),
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn __purust_borrow_before(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_after_before(r) => r.before.as_ref().unwrap(),
            Value::Record_a(r) => r.before.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("before").expect("Missing record field"),
            _ => panic!("Expected record with field before"),
        }
    }
    pub fn __purust_borrow_biapply(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bifunctor0_biapply(r) => r.biapply.as_ref().unwrap(),
            Value::Record_a(r) => r.biapply.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("biapply").expect("Missing record field"),
            _ => panic!("Expected record with field biapply"),
        }
    }
    pub fn __purust_borrow_bifoldMap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldMap.as_ref().unwrap(),
            Value::Record_a(r) => r.bifoldMap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldMap").expect("Missing record field"),
            _ => panic!("Expected record with field bifoldMap"),
        }
    }
    pub fn __purust_borrow_bifoldl(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldl.as_ref().unwrap(),
            Value::Record_a(r) => r.bifoldl.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldl").expect("Missing record field"),
            _ => panic!("Expected record with field bifoldl"),
        }
    }
    pub fn __purust_borrow_bifoldr(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => r.bifoldr.as_ref().unwrap(),
            Value::Record_a(r) => r.bifoldr.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bifoldr").expect("Missing record field"),
            _ => panic!("Expected record with field bifoldr"),
        }
    }
    pub fn __purust_borrow_bimap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_bimap(r) => r.bimap.as_ref().unwrap(),
            Value::Record_a(r) => r.bimap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bimap").expect("Missing record field"),
            _ => panic!("Expected record with field bimap"),
        }
    }
    pub fn __purust_borrow_bind(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_bind(r) => r.bind.as_ref().unwrap(),
            Value::Record_a(r) => r.bind.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bind").expect("Missing record field"),
            _ => panic!("Expected record with field bind"),
        }
    }
    pub fn __purust_borrow_bipure(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Biapply0_bipure(r) => r.bipure.as_ref().unwrap(),
            Value::Record_a(r) => r.bipure.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bipure").expect("Missing record field"),
            _ => panic!("Expected record with field bipure"),
        }
    }
    pub fn __purust_borrow_bisequence(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.bisequence.as_ref().unwrap(),
            Value::Record_a(r) => r.bisequence.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bisequence").expect("Missing record field"),
            _ => panic!("Expected record with field bisequence"),
        }
    }
    pub fn __purust_borrow_bitraverse(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => r.bitraverse.as_ref().unwrap(),
            Value::Record_a(r) => r.bitraverse.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bitraverse").expect("Missing record field"),
            _ => panic!("Expected record with field bitraverse"),
        }
    }
    pub fn __purust_borrow_bottom(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.bottom.as_ref().unwrap(),
            Value::Record_a(r) => r.bottom.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bottom").expect("Missing record field"),
            _ => panic!("Expected record with field bottom"),
        }
    }
    pub fn __purust_borrow_bottomRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.bottomRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.bottomRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("bottomRecord").expect("Missing record field"),
            _ => panic!("Expected record with field bottomRecord"),
        }
    }
    pub fn __purust_borrow_c(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_c(r) => r.c.as_ref().unwrap(),
            Value::Record_a_b_c_d_e(r) => r.c.as_ref().unwrap(),
            Value::Record_c_d(r) => r.c.as_ref().unwrap(),
            Value::Record_a(r) => r.c.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("c").expect("Missing record field"),
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn __purust_borrow_callCC(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_callCC(r) => r.callCC.as_ref().unwrap(),
            Value::Record_a(r) => r.callCC.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("callCC").expect("Missing record field"),
            _ => panic!("Expected record with field callCC"),
        }
    }
    pub fn __purust_borrow_cardinality(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.cardinality.as_ref().unwrap(),
            Value::Record_a(r) => r.cardinality.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("cardinality").expect("Missing record field"),
            _ => panic!("Expected record with field cardinality"),
        }
    }
    pub fn __purust_borrow_catchError(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadThrow0_catchError(r) => r.catchError.as_ref().unwrap(),
            Value::Record_a(r) => r.catchError.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("catchError").expect("Missing record field"),
            _ => panic!("Expected record with field catchError"),
        }
    }
    pub fn __purust_borrow_choose(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Divide0_choose(r) => r.choose.as_ref().unwrap(),
            Value::Record_a(r) => r.choose.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("choose").expect("Missing record field"),
            _ => panic!("Expected record with field choose"),
        }
    }
    pub fn __purust_borrow_closed(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_closed(r) => r.closed.as_ref().unwrap(),
            Value::Record_a(r) => r.closed.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("closed").expect("Missing record field"),
            _ => panic!("Expected record with field closed"),
        }
    }
    pub fn __purust_borrow_cmap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_cmap(r) => r.cmap.as_ref().unwrap(),
            Value::Record_a(r) => r.cmap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("cmap").expect("Missing record field"),
            _ => panic!("Expected record with field cmap"),
        }
    }
    pub fn __purust_borrow_collect(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_collect_distribute(r) => r.collect.as_ref().unwrap(),
            Value::Record_a(r) => r.collect.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("collect").expect("Missing record field"),
            _ => panic!("Expected record with field collect"),
        }
    }
    pub fn __purust_borrow_compare(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Eq0_compare(r) => r.compare.as_ref().unwrap(),
            Value::Record_a(r) => r.compare.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("compare").expect("Missing record field"),
            _ => panic!("Expected record with field compare"),
        }
    }
    pub fn __purust_borrow_compare1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Eq10_compare1(r) => r.compare1.as_ref().unwrap(),
            Value::Record_a(r) => r.compare1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("compare1").expect("Missing record field"),
            _ => panic!("Expected record with field compare1"),
        }
    }
    pub fn __purust_borrow_compareRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_EqRecord0_compareRecord(r) => r.compareRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.compareRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("compareRecord").expect("Missing record field"),
            _ => panic!("Expected record with field compareRecord"),
        }
    }
    pub fn __purust_borrow_completed(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.completed.as_ref().unwrap(),
            Value::Record_a(r) => r.completed.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("completed").expect("Missing record field"),
            _ => panic!("Expected record with field completed"),
        }
    }
    pub fn __purust_borrow_compose(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_compose(r) => r.compose.as_ref().unwrap(),
            Value::Record_a(r) => r.compose.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("compose").expect("Missing record field"),
            _ => panic!("Expected record with field compose"),
        }
    }
    pub fn __purust_borrow_conj(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.conj.as_ref().unwrap(),
            Value::Record_a(r) => r.conj.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("conj").expect("Missing record field"),
            _ => panic!("Expected record with field conj"),
        }
    }
    pub fn __purust_borrow_conjRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.conjRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.conjRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("conjRecord").expect("Missing record field"),
            _ => panic!("Expected record with field conjRecord"),
        }
    }
    pub fn __purust_borrow_conquer(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Divide0_conquer(r) => r.conquer.as_ref().unwrap(),
            Value::Record_a(r) => r.conquer.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("conquer").expect("Missing record field"),
            _ => panic!("Expected record with field conquer"),
        }
    }
    pub fn __purust_borrow_d(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_c_d_e(r) => r.d.as_ref().unwrap(),
            Value::Record_c_d(r) => r.d.as_ref().unwrap(),
            Value::Record_a(r) => r.d.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("d").expect("Missing record field"),
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn __purust_borrow_day(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.day.as_ref().unwrap(),
            Value::Record_a(r) => r.day.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("day").expect("Missing record field"),
            _ => panic!("Expected record with field day"),
        }
    }
    pub fn __purust_borrow_defer(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_defer(r) => r.defer.as_ref().unwrap(),
            Value::Record_a(r) => r.defer.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("defer").expect("Missing record field"),
            _ => panic!("Expected record with field defer"),
        }
    }
    pub fn __purust_borrow_degree(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.degree.as_ref().unwrap(),
            Value::Record_a(r) => r.degree.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("degree").expect("Missing record field"),
            _ => panic!("Expected record with field degree"),
        }
    }
    pub fn __purust_borrow_dimap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dimap(r) => r.dimap.as_ref().unwrap(),
            Value::Record_a(r) => r.dimap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("dimap").expect("Missing record field"),
            _ => panic!("Expected record with field dimap"),
        }
    }
    pub fn __purust_borrow_discard(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_discard(r) => r.discard.as_ref().unwrap(),
            Value::Record_a(r) => r.discard.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("discard").expect("Missing record field"),
            _ => panic!("Expected record with field discard"),
        }
    }
    pub fn __purust_borrow_disj(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.disj.as_ref().unwrap(),
            Value::Record_a(r) => r.disj.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("disj").expect("Missing record field"),
            _ => panic!("Expected record with field disj"),
        }
    }
    pub fn __purust_borrow_disjRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.disjRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.disjRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("disjRecord").expect("Missing record field"),
            _ => panic!("Expected record with field disjRecord"),
        }
    }
    pub fn __purust_borrow_distribute(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_collect_distribute(r) => r.distribute.as_ref().unwrap(),
            Value::Record_a(r) => r.distribute.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("distribute").expect("Missing record field"),
            _ => panic!("Expected record with field distribute"),
        }
    }
    pub fn __purust_borrow_div(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.div.as_ref().unwrap(),
            Value::Record_a(r) => r.div.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("div").expect("Missing record field"),
            _ => panic!("Expected record with field div"),
        }
    }
    pub fn __purust_borrow_divide(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Contravariant0_divide(r) => r.divide.as_ref().unwrap(),
            Value::Record_a(r) => r.divide.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("divide").expect("Missing record field"),
            _ => panic!("Expected record with field divide"),
        }
    }
    pub fn __purust_borrow_dotAll(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.dotAll.as_ref().unwrap(),
            Value::Record_a(r) => r.dotAll.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("dotAll").expect("Missing record field"),
            _ => panic!("Expected record with field dotAll"),
        }
    }
    pub fn __purust_borrow_e(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_c_d_e(r) => r.e.as_ref().unwrap(),
            Value::Record_e_f(r) => r.e.as_ref().unwrap(),
            Value::Record_a(r) => r.e.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("e").expect("Missing record field"),
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn __purust_borrow_elem(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_elem_pos(r) => r.elem.as_ref().unwrap(),
            Value::Record_a(r) => r.elem.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("elem").expect("Missing record field"),
            _ => panic!("Expected record with field elem"),
        }
    }
    pub fn __purust_borrow_empty(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Alt0_empty(r) => r.empty.as_ref().unwrap(),
            Value::Record_a(r) => r.empty.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("empty").expect("Missing record field"),
            _ => panic!("Expected record with field empty"),
        }
    }
    pub fn __purust_borrow_eq(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_eq(r) => r.eq.as_ref().unwrap(),
            Value::Record_a(r) => r.eq.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("eq").expect("Missing record field"),
            _ => panic!("Expected record with field eq"),
        }
    }
    pub fn __purust_borrow_eq1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_eq1(r) => r.eq1.as_ref().unwrap(),
            Value::Record_a(r) => r.eq1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("eq1").expect("Missing record field"),
            _ => panic!("Expected record with field eq1"),
        }
    }
    pub fn __purust_borrow_eqRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_eqRecord(r) => r.eqRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.eqRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("eqRecord").expect("Missing record field"),
            _ => panic!("Expected record with field eqRecord"),
        }
    }
    pub fn __purust_borrow_extend(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_extend(r) => r.extend.as_ref().unwrap(),
            Value::Record_a(r) => r.extend.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("extend").expect("Missing record field"),
            _ => panic!("Expected record with field extend"),
        }
    }
    pub fn __purust_borrow_extract(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Extend0_extract(r) => r.extract.as_ref().unwrap(),
            Value::Record_a(r) => r.extract.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("extract").expect("Missing record field"),
            _ => panic!("Expected record with field extract"),
        }
    }
    pub fn __purust_borrow_f(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.f.as_ref().unwrap(),
            Value::Record_a(r) => r.f.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("f").expect("Missing record field"),
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn __purust_borrow_failed(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.failed.as_ref().unwrap(),
            Value::Record_a(r) => r.failed.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("failed").expect("Missing record field"),
            _ => panic!("Expected record with field failed"),
        }
    }
    pub fn __purust_borrow_ff(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.ff.as_ref().unwrap(),
            Value::Record_a(r) => r.ff.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ff").expect("Missing record field"),
            _ => panic!("Expected record with field ff"),
        }
    }
    pub fn __purust_borrow_ffRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.ffRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.ffRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ffRecord").expect("Missing record field"),
            _ => panic!("Expected record with field ffRecord"),
        }
    }
    pub fn __purust_borrow_fiber(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fiber_supervisor(r) => r.fiber.as_ref().unwrap(),
            Value::Record_a(r) => r.fiber.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("fiber").expect("Missing record field"),
            _ => panic!("Expected record with field fiber"),
        }
    }
    pub fn __purust_borrow_first(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_first_second(r) => r.first.as_ref().unwrap(),
            Value::Record_a(r) => r.first.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("first").expect("Missing record field"),
            _ => panic!("Expected record with field first"),
        }
    }
    pub fn __purust_borrow_foldMap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldMap.as_ref().unwrap(),
            Value::Record_a(r) => r.foldMap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMap").expect("Missing record field"),
            _ => panic!("Expected record with field foldMap"),
        }
    }
    pub fn __purust_borrow_foldMap1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldMap1.as_ref().unwrap(),
            Value::Record_a(r) => r.foldMap1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMap1").expect("Missing record field"),
            _ => panic!("Expected record with field foldMap1"),
        }
    }
    pub fn __purust_borrow_foldMapWithIndex(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldMapWithIndex.as_ref().unwrap(),
            Value::Record_a(r) => r.foldMapWithIndex.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldMapWithIndex").expect("Missing record field"),
            _ => panic!("Expected record with field foldMapWithIndex"),
        }
    }
    pub fn __purust_borrow_foldl(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldl.as_ref().unwrap(),
            Value::Record_a(r) => r.foldl.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldl").expect("Missing record field"),
            _ => panic!("Expected record with field foldl"),
        }
    }
    pub fn __purust_borrow_foldl1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldl1.as_ref().unwrap(),
            Value::Record_a(r) => r.foldl1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldl1").expect("Missing record field"),
            _ => panic!("Expected record with field foldl1"),
        }
    }
    pub fn __purust_borrow_foldlWithIndex(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldlWithIndex.as_ref().unwrap(),
            Value::Record_a(r) => r.foldlWithIndex.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldlWithIndex").expect("Missing record field"),
            _ => panic!("Expected record with field foldlWithIndex"),
        }
    }
    pub fn __purust_borrow_foldr(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_foldMap_foldl_foldr(r) => r.foldr.as_ref().unwrap(),
            Value::Record_a(r) => r.foldr.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldr").expect("Missing record field"),
            _ => panic!("Expected record with field foldr"),
        }
    }
    pub fn __purust_borrow_foldr1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => r.foldr1.as_ref().unwrap(),
            Value::Record_a(r) => r.foldr1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldr1").expect("Missing record field"),
            _ => panic!("Expected record with field foldr1"),
        }
    }
    pub fn __purust_borrow_foldrWithIndex(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => r.foldrWithIndex.as_ref().unwrap(),
            Value::Record_a(r) => r.foldrWithIndex.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("foldrWithIndex").expect("Missing record field"),
            _ => panic!("Expected record with field foldrWithIndex"),
        }
    }
    pub fn __purust_borrow_found(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_found_result(r) => r.found.as_ref().unwrap(),
            Value::Record_a(r) => r.found.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("found").expect("Missing record field"),
            _ => panic!("Expected record with field found"),
        }
    }
    pub fn __purust_borrow_from(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_from_to(r) => r.from.as_ref().unwrap(),
            Value::Record_a(r) => r.from.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("from").expect("Missing record field"),
            _ => panic!("Expected record with field from"),
        }
    }
    pub fn __purust_borrow_fromDuration(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fromDuration_toDuration(r) => r.fromDuration.as_ref().unwrap(),
            Value::Record_a(r) => r.fromDuration.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("fromDuration").expect("Missing record field"),
            _ => panic!("Expected record with field fromDuration"),
        }
    }
    pub fn __purust_borrow_fromEnum(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.fromEnum.as_ref().unwrap(),
            Value::Record_a(r) => r.fromEnum.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("fromEnum").expect("Missing record field"),
            _ => panic!("Expected record with field fromEnum"),
        }
    }
    pub fn __purust_borrow_fromLeft(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.fromLeft.as_ref().unwrap(),
            Value::Record_a(r) => r.fromLeft.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("fromLeft").expect("Missing record field"),
            _ => panic!("Expected record with field fromLeft"),
        }
    }
    pub fn __purust_borrow_fromRight(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.fromRight.as_ref().unwrap(),
            Value::Record_a(r) => r.fromRight.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("fromRight").expect("Missing record field"),
            _ => panic!("Expected record with field fromRight"),
        }
    }
    pub fn __purust_borrow_genericAdd_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericAdd_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericAdd_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericAdd'").expect("Missing record field"),
            _ => panic!("Expected record with field genericAdd_prime"),
        }
    }
    pub fn __purust_borrow_genericAppend_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericAppend_prime(r) => r.genericAppend_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericAppend_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericAppend'").expect("Missing record field"),
            _ => panic!("Expected record with field genericAppend_prime"),
        }
    }
    pub fn __purust_borrow_genericBottom_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericBottom_prime(r) => r.genericBottom_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericBottom_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericBottom'").expect("Missing record field"),
            _ => panic!("Expected record with field genericBottom_prime"),
        }
    }
    pub fn __purust_borrow_genericCardinality_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericCardinality_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericCardinality_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericCardinality'").expect("Missing record field"),
            _ => panic!("Expected record with field genericCardinality_prime"),
        }
    }
    pub fn __purust_borrow_genericCompare_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericCompare_prime(r) => r.genericCompare_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericCompare_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericCompare'").expect("Missing record field"),
            _ => panic!("Expected record with field genericCompare_prime"),
        }
    }
    pub fn __purust_borrow_genericConj_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericConj_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericConj_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericConj'").expect("Missing record field"),
            _ => panic!("Expected record with field genericConj_prime"),
        }
    }
    pub fn __purust_borrow_genericDisj_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericDisj_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericDisj_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericDisj'").expect("Missing record field"),
            _ => panic!("Expected record with field genericDisj_prime"),
        }
    }
    pub fn __purust_borrow_genericEq_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericEq_prime(r) => r.genericEq_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericEq_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericEq'").expect("Missing record field"),
            _ => panic!("Expected record with field genericEq_prime"),
        }
    }
    pub fn __purust_borrow_genericFF_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericFF_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericFF_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericFF'").expect("Missing record field"),
            _ => panic!("Expected record with field genericFF_prime"),
        }
    }
    pub fn __purust_borrow_genericFromEnum_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericFromEnum_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericFromEnum_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericFromEnum'").expect("Missing record field"),
            _ => panic!("Expected record with field genericFromEnum_prime"),
        }
    }
    pub fn __purust_borrow_genericImplies_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericImplies_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericImplies_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericImplies'").expect("Missing record field"),
            _ => panic!("Expected record with field genericImplies_prime"),
        }
    }
    pub fn __purust_borrow_genericMempty_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericMempty_prime(r) => r.genericMempty_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericMempty_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericMempty'").expect("Missing record field"),
            _ => panic!("Expected record with field genericMempty_prime"),
        }
    }
    pub fn __purust_borrow_genericMul_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericMul_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericMul_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericMul'").expect("Missing record field"),
            _ => panic!("Expected record with field genericMul_prime"),
        }
    }
    pub fn __purust_borrow_genericNot_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericNot_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericNot_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericNot'").expect("Missing record field"),
            _ => panic!("Expected record with field genericNot_prime"),
        }
    }
    pub fn __purust_borrow_genericOne_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericOne_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericOne_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericOne'").expect("Missing record field"),
            _ => panic!("Expected record with field genericOne_prime"),
        }
    }
    pub fn __purust_borrow_genericPred_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericPred_prime_genericSucc_prime(r) => r.genericPred_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericPred_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericPred'").expect("Missing record field"),
            _ => panic!("Expected record with field genericPred_prime"),
        }
    }
    pub fn __purust_borrow_genericShow_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericShow_prime(r) => r.genericShow_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericShow_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericShow'").expect("Missing record field"),
            _ => panic!("Expected record with field genericShow_prime"),
        }
    }
    pub fn __purust_borrow_genericShowArgs(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericShowArgs(r) => r.genericShowArgs.as_ref().unwrap(),
            Value::Record_a(r) => r.genericShowArgs.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericShowArgs").expect("Missing record field"),
            _ => panic!("Expected record with field genericShowArgs"),
        }
    }
    pub fn __purust_borrow_genericSub_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericSub_prime(r) => r.genericSub_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericSub_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericSub'").expect("Missing record field"),
            _ => panic!("Expected record with field genericSub_prime"),
        }
    }
    pub fn __purust_borrow_genericSucc_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericPred_prime_genericSucc_prime(r) => r.genericSucc_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericSucc_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericSucc'").expect("Missing record field"),
            _ => panic!("Expected record with field genericSucc_prime"),
        }
    }
    pub fn __purust_borrow_genericTT_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => r.genericTT_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericTT_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericTT'").expect("Missing record field"),
            _ => panic!("Expected record with field genericTT_prime"),
        }
    }
    pub fn __purust_borrow_genericToEnum_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => r.genericToEnum_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericToEnum_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericToEnum'").expect("Missing record field"),
            _ => panic!("Expected record with field genericToEnum_prime"),
        }
    }
    pub fn __purust_borrow_genericTop_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericTop_prime(r) => r.genericTop_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericTop_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericTop'").expect("Missing record field"),
            _ => panic!("Expected record with field genericTop_prime"),
        }
    }
    pub fn __purust_borrow_genericZero_prime(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => r.genericZero_prime.as_ref().unwrap(),
            Value::Record_a(r) => r.genericZero_prime.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("genericZero'").expect("Missing record field"),
            _ => panic!("Expected record with field genericZero_prime"),
        }
    }
    pub fn __purust_borrow_global(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.global.as_ref().unwrap(),
            Value::Record_a(r) => r.global.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("global").expect("Missing record field"),
            _ => panic!("Expected record with field global"),
        }
    }
    pub fn __purust_borrow_handler(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_handler_rethrow(r) => r.handler.as_ref().unwrap(),
            Value::Record_a(r) => r.handler.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("handler").expect("Missing record field"),
            _ => panic!("Expected record with field handler"),
        }
    }
    pub fn __purust_borrow_head(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_head_tail(r) => r.head.as_ref().unwrap(),
            Value::Record_a(r) => r.head.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("head").expect("Missing record field"),
            _ => panic!("Expected record with field head"),
        }
    }
    pub fn __purust_borrow_hour(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.hour.as_ref().unwrap(),
            Value::Record_a(r) => r.hour.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("hour").expect("Missing record field"),
            _ => panic!("Expected record with field hour"),
        }
    }
    pub fn __purust_borrow_identity(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Semigroupoid0_identity(r) => r.identity.as_ref().unwrap(),
            Value::Record_a(r) => r.identity.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("identity").expect("Missing record field"),
            _ => panic!("Expected record with field identity"),
        }
    }
    pub fn __purust_borrow_ignoreCase(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.ignoreCase.as_ref().unwrap(),
            Value::Record_a(r) => r.ignoreCase.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ignoreCase").expect("Missing record field"),
            _ => panic!("Expected record with field ignoreCase"),
        }
    }
    pub fn __purust_borrow_imap(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_imap(r) => r.imap.as_ref().unwrap(),
            Value::Record_a(r) => r.imap.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("imap").expect("Missing record field"),
            _ => panic!("Expected record with field imap"),
        }
    }
    pub fn __purust_borrow_implies(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.implies.as_ref().unwrap(),
            Value::Record_a(r) => r.implies.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("implies").expect("Missing record field"),
            _ => panic!("Expected record with field implies"),
        }
    }
    pub fn __purust_borrow_impliesRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.impliesRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.impliesRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("impliesRecord").expect("Missing record field"),
            _ => panic!("Expected record with field impliesRecord"),
        }
    }
    pub fn __purust_borrow_index(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_index_value(r) => r.index.as_ref().unwrap(),
            Value::Record_a(r) => r.index.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("index").expect("Missing record field"),
            _ => panic!("Expected record with field index"),
        }
    }
    pub fn __purust_borrow_init(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_acc_init(r) => r.init.as_ref().unwrap(),
            Value::Record_init_last(r) => r.init.as_ref().unwrap(),
            Value::Record_init_rest(r) => r.init.as_ref().unwrap(),
            Value::Record_a(r) => r.init.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("init").expect("Missing record field"),
            _ => panic!("Expected record with field init"),
        }
    }
    pub fn __purust_borrow_inj(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_inj_prj(r) => r.inj.as_ref().unwrap(),
            Value::Record_a(r) => r.inj.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("inj").expect("Missing record field"),
            _ => panic!("Expected record with field inj"),
        }
    }
    pub fn __purust_borrow_isLeft(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.isLeft.as_ref().unwrap(),
            Value::Record_a(r) => r.isLeft.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("isLeft").expect("Missing record field"),
            _ => panic!("Expected record with field isLeft"),
        }
    }
    pub fn __purust_borrow_isSuspended(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.isSuspended.as_ref().unwrap(),
            Value::Record_a(r) => r.isSuspended.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("isSuspended").expect("Missing record field"),
            _ => panic!("Expected record with field isSuspended"),
        }
    }
    pub fn __purust_borrow_join(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.join.as_ref().unwrap(),
            Value::Record_a(r) => r.join.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("join").expect("Missing record field"),
            _ => panic!("Expected record with field join"),
        }
    }
    pub fn __purust_borrow_key(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_key_value(r) => r.key.as_ref().unwrap(),
            Value::Record_a(r) => r.key.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("key").expect("Missing record field"),
            _ => panic!("Expected record with field key"),
        }
    }
    pub fn __purust_borrow_keysImpl(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_keysImpl(r) => r.keysImpl.as_ref().unwrap(),
            Value::Record_a(r) => r.keysImpl.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("keysImpl").expect("Missing record field"),
            _ => panic!("Expected record with field keysImpl"),
        }
    }
    pub fn __purust_borrow_kill(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.kill.as_ref().unwrap(),
            Value::Record_a(r) => r.kill.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("kill").expect("Missing record field"),
            _ => panic!("Expected record with field kill"),
        }
    }
    pub fn __purust_borrow_killed(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_completed_failed_killed(r) => r.killed.as_ref().unwrap(),
            Value::Record_a(r) => r.killed.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("killed").expect("Missing record field"),
            _ => panic!("Expected record with field killed"),
        }
    }
    pub fn __purust_borrow_last(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_init_last(r) => r.last.as_ref().unwrap(),
            Value::Record_last_revInit(r) => r.last.as_ref().unwrap(),
            Value::Record_a(r) => r.last.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("last").expect("Missing record field"),
            _ => panic!("Expected record with field last"),
        }
    }
    pub fn __purust_borrow_left(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_left_right(r) => r.left.as_ref().unwrap(),
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.left.as_ref().unwrap(),
            Value::Record_a(r) => r.left.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("left").expect("Missing record field"),
            _ => panic!("Expected record with field left"),
        }
    }
    pub fn __purust_borrow_lift(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_lift(r) => r.lift.as_ref().unwrap(),
            Value::Record_a(r) => r.lift.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("lift").expect("Missing record field"),
            _ => panic!("Expected record with field lift"),
        }
    }
    pub fn __purust_borrow_liftAff(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadEffect0_liftAff(r) => r.liftAff.as_ref().unwrap(),
            Value::Record_a(r) => r.liftAff.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("liftAff").expect("Missing record field"),
            _ => panic!("Expected record with field liftAff"),
        }
    }
    pub fn __purust_borrow_liftEffect(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_liftEffect(r) => r.liftEffect.as_ref().unwrap(),
            Value::Record_a(r) => r.liftEffect.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("liftEffect").expect("Missing record field"),
            _ => panic!("Expected record with field liftEffect"),
        }
    }
    pub fn __purust_borrow_liftST(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_liftST(r) => r.liftST.as_ref().unwrap(),
            Value::Record_a(r) => r.liftST.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("liftST").expect("Missing record field"),
            _ => panic!("Expected record with field liftST"),
        }
    }
    pub fn __purust_borrow_listen(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.listen.as_ref().unwrap(),
            Value::Record_a(r) => r.listen.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("listen").expect("Missing record field"),
            _ => panic!("Expected record with field listen"),
        }
    }
    pub fn __purust_borrow_local(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_ComonadAsk0_local(r) => r.local.as_ref().unwrap(),
            Value::Record_MonadAsk0_local(r) => r.local.as_ref().unwrap(),
            Value::Record_a(r) => r.local.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("local").expect("Missing record field"),
            _ => panic!("Expected record with field local"),
        }
    }
    pub fn __purust_borrow_lose(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Decide0_Divisible1_lose(r) => r.lose.as_ref().unwrap(),
            Value::Record_a(r) => r.lose.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("lose").expect("Missing record field"),
            _ => panic!("Expected record with field lose"),
        }
    }
    pub fn __purust_borrow_lower(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_lower(r) => r.lower.as_ref().unwrap(),
            Value::Record_a(r) => r.lower.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("lower").expect("Missing record field"),
            _ => panic!("Expected record with field lower"),
        }
    }
    pub fn __purust_borrow_map(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_map(r) => r.map.as_ref().unwrap(),
            Value::Record_a(r) => r.map.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("map").expect("Missing record field"),
            _ => panic!("Expected record with field map"),
        }
    }
    pub fn __purust_borrow_mapWithIndex(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Functor0_mapWithIndex(r) => r.mapWithIndex.as_ref().unwrap(),
            Value::Record_a(r) => r.mapWithIndex.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mapWithIndex").expect("Missing record field"),
            _ => panic!("Expected record with field mapWithIndex"),
        }
    }
    pub fn __purust_borrow_mappend_(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_mappend__mempty_(r) => r.mappend_.as_ref().unwrap(),
            Value::Record_a(r) => r.mappend_.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mappend_").expect("Missing record field"),
            _ => panic!("Expected record with field mappend_"),
        }
    }
    pub fn __purust_borrow_mempty(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Semigroup0_mempty(r) => r.mempty.as_ref().unwrap(),
            Value::Record_a(r) => r.mempty.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mempty").expect("Missing record field"),
            _ => panic!("Expected record with field mempty"),
        }
    }
    pub fn __purust_borrow_memptyRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_SemigroupRecord0_memptyRecord(r) => r.memptyRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.memptyRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("memptyRecord").expect("Missing record field"),
            _ => panic!("Expected record with field memptyRecord"),
        }
    }
    pub fn __purust_borrow_mempty_(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_mappend__mempty_(r) => r.mempty_.as_ref().unwrap(),
            Value::Record_a(r) => r.mempty_.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mempty_").expect("Missing record field"),
            _ => panic!("Expected record with field mempty_"),
        }
    }
    pub fn __purust_borrow_millisecond(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.millisecond.as_ref().unwrap(),
            Value::Record_a(r) => r.millisecond.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("millisecond").expect("Missing record field"),
            _ => panic!("Expected record with field millisecond"),
        }
    }
    pub fn __purust_borrow_minute(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.minute.as_ref().unwrap(),
            Value::Record_a(r) => r.minute.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("minute").expect("Missing record field"),
            _ => panic!("Expected record with field minute"),
        }
    }
    pub fn __purust_borrow_mod_kw(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => r.mod_kw.as_ref().unwrap(),
            Value::Record_a(r) => r.mod_kw.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mod").expect("Missing record field"),
            _ => panic!("Expected record with field mod_kw"),
        }
    }
    pub fn __purust_borrow_month(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.month.as_ref().unwrap(),
            Value::Record_a(r) => r.month.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("month").expect("Missing record field"),
            _ => panic!("Expected record with field month"),
        }
    }
    pub fn __purust_borrow_mul(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.mul.as_ref().unwrap(),
            Value::Record_a(r) => r.mul.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mul").expect("Missing record field"),
            _ => panic!("Expected record with field mul"),
        }
    }
    pub fn __purust_borrow_mulRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.mulRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.mulRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("mulRecord").expect("Missing record field"),
            _ => panic!("Expected record with field mulRecord"),
        }
    }
    pub fn __purust_borrow_multiline(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.multiline.as_ref().unwrap(),
            Value::Record_a(r) => r.multiline.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("multiline").expect("Missing record field"),
            _ => panic!("Expected record with field multiline"),
        }
    }
    pub fn __purust_borrow_myMethod(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_myMethod(r) => r.myMethod.as_ref().unwrap(),
            Value::Record_a(r) => r.myMethod.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("myMethod").expect("Missing record field"),
            _ => panic!("Expected record with field myMethod"),
        }
    }
    pub fn __purust_borrow_nes(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_nes(r) => r.nes.as_ref().unwrap(),
            Value::Record_a(r) => r.nes.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("nes").expect("Missing record field"),
            _ => panic!("Expected record with field nes"),
        }
    }
    pub fn __purust_borrow_no(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_no_yes(r) => r.no.as_ref().unwrap(),
            Value::Record_a(r) => r.no.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("no").expect("Missing record field"),
            _ => panic!("Expected record with field no"),
        }
    }
    pub fn __purust_borrow_not(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.not.as_ref().unwrap(),
            Value::Record_a(r) => r.not.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("not").expect("Missing record field"),
            _ => panic!("Expected record with field not"),
        }
    }
    pub fn __purust_borrow_notRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.notRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.notRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("notRecord").expect("Missing record field"),
            _ => panic!("Expected record with field notRecord"),
        }
    }
    pub fn __purust_borrow_onComplete(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.onComplete.as_ref().unwrap(),
            Value::Record_a(r) => r.onComplete.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("onComplete").expect("Missing record field"),
            _ => panic!("Expected record with field onComplete"),
        }
    }
    pub fn __purust_borrow_one(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.one.as_ref().unwrap(),
            Value::Record_a(r) => r.one.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("one").expect("Missing record field"),
            _ => panic!("Expected record with field one"),
        }
    }
    pub fn __purust_borrow_oneRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.oneRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.oneRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("oneRecord").expect("Missing record field"),
            _ => panic!("Expected record with field oneRecord"),
        }
    }
    pub fn __purust_borrow_parallel(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.parallel.as_ref().unwrap(),
            Value::Record_a(r) => r.parallel.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("parallel").expect("Missing record field"),
            _ => panic!("Expected record with field parallel"),
        }
    }
    pub fn __purust_borrow_pass(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => r.pass.as_ref().unwrap(),
            Value::Record_a(r) => r.pass.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("pass").expect("Missing record field"),
            _ => panic!("Expected record with field pass"),
        }
    }
    pub fn __purust_borrow_peek(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_peek_pos(r) => r.peek.as_ref().unwrap(),
            Value::Record_a(r) => r.peek.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("peek").expect("Missing record field"),
            _ => panic!("Expected record with field peek"),
        }
    }
    pub fn __purust_borrow_pos(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_peek_pos(r) => r.pos.as_ref().unwrap(),
            Value::Record_elem_pos(r) => r.pos.as_ref().unwrap(),
            Value::Record_a(r) => r.pos.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("pos").expect("Missing record field"),
            _ => panic!("Expected record with field pos"),
        }
    }
    pub fn __purust_borrow_pred(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ord0_pred_succ(r) => r.pred.as_ref().unwrap(),
            Value::Record_a(r) => r.pred.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("pred").expect("Missing record field"),
            _ => panic!("Expected record with field pred"),
        }
    }
    pub fn __purust_borrow_prj(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_inj_prj(r) => r.prj.as_ref().unwrap(),
            Value::Record_a(r) => r.prj.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("prj").expect("Missing record field"),
            _ => panic!("Expected record with field prj"),
        }
    }
    pub fn __purust_borrow_proof(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Coercible0_proof(r) => r.proof.as_ref().unwrap(),
            Value::Record_a(r) => r.proof.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("proof").expect("Missing record field"),
            _ => panic!("Expected record with field proof"),
        }
    }
    pub fn __purust_borrow_ps_minus_rust_minus_test(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_ps_minus_rust_minus_test(r) => r.ps_minus_rust_minus_test.as_ref().unwrap(),
            Value::Record_a(r) => r.ps_minus_rust_minus_test.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ps-rust-test").expect("Missing record field"),
            _ => panic!("Expected record with field ps_minus_rust_minus_test"),
        }
    }
    pub fn __purust_borrow_pure(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_pure(r) => r.pure.as_ref().unwrap(),
            Value::Record_a(r) => r.pure.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("pure").expect("Missing record field"),
            _ => panic!("Expected record with field pure"),
        }
    }
    pub fn __purust_borrow_recip(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ring0_recip(r) => r.recip.as_ref().unwrap(),
            Value::Record_a(r) => r.recip.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("recip").expect("Missing record field"),
            _ => panic!("Expected record with field recip"),
        }
    }
    pub fn __purust_borrow_reflectSymbol(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_reflectSymbol(r) => r.reflectSymbol.as_ref().unwrap(),
            Value::Record_a(r) => r.reflectSymbol.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("reflectSymbol").expect("Missing record field"),
            _ => panic!("Expected record with field reflectSymbol"),
        }
    }
    pub fn __purust_borrow_reflectType(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_reflectType(r) => r.reflectType.as_ref().unwrap(),
            Value::Record_a(r) => r.reflectType.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("reflectType").expect("Missing record field"),
            _ => panic!("Expected record with field reflectType"),
        }
    }
    pub fn __purust_borrow_rest(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_init_rest(r) => r.rest.as_ref().unwrap(),
            Value::Record_a(r) => r.rest.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("rest").expect("Missing record field"),
            _ => panic!("Expected record with field rest"),
        }
    }
    pub fn __purust_borrow_result(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_found_result(r) => r.result.as_ref().unwrap(),
            Value::Record_a(r) => r.result.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("result").expect("Missing record field"),
            _ => panic!("Expected record with field result"),
        }
    }
    pub fn __purust_borrow_rethrow(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_handler_rethrow(r) => r.rethrow.as_ref().unwrap(),
            Value::Record_a(r) => r.rethrow.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("rethrow").expect("Missing record field"),
            _ => panic!("Expected record with field rethrow"),
        }
    }
    pub fn __purust_borrow_revInit(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_last_revInit(r) => r.revInit.as_ref().unwrap(),
            Value::Record_a(r) => r.revInit.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("revInit").expect("Missing record field"),
            _ => panic!("Expected record with field revInit"),
        }
    }
    pub fn __purust_borrow_right(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_left_right(r) => r.right.as_ref().unwrap(),
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => r.right.as_ref().unwrap(),
            Value::Record_a(r) => r.right.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("right").expect("Missing record field"),
            _ => panic!("Expected record with field right"),
        }
    }
    pub fn __purust_borrow_run(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => r.run.as_ref().unwrap(),
            Value::Record_a(r) => r.run.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("run").expect("Missing record field"),
            _ => panic!("Expected record with field run"),
        }
    }
    pub fn __purust_borrow_second(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Profunctor0_first_second(r) => r.second.as_ref().unwrap(),
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.second.as_ref().unwrap(),
            Value::Record_a(r) => r.second.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("second").expect("Missing record field"),
            _ => panic!("Expected record with field second"),
        }
    }
    pub fn __purust_borrow_sequence(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.sequence.as_ref().unwrap(),
            Value::Record_a(r) => r.sequence.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("sequence").expect("Missing record field"),
            _ => panic!("Expected record with field sequence"),
        }
    }
    pub fn __purust_borrow_sequence1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.sequence1.as_ref().unwrap(),
            Value::Record_a(r) => r.sequence1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("sequence1").expect("Missing record field"),
            _ => panic!("Expected record with field sequence1"),
        }
    }
    pub fn __purust_borrow_sequential(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => r.sequential.as_ref().unwrap(),
            Value::Record_a(r) => r.sequential.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("sequential").expect("Missing record field"),
            _ => panic!("Expected record with field sequential"),
        }
    }
    pub fn __purust_borrow_show(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_show(r) => r.show.as_ref().unwrap(),
            Value::Record_a(r) => r.show.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("show").expect("Missing record field"),
            _ => panic!("Expected record with field show"),
        }
    }
    pub fn __purust_borrow_showRecordFields(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_showRecordFields(r) => r.showRecordFields.as_ref().unwrap(),
            Value::Record_a(r) => r.showRecordFields.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("showRecordFields").expect("Missing record field"),
            _ => panic!("Expected record with field showRecordFields"),
        }
    }
    pub fn __purust_borrow_state(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_state(r) => r.state.as_ref().unwrap(),
            Value::Record_state_val(r) => r.state.as_ref().unwrap(),
            Value::Record_state_value(r) => r.state.as_ref().unwrap(),
            Value::Record_a(r) => r.state.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("state").expect("Missing record field"),
            _ => panic!("Expected record with field state"),
        }
    }
    pub fn __purust_borrow_sticky(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.sticky.as_ref().unwrap(),
            Value::Record_a(r) => r.sticky.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("sticky").expect("Missing record field"),
            _ => panic!("Expected record with field sticky"),
        }
    }
    pub fn __purust_borrow_sub(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Semiring0_sub(r) => r.sub.as_ref().unwrap(),
            Value::Record_a(r) => r.sub.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("sub").expect("Missing record field"),
            _ => panic!("Expected record with field sub"),
        }
    }
    pub fn __purust_borrow_subRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_SemiringRecord0_subRecord(r) => r.subRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.subRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("subRecord").expect("Missing record field"),
            _ => panic!("Expected record with field subRecord"),
        }
    }
    pub fn __purust_borrow_succ(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ord0_pred_succ(r) => r.succ.as_ref().unwrap(),
            Value::Record_a(r) => r.succ.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("succ").expect("Missing record field"),
            _ => panic!("Expected record with field succ"),
        }
    }
    pub fn __purust_borrow_supervisor(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fiber_supervisor(r) => r.supervisor.as_ref().unwrap(),
            Value::Record_a(r) => r.supervisor.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("supervisor").expect("Missing record field"),
            _ => panic!("Expected record with field supervisor"),
        }
    }
    pub fn __purust_borrow_tail(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_head_tail(r) => r.tail.as_ref().unwrap(),
            Value::Record_a(r) => r.tail.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("tail").expect("Missing record field"),
            _ => panic!("Expected record with field tail"),
        }
    }
    pub fn __purust_borrow_tailRecM(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_tailRecM(r) => r.tailRecM.as_ref().unwrap(),
            Value::Record_a(r) => r.tailRecM.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("tailRecM").expect("Missing record field"),
            _ => panic!("Expected record with field tailRecM"),
        }
    }
    pub fn __purust_borrow_tell(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad1_Semigroup0_tell(r) => r.tell.as_ref().unwrap(),
            Value::Record_a(r) => r.tell.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("tell").expect("Missing record field"),
            _ => panic!("Expected record with field tell"),
        }
    }
    pub fn __purust_borrow_throwError(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Monad0_throwError(r) => r.throwError.as_ref().unwrap(),
            Value::Record_a(r) => r.throwError.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("throwError").expect("Missing record field"),
            _ => panic!("Expected record with field throwError"),
        }
    }
    pub fn __purust_borrow_to(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_from_to(r) => r.to.as_ref().unwrap(),
            Value::Record_a(r) => r.to.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("to").expect("Missing record field"),
            _ => panic!("Expected record with field to"),
        }
    }
    pub fn __purust_borrow_toDuration(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_fromDuration_toDuration(r) => r.toDuration.as_ref().unwrap(),
            Value::Record_a(r) => r.toDuration.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("toDuration").expect("Missing record field"),
            _ => panic!("Expected record with field toDuration"),
        }
    }
    pub fn __purust_borrow_toEnum(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => r.toEnum.as_ref().unwrap(),
            Value::Record_a(r) => r.toEnum.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("toEnum").expect("Missing record field"),
            _ => panic!("Expected record with field toEnum"),
        }
    }
    pub fn __purust_borrow_top(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Ord0_bottom_top(r) => r.top.as_ref().unwrap(),
            Value::Record_a(r) => r.top.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("top").expect("Missing record field"),
            _ => panic!("Expected record with field top"),
        }
    }
    pub fn __purust_borrow_topRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => r.topRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.topRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("topRecord").expect("Missing record field"),
            _ => panic!("Expected record with field topRecord"),
        }
    }
    pub fn __purust_borrow_track(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Comonad0_track(r) => r.track.as_ref().unwrap(),
            Value::Record_a(r) => r.track.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("track").expect("Missing record field"),
            _ => panic!("Expected record with field track"),
        }
    }
    pub fn __purust_borrow_traverse(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => r.traverse.as_ref().unwrap(),
            Value::Record_a(r) => r.traverse.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("traverse").expect("Missing record field"),
            _ => panic!("Expected record with field traverse"),
        }
    }
    pub fn __purust_borrow_traverse1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => r.traverse1.as_ref().unwrap(),
            Value::Record_a(r) => r.traverse1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("traverse1").expect("Missing record field"),
            _ => panic!("Expected record with field traverse1"),
        }
    }
    pub fn __purust_borrow_traverseWithIndex(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => r.traverseWithIndex.as_ref().unwrap(),
            Value::Record_a(r) => r.traverseWithIndex.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("traverseWithIndex").expect("Missing record field"),
            _ => panic!("Expected record with field traverseWithIndex"),
        }
    }
    pub fn __purust_borrow_tt(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conj_disj_ff_implies_not_tt(r) => r.tt.as_ref().unwrap(),
            Value::Record_a(r) => r.tt.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("tt").expect("Missing record field"),
            _ => panic!("Expected record with field tt"),
        }
    }
    pub fn __purust_borrow_ttRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => r.ttRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.ttRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("ttRecord").expect("Missing record field"),
            _ => panic!("Expected record with field ttRecord"),
        }
    }
    pub fn __purust_borrow_unfoldr(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_Unfoldable10_unfoldr(r) => r.unfoldr.as_ref().unwrap(),
            Value::Record_a(r) => r.unfoldr.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("unfoldr").expect("Missing record field"),
            _ => panic!("Expected record with field unfoldr"),
        }
    }
    pub fn __purust_borrow_unfoldr1(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_unfoldr1(r) => r.unfoldr1.as_ref().unwrap(),
            Value::Record_a(r) => r.unfoldr1.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("unfoldr1").expect("Missing record field"),
            _ => panic!("Expected record with field unfoldr1"),
        }
    }
    pub fn __purust_borrow_unicode(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => r.unicode.as_ref().unwrap(),
            Value::Record_a(r) => r.unicode.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("unicode").expect("Missing record field"),
            _ => panic!("Expected record with field unicode"),
        }
    }
    pub fn __purust_borrow_val(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_acc_val(r) => r.val.as_ref().unwrap(),
            Value::Record_state_val(r) => r.val.as_ref().unwrap(),
            Value::Record_a(r) => r.val.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("val").expect("Missing record field"),
            _ => panic!("Expected record with field val"),
        }
    }
    pub fn __purust_borrow_value(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_accum_value(r) => r.value.as_ref().unwrap(),
            Value::Record_index_value(r) => r.value.as_ref().unwrap(),
            Value::Record_key_value(r) => r.value.as_ref().unwrap(),
            Value::Record_state_value(r) => r.value.as_ref().unwrap(),
            Value::Record_a(r) => r.value.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("value").expect("Missing record field"),
            _ => panic!("Expected record with field value"),
        }
    }
    pub fn __purust_borrow_year(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => r.year.as_ref().unwrap(),
            Value::Record_a(r) => r.year.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("year").expect("Missing record field"),
            _ => panic!("Expected record with field year"),
        }
    }
    pub fn __purust_borrow_yes(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_no_yes(r) => r.yes.as_ref().unwrap(),
            Value::Record_a(r) => r.yes.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("yes").expect("Missing record field"),
            _ => panic!("Expected record with field yes"),
        }
    }
    pub fn __purust_borrow_zero(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_add_mul_one_zero(r) => r.zero.as_ref().unwrap(),
            Value::Record_a(r) => r.zero.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("zero").expect("Missing record field"),
            _ => panic!("Expected record with field zero"),
        }
    }
    pub fn __purust_borrow_zeroRecord(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => r.zeroRecord.as_ref().unwrap(),
            Value::Record_a(r) => r.zeroRecord.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("zeroRecord").expect("Missing record field"),
            _ => panic!("Expected record with field zeroRecord"),
        }
    }
    pub fn set_Alt0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Alt0_empty(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Alt0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Alt0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Alt0".to_owned(), val); },
            _ => panic!("Expected record with field Alt0"),
        }
    }
    pub fn set_Alternative1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Alternative1_Monad0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Alternative1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Alternative1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Alternative1".to_owned(), val); },
            _ => panic!("Expected record with field Alternative1"),
        }
    }
    pub fn set_Applicative0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Applicative0_Bind1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Applicative0 = Some(val);
            },
            Value::Record_Applicative0_Plus1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Applicative0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Applicative0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Applicative0".to_owned(), val); },
            _ => panic!("Expected record with field Applicative0"),
        }
    }
    pub fn set_Apply0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply0 = Some(val);
            },
            Value::Record_Apply0_bind(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply0 = Some(val);
            },
            Value::Record_Apply0_pure(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Apply0".to_owned(), val); },
            _ => panic!("Expected record with field Apply0"),
        }
    }
    pub fn set_Apply1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Apply1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Apply1".to_owned(), val); },
            _ => panic!("Expected record with field Apply1"),
        }
    }
    pub fn set_Biapply0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Biapply0_bipure(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Biapply0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Biapply0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Biapply0".to_owned(), val); },
            _ => panic!("Expected record with field Biapply0"),
        }
    }
    pub fn set_Bifoldable1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bifoldable1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bifoldable1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Bifoldable1".to_owned(), val); },
            _ => panic!("Expected record with field Bifoldable1"),
        }
    }
    pub fn set_Bifunctor0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bifunctor0 = Some(val);
            },
            Value::Record_Bifunctor0_biapply(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bifunctor0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bifunctor0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Bifunctor0".to_owned(), val); },
            _ => panic!("Expected record with field Bifunctor0"),
        }
    }
    pub fn set_Bind1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Applicative0_Bind1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bind1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bind1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Bind1".to_owned(), val); },
            _ => panic!("Expected record with field Bind1"),
        }
    }
    pub fn set_Bounded0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bounded0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Bounded0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Bounded0".to_owned(), val); },
            _ => panic!("Expected record with field Bounded0"),
        }
    }
    pub fn set_Coercible0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Coercible0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Coercible0 = Some(val);
            },
            Value::Record_Coercible0_proof(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Coercible0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Coercible0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Coercible0".to_owned(), val); },
            _ => panic!("Expected record with field Coercible0"),
        }
    }
    pub fn set_CommutativeRing0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.CommutativeRing0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.CommutativeRing0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("CommutativeRing0".to_owned(), val); },
            _ => panic!("Expected record with field CommutativeRing0"),
        }
    }
    pub fn set_Comonad0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Comonad0_ask(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Comonad0 = Some(val);
            },
            Value::Record_Comonad0_peek_pos(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Comonad0 = Some(val);
            },
            Value::Record_Comonad0_track(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Comonad0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Comonad0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Comonad0".to_owned(), val); },
            _ => panic!("Expected record with field Comonad0"),
        }
    }
    pub fn set_ComonadAsk0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_ComonadAsk0_local(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ComonadAsk0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ComonadAsk0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ComonadAsk0".to_owned(), val); },
            _ => panic!("Expected record with field ComonadAsk0"),
        }
    }
    pub fn set_Contravariant0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Contravariant0_divide(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Contravariant0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Contravariant0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Contravariant0".to_owned(), val); },
            _ => panic!("Expected record with field Contravariant0"),
        }
    }
    pub fn set_Decide0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Decide0_Divisible1_lose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Decide0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Decide0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Decide0".to_owned(), val); },
            _ => panic!("Expected record with field Decide0"),
        }
    }
    pub fn set_Divide0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Divide0_choose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Divide0 = Some(val);
            },
            Value::Record_Divide0_conquer(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Divide0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Divide0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Divide0".to_owned(), val); },
            _ => panic!("Expected record with field Divide0"),
        }
    }
    pub fn set_Divisible1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Decide0_Divisible1_lose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Divisible1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Divisible1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Divisible1".to_owned(), val); },
            _ => panic!("Expected record with field Divisible1"),
        }
    }
    pub fn set_DivisionRing1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_DivisionRing1_EuclideanRing0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.DivisionRing1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.DivisionRing1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("DivisionRing1".to_owned(), val); },
            _ => panic!("Expected record with field DivisionRing1"),
        }
    }
    pub fn set_Enum1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Enum1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Enum1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Enum1".to_owned(), val); },
            _ => panic!("Expected record with field Enum1"),
        }
    }
    pub fn set_Eq0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Eq0_compare(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Eq0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Eq0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Eq0".to_owned(), val); },
            _ => panic!("Expected record with field Eq0"),
        }
    }
    pub fn set_Eq10(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Eq10_compare1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Eq10 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Eq10 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Eq10".to_owned(), val); },
            _ => panic!("Expected record with field Eq10"),
        }
    }
    pub fn set_EqRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_EqRecord0_compareRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.EqRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.EqRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("EqRecord0".to_owned(), val); },
            _ => panic!("Expected record with field EqRecord0"),
        }
    }
    pub fn set_EuclideanRing0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_DivisionRing1_EuclideanRing0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.EuclideanRing0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.EuclideanRing0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("EuclideanRing0".to_owned(), val); },
            _ => panic!("Expected record with field EuclideanRing0"),
        }
    }
    pub fn set_Extend0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Extend0_extract(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Extend0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Extend0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Extend0".to_owned(), val); },
            _ => panic!("Expected record with field Extend0"),
        }
    }
    pub fn set_Foldable0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable0 = Some(val);
            },
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Foldable0".to_owned(), val); },
            _ => panic!("Expected record with field Foldable0"),
        }
    }
    pub fn set_Foldable1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Foldable1".to_owned(), val); },
            _ => panic!("Expected record with field Foldable1"),
        }
    }
    pub fn set_Foldable10(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable10 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Foldable10 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Foldable10".to_owned(), val); },
            _ => panic!("Expected record with field Foldable10"),
        }
    }
    pub fn set_FoldableWithIndex1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.FoldableWithIndex1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.FoldableWithIndex1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("FoldableWithIndex1".to_owned(), val); },
            _ => panic!("Expected record with field FoldableWithIndex1"),
        }
    }
    pub fn set_Functor0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_Functor0_alt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_Functor0_apply(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_Functor0_collect_distribute(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_Functor0_extend(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_Functor0_mapWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Functor0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Functor0".to_owned(), val); },
            _ => panic!("Expected record with field Functor0"),
        }
    }
    pub fn set_FunctorWithIndex0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.FunctorWithIndex0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.FunctorWithIndex0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("FunctorWithIndex0".to_owned(), val); },
            _ => panic!("Expected record with field FunctorWithIndex0"),
        }
    }
    pub fn set_HeytingAlgebra0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_HeytingAlgebra0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.HeytingAlgebra0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.HeytingAlgebra0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("HeytingAlgebra0".to_owned(), val); },
            _ => panic!("Expected record with field HeytingAlgebra0"),
        }
    }
    pub fn set_HeytingAlgebraRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_HeytingAlgebraRecord0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.HeytingAlgebraRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.HeytingAlgebraRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("HeytingAlgebraRecord0".to_owned(), val); },
            _ => panic!("Expected record with field HeytingAlgebraRecord0"),
        }
    }
    pub fn set_Monad0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Alternative1_Monad0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_ask(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_callCC(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_liftEffect(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_liftST(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_state(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_tailRecM(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_Monad0_throwError(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Monad0".to_owned(), val); },
            _ => panic!("Expected record with field Monad0"),
        }
    }
    pub fn set_Monad1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad1_Semigroup0_tell(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monad1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Monad1".to_owned(), val); },
            _ => panic!("Expected record with field Monad1"),
        }
    }
    pub fn set_MonadAsk0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadAsk0_local(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadAsk0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadAsk0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("MonadAsk0".to_owned(), val); },
            _ => panic!("Expected record with field MonadAsk0"),
        }
    }
    pub fn set_MonadEffect0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadEffect0_liftAff(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadEffect0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadEffect0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("MonadEffect0".to_owned(), val); },
            _ => panic!("Expected record with field MonadEffect0"),
        }
    }
    pub fn set_MonadTell1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadTell1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadTell1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("MonadTell1".to_owned(), val); },
            _ => panic!("Expected record with field MonadTell1"),
        }
    }
    pub fn set_MonadThrow0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadThrow0_catchError(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadThrow0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.MonadThrow0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("MonadThrow0".to_owned(), val); },
            _ => panic!("Expected record with field MonadThrow0"),
        }
    }
    pub fn set_Monoid0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monoid0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Monoid0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Monoid0".to_owned(), val); },
            _ => panic!("Expected record with field Monoid0"),
        }
    }
    pub fn set_Ord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ord0_bottom_top(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ord0 = Some(val);
            },
            Value::Record_Ord0_pred_succ(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Ord0".to_owned(), val); },
            _ => panic!("Expected record with field Ord0"),
        }
    }
    pub fn set_OrdRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.OrdRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.OrdRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("OrdRecord0".to_owned(), val); },
            _ => panic!("Expected record with field OrdRecord0"),
        }
    }
    pub fn set_Plus1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Applicative0_Plus1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Plus1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Plus1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Plus1".to_owned(), val); },
            _ => panic!("Expected record with field Plus1"),
        }
    }
    pub fn set_Profunctor0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_closed(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Profunctor0 = Some(val);
            },
            Value::Record_Profunctor0_first_second(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Profunctor0 = Some(val);
            },
            Value::Record_Profunctor0_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Profunctor0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Profunctor0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Profunctor0".to_owned(), val); },
            _ => panic!("Expected record with field Profunctor0"),
        }
    }
    pub fn set_Ring0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ring0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ring0 = Some(val);
            },
            Value::Record_Ring0_recip(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ring0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Ring0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Ring0".to_owned(), val); },
            _ => panic!("Expected record with field Ring0"),
        }
    }
    pub fn set_RingRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_RingRecord0(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.RingRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.RingRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("RingRecord0".to_owned(), val); },
            _ => panic!("Expected record with field RingRecord0"),
        }
    }
    pub fn set_Semigroup0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad1_Semigroup0_tell(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semigroup0 = Some(val);
            },
            Value::Record_Semigroup0_mempty(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semigroup0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semigroup0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Semigroup0".to_owned(), val); },
            _ => panic!("Expected record with field Semigroup0"),
        }
    }
    pub fn set_SemigroupRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_SemigroupRecord0_memptyRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.SemigroupRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.SemigroupRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("SemigroupRecord0".to_owned(), val); },
            _ => panic!("Expected record with field SemigroupRecord0"),
        }
    }
    pub fn set_Semigroupoid0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Semigroupoid0_identity(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semigroupoid0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semigroupoid0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Semigroupoid0".to_owned(), val); },
            _ => panic!("Expected record with field Semigroupoid0"),
        }
    }
    pub fn set_Semiring0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Semiring0_sub(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semiring0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Semiring0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Semiring0".to_owned(), val); },
            _ => panic!("Expected record with field Semiring0"),
        }
    }
    pub fn set_SemiringRecord0(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_SemiringRecord0_subRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.SemiringRecord0 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.SemiringRecord0 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("SemiringRecord0".to_owned(), val); },
            _ => panic!("Expected record with field SemiringRecord0"),
        }
    }
    pub fn set_Traversable1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Traversable1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Traversable1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Traversable1".to_owned(), val); },
            _ => panic!("Expected record with field Traversable1"),
        }
    }
    pub fn set_Traversable2(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Traversable2 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Traversable2 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Traversable2".to_owned(), val); },
            _ => panic!("Expected record with field Traversable2"),
        }
    }
    pub fn set_Unfoldable10(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Unfoldable10_unfoldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Unfoldable10 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.Unfoldable10 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("Unfoldable10".to_owned(), val); },
            _ => panic!("Expected record with field Unfoldable10"),
        }
    }
    pub fn set_a(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::Record_a_b_c(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::Record_a_b_c_d_e(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("a".to_owned(), val); },
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn set_acc(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_acc_init(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.acc = Some(val);
            },
            Value::Record_acc_val(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.acc = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.acc = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("acc".to_owned(), val); },
            _ => panic!("Expected record with field acc"),
        }
    }
    pub fn set_accum(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_accum_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.accum = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.accum = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("accum".to_owned(), val); },
            _ => panic!("Expected record with field accum"),
        }
    }
    pub fn set_add(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_add_mul_one_zero(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.add = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.add = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("add".to_owned(), val); },
            _ => panic!("Expected record with field add"),
        }
    }
    pub fn set_addRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.addRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.addRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("addRecord".to_owned(), val); },
            _ => panic!("Expected record with field addRecord"),
        }
    }
    pub fn set_after(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_after_before(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.after = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.after = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("after".to_owned(), val); },
            _ => panic!("Expected record with field after"),
        }
    }
    pub fn set_alt(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_alt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.alt = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.alt = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("alt".to_owned(), val); },
            _ => panic!("Expected record with field alt"),
        }
    }
    pub fn set_append(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_append(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.append = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.append = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("append".to_owned(), val); },
            _ => panic!("Expected record with field append"),
        }
    }
    pub fn set_appendRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_appendRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.appendRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.appendRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("appendRecord".to_owned(), val); },
            _ => panic!("Expected record with field appendRecord"),
        }
    }
    pub fn set_apply(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_apply(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.apply = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.apply = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("apply".to_owned(), val); },
            _ => panic!("Expected record with field apply"),
        }
    }
    pub fn set_asList(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_asList_asMap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.asList = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.asList = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("asList".to_owned(), val); },
            _ => panic!("Expected record with field asList"),
        }
    }
    pub fn set_asMap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_asList_asMap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.asMap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.asMap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("asMap".to_owned(), val); },
            _ => panic!("Expected record with field asMap"),
        }
    }
    pub fn set_ask(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Comonad0_ask(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ask = Some(val);
            },
            Value::Record_Monad0_ask(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ask = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ask = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ask".to_owned(), val); },
            _ => panic!("Expected record with field ask"),
        }
    }
    pub fn set_b(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::Record_a_b_c(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::Record_a_b_c_d_e(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("b".to_owned(), val); },
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn set_before(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_after_before(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.before = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.before = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("before".to_owned(), val); },
            _ => panic!("Expected record with field before"),
        }
    }
    pub fn set_biapply(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bifunctor0_biapply(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.biapply = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.biapply = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("biapply".to_owned(), val); },
            _ => panic!("Expected record with field biapply"),
        }
    }
    pub fn set_bifoldMap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldMap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldMap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bifoldMap".to_owned(), val); },
            _ => panic!("Expected record with field bifoldMap"),
        }
    }
    pub fn set_bifoldl(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldl = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldl = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bifoldl".to_owned(), val); },
            _ => panic!("Expected record with field bifoldl"),
        }
    }
    pub fn set_bifoldr(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_bifoldMap_bifoldl_bifoldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldr = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bifoldr = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bifoldr".to_owned(), val); },
            _ => panic!("Expected record with field bifoldr"),
        }
    }
    pub fn set_bimap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_bimap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bimap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bimap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bimap".to_owned(), val); },
            _ => panic!("Expected record with field bimap"),
        }
    }
    pub fn set_bind(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_bind(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bind = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bind = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bind".to_owned(), val); },
            _ => panic!("Expected record with field bind"),
        }
    }
    pub fn set_bipure(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Biapply0_bipure(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bipure = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bipure = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bipure".to_owned(), val); },
            _ => panic!("Expected record with field bipure"),
        }
    }
    pub fn set_bisequence(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bisequence = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bisequence = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bisequence".to_owned(), val); },
            _ => panic!("Expected record with field bisequence"),
        }
    }
    pub fn set_bitraverse(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bifoldable1_Bifunctor0_bisequence_bitraverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bitraverse = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bitraverse = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bitraverse".to_owned(), val); },
            _ => panic!("Expected record with field bitraverse"),
        }
    }
    pub fn set_bottom(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ord0_bottom_top(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bottom = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bottom = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bottom".to_owned(), val); },
            _ => panic!("Expected record with field bottom"),
        }
    }
    pub fn set_bottomRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bottomRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.bottomRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("bottomRecord".to_owned(), val); },
            _ => panic!("Expected record with field bottomRecord"),
        }
    }
    pub fn set_c(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_c(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::Record_a_b_c_d_e(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::Record_c_d(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("c".to_owned(), val); },
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn set_callCC(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_callCC(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.callCC = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.callCC = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("callCC".to_owned(), val); },
            _ => panic!("Expected record with field callCC"),
        }
    }
    pub fn set_cardinality(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.cardinality = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.cardinality = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("cardinality".to_owned(), val); },
            _ => panic!("Expected record with field cardinality"),
        }
    }
    pub fn set_catchError(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadThrow0_catchError(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.catchError = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.catchError = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("catchError".to_owned(), val); },
            _ => panic!("Expected record with field catchError"),
        }
    }
    pub fn set_choose(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Divide0_choose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.choose = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.choose = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("choose".to_owned(), val); },
            _ => panic!("Expected record with field choose"),
        }
    }
    pub fn set_closed(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_closed(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.closed = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.closed = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("closed".to_owned(), val); },
            _ => panic!("Expected record with field closed"),
        }
    }
    pub fn set_cmap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_cmap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.cmap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.cmap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("cmap".to_owned(), val); },
            _ => panic!("Expected record with field cmap"),
        }
    }
    pub fn set_collect(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_collect_distribute(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.collect = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.collect = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("collect".to_owned(), val); },
            _ => panic!("Expected record with field collect"),
        }
    }
    pub fn set_compare(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Eq0_compare(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compare = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compare = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("compare".to_owned(), val); },
            _ => panic!("Expected record with field compare"),
        }
    }
    pub fn set_compare1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Eq10_compare1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compare1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compare1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("compare1".to_owned(), val); },
            _ => panic!("Expected record with field compare1"),
        }
    }
    pub fn set_compareRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_EqRecord0_compareRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compareRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compareRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("compareRecord".to_owned(), val); },
            _ => panic!("Expected record with field compareRecord"),
        }
    }
    pub fn set_completed(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_completed_failed_killed(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.completed = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.completed = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("completed".to_owned(), val); },
            _ => panic!("Expected record with field completed"),
        }
    }
    pub fn set_compose(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_compose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compose = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.compose = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("compose".to_owned(), val); },
            _ => panic!("Expected record with field compose"),
        }
    }
    pub fn set_conj(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conj = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conj = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("conj".to_owned(), val); },
            _ => panic!("Expected record with field conj"),
        }
    }
    pub fn set_conjRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conjRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conjRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("conjRecord".to_owned(), val); },
            _ => panic!("Expected record with field conjRecord"),
        }
    }
    pub fn set_conquer(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Divide0_conquer(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conquer = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.conquer = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("conquer".to_owned(), val); },
            _ => panic!("Expected record with field conquer"),
        }
    }
    pub fn set_d(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_c_d_e(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.d = Some(val);
            },
            Value::Record_c_d(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.d = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.d = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("d".to_owned(), val); },
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn set_day(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.day = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.day = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("day".to_owned(), val); },
            _ => panic!("Expected record with field day"),
        }
    }
    pub fn set_defer(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_defer(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.defer = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.defer = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("defer".to_owned(), val); },
            _ => panic!("Expected record with field defer"),
        }
    }
    pub fn set_degree(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.degree = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.degree = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("degree".to_owned(), val); },
            _ => panic!("Expected record with field degree"),
        }
    }
    pub fn set_dimap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dimap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.dimap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.dimap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("dimap".to_owned(), val); },
            _ => panic!("Expected record with field dimap"),
        }
    }
    pub fn set_discard(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_discard(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.discard = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.discard = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("discard".to_owned(), val); },
            _ => panic!("Expected record with field discard"),
        }
    }
    pub fn set_disj(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.disj = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.disj = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("disj".to_owned(), val); },
            _ => panic!("Expected record with field disj"),
        }
    }
    pub fn set_disjRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.disjRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.disjRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("disjRecord".to_owned(), val); },
            _ => panic!("Expected record with field disjRecord"),
        }
    }
    pub fn set_distribute(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_collect_distribute(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.distribute = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.distribute = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("distribute".to_owned(), val); },
            _ => panic!("Expected record with field distribute"),
        }
    }
    pub fn set_div(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.div = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.div = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("div".to_owned(), val); },
            _ => panic!("Expected record with field div"),
        }
    }
    pub fn set_divide(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Contravariant0_divide(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.divide = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.divide = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("divide".to_owned(), val); },
            _ => panic!("Expected record with field divide"),
        }
    }
    pub fn set_dotAll(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.dotAll = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.dotAll = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("dotAll".to_owned(), val); },
            _ => panic!("Expected record with field dotAll"),
        }
    }
    pub fn set_e(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_c_d_e(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.e = Some(val);
            },
            Value::Record_e_f(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.e = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.e = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("e".to_owned(), val); },
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn set_elem(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_elem_pos(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.elem = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.elem = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("elem".to_owned(), val); },
            _ => panic!("Expected record with field elem"),
        }
    }
    pub fn set_empty(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Alt0_empty(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.empty = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.empty = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("empty".to_owned(), val); },
            _ => panic!("Expected record with field empty"),
        }
    }
    pub fn set_eq(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_eq(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eq = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eq = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("eq".to_owned(), val); },
            _ => panic!("Expected record with field eq"),
        }
    }
    pub fn set_eq1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_eq1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eq1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eq1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("eq1".to_owned(), val); },
            _ => panic!("Expected record with field eq1"),
        }
    }
    pub fn set_eqRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_eqRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eqRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.eqRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("eqRecord".to_owned(), val); },
            _ => panic!("Expected record with field eqRecord"),
        }
    }
    pub fn set_extend(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_extend(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.extend = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.extend = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("extend".to_owned(), val); },
            _ => panic!("Expected record with field extend"),
        }
    }
    pub fn set_extract(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Extend0_extract(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.extract = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.extract = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("extract".to_owned(), val); },
            _ => panic!("Expected record with field extract"),
        }
    }
    pub fn set_f(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_e_f(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.f = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.f = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("f".to_owned(), val); },
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn set_failed(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_completed_failed_killed(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.failed = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.failed = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("failed".to_owned(), val); },
            _ => panic!("Expected record with field failed"),
        }
    }
    pub fn set_ff(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ff = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ff = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ff".to_owned(), val); },
            _ => panic!("Expected record with field ff"),
        }
    }
    pub fn set_ffRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ffRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ffRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ffRecord".to_owned(), val); },
            _ => panic!("Expected record with field ffRecord"),
        }
    }
    pub fn set_fiber(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fiber_supervisor(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fiber = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fiber = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("fiber".to_owned(), val); },
            _ => panic!("Expected record with field fiber"),
        }
    }
    pub fn set_first(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_first_second(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.first = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.first = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("first".to_owned(), val); },
            _ => panic!("Expected record with field first"),
        }
    }
    pub fn set_foldMap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_foldMap_foldl_foldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldMap".to_owned(), val); },
            _ => panic!("Expected record with field foldMap"),
        }
    }
    pub fn set_foldMap1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMap1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMap1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldMap1".to_owned(), val); },
            _ => panic!("Expected record with field foldMap1"),
        }
    }
    pub fn set_foldMapWithIndex(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMapWithIndex = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldMapWithIndex = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldMapWithIndex".to_owned(), val); },
            _ => panic!("Expected record with field foldMapWithIndex"),
        }
    }
    pub fn set_foldl(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_foldMap_foldl_foldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldl = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldl = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldl".to_owned(), val); },
            _ => panic!("Expected record with field foldl"),
        }
    }
    pub fn set_foldl1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldl1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldl1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldl1".to_owned(), val); },
            _ => panic!("Expected record with field foldl1"),
        }
    }
    pub fn set_foldlWithIndex(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldlWithIndex = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldlWithIndex = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldlWithIndex".to_owned(), val); },
            _ => panic!("Expected record with field foldlWithIndex"),
        }
    }
    pub fn set_foldr(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_foldMap_foldl_foldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldr = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldr = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldr".to_owned(), val); },
            _ => panic!("Expected record with field foldr"),
        }
    }
    pub fn set_foldr1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMap1_foldl1_foldr1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldr1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldr1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldr1".to_owned(), val); },
            _ => panic!("Expected record with field foldr1"),
        }
    }
    pub fn set_foldrWithIndex(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldrWithIndex = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.foldrWithIndex = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("foldrWithIndex".to_owned(), val); },
            _ => panic!("Expected record with field foldrWithIndex"),
        }
    }
    pub fn set_found(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_found_result(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.found = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.found = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("found".to_owned(), val); },
            _ => panic!("Expected record with field found"),
        }
    }
    pub fn set_from(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_from_to(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.from = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.from = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("from".to_owned(), val); },
            _ => panic!("Expected record with field from"),
        }
    }
    pub fn set_fromDuration(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fromDuration_toDuration(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromDuration = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromDuration = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("fromDuration".to_owned(), val); },
            _ => panic!("Expected record with field fromDuration"),
        }
    }
    pub fn set_fromEnum(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromEnum = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromEnum = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("fromEnum".to_owned(), val); },
            _ => panic!("Expected record with field fromEnum"),
        }
    }
    pub fn set_fromLeft(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromLeft = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromLeft = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("fromLeft".to_owned(), val); },
            _ => panic!("Expected record with field fromLeft"),
        }
    }
    pub fn set_fromRight(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromRight = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.fromRight = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("fromRight".to_owned(), val); },
            _ => panic!("Expected record with field fromRight"),
        }
    }
    pub fn set_genericAdd_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericAdd_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericAdd_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericAdd'".to_owned(), val); },
            _ => panic!("Expected record with field genericAdd_prime"),
        }
    }
    pub fn set_genericAppend_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericAppend_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericAppend_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericAppend_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericAppend'".to_owned(), val); },
            _ => panic!("Expected record with field genericAppend_prime"),
        }
    }
    pub fn set_genericBottom_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericBottom_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericBottom_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericBottom_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericBottom'".to_owned(), val); },
            _ => panic!("Expected record with field genericBottom_prime"),
        }
    }
    pub fn set_genericCardinality_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericCardinality_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericCardinality_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericCardinality'".to_owned(), val); },
            _ => panic!("Expected record with field genericCardinality_prime"),
        }
    }
    pub fn set_genericCompare_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericCompare_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericCompare_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericCompare_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericCompare'".to_owned(), val); },
            _ => panic!("Expected record with field genericCompare_prime"),
        }
    }
    pub fn set_genericConj_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericConj_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericConj_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericConj'".to_owned(), val); },
            _ => panic!("Expected record with field genericConj_prime"),
        }
    }
    pub fn set_genericDisj_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericDisj_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericDisj_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericDisj'".to_owned(), val); },
            _ => panic!("Expected record with field genericDisj_prime"),
        }
    }
    pub fn set_genericEq_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericEq_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericEq_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericEq_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericEq'".to_owned(), val); },
            _ => panic!("Expected record with field genericEq_prime"),
        }
    }
    pub fn set_genericFF_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericFF_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericFF_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericFF'".to_owned(), val); },
            _ => panic!("Expected record with field genericFF_prime"),
        }
    }
    pub fn set_genericFromEnum_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericFromEnum_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericFromEnum_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericFromEnum'".to_owned(), val); },
            _ => panic!("Expected record with field genericFromEnum_prime"),
        }
    }
    pub fn set_genericImplies_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericImplies_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericImplies_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericImplies'".to_owned(), val); },
            _ => panic!("Expected record with field genericImplies_prime"),
        }
    }
    pub fn set_genericMempty_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericMempty_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericMempty_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericMempty_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericMempty'".to_owned(), val); },
            _ => panic!("Expected record with field genericMempty_prime"),
        }
    }
    pub fn set_genericMul_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericMul_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericMul_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericMul'".to_owned(), val); },
            _ => panic!("Expected record with field genericMul_prime"),
        }
    }
    pub fn set_genericNot_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericNot_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericNot_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericNot'".to_owned(), val); },
            _ => panic!("Expected record with field genericNot_prime"),
        }
    }
    pub fn set_genericOne_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericOne_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericOne_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericOne'".to_owned(), val); },
            _ => panic!("Expected record with field genericOne_prime"),
        }
    }
    pub fn set_genericPred_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericPred_prime_genericSucc_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericPred_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericPred_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericPred'".to_owned(), val); },
            _ => panic!("Expected record with field genericPred_prime"),
        }
    }
    pub fn set_genericShow_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericShow_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericShow_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericShow_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericShow'".to_owned(), val); },
            _ => panic!("Expected record with field genericShow_prime"),
        }
    }
    pub fn set_genericShowArgs(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericShowArgs(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericShowArgs = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericShowArgs = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericShowArgs".to_owned(), val); },
            _ => panic!("Expected record with field genericShowArgs"),
        }
    }
    pub fn set_genericSub_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericSub_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericSub_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericSub_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericSub'".to_owned(), val); },
            _ => panic!("Expected record with field genericSub_prime"),
        }
    }
    pub fn set_genericSucc_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericPred_prime_genericSucc_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericSucc_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericSucc_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericSucc'".to_owned(), val); },
            _ => panic!("Expected record with field genericSucc_prime"),
        }
    }
    pub fn set_genericTT_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericTT_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericTT_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericTT'".to_owned(), val); },
            _ => panic!("Expected record with field genericTT_prime"),
        }
    }
    pub fn set_genericToEnum_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericToEnum_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericToEnum_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericToEnum'".to_owned(), val); },
            _ => panic!("Expected record with field genericToEnum_prime"),
        }
    }
    pub fn set_genericTop_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericTop_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericTop_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericTop_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericTop'".to_owned(), val); },
            _ => panic!("Expected record with field genericTop_prime"),
        }
    }
    pub fn set_genericZero_prime(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericZero_prime = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.genericZero_prime = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("genericZero'".to_owned(), val); },
            _ => panic!("Expected record with field genericZero_prime"),
        }
    }
    pub fn set_global(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.global = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.global = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("global".to_owned(), val); },
            _ => panic!("Expected record with field global"),
        }
    }
    pub fn set_handler(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_handler_rethrow(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.handler = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.handler = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("handler".to_owned(), val); },
            _ => panic!("Expected record with field handler"),
        }
    }
    pub fn set_head(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_head_tail(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.head = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.head = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("head".to_owned(), val); },
            _ => panic!("Expected record with field head"),
        }
    }
    pub fn set_hour(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.hour = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.hour = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("hour".to_owned(), val); },
            _ => panic!("Expected record with field hour"),
        }
    }
    pub fn set_identity(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Semigroupoid0_identity(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.identity = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.identity = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("identity".to_owned(), val); },
            _ => panic!("Expected record with field identity"),
        }
    }
    pub fn set_ignoreCase(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ignoreCase = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ignoreCase = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ignoreCase".to_owned(), val); },
            _ => panic!("Expected record with field ignoreCase"),
        }
    }
    pub fn set_imap(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_imap(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.imap = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.imap = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("imap".to_owned(), val); },
            _ => panic!("Expected record with field imap"),
        }
    }
    pub fn set_implies(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.implies = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.implies = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("implies".to_owned(), val); },
            _ => panic!("Expected record with field implies"),
        }
    }
    pub fn set_impliesRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.impliesRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.impliesRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("impliesRecord".to_owned(), val); },
            _ => panic!("Expected record with field impliesRecord"),
        }
    }
    pub fn set_index(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_index_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.index = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.index = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("index".to_owned(), val); },
            _ => panic!("Expected record with field index"),
        }
    }
    pub fn set_init(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_acc_init(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.init = Some(val);
            },
            Value::Record_init_last(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.init = Some(val);
            },
            Value::Record_init_rest(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.init = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.init = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("init".to_owned(), val); },
            _ => panic!("Expected record with field init"),
        }
    }
    pub fn set_inj(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_inj_prj(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.inj = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.inj = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("inj".to_owned(), val); },
            _ => panic!("Expected record with field inj"),
        }
    }
    pub fn set_isLeft(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.isLeft = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.isLeft = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("isLeft".to_owned(), val); },
            _ => panic!("Expected record with field isLeft"),
        }
    }
    pub fn set_isSuspended(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.isSuspended = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.isSuspended = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("isSuspended".to_owned(), val); },
            _ => panic!("Expected record with field isSuspended"),
        }
    }
    pub fn set_join(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.join = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.join = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("join".to_owned(), val); },
            _ => panic!("Expected record with field join"),
        }
    }
    pub fn set_key(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_key_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.key = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.key = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("key".to_owned(), val); },
            _ => panic!("Expected record with field key"),
        }
    }
    pub fn set_keysImpl(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_keysImpl(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.keysImpl = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.keysImpl = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("keysImpl".to_owned(), val); },
            _ => panic!("Expected record with field keysImpl"),
        }
    }
    pub fn set_kill(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.kill = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.kill = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("kill".to_owned(), val); },
            _ => panic!("Expected record with field kill"),
        }
    }
    pub fn set_killed(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_completed_failed_killed(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.killed = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.killed = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("killed".to_owned(), val); },
            _ => panic!("Expected record with field killed"),
        }
    }
    pub fn set_last(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_init_last(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.last = Some(val);
            },
            Value::Record_last_revInit(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.last = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.last = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("last".to_owned(), val); },
            _ => panic!("Expected record with field last"),
        }
    }
    pub fn set_left(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.left = Some(val);
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.left = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.left = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("left".to_owned(), val); },
            _ => panic!("Expected record with field left"),
        }
    }
    pub fn set_lift(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_lift(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lift = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lift = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("lift".to_owned(), val); },
            _ => panic!("Expected record with field lift"),
        }
    }
    pub fn set_liftAff(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadEffect0_liftAff(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftAff = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftAff = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("liftAff".to_owned(), val); },
            _ => panic!("Expected record with field liftAff"),
        }
    }
    pub fn set_liftEffect(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_liftEffect(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftEffect = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftEffect = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("liftEffect".to_owned(), val); },
            _ => panic!("Expected record with field liftEffect"),
        }
    }
    pub fn set_liftST(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_liftST(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftST = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.liftST = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("liftST".to_owned(), val); },
            _ => panic!("Expected record with field liftST"),
        }
    }
    pub fn set_listen(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.listen = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.listen = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("listen".to_owned(), val); },
            _ => panic!("Expected record with field listen"),
        }
    }
    pub fn set_local(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_ComonadAsk0_local(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.local = Some(val);
            },
            Value::Record_MonadAsk0_local(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.local = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.local = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("local".to_owned(), val); },
            _ => panic!("Expected record with field local"),
        }
    }
    pub fn set_lose(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Decide0_Divisible1_lose(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lose = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lose = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("lose".to_owned(), val); },
            _ => panic!("Expected record with field lose"),
        }
    }
    pub fn set_lower(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_lower(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lower = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.lower = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("lower".to_owned(), val); },
            _ => panic!("Expected record with field lower"),
        }
    }
    pub fn set_map(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_map(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.map = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.map = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("map".to_owned(), val); },
            _ => panic!("Expected record with field map"),
        }
    }
    pub fn set_mapWithIndex(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Functor0_mapWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mapWithIndex = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mapWithIndex = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mapWithIndex".to_owned(), val); },
            _ => panic!("Expected record with field mapWithIndex"),
        }
    }
    pub fn set_mappend_(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_mappend__mempty_(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mappend_ = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mappend_ = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mappend_".to_owned(), val); },
            _ => panic!("Expected record with field mappend_"),
        }
    }
    pub fn set_mempty(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Semigroup0_mempty(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mempty = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mempty = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mempty".to_owned(), val); },
            _ => panic!("Expected record with field mempty"),
        }
    }
    pub fn set_memptyRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_SemigroupRecord0_memptyRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.memptyRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.memptyRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("memptyRecord".to_owned(), val); },
            _ => panic!("Expected record with field memptyRecord"),
        }
    }
    pub fn set_mempty_(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_mappend__mempty_(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mempty_ = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mempty_ = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mempty_".to_owned(), val); },
            _ => panic!("Expected record with field mempty_"),
        }
    }
    pub fn set_millisecond(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.millisecond = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.millisecond = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("millisecond".to_owned(), val); },
            _ => panic!("Expected record with field millisecond"),
        }
    }
    pub fn set_minute(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.minute = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.minute = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("minute".to_owned(), val); },
            _ => panic!("Expected record with field minute"),
        }
    }
    pub fn set_mod_kw(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_CommutativeRing0_degree_div_mod_kw(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mod_kw = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mod_kw = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mod".to_owned(), val); },
            _ => panic!("Expected record with field mod_kw"),
        }
    }
    pub fn set_month(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.month = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.month = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("month".to_owned(), val); },
            _ => panic!("Expected record with field month"),
        }
    }
    pub fn set_mul(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_add_mul_one_zero(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mul = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mul = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mul".to_owned(), val); },
            _ => panic!("Expected record with field mul"),
        }
    }
    pub fn set_mulRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mulRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.mulRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("mulRecord".to_owned(), val); },
            _ => panic!("Expected record with field mulRecord"),
        }
    }
    pub fn set_multiline(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.multiline = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.multiline = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("multiline".to_owned(), val); },
            _ => panic!("Expected record with field multiline"),
        }
    }
    pub fn set_myMethod(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_myMethod(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.myMethod = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.myMethod = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("myMethod".to_owned(), val); },
            _ => panic!("Expected record with field myMethod"),
        }
    }
    pub fn set_nes(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_nes(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.nes = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.nes = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("nes".to_owned(), val); },
            _ => panic!("Expected record with field nes"),
        }
    }
    pub fn set_no(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_no_yes(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.no = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.no = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("no".to_owned(), val); },
            _ => panic!("Expected record with field no"),
        }
    }
    pub fn set_not(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.not = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.not = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("not".to_owned(), val); },
            _ => panic!("Expected record with field not"),
        }
    }
    pub fn set_notRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.notRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.notRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("notRecord".to_owned(), val); },
            _ => panic!("Expected record with field notRecord"),
        }
    }
    pub fn set_onComplete(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.onComplete = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.onComplete = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("onComplete".to_owned(), val); },
            _ => panic!("Expected record with field onComplete"),
        }
    }
    pub fn set_one(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_add_mul_one_zero(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.one = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.one = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("one".to_owned(), val); },
            _ => panic!("Expected record with field one"),
        }
    }
    pub fn set_oneRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.oneRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.oneRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("oneRecord".to_owned(), val); },
            _ => panic!("Expected record with field oneRecord"),
        }
    }
    pub fn set_parallel(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.parallel = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.parallel = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("parallel".to_owned(), val); },
            _ => panic!("Expected record with field parallel"),
        }
    }
    pub fn set_pass(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_MonadTell1_Monoid0_listen_pass(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pass = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pass = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("pass".to_owned(), val); },
            _ => panic!("Expected record with field pass"),
        }
    }
    pub fn set_peek(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Comonad0_peek_pos(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.peek = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.peek = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("peek".to_owned(), val); },
            _ => panic!("Expected record with field peek"),
        }
    }
    pub fn set_pos(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Comonad0_peek_pos(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pos = Some(val);
            },
            Value::Record_elem_pos(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pos = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pos = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("pos".to_owned(), val); },
            _ => panic!("Expected record with field pos"),
        }
    }
    pub fn set_pred(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ord0_pred_succ(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pred = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pred = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("pred".to_owned(), val); },
            _ => panic!("Expected record with field pred"),
        }
    }
    pub fn set_prj(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_inj_prj(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.prj = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.prj = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("prj".to_owned(), val); },
            _ => panic!("Expected record with field prj"),
        }
    }
    pub fn set_proof(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Coercible0_proof(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.proof = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.proof = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("proof".to_owned(), val); },
            _ => panic!("Expected record with field proof"),
        }
    }
    pub fn set_ps_minus_rust_minus_test(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_ps_minus_rust_minus_test(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ps_minus_rust_minus_test = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ps_minus_rust_minus_test = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ps-rust-test".to_owned(), val); },
            _ => panic!("Expected record with field ps_minus_rust_minus_test"),
        }
    }
    pub fn set_pure(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_pure(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pure = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.pure = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("pure".to_owned(), val); },
            _ => panic!("Expected record with field pure"),
        }
    }
    pub fn set_recip(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ring0_recip(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.recip = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.recip = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("recip".to_owned(), val); },
            _ => panic!("Expected record with field recip"),
        }
    }
    pub fn set_reflectSymbol(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_reflectSymbol(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.reflectSymbol = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.reflectSymbol = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("reflectSymbol".to_owned(), val); },
            _ => panic!("Expected record with field reflectSymbol"),
        }
    }
    pub fn set_reflectType(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_reflectType(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.reflectType = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.reflectType = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("reflectType".to_owned(), val); },
            _ => panic!("Expected record with field reflectType"),
        }
    }
    pub fn set_rest(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_init_rest(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.rest = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.rest = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("rest".to_owned(), val); },
            _ => panic!("Expected record with field rest"),
        }
    }
    pub fn set_result(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_found_result(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.result = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.result = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("result".to_owned(), val); },
            _ => panic!("Expected record with field result"),
        }
    }
    pub fn set_rethrow(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_handler_rethrow(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.rethrow = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.rethrow = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("rethrow".to_owned(), val); },
            _ => panic!("Expected record with field rethrow"),
        }
    }
    pub fn set_revInit(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_last_revInit(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.revInit = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.revInit = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("revInit".to_owned(), val); },
            _ => panic!("Expected record with field revInit"),
        }
    }
    pub fn set_right(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.right = Some(val);
            },
            Value::Record_fromLeft_fromRight_isLeft_left_right(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.right = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.right = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("right".to_owned(), val); },
            _ => panic!("Expected record with field right"),
        }
    }
    pub fn set_run(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_isSuspended_join_kill_onComplete_run(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.run = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.run = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("run".to_owned(), val); },
            _ => panic!("Expected record with field run"),
        }
    }
    pub fn set_second(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Profunctor0_first_second(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.second = Some(val);
            },
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.second = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.second = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("second".to_owned(), val); },
            _ => panic!("Expected record with field second"),
        }
    }
    pub fn set_sequence(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequence = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequence = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("sequence".to_owned(), val); },
            _ => panic!("Expected record with field sequence"),
        }
    }
    pub fn set_sequence1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequence1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequence1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("sequence1".to_owned(), val); },
            _ => panic!("Expected record with field sequence1"),
        }
    }
    pub fn set_sequential(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Apply0_Apply1_parallel_sequential(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequential = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sequential = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("sequential".to_owned(), val); },
            _ => panic!("Expected record with field sequential"),
        }
    }
    pub fn set_show(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_show(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.show = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.show = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("show".to_owned(), val); },
            _ => panic!("Expected record with field show"),
        }
    }
    pub fn set_showRecordFields(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_showRecordFields(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.showRecordFields = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.showRecordFields = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("showRecordFields".to_owned(), val); },
            _ => panic!("Expected record with field showRecordFields"),
        }
    }
    pub fn set_state(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_state(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.state = Some(val);
            },
            Value::Record_state_val(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.state = Some(val);
            },
            Value::Record_state_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.state = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.state = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("state".to_owned(), val); },
            _ => panic!("Expected record with field state"),
        }
    }
    pub fn set_sticky(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sticky = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sticky = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("sticky".to_owned(), val); },
            _ => panic!("Expected record with field sticky"),
        }
    }
    pub fn set_sub(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Semiring0_sub(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sub = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.sub = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("sub".to_owned(), val); },
            _ => panic!("Expected record with field sub"),
        }
    }
    pub fn set_subRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_SemiringRecord0_subRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.subRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.subRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("subRecord".to_owned(), val); },
            _ => panic!("Expected record with field subRecord"),
        }
    }
    pub fn set_succ(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ord0_pred_succ(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.succ = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.succ = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("succ".to_owned(), val); },
            _ => panic!("Expected record with field succ"),
        }
    }
    pub fn set_supervisor(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fiber_supervisor(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.supervisor = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.supervisor = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("supervisor".to_owned(), val); },
            _ => panic!("Expected record with field supervisor"),
        }
    }
    pub fn set_tail(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_head_tail(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tail = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tail = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("tail".to_owned(), val); },
            _ => panic!("Expected record with field tail"),
        }
    }
    pub fn set_tailRecM(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_tailRecM(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tailRecM = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tailRecM = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("tailRecM".to_owned(), val); },
            _ => panic!("Expected record with field tailRecM"),
        }
    }
    pub fn set_tell(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad1_Semigroup0_tell(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tell = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tell = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("tell".to_owned(), val); },
            _ => panic!("Expected record with field tell"),
        }
    }
    pub fn set_throwError(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Monad0_throwError(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.throwError = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.throwError = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("throwError".to_owned(), val); },
            _ => panic!("Expected record with field throwError"),
        }
    }
    pub fn set_to(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_from_to(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.to = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.to = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("to".to_owned(), val); },
            _ => panic!("Expected record with field to"),
        }
    }
    pub fn set_toDuration(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_fromDuration_toDuration(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.toDuration = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.toDuration = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("toDuration".to_owned(), val); },
            _ => panic!("Expected record with field toDuration"),
        }
    }
    pub fn set_toEnum(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Bounded0_Enum1_cardinality_fromEnum_toEnum(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.toEnum = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.toEnum = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("toEnum".to_owned(), val); },
            _ => panic!("Expected record with field toEnum"),
        }
    }
    pub fn set_top(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Ord0_bottom_top(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.top = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.top = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("top".to_owned(), val); },
            _ => panic!("Expected record with field top"),
        }
    }
    pub fn set_topRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_OrdRecord0_bottomRecord_topRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.topRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.topRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("topRecord".to_owned(), val); },
            _ => panic!("Expected record with field topRecord"),
        }
    }
    pub fn set_track(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Comonad0_track(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.track = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.track = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("track".to_owned(), val); },
            _ => panic!("Expected record with field track"),
        }
    }
    pub fn set_traverse(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable1_Functor0_sequence_traverse(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverse = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverse = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("traverse".to_owned(), val); },
            _ => panic!("Expected record with field traverse"),
        }
    }
    pub fn set_traverse1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Foldable10_Traversable1_sequence1_traverse1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverse1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverse1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("traverse1".to_owned(), val); },
            _ => panic!("Expected record with field traverse1"),
        }
    }
    pub fn set_traverseWithIndex(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverseWithIndex = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.traverseWithIndex = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("traverseWithIndex".to_owned(), val); },
            _ => panic!("Expected record with field traverseWithIndex"),
        }
    }
    pub fn set_tt(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conj_disj_ff_implies_not_tt(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tt = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.tt = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("tt".to_owned(), val); },
            _ => panic!("Expected record with field tt"),
        }
    }
    pub fn set_ttRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ttRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.ttRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("ttRecord".to_owned(), val); },
            _ => panic!("Expected record with field ttRecord"),
        }
    }
    pub fn set_unfoldr(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_Unfoldable10_unfoldr(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unfoldr = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unfoldr = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("unfoldr".to_owned(), val); },
            _ => panic!("Expected record with field unfoldr"),
        }
    }
    pub fn set_unfoldr1(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_unfoldr1(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unfoldr1 = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unfoldr1 = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("unfoldr1".to_owned(), val); },
            _ => panic!("Expected record with field unfoldr1"),
        }
    }
    pub fn set_unicode(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_dotAll_global_ignoreCase_multiline_sticky_unicode(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unicode = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.unicode = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("unicode".to_owned(), val); },
            _ => panic!("Expected record with field unicode"),
        }
    }
    pub fn set_val(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_acc_val(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.val = Some(val);
            },
            Value::Record_state_val(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.val = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.val = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("val".to_owned(), val); },
            _ => panic!("Expected record with field val"),
        }
    }
    pub fn set_value(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_accum_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.value = Some(val);
            },
            Value::Record_index_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.value = Some(val);
            },
            Value::Record_key_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.value = Some(val);
            },
            Value::Record_state_value(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.value = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.value = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("value".to_owned(), val); },
            _ => panic!("Expected record with field value"),
        }
    }
    pub fn set_year(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_day_hour_millisecond_minute_month_second_year(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.year = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.year = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("year".to_owned(), val); },
            _ => panic!("Expected record with field year"),
        }
    }
    pub fn set_yes(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_no_yes(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.yes = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.yes = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("yes".to_owned(), val); },
            _ => panic!("Expected record with field yes"),
        }
    }
    pub fn set_zero(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_add_mul_one_zero(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.zero = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.zero = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("zero".to_owned(), val); },
            _ => panic!("Expected record with field zero"),
        }
    }
    pub fn set_zeroRecord(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_addRecord_mulRecord_oneRecord_zeroRecord(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.zeroRecord = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.zeroRecord = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("zeroRecord".to_owned(), val); },
            _ => panic!("Expected record with field zeroRecord"),
        }
    }
}

pub type UnknownType = Value;


// Internal strings store one Rust scalar per UTF-16 code unit, in code-unit
// order. Scalars >= D800 are shifted past Rust's surrogate hole.
#[inline]
pub fn purust_char_from_code_unit(unit: u16) -> char {
    let value = unit as u32;
    char::from_u32(if value < 0xd800 { value } else { value + 0x800 }).unwrap()
}

#[inline]
pub fn purust_char_to_code_unit(value: char) -> u16 {
    let value = value as u32;
    assert!(value <= 0x107ff, "Expected an encoded UTF-16 code unit");
    (if value < 0xd800 { value } else { value - 0x800 }) as u16
}

pub fn purust_string_from_utf16(units: &[u16]) -> String {
    units.iter().copied().map(purust_char_from_code_unit).collect()
}

pub fn purust_string_to_utf16(value: &str) -> Vec<u16> {
    value.chars().map(purust_char_to_code_unit).collect()
}

pub fn purust_string_from_utf8(value: &str) -> String {
    if value.is_ascii() { return value.to_owned(); }
    value.encode_utf16().map(purust_char_from_code_unit).collect()
}

pub fn purust_string_to_utf8_lossy(value: &str) -> std::string::String {
    if value.is_ascii() { return value.to_owned(); }
    std::string::String::from_utf16_lossy(&purust_string_to_utf16(value))
}

pub mod module_values {
    use std::collections::HashMap;
    use std::panic::{catch_unwind, resume_unwind, AssertUnwindSafe};
    use std::sync::{Condvar, Mutex, OnceLock};
    use std::thread::{self, ThreadId};

    enum State { Empty, Running(ThreadId), Ready, Poisoned }
    struct Wait { owner: ThreadId, cell: usize }
    static WAITS: OnceLock<Mutex<HashMap<ThreadId, Wait>>> = OnceLock::new();

    fn waits() -> &'static Mutex<HashMap<ThreadId, Wait>> {
        WAITS.get_or_init(|| Mutex::new(HashMap::new()))
    }

    struct Waiting { thread: ThreadId }
    impl Waiting {
        // Called with this cell's state locked, so its owner cannot complete
        // between registering the edge and entering the condition-variable wait.
        fn enter(thread: ThreadId, owner: ThreadId, cell: usize) -> Option<Self> {
            let mut graph = waits().lock().unwrap();
            let mut next = owner;
            loop {
                if next == thread { return None; }
                match graph.get(&next) {
                    Some(wait) => next = wait.owner,
                    None => break,
                }
            }
            graph.insert(thread, Wait { owner, cell });
            Some(Self { thread })
        }
    }
    impl Drop for Waiting {
        fn drop(&mut self) { waits().lock().unwrap().remove(&self.thread); }
    }

    pub struct Cell<T> {
        value: OnceLock<T>,
        state: Mutex<State>,
        ready: Condvar,
    }
    impl<T> Cell<T> {
        pub const fn new() -> Self {
            Self { value: OnceLock::new(), state: Mutex::new(State::Empty), ready: Condvar::new() }
        }

        pub fn get_or_init(&self, name: &'static str, init: impl FnOnce() -> T) -> &T {
            if let Some(value) = self.value.get() { return value; }
            let current = thread::current().id();
            let key = self as *const Self as usize;
            let mut state = self.state.lock().unwrap();
            loop {
                match *state {
                    State::Ready => return self.value.get().unwrap(),
                    State::Poisoned => {
                        drop(state);
                        panic!("module value initialization previously failed: {}", name);
                    }
                    State::Running(owner) => {
                        let waiting = match Waiting::enter(current, owner, key) {
                            Some(waiting) => waiting,
                            None => {
                                drop(state);
                                panic!("cyclic module value initialization: {}", name);
                            }
                        };
                        state = self.ready.wait(state).unwrap();
                        drop(waiting);
                    }
                    State::Empty => { *state = State::Running(current); break; }
                }
            }
            drop(state);
            // No state/graph mutex is held while executing user code, cloning
            // its result, or unwinding its destructors. Failure is permanent:
            // callers never silently replay a partially executed initializer.
            let outcome = catch_unwind(AssertUnwindSafe(init));
            let failure = match outcome {
                Ok(value) => {
                    assert!(self.value.set(value).is_ok(), "module value initialized twice");
                    None
                }
                Err(error) => Some(error),
            };
            let mut state = self.state.lock().unwrap();
            *state = if failure.is_some() { State::Poisoned } else { State::Ready };
            // Remove completed dependencies before this owner can start another
            // initialization. Otherwise stale wait edges could report a cycle.
            waits().lock().unwrap().retain(|_, wait| wait.cell != key);
            self.ready.notify_all();
            drop(state);
            if let Some(error) = failure { resume_unwind(error); }
            self.value.get().unwrap()
        }
    }
}

// Own properties retain insertion order. Enumeration puts array indices first,
// as JS Object.keys does; replacing a value never moves its property.
#[derive(Clone, Default)]
pub struct RecordFields(Vec<(String, Value)>);

impl RecordFields {
    pub fn new() -> Self { Self::default() }
    pub fn get(&self, name: &str) -> Option<&Value> {
        self.0.iter().find(|(key, _)| key == name).map(|(_, value)| value)
    }
    pub fn insert(&mut self, name: String, value: Value) -> Option<Value> {
        if let Some((_, old)) = self.0.iter_mut().find(|(key, _)| key == &name) {
            return Some(std::mem::replace(old, value));
        }
        self.0.push((name, value));
        None
    }
    pub fn remove(&mut self, name: &str) -> Option<Value> {
        self.0.iter().position(|(key, _)| key == name).map(|i| self.0.remove(i).1)
    }
    pub fn entries(&self) -> Vec<(String, Value)> {
        let mut entries = self.0.clone();
        entries.sort_by_key(|(key, _)| {
            key.parse::<u32>().ok().filter(|n| *n != u32::MAX && n.to_string() == *key)
                .map(|n| (0, n)).unwrap_or((1, 0))
        });
        entries
    }
}

// Shared mutable own-property storage used by native object FFI. Keeping the
// carrier here lets Foreign readers inspect it without a library dependency cycle.
pub struct SharedRecord(std::sync::Mutex<RecordFields>);
impl SharedRecord {
    pub fn empty() -> Self { Self(std::sync::Mutex::new(RecordFields::new())) }
    pub fn from_entries(entries: Vec<(String, Value)>) -> Self {
        let mut fields = RecordFields::new();
        for (key, value) in entries { fields.insert(key, value); }
        Self(std::sync::Mutex::new(fields))
    }
    pub fn snapshot(&self) -> Self { Self(std::sync::Mutex::new(self.lock().clone())) }
    pub fn get(&self, key: &str) -> Option<Value> { self.lock().get(key).cloned() }
    pub fn entries(&self) -> Vec<(String, Value)> { self.lock().entries() }
    pub fn insert(&self, key: String, value: Value) -> Option<Value> { self.lock().insert(key, value) }
    pub fn remove(&self, key: &str) -> Option<Value> { self.lock().remove(key) }
    fn lock(&self) -> std::sync::MutexGuard<'_, RecordFields> {
        self.0.lock().unwrap_or_else(|poisoned| poisoned.into_inner())
    }
}

impl Value {
    // Immutable PureScript records can enter Foreign.Object through Foreign
    // readers. Keep existing object handles shared; materialize own fields
    // only when crossing from a native immutable record representation.
    pub fn __purust_foreign_object(&self) -> std::rc::Rc<SharedRecord> {
        if let Value::Class(native) = self.resolve() {
            return native.downcast_ref::<std::rc::Rc<SharedRecord>>()
                .expect("Expected a Foreign.Object handle").clone();
        }
        let fields = self.__purust_record_fields().expect("Expected an object or record");
        std::rc::Rc::new(SharedRecord::from_entries(fields.entries()))
    }
}
pub fn mk_unit(_val: ()) -> UnknownType { Value::Unit }
pub fn mk_int(val: i64) -> UnknownType { Value::Int(val) }
pub fn mk_bool(val: bool) -> UnknownType { Value::Bool(val) }
pub fn mk_number(val: f64) -> UnknownType { Value::Number(val) }
pub fn mk_string(val: &str) -> UnknownType { Value::String(val.to_string()) }
pub fn mk_char(val: char) -> UnknownType { Value::Char(val) }
pub fn mk_array(val: Vec<UnknownType>) -> UnknownType { Value::Array(std::rc::Rc::new(val)) }

#[derive(Clone, Default)]
pub struct Thunk {
    pub value: std::sync::OnceLock<Value>,
}

#[derive(Clone, Default)]
pub struct Record_a {
    pub tag: &'static str,
    pub vals: Option<std::rc::Rc<Vec<UnknownType>>>,
    pub call: Option<Func1<UnknownType, UnknownType>>,
    pub Alt0: Option<UnknownType>,
    pub Alternative1: Option<UnknownType>,
    pub Applicative0: Option<UnknownType>,
    pub Apply0: Option<UnknownType>,
    pub Apply1: Option<UnknownType>,
    pub Biapply0: Option<UnknownType>,
    pub Bifoldable1: Option<UnknownType>,
    pub Bifunctor0: Option<UnknownType>,
    pub Bind1: Option<UnknownType>,
    pub Bounded0: Option<UnknownType>,
    pub Coercible0: Option<UnknownType>,
    pub CommutativeRing0: Option<UnknownType>,
    pub Comonad0: Option<UnknownType>,
    pub ComonadAsk0: Option<UnknownType>,
    pub Contravariant0: Option<UnknownType>,
    pub Decide0: Option<UnknownType>,
    pub Divide0: Option<UnknownType>,
    pub Divisible1: Option<UnknownType>,
    pub DivisionRing1: Option<UnknownType>,
    pub Enum1: Option<UnknownType>,
    pub Eq0: Option<UnknownType>,
    pub Eq10: Option<UnknownType>,
    pub EqRecord0: Option<UnknownType>,
    pub EuclideanRing0: Option<UnknownType>,
    pub Extend0: Option<UnknownType>,
    pub Foldable0: Option<UnknownType>,
    pub Foldable1: Option<UnknownType>,
    pub Foldable10: Option<UnknownType>,
    pub FoldableWithIndex1: Option<UnknownType>,
    pub Functor0: Option<UnknownType>,
    pub FunctorWithIndex0: Option<UnknownType>,
    pub HeytingAlgebra0: Option<UnknownType>,
    pub HeytingAlgebraRecord0: Option<UnknownType>,
    pub Monad0: Option<UnknownType>,
    pub Monad1: Option<UnknownType>,
    pub MonadAsk0: Option<UnknownType>,
    pub MonadEffect0: Option<UnknownType>,
    pub MonadTell1: Option<UnknownType>,
    pub MonadThrow0: Option<UnknownType>,
    pub Monoid0: Option<UnknownType>,
    pub Ord0: Option<UnknownType>,
    pub OrdRecord0: Option<UnknownType>,
    pub Plus1: Option<UnknownType>,
    pub Profunctor0: Option<UnknownType>,
    pub Ring0: Option<UnknownType>,
    pub RingRecord0: Option<UnknownType>,
    pub Semigroup0: Option<UnknownType>,
    pub SemigroupRecord0: Option<UnknownType>,
    pub Semigroupoid0: Option<UnknownType>,
    pub Semiring0: Option<UnknownType>,
    pub SemiringRecord0: Option<UnknownType>,
    pub Traversable1: Option<UnknownType>,
    pub Traversable2: Option<UnknownType>,
    pub Unfoldable10: Option<UnknownType>,
    pub a: Option<UnknownType>,
    pub acc: Option<UnknownType>,
    pub accum: Option<UnknownType>,
    pub add: Option<UnknownType>,
    pub addRecord: Option<UnknownType>,
    pub after: Option<UnknownType>,
    pub alt: Option<UnknownType>,
    pub append: Option<UnknownType>,
    pub appendRecord: Option<UnknownType>,
    pub apply: Option<UnknownType>,
    pub asList: Option<UnknownType>,
    pub asMap: Option<UnknownType>,
    pub ask: Option<UnknownType>,
    pub b: Option<UnknownType>,
    pub before: Option<UnknownType>,
    pub biapply: Option<UnknownType>,
    pub bifoldMap: Option<UnknownType>,
    pub bifoldl: Option<UnknownType>,
    pub bifoldr: Option<UnknownType>,
    pub bimap: Option<UnknownType>,
    pub bind: Option<UnknownType>,
    pub bipure: Option<UnknownType>,
    pub bisequence: Option<UnknownType>,
    pub bitraverse: Option<UnknownType>,
    pub bottom: Option<UnknownType>,
    pub bottomRecord: Option<UnknownType>,
    pub c: Option<UnknownType>,
    pub callCC: Option<UnknownType>,
    pub cardinality: Option<UnknownType>,
    pub catchError: Option<UnknownType>,
    pub choose: Option<UnknownType>,
    pub closed: Option<UnknownType>,
    pub cmap: Option<UnknownType>,
    pub collect: Option<UnknownType>,
    pub compare: Option<UnknownType>,
    pub compare1: Option<UnknownType>,
    pub compareRecord: Option<UnknownType>,
    pub completed: Option<UnknownType>,
    pub compose: Option<UnknownType>,
    pub conj: Option<UnknownType>,
    pub conjRecord: Option<UnknownType>,
    pub conquer: Option<UnknownType>,
    pub d: Option<UnknownType>,
    pub day: Option<UnknownType>,
    pub defer: Option<UnknownType>,
    pub degree: Option<UnknownType>,
    pub dimap: Option<UnknownType>,
    pub discard: Option<UnknownType>,
    pub disj: Option<UnknownType>,
    pub disjRecord: Option<UnknownType>,
    pub distribute: Option<UnknownType>,
    pub div: Option<UnknownType>,
    pub divide: Option<UnknownType>,
    pub dotAll: Option<UnknownType>,
    pub e: Option<UnknownType>,
    pub elem: Option<UnknownType>,
    pub empty: Option<UnknownType>,
    pub eq: Option<UnknownType>,
    pub eq1: Option<UnknownType>,
    pub eqRecord: Option<UnknownType>,
    pub extend: Option<UnknownType>,
    pub extract: Option<UnknownType>,
    pub f: Option<UnknownType>,
    pub failed: Option<UnknownType>,
    pub ff: Option<UnknownType>,
    pub ffRecord: Option<UnknownType>,
    pub fiber: Option<UnknownType>,
    pub first: Option<UnknownType>,
    pub foldMap: Option<UnknownType>,
    pub foldMap1: Option<UnknownType>,
    pub foldMapWithIndex: Option<UnknownType>,
    pub foldl: Option<UnknownType>,
    pub foldl1: Option<UnknownType>,
    pub foldlWithIndex: Option<UnknownType>,
    pub foldr: Option<UnknownType>,
    pub foldr1: Option<UnknownType>,
    pub foldrWithIndex: Option<UnknownType>,
    pub found: Option<UnknownType>,
    pub from: Option<UnknownType>,
    pub fromDuration: Option<UnknownType>,
    pub fromEnum: Option<UnknownType>,
    pub fromLeft: Option<UnknownType>,
    pub fromRight: Option<UnknownType>,
    pub genericAdd_prime: Option<UnknownType>,
    pub genericAppend_prime: Option<UnknownType>,
    pub genericBottom_prime: Option<UnknownType>,
    pub genericCardinality_prime: Option<UnknownType>,
    pub genericCompare_prime: Option<UnknownType>,
    pub genericConj_prime: Option<UnknownType>,
    pub genericDisj_prime: Option<UnknownType>,
    pub genericEq_prime: Option<UnknownType>,
    pub genericFF_prime: Option<UnknownType>,
    pub genericFromEnum_prime: Option<UnknownType>,
    pub genericImplies_prime: Option<UnknownType>,
    pub genericMempty_prime: Option<UnknownType>,
    pub genericMul_prime: Option<UnknownType>,
    pub genericNot_prime: Option<UnknownType>,
    pub genericOne_prime: Option<UnknownType>,
    pub genericPred_prime: Option<UnknownType>,
    pub genericShow_prime: Option<UnknownType>,
    pub genericShowArgs: Option<UnknownType>,
    pub genericSub_prime: Option<UnknownType>,
    pub genericSucc_prime: Option<UnknownType>,
    pub genericTT_prime: Option<UnknownType>,
    pub genericToEnum_prime: Option<UnknownType>,
    pub genericTop_prime: Option<UnknownType>,
    pub genericZero_prime: Option<UnknownType>,
    pub global: Option<UnknownType>,
    pub handler: Option<UnknownType>,
    pub head: Option<UnknownType>,
    pub hour: Option<UnknownType>,
    pub identity: Option<UnknownType>,
    pub ignoreCase: Option<UnknownType>,
    pub imap: Option<UnknownType>,
    pub implies: Option<UnknownType>,
    pub impliesRecord: Option<UnknownType>,
    pub index: Option<UnknownType>,
    pub init: Option<UnknownType>,
    pub inj: Option<UnknownType>,
    pub isLeft: Option<UnknownType>,
    pub isSuspended: Option<UnknownType>,
    pub join: Option<UnknownType>,
    pub key: Option<UnknownType>,
    pub keysImpl: Option<UnknownType>,
    pub kill: Option<UnknownType>,
    pub killed: Option<UnknownType>,
    pub last: Option<UnknownType>,
    pub left: Option<UnknownType>,
    pub lift: Option<UnknownType>,
    pub liftAff: Option<UnknownType>,
    pub liftEffect: Option<UnknownType>,
    pub liftST: Option<UnknownType>,
    pub listen: Option<UnknownType>,
    pub local: Option<UnknownType>,
    pub lose: Option<UnknownType>,
    pub lower: Option<UnknownType>,
    pub map: Option<UnknownType>,
    pub mapWithIndex: Option<UnknownType>,
    pub mappend_: Option<UnknownType>,
    pub mempty: Option<UnknownType>,
    pub memptyRecord: Option<UnknownType>,
    pub mempty_: Option<UnknownType>,
    pub millisecond: Option<UnknownType>,
    pub minute: Option<UnknownType>,
    pub mod_kw: Option<UnknownType>,
    pub month: Option<UnknownType>,
    pub mul: Option<UnknownType>,
    pub mulRecord: Option<UnknownType>,
    pub multiline: Option<UnknownType>,
    pub myMethod: Option<UnknownType>,
    pub nes: Option<UnknownType>,
    pub no: Option<UnknownType>,
    pub not: Option<UnknownType>,
    pub notRecord: Option<UnknownType>,
    pub onComplete: Option<UnknownType>,
    pub one: Option<UnknownType>,
    pub oneRecord: Option<UnknownType>,
    pub parallel: Option<UnknownType>,
    pub pass: Option<UnknownType>,
    pub peek: Option<UnknownType>,
    pub pos: Option<UnknownType>,
    pub pred: Option<UnknownType>,
    pub prj: Option<UnknownType>,
    pub proof: Option<UnknownType>,
    pub ps_minus_rust_minus_test: Option<UnknownType>,
    pub pure: Option<UnknownType>,
    pub recip: Option<UnknownType>,
    pub reflectSymbol: Option<UnknownType>,
    pub reflectType: Option<UnknownType>,
    pub rest: Option<UnknownType>,
    pub result: Option<UnknownType>,
    pub rethrow: Option<UnknownType>,
    pub revInit: Option<UnknownType>,
    pub right: Option<UnknownType>,
    pub run: Option<UnknownType>,
    pub second: Option<UnknownType>,
    pub sequence: Option<UnknownType>,
    pub sequence1: Option<UnknownType>,
    pub sequential: Option<UnknownType>,
    pub show: Option<UnknownType>,
    pub showRecordFields: Option<UnknownType>,
    pub state: Option<UnknownType>,
    pub sticky: Option<UnknownType>,
    pub sub: Option<UnknownType>,
    pub subRecord: Option<UnknownType>,
    pub succ: Option<UnknownType>,
    pub supervisor: Option<UnknownType>,
    pub tail: Option<UnknownType>,
    pub tailRecM: Option<UnknownType>,
    pub tell: Option<UnknownType>,
    pub throwError: Option<UnknownType>,
    pub to: Option<UnknownType>,
    pub toDuration: Option<UnknownType>,
    pub toEnum: Option<UnknownType>,
    pub top: Option<UnknownType>,
    pub topRecord: Option<UnknownType>,
    pub track: Option<UnknownType>,
    pub traverse: Option<UnknownType>,
    pub traverse1: Option<UnknownType>,
    pub traverseWithIndex: Option<UnknownType>,
    pub tt: Option<UnknownType>,
    pub ttRecord: Option<UnknownType>,
    pub unfoldr: Option<UnknownType>,
    pub unfoldr1: Option<UnknownType>,
    pub unicode: Option<UnknownType>,
    pub val: Option<UnknownType>,
    pub value: Option<UnknownType>,
    pub year: Option<UnknownType>,
    pub yes: Option<UnknownType>,
    pub zero: Option<UnknownType>,
    pub zeroRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Alt0_empty {
    pub Alt0: Option<UnknownType>,
    pub empty: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Alternative1_Monad0 {
    pub Alternative1: Option<UnknownType>,
    pub Monad0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Applicative0_Bind1 {
    pub Applicative0: Option<UnknownType>,
    pub Bind1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Applicative0_Plus1 {
    pub Applicative0: Option<UnknownType>,
    pub Plus1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Apply0_Apply1_parallel_sequential {
    pub Apply0: Option<UnknownType>,
    pub Apply1: Option<UnknownType>,
    pub parallel: Option<UnknownType>,
    pub sequential: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Apply0_bind {
    pub Apply0: Option<UnknownType>,
    pub bind: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Apply0_pure {
    pub Apply0: Option<UnknownType>,
    pub pure: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Biapply0_bipure {
    pub Biapply0: Option<UnknownType>,
    pub bipure: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Bifoldable1_Bifunctor0_bisequence_bitraverse {
    pub Bifoldable1: Option<UnknownType>,
    pub Bifunctor0: Option<UnknownType>,
    pub bisequence: Option<UnknownType>,
    pub bitraverse: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Bifunctor0_biapply {
    pub Bifunctor0: Option<UnknownType>,
    pub biapply: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Bounded0_Enum1_cardinality_fromEnum_toEnum {
    pub Bounded0: Option<UnknownType>,
    pub Enum1: Option<UnknownType>,
    pub cardinality: Option<UnknownType>,
    pub fromEnum: Option<UnknownType>,
    pub toEnum: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Coercible0 {
    pub Coercible0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Coercible0_proof {
    pub Coercible0: Option<UnknownType>,
    pub proof: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_CommutativeRing0_degree_div_mod_kw {
    pub CommutativeRing0: Option<UnknownType>,
    pub degree: Option<UnknownType>,
    pub div: Option<UnknownType>,
    pub mod_kw: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Comonad0_ask {
    pub Comonad0: Option<UnknownType>,
    pub ask: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Comonad0_peek_pos {
    pub Comonad0: Option<UnknownType>,
    pub peek: Option<UnknownType>,
    pub pos: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Comonad0_track {
    pub Comonad0: Option<UnknownType>,
    pub track: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_ComonadAsk0_local {
    pub ComonadAsk0: Option<UnknownType>,
    pub local: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Contravariant0_divide {
    pub Contravariant0: Option<UnknownType>,
    pub divide: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Decide0_Divisible1_lose {
    pub Decide0: Option<UnknownType>,
    pub Divisible1: Option<UnknownType>,
    pub lose: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Divide0_choose {
    pub Divide0: Option<UnknownType>,
    pub choose: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Divide0_conquer {
    pub Divide0: Option<UnknownType>,
    pub conquer: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_DivisionRing1_EuclideanRing0 {
    pub DivisionRing1: Option<UnknownType>,
    pub EuclideanRing0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Eq0_compare {
    pub Eq0: Option<UnknownType>,
    pub compare: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Eq10_compare1 {
    pub Eq10: Option<UnknownType>,
    pub compare1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_EqRecord0_compareRecord {
    pub EqRecord0: Option<UnknownType>,
    pub compareRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Extend0_extract {
    pub Extend0: Option<UnknownType>,
    pub extract: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Foldable0_foldMap1_foldl1_foldr1 {
    pub Foldable0: Option<UnknownType>,
    pub foldMap1: Option<UnknownType>,
    pub foldl1: Option<UnknownType>,
    pub foldr1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Foldable0_foldMapWithIndex_foldlWithIndex_foldrWithIndex {
    pub Foldable0: Option<UnknownType>,
    pub foldMapWithIndex: Option<UnknownType>,
    pub foldlWithIndex: Option<UnknownType>,
    pub foldrWithIndex: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Foldable1_Functor0_sequence_traverse {
    pub Foldable1: Option<UnknownType>,
    pub Functor0: Option<UnknownType>,
    pub sequence: Option<UnknownType>,
    pub traverse: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Foldable10_Traversable1_sequence1_traverse1 {
    pub Foldable10: Option<UnknownType>,
    pub Traversable1: Option<UnknownType>,
    pub sequence1: Option<UnknownType>,
    pub traverse1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_FoldableWithIndex1_FunctorWithIndex0_Traversable2_traverseWithIndex {
    pub FoldableWithIndex1: Option<UnknownType>,
    pub FunctorWithIndex0: Option<UnknownType>,
    pub Traversable2: Option<UnknownType>,
    pub traverseWithIndex: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Functor0_alt {
    pub Functor0: Option<UnknownType>,
    pub alt: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Functor0_apply {
    pub Functor0: Option<UnknownType>,
    pub apply: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Functor0_collect_distribute {
    pub Functor0: Option<UnknownType>,
    pub collect: Option<UnknownType>,
    pub distribute: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Functor0_extend {
    pub Functor0: Option<UnknownType>,
    pub extend: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Functor0_mapWithIndex {
    pub Functor0: Option<UnknownType>,
    pub mapWithIndex: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_HeytingAlgebra0 {
    pub HeytingAlgebra0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_HeytingAlgebraRecord0 {
    pub HeytingAlgebraRecord0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_ask {
    pub Monad0: Option<UnknownType>,
    pub ask: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_callCC {
    pub Monad0: Option<UnknownType>,
    pub callCC: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_liftEffect {
    pub Monad0: Option<UnknownType>,
    pub liftEffect: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_liftST {
    pub Monad0: Option<UnknownType>,
    pub liftST: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_state {
    pub Monad0: Option<UnknownType>,
    pub state: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_tailRecM {
    pub Monad0: Option<UnknownType>,
    pub tailRecM: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad0_throwError {
    pub Monad0: Option<UnknownType>,
    pub throwError: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Monad1_Semigroup0_tell {
    pub Monad1: Option<UnknownType>,
    pub Semigroup0: Option<UnknownType>,
    pub tell: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_MonadAsk0_local {
    pub MonadAsk0: Option<UnknownType>,
    pub local: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_MonadEffect0_liftAff {
    pub MonadEffect0: Option<UnknownType>,
    pub liftAff: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_MonadTell1_Monoid0_listen_pass {
    pub MonadTell1: Option<UnknownType>,
    pub Monoid0: Option<UnknownType>,
    pub listen: Option<UnknownType>,
    pub pass: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_MonadThrow0_catchError {
    pub MonadThrow0: Option<UnknownType>,
    pub catchError: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Ord0_bottom_top {
    pub Ord0: Option<UnknownType>,
    pub bottom: Option<UnknownType>,
    pub top: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Ord0_pred_succ {
    pub Ord0: Option<UnknownType>,
    pub pred: Option<UnknownType>,
    pub succ: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_OrdRecord0_bottomRecord_topRecord {
    pub OrdRecord0: Option<UnknownType>,
    pub bottomRecord: Option<UnknownType>,
    pub topRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Profunctor0_closed {
    pub Profunctor0: Option<UnknownType>,
    pub closed: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Profunctor0_first_second {
    pub Profunctor0: Option<UnknownType>,
    pub first: Option<UnknownType>,
    pub second: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Profunctor0_left_right {
    pub Profunctor0: Option<UnknownType>,
    pub left: Option<UnknownType>,
    pub right: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Ring0 {
    pub Ring0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Ring0_recip {
    pub Ring0: Option<UnknownType>,
    pub recip: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_RingRecord0 {
    pub RingRecord0: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Semigroup0_mempty {
    pub Semigroup0: Option<UnknownType>,
    pub mempty: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_SemigroupRecord0_memptyRecord {
    pub SemigroupRecord0: Option<UnknownType>,
    pub memptyRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Semigroupoid0_identity {
    pub Semigroupoid0: Option<UnknownType>,
    pub identity: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Semiring0_sub {
    pub Semiring0: Option<UnknownType>,
    pub sub: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_SemiringRecord0_subRecord {
    pub SemiringRecord0: Option<UnknownType>,
    pub subRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_Unfoldable10_unfoldr {
    pub Unfoldable10: Option<UnknownType>,
    pub unfoldr: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_a_b {
    pub a: Option<UnknownType>,
    pub b: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_a_b_c {
    pub a: Option<UnknownType>,
    pub b: Option<UnknownType>,
    pub c: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_a_b_c_d_e {
    pub a: Option<UnknownType>,
    pub b: Option<UnknownType>,
    pub c: Option<UnknownType>,
    pub d: Option<UnknownType>,
    pub e: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_acc_init {
    pub acc: Option<UnknownType>,
    pub init: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_acc_val {
    pub acc: Option<UnknownType>,
    pub val: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_accum_value {
    pub accum: Option<UnknownType>,
    pub value: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_add_mul_one_zero {
    pub add: Option<UnknownType>,
    pub mul: Option<UnknownType>,
    pub one: Option<UnknownType>,
    pub zero: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_addRecord_mulRecord_oneRecord_zeroRecord {
    pub addRecord: Option<UnknownType>,
    pub mulRecord: Option<UnknownType>,
    pub oneRecord: Option<UnknownType>,
    pub zeroRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_after_before {
    pub after: Option<UnknownType>,
    pub before: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_append {
    pub append: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_appendRecord {
    pub appendRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_asList_asMap {
    pub asList: Option<UnknownType>,
    pub asMap: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_bifoldMap_bifoldl_bifoldr {
    pub bifoldMap: Option<UnknownType>,
    pub bifoldl: Option<UnknownType>,
    pub bifoldr: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_bimap {
    pub bimap: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_c_d {
    pub c: Option<UnknownType>,
    pub d: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_cmap {
    pub cmap: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_completed_failed_killed {
    pub completed: Option<UnknownType>,
    pub failed: Option<UnknownType>,
    pub killed: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_compose {
    pub compose: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_conj_disj_ff_implies_not_tt {
    pub conj: Option<UnknownType>,
    pub disj: Option<UnknownType>,
    pub ff: Option<UnknownType>,
    pub implies: Option<UnknownType>,
    pub not: Option<UnknownType>,
    pub tt: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_conjRecord_disjRecord_ffRecord_impliesRecord_notRecord_ttRecord {
    pub conjRecord: Option<UnknownType>,
    pub disjRecord: Option<UnknownType>,
    pub ffRecord: Option<UnknownType>,
    pub impliesRecord: Option<UnknownType>,
    pub notRecord: Option<UnknownType>,
    pub ttRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_day_hour_millisecond_minute_month_second_year {
    pub day: Option<UnknownType>,
    pub hour: Option<UnknownType>,
    pub millisecond: Option<UnknownType>,
    pub minute: Option<UnknownType>,
    pub month: Option<UnknownType>,
    pub second: Option<UnknownType>,
    pub year: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_defer {
    pub defer: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_dimap {
    pub dimap: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_discard {
    pub discard: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_dotAll_global_ignoreCase_multiline_sticky_unicode {
    pub dotAll: Option<UnknownType>,
    pub global: Option<UnknownType>,
    pub ignoreCase: Option<UnknownType>,
    pub multiline: Option<UnknownType>,
    pub sticky: Option<UnknownType>,
    pub unicode: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_e_f {
    pub e: Option<UnknownType>,
    pub f: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_elem_pos {
    pub elem: Option<UnknownType>,
    pub pos: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_eq {
    pub eq: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_eq1 {
    pub eq1: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_eqRecord {
    pub eqRecord: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_fiber_supervisor {
    pub fiber: Option<UnknownType>,
    pub supervisor: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_foldMap_foldl_foldr {
    pub foldMap: Option<UnknownType>,
    pub foldl: Option<UnknownType>,
    pub foldr: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_found_result {
    pub found: Option<UnknownType>,
    pub result: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_from_to {
    pub from: Option<UnknownType>,
    pub to: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_fromDuration_toDuration {
    pub fromDuration: Option<UnknownType>,
    pub toDuration: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_fromLeft_fromRight_isLeft_left_right {
    pub fromLeft: Option<UnknownType>,
    pub fromRight: Option<UnknownType>,
    pub isLeft: Option<UnknownType>,
    pub left: Option<UnknownType>,
    pub right: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericAdd_prime_genericMul_prime_genericOne_prime_genericZero_prime {
    pub genericAdd_prime: Option<UnknownType>,
    pub genericMul_prime: Option<UnknownType>,
    pub genericOne_prime: Option<UnknownType>,
    pub genericZero_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericAppend_prime {
    pub genericAppend_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericBottom_prime {
    pub genericBottom_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericCardinality_prime_genericFromEnum_prime_genericToEnum_prime {
    pub genericCardinality_prime: Option<UnknownType>,
    pub genericFromEnum_prime: Option<UnknownType>,
    pub genericToEnum_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericCompare_prime {
    pub genericCompare_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericConj_prime_genericDisj_prime_genericFF_prime_genericImplies_prime_genericNot_prime_genericTT_prime {
    pub genericConj_prime: Option<UnknownType>,
    pub genericDisj_prime: Option<UnknownType>,
    pub genericFF_prime: Option<UnknownType>,
    pub genericImplies_prime: Option<UnknownType>,
    pub genericNot_prime: Option<UnknownType>,
    pub genericTT_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericEq_prime {
    pub genericEq_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericMempty_prime {
    pub genericMempty_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericPred_prime_genericSucc_prime {
    pub genericPred_prime: Option<UnknownType>,
    pub genericSucc_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericShow_prime {
    pub genericShow_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericShowArgs {
    pub genericShowArgs: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericSub_prime {
    pub genericSub_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_genericTop_prime {
    pub genericTop_prime: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_handler_rethrow {
    pub handler: Option<UnknownType>,
    pub rethrow: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_head_tail {
    pub head: Option<UnknownType>,
    pub tail: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_imap {
    pub imap: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_index_value {
    pub index: Option<UnknownType>,
    pub value: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_init_last {
    pub init: Option<UnknownType>,
    pub last: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_init_rest {
    pub init: Option<UnknownType>,
    pub rest: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_inj_prj {
    pub inj: Option<UnknownType>,
    pub prj: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_isSuspended_join_kill_onComplete_run {
    pub isSuspended: Option<UnknownType>,
    pub join: Option<UnknownType>,
    pub kill: Option<UnknownType>,
    pub onComplete: Option<UnknownType>,
    pub run: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_key_value {
    pub key: Option<UnknownType>,
    pub value: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_keysImpl {
    pub keysImpl: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_last_revInit {
    pub last: Option<UnknownType>,
    pub revInit: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_lift {
    pub lift: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_lower {
    pub lower: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_map {
    pub map: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_mappend__mempty_ {
    pub mappend_: Option<UnknownType>,
    pub mempty_: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_myMethod {
    pub myMethod: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_nes {
    pub nes: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_no_yes {
    pub no: Option<UnknownType>,
    pub yes: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_ps_minus_rust_minus_test {
    pub ps_minus_rust_minus_test: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_reflectSymbol {
    pub reflectSymbol: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_reflectType {
    pub reflectType: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_show {
    pub show: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_showRecordFields {
    pub showRecordFields: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_state_val {
    pub state: Option<UnknownType>,
    pub val: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_state_value {
    pub state: Option<UnknownType>,
    pub value: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_unfoldr1 {
    pub unfoldr1: Option<UnknownType>,
}



#[derive(Clone)]
pub enum Func1<T1, R> {
    Static(fn(T1) -> R),
    Shared(std::rc::Rc<dyn Fn(T1) -> R>),
}

impl<T1: 'static, R: 'static> std::ops::Deref for Func1<T1, R> {
    type Target = dyn Fn(T1) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func1::Static(f) => f,
            Func1::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func2<T1, T2, R> {
    Static(fn(T1, T2) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2) -> R>),
}

impl<T1: 'static, T2: 'static, R: 'static> std::ops::Deref for Func2<T1, T2, R> {
    type Target = dyn Fn(T1, T2) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func2::Static(f) => f,
            Func2::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func3<T1, T2, T3, R> {
    Static(fn(T1, T2, T3) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, R: 'static> std::ops::Deref for Func3<T1, T2, T3, R> {
    type Target = dyn Fn(T1, T2, T3) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func3::Static(f) => f,
            Func3::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func4<T1, T2, T3, T4, R> {
    Static(fn(T1, T2, T3, T4) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, R: 'static> std::ops::Deref for Func4<T1, T2, T3, T4, R> {
    type Target = dyn Fn(T1, T2, T3, T4) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func4::Static(f) => f,
            Func4::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func5<T1, T2, T3, T4, T5, R> {
    Static(fn(T1, T2, T3, T4, T5) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, R: 'static> std::ops::Deref for Func5<T1, T2, T3, T4, T5, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func5::Static(f) => f,
            Func5::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func6<T1, T2, T3, T4, T5, T6, R> {
    Static(fn(T1, T2, T3, T4, T5, T6) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, R: 'static> std::ops::Deref for Func6<T1, T2, T3, T4, T5, T6, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func6::Static(f) => f,
            Func6::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func7<T1, T2, T3, T4, T5, T6, T7, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, R: 'static> std::ops::Deref for Func7<T1, T2, T3, T4, T5, T6, T7, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func7::Static(f) => f,
            Func7::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func8<T1, T2, T3, T4, T5, T6, T7, T8, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, R: 'static> std::ops::Deref for Func8<T1, T2, T3, T4, T5, T6, T7, T8, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func8::Static(f) => f,
            Func8::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func9<T1, T2, T3, T4, T5, T6, T7, T8, T9, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, R: 'static> std::ops::Deref for Func9<T1, T2, T3, T4, T5, T6, T7, T8, T9, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func9::Static(f) => f,
            Func9::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func10<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, R: 'static> std::ops::Deref for Func10<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func10::Static(f) => f,
            Func10::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func11<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, T11: 'static, R: 'static> std::ops::Deref for Func11<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func11::Static(f) => f,
            Func11::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func12<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, T11: 'static, T12: 'static, R: 'static> std::ops::Deref for Func12<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func12::Static(f) => f,
            Func12::Shared(rc) => rc.as_ref(),
        }
    }
}


pub mod microtasks {
// Promise reactions run at a checkpoint, never on the resolving worker's stack.
// Synchronous Aff turns may run concurrently; a checkpoint waits until they end.
use std::rc::Rc;
use std::collections::VecDeque;
use std::sync::{Condvar, Mutex};

type Job = Box<dyn FnOnce() + 'static>;
struct State {
    jobs: VecDeque<Job>,
    checks: VecDeque<Job>,
    turns: usize,
    draining: bool,
}
pub struct Queue {
    state: Mutex<State>,
    idle: Condvar,
    wake: Box<dyn Fn() + 'static>,
}
thread_local! {
    static CURRENT: std::cell::RefCell<Option<Rc<Queue>>> = const { std::cell::RefCell::new(None) };
    static TURN: std::cell::RefCell<Option<Rc<Queue>>> = const { std::cell::RefCell::new(None) };
}
pub struct Scope(Option<Rc<Queue>>);
impl Scope {
    pub fn enter(queue: Rc<Queue>) -> Self {
        Self(CURRENT.with(|current| current.replace(Some(queue))))
    }
}
impl Drop for Scope {
    fn drop(&mut self) { CURRENT.with(|current| current.replace(self.0.take())); }
}
pub fn current() -> Rc<Queue> {
    CURRENT.with(|current| current.borrow().clone()).expect("Promise requires a purust microtask scope")
}
struct Turn {
    queue: Rc<Queue>,
    previous: Option<Rc<Queue>>,
    checkpoint: bool,
}
impl Drop for Turn {
    fn drop(&mut self) {
        TURN.with(|current| current.replace(self.previous.take()));
        let mut state = self.queue.state.lock().unwrap();
        if self.checkpoint { state.draining = false; } else { state.turns -= 1; }
        let wake = (!state.jobs.is_empty() || !state.checks.is_empty()) && state.turns == 0;
        drop(state);
        self.queue.idle.notify_all();
        if wake { (self.queue.wake)(); }
    }
}
impl Queue {
    pub fn new(wake: impl Fn() + 'static) -> Rc<Self> {
        Rc::new(Self { state: Mutex::new(State { jobs: VecDeque::new(), checks: VecDeque::new(), turns: 0, draining: false }),
            idle: Condvar::new(), wake: Box::new(wake) })
    }
    pub fn enqueue(&self, job: impl FnOnce() + 'static) {
        self.state.lock().unwrap().jobs.push_back(Box::new(job));
        (self.wake)();
    }
    pub fn after_checkpoint(&self, check: impl FnOnce() + 'static) {
        self.state.lock().unwrap().checks.push_back(Box::new(check));
        (self.wake)();
    }
    pub fn turn<R>(self: &Rc<Self>, action: impl FnOnce() -> R) -> R {
        let _scope = Scope::enter(self.clone());
        if TURN.with(|current| current.borrow().as_ref().is_some_and(|queue| Rc::ptr_eq(queue, self))) {
            return action();
        }
        let mut state = self.state.lock().unwrap();
        while state.draining { state = self.idle.wait(state).unwrap(); }
        state.turns += 1;
        drop(state);
        let _turn = Turn { queue: self.clone(), checkpoint: false,
            previous: TURN.with(|current| current.replace(Some(self.clone()))) };
        action()
    }
    // A busy turn wakes the executor on exit. Do not block the executor here.
    pub fn drain(self: &Rc<Self>) {
        let mut state = self.state.lock().unwrap();
        if state.draining || state.turns != 0 || (state.jobs.is_empty() && state.checks.is_empty()) { return; }
        state.draining = true;
        drop(state);
        let _scope = Scope::enter(self.clone());
        let _turn = Turn { queue: self.clone(), checkpoint: true,
            previous: TURN.with(|current| current.replace(Some(self.clone()))) };
        loop {
            let job = {
                let mut state = self.state.lock().unwrap();
                state.jobs.pop_front().or_else(|| state.checks.pop_front())
            };
            match job { Some(job) => job(), None => break }
        }
    }
    pub fn has_jobs(&self) -> bool {
        let state = self.state.lock().unwrap();
        !state.jobs.is_empty() || !state.checks.is_empty()
    }
}
pub fn run_main<R>(main: impl FnOnce() -> R) -> R {
    let queue = Queue::new(|| {});
    let result = queue.turn(main);
    queue.drain();
    result
}

}
