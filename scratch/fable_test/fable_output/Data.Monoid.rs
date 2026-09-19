pub mod PureScript_Data_Monoid {
    use super::*;
    use fable_library_rust::Interfaces_::System::IComparable;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::NativeArray_::new_empty;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Monoid_semigroupRecord() -> &dyn Any {
        static Data_Monoid_semigroupRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_semigroupRecord.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupRecord(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Monoid_MonoidRecordusd_Dict() -> &dyn Any {
        static Data_Monoid_MonoidRecordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_MonoidRecordusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Monoid_Monoidusd_Dict() -> &dyn Any {
        static Data_Monoid_Monoidusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Monoidusd_Dict.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Monoid_monoidUnit() -> &dyn Any {
        static Data_Monoid_monoidUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidUnit.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                &&&add(string("mempty"),
                                                                                       &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                       add(string("Semigroup0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Semigroup::Data_Semigroup_semigroupUnit()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Monoid_monoidString() -> &dyn Any {
        static Data_Monoid_monoidString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidString.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                  &&&add(string("mempty"),
                                                                                         &&string(""),
                                                                                         add(string("Semigroup0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Monoid_monoidRecordNil() -> &dyn Any {
        static Data_Monoid_monoidRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidRecordNil.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_MonoidRecordusd_Dict(),
                                                                                     &&&add(string("memptyRecord"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             &empty::<LrcPtr<dyn IComparable>,
                                                                                                                      &dyn Any>()),
                                                                                            add(string("SemigroupRecord0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Semigroup::Data_Semigroup_semigroupRecordNil()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Monoid_monoidOrdering() -> &dyn Any {
        static Data_Monoid_monoidOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidOrdering.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                    &&&add(string("mempty"),
                                                                                           &&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor),
                                                                                           add(string("Semigroup0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Ordering::Data_Ordering_semigroupOrdering()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_monoidArray() -> &dyn Any {
        static Data_Monoid_monoidArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidArray.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                 &&&add(string("mempty"),
                                                                                        &&new_empty::<&dyn Any>(),
                                                                                        add(string("Semigroup0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Monoid_memptyRecord() -> &dyn Any {
        static Data_Monoid_memptyRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_memptyRecord.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("memptyRecord"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Monoid_monoidRecord() -> &dyn Any {
        static Data_Monoid_monoidRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidRecord.get_or_init(||
                                                 &Func1::new(move
                                                                 |usd__unused|
                                                                 &Func1::new(move
                                                                                 |dictMonoidRecord|
                                                                                 {
                                                                                     let semigroupRecord1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_semigroupRecord(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("SemigroupRecord0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictMonoidRecord)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                      &&&add(string("mempty"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_memptyRecord(),
                                                                                                                                                                                                  dictMonoidRecord),
                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                             add(string("Semigroup0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let semigroupRecord1
                                                                                                                                                      =
                                                                                                                                                      semigroupRecord1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused_1|
                                                                                                                                                      &semigroupRecord1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 })))
    }
    pub fn Data_Monoid_mempty() -> &dyn Any {
        static Data_Monoid_mempty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_mempty.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("mempty"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Monoid_monoidFn() -> &dyn Any {
        static Data_Monoid_monoidFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidFn.get_or_init(||
                                             &Func1::new(move |dictMonoid|
                                                             {
                                                                 let semigroupFn =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupFn(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                  &&&add(string("mempty"),
                                                                                                         &&Func1::new({
                                                                                                                          let dictMonoid
                                                                                                                              =
                                                                                                                              dictMonoid.clone();
                                                                                                                          move
                                                                                                                              |v|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                               &&&dictMonoid)
                                                                                                                      }),
                                                                                                         add(string("Semigroup0"),
                                                                                                             &&Func1::new({
                                                                                                                              let semigroupFn
                                                                                                                                  =
                                                                                                                                  semigroupFn.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &semigroupFn
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Monoid_monoidRecordCons() -> &dyn Any {
        static Data_Monoid_monoidRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_monoidRecordCons.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictIsSymbol|
                                                                     {
                                                                         let semigroupRecordCons =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupRecordCons(),
                                                                                                                                                 dictIsSymbol),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         &Func1::new({
                                                                                         let dictIsSymbol
                                                                                             =
                                                                                             dictIsSymbol.clone();
                                                                                         let semigroupRecordCons
                                                                                             =
                                                                                             semigroupRecordCons.clone();
                                                                                         move
                                                                                             |dictMonoid|
                                                                                             {
                                                                                                 let mempty1 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                      dictMonoid);
                                                                                                 let Semigroup0 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                 &Func1::new({
                                                                                                                 let Semigroup0
                                                                                                                     =
                                                                                                                     Semigroup0.clone();
                                                                                                                 let mempty1
                                                                                                                     =
                                                                                                                     mempty1.clone();
                                                                                                                 move
                                                                                                                     |usd__unused|
                                                                                                                     &Func1::new(move
                                                                                                                                     |dictMonoidRecord|
                                                                                                                                     {
                                                                                                                                         let semigroupRecordCons1 =
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&semigroupRecordCons,
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("SemigroupRecord0"),
                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonoidRecord)),
                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                              &&&Semigroup0);
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_MonoidRecordusd_Dict(),
                                                                                                                                                                          &&&add(string("memptyRecord"),
                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                  let dictMonoidRecord
                                                                                                                                                                                                      =
                                                                                                                                                                                                      dictMonoidRecord.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |v|
                                                                                                                                                                                                      {
                                                                                                                                                                                                          let tail =
                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_memptyRecord(),
                                                                                                                                                                                                                                                                                  &&&dictMonoidRecord),
                                                                                                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                          let key =
                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                  &&&dictIsSymbol),
                                                                                                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                  &&&key),
                                                                                                                                                                                                                                                                              &&&mempty1),
                                                                                                                                                                                                                                           &&&tail)
                                                                                                                                                                                                      }
                                                                                                                                                                                              }),
                                                                                                                                                                                 add(string("SemigroupRecord0"),
                                                                                                                                                                                     &&Func1::new({
                                                                                                                                                                                                      let semigroupRecordCons1
                                                                                                                                                                                                          =
                                                                                                                                                                                                          semigroupRecordCons1.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |usd__unused_1|
                                                                                                                                                                                                          &semigroupRecordCons1
                                                                                                                                                                                                  }),
                                                                                                                                                                                     empty::<string,
                                                                                                                                                                                             &dyn Any>())))
                                                                                                                                     })
                                                                                                             })
                                                                                             }
                                                                                     })
                                                                     }))
    }
    pub fn Data_Monoid_power() -> &dyn Any {
        static Data_Monoid_power: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_power.get_or_init(||
                                          &Func1::new(move |dictMonoid|
                                                          {
                                                              let mempty1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                   dictMonoid);
                                                              let Semigroup0 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                          Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                              &Func1::new({
                                                                              let Semigroup0
                                                                                  =
                                                                                  Semigroup0.clone();
                                                                              let mempty1
                                                                                  =
                                                                                  mempty1.clone();
                                                                              move
                                                                                  |x|
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
                                                                                                                                 |p|
                                                                                                                                 go_tco(p.clone())
                                                                                                                         })
                                                                                                     });
                                                                                      let go_1 =
                                                                                          Lazy(go_2);
                                                                                      let go_tco =
                                                                                          Func1::new({
                                                                                                         let x
                                                                                                             =
                                                                                                             x.clone();
                                                                                                         move
                                                                                                             |p_1|
                                                                                                             fix1(&(move
                                                                                                                        |go_tco,
                                                                                                                         p_1|
                                                                                                                        {
                                                                                                                            let matchValue =
                                                                                                                                Sharpurs_Prelude::unbox(p_1);
                                                                                                                            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                         &&&0_i32))
                                                                                                                               {
                                                                                                                                &mempty1
                                                                                                                            } else {
                                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                             &&&1_i32))
                                                                                                                                   {
                                                                                                                                    &x
                                                                                                                                } else {
                                                                                                                                    if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                                                       &&&2_i32)),
                                                                                                                                                                                                 &&&0_i32))
                                                                                                                                       {
                                                                                                                                        let x_prime =
                                                                                                                                            go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                    &&&2_i32));
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                               &&&Semigroup0),
                                                                                                                                                                                                            &&&x_prime),
                                                                                                                                                                         &&&x_prime)
                                                                                                                                    } else {
                                                                                                                                        if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                           {
                                                                                                                                            let x_prime_1 =
                                                                                                                                                go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                        &&&2_i32));
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                   &&&Semigroup0),
                                                                                                                                                                                                                &&&x_prime_1),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                      &&&Semigroup0),
                                                                                                                                                                                                                                                   &&&x_prime_1),
                                                                                                                                                                                                                &&&x))
                                                                                                                                        } else {
                                                                                                                                            panic!("{}",
                                                                                                                                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Monoid.fs"),
                                  Data1: 35_i32,
                                  Data2: 240_i32,}).get_Message(),)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }),
                                                                                                                  p_1.clone())
                                                                                                     });
                                                                                      let go =
                                                                                          go_1.Value;
                                                                                      &go
                                                                                  }
                                                                          })
                                                          }))
    }
    pub fn Data_Monoid_guard() -> &dyn Any {
        static Data_Monoid_guard: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_guard.get_or_init(||
                                          &Func1::new(move |dictMonoid|
                                                          {
                                                              let mempty1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                   dictMonoid);
                                                              &Func1::new({
                                                                              let mempty1
                                                                                  =
                                                                                  mempty1.clone();
                                                                              move
                                                                                  |v|
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
                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                           &matchValue)
                                                                                                              {
                                                                                                              0_i32
                                                                                                              =>
                                                                                                              a.clone(),
                                                                                                              _
                                                                                                              =>
                                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                               &matchValue)
                                                                                                                  {
                                                                                                                  0_i32
                                                                                                                  =>
                                                                                                                  &mempty1,
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  panic!("{}",
                                                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Monoid.fs"),
                                  Data1: 41_i32,
                                  Data2: 201_i32,}).get_Message(),),
                                                                                                              },
                                                                                                          }
                                                                                                      }
                                                                                              })
                                                                          })
                                                          }))
    }
}
