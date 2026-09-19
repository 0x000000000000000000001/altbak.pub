pub mod PureScript_Test_RBTreeFFICheatcode {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Test_RBTreeFFICheatcode_FFI {
        use super::*;
        use fable_library_rust::Exception_::finally;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerable_1;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerator_1;
        use fable_library_rust::Interfaces_::System::IDisposable;
        use fable_library_rust::Native_::Lrc;
        use fable_library_rust::Native_::LrcPtr;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Range_::rangeNumeric;
        #[derive(Clone, Debug, Default,)]
        pub struct Node {
            Value_0040: i32,
            Black_0040: MutCell<bool>,
            Left_0040: MutCell<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>>,
            Right_0040: MutCell<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>>,
        }
        impl PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node
         {
            pub fn _ctor__Z524259A4(value: i32)
             ->
                 LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
                let Value_0040: i32;
                let Black_0040: bool;
                let Left_0040:
                        LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>;
                let Right_0040:
                        LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>;
                ();
                Value_0040 = value;
                Black_0040 = false;
                Left_0040 = defaultOf();
                Right_0040 = defaultOf();
                ();
                LrcPtr::new(PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node{Value_0040:
                                                                                                      Value_0040,
                                                                                                  Black_0040:
                                                                                                      MutCell::new(Black_0040),
                                                                                                  Left_0040:
                                                                                                      MutCell::new(Left_0040),
                                                                                                  Right_0040:
                                                                                                      MutCell::new(Right_0040),})
            }
            pub fn get_Value(&self) -> i32 { self.Value_0040.clone() }
            pub fn get_Black(&self) -> bool { self.Black_0040.get().clone() }
            pub fn set_Black_Z1FBCCD16(&self, v: bool) {
                self.Black_0040.set(v);
            }
            pub fn get_Left(self: &Lrc<Self>)
             ->
                 LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
                self.Left_0040.get().clone()
            }
            pub fn set_Left_ZD62728F(&self,
                                     v:
                                         LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>) {
                self.Left_0040.set(v);
            }
            pub fn get_Right(self: &Lrc<Self>)
             ->
                 LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
                self.Right_0040.get().clone()
            }
            pub fn set_Right_ZD62728F(&self,
                                      v:
                                          LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>) {
                self.Right_0040.set(v);
            }
        }
        impl core::fmt::Display for
         PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn red(tree:
                       LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>)
         -> bool {
            if !(tree.clone() ==
                     defaultOf::<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>>())
               {
                !tree.get_Black()
            } else { false }
        }
        pub fn balance(tree:
                           LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>)
         ->
             LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
            if !tree.get_Black() {
                tree.clone()
            } else {
                if if PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red(tree.get_Left())
                      {
                       PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red((tree.get_Left()).get_Left())
                   } else { false } {
                    let root:
                            LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                        tree.get_Left();
                    tree.set_Left_ZD62728F(root.get_Right());
                    root.set_Right_ZD62728F(tree.clone());
                    (root.get_Left()).set_Black_Z1FBCCD16(true);
                    root.set_Black_Z1FBCCD16(false);
                    root
                } else {
                    if if PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red(tree.get_Left())
                          {
                           PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red((tree.get_Left()).get_Right())
                       } else { false } {
                        let child:
                                LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                            tree.get_Left();
                        let root_1:
                                LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                            child.get_Right();
                        child.set_Right_ZD62728F(root_1.get_Left());
                        tree.set_Left_ZD62728F(root_1.get_Right());
                        root_1.set_Left_ZD62728F(child.clone());
                        root_1.set_Right_ZD62728F(tree.clone());
                        child.set_Black_Z1FBCCD16(true);
                        root_1.set_Black_Z1FBCCD16(false);
                        root_1
                    } else {
                        if if PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red(tree.get_Right())
                              {
                               PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red((tree.get_Right()).get_Left())
                           } else { false } {
                            let child_1:
                                    LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                                tree.get_Right();
                            let root_2:
                                    LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                                child_1.get_Left();
                            tree.set_Right_ZD62728F(root_2.get_Left());
                            child_1.set_Left_ZD62728F(root_2.get_Right());
                            root_2.set_Left_ZD62728F(tree.clone());
                            root_2.set_Right_ZD62728F(child_1.clone());
                            child_1.set_Black_Z1FBCCD16(true);
                            root_2.set_Black_Z1FBCCD16(false);
                            root_2
                        } else {
                            if if PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red(tree.get_Right())
                                  {
                                   PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::red((tree.get_Right()).get_Right())
                               } else { false } {
                                let root_3:
                                        LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                                    tree.get_Right();
                                tree.set_Right_ZD62728F(root_3.get_Left());
                                root_3.set_Left_ZD62728F(tree.clone());
                                (root_3.get_Right()).set_Black_Z1FBCCD16(true);
                                root_3.set_Black_Z1FBCCD16(false);
                                root_3
                            } else { tree }
                        }
                    }
                }
            }
        }
        pub fn insertRed(value: i32,
                         tree:
                             LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>)
         ->
             LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
            if tree.clone() ==
                   defaultOf::<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>>()
               {
                PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node::_ctor__Z524259A4(value)
            } else {
                if value < tree.get_Value() {
                    tree.set_Left_ZD62728F(PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::insertRed(value,
                                                                                                                      tree.get_Left()));
                    PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::balance(tree.clone())
                } else {
                    if value > tree.get_Value() {
                        tree.set_Right_ZD62728F(PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::insertRed(value,
                                                                                                                           tree.get_Right()));
                        PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::balance(tree.clone())
                    } else { tree }
                }
            }
        }
        pub fn insert(value: i32,
                      tree:
                          LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>)
         ->
             LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> {
            let root:
                    LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node> =
                PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::insertRed(value,
                                                                                           tree);
            root.set_Black_Z1FBCCD16(true);
            root
        }
        pub fn depth(tree:
                         LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>)
         -> i32 {
            if tree.clone() ==
                   defaultOf::<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>>()
               {
                0_i32
            } else {
                1_i32 +
                    PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::depth(tree.get_Left()).max(PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::depth(tree.get_Right()))
            }
        }
        pub fn runRBTreeFFICheatcode(input: &dyn Any) -> &dyn Any {
            let tree:
                    MutCell<LrcPtr<PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::Node>> =
                MutCell::new(defaultOf());
            {
                let inputSequence: LrcPtr<dyn IEnumerable_1<i32>> =
                    rangeNumeric(Sharpurs_Prelude::unbox(input), -1_i32,
                                 1_i32);
                let enumerator: LrcPtr<dyn IEnumerator_1<i32>> =
                    IEnumerable_1::GetEnumerator(inputSequence.as_ref());
                {
                    finally(||
                                if ((&enumerator) as
                                        &dyn Any).is::<LrcPtr<dyn IDisposable>>()
                                   {
                                    (&enumerator).Dispose();
                                });
                    while IEnumerator_1::MoveNext(enumerator.as_ref()) {
                        let value: i32 =
                            IEnumerator_1::get_Current(enumerator.as_ref());
                        tree.set(PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::insert(value,
                                                                                                         tree.get()))
                    }
                }
            }
            &PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::depth(tree.get())
        }
    }
    pub fn Test_RBTreeFFICheatcode_runRBTreeFFICheatcode() -> &dyn Any {
        static Test_RBTreeFFICheatcode_runRBTreeFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFICheatcode_runRBTreeFFICheatcode.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |arg0|
                                                                                      &PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_FFI::runRBTreeFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_RBTreeFFICheatcode_describe() -> &dyn Any {
        static Test_RBTreeFFICheatcode_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFICheatcode_describe.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                          &&&string("Red-Black Tree FFICheatcode (100k Worst-Case Insertions):")))
    }
    pub fn Test_RBTreeFFICheatcode_act() -> &dyn Any {
        static Test_RBTreeFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFICheatcode_act.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                           &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                           &&&100000_i32)),
                                                                                     &&&Func1::new(move
                                                                                                       |dummy|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                           &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                              &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RBTreeFFICheatcode::Test_RBTreeFFICheatcode_runRBTreeFFICheatcode(),
                                                                                                                                                                                                              dummy))))))
    }
}
