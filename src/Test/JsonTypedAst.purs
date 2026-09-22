module Test.JsonTypedAst where

import Prelude hiding (bind)
import Data.Argonaut.Core (Json, stringify)
import Data.Argonaut.Core as Json
import Data.Argonaut.Parser (jsonParser)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Array as Array
import Data.Either (Either(..))
import Data.Int as Int
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String.CodeUnits as CodeUnits
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Partial.Unsafe (unsafeCrashWith)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module)
import PureScript.Backend.Optimizer.CoreFn as C
import PureScript.Backend.Optimizer.CoreFn.Json (decodeModule)

foreign import drive :: forall a b. (String -> a) -> (a -> b) -> (a -> String) -> (b -> String) -> Effect Unit

parse :: String -> Json
parse input = case jsonParser input of
  Left err -> unsafeCrashWith err
  Right value -> value

decode :: Json -> Module Ann
decode input = case decodeModule input of
  Left err -> unsafeCrashWith (printJsonDecodeError err)
  Right value -> value

main :: Effect Unit
main = drive parse decode stringify fingerprint

-- | An exhaustive, ordered representation of the decoded module. Arrays avoid
-- | depending on JSON object iteration order; only the foreign Map is sorted.
-- | Compute this string and its SHA-256 outside the measured decoding interval.
fingerprint :: C.Module C.Ann -> String
fingerprint = Json.stringify <<< moduleJson

tag :: String -> Array Json -> Json
tag name fields = Json.fromArray (Array.cons (Json.fromString name) fields)

array :: forall a. (a -> Json) -> Array a -> Json
array encode = Json.fromArray <<< map encode

optional :: forall a. (a -> Json) -> Maybe a -> Json
optional encode = case _ of
  Nothing -> tag "Nothing" []
  Just value -> tag "Just" [ encode value ]

pair :: forall a b. (a -> Json) -> (b -> Json) -> Tuple a b -> Json
pair encodeA encodeB (Tuple a b) = Json.fromArray [ encodeA a, encodeB b ]

int :: Int -> Json
int = Json.fromNumber <<< Int.toNumber

strings :: Array String -> Json
strings = array Json.fromString

ident :: C.Ident -> Json
ident (C.Ident value) = Json.fromString value

properName :: C.ProperName -> Json
properName (C.ProperName value) = Json.fromString value

moduleName :: C.ModuleName -> Json
moduleName (C.ModuleName value) = Json.fromString value

qualified :: forall a. (a -> Json) -> C.Qualified a -> Json
qualified encode (C.Qualified name value) =
  tag "Qualified" [ optional moduleName name, encode value ]

sourcePos :: C.SourcePos -> Json
sourcePos value = tag "SourcePos" [ int value.line, int value.column ]

sourceSpan :: C.SourceSpan -> Json
sourceSpan value = tag "SourceSpan"
  [ Json.fromString value.path, sourcePos value.start, sourcePos value.end ]

annotation :: C.Ann -> Json
annotation (C.Ann value) = tag "Ann"
  [ sourceSpan value.span
  , optional meta value.meta
  , optional exprType value.type
  , optional sourceUsage value.sourceUsage
  ]

sourceBindingId :: C.SourceBindingId -> Json
sourceBindingId (C.SourceBindingId value) = tag "SourceBindingId"
  [ moduleName value.moduleName, int value.bindingId ]

bindingUsage :: C.BindingUsage -> Json
bindingUsage value = tag "BindingUsage"
  [ sourceBindingId value.binding
  , optional int value.maxUses
  , optional Json.fromBoolean value.hasEscapingUseContext
  ]

variableUse :: C.VariableUse -> Json
variableUse value = tag "VariableUse"
  [ sourceBindingId value.binding, optional Json.fromBoolean value.lastLocalUse ]

sourceUsage :: C.SourceUsage -> Json
sourceUsage value = tag "SourceUsage"
  [ optional bindingUsage value.bindingUsage, optional variableUse value.variableUse ]

exprType :: C.ExprType -> Json
exprType = case _ of
  C.Int -> tag "Int" []
  C.Number -> tag "Number" []
  C.String -> tag "String" []
  C.Char -> tag "Char" []
  C.Boolean -> tag "Boolean" []
  C.Unit -> tag "Unit" []
  C.Any -> tag "Any" []
  C.TypeLevelString value -> tag "TypeLevelString" [ Json.fromString value ]
  C.Array value -> tag "Array" [ exprType value ]
  C.TypeVar value -> tag "TypeVar" [ Json.fromString value ]
  C.ADT name parts args -> tag "ADT"
    [ Json.fromString name, strings parts, array exprType args ]
  C.TypeApp base args -> tag "TypeApp" [ exprType base, array exprType args ]
  C.Func args result -> tag "Func" [ array exprType args, exprType result ]
  C.Row fields tail -> tag "Row"
    [ array (pair Json.fromString exprType) fields, optional exprType tail ]
  C.Record value -> tag "Record" [ exprType value ]
  C.ForAll vars value -> tag "ForAll" [ strings vars, exprType value ]
  C.ConstrainedType constraints value -> tag "ConstrainedType"
    [ array constraint constraints, exprType value ]

constraint :: Tuple (Array String) (Array C.ExprType) -> Json
constraint value = pair strings (array exprType) value

meta :: C.Meta -> Json
meta = case _ of
  C.IsConstructor kind fields -> tag "IsConstructor"
    [ constructorType kind, array ident fields ]
  C.IsNewtype -> tag "IsNewtype" []
  C.IsTypeClassConstructor -> tag "IsTypeClassConstructor" []
  C.IsForeign -> tag "IsForeign" []
  C.IsWhere -> tag "IsWhere" []
  C.IsSyntheticApp -> tag "IsSyntheticApp" []

