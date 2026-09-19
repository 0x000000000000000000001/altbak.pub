pub mod PureScript_Test_RBTreeFFI {
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
    pub mod Test_RBTreeFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::LrcPtr;
        use fable_library_rust::Native_::fix1;
        #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
        pub enum Color { Red, Black, }
        impl core::fmt::Display for
         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
        pub enum Tree {
            Empty,
            Node(LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color>,
                 LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>,
                 i32,
                 LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>),
        }
        impl core::fmt::Display for
         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn balance(_arg1_:
                           LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color>,
                       _arg1__1:
                           LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>,
                       _arg1__2: i32,
                       _arg1__3:
                           LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>)
         -> LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> {
            let _arg:
                    LrcPtr<(LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color>,
                            LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>,
                            i32,
                            LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>)> =
                LrcPtr::new((_arg1_, _arg1__1, _arg1__2, _arg1__3));
            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black
                   = _arg.0.clone().as_ref() {
                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                       = _arg.1.clone().as_ref() {
                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                           =
                           match _arg.1.clone().as_ref() {
                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                         _,
                                                                                         _,
                                                                                         _)
                               => x.clone(),
                               _ => unreachable!(),
                           }.as_ref() {
                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                               =
                               match _arg.1.clone().as_ref() {
                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                             x,
                                                                                             _,
                                                                                             _)
                                   => x.clone(),
                                   _ => unreachable!(),
                               }.as_ref() {
                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                   =
                                   match match _arg.1.clone().as_ref() {
                                             PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                       x,
                                                                                                       _,
                                                                                                       _)
                                             => x.clone(),
                                             _ => unreachable!(),
                                         }.as_ref() {
                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                 _,
                                                                                                 _,
                                                                                                 _)
                                       => x.clone(),
                                       _ => unreachable!(),
                                   }.as_ref() {
                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                            match match _arg.1.clone().as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                _)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          _)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            },
                                                                                                                                                                            match match _arg.1.clone().as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                _)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                          _)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            },
                                                                                                                                                                            match match _arg.1.clone().as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                _)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            })),
                                                                                                      match _arg.1.clone().as_ref()
                                                                                                          {
                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                    _,
                                                                                                                                                                    x,
                                                                                                                                                                    _)
                                                                                                          =>
                                                                                                          x.clone(),
                                                                                                          _
                                                                                                          =>
                                                                                                          unreachable!(),
                                                                                                      },
                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                            match _arg.1.clone().as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            },
                                                                                                                                                                            _arg.2.clone(),
                                                                                                                                                                            _arg.3.clone()))))
                            } else {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                       =
                                       match _arg.1.clone().as_ref() {
                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                     _,
                                                                                                     _,
                                                                                                     x)
                                           => x.clone(),
                                           _ => unreachable!(),
                                       }.as_ref() {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                           =
                                           match match _arg.1.clone().as_ref()
                                                     {
                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                               _,
                                                                                                               _,
                                                                                                               x)
                                                     => x.clone(),
                                                     _ => unreachable!(),
                                                 }.as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                         _,
                                                                                                         _,
                                                                                                         _)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    match _arg.1.clone().as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match _arg.1.clone().as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match match _arg.1.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    })),
                                                                                                              match match _arg.1.clone().as_ref()
                                                                                                                        {
                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                  _,
                                                                                                                                                                                  _,
                                                                                                                                                                                  x)
                                                                                                                        =>
                                                                                                                        x.clone(),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        unreachable!(),
                                                                                                                    }.as_ref()
                                                                                                                  {
                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                            _,
                                                                                                                                                                            x,
                                                                                                                                                                            _)
                                                                                                                  =>
                                                                                                                  x.clone(),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  unreachable!(),
                                                                                                              },
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    match match _arg.1.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                    _arg.3.clone()))))
                                    } else {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                               = _arg.3.clone().as_ref() {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                   =
                                                   match _arg.3.clone().as_ref()
                                                       {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                 _,
                                                                                                                 _,
                                                                                                                 _)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                       =
                                                       match _arg.3.clone().as_ref()
                                                           {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                     x,
                                                                                                                     _,
                                                                                                                     _)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                           =
                                                           match match _arg.3.clone().as_ref()
                                                                     {
                                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                               x,
                                                                                                                               _,
                                                                                                                               _)
                                                                     =>
                                                                     x.clone(),
                                                                     _ =>
                                                                     unreachable!(),
                                                                 }.as_ref() {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         _)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    })),
                                                                                                                              match match _arg.3.clone().as_ref()
                                                                                                                                        {
                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                  x,
                                                                                                                                                                                                  _,
                                                                                                                                                                                                  _)
                                                                                                                                        =>
                                                                                                                                        x.clone(),
                                                                                                                                        _
                                                                                                                                        =>
                                                                                                                                        unreachable!(),
                                                                                                                                    }.as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                            _,
                                                                                                                                                                                            x,
                                                                                                                                                                                            _)
                                                                                                                                  =>
                                                                                                                                  x.clone(),
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  unreachable!(),
                                                                                                                              },
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    }))))
                                                    } else {
                                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                               =
                                                               match _arg.3.clone().as_ref()
                                                                   {
                                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                             _,
                                                                                                                             _,
                                                                                                                             x)
                                                                   =>
                                                                   x.clone(),
                                                                   _ =>
                                                                   unreachable!(),
                                                               }.as_ref() {
                                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                                   =
                                                                   match match _arg.3.clone().as_ref()
                                                                             {
                                                                             PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                       _,
                                                                                                                                       _,
                                                                                                                                       x)
                                                                             =>
                                                                             x.clone(),
                                                                             _
                                                                             =>
                                                                             unreachable!(),
                                                                         }.as_ref()
                                                                       {
                                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                                 _,
                                                                                                                                 _,
                                                                                                                                 _)
                                                                       =>
                                                                       x.clone(),
                                                                       _ =>
                                                                       unreachable!(),
                                                                   }.as_ref()
                                                               {
                                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                            _arg.1.clone(),
                                                                                                                                                                                                            _arg.2.clone(),
                                                                                                                                                                                                            match _arg.3.clone().as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                _
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                            })),
                                                                                                                                      match _arg.3.clone().as_ref()
                                                                                                                                          {
                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                    _,
                                                                                                                                                                                                    x,
                                                                                                                                                                                                    _)
                                                                                                                                          =>
                                                                                                                                          x.clone(),
                                                                                                                                          _
                                                                                                                                          =>
                                                                                                                                          unreachable!(),
                                                                                                                                      },
                                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                _
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                            },
                                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                _
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                            },
                                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                _
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                            }))))
                                                            } else {
                                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                                      _arg.1.clone(),
                                                                                                                                      _arg.2.clone(),
                                                                                                                                      _arg.3.clone()))
                                                            }
                                                        } else {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                                  _arg.1.clone(),
                                                                                                                                  _arg.2.clone(),
                                                                                                                                  _arg.3.clone()))
                                                        }
                                                    }
                                                } else {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                           =
                                                           match _arg.3.clone().as_ref()
                                                               {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         x)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                               =
                                                               match match _arg.3.clone().as_ref()
                                                                         {
                                                                         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                   _,
                                                                                                                                   _,
                                                                                                                                   x)
                                                                         =>
                                                                         x.clone(),
                                                                         _ =>
                                                                         unreachable!(),
                                                                     }.as_ref()
                                                                   {
                                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                             _,
                                                                                                                             _,
                                                                                                                             _)
                                                                   =>
                                                                   x.clone(),
                                                                   _ =>
                                                                   unreachable!(),
                                                               }.as_ref() {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        _arg.1.clone(),
                                                                                                                                                                                                        _arg.2.clone(),
                                                                                                                                                                                                        match _arg.3.clone().as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        })),
                                                                                                                                  match _arg.3.clone().as_ref()
                                                                                                                                      {
                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                _,
                                                                                                                                                                                                x,
                                                                                                                                                                                                _)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      unreachable!(),
                                                                                                                                  },
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        }))))
                                                        } else {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                                  _arg.1.clone(),
                                                                                                                                  _arg.2.clone(),
                                                                                                                                  _arg.3.clone()))
                                                        }
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                }
                                            } else {
                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                      _arg.1.clone(),
                                                                                                                      _arg.2.clone(),
                                                                                                                      _arg.3.clone()))
                                            }
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    }
                                } else {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                           = _arg.3.clone().as_ref() {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                               =
                                               match _arg.3.clone().as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                             _,
                                                                                                             _,
                                                                                                             _)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                   =
                                                   match _arg.3.clone().as_ref()
                                                       {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                 x,
                                                                                                                 _,
                                                                                                                 _)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                       =
                                                       match match _arg.3.clone().as_ref()
                                                                 {
                                                                 PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                           x,
                                                                                                                           _,
                                                                                                                           _)
                                                                 => x.clone(),
                                                                 _ =>
                                                                 unreachable!(),
                                                             }.as_ref() {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     _)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                _arg.1.clone(),
                                                                                                                                                                                                _arg.2.clone(),
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                })),
                                                                                                                          match match _arg.3.clone().as_ref()
                                                                                                                                    {
                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                              x,
                                                                                                                                                                                              _,
                                                                                                                                                                                              _)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                    _
                                                                                                                                    =>
                                                                                                                                    unreachable!(),
                                                                                                                                }.as_ref()
                                                                                                                              {
                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                        _,
                                                                                                                                                                                        x,
                                                                                                                                                                                        _)
                                                                                                                              =>
                                                                                                                              x.clone(),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              unreachable!(),
                                                                                                                          },
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                }))))
                                                } else {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                           =
                                                           match _arg.3.clone().as_ref()
                                                               {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         x)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                               =
                                                               match match _arg.3.clone().as_ref()
                                                                         {
                                                                         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                   _,
                                                                                                                                   _,
                                                                                                                                   x)
                                                                         =>
                                                                         x.clone(),
                                                                         _ =>
                                                                         unreachable!(),
                                                                     }.as_ref()
                                                                   {
                                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                             _,
                                                                                                                             _,
                                                                                                                             _)
                                                                   =>
                                                                   x.clone(),
                                                                   _ =>
                                                                   unreachable!(),
                                                               }.as_ref() {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        _arg.1.clone(),
                                                                                                                                                                                                        _arg.2.clone(),
                                                                                                                                                                                                        match _arg.3.clone().as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        })),
                                                                                                                                  match _arg.3.clone().as_ref()
                                                                                                                                      {
                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                _,
                                                                                                                                                                                                x,
                                                                                                                                                                                                _)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      unreachable!(),
                                                                                                                                  },
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        }))))
                                                        } else {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                                  _arg.1.clone(),
                                                                                                                                  _arg.2.clone(),
                                                                                                                                  _arg.3.clone()))
                                                        }
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                }
                                            } else {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                       =
                                                       match _arg.3.clone().as_ref()
                                                           {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     x)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                           =
                                                           match match _arg.3.clone().as_ref()
                                                                     {
                                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                               _,
                                                                                                                               _,
                                                                                                                               x)
                                                                     =>
                                                                     x.clone(),
                                                                     _ =>
                                                                     unreachable!(),
                                                                 }.as_ref() {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         _)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    })),
                                                                                                                              match _arg.3.clone().as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                            _,
                                                                                                                                                                                            x,
                                                                                                                                                                                            _)
                                                                                                                                  =>
                                                                                                                                  x.clone(),
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  unreachable!(),
                                                                                                                              },
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    }))))
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                } else {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                          _arg.1.clone(),
                                                                                                                          _arg.2.clone(),
                                                                                                                          _arg.3.clone()))
                                                }
                                            }
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                }
                            }
                        } else {
                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                   =
                                   match _arg.1.clone().as_ref() {
                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                 _,
                                                                                                 _,
                                                                                                 x)
                                       => x.clone(),
                                       _ => unreachable!(),
                                   }.as_ref() {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                       =
                                       match match _arg.1.clone().as_ref() {
                                                 PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                           _,
                                                                                                           _,
                                                                                                           x)
                                                 => x.clone(),
                                                 _ => unreachable!(),
                                             }.as_ref() {
                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                     _,
                                                                                                     _,
                                                                                                     _)
                                           => x.clone(),
                                           _ => unreachable!(),
                                       }.as_ref() {
                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                match _arg.1.clone().as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                },
                                                                                                                                                                                match _arg.1.clone().as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                              _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                },
                                                                                                                                                                                match match _arg.1.clone().as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                          =>
                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                })),
                                                                                                          match match _arg.1.clone().as_ref()
                                                                                                                    {
                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                              _,
                                                                                                                                                                              _,
                                                                                                                                                                              x)
                                                                                                                    =>
                                                                                                                    x.clone(),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    unreachable!(),
                                                                                                                }.as_ref()
                                                                                                              {
                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                        _,
                                                                                                                                                                        x,
                                                                                                                                                                        _)
                                                                                                              =>
                                                                                                              x.clone(),
                                                                                                              _
                                                                                                              =>
                                                                                                              unreachable!(),
                                                                                                          },
                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                match match _arg.1.clone().as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                          =>
                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              x)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                },
                                                                                                                                                                                _arg.2.clone(),
                                                                                                                                                                                _arg.3.clone()))))
                                } else {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                           = _arg.3.clone().as_ref() {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                               =
                                               match _arg.3.clone().as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                             _,
                                                                                                             _,
                                                                                                             _)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                   =
                                                   match _arg.3.clone().as_ref()
                                                       {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                 x,
                                                                                                                 _,
                                                                                                                 _)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                       =
                                                       match match _arg.3.clone().as_ref()
                                                                 {
                                                                 PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                           x,
                                                                                                                           _,
                                                                                                                           _)
                                                                 => x.clone(),
                                                                 _ =>
                                                                 unreachable!(),
                                                             }.as_ref() {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     _)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                _arg.1.clone(),
                                                                                                                                                                                                _arg.2.clone(),
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                })),
                                                                                                                          match match _arg.3.clone().as_ref()
                                                                                                                                    {
                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                              x,
                                                                                                                                                                                              _,
                                                                                                                                                                                              _)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                    _
                                                                                                                                    =>
                                                                                                                                    unreachable!(),
                                                                                                                                }.as_ref()
                                                                                                                              {
                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                        _,
                                                                                                                                                                                        x,
                                                                                                                                                                                        _)
                                                                                                                              =>
                                                                                                                              x.clone(),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              unreachable!(),
                                                                                                                          },
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                }))))
                                                } else {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                           =
                                                           match _arg.3.clone().as_ref()
                                                               {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         x)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                               =
                                                               match match _arg.3.clone().as_ref()
                                                                         {
                                                                         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                   _,
                                                                                                                                   _,
                                                                                                                                   x)
                                                                         =>
                                                                         x.clone(),
                                                                         _ =>
                                                                         unreachable!(),
                                                                     }.as_ref()
                                                                   {
                                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                             _,
                                                                                                                             _,
                                                                                                                             _)
                                                                   =>
                                                                   x.clone(),
                                                                   _ =>
                                                                   unreachable!(),
                                                               }.as_ref() {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        _arg.1.clone(),
                                                                                                                                                                                                        _arg.2.clone(),
                                                                                                                                                                                                        match _arg.3.clone().as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        })),
                                                                                                                                  match _arg.3.clone().as_ref()
                                                                                                                                      {
                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                _,
                                                                                                                                                                                                x,
                                                                                                                                                                                                _)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      unreachable!(),
                                                                                                                                  },
                                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        },
                                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                  _
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                            _
                                                                                                                                                                                                            =>
                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                        }))))
                                                        } else {
                                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                                  _arg.1.clone(),
                                                                                                                                  _arg.2.clone(),
                                                                                                                                  _arg.3.clone()))
                                                        }
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                }
                                            } else {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                       =
                                                       match _arg.3.clone().as_ref()
                                                           {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     x)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                           =
                                                           match match _arg.3.clone().as_ref()
                                                                     {
                                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                               _,
                                                                                                                               _,
                                                                                                                               x)
                                                                     =>
                                                                     x.clone(),
                                                                     _ =>
                                                                     unreachable!(),
                                                                 }.as_ref() {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         _)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    })),
                                                                                                                              match _arg.3.clone().as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                            _,
                                                                                                                                                                                            x,
                                                                                                                                                                                            _)
                                                                                                                                  =>
                                                                                                                                  x.clone(),
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  unreachable!(),
                                                                                                                              },
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    }))))
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                } else {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                          _arg.1.clone(),
                                                                                                                          _arg.2.clone(),
                                                                                                                          _arg.3.clone()))
                                                }
                                            }
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                }
                            } else {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                       = _arg.3.clone().as_ref() {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                           =
                                           match _arg.3.clone().as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                         _,
                                                                                                         _,
                                                                                                         _)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                               =
                                               match _arg.3.clone().as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                             x,
                                                                                                             _,
                                                                                                             _)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                   =
                                                   match match _arg.3.clone().as_ref()
                                                             {
                                                             PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                       x,
                                                                                                                       _,
                                                                                                                       _)
                                                             => x.clone(),
                                                             _ =>
                                                             unreachable!(),
                                                         }.as_ref() {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                 _,
                                                                                                                 _,
                                                                                                                 _)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                            _arg.1.clone(),
                                                                                                                                                                                            _arg.2.clone(),
                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            })),
                                                                                                                      match match _arg.3.clone().as_ref()
                                                                                                                                {
                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                          x,
                                                                                                                                                                                          _,
                                                                                                                                                                                          _)
                                                                                                                                =>
                                                                                                                                x.clone(),
                                                                                                                                _
                                                                                                                                =>
                                                                                                                                unreachable!(),
                                                                                                                            }.as_ref()
                                                                                                                          {
                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                    _,
                                                                                                                                                                                    x,
                                                                                                                                                                                    _)
                                                                                                                          =>
                                                                                                                          x.clone(),
                                                                                                                          _
                                                                                                                          =>
                                                                                                                          unreachable!(),
                                                                                                                      },
                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            },
                                                                                                                                                                                            match _arg.3.clone().as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            },
                                                                                                                                                                                            match _arg.3.clone().as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            }))))
                                            } else {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                       =
                                                       match _arg.3.clone().as_ref()
                                                           {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     x)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                           =
                                                           match match _arg.3.clone().as_ref()
                                                                     {
                                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                               _,
                                                                                                                               _,
                                                                                                                               x)
                                                                     =>
                                                                     x.clone(),
                                                                     _ =>
                                                                     unreachable!(),
                                                                 }.as_ref() {
                                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                         _,
                                                                                                                         _,
                                                                                                                         _)
                                                               => x.clone(),
                                                               _ =>
                                                               unreachable!(),
                                                           }.as_ref() {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    })),
                                                                                                                              match _arg.3.clone().as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                            _,
                                                                                                                                                                                            x,
                                                                                                                                                                                            _)
                                                                                                                                  =>
                                                                                                                                  x.clone(),
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  unreachable!(),
                                                                                                                              },
                                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    },
                                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                                        {
                                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                        _
                                                                                                                                                                                                        =>
                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                    }))))
                                                    } else {
                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                              _arg.1.clone(),
                                                                                                                              _arg.2.clone(),
                                                                                                                              _arg.3.clone()))
                                                    }
                                                } else {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                          _arg.1.clone(),
                                                                                                                          _arg.2.clone(),
                                                                                                                          _arg.3.clone()))
                                                }
                                            }
                                        } else {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                                   =
                                                   match _arg.3.clone().as_ref()
                                                       {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                 _,
                                                                                                                 _,
                                                                                                                 x)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                       =
                                                       match match _arg.3.clone().as_ref()
                                                                 {
                                                                 PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                           _,
                                                                                                                           _,
                                                                                                                           x)
                                                                 => x.clone(),
                                                                 _ =>
                                                                 unreachable!(),
                                                             }.as_ref() {
                                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                     _,
                                                                                                                     _,
                                                                                                                     _)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       }.as_ref() {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                _arg.1.clone(),
                                                                                                                                                                                                _arg.2.clone(),
                                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                })),
                                                                                                                          match _arg.3.clone().as_ref()
                                                                                                                              {
                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                        _,
                                                                                                                                                                                        x,
                                                                                                                                                                                        _)
                                                                                                                              =>
                                                                                                                              x.clone(),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              unreachable!(),
                                                                                                                          },
                                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                },
                                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                }))))
                                                } else {
                                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                          _arg.1.clone(),
                                                                                                                          _arg.2.clone(),
                                                                                                                          _arg.3.clone()))
                                                }
                                            } else {
                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                      _arg.1.clone(),
                                                                                                                      _arg.2.clone(),
                                                                                                                      _arg.3.clone()))
                                            }
                                        }
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                } else {
                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                          _arg.1.clone(),
                                                                                                          _arg.2.clone(),
                                                                                                          _arg.3.clone()))
                                }
                            }
                        }
                    } else {
                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                               = _arg.3.clone().as_ref() {
                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                   =
                                   match _arg.3.clone().as_ref() {
                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                 _,
                                                                                                 _,
                                                                                                 _)
                                       => x.clone(),
                                       _ => unreachable!(),
                                   }.as_ref() {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                       =
                                       match _arg.3.clone().as_ref() {
                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                     x,
                                                                                                     _,
                                                                                                     _)
                                           => x.clone(),
                                           _ => unreachable!(),
                                       }.as_ref() {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                           =
                                           match match _arg.3.clone().as_ref()
                                                     {
                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                               x,
                                                                                                               _,
                                                                                                               _)
                                                     => x.clone(),
                                                     _ => unreachable!(),
                                                 }.as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                         _,
                                                                                                         _,
                                                                                                         _)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    })),
                                                                                                              match match _arg.3.clone().as_ref()
                                                                                                                        {
                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                  x,
                                                                                                                                                                                  _,
                                                                                                                                                                                  _)
                                                                                                                        =>
                                                                                                                        x.clone(),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        unreachable!(),
                                                                                                                    }.as_ref()
                                                                                                                  {
                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                            _,
                                                                                                                                                                            x,
                                                                                                                                                                            _)
                                                                                                                  =>
                                                                                                                  x.clone(),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  unreachable!(),
                                                                                                              },
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    }))))
                                    } else {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                               =
                                               match _arg.3.clone().as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                             _,
                                                                                                             _,
                                                                                                             x)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                                   =
                                                   match match _arg.3.clone().as_ref()
                                                             {
                                                             PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                       _,
                                                                                                                       _,
                                                                                                                       x)
                                                             => x.clone(),
                                                             _ =>
                                                             unreachable!(),
                                                         }.as_ref() {
                                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                                 _,
                                                                                                                 _,
                                                                                                                 _)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   }.as_ref() {
                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                            _arg.1.clone(),
                                                                                                                                                                                            _arg.2.clone(),
                                                                                                                                                                                            match _arg.3.clone().as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            })),
                                                                                                                      match _arg.3.clone().as_ref()
                                                                                                                          {
                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                    _,
                                                                                                                                                                                    x,
                                                                                                                                                                                    _)
                                                                                                                          =>
                                                                                                                          x.clone(),
                                                                                                                          _
                                                                                                                          =>
                                                                                                                          unreachable!(),
                                                                                                                      },
                                                                                                                      LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            },
                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            },
                                                                                                                                                                                            match match _arg.3.clone().as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                  }.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            }))))
                                            } else {
                                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                      _arg.1.clone(),
                                                                                                                      _arg.2.clone(),
                                                                                                                      _arg.3.clone()))
                                            }
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    }
                                } else {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                           =
                                           match _arg.3.clone().as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                         _,
                                                                                                         _,
                                                                                                         x)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                               =
                                               match match _arg.3.clone().as_ref()
                                                         {
                                                         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                   _,
                                                                                                                   _,
                                                                                                                   x)
                                                         => x.clone(),
                                                         _ => unreachable!(),
                                                     }.as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                             _,
                                                                                                             _,
                                                                                                             _)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                        _arg.1.clone(),
                                                                                                                                                                                        _arg.2.clone(),
                                                                                                                                                                                        match _arg.3.clone().as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        })),
                                                                                                                  match _arg.3.clone().as_ref()
                                                                                                                      {
                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                _,
                                                                                                                                                                                x,
                                                                                                                                                                                _)
                                                                                                                      =>
                                                                                                                      x.clone(),
                                                                                                                      _
                                                                                                                      =>
                                                                                                                      unreachable!(),
                                                                                                                  },
                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        },
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        },
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        }))))
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                }
                            } else {
                                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                      _arg.1.clone(),
                                                                                                      _arg.2.clone(),
                                                                                                      _arg.3.clone()))
                            }
                        } else {
                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                  _arg.1.clone(),
                                                                                                  _arg.2.clone(),
                                                                                                  _arg.3.clone()))
                        }
                    }
                } else {
                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                           = _arg.3.clone().as_ref() {
                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                               =
                               match _arg.3.clone().as_ref() {
                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                             _,
                                                                                             _,
                                                                                             _)
                                   => x.clone(),
                                   _ => unreachable!(),
                               }.as_ref() {
                            if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                   =
                                   match _arg.3.clone().as_ref() {
                                       PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                 x,
                                                                                                 _,
                                                                                                 _)
                                       => x.clone(),
                                       _ => unreachable!(),
                                   }.as_ref() {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                       =
                                       match match _arg.3.clone().as_ref() {
                                                 PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                           x,
                                                                                                           _,
                                                                                                           _)
                                                 => x.clone(),
                                                 _ => unreachable!(),
                                             }.as_ref() {
                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                     _,
                                                                                                     _,
                                                                                                     _)
                                           => x.clone(),
                                           _ => unreachable!(),
                                       }.as_ref() {
                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                _arg.1.clone(),
                                                                                                                                                                                _arg.2.clone(),
                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                          =>
                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                })),
                                                                                                          match match _arg.3.clone().as_ref()
                                                                                                                    {
                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                              x,
                                                                                                                                                                              _,
                                                                                                                                                                              _)
                                                                                                                    =>
                                                                                                                    x.clone(),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    unreachable!(),
                                                                                                                }.as_ref()
                                                                                                              {
                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                        _,
                                                                                                                                                                        x,
                                                                                                                                                                        _)
                                                                                                              =>
                                                                                                              x.clone(),
                                                                                                              _
                                                                                                              =>
                                                                                                              unreachable!(),
                                                                                                          },
                                                                                                          LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                match match _arg.3.clone().as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                    x,
                                                                                                                                                                                                                                                    _,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                          =>
                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                      }.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              x)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                },
                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                              _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                },
                                                                                                                                                                                match _arg.3.clone().as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                              x)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                }))))
                                } else {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                           =
                                           match _arg.3.clone().as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                         _,
                                                                                                         _,
                                                                                                         x)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                               =
                                               match match _arg.3.clone().as_ref()
                                                         {
                                                         PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                   _,
                                                                                                                   _,
                                                                                                                   x)
                                                         => x.clone(),
                                                         _ => unreachable!(),
                                                     }.as_ref() {
                                                   PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                             _,
                                                                                                             _,
                                                                                                             _)
                                                   => x.clone(),
                                                   _ => unreachable!(),
                                               }.as_ref() {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                        _arg.1.clone(),
                                                                                                                                                                                        _arg.2.clone(),
                                                                                                                                                                                        match _arg.3.clone().as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        })),
                                                                                                                  match _arg.3.clone().as_ref()
                                                                                                                      {
                                                                                                                      PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                _,
                                                                                                                                                                                x,
                                                                                                                                                                                _)
                                                                                                                      =>
                                                                                                                      x.clone(),
                                                                                                                      _
                                                                                                                      =>
                                                                                                                      unreachable!(),
                                                                                                                  },
                                                                                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        },
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        },
                                                                                                                                                                                        match match _arg.3.clone().as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      _,
                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        }))))
                                        } else {
                                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                                  _arg.1.clone(),
                                                                                                                  _arg.2.clone(),
                                                                                                                  _arg.3.clone()))
                                        }
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                }
                            } else {
                                if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_)
                                       =
                                       match _arg.3.clone().as_ref() {
                                           PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                     _,
                                                                                                     _,
                                                                                                     x)
                                           => x.clone(),
                                           _ => unreachable!(),
                                       }.as_ref() {
                                    if let PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red
                                           =
                                           match match _arg.3.clone().as_ref()
                                                     {
                                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                               _,
                                                                                                               _,
                                                                                                               x)
                                                     => x.clone(),
                                                     _ => unreachable!(),
                                                 }.as_ref() {
                                               PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(x,
                                                                                                         _,
                                                                                                         _,
                                                                                                         _)
                                               => x.clone(),
                                               _ => unreachable!(),
                                           }.as_ref() {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    _arg.1.clone(),
                                                                                                                                                                                    _arg.2.clone(),
                                                                                                                                                                                    match _arg.3.clone().as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    })),
                                                                                                              match _arg.3.clone().as_ref()
                                                                                                                  {
                                                                                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                            _,
                                                                                                                                                                            x,
                                                                                                                                                                            _)
                                                                                                                  =>
                                                                                                                  x.clone(),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  unreachable!(),
                                                                                                              },
                                                                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x,
                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    },
                                                                                                                                                                                    match match _arg.3.clone().as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    }))))
                                    } else {
                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                              _arg.1.clone(),
                                                                                                              _arg.2.clone(),
                                                                                                              _arg.3.clone()))
                                    }
                                } else {
                                    LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                          _arg.1.clone(),
                                                                                                          _arg.2.clone(),
                                                                                                          _arg.3.clone()))
                                }
                            }
                        } else {
                            LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                                  _arg.1.clone(),
                                                                                                  _arg.2.clone(),
                                                                                                  _arg.3.clone()))
                        }
                    } else {
                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                              _arg.1.clone(),
                                                                                              _arg.2.clone(),
                                                                                              _arg.3.clone()))
                    }
                }
            } else {
                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg.0.clone(),
                                                                                      _arg.1.clone(),
                                                                                      _arg.2.clone(),
                                                                                      _arg.3.clone()))
            }
        }
        pub fn insert(x: i32,
                      t:
                          LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>)
         -> LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> {
            let ins =
                Func1::new({
                               let x = x.clone();
                               move
                                   |_arg:
                                        LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>|
                                   fix1(&(move
                                              |ins,
                                               _arg:
                                                   LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>|
                                              match _arg.as_ref() {
                                                  PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg_1_0,
                                                                                                            _arg_1_1,
                                                                                                            _arg_1_2,
                                                                                                            _arg_1_3)
                                                  => {
                                                      let y: i32 =
                                                          _arg_1_2.clone();
                                                      let c:
                                                              LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color> =
                                                          _arg_1_0.clone();
                                                      let b:
                                                              LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> =
                                                          _arg_1_3.clone();
                                                      let a:
                                                              LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> =
                                                          _arg_1_1.clone();
                                                      if x < y {
                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::balance(c.clone(),
                                                                                                                 ins(a.clone()),
                                                                                                                 y,
                                                                                                                 b.clone())
                                                      } else {
                                                          if x > y {
                                                              PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::balance(c.clone(),
                                                                                                                     a.clone(),
                                                                                                                     y,
                                                                                                                     ins(b.clone()))
                                                          } else {
                                                              LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(c,
                                                                                                                                    a,
                                                                                                                                    y,
                                                                                                                                    b))
                                                          }
                                                      }
                                                  }
                                                  _ =>
                                                  LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Red),
                                                                                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Empty),
                                                                                                                        x,
                                                                                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Empty))),
                                              }), _arg.clone())
                           });
            let matchValue:
                    LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> =
                ins(t);
            match matchValue.as_ref() {
                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Empty =>
                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Empty),
                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(matchValue_1_0,
                                                                          matchValue_1_1,
                                                                          matchValue_1_2,
                                                                          matchValue_1_3)
                =>
                LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Color::Black),
                                                                                      match matchValue.as_ref()
                                                                                          {
                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                    x,
                                                                                                                                                    _,
                                                                                                                                                    _)
                                                                                          =>
                                                                                          x.clone(),
                                                                                          _
                                                                                          =>
                                                                                          unreachable!(),
                                                                                      },
                                                                                      match matchValue.as_ref()
                                                                                          {
                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                    _,
                                                                                                                                                    x,
                                                                                                                                                    _)
                                                                                          =>
                                                                                          x.clone(),
                                                                                          _
                                                                                          =>
                                                                                          unreachable!(),
                                                                                      },
                                                                                      match matchValue.as_ref()
                                                                                          {
                                                                                          PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_,
                                                                                                                                                    _,
                                                                                                                                                    _,
                                                                                                                                                    x)
                                                                                          =>
                                                                                          x.clone(),
                                                                                          _
                                                                                          =>
                                                                                          unreachable!(),
                                                                                      })),
            }
        }
        pub fn depth(_arg:
                         LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>)
         -> i32 {
            match _arg.as_ref() {
                PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Node(_arg_1_0,
                                                                          _arg_1_1,
                                                                          _arg_1_2,
                                                                          _arg_1_3)
                =>
                1_i32 +
                    PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::depth(_arg_1_1.clone()).max(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::depth(_arg_1_3.clone())),
                _ => 0_i32,
            }
        }
        pub fn runRBTreeFFI(n: &dyn Any) -> &dyn Any {
            fn build(i: i32,
                     t:
                         LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>)
             -> LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> {
                let i: MutCell<i32> = MutCell::new(i);
                let t:
                        MutCell<LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree>> =
                    MutCell::new(t.clone());
                '_build:
                    loop  {
                        break '_build
                            (if i.get() == 0_i32 {
                                 t.get()
                             } else {
                                 let i_temp: i32 = i.get() - 1_i32;
                                 let t_temp:
                                         LrcPtr<PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree> =
                                     PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::insert(i.get(),
                                                                                           t.get());
                                 i.set(i_temp);
                                 t.set(t_temp);
                                 continue '_build
                             }) ;
                    }
            }
            &PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::depth(build(Sharpurs_Prelude::unbox(n),
                                                                        LrcPtr::new(PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::Tree::Empty)))
        }
    }
    pub fn Test_RBTreeFFI_runRBTreeFFI() -> &dyn Any {
        static Test_RBTreeFFI_runRBTreeFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFI_runRBTreeFFI.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &PureScript_Test_RBTreeFFI::Test_RBTreeFFI_FFI::runRBTreeFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_RBTreeFFI_describe() -> &dyn Any {
        static Test_RBTreeFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFI_describe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                 &&&string("Red-Black Tree FFI (100k Worst-Case Insertions):")))
    }
    pub fn Test_RBTreeFFI_act() -> &dyn Any {
        static Test_RBTreeFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RBTreeFFI_act.get_or_init(||
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
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RBTreeFFI::Test_RBTreeFFI_runRBTreeFFI(),
                                                                                                                                                                                                     dummy))))))
    }
}
