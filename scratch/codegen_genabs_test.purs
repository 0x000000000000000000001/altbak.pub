module Scratch.CodeGenGenAbsTest where

import Prelude
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Tuple (Tuple(..))
import Data.Array as Array

data ExprType = Int | Boolean | Number | String | Char | Func (Array ExprType) ExprType | Any | ForAll String ExprType | ConstrainedType (Array String) ExprType

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
  Func args ret -> 
    let retStr = if Array.length args > 1 then codegenExprType true (Func (Array.drop 1 args) ret) else codegenExprType true ret
    in "std::rc::Rc<dyn Fn(" <> codegenExprType false (fromMaybe Any (Array.head args)) <> ") -> " <> retStr <> ">"
  _ -> "crate::Value"

extractAllArgTypes :: ExprType -> Array ExprType
extractAllArgTypes (ForAll _ t) = extractAllArgTypes t
extractAllArgTypes (ConstrainedType cs t) = Array.replicate (Array.length cs) Any <> extractAllArgTypes t
extractAllArgTypes (Func args t) = args <> extractAllArgTypes t
extractAllArgTypes _ = []

extractFinalRetType :: ExprType -> ExprType
extractFinalRetType (ForAll _ t) = extractFinalRetType t
extractFinalRetType (ConstrainedType _ t) = extractFinalRetType t
extractFinalRetType (Func _ t) = extractFinalRetType t
extractFinalRetType t = t
