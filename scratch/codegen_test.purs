module Scratch.CodeGenTest where

import Prelude
import Data.Maybe (Maybe(..))
import Data.String as String

data ExprType = Int | Boolean | Number | String | Char | Func ExprType ExprType | Any | ForAll String ExprType | ConstrainedType (Array String) ExprType

unwrapType :: ExprType -> ExprType
unwrapType (ForAll _ t) = unwrapType t
unwrapType (ConstrainedType _ t) = unwrapType t
unwrapType t = t

codegenExprType :: Boolean -> ExprType -> String
codegenExprType isRet ty = case unwrapType ty of
  Int -> "i64"
  Boolean -> "bool"
  Number -> "f64"
  String -> "String"
  Char -> "char"
  Func arg ret -> "std::rc::Rc<dyn Fn(" <> codegenExprType false arg <> ") -> " <> codegenExprType true ret <> ">"
  _ -> "crate::Value"

boxUnbox :: ExprType -> ExprType -> String -> String
boxUnbox expected actual code =
  let
    expStr = codegenExprType true expected
    actStr = codegenExprType true actual
  in
    if expStr == actStr then code
    else case unwrapType expected, unwrapType actual of
      Func expArg expRet, Func actArg actRet ->
        "std::rc::Rc::new({ let _f = (" <> code <> ").clone(); move |mut _a: " <> codegenExprType false expArg <> "| -> " <> codegenExprType true expRet <> " { " <> boxUnbox expRet actRet ("_f(" <> boxUnbox actArg expArg "_a" <> ")") <> " } })"
      
      Func expArg expRet, _ ->
        if actStr == "crate::UnknownType" || actStr == "crate::Value" then
          "std::rc::Rc::new({ let _f = (" <> code <> ").unwrap_func(); move |mut _a: " <> codegenExprType false expArg <> "| -> " <> codegenExprType true expRet <> " { " <> boxUnbox expRet Any ("_f(" <> boxUnbox Any expArg "_a" <> ")") <> " } })"
        else code

      _, Func actArg actRet ->
        if expStr == "crate::UnknownType" || expStr == "crate::Value" then
          "crate::Value::Func(std::rc::Rc::new({ let _f = (" <> code <> ").clone(); move |mut _a: crate::Value| -> crate::Value { " <> boxUnbox Any actRet ("_f(" <> boxUnbox actArg Any "_a" <> ")") <> " } }))"
        else code

      _, _ ->
        if expStr == "i64" && (actStr == "crate::UnknownType" || actStr == "crate::Value") then "(" <> code <> ").unwrap_int()"
        else if (expStr == "crate::UnknownType" || expStr == "crate::Value") && actStr == "i64" then "crate::Value::Int(" <> code <> ")"
        else if expStr == "bool" && (actStr == "crate::UnknownType" || actStr == "crate::Value") then "(" <> code <> ").unwrap_bool()"
        else if (expStr == "crate::UnknownType" || expStr == "crate::Value") && actStr == "bool" then "crate::Value::Bool(" <> code <> ")"
        else if expStr == "f64" && (actStr == "crate::UnknownType" || actStr == "crate::Value") then "(" <> code <> ").unwrap_number()"
        else if (expStr == "crate::UnknownType" || expStr == "crate::Value") && actStr == "f64" then "crate::Value::Number(" <> code <> ")"
        else if expStr == "char" && (actStr == "crate::UnknownType" || actStr == "crate::Value") then "(" <> code <> ").unwrap_char()"
        else if (expStr == "crate::UnknownType" || expStr == "crate::Value") && actStr == "char" then "crate::Value::Char(" <> code <> ")"
        else if expStr == "String" && (actStr == "crate::UnknownType" || actStr == "crate::Value") then "(" <> code <> ").unwrap_string()"
        else if (expStr == "crate::UnknownType" || expStr == "crate::Value") && actStr == "String" then "crate::Value::String((" <> code <> ").clone())"
        else code
