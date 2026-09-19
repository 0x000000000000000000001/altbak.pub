#![allow(dead_code,)]
#![allow(non_camel_case_types,)]
#![allow(non_snake_case,)]
#![allow(non_upper_case_globals,)]
#![allow(unreachable_code,)]
#![allow(unused_attributes,)]
#![allow(unused_imports,)]
#![allow(unused_macros,)]
#![allow(unused_parens,)]
#![allow(unused_variables,)]
use fable_library_rust::NativeArray_::array_from;
use fable_library_rust::String_::fromString;
pub mod Sharpurs_EntryPoint {
    use super::*;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::defaultOf;
    use fable_library_rust::NativeArray_::Array;
    use fable_library_rust::String_::string;
    use fable_library_rust::Thread_::Thread;
    use fable_library_rust::Thread_::new;
    use crate::module_33161eca::PureScript_App;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn main(argv: Array<string>) -> i32 {
        let thread: LrcPtr<Thread> =
            new(Func0::new(move ||
                               {
                                   let value =
                                       (Sharpurs_Prelude::unbox(&PureScript_App::App_main()))(defaultOf());
                                   ()
                               }), 1024_i32 * 1024_i32 * 1024_i32);
        thread.start();
        thread.join();
        0_i32
    }
}
#[path = "./App.rs"]
mod module_33161eca;
pub use module_33161eca::*;
#[path = "./AppFFI.rs"]
mod module_4304fb83;
pub use module_4304fb83::*;
#[path = "./AppFFICheatcode.rs"]
mod module_2c2447d5;
pub use module_2c2447d5::*;
#[path = "./AppX.rs"]
mod module_961a3df2;
pub use module_961a3df2::*;
#[path = "./Bench.rs"]
mod module_4da9e9e9;
pub use module_4da9e9e9::*;
#[path = "./Control.Alt.rs"]
mod module_73921b5b;
pub use module_73921b5b::*;
#[path = "./Control.Alternative.rs"]
mod module_9699daad;
pub use module_9699daad::*;
#[path = "./Control.Applicative.rs"]
mod module_ce6bdf8a;
pub use module_ce6bdf8a::*;
#[path = "./Control.Apply.cs"]
mod module_bb42ea47;
pub use module_bb42ea47::*;
#[path = "./Control.Apply.rs"]
mod module_bb42e836;
pub use module_bb42e836::*;
#[path = "./Control.Biapplicative.rs"]
mod module_2f2d2d01;
pub use module_2f2d2d01::*;
#[path = "./Control.Biapply.rs"]
mod module_c52ab4fd;
pub use module_c52ab4fd::*;
#[path = "./Control.Bind.cs"]
mod module_6c4d44b2;
pub use module_6c4d44b2::*;
#[path = "./Control.Bind.rs"]
mod module_6c4d46c3;
pub use module_6c4d46c3::*;
#[path = "./Control.Category.rs"]
mod module_86d6df2;
pub use module_86d6df2::*;
#[path = "./Control.Comonad.Env.Class.rs"]
mod module_e332d6f4;
pub use module_e332d6f4::*;
#[path = "./Control.Comonad.Env.Trans.rs"]
mod module_ca268cc0;
pub use module_ca268cc0::*;
#[path = "./Control.Comonad.Env.rs"]
mod module_3c1181b4;
pub use module_3c1181b4::*;
#[path = "./Control.Comonad.Store.Class.rs"]
mod module_374f55b6;
pub use module_374f55b6::*;
#[path = "./Control.Comonad.Store.Trans.rs"]
mod module_a7b6ae02;
pub use module_a7b6ae02::*;
#[path = "./Control.Comonad.Store.rs"]
mod module_6c234af6;
pub use module_6c234af6::*;
#[path = "./Control.Comonad.Traced.Class.rs"]
mod module_e69784cc;
pub use module_e69784cc::*;
#[path = "./Control.Comonad.Traced.Trans.rs"]
mod module_7d1f6438;
pub use module_7d1f6438::*;
#[path = "./Control.Comonad.Traced.rs"]
mod module_da1c91cc;
pub use module_da1c91cc::*;
#[path = "./Control.Comonad.Trans.Class.rs"]
mod module_35294a53;
pub use module_35294a53::*;
#[path = "./Control.Comonad.rs"]
mod module_df3c4667;
pub use module_df3c4667::*;
#[path = "./Control.Extend.rs"]
mod module_32f29804;
pub use module_32f29804::*;
#[path = "./Control.Lazy.rs"]
mod module_209e0d2c;
pub use module_209e0d2c::*;
#[path = "./Control.Monad.Cont.Class.rs"]
mod module_c6e4dcf3;
pub use module_c6e4dcf3::*;
#[path = "./Control.Monad.Cont.Trans.rs"]
mod module_dbc0dd87;
pub use module_dbc0dd87::*;
#[path = "./Control.Monad.Cont.rs"]
mod module_95133bb3;
pub use module_95133bb3::*;
#[path = "./Control.Monad.Error.Class.rs"]
mod module_7a9a81dd;
pub use module_7a9a81dd::*;
#[path = "./Control.Monad.Except.Trans.rs"]
mod module_8ea4b64e;
pub use module_8ea4b64e::*;
#[path = "./Control.Monad.Except.rs"]
mod module_17353fba;
pub use module_17353fba::*;
#[path = "./Control.Monad.Gen.Class.rs"]
mod module_851afbc9;
pub use module_851afbc9::*;
#[path = "./Control.Monad.Gen.Common.rs"]
mod module_96ec158a;
pub use module_96ec158a::*;
#[path = "./Control.Monad.Gen.rs"]
mod module_e72de349;
pub use module_e72de349::*;
#[path = "./Control.Monad.Identity.Trans.rs"]
mod module_bc5f9127;
pub use module_bc5f9127::*;
#[path = "./Control.Monad.List.Trans.rs"]
mod module_34515613;
pub use module_34515613::*;
#[path = "./Control.Monad.Maybe.Trans.rs"]
mod module_9b44f963;
pub use module_9b44f963::*;
#[path = "./Control.Monad.RWS.Trans.rs"]
mod module_9e2ffde7;
pub use module_9e2ffde7::*;
#[path = "./Control.Monad.RWS.rs"]
mod module_2818553;
pub use module_2818553::*;
#[path = "./Control.Monad.Reader.Class.rs"]
mod module_6641b520;
pub use module_6641b520::*;
#[path = "./Control.Monad.Reader.Trans.rs"]
mod module_8ea61114;
pub use module_8ea61114::*;
#[path = "./Control.Monad.Reader.rs"]
mod module_4a1dd860;
pub use module_4a1dd860::*;
#[path = "./Control.Monad.Rec.Class.rs"]
mod module_e6eed311;
pub use module_e6eed311::*;
#[path = "./Control.Monad.ST.Class.rs"]
mod module_8aa70462;
pub use module_8aa70462::*;
#[path = "./Control.Monad.ST.Global.rs"]
mod module_5b0d6b47;
pub use module_5b0d6b47::*;
#[path = "./Control.Monad.ST.Internal.cs"]
mod module_fcf3045a;
pub use module_fcf3045a::*;
#[path = "./Control.Monad.ST.Internal.rs"]
mod module_fcf3066b;
pub use module_fcf3066b::*;
#[path = "./Control.Monad.ST.Ref.rs"]
mod module_5e390b9d;
pub use module_5e390b9d::*;
#[path = "./Control.Monad.ST.Uncurried.cs"]
mod module_1d594158;
pub use module_1d594158::*;
#[path = "./Control.Monad.ST.Uncurried.rs"]
mod module_1d594369;
pub use module_1d594369::*;
#[path = "./Control.Monad.ST.rs"]
mod module_7e70e22;
pub use module_7e70e22::*;
#[path = "./Control.Monad.State.Class.rs"]
mod module_e3c9db92;
pub use module_e3c9db92::*;
#[path = "./Control.Monad.State.Trans.rs"]
mod module_3bc71566;
pub use module_3bc71566::*;
#[path = "./Control.Monad.State.rs"]
mod module_d006f612;
pub use module_d006f612::*;
#[path = "./Control.Monad.Trans.Class.rs"]
mod module_f5fe307f;
pub use module_f5fe307f::*;
#[path = "./Control.Monad.Writer.Class.rs"]
mod module_c8a91fca;
pub use module_c8a91fca::*;
#[path = "./Control.Monad.Writer.Trans.rs"]
mod module_2b2441be;
pub use module_2b2441be::*;
#[path = "./Control.Monad.Writer.rs"]
mod module_dc2f25ca;
pub use module_dc2f25ca::*;
#[path = "./Control.Monad.rs"]
mod module_ed1bca0b;
pub use module_ed1bca0b::*;
#[path = "./Control.MonadPlus.rs"]
mod module_53a2e11;
pub use module_53a2e11::*;
#[path = "./Control.Parallel.Class.rs"]
mod module_255d4e69;
pub use module_255d4e69::*;
#[path = "./Control.Parallel.rs"]
mod module_cd8e8e29;
pub use module_cd8e8e29::*;
#[path = "./Control.Plus.rs"]
mod module_6afec8d8;
pub use module_6afec8d8::*;
#[path = "./Control.Semigroupoid.rs"]
mod module_655da3ed;
pub use module_655da3ed::*;
#[path = "./Data.Array.NonEmpty.Internal.rs"]
mod module_ba715d91;
pub use module_ba715d91::*;
#[path = "./Data.Array.NonEmpty.rs"]
mod module_34a61018;
pub use module_34a61018::*;
#[path = "./Data.Array.Partial.rs"]
mod module_5ccc2071;
pub use module_5ccc2071::*;
#[path = "./Data.Array.ST.Iterator.rs"]
mod module_e9f7eca9;
pub use module_e9f7eca9::*;
#[path = "./Data.Array.ST.Partial.cs"]
mod module_48c30649;
pub use module_48c30649::*;
#[path = "./Data.Array.ST.Partial.rs"]
mod module_48c30438;
pub use module_48c30438::*;
#[path = "./Data.Array.ST.cs"]
mod module_acc76a94;
pub use module_acc76a94::*;
#[path = "./Data.Array.ST.rs"]
mod module_acc76ca5;
pub use module_acc76ca5::*;
#[path = "./Data.Array.cs"]
mod module_2d8de9d;
pub use module_2d8de9d::*;
#[path = "./Data.Array.rs"]
mod module_2d8e16c;
pub use module_2d8e16c::*;
#[path = "./Data.Bifoldable.rs"]
mod module_637e1ff5;
pub use module_637e1ff5::*;
#[path = "./Data.Bifunctor.Join.rs"]
mod module_6273045;
pub use module_6273045::*;
#[path = "./Data.Bifunctor.rs"]
mod module_c4b10869;
pub use module_c4b10869::*;
#[path = "./Data.Bitraversable.rs"]
mod module_bdb37be1;
pub use module_bdb37be1::*;
#[path = "./Data.Boolean.rs"]
mod module_a4848631;
pub use module_a4848631::*;
#[path = "./Data.BooleanAlgebra.rs"]
mod module_13840e4f;
pub use module_13840e4f::*;
#[path = "./Data.Bounded.Generic.rs"]
mod module_7f163039;
pub use module_7f163039::*;
#[path = "./Data.Bounded.cs"]
mod module_d89c2d77;
pub use module_d89c2d77::*;
#[path = "./Data.Bounded.rs"]
mod module_d89c2f46;
pub use module_d89c2f46::*;
#[path = "./Data.Char.Gen.rs"]
mod module_5ec638cf;
pub use module_5ec638cf::*;
#[path = "./Data.Char.rs"]
mod module_893d708d;
pub use module_893d708d::*;
#[path = "./Data.CommutativeRing.rs"]
mod module_cb047b05;
pub use module_cb047b05::*;
#[path = "./Data.Comparison.rs"]
mod module_18b8108c;
pub use module_18b8108c::*;
#[path = "./Data.Const.rs"]
mod module_a8445950;
pub use module_a8445950::*;
#[path = "./Data.Date.Component.Gen.rs"]
mod module_bdfbe862;
pub use module_bdfbe862::*;
#[path = "./Data.Date.Component.rs"]
mod module_b6aac8e0;
pub use module_b6aac8e0::*;
#[path = "./Data.Date.Gen.rs"]
mod module_41690983;
pub use module_41690983::*;
#[path = "./Data.Date.cs"]
mod module_d96ec0b0;
pub use module_d96ec0b0::*;
#[path = "./Data.Date.rs"]
mod module_d96ec2c1;
pub use module_d96ec2c1::*;
#[path = "./Data.DateTime.Gen.rs"]
mod module_a5b31cb6;
pub use module_a5b31cb6::*;
#[path = "./Data.DateTime.Instant.cs"]
mod module_47c4d50;
pub use module_47c4d50::*;
#[path = "./Data.DateTime.Instant.rs"]
mod module_47c4f61;
pub use module_47c4f61::*;
#[path = "./Data.DateTime.cs"]
mod module_867837c5;
pub use module_867837c5::*;
#[path = "./Data.DateTime.rs"]
mod module_867835b4;
pub use module_867835b4::*;
#[path = "./Data.Decidable.rs"]
mod module_e3f65a70;
pub use module_e3f65a70::*;
#[path = "./Data.Decide.rs"]
mod module_abb3d73f;
pub use module_abb3d73f::*;
#[path = "./Data.Distributive.rs"]
mod module_83e4823d;
pub use module_83e4823d::*;
#[path = "./Data.Divide.rs"]
mod module_d8620aa6;
pub use module_d8620aa6::*;
#[path = "./Data.Divisible.rs"]
mod module_747333f6;
pub use module_747333f6::*;
#[path = "./Data.DivisionRing.rs"]
mod module_5ba8cfce;
pub use module_5ba8cfce::*;
#[path = "./Data.Either.Inject.rs"]
mod module_7c3591a3;
pub use module_7c3591a3::*;
#[path = "./Data.Either.Nested.rs"]
mod module_e29c3231;
pub use module_e29c3231::*;
#[path = "./Data.Either.rs"]
mod module_173929b2;
pub use module_173929b2::*;
#[path = "./Data.Enum.Gen.rs"]
mod module_25dacb64;
pub use module_25dacb64::*;
#[path = "./Data.Enum.Generic.rs"]
mod module_5a661619;
pub use module_5a661619::*;
#[path = "./Data.Enum.cs"]
mod module_6a1c5a17;
pub use module_6a1c5a17::*;
#[path = "./Data.Enum.rs"]
mod module_6a1c5ce6;
pub use module_6a1c5ce6::*;
#[path = "./Data.Eq.Generic.rs"]
mod module_63dab5e;
pub use module_63dab5e::*;
#[path = "./Data.Eq.cs"]
mod module_8617f750;
pub use module_8617f750::*;
#[path = "./Data.Eq.rs"]
mod module_8617f961;
pub use module_8617f961::*;
#[path = "./Data.Equivalence.rs"]
mod module_59dcd94b;
pub use module_59dcd94b::*;
#[path = "./Data.EuclideanRing.cs"]
mod module_26da002e;
pub use module_26da002e::*;
#[path = "./Data.EuclideanRing.rs"]
mod module_26d9fe5f;
pub use module_26d9fe5f::*;
#[path = "./Data.Exists.rs"]
mod module_c353b055;
pub use module_c353b055::*;
#[path = "./Data.Field.rs"]
mod module_68f04657;
pub use module_68f04657::*;
#[path = "./Data.Foldable.cs"]
mod module_419ed0af;
pub use module_419ed0af::*;
#[path = "./Data.Foldable.rs"]
mod module_419ece9e;
pub use module_419ece9e::*;
#[path = "./Data.FoldableWithIndex.rs"]
mod module_9201da02;
pub use module_9201da02::*;
#[path = "./Data.Function.Uncurried.cs"]
mod module_d2b7e70d;
pub use module_d2b7e70d::*;
#[path = "./Data.Function.Uncurried.rs"]
mod module_d2b7e5fc;
pub use module_d2b7e5fc::*;
#[path = "./Data.Function.rs"]
mod module_6e7709b7;
pub use module_6e7709b7::*;
#[path = "./Data.Functor.App.rs"]
mod module_3ae611ad;
pub use module_3ae611ad::*;
#[path = "./Data.Functor.Clown.rs"]
mod module_58c69d5;
pub use module_58c69d5::*;
#[path = "./Data.Functor.Compose.rs"]
mod module_ea451784;
pub use module_ea451784::*;
#[path = "./Data.Functor.Contravariant.rs"]
mod module_cf56105e;
pub use module_cf56105e::*;
#[path = "./Data.Functor.Coproduct.Inject.rs"]
mod module_e057f0da;
pub use module_e057f0da::*;
#[path = "./Data.Functor.Coproduct.Nested.rs"]
mod module_5dd720c8;
pub use module_5dd720c8::*;
#[path = "./Data.Functor.Coproduct.rs"]
mod module_97eca9ab;
pub use module_97eca9ab::*;
#[path = "./Data.Functor.Costar.rs"]
mod module_22da7174;
pub use module_22da7174::*;
#[path = "./Data.Functor.Flip.rs"]
mod module_f8895b9f;
pub use module_f8895b9f::*;
#[path = "./Data.Functor.Invariant.rs"]
mod module_ddf66a9c;
pub use module_ddf66a9c::*;
#[path = "./Data.Functor.Joker.rs"]
mod module_8377b1b5;
pub use module_8377b1b5::*;
#[path = "./Data.Functor.Product.Nested.rs"]
mod module_eb2d88e4;
pub use module_eb2d88e4::*;
#[path = "./Data.Functor.Product.rs"]
mod module_4f0b17c7;
pub use module_4f0b17c7::*;
#[path = "./Data.Functor.Product2.rs"]
mod module_3079b2b5;
pub use module_3079b2b5::*;
#[path = "./Data.Functor.cs"]
mod module_b38ced3;
pub use module_b38ced3::*;
#[path = "./Data.Functor.rs"]
mod module_b38d0a2;
pub use module_b38d0a2::*;
#[path = "./Data.FunctorWithIndex.cs"]
mod module_f8b1f58f;
pub use module_f8b1f58f::*;
#[path = "./Data.FunctorWithIndex.rs"]
mod module_f8b1f47e;
pub use module_f8b1f47e::*;
#[path = "./Data.Generic.Rep.rs"]
mod module_b37db3cd;
pub use module_b37db3cd::*;
#[path = "./Data.HeytingAlgebra.Generic.rs"]
mod module_882c80d4;
pub use module_882c80d4::*;
#[path = "./Data.HeytingAlgebra.cs"]
mod module_c0687dda;
pub use module_c0687dda::*;
#[path = "./Data.HeytingAlgebra.rs"]
mod module_c0687feb;
pub use module_c0687feb::*;
#[path = "./Data.Identity.rs"]
mod module_1becb483;
pub use module_1becb483::*;
#[path = "./Data.Int.Bits.cs"]
mod module_e83e1f95;
pub use module_e83e1f95::*;
#[path = "./Data.Int.Bits.rs"]
mod module_e83e2264;
pub use module_e83e2264::*;
#[path = "./Data.Int.cs"]
mod module_2eb6dcf7;
pub use module_2eb6dcf7::*;
#[path = "./Data.Int.rs"]
mod module_2eb6dec6;
pub use module_2eb6dec6::*;
#[path = "./Data.Interval.Duration.Iso.rs"]
mod module_7baaee01;
pub use module_7baaee01::*;
#[path = "./Data.Interval.Duration.rs"]
mod module_8669b9ba;
pub use module_8669b9ba::*;
#[path = "./Data.Interval.rs"]
mod module_63eb8d6a;
pub use module_63eb8d6a::*;
#[path = "./Data.Lazy.rs"]
mod module_720e12db;
pub use module_720e12db::*;
#[path = "./Data.List.Internal.rs"]
mod module_4df90a9e;
pub use module_4df90a9e::*;
#[path = "./Data.List.Lazy.NonEmpty.rs"]
mod module_6b572c63;
pub use module_6b572c63::*;
#[path = "./Data.List.Lazy.Types.rs"]
mod module_44df2f32;
pub use module_44df2f32::*;
#[path = "./Data.List.Lazy.rs"]
mod module_529acc77;
pub use module_529acc77::*;
#[path = "./Data.List.NonEmpty.rs"]
mod module_eb82fa3;
pub use module_eb82fa3::*;
#[path = "./Data.List.Partial.rs"]
mod module_ba7076aa;
pub use module_ba7076aa::*;
#[path = "./Data.List.Types.rs"]
mod module_d662adf2;
pub use module_d662adf2::*;
#[path = "./Data.List.ZipList.rs"]
mod module_e8f004d8;
pub use module_e8f004d8::*;
#[path = "./Data.List.rs"]
mod module_843b47b7;
pub use module_843b47b7::*;
#[path = "./Data.Map.Gen.rs"]
mod module_e3e1610b;
pub use module_e3e1610b::*;
#[path = "./Data.Map.Internal.rs"]
mod module_ed2bf3e0;
pub use module_ed2bf3e0::*;
#[path = "./Data.Map.rs"]
mod module_37b01849;
pub use module_37b01849::*;
#[path = "./Data.Maybe.First.rs"]
mod module_bde357b3;
pub use module_bde357b3::*;
#[path = "./Data.Maybe.Last.rs"]
mod module_a305e0e3;
pub use module_a305e0e3::*;
#[path = "./Data.Maybe.rs"]
mod module_f879ba47;
pub use module_f879ba47::*;
#[path = "./Data.Monoid.Additive.rs"]
mod module_3b72fe33;
pub use module_3b72fe33::*;
#[path = "./Data.Monoid.Alternate.rs"]
mod module_ee682245;
pub use module_ee682245::*;
#[path = "./Data.Monoid.Conj.rs"]
mod module_89cc47bd;
pub use module_89cc47bd::*;
#[path = "./Data.Monoid.Disj.rs"]
mod module_6a9d0e21;
pub use module_6a9d0e21::*;
#[path = "./Data.Monoid.Dual.rs"]
mod module_5023b7e9;
pub use module_5023b7e9::*;
#[path = "./Data.Monoid.Endo.rs"]
mod module_f1079b5;
pub use module_f1079b5::*;
#[path = "./Data.Monoid.Generic.rs"]
mod module_c86e38e4;
pub use module_c86e38e4::*;
#[path = "./Data.Monoid.Multiplicative.rs"]
mod module_7d83a5e5;
pub use module_7d83a5e5::*;
#[path = "./Data.Monoid.rs"]
mod module_5e99fcdb;
pub use module_5e99fcdb::*;
#[path = "./Data.NaturalTransformation.rs"]
mod module_47768895;
pub use module_47768895::*;
#[path = "./Data.Newtype.rs"]
mod module_146b7611;
pub use module_146b7611::*;
#[path = "./Data.NonEmpty.rs"]
mod module_d6a130cf;
pub use module_d6a130cf::*;
#[path = "./Data.Number.Approximate.rs"]
mod module_ef2d6ae8;
pub use module_ef2d6ae8::*;
#[path = "./Data.Number.Format.cs"]
mod module_c0b7ca2a;
pub use module_c0b7ca2a::*;
#[path = "./Data.Number.Format.rs"]
mod module_c0b7c85b;
pub use module_c0b7c85b::*;
#[path = "./Data.Number.cs"]
mod module_91782cc7;
pub use module_91782cc7::*;
#[path = "./Data.Number.rs"]
mod module_91782ab6;
pub use module_91782ab6::*;
#[path = "./Data.Op.rs"]
mod module_851c93ca;
pub use module_851c93ca::*;
#[path = "./Data.Ord.Down.rs"]
mod module_2d58bb50;
pub use module_2d58bb50::*;
#[path = "./Data.Ord.Generic.rs"]
mod module_1f9c8bf3;
pub use module_1f9c8bf3::*;
#[path = "./Data.Ord.Max.rs"]
mod module_a3a23b56;
pub use module_a3a23b56::*;
#[path = "./Data.Ord.Min.rs"]
mod module_a424d288;
pub use module_a424d288::*;
#[path = "./Data.Ord.cs"]
mod module_28ab9abd;
pub use module_28ab9abd::*;
#[path = "./Data.Ord.rs"]
mod module_28ab9c8c;
pub use module_28ab9c8c::*;
#[path = "./Data.Ordering.rs"]
mod module_5f769efb;
pub use module_5f769efb::*;
#[path = "./Data.Predicate.rs"]
mod module_d23c04ec;
pub use module_d23c04ec::*;
#[path = "./Data.Profunctor.Choice.rs"]
mod module_8c6f434a;
pub use module_8c6f434a::*;
#[path = "./Data.Profunctor.Closed.rs"]
mod module_6da15eb3;
pub use module_6da15eb3::*;
#[path = "./Data.Profunctor.Cochoice.rs"]
mod module_2c43c026;
pub use module_2c43c026::*;
#[path = "./Data.Profunctor.Costrong.rs"]
mod module_a7d13e7e;
pub use module_a7d13e7e::*;
#[path = "./Data.Profunctor.Join.rs"]
mod module_1f049ce3;
pub use module_1f049ce3::*;
#[path = "./Data.Profunctor.Split.rs"]
mod module_5575a6d3;
pub use module_5575a6d3::*;
#[path = "./Data.Profunctor.Star.rs"]
mod module_cba67fd5;
pub use module_cba67fd5::*;
#[path = "./Data.Profunctor.Strong.rs"]
mod module_15730292;
pub use module_15730292::*;
#[path = "./Data.Profunctor.rs"]
mod module_2db53acf;
pub use module_2db53acf::*;
#[path = "./Data.Reflectable.cs"]
mod module_78857e41;
pub use module_78857e41::*;
#[path = "./Data.Reflectable.rs"]
mod module_78857c30;
pub use module_78857c30::*;
#[path = "./Data.Ring.Generic.rs"]
mod module_89cee7b8;
pub use module_89cee7b8::*;
#[path = "./Data.Ring.cs"]
mod module_111eaf6;
pub use module_111eaf6::*;
#[path = "./Data.Ring.rs"]
mod module_111ec07;
pub use module_111ec07::*;
#[path = "./Data.Semigroup.First.rs"]
mod module_9f32becc;
pub use module_9f32becc::*;
#[path = "./Data.Semigroup.Foldable.rs"]
mod module_2563115d;
pub use module_2563115d::*;
#[path = "./Data.Semigroup.Generic.rs"]
mod module_2b634507;
pub use module_2b634507::*;
#[path = "./Data.Semigroup.Last.rs"]
mod module_8e57e1bc;
pub use module_8e57e1bc::*;
#[path = "./Data.Semigroup.Traversable.rs"]
mod module_abab3d09;
pub use module_abab3d09::*;
#[path = "./Data.Semigroup.cs"]
mod module_ab537a09;
pub use module_ab537a09::*;
#[path = "./Data.Semigroup.rs"]
mod module_ab5378f8;
pub use module_ab5378f8::*;
#[path = "./Data.Semiring.Generic.rs"]
mod module_682be16a;
pub use module_682be16a::*;
#[path = "./Data.Semiring.cs"]
mod module_be45b324;
pub use module_be45b324::*;
#[path = "./Data.Semiring.rs"]
mod module_be45b155;
pub use module_be45b155::*;
#[path = "./Data.Set.NonEmpty.rs"]
mod module_9672c243;
pub use module_9672c243::*;
#[path = "./Data.Set.rs"]
mod module_e534597;
pub use module_e534597::*;
#[path = "./Data.Show.Generic.cs"]
mod module_afe20778;
pub use module_afe20778::*;
#[path = "./Data.Show.Generic.rs"]
mod module_afe20889;
pub use module_afe20889::*;
#[path = "./Data.Show.cs"]
mod module_baebcc87;
pub use module_baebcc87::*;
#[path = "./Data.Show.rs"]
mod module_baebcb76;
pub use module_baebcb76::*;
#[path = "./Data.String.CaseInsensitive.rs"]
mod module_ae3b0e31;
pub use module_ae3b0e31::*;
#[path = "./Data.String.CodePoints.cs"]
mod module_2f0084ed;
pub use module_2f0084ed::*;
#[path = "./Data.String.CodePoints.rs"]
mod module_2f0082dc;
pub use module_2f0082dc::*;
#[path = "./Data.String.CodeUnits.cs"]
mod module_eea318e7;
pub use module_eea318e7::*;
#[path = "./Data.String.CodeUnits.rs"]
mod module_eea316d6;
pub use module_eea316d6::*;
#[path = "./Data.String.Common.cs"]
mod module_eb405752;
pub use module_eb405752::*;
#[path = "./Data.String.Common.rs"]
mod module_eb405963;
pub use module_eb405963::*;
#[path = "./Data.String.Gen.rs"]
mod module_50f032c2;
pub use module_50f032c2::*;
#[path = "./Data.String.NonEmpty.CaseInsensitive.rs"]
mod module_833f1e05;
pub use module_833f1e05::*;
#[path = "./Data.String.NonEmpty.CodePoints.rs"]
mod module_4c773b68;
pub use module_4c773b68::*;
#[path = "./Data.String.NonEmpty.CodeUnits.rs"]
mod module_6fbe5722;
pub use module_6fbe5722::*;
#[path = "./Data.String.NonEmpty.Internal.rs"]
mod module_d71935fd;
pub use module_d71935fd::*;
#[path = "./Data.String.NonEmpty.rs"]
mod module_8d95974;
pub use module_8d95974::*;
#[path = "./Data.String.Pattern.rs"]
mod module_91241446;
pub use module_91241446::*;
#[path = "./Data.String.Regex.Flags.rs"]
mod module_df7b54b2;
pub use module_df7b54b2::*;
#[path = "./Data.String.Regex.Unsafe.rs"]
mod module_f38e6187;
pub use module_f38e6187::*;
#[path = "./Data.String.Regex.rs"]
mod module_db031083;
pub use module_db031083::*;
#[path = "./Data.String.Unsafe.cs"]
mod module_b0fd7715;
pub use module_b0fd7715::*;
#[path = "./Data.String.Unsafe.rs"]
mod module_b0fd79e4;
pub use module_b0fd79e4::*;
#[path = "./Data.String.rs"]
mod module_1e5a4c80;
pub use module_1e5a4c80::*;
#[path = "./Data.Symbol.cs"]
mod module_20f33982;
pub use module_20f33982::*;
#[path = "./Data.Symbol.rs"]
mod module_20f337b3;
pub use module_20f337b3::*;
#[path = "./Data.Time.Component.Gen.rs"]
mod module_c0ca5c63;
pub use module_c0ca5c63::*;
#[path = "./Data.Time.Component.rs"]
mod module_f62dae61;
pub use module_f62dae61::*;
#[path = "./Data.Time.Duration.Gen.rs"]
mod module_73de7652;
pub use module_73de7652::*;
#[path = "./Data.Time.Duration.rs"]
mod module_c4b34810;
pub use module_c4b34810::*;
#[path = "./Data.Time.Gen.rs"]
mod module_b10b4242;
pub use module_b10b4242::*;
#[path = "./Data.Time.rs"]
mod module_f8599c00;
pub use module_f8599c00::*;
#[path = "./Data.Traversable.Accum.Internal.rs"]
mod module_833c8c54;
pub use module_833c8c54::*;
#[path = "./Data.Traversable.Accum.rs"]
mod module_26e4d8bd;
pub use module_26e4d8bd::*;
#[path = "./Data.Traversable.cs"]
mod module_92875c5b;
pub use module_92875c5b::*;
#[path = "./Data.Traversable.rs"]
mod module_92875e2a;
pub use module_92875e2a::*;
#[path = "./Data.TraversableWithIndex.rs"]
mod module_829cacf6;
pub use module_829cacf6::*;
#[path = "./Data.Tuple.Nested.rs"]
mod module_dfff516e;
pub use module_dfff516e::*;
#[path = "./Data.Tuple.rs"]
mod module_e7bd458d;
pub use module_e7bd458d::*;
#[path = "./Data.Unfoldable.rs"]
mod module_98c530c5;
pub use module_98c530c5::*;
#[path = "./Data.Unfoldable1.cs"]
mod module_b1755125;
pub use module_b1755125::*;
#[path = "./Data.Unfoldable1.rs"]
mod module_b1754f14;
pub use module_b1754f14::*;
#[path = "./Data.Unit.cs"]
mod module_9be98062;
pub use module_9be98062::*;
#[path = "./Data.Unit.rs"]
mod module_9be97d93;
pub use module_9be97d93::*;
#[path = "./Data.Void.rs"]
mod module_38c1e8e1;
pub use module_38c1e8e1::*;
#[path = "./Effect.Aff.Class.rs"]
mod module_a62e71f3;
pub use module_a62e71f3::*;
#[path = "./Effect.Aff.Compat.rs"]
mod module_43c1a259;
pub use module_43c1a259::*;
#[path = "./Effect.Aff.cs"]
mod module_36111282;
pub use module_36111282::*;
#[path = "./Effect.Aff.rs"]
mod module_361110b3;
pub use module_361110b3::*;
#[path = "./Effect.Class.Console.rs"]
mod module_b25f9165;
pub use module_b25f9165::*;
#[path = "./Effect.Class.rs"]
mod module_21d6b3bc;
pub use module_21d6b3bc::*;
#[path = "./Effect.Console.cs"]
mod module_3ed61a14;
pub use module_3ed61a14::*;
#[path = "./Effect.Console.rs"]
mod module_3ed61c25;
pub use module_3ed61c25::*;
#[path = "./Effect.Exception.Unsafe.rs"]
mod module_96b36061;
pub use module_96b36061::*;
#[path = "./Effect.Exception.cs"]
mod module_ceb94194;
pub use module_ceb94194::*;
#[path = "./Effect.Exception.rs"]
mod module_ceb943a5;
pub use module_ceb943a5::*;
#[path = "./Effect.Now.rs"]
mod module_14601aa4;
pub use module_14601aa4::*;
#[path = "./Effect.Ref.cs"]
mod module_9d127b2;
pub use module_9d127b2::*;
#[path = "./Effect.Ref.rs"]
mod module_9d129c3;
pub use module_9d129c3::*;
#[path = "./Effect.Uncurried.cs"]
mod module_edf44d86;
pub use module_edf44d86::*;
#[path = "./Effect.Uncurried.rs"]
mod module_edf44bb7;
pub use module_edf44bb7::*;
#[path = "./Effect.Unsafe.cs"]
mod module_90c22ee9;
pub use module_90c22ee9::*;
#[path = "./Effect.Unsafe.rs"]
mod module_90c22cd8;
pub use module_90c22cd8::*;
#[path = "./Effect.cs"]
mod module_5706e30d;
pub use module_5706e30d::*;
#[path = "./Effect.rs"]
mod module_5706e1fc;
pub use module_5706e1fc::*;
#[path = "./Main.rs"]
mod module_a74a3de0;
pub use module_a74a3de0::*;
#[path = "./Partial.Unsafe.cs"]
mod module_1180104d;
pub use module_1180104d::*;
#[path = "./Partial.Unsafe.rs"]
mod module_11800e3c;
pub use module_11800e3c::*;
#[path = "./Partial.cs"]
mod module_c2818e9;
pub use module_c2818e9::*;
#[path = "./Partial.rs"]
mod module_c2816d8;
pub use module_c2816d8::*;
#[path = "./Prelude.rs"]
mod module_692e3d4;
pub use module_692e3d4::*;
#[path = "./Record.Unsafe.cs"]
mod module_312cfa53;
pub use module_312cfa53::*;
#[path = "./Record.Unsafe.rs"]
mod module_312cfc22;
pub use module_312cfc22::*;
#[path = "./Safe.Coerce.rs"]
mod module_285149e9;
pub use module_285149e9::*;
#[path = "./Sharpurs_Prelude.rs"]
mod module_aa21d1e7;
pub use module_aa21d1e7::*;
#[path = "./Spago.Generated.BuildInfo.rs"]
mod module_e96f94d6;
pub use module_e96f94d6::*;
#[path = "./Test.Ackermann.rs"]
mod module_b3bd47a1;
pub use module_b3bd47a1::*;
#[path = "./Test.AckermannFFI.rs"]
mod module_db4e0a48;
pub use module_db4e0a48::*;
#[path = "./Test.AckermannFFICheatcode.rs"]
mod module_35c6b5be;
pub use module_35c6b5be::*;
#[path = "./Test.ArrayOps.rs"]
mod module_d5481526;
pub use module_d5481526::*;
#[path = "./Test.ArrayOpsFFI.rs"]
mod module_6ae0516f;
pub use module_6ae0516f::*;
#[path = "./Test.ArrayOpsFFICheatcode.rs"]
mod module_4aebffb9;
pub use module_4aebffb9::*;
#[path = "./Test.AstTree.rs"]
mod module_68d0f733;
pub use module_68d0f733::*;
#[path = "./Test.AstTreeFFI.rs"]
mod module_2bde461a;
pub use module_2bde461a::*;
#[path = "./Test.AstTreeFFICheatcode.rs"]
mod module_3748caac;
pub use module_3748caac::*;
#[path = "./Test.BenchCheck.rs"]
mod module_fbfa7037;
pub use module_fbfa7037::*;
#[path = "./Test.Church.rs"]
mod module_83b49eb4;
pub use module_83b49eb4::*;
#[path = "./Test.ChurchFFI.rs"]
mod module_df6db8bd;
pub use module_df6db8bd::*;
#[path = "./Test.ChurchFFICheatcode.rs"]
mod module_43218cab;
pub use module_43218cab::*;
#[path = "./Test.Fib.rs"]
mod module_13ca489e;
pub use module_13ca489e::*;
#[path = "./Test.FibFFI.rs"]
mod module_c5354257;
pub use module_c5354257::*;
#[path = "./Test.FibFFICheatcode.rs"]
mod module_7fb6d281;
pub use module_7fb6d281::*;
#[path = "./Test.Implicit.rs"]
mod module_8a4f977c;
pub use module_8a4f977c::*;
#[path = "./Test.LazyEvaluation.rs"]
mod module_c7892d4b;
pub use module_c7892d4b::*;
#[path = "./Test.LazyEvaluationFFI.rs"]
mod module_72ccf622;
pub use module_72ccf622::*;
#[path = "./Test.LazyEvaluationFFICheatcode.rs"]
mod module_7c9fd554;
pub use module_7c9fd554::*;
#[path = "./Test.ListOps.rs"]
mod module_996c04dd;
pub use module_996c04dd::*;
#[path = "./Test.ListOpsFFI.rs"]
mod module_f72efaf4;
pub use module_f72efaf4::*;
#[path = "./Test.ListOpsFFICheatcode.rs"]
mod module_a0d6b42;
pub use module_a0d6b42::*;
#[path = "./Test.Main.rs"]
mod module_7364c658;
pub use module_7364c658::*;
#[path = "./Test.Polymorphism.rs"]
mod module_2bcb98c6;
pub use module_2bcb98c6::*;
#[path = "./Test.PolymorphismFFI.rs"]
mod module_e3376a0f;
pub use module_e3376a0f::*;
#[path = "./Test.PolymorphismFFICheatcode.rs"]
mod module_4ec4c4d9;
pub use module_4ec4c4d9::*;
#[path = "./Test.Primes.rs"]
mod module_55d52a3;
pub use module_55d52a3::*;
#[path = "./Test.PrimesFFI.rs"]
mod module_3cdcab8a;
pub use module_3cdcab8a::*;
#[path = "./Test.PrimesFFICheatcode.rs"]
mod module_93de623c;
pub use module_93de623c::*;
#[path = "./Test.RBTree.rs"]
mod module_81195d05;
pub use module_81195d05::*;
#[path = "./Test.RBTreeFFI.rs"]
mod module_973c6bac;
pub use module_973c6bac::*;
#[path = "./Test.RBTreeFFICheatcode.rs"]
mod module_5196879a;
pub use module_5196879a::*;
#[path = "./Test.Records.rs"]
mod module_2992b04d;
pub use module_2992b04d::*;
#[path = "./Test.RecordsFFI.rs"]
mod module_b965b064;
pub use module_b965b064::*;
#[path = "./Test.RecordsFFICheatcode.rs"]
mod module_443552d2;
pub use module_443552d2::*;
#[path = "./Test.RowToList.rs"]
mod module_fefd76a0;
pub use module_fefd76a0::*;
#[path = "./Test.RowToListFFI.rs"]
mod module_6a7bbb29;
pub use module_6a7bbb29::*;
#[path = "./Test.RowToListFFICheatcode.rs"]
mod module_fcb563bf;
pub use module_fcb563bf::*;
#[path = "./Test.StateMonad.rs"]
mod module_9f617d6d;
pub use module_9f617d6d::*;
#[path = "./Test.StateMonadFFI.rs"]
mod module_e406b004;
pub use module_e406b004::*;
#[path = "./Test.StateMonadFFICheatcode.rs"]
mod module_e6b88172;
pub use module_e6b88172::*;
#[path = "./Test.TCO.rs"]
mod module_e3c8f2ab;
pub use module_e3c8f2ab::*;
#[path = "./Test.TCOFFI.rs"]
mod module_39dea682;
pub use module_39dea682::*;
#[path = "./Test.TCOFFICheatcode.rs"]
mod module_dcb2fc34;
pub use module_dcb2fc34::*;
#[path = "./Type.Equality.rs"]
mod module_6b8c0d95;
pub use module_6b8c0d95::*;
#[path = "./Type.Proxy.rs"]
mod module_48ec9431;
pub use module_48ec9431::*;
#[path = "./Unsafe.Coerce.cs"]
mod module_2a7664e3;
pub use module_2a7664e3::*;
#[path = "./Unsafe.Coerce.rs"]
mod module_2a7662d2;
pub use module_2a7662d2::*;
pub fn main() {
    let args = std::env::args().skip(1).map(fromString).collect();
    Sharpurs_EntryPoint::main(array_from(args));
}
