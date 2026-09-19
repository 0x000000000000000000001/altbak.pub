pub mod PureScript_Control_Parallel {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_255d4e69::PureScript_Control_Parallel_Class;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Parallel_identity() -> &dyn Any {
        static Control_Parallel_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_identity.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Control_Parallel_parTraverse_() -> &dyn Any {
        static Control_Parallel_parTraverse_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parTraverse_.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictParallel|
                                                                      {
                                                                          let sequential =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_sequential(),
                                                                                                               dictParallel);
                                                                          let parallel =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                               dictParallel);
                                                                          &Func1::new({
                                                                                          let parallel
                                                                                              =
                                                                                              parallel.clone();
                                                                                          let sequential
                                                                                              =
                                                                                              sequential.clone();
                                                                                          move
                                                                                              |dictApplicative|
                                                                                              &Func1::new({
                                                                                                              let dictApplicative
                                                                                                                  =
                                                                                                                  dictApplicative.clone();
                                                                                                              move
                                                                                                                  |dictFoldable|
                                                                                                                  &Func1::new({
                                                                                                                                  let dictFoldable
                                                                                                                                      =
                                                                                                                                      dictFoldable.clone();
                                                                                                                                  move
                                                                                                                                      |f|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                          &&&sequential),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_traverse_(),
                                                                                                                                                                                                                                                                                &&&dictApplicative),
                                                                                                                                                                                                                                             &&&dictFoldable),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                &&&parallel),
                                                                                                                                                                                                                                             f)))
                                                                                                                              })
                                                                                                          })
                                                                                      })
                                                                      }))
    }
    pub fn Control_Parallel_parTraverse() -> &dyn Any {
        static Control_Parallel_parTraverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parTraverse.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictParallel|
                                                                     {
                                                                         let sequential =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_sequential(),
                                                                                                              dictParallel);
                                                                         let parallel =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                              dictParallel);
                                                                         &Func1::new({
                                                                                         let parallel
                                                                                             =
                                                                                             parallel.clone();
                                                                                         let sequential
                                                                                             =
                                                                                             sequential.clone();
                                                                                         move
                                                                                             |dictApplicative|
                                                                                             &Func1::new({
                                                                                                             let dictApplicative
                                                                                                                 =
                                                                                                                 dictApplicative.clone();
                                                                                                             move
                                                                                                                 |dictTraversable|
                                                                                                                 &Func1::new({
                                                                                                                                 let dictTraversable
                                                                                                                                     =
                                                                                                                                     dictTraversable.clone();
                                                                                                                                 move
                                                                                                                                     |f|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                         &&&sequential),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                               &&&dictTraversable),
                                                                                                                                                                                                                                            &&&dictApplicative),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                               &&&parallel),
                                                                                                                                                                                                                                            f)))
                                                                                                                             })
                                                                                                         })
                                                                                     })
                                                                     }))
    }
    pub fn Control_Parallel_parSequence_() -> &dyn Any {
        static Control_Parallel_parSequence_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parSequence_.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictParallel|
                                                                      &Func1::new({
                                                                                      let dictParallel
                                                                                          =
                                                                                          dictParallel.clone();
                                                                                      move
                                                                                          |dictApplicative|
                                                                                          &Func1::new({
                                                                                                          let dictApplicative
                                                                                                              =
                                                                                                              dictApplicative.clone();
                                                                                                          move
                                                                                                              |dictFoldable|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel::Control_Parallel_parTraverse_(),
                                                                                                                                                                                                                                                        &&&dictParallel),
                                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                                  dictFoldable),
                                                                                                                                               &&&PureScript_Control_Parallel::Control_Parallel_identity())
                                                                                                      })
                                                                                  })))
    }
    pub fn Control_Parallel_parSequence() -> &dyn Any {
        static Control_Parallel_parSequence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parSequence.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictParallel|
                                                                     &Func1::new({
                                                                                     let dictParallel
                                                                                         =
                                                                                         dictParallel.clone();
                                                                                     move
                                                                                         |dictApplicative|
                                                                                         &Func1::new({
                                                                                                         let dictApplicative
                                                                                                             =
                                                                                                             dictApplicative.clone();
                                                                                                         move
                                                                                                             |dictTraversable|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel::Control_Parallel_parTraverse(),
                                                                                                                                                                                                                                                       &&&dictParallel),
                                                                                                                                                                                                                    &&&dictApplicative),
                                                                                                                                                                                 dictTraversable),
                                                                                                                                              &&&PureScript_Control_Parallel::Control_Parallel_identity())
                                                                                                     })
                                                                                 })))
    }
    pub fn Control_Parallel_parOneOfMap() -> &dyn Any {
        static Control_Parallel_parOneOfMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parOneOfMap.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictParallel|
                                                                     {
                                                                         let sequential =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_sequential(),
                                                                                                              dictParallel);
                                                                         let parallel =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                              dictParallel);
                                                                         &Func1::new({
                                                                                         let parallel
                                                                                             =
                                                                                             parallel.clone();
                                                                                         let sequential
                                                                                             =
                                                                                             sequential.clone();
                                                                                         move
                                                                                             |dictAlternative|
                                                                                             {
                                                                                                 let Plus1 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                 &Func1::new({
                                                                                                                 let Plus1
                                                                                                                     =
                                                                                                                     Plus1.clone();
                                                                                                                 move
                                                                                                                     |dictFoldable|
                                                                                                                     &Func1::new({
                                                                                                                                     let dictFoldable
                                                                                                                                         =
                                                                                                                                         dictFoldable.clone();
                                                                                                                                     move
                                                                                                                                         |dictFunctor|
                                                                                                                                         &Func1::new(move
                                                                                                                                                         |f|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                             &&&sequential),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_oneOfMap(),
                                                                                                                                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                                                                                                                                &&&Plus1),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                   &&&parallel),
                                                                                                                                                                                                                                                                f))))
                                                                                                                                 })
                                                                                                             })
                                                                                             }
                                                                                     })
                                                                     }))
    }
    pub fn Control_Parallel_parOneOf() -> &dyn Any {
        static Control_Parallel_parOneOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parOneOf.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictParallel|
                                                                  {
                                                                      let sequential =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_sequential(),
                                                                                                           dictParallel);
                                                                      let parallel =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                           dictParallel);
                                                                      &Func1::new({
                                                                                      let parallel
                                                                                          =
                                                                                          parallel.clone();
                                                                                      let sequential
                                                                                          =
                                                                                          sequential.clone();
                                                                                      move
                                                                                          |dictAlternative|
                                                                                          {
                                                                                              let Plus1 =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                          Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                                              &Func1::new({
                                                                                                              let Plus1
                                                                                                                  =
                                                                                                                  Plus1.clone();
                                                                                                              move
                                                                                                                  |dictFoldable|
                                                                                                                  &Func1::new({
                                                                                                                                  let dictFoldable
                                                                                                                                      =
                                                                                                                                      dictFoldable.clone();
                                                                                                                                  move
                                                                                                                                      |dictFunctor|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                          &&&sequential),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_oneOfMap(),
                                                                                                                                                                                                                                                                                &&&dictFoldable),
                                                                                                                                                                                                                                             &&&Plus1),
                                                                                                                                                                                                          &&&parallel))
                                                                                                                              })
                                                                                                          })
                                                                                          }
                                                                                  })
                                                                  }))
    }
    pub fn Control_Parallel_parApply() -> &dyn Any {
        static Control_Parallel_parApply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Parallel_parApply.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictParallel|
                                                                  {
                                                                      let Apply1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply1"),
                                                                                                                  Sharpurs_Prelude::unbox(dictParallel)),
                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                      &Func1::new({
                                                                                      let Apply1
                                                                                          =
                                                                                          Apply1.clone();
                                                                                      let dictParallel
                                                                                          =
                                                                                          dictParallel.clone();
                                                                                      move
                                                                                          |mf|
                                                                                          &Func1::new({
                                                                                                          let mf
                                                                                                              =
                                                                                                              mf.clone();
                                                                                                          move
                                                                                                              |ma|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_sequential(),
                                                                                                                                                                                  &&&dictParallel),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                        &&&Apply1),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                                                                                                                                                                                                           &&&dictParallel),
                                                                                                                                                                                                                                                        &&&mf)),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Parallel_Class::Control_Parallel_Class_parallel(),
                                                                                                                                                                                                                                                        &&&dictParallel),
                                                                                                                                                                                                                     ma)))
                                                                                                      })
                                                                                  })
                                                                  }))
    }
}
