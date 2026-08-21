codegenPrelude :: Array (Module Ann) -> Set.Set String -> String
codegenPrelude modules allFields =
  let
    allAdtsStr = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let adtName = "Adt_" <> modPrefix <> sanitizeIdent decl.name
               enumName = adtName <> "_enum"
               
               ctorsStr = Array.foldMap (\ctor ->
                 let ctorName = sanitizeIdent ctor.name
                     fieldsStr = if Array.length ctor.fields == 0 then "" else "(" <> String.joinWith ", " (Array.replicate (Array.length ctor.fields) "crate::Value") <> ")"
                 in "    " <> ctorName <> fieldsStr <> ",\n"
               ) decl.constructors
               
           in "#[derive(Clone)]\npub enum " <> enumName <> " {\n" <> ctorsStr <> "}\n" <>
              "pub type " <> adtName <> " = perceus_ptr::PerceusPtr<" <> enumName <> ">;\n\n"
         ) m.dataDecls
    ) modules

    allClassesStr = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let className = "Adt_" <> modPrefix <> sanitizeIdent decl.name
               
               superNames = Array.mapWithIndex (\i (Tuple fqn _) -> 
                 Tuple ((case Array.last fqn of
                    Just sc -> sc
                    Nothing -> "Super") <> show i) Any
               ) decl.superclasses
               methodNames = map (\(Tuple mName mTy) -> Tuple (sanitizeIdent mName) mTy) decl.methods
               allFieldsArr = Array.concat [superNames, methodNames]
               
               fieldsStr = Array.foldMap (\(Tuple f _) ->
                 "    pub " <> sanitizeIdent f <> ": crate::Value,\n"
               ) allFieldsArr
               
           in "#[derive(Clone)]\npub struct " <> className <> "Dict {\n" <> fieldsStr <> "}\n" <>
              "pub type " <> className <> " = perceus_ptr::PerceusPtr<" <> className <> "Dict>;\n\n"
         ) m.classDecls
    ) modules

    valueVariantsAdts = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let adtName = "Adt_" <> modPrefix <> sanitizeIdent decl.name
           in "    " <> adtName <> "(" <> adtName <> "),\n"
         ) m.dataDecls
    ) modules
    valueVariantsClasses = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let className = "Adt_" <> modPrefix <> sanitizeIdent decl.name
           in "    " <> className <> "(" <> className <> "),\n"
         ) m.classDecls
    ) modules

    valueUnwrapsAdts = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let adtName = "Adt_" <> modPrefix <> sanitizeIdent decl.name
           in "    pub fn unwrap_" <> adtName <> "(&self) -> " <> adtName <> " {\n" <>
              "        if let Value::" <> adtName <> "(v) = self { v.clone() } else { panic!(\"Expected " <> adtName <> " but got another type\"); }\n" <>
              "    }\n" <>
              "    pub fn as_" <> adtName <> "_mut(&mut self) -> &mut " <> adtName <> " {\n" <>
              "        if let Value::" <> adtName <> "(v) = self { v } else { panic!(\"Expected " <> adtName <> "\"); }\n" <>
              "    }\n"
         ) m.dataDecls
    ) modules
    valueUnwrapsClasses = Array.foldMap (\(Module m) ->
      let modPrefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name) <> "_"
      in Array.foldMap (\decl ->
           let className = "Adt_" <> modPrefix <> sanitizeIdent decl.name
           in "    pub fn unwrap_" <> className <> "(&self) -> " <> className <> " {\n" <>
              "        if let Value::" <> className <> "(v) = self { v.clone() } else { panic!(\"Expected " <> className <> "\"); }\n" <>
              "    }\n"
         ) m.classDecls
    ) modules

  in
    "#![allow(warnings)]\n\n" <>
    "use perceus_ptr::PerceusPtr;\n\n" <>
    allAdtsStr <>
    allClassesStr <>
    "#[derive(Clone)]\n" <>
    "pub enum Value {\n" <>
    "    Int(i64),\n" <>
    "    Number(f64),\n" <>
    "    Bool(bool),\n" <>
    "    String(String),\n" <>
    "    Char(char),\n" <>
    "    Array(std::rc::Rc<Vec<crate::Value>>),\n" <>
    "    Func(std::rc::Rc<dyn Fn(crate::Value) -> crate::Value>),\n" <>
    "    Record(perceus_ptr::PerceusPtr<Record_a>),\n" <>
    valueVariantsAdts <>
    valueVariantsClasses <>
    "}\n\n" <>
    "pub type UnknownType = Value;\n\n" <>
    "impl Value {\n" <>
    "    pub fn unwrap_int(&self) -> i64 {\n" <>
    "        if let Value::Int(v) = self { *v } else { panic!(\"Expected Int\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_number(&self) -> f64 {\n" <>
    "        if let Value::Number(v) = self { *v } else { panic!(\"Expected Number\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_bool(&self) -> bool {\n" <>
    "        if let Value::Bool(v) = self { *v } else { panic!(\"Expected Bool\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_string(&self) -> String {\n" <>
    "        if let Value::String(v) = self { v.clone() } else { panic!(\"Expected String\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_char(&self) -> char {\n" <>
    "        if let Value::Char(v) = self { *v } else { panic!(\"Expected Char\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_array(&self) -> std::rc::Rc<Vec<crate::Value>> {\n" <>
    "        if let Value::Array(v) = self { v.clone() } else { panic!(\"Expected Array\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_func(&self) -> std::rc::Rc<dyn Fn(crate::Value) -> crate::Value> {\n" <>
    "        if let Value::Func(v) = self { v.clone() } else if let Value::Record(v) = self { v.call.clone().unwrap() } else { panic!(\"Expected Func\"); }\n" <>
    "    }\n" <>
    "    pub fn unwrap_record(&self) -> perceus_ptr::PerceusPtr<Record_a> {\n" <>
    "        if let Value::Record(v) = self { v.clone() } else { panic!(\"Expected Record\"); }\n" <>
    "    }\n" <>
    "    pub fn as_record_mut(&mut self) -> &mut perceus_ptr::PerceusPtr<Record_a> {\n" <>
    "        if let Value::Record(v) = self { v } else { panic!(\"Expected Record\"); }\n" <>
    "    }\n" <>
    "    pub fn drop_explicit(self) {\n" <>
    "        if let Value::Record(v) = self { v.drop_explicit(); }\n" <>
    "    }\n" <>
    "    pub fn new(r: Record_a) -> Self {\n" <>
    "        Value::Record(perceus_ptr::PerceusPtr::new(r))\n" <>
    "    }\n" <>
    valueUnwrapsAdts <>
    valueUnwrapsClasses <>
    "}\n\n" <>
    "pub fn mk_int(val: i64) -> UnknownType { Value::Int(val) }\n" <>
    "pub fn mk_bool(val: bool) -> UnknownType { Value::Bool(val) }\n" <>
    "pub fn mk_number(val: f64) -> UnknownType { Value::Number(val) }\n" <>
    "pub fn mk_string(val: &str) -> UnknownType { Value::String(val.to_string()) }\n" <>
    "pub fn mk_char(val: char) -> UnknownType { Value::Char(val) }\n" <>
    "pub fn mk_array(val: Vec<UnknownType>) -> UnknownType { Value::Array(std::rc::Rc::new(val)) }\n\n" <>
    "#[derive(Clone, Default)]\npub struct Record_a {\n" <>
    "    pub tag: &'static str,\n" <>
    "    pub vals: Option<std::rc::Rc<Vec<UnknownType>>>,\n" <>
    "    pub call: Option<std::rc::Rc<dyn Fn(UnknownType) -> UnknownType>>,\n" <>
    Array.foldMap (\field ->
      let
        ignore = Set.fromFoldable ["unwrap", "clone", "as_ref", "tag", "vals", "call"]
      in if Set.member field ignore then "" else "    pub " <> sanitizeIdent field <> ": Option<UnknownType>,\n"
    ) (Array.fromFoldable allFields) <>
    "}\n\n"