constructorType :: C.ConstructorType -> Json
constructorType = case _ of
  C.ProductType -> tag "ProductType" []
  C.SumType -> tag "SumType" []

comment :: C.Comment -> Json
comment = case _ of
  C.LineComment value -> tag "LineComment" [ Json.fromString value ]
  C.BlockComment value -> tag "BlockComment" [ Json.fromString value ]

dataConstructor :: C.DataConstructor -> Json
dataConstructor value = tag "DataConstructor"
  [ Json.fromString value.name, array exprType value.fields ]

dataDecl :: C.DataDecl -> Json
dataDecl value = tag "DataDecl"
  [ Json.fromString value.name
  , strings value.vars
  , array dataConstructor value.constructors
  ]

classDecl :: C.ClassDecl -> Json
classDecl value = tag "ClassDecl"
  [ Json.fromString value.name
  , strings value.vars
  , array constraint value.superclasses
  , array (pair Json.fromString exprType) value.methods
  ]

moduleJson :: C.Module C.Ann -> Json
moduleJson (C.Module value) = tag "Module"
  [ moduleName value.name
  , Json.fromString value.path
  , sourceSpan value.span
  , array importJson value.imports
  , array ident value.exports
  , array reExport value.reExports
  , array dataDecl value.dataDecls
  , array classDecl value.classDecls
  , array bind value.decls
  , array (pair ident (optional exprType)) (Map.toUnfoldable value.foreign)
  , array comment value.comments
  ]

importJson :: C.Import C.Ann -> Json
importJson (C.Import ann name) = tag "Import" [ annotation ann, moduleName name ]

reExport :: C.ReExport -> Json
reExport (C.ReExport name value) = tag "ReExport" [ moduleName name, ident value ]

bind :: C.Bind C.Ann -> Json
bind = case _ of
  C.NonRec value -> tag "NonRec" [ binding value ]
  C.Rec values -> tag "Rec" [ array binding values ]

binding :: C.Binding C.Ann -> Json
binding (C.Binding ann name value) = tag "Binding"
  [ annotation ann, ident name, expr value ]

expr :: C.Expr C.Ann -> Json
expr = case _ of
  C.ExprVar ann value -> tag "ExprVar" [ annotation ann, qualified ident value ]
  C.ExprLit ann value -> tag "ExprLit" [ annotation ann, literal expr value ]
  C.ExprConstructor ann name ctor fields -> tag "ExprConstructor"
    [ annotation ann, properName name, ident ctor, strings fields ]
  C.ExprAccessor ann value key -> tag "ExprAccessor"
    [ annotation ann, expr value, Json.fromString key ]
  C.ExprUpdate ann value fields -> tag "ExprUpdate"
    [ annotation ann, expr value, array (prop expr) fields ]
  C.ExprAbs ann name body -> tag "ExprAbs" [ annotation ann, ident name, expr body ]
  C.ExprApp ann fn arg -> tag "ExprApp" [ annotation ann, expr fn, expr arg ]
  C.ExprCase ann values alternatives -> tag "ExprCase"
    [ annotation ann, array expr values, array caseAlternative alternatives ]
  C.ExprLet ann bindings body -> tag "ExprLet"
    [ annotation ann, array bind bindings, expr body ]
  C.ExprTypeApp ann value arg -> tag "ExprTypeApp"
    [ annotation ann, expr value, exprType arg ]

caseAlternative :: C.CaseAlternative C.Ann -> Json
caseAlternative (C.CaseAlternative binders result) = tag "CaseAlternative"
  [ array binder binders, caseGuard result ]

caseGuard :: C.CaseGuard C.Ann -> Json
caseGuard = case _ of
  C.Unconditional value -> tag "Unconditional" [ expr value ]
  C.Guarded values -> tag "Guarded" [ array guard values ]

guard :: C.Guard C.Ann -> Json
guard (C.Guard condition value) = tag "Guard" [ expr condition, expr value ]

prop :: forall a. (a -> Json) -> C.Prop a -> Json
prop encode (C.Prop key value) = tag "Prop" [ Json.fromString key, encode value ]

literal :: forall a. (a -> Json) -> C.Literal a -> Json
literal encode = case _ of
  C.LitInt value -> tag "LitInt" [ int value ]
  C.LitNumber value -> tag "LitNumber" [ Json.fromNumber value ]
  C.LitString value -> tag "LitString" [ Json.fromString value ]
  C.LitChar value -> tag "LitChar" [ Json.fromString (CodeUnits.singleton value) ]
  C.LitBoolean value -> tag "LitBoolean" [ Json.fromBoolean value ]
  C.LitArray values -> tag "LitArray" [ array encode values ]
  C.LitRecord values -> tag "LitRecord" [ array (prop encode) values ]

binder :: C.Binder C.Ann -> Json
binder = case _ of
  C.BinderNull ann -> tag "BinderNull" [ annotation ann ]
  C.BinderVar ann name -> tag "BinderVar" [ annotation ann, ident name ]
  C.BinderNamed ann name value -> tag "BinderNamed"
    [ annotation ann, ident name, binder value ]
  C.BinderLit ann value -> tag "BinderLit" [ annotation ann, literal binder value ]
  C.BinderConstructor ann name ctor fields -> tag "BinderConstructor"
    [ annotation ann
    , qualified properName name
    , qualified ident ctor
    , array binder fields
    ]
