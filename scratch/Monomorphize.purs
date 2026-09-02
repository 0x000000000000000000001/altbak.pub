-- | Monomorphisation, Spécialisation et DPE (Dictionary Passing Elimination)
-- | L'une des passes d'optimisation les plus sophistiquées.
-- | Elle a pour rôle d'éliminer l'overhead de l'abstraction fonctionnelle :
-- | 1. En spécialisant le code polymorphe (générique) pour des types concrets.
-- | 2. En inlinant les dictionnaires de Type Classes pour transformer des appels dynamiques en appels directs statiques (DPE).
-- | 3. En préparant la décurryfication (via eta-expansion) pour s'assurer que l'arité des fonctions reste visible après l'élimination des dictionnaires.

module PureScript.Backend.Optimizer.Monomorphize
  ( InstantiationMap
  , collectInstantiations
  , collectAllTypes
  , mangleType
  , defaultToAny
  , collectSpine
  , SpineArg(..)
  , getExprAnn
  , inferExprType
  , monomorphize
  , extractFuncType
  , transitiveCollect
  , applyStaticArgs
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Map (Map)
import Data.Map as Map
import Data.String as String
import Data.Maybe (Maybe(..), fromMaybe, maybe, isJust)
import Debug (trace)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Debug (trace)
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binder(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Ident(..), Literal(..), Module(..), ModuleName(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import PureScript.Backend.Optimizer.Substitute (substituteExprType, unify)
import Node.Encoding (Encoding(..))
import Effect.Unsafe (unsafePerformEffect)
import Effect.Console as Console

type InstantiationMap = Map String (Map String { instType :: ExprType, dictArgs :: Array (Expr Ann), normalArgs :: Array (Expr Ann), callers :: Set String, subst :: Map String ExprType })

isStatic :: Expr Ann -> Boolean
isStatic = case _ of
  ExprVar _ (Qualified (Just _) _) -> true
  ExprVar _ (Qualified Nothing _) -> false
  ExprApp _ f arg -> isStatic f && isStatic arg
  ExprAccessor _ e _ -> isStatic e
  -- We deliberately do NOT consider data values (literals, constructors) as static.
  -- Specializing on data values breaks recursive functions where the accumulator 
  -- changes (e.g. `foldl`), because PBO blindly substitutes the initial static value 
  -- everywhere in the body, ruining the loop state.
  ExprLit _ _ -> false
  ExprConstructor _ _ _ _ -> false
  _ -> false

defaultToAny :: ExprType -> ExprType
defaultToAny = case _ of
  TypeVar _ -> Any
  Array t -> Array (defaultToAny t)
  Func args ret -> Func (map defaultToAny args) (defaultToAny ret)
  Record row -> Record (defaultToAny row)
  Row props tail -> Row (map (\(Tuple k v) -> Tuple k (defaultToAny v)) props) (map defaultToAny tail)
  TypeApp c args -> TypeApp (defaultToAny c) (map defaultToAny args)
  ForAll vars body -> ForAll vars (defaultToAny body)
  ConstrainedType constraints body -> ConstrainedType (map (\(Tuple c a) -> Tuple c (map defaultToAny a)) constraints) (defaultToAny body)
  ADT fn names args -> ADT fn names (map defaultToAny args)
  t -> t

mangleType :: ExprType -> String
mangleType Int = "Int"
mangleType Number = "Number"
mangleType String = "String"
mangleType Char = "Char"
mangleType Boolean = "Boolean"
mangleType Unit = "Unit"
mangleType Any = "Any"
mangleType (TypeLevelString s) = "TypeLevelString_" <> s
mangleType (Array t) = "Array_" <> mangleType t
mangleType (Func args ret) = "Func_" <> String.joinWith "_" (map mangleType args) <> "_" <> mangleType ret
mangleType (Record row) = "Record_" <> mangleType row
mangleType (Row props tail) = "Row_" <> String.joinWith "_" (map (\(Tuple k v) -> k <> "_" <> mangleType v) props) <> "_" <> maybe "Empty" mangleType tail
mangleType (TypeApp c args) = "TypeApp_" <> mangleType c <> "_" <> String.joinWith "_" (map mangleType args)
mangleType (ForAll vars body) = "ForAll_" <> String.joinWith "_" vars <> "_" <> mangleType body
mangleType (ConstrainedType constraints body) = "ConstrainedType_" <> String.joinWith "_" (map (\(Tuple c a) -> String.joinWith "_" c <> "_" <> String.joinWith "_" (map mangleType a)) constraints) <> "_" <> mangleType body
mangleType (ADT _ names args) = "ADT_" <> String.joinWith "_" names <> (if Array.length args == 0 then "" else "_" <> String.joinWith "_" (map mangleType args))
mangleType (TypeVar name) = "Var_" <> name


mangleExpr :: forall a. Expr a -> String
mangleExpr = case _ of
  ExprVar _ (Qualified mbMod (Ident name)) ->
    "Var_" <> maybe "" (\(ModuleName mn) -> mn <> "_") mbMod <> name
  ExprLit _ (LitInt i) -> "LitInt_" <> show i
  ExprLit _ (LitNumber n) -> "LitNum_" <> show n
  ExprLit _ (LitString s) -> "LitStr_" <> s
  ExprLit _ (LitChar _) -> "LitChar"
  ExprLit _ (LitBoolean b) -> "LitBool_" <> show b
  ExprLit _ _ -> "LitUnk"
  ExprApp _ f arg -> "App_" <> mangleExpr f <> "_" <> mangleExpr arg
  ExprAccessor _ e prop -> "Acc_" <> prop <> "_" <> mangleExpr e
  ExprConstructor _ _ (Ident c) _ -> "Ctor_" <> c
  _ -> "Unk"

collectInstantiations :: Map String (Binding Ann) -> InstantiationMap -> Module Ann -> InstantiationMap
collectInstantiations globalAstMap acc (Module m) =
  let
    modNameStr = unwrap m.name
  in
    foldl (collectBind globalAstMap modNameStr) acc m.decls

collectBind :: Map String (Binding Ann) -> String -> InstantiationMap -> Bind Ann -> InstantiationMap
collectBind globalAstMap modName acc (NonRec binding) = collectBinding globalAstMap modName acc binding
collectBind globalAstMap modName acc (Rec bindings) = foldl (collectBinding globalAstMap modName) acc bindings

collectBinding :: Map String (Binding Ann) -> String -> InstantiationMap -> Binding Ann -> InstantiationMap
collectBinding globalAstMap modName acc (Binding _ _ expr) = collectExpr globalAstMap modName acc expr

getExprAnn :: Expr Ann -> Ann
getExprAnn = case _ of
  ExprVar ann _ -> ann
  ExprLit ann _ -> ann
  ExprApp ann _ _ -> ann
  ExprAbs ann _ _ -> ann
  ExprLet ann _ _ -> ann
  ExprTypeApp ann _ _ -> ann
  ExprCase ann _ _ -> ann
  ExprConstructor ann _ _ _ -> ann
  ExprAccessor ann _ _ -> ann
  ExprUpdate ann _ _ -> ann

inferExprType :: Expr Ann -> Maybe ExprType
inferExprType expr = let (Ann ann) = getExprAnn expr in ann.type

extractFuncType :: ExprType -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractFuncType = case _ of
  Func fArgs fRet -> Just { fArgs, fRet }
  ForAll _ body -> extractFuncType body
  ConstrainedType _ body -> extractFuncType body
  _ -> Nothing

data SpineArg = SpineApp (Expr Ann) | SpineTypeApp ExprType

collectSpine :: Expr Ann -> { f_var :: Expr Ann, spine :: Array SpineArg }
collectSpine = go []
  where
  go acc (ExprApp _ f x) = go (Array.cons (SpineApp x) acc) f
  go acc (ExprTypeApp _ f t) = go (Array.cons (SpineTypeApp t) acc) f
  go acc f = { f_var: f, spine: acc }

getSpineArgs :: Array SpineArg -> Array (Expr Ann)
getSpineArgs = Array.mapMaybe (case _ of
  SpineApp e -> Just e
  _ -> Nothing)

getSpineTypeArgs :: Array SpineArg -> Array ExprType
getSpineTypeArgs = Array.mapMaybe (case _ of
  SpineTypeApp t -> Just t
  _ -> Nothing)

buildSubst :: ExprType -> Array ExprType -> Map String ExprType
buildSubst (ForAll vars body) typeArgs =
  let
    subst1 = Map.fromFoldable (Array.zip vars typeArgs)
    subst2 = buildSubst body (Array.drop (Array.length vars) typeArgs)
  in Map.union subst1 subst2
buildSubst (ConstrainedType _ body) typeArgs = buildSubst body typeArgs
buildSubst _ _ = Map.empty

partitionArgs :: ExprType -> Array (Expr Ann) -> { dictArgs :: Array (Expr Ann), normalArgs :: Array (Expr Ann) }
partitionArgs (ConstrainedType constraints _) args =
  let
    numDicts = Array.length constraints
    dictArgs = Array.take numDicts args
    normalArgs = Array.drop numDicts args
  in
    { dictArgs, normalArgs }
partitionArgs (ForAll _ body) args = partitionArgs body args
partitionArgs _ args = { dictArgs: [], normalArgs: args }



stripStaticConstraints :: Array (Expr Ann) -> ExprType -> ExprType
stripStaticConstraints dictArgs = case _ of
  ConstrainedType constraints body ->
    let
      newConstraints = Array.mapMaybe (\(Tuple d c) -> if isStatic d then Nothing else Just c) (Array.zip dictArgs constraints)
      remainingConstraints = Array.drop (Array.length dictArgs) constraints
      finalConstraints = newConstraints <> remainingConstraints
    in
      if Array.length finalConstraints == 0 then stripStaticConstraints dictArgs body else ConstrainedType finalConstraints (stripStaticConstraints dictArgs body)
  ForAll vars body -> ForAll vars (stripStaticConstraints dictArgs body)
  t -> t

collectExpr :: Map String (Binding Ann) -> String -> InstantiationMap -> Expr Ann -> InstantiationMap
collectExpr globalAstMap modName acc expr = case expr of
  ExprVar (Ann ann) (Qualified mbMod (Ident name)) ->
    let
      qualName = case mbMod of
        Just mod -> unwrap mod <> "." <> name
        Nothing -> modName <> "." <> name
      trueGenericType = case Map.lookup qualName globalAstMap of
        Just (Binding (Ann bAnn) _ _) -> bAnn.type
        Nothing -> ann.type
    in
      case trueGenericType of
        Just t ->
          Map.insertWith (\new old -> Map.unionWith (\a b -> { instType: a.instType, dictArgs: a.dictArgs, normalArgs: a.normalArgs, callers: Set.union a.callers b.callers, subst: a.subst }) new old) qualName (Map.singleton (mangleType (defaultToAny t)) { instType: defaultToAny t, dictArgs: [], normalArgs: [], callers: Set.singleton modName, subst: Map.empty }) acc
        Nothing -> acc
  ExprApp _ _ _ ->
    let
      { f_var, spine } = collectSpine expr
      args = getSpineArgs spine
      typeArgs = getSpineTypeArgs spine
      acc1 = collectExpr globalAstMap modName acc f_var
      acc2 = foldl (collectExpr globalAstMap modName) acc1 args
    in
      case f_var of
        ExprVar (Ann varAnn) (Qualified mbMod (Ident name)) ->
          let
             qualName = case mbMod of
               Just mod -> unwrap mod <> "." <> name
               Nothing -> modName <> "." <> name
             trueGenericType = case Map.lookup qualName globalAstMap of
               Just (Binding (Ann bAnn) _ _) -> case bAnn.type of
                 Just t -> t
                 Nothing -> fromMaybe Any varAnn.type
               Nothing -> fromMaybe Any varAnn.type
             genericType = trueGenericType
             substFromTypeArgs = buildSubst genericType typeArgs
             
             printTy :: ExprType -> String
             printTy Any = "Any"
             printTy Int = "Int"
             printTy (TypeVar v) = "TypeVar " <> v
             printTy (ForAll v t) = "ForAll [" <> String.joinWith ", " v <> "] (" <> printTy t <> ")"
             printTy (Func args' ret') = "Func [" <> String.joinWith ", " (map printTy args') <> "] (" <> printTy ret' <> ")"
             printTy (ADT name _ args') = "ADT " <> name <> " [" <> String.joinWith ", " (map printTy args') <> "]"
             printTy _ = "Other"

             unifySpine :: ExprType -> Array (Expr Ann) -> Map String ExprType -> Map String ExprType
             unifySpine _ [] s = s
             unifySpine (ForAll _ t) args' s = unifySpine t args' s
             unifySpine (ConstrainedType constraints t) args' s =
               let numConstraints = Array.length constraints
               in unifySpine t (Array.drop numConstraints args') s
             unifySpine (Func paramTypes ret) args' s =
               let
                 numParams = Array.length paramTypes
                 appliedArgs = Array.take numParams args'
                 remainingArgs = Array.drop numParams args'
                 s1 = foldl (\acc (Tuple paramType arg) ->
                        let actualType = case getExprAnn arg of Ann a -> fromMaybe Any a.type
                            _ = if String.contains (Pattern "filter") qualName then trace ("  unify paramType: " <> printTy paramType <> " with actualType: " <> printTy actualType) \_ -> unit else unit
                        in unify paramType actualType acc
                      ) s (Array.zip paramTypes appliedArgs)
               in
                 unifySpine ret remainingArgs s1
             unifySpine _ _ s = s
             
             subst = unifySpine genericType args substFromTypeArgs
             stripForAlls = case _ of
               ForAll _ b -> stripForAlls b
               x -> x
             instType = stripTypeVariables (substituteExprType subst (stripForAlls genericType))
             _ = if String.contains (Pattern "filter") qualName then trace ("\n\n[TRACE 1.6.3.3] collectExpr ExprApp for " <> qualName <> "\n  is in globalAstMap? " <> show (Map.member qualName globalAstMap) <> "\n  hasTypeVars instType: " <> show (hasTypeVariables instType) <> "\n  subst size: " <> show (Map.size subst) <> "\n  genericType: " <> printTy genericType <> "\n  args len: " <> show (Array.length args)) \_ -> unit else unit
             { dictArgs, normalArgs } = partitionArgs genericType args
             staticNormalArgs = Array.filter isStatic normalArgs
             specKey = mangleType (defaultToAny instType) <> "_" <> String.joinWith "_" (map mangleExpr staticNormalArgs)
          in
             if not (hasTypeVariables genericType) then acc2
             else if hasTypeVariables instType then acc2
             else Map.insertWith (\new old -> Map.unionWith (\a b -> { instType: a.instType, dictArgs: a.dictArgs, normalArgs: a.normalArgs, callers: Set.union a.callers b.callers, subst: a.subst }) new old) qualName (Map.singleton specKey { instType: defaultToAny instType, dictArgs, normalArgs, callers: Set.singleton modName, subst }) acc2
        _ -> acc2

  ExprLit _ lit -> foldl (collectExpr globalAstMap modName) acc lit
  ExprConstructor _ _ _ _ -> acc
  ExprAccessor _ e _ -> collectExpr globalAstMap modName acc e
  ExprUpdate _ e props -> foldl (collectProp globalAstMap modName) (collectExpr globalAstMap modName acc e) props
  ExprAbs _ _ e -> collectExpr globalAstMap modName acc e
  ExprTypeApp _ e _ -> collectExpr globalAstMap modName acc e
  ExprCase _ exprs alts -> foldl (collectAlt globalAstMap modName) (foldl (collectExpr globalAstMap modName) acc exprs) alts
  ExprLet _ binds e -> foldl (collectBind globalAstMap modName) (collectExpr globalAstMap modName acc e) binds

collectProp :: Map String (Binding Ann) -> String -> InstantiationMap -> Prop (Expr Ann) -> InstantiationMap
collectProp globalAstMap modName acc (Prop _ e) = collectExpr globalAstMap modName acc e

collectAlt :: Map String (Binding Ann) -> String -> InstantiationMap -> CaseAlternative Ann -> InstantiationMap
collectAlt globalAstMap modName acc (CaseAlternative _ cg) = case cg of
  Unconditional e -> collectExpr globalAstMap modName acc e
  Guarded guards -> foldl (collectGuard globalAstMap modName) acc guards

collectGuard :: Map String (Binding Ann) -> String -> InstantiationMap -> Guard Ann -> InstantiationMap
collectGuard globalAstMap modName acc (Guard e1 e2) = collectExpr globalAstMap modName (collectExpr globalAstMap modName acc e1) e2

collectAllTypes :: Module Ann -> Set ExprType
collectAllTypes (Module m) = foldl (\a b -> collectTypesFromBind b a) Set.empty m.decls

type LocalInstMap = Map Ident { genericType :: ExprType, insts :: Array (Map String ExprType) }

collectLocalExpr :: Set Ident -> LocalInstMap -> Expr Ann -> LocalInstMap
collectLocalExpr targets acc expr = case expr of
  ExprApp ann f arg ->
    let
      spineRec = collectSpine (ExprApp ann f arg)
      f_var = spineRec.f_var
      spine = spineRec.spine
      typeArgs = getSpineTypeArgs spine
      args = getSpineArgs spine
      
      acc1 = case f_var of
        ExprVar _ (Qualified Nothing id) | Set.member id targets ->
          let
            genericType = case getExprAnn f_var of Ann a -> fromMaybe Any a.type
            _ = if id == Ident "go" || id == Ident "$go" then trace ("\n\n[TRACE 1.6.3.1] collectLocalExpr found go call!\n  genericType: " <> show (hasTypeVariables genericType)) \_ -> unit else unit
            substFromTypeArgs = buildSubst genericType typeArgs
            
            unifySpine :: ExprType -> Array (Expr Ann) -> Map String ExprType -> Map String ExprType
            unifySpine _ [] s = s
            unifySpine (ForAll _ t) args' s = unifySpine t args' s
            unifySpine (ConstrainedType constraints t) args' s =
              let numConstraints = Array.length constraints
              in unifySpine t (Array.drop numConstraints args') s
            unifySpine (Func paramTypes ret) args' s =
              let
                numParams = Array.length paramTypes
                appliedArgs = Array.take numParams args'
                remainingArgs = Array.drop numParams args'
                s1 = foldl (\acc (Tuple paramType arg) ->
                       let actualType = case getExprAnn arg of Ann a -> fromMaybe Any a.type
                           
                           printTy :: ExprType -> String
                           printTy Int = "Int"
                           printTy Number = "Number"
                           printTy String = "String"
                           printTy Char = "Char"
                           printTy Boolean = "Boolean"
                           printTy Unit = "Unit"
                           printTy Any = "Any"
                           printTy (TypeVar v) = "TypeVar " <> v
                           printTy (TypeApp c args') = "TypeApp (" <> printTy c <> ") [" <> String.joinWith ", " (map printTy args') <> "]"
                           printTy (ForAll v t) = "ForAll [" <> String.joinWith ", " v <> "] (" <> printTy t <> ")"
                           printTy (Func args' ret') = "Func [" <> String.joinWith ", " (map printTy args') <> "] (" <> printTy ret' <> ")"
                           printTy (ADT name _ args') = "ADT " <> name <> " [" <> String.joinWith ", " (map printTy args') <> "]"
                           printTy _ = "Other"

                           _ = if id == Ident "go" || id == Ident "$go" then trace ("\n\n[TRACE 1.6.3.2] unifySpine go arg" <> "\n  paramType:  " <> printTy paramType <> "\n  actualType: " <> printTy actualType) \_ -> unit else unit
                       in unify paramType actualType acc
                     ) s (Array.zip paramTypes appliedArgs)
              in
                unifySpine ret remainingArgs s1
            unifySpine _ _ s = s
            
            finalSubst = unifySpine genericType args substFromTypeArgs
          in
            if not (Map.isEmpty finalSubst) then
              let
                stripForAlls = case _ of
                  ForAll _ b -> stripForAlls b
                  x -> x
                instType = stripTypeVariables (substituteExprType finalSubst (stripForAlls genericType))
              in
                if (hasTypeVariables genericType) && (hasTypeVariables instType && instType == stripForAlls genericType) then
                  acc
                else
                  Map.insertWith (\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [finalSubst] } acc
            else
              -- Even if finalSubst is empty, we must record the genericType so it can be restored!
              Map.insertWith (\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [Map.empty] } acc
        _ -> collectLocalExpr targets acc f_var
      
      acc2 = foldl (collectLocalExpr targets) acc1 args
    in
      acc2
  ExprLit _ lit -> foldl (collectLocalExpr targets) acc lit
  ExprConstructor _ _ _ _ -> acc
  ExprAccessor _ e _ -> collectLocalExpr targets acc e
  ExprUpdate _ e props -> foldl (\a (Prop _ v) -> collectLocalExpr targets a v) (collectLocalExpr targets acc e) props
  ExprAbs _ _ e -> collectLocalExpr targets acc e
  ExprTypeApp _ e _ -> collectLocalExpr targets acc e
  ExprCase _ exprs alts ->
    foldl (\a (CaseAlternative _ cg) -> case cg of
      Unconditional e' -> collectLocalExpr targets a e'
      Guarded guards -> foldl (\a2 (Guard e1 e2) -> collectLocalExpr targets (collectLocalExpr targets a2 e1) e2) a guards
    ) (foldl (collectLocalExpr targets) acc exprs) alts
  ExprLet _ binds e -> foldl (\a b -> collectLocalBind targets a b) (collectLocalExpr targets acc e) binds
  ExprVar _ _ -> acc

collectLocalBind :: Set Ident -> LocalInstMap -> Bind Ann -> LocalInstMap
collectLocalBind targets acc = case _ of
  NonRec (Binding _ _ e) -> collectLocalExpr targets acc e
  Rec binds -> foldl (\a (Binding _ _ e) -> collectLocalExpr targets a e) acc binds

collectTypesFromExpr :: Expr Ann -> Set ExprType -> Set ExprType
collectTypesFromExpr expr acc = case expr of
  ExprVar (Ann ann) _ -> maybe acc (\t -> Set.insert t acc) ann.type
  ExprLit (Ann ann) lit -> foldl (\a e -> collectTypesFromExpr e a) (maybe acc (\t -> Set.insert t acc) ann.type) lit
  ExprApp (Ann ann) f arg -> collectTypesFromExpr arg (collectTypesFromExpr f (maybe acc (\t -> Set.insert t acc) ann.type))
  ExprAbs (Ann ann) _ e -> collectTypesFromExpr e (maybe acc (\t -> Set.insert t acc) ann.type)
  ExprLet (Ann ann) binds e -> foldl (\a b -> collectTypesFromBind b a) (collectTypesFromExpr e (maybe acc (\t -> Set.insert t acc) ann.type)) binds
  ExprTypeApp (Ann ann) e _ -> collectTypesFromExpr e (maybe acc (\t -> Set.insert t acc) ann.type)
  ExprCase (Ann ann) exprs alts -> foldl (\a alt -> collectTypesFromAlt alt a) (foldl (flip collectTypesFromExpr) (maybe acc (\t -> Set.insert t acc) ann.type) exprs) alts
  ExprConstructor (Ann ann) _ _ _ -> maybe acc (\t -> Set.insert t acc) ann.type
  ExprAccessor (Ann ann) e _ -> collectTypesFromExpr e (maybe acc (\t -> Set.insert t acc) ann.type)
  ExprUpdate (Ann ann) e props -> foldl (\a (Prop _ v) -> collectTypesFromExpr v a) (collectTypesFromExpr e (maybe acc (\t -> Set.insert t acc) ann.type)) props

collectTypesFromBind :: Bind Ann -> Set ExprType -> Set ExprType
collectTypesFromBind (NonRec (Binding _ _ e)) acc = collectTypesFromExpr e acc
collectTypesFromBind (Rec binds) acc = foldl (\a (Binding _ _ e) -> collectTypesFromExpr e a) acc binds

collectTypesFromAlt :: CaseAlternative Ann -> Set ExprType -> Set ExprType
collectTypesFromAlt (CaseAlternative _ (Unconditional e)) acc = collectTypesFromExpr e acc
collectTypesFromAlt (CaseAlternative _ (Guarded guards)) acc = foldl (\a (Guard e1 e2) -> collectTypesFromExpr e2 (collectTypesFromExpr e1 a)) acc guards

mapAnn :: (ExprType -> ExprType) -> Ann -> Ann
mapAnn f (Ann ann) = Ann (ann { type = map f ann.type })

rewriteExpr :: Map String (Binding Ann) -> Map Ident (Expr Ann) -> Map String (Expr Ann) -> (ExprType -> ExprType) -> Expr Ann -> Expr Ann
rewriteExpr globalAstMap = goLocals
  where
  goLocals locals globalSubst f = go
    where
    printTy :: ExprType -> String
    printTy Any = "Any"
    printTy Int = "Int"
    printTy (TypeVar v) = "TypeVar " <> v
    printTy (ForAll v t) = "ForAll [" <> String.joinWith ", " v <> "] (" <> printTy t <> ")"
    printTy (Func args' ret') = "Func [" <> String.joinWith ", " (map printTy args') <> "] (" <> printTy ret' <> ")"
    printTy (ADT name _ args') = "ADT " <> name <> " [" <> String.joinWith ", " (map printTy args') <> "]"
    printTy _ = "Other"

    go expr = case expr of
      ExprVar ann q@(Qualified mbMod (Ident name)) ->
        let
          qualName = case mbMod of
            Just mod -> unwrap mod <> "." <> name
            Nothing -> name
        in case Map.lookup qualName globalSubst of
          Just newExpr -> newExpr
          Nothing -> 
            let 
              oldType = case ann of Ann a -> fromMaybe Any a.type
              newAnn = mapAnn f ann
              newType = case newAnn of Ann a -> fromMaybe Any a.type
              _ = if name == "go" || name == "$go" then trace ("\n\n[TRACE 1.6.3.3] rewriteExpr ExprVar for " <> name <> "\n  oldType: " <> printTy oldType <> "\n  newType: " <> printTy newType) \_ -> unit else unit
            in ExprVar newAnn q
      ExprLit ann lit -> ExprLit (mapAnn f ann) (map go lit)
      ExprApp ann e1 e2 -> ExprApp (mapAnn f ann) (go e1) (go e2)
      ExprAbs ann id e -> ExprAbs (mapAnn f ann) id (goLocals (Map.delete id locals) globalSubst f e)
      ExprLet ann binds e ->
        let
          binds' = map goBind binds
          newLocals = foldl
            ( \acc b -> case b of
                NonRec (Binding _ ident val) -> Map.insert ident val acc
                _ -> acc
            )
            locals
            binds'
          e' = goLocals newLocals globalSubst f e

          foldFn b acc = case b of
            NonRec (Binding _ ident val) ->
              if Set.member ident acc.used then
                { used: Set.union (Set.delete ident acc.used) (collectFreeVars val), binds: Array.cons b acc.binds }
              else
                acc
            Rec bs ->
              let
                bound = foldl (\a (Binding _ ident _) -> Set.insert ident a) Set.empty bs
                isUsed = Array.any (\(Binding _ ident _) -> Set.member ident acc.used) bs
              in
                if isUsed then
                  let
                    usedByRec = foldl (\a (Binding _ _ val) -> Set.union a (collectFreeVars val)) acc.used bs
                    newUsed = Set.difference usedByRec bound
                  in
                    { used: newUsed, binds: Array.cons b acc.binds }
                else
                  acc

          filtered = Array.foldr foldFn { used: collectFreeVars e', binds: [] } binds'
        in
          if Array.length filtered.binds == 0 then e' else ExprLet (mapAnn f ann) filtered.binds e'
      ExprTypeApp ann e t ->
        ExprTypeApp (mapAnn f ann) (go e) (f t)
      ExprCase ann exprs alts ->
        let
          exprs' = map go exprs
        in
          case exprs', alts of
            [ e1 ], [ CaseAlternative [ BinderConstructor _ _ _ [ BinderVar _ ident ] ] (Unconditional e2) ] ->
              case resolveDict e1 of
                Just lit -> goLocals (Map.insert ident (ExprLit (getExprAnn e1) lit) locals) globalSubst f e2
                Nothing -> ExprCase (mapAnn f ann) exprs' (map goAlt alts)
            _, _ -> ExprCase (mapAnn f ann) exprs' (map goAlt alts)
      ExprConstructor ann t c ids -> ExprConstructor (mapAnn f ann) t c ids
      ExprAccessor ann e prop ->
        let
          e' = go e
        in
          case resolveDict e' of
            Just (LitRecord props) ->
              case Array.find (\(Prop p _) -> p == prop) props of
                Just (Prop _ val) -> go val
                Nothing -> ExprAccessor (mapAnn f ann) e' prop
            Just (LitArray _) ->
              ExprAccessor (mapAnn f ann) e' prop
            _ -> ExprAccessor (mapAnn f ann) e' prop
      ExprUpdate ann e props -> ExprUpdate (mapAnn f ann) (go e) (map goProp props)

    resolveDict :: Expr Ann -> Maybe (Literal (Expr Ann))
    resolveDict e = case e of
      ExprLit _ lit@(LitRecord _) -> Just lit
      ExprApp _ (ExprConstructor _ _ _ _) e' -> resolveDict e'
      ExprVar _ (Qualified mbMod ident) ->
        case mbMod of
          Nothing -> case Map.lookup ident locals of
            Just val -> resolveDict val
            Nothing -> Nothing
          Just (ModuleName mn) ->
            let
              fullName = mn <> "." <> (\(Ident n) -> n) ident
            in
              case Map.lookup fullName globalAstMap of
                Just (Binding _ _ val) -> resolveDict val
                Nothing -> Nothing
      ExprApp _ (ExprVar _ (Qualified _ (Ident ctorName))) arg | String.contains (Pattern "$Dict") ctorName -> resolveDict arg
      _ -> Nothing

    goBind (NonRec b) = NonRec (goBinding b)
    goBind (Rec bs) = Rec (map goBinding bs)

    goBinding (Binding ann id e) =
      let
        rewrittenE = go e
        eAnn = case getExprAnn rewrittenE of Ann a -> a.type
        inferredType = case eAnn of
          Just t -> Just t
          Nothing -> case ann of Ann a -> a.type
        ann' = mapAnn f (Ann (case ann of Ann a -> a { type = inferredType }))
        name = case id of Ident n -> n
        isGo = name == "go" || name == "$go"
        showHasVars = case _ of
          Just t -> show (hasTypeVariables t)
          Nothing -> "Nothing"
        _ = if isGo then trace ("\n\n[TRACE 1.6.3.1] rewriteExpr goBinding for " <> name <> "\n  orig vars: " <> showHasVars (case ann of Ann a -> a.type) <> "\n  new vars:  " <> showHasVars (case ann' of Ann a -> a.type)) \_ -> unit else unit
      in
        Binding ann' id rewrittenE

    goAlt (CaseAlternative binders cg) = CaseAlternative (map goBinder binders) (goCaseGuard cg)

    goBinder binder = map (mapAnn f) binder

    goCaseGuard (Unconditional e) = Unconditional (go e)
    goCaseGuard (Guarded guards) = Guarded (map goGuard guards)

    goGuard (Guard e1 e2) = Guard (go e1) (go e2)

    goProp (Prop p e) = Prop p (go e)


substituteVars :: Map Ident (Expr Ann) -> Expr Ann -> Expr Ann
substituteVars subst = go
  where
  go = case _ of
    ExprVar ann (Qualified Nothing ident) ->
      case Map.lookup ident subst of
        Just e -> e
        Nothing -> ExprVar ann (Qualified Nothing ident)
    ExprVar ann q -> ExprVar ann q
    ExprLit ann lit -> ExprLit ann (map go lit)
    ExprApp ann f arg -> ExprApp ann (go f) (go arg)
    ExprAbs ann ident body ->
      ExprAbs ann ident (substituteVars (Map.delete ident subst) body)
    ExprLet ann binds body ->
      let
        bound = foldl
          ( \acc -> case _ of
              NonRec (Binding _ ident _) -> Set.insert ident acc
              Rec bs -> foldl (\acc2 (Binding _ ident _) -> Set.insert ident acc2) acc bs
          )
          Set.empty
          binds
        subst' = Map.filterKeys (\k -> not (Set.member k bound)) subst
      in
        ExprLet ann
          (map (\b -> case b of
              NonRec (Binding annB ident expr) -> NonRec (Binding annB ident (substituteVars subst' expr))
              Rec bs -> Rec (map (\(Binding annB ident expr) -> Binding annB ident (substituteVars subst' expr)) bs)
          ) binds)
          (substituteVars subst' body)
    ExprConstructor ann t c idents -> ExprConstructor ann t c idents
    ExprTypeApp ann e t -> ExprTypeApp ann (go e) t
    ExprCase ann exprs alts -> ExprCase ann (map go exprs) (map (\(CaseAlternative binders guard) -> CaseAlternative binders (goGuard subst guard)) alts)
    ExprAccessor ann e prop -> ExprAccessor ann (go e) prop
    ExprUpdate ann e props -> ExprUpdate ann (go e) (map (\(Prop p v) -> Prop p (go v)) props)

  goGuard subst (Unconditional e) = Unconditional (substituteVars subst e)
  goGuard subst (Guarded guards) = Guarded (map (\(Guard e1 e2) -> Guard (substituteVars subst e1) (substituteVars subst e2)) guards)

applyStaticArgs :: Array (Expr Ann) -> Array (Expr Ann) -> Expr Ann -> Expr Ann
applyStaticArgs dictArgs normalArgs body =
  let
    -- collect dictionary arguments
    resDicts = goCollect "dict" dictArgs body
    -- collect normal arguments
    resNorms = goCollect "norm" normalArgs resDicts.expr
    
    subst = Map.union resDicts.subst resNorms.subst
    
    -- substitute variables
    substBody = substituteVars subst resNorms.expr
  in
    -- re-wrap with unused normal arguments to keep the same arity,
    -- but DO NOT re-wrap dicts because they are removed from the caller's spine!
    wrapUnused resNorms.unusedIds substBody
  where
  goCollect prefix args e = case Array.uncons args of
    Nothing -> { subst: Map.empty, unusedIds: [], expr: e }
    Just { head: a, tail: as' } ->
      if isStatic a then
        case e of
          ExprAbs ann id b ->
            let
              rest = goCollect prefix as' b
            in
              { subst: Map.insert id a rest.subst, unusedIds: Array.cons (Tuple ann (Ident (unwrap id <> "_unused"))) rest.unusedIds, expr: rest.expr }
          _ -> 
            let
              freshId = Ident ("__eta_" <> prefix <> "_" <> show (Array.length as'))
              ann = getExprAnn e
              etaBody = ExprApp ann e (ExprVar ann (Qualified Nothing freshId))
              rest = goCollect prefix as' etaBody
            in
              { subst: Map.insert freshId a rest.subst, unusedIds: Array.cons (Tuple ann (Ident (unwrap freshId <> "_unused"))) rest.unusedIds, expr: rest.expr }
      else
        case e of
          ExprAbs _ id b -> 
            let rest = goCollect prefix as' b
            in { subst: rest.subst, unusedIds: rest.unusedIds, expr: ExprAbs (getExprAnn e) id rest.expr }
          _ ->
            let
              freshId = Ident ("__eta_" <> prefix <> "_" <> show (Array.length as'))
              ann = getExprAnn e
              etaBody = ExprApp ann e (ExprVar ann (Qualified Nothing freshId))
              rest = goCollect prefix as' etaBody
            in
              { subst: rest.subst, unusedIds: rest.unusedIds, expr: ExprAbs ann freshId rest.expr }

  wrapUnused ids e = Array.foldr (\(Tuple ann id) acc -> ExprAbs ann id acc) e ids

getBindIdents :: Array (Bind Ann) -> Array Ident
getBindIdents = Array.concatMap case _ of
  NonRec (Binding _ id _) -> [ id ]
  Rec binds -> map (\(Binding _ id _) -> id) binds

monomorphize :: Map String (Binding Ann) -> InstantiationMap -> Module Ann -> Module Ann
monomorphize globalAstMap instMap (Module m) =
  let
    modNameStr = unwrap m.name

    getInjectedBindsFor qualName = case Map.lookup qualName instMap of
      Just typeMap ->
        let
          processBinding = case Map.lookup qualName globalAstMap of
            Just (Binding ann (Ident name) expr) ->
              Array.mapMaybe
                ( \(Tuple specKey info) ->
                    if hasTypeVariables info.instType then Nothing
                    else
                      let
                        definerMod = case String.split (Pattern ".") qualName of
                          parts -> String.joinWith "." (fromMaybe [] (Array.init parts))
                      in
                        if modNameStr == definerMod then
                          let
                            stripForAlls = case _ of
                              ForAll _ b -> stripForAlls b
                              x -> x
                            substFn t = stripStaticConstraints info.dictArgs (substituteExprType info.subst (stripForAlls t))
                            astSubstFn t = substituteExprType info.subst (stripForAlls t)

                            exprWithDicts = applyStaticArgs info.dictArgs info.normalArgs expr
                            resolvedExpr = resolveGlobals definerMod Set.empty exprWithDicts

                            specializedVar = ExprVar ann (Qualified (Just (ModuleName definerMod)) (Ident (name <> "__" <> hashString specKey)))
                            globalSubst = Map.fromFoldable [ Tuple qualName specializedVar, Tuple name specializedVar ]
                            finalTy = stripTypeVariables (substFn info.instType)

                            rewrittenExpr = rewriteExpr globalAstMap Map.empty globalSubst astSubstFn resolvedExpr
                            localInstMap = collectAllLocalInsts [] Map.empty rewrittenExpr
                            specializedExpr = monomorphizeExpr modNameStr instMap localInstMap Map.empty rewrittenExpr

                            etaExpandedExpr = case specializedExpr of
                              ExprAbs _ _ _ -> specializedExpr
                              _ | Array.length info.dictArgs == 0 && Array.length info.normalArgs == 0 -> specializedExpr
                              _ ->
                                let
                                  monomorphizedAnn = mapAnn (\_ -> finalTy) ann
                                in
                                  case extractFuncType finalTy of
                                    Just { fArgs } ->
                                      let
                                        idents = Array.mapWithIndex (\i _ -> Ident ("__eta" <> show i)) fArgs
                                        vars = map (\id -> ExprVar monomorphizedAnn (Qualified Nothing id)) idents
                                        app = foldl (\acc v -> ExprApp monomorphizedAnn acc v) specializedExpr vars
                                      in
                                        Array.foldr (\id acc -> ExprAbs monomorphizedAnn id acc) app idents
                                    Nothing -> specializedExpr

                            newName = Ident (name <> "__" <> hashString specKey)
                            newBinding = Rec [ Binding (mapAnn (\t -> stripTypeVariables (substFn t)) ann) newName etaExpandedExpr ]
                          in
                            Just newBinding
                        else Nothing
                )
                (Map.toUnfoldable typeMap :: Array _)
            Nothing -> []
        in
          processBinding
      Nothing -> []

    processDecl bind =
      let
        originalBinds = monomorphizeBind modNameStr instMap Map.empty bind
        injectedBinds = case bind of
          NonRec (Binding _ (Ident name) _) -> getInjectedBindsFor (modNameStr <> "." <> name)
          Rec bindings -> Array.concatMap (\(Binding _ (Ident name) _) -> getInjectedBindsFor (modNameStr <> "." <> name)) bindings
      in originalBinds <> injectedBinds

    finalDecls = Array.concatMap processDecl m.decls
    newIdents = getBindIdents finalDecls
  in
    Module (m { decls = finalDecls, exports = newIdents })

monomorphizeBind :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Bind Ann -> Array (Bind Ann)
monomorphizeBind modName instMap localDicts (NonRec binding) =
  [ NonRec (monomorphizeBinding modName instMap localDicts binding) ]
monomorphizeBind modName instMap localDicts (Rec bindings) =
  [ Rec (map (monomorphizeBinding modName instMap localDicts) bindings) ]

monomorphizeBinding :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Binding Ann -> Binding Ann
monomorphizeBinding modName instMap localDicts (Binding ann (Ident name) expr) =
  let
    localInstMap = collectAllLocalInsts [] Map.empty expr
  in
    Binding ann (Ident name) (monomorphizeExpr modName instMap localInstMap localDicts expr)

collectFreeVars :: Expr Ann -> Set Ident
collectFreeVars = case _ of
  ExprVar _ (Qualified Nothing ident) -> Set.singleton ident
  ExprVar _ _ -> Set.empty
  ExprLit _ lit -> foldl (\acc e -> Set.union acc (collectFreeVars e)) Set.empty lit
  ExprApp _ e1 e2 -> Set.union (collectFreeVars e1) (collectFreeVars e2)
  ExprAbs _ ident e -> Set.delete ident (collectFreeVars e)
  ExprLet _ binds e ->
    let
      bound = foldl
        ( \acc b -> case b of
            NonRec (Binding _ ident _) -> Set.insert ident acc
            Rec bs -> foldl (\a (Binding _ ident _) -> Set.insert ident a) acc bs
        )
        Set.empty
        binds
      used = foldl
        ( \acc b -> case b of
            NonRec (Binding _ _ val) -> Set.union acc (collectFreeVars val)
            Rec bs -> foldl (\a (Binding _ _ val) -> Set.union a (collectFreeVars val)) acc bs
        )
        (collectFreeVars e)
        binds
    in
      Set.difference used bound
  ExprTypeApp _ e _ ->
    collectFreeVars e
  ExprCase _ exprs alts ->
    foldl (\acc alt -> Set.union acc (collectFreeVarsAlt alt)) (foldl (\acc e -> Set.union acc (collectFreeVars e)) Set.empty exprs) alts
  ExprConstructor _ _ _ _ -> Set.empty
  ExprAccessor _ e _ -> collectFreeVars e
  ExprUpdate _ e props -> foldl (\acc (Prop _ val) -> Set.union acc (collectFreeVars val)) (collectFreeVars e) props

collectFreeVarsAlt :: CaseAlternative Ann -> Set Ident
collectFreeVarsAlt (CaseAlternative binders cg) =
  let
    bound = foldl (\acc binder -> Set.union acc (binderIdents binder)) Set.empty binders
    used = case cg of
      Unconditional e -> collectFreeVars e
      Guarded guards -> foldl (\acc (Guard g e) -> Set.union acc (Set.union (collectFreeVars g) (collectFreeVars e))) Set.empty guards
  in
    Set.difference used bound

monomorphizeExpr :: String -> InstantiationMap -> LocalInstMap -> Map Ident (Expr Ann) -> Expr Ann -> Expr Ann
monomorphizeExpr modName instMap localInstMap localDicts expr = case expr of
  ExprVar ann ident@(Qualified mbMod (Ident name)) ->
    case mbMod of
      Nothing -> case Map.lookup (Ident name) localDicts of
        Just d -> d
        Nothing -> ExprVar ann ident
      _ -> ExprVar ann ident
  expr | isAppOrTypeApp expr ->
    let
      Ann ann = getExprAnn expr
      { f_var, spine } = collectSpine expr

      typeArgs = getSpineTypeArgs spine
      
      f_var' = monomorphizeExpr modName instMap localInstMap localDicts f_var
      spine' = map (case _ of
                       SpineApp e -> SpineApp (monomorphizeExpr modName instMap localInstMap localDicts e)
                       SpineTypeApp t -> SpineTypeApp t) spine
      args' = getSpineArgs spine'
    in
      case f_var' of
        ExprVar (Ann varAnn) (Qualified mbMod (Ident name)) ->
          let
             genericType = fromMaybe Any varAnn.type
             subst = buildSubst genericType typeArgs
             stripForAlls = case _ of
               ForAll _ b -> stripForAlls b
               x -> x
             instType = stripTypeVariables (substituteExprType subst (stripForAlls genericType))
             { dictArgs, normalArgs } = partitionArgs genericType args'
             qualName = case mbMod of
               Just mod -> unwrap mod <> "." <> name
               Nothing -> modName <> "." <> name
             filteredArgs = Array.filter (\d -> not (isStatic d)) dictArgs <> normalArgs
             
          in
             let 
               isReverse = name == "reverse" || name == "reverse__Int64"
               _ = if isReverse then trace ("\n\n[TRACE 1.6.1] ExprTypeApp for " <> name <> "\n  genericType=" <> show (hasTypeVariables genericType) <> "\n  instType=" <> show (hasTypeVariables instType) <> "\n  varAnn.type=" <> if isJust varAnn.type then "Just" else "Nothing") \_ -> unit else unit
             in
             if hasTypeVariables instType then
               rebuildSpine (Ann ann) f_var' spine'
             else
               let staticNormalArgs = Array.filter isStatic normalArgs
                   specKey = mangleType (defaultToAny instType) <> "_" <> String.joinWith "_" (map mangleExpr staticNormalArgs)
                   specializedName = Ident (name <> "__" <> hashString specKey)
               in case Map.lookup qualName instMap of
                    Just typeMap ->
                      case Map.lookup specKey typeMap of
                        Just info -> 
                          let
                             stripForAlls2 = case _ of
                               ForAll _ b -> stripForAlls2 b
                               x -> x
                             substFn t = stripStaticConstraints dictArgs (substituteExprType info.subst (stripForAlls2 t))
                             newAnn = varAnn { type = map (\t -> stripTypeVariables (substFn t)) varAnn.type }
                             definerMod = case String.split (Pattern ".") qualName of
                               parts -> String.joinWith "." (fromMaybe [] (Array.init parts))
                             resolvedMod = Just (ModuleName definerMod)
                             specializedVar = ExprVar (Ann newAnn) (Qualified resolvedMod specializedName)
                          in
                             rebuildSpine (Ann ann) specializedVar (map SpineApp filteredArgs)
                        Nothing -> rebuildSpine (Ann ann) f_var' spine'
                    Nothing -> rebuildSpine (Ann ann) f_var' spine'
        _ -> rebuildSpine (Ann ann) f_var' spine'

  ExprLit ann lit -> ExprLit ann (map (monomorphizeExpr modName instMap localInstMap localDicts) lit)
  ExprAbs ann id e -> ExprAbs ann id (monomorphizeExpr modName instMap localInstMap localDicts e)
  ExprLet ann binds e ->
    let
      newLocalDicts = Array.foldl
        ( \acc b -> case b of
            NonRec (Binding _ id bindExpr) | isStatic bindExpr -> Map.insert id bindExpr acc
            _ -> acc
        )
        localDicts
        binds

      boundIds = foldl (\acc -> case _ of
          NonRec (Binding _ id _) -> Set.insert id acc
          Rec bs -> foldl (\a (Binding _ id _) -> Set.insert id a) acc bs
        ) Set.empty binds
        
      localInstMap = collectLocalExpr boundIds Map.empty (ExprLet ann binds e)
      
      fastPath = ExprLet ann (map (monomorphizeBindLocal modName instMap localInstMap newLocalDicts) binds) (monomorphizeExpr modName instMap localInstMap newLocalDicts e)
    in
      if Map.isEmpty localInstMap then fastPath
      else
        let
          injectType ty expr =
            let
              stripForAlls = case _ of
                ForAll _ b -> stripForAlls b
                x -> x
            in
              case stripForAlls ty, expr of
                Func paramTypes ret, ExprAbs (Ann a) argId bodyExpr ->
                  if Array.length paramTypes > 1 then
                    let nextTy = Func (Array.drop 1 paramTypes) ret
                    in ExprAbs (Ann (a { type = Just ty })) argId (injectType nextTy bodyExpr)
                  else if Array.length paramTypes == 1 then
                    ExprAbs (Ann (a { type = Just ty })) argId (injectType ret bodyExpr)
                  else
                    ExprAbs (Ann (a { type = Just ty })) argId bodyExpr
                _, ExprAbs (Ann a) argId bodyExpr ->
                  ExprAbs (Ann (a { type = Just ty })) argId bodyExpr
                _, _ -> expr

          processBinds = foldl (\acc bind -> 
            case bind of
              NonRec b@(Binding bAnn id bExpr) ->
                case Map.lookup id localInstMap of
                  Just { genericType, insts: substList } ->
                     let
                       fixedAnn = case bAnn of Ann a -> Ann (a { type = Just genericType })
                       fixedExpr = injectType genericType bExpr
                       specs = Array.mapMaybe (\substType ->
                           let
                             instType = stripTypeVariables (substituteExprType substType genericType)
                             specKey = mangleType (defaultToAny instType)
                             mangledId = Ident (unwrap id <> "__" <> hashString specKey)
                             specExpr = rewriteExpr Map.empty Map.empty Map.empty (\t -> substituteExprType substType (stripTypeVariables t)) fixedExpr
                             newBind = Binding (mapAnn (\t -> substituteExprType substType (stripTypeVariables t)) fixedAnn) mangledId specExpr
                           in
                             Just { specKey, mangledId, newBind }
                       ) substList
                       
                       newBinds = [NonRec b] <> map (\s -> NonRec s.newBind) specs
                       polyMapEntry = Map.fromFoldable (map (\s -> Tuple s.specKey s.mangledId) specs)
                     in
                       { binds: acc.binds <> newBinds, polyMap: Map.insert id polyMapEntry acc.polyMap }
                  Nothing -> { binds: Array.snoc acc.binds bind, polyMap: acc.polyMap }
              
              Rec bs ->
                let
                  specs = foldl (\accRec (Binding bAnn id bExpr) ->
                    case Map.lookup id localInstMap of
                      Just { genericType, insts: substList } ->
                         let
                           fixedAnn = case bAnn of Ann a -> Ann (a { type = Just genericType })
                           fixedExpr = injectType genericType bExpr
                           s = Array.mapMaybe (\substType ->
                               let
                                 instType = stripTypeVariables (substituteExprType substType genericType)
                                 specKey = mangleType (defaultToAny instType)
                                 mangledId = Ident (unwrap id <> "__" <> hashString specKey)
                                 specExpr = rewriteExpr Map.empty Map.empty Map.empty (\t -> substituteExprType substType (stripTypeVariables t)) fixedExpr
                                 newBind = Binding (mapAnn (\t -> substituteExprType substType (stripTypeVariables t)) fixedAnn) mangledId specExpr
                               in
                                 Just { specKey, mangledId, newBind }
                           ) substList
                         in
                           { newBinds: accRec.newBinds <> map _.newBind s
                           , polyMap: Map.insert id (Map.fromFoldable (map (\x -> Tuple x.specKey x.mangledId) s)) accRec.polyMap 
                           }
                      Nothing -> accRec
                  ) { newBinds: [], polyMap: acc.polyMap } bs
                in
                  if Array.length bs == 1 then
                    { binds: acc.binds <> map (\b -> Rec [b]) specs.newBinds <> [Rec bs], polyMap: specs.polyMap }
                  else if Array.length specs.newBinds > 0 then
                    { binds: Array.snoc acc.binds (Rec (specs.newBinds <> bs)), polyMap: specs.polyMap }
                  else
                    { binds: Array.snoc acc.binds (Rec bs), polyMap: acc.polyMap }
          ) { binds: [], polyMap: Map.empty } binds

          polys = processBinds.polyMap
          
          go expr = case expr of
            ExprApp annApp f arg ->
              let
                spineRec = collectSpine expr
                f_var = spineRec.f_var
                spine = spineRec.spine
              in
                case f_var of
                  ExprVar varAnn (Qualified Nothing id) | Just insts <- Map.lookup id polys ->
                    let
                      typeArgs = getSpineTypeArgs spine
                      args = getSpineArgs spine
                      
                      genericType = case getExprAnn f_var of Ann a -> fromMaybe Any a.type
                      substFromTypeArgs = buildSubst genericType typeArgs
                      
                      unifySpine :: ExprType -> Array (Expr Ann) -> Map String ExprType -> Map String ExprType
                      unifySpine _ [] s = s
                      unifySpine (ForAll _ t) args' s = unifySpine t args' s
                      unifySpine (ConstrainedType constraints t) args' s =
                        let numConstraints = Array.length constraints
                        in unifySpine t (Array.drop numConstraints args') s
                      unifySpine (Func paramTypes ret) args' s =
                        let
                          numParams = Array.length paramTypes
                          appliedArgs = Array.take numParams args'
                          remainingArgs = Array.drop numParams args'
                          s1 = foldl (\acc (Tuple paramType arg) ->
                                 let actualType = case getExprAnn arg of Ann a -> fromMaybe Any a.type
                                 in unify paramType actualType acc
                               ) s (Array.zip paramTypes appliedArgs)
                        in
                          unifySpine ret remainingArgs s1
                      unifySpine _ _ s = s
                      
                      substType = unifySpine genericType args substFromTypeArgs
                    in
                        let
                          instType = stripTypeVariables (substituteExprType substType genericType)
                          specKey = mangleType (defaultToAny instType)
                        in
                          case Map.lookup specKey insts of
                            Just mangledId ->
                              let
                                newVar = ExprVar (mapAnn (\_ -> defaultToAny instType) varAnn) (Qualified Nothing mangledId)
                              in
                                foldl (\acc a -> ExprApp annApp acc (go a)) newVar args
                            Nothing ->
                              let
                                app1 = foldl (\acc t -> ExprTypeApp annApp acc t) (go f_var) typeArgs
                              in
                                foldl (\acc a -> ExprApp annApp acc (go a)) app1 args
                  _ -> 
                     let
                       typeArgs = getSpineTypeArgs spine
                       args = getSpineArgs spine
                       app1 = foldl (\acc t -> ExprTypeApp annApp acc t) (go f_var) typeArgs
                     in
                       foldl (\acc a -> ExprApp annApp acc (go a)) app1 args
            ExprLit annLit lit -> ExprLit annLit (map go lit)
            ExprAbs annAbs id e' -> ExprAbs annAbs id (go e')
            ExprLet annLet binds' e' -> ExprLet annLet (map goBind binds') (go e')
            ExprTypeApp annTy e' t -> ExprTypeApp annTy (go e') t
            ExprCase annCase exprs alts -> ExprCase annCase (map go exprs) (map goAlt alts)
            ExprConstructor annCtor t c ids -> ExprConstructor annCtor t c ids
            ExprAccessor annAcc e' prop -> ExprAccessor annAcc (go e') prop
            ExprUpdate annUp e' props -> ExprUpdate annUp (go e') (map goProp props)
            ExprVar _ _ -> expr

          goBind (NonRec (Binding bAnn id e')) = NonRec (Binding bAnn id (go e'))
          goBind (Rec binds') = Rec (map (\(Binding bAnn id e') -> Binding bAnn id (go e')) binds')
          goAlt (CaseAlternative binders cg) = CaseAlternative binders (goCaseGuard cg)
          goCaseGuard (Unconditional e') = Unconditional (go e')
          goCaseGuard (Guarded guards) = Guarded (map (\(Guard e1 e2) -> Guard (go e1) (go e2)) guards)
          goProp (Prop p e') = Prop p (go e')

          rewrittenBinds = map goBind processBinds.binds
          rewrittenE = go e
        in
          ExprLet ann (map (monomorphizeBindLocal modName instMap localInstMap newLocalDicts) rewrittenBinds) (monomorphizeExpr modName instMap localInstMap newLocalDicts rewrittenE)
  ExprCase ann exprs alts -> ExprCase ann (map (monomorphizeExpr modName instMap localInstMap localDicts) exprs) (map (monomorphizeAlt modName instMap localInstMap localDicts) alts)
  ExprConstructor ann t c ids -> ExprConstructor ann t c ids
  ExprAccessor ann e prop -> ExprAccessor ann (monomorphizeExpr modName instMap localInstMap localDicts e) prop
  ExprUpdate ann e props -> ExprUpdate ann (monomorphizeExpr modName instMap localInstMap localDicts e) (map (monomorphizeProp modName instMap localInstMap localDicts) props)
  _ -> expr
  where
  isAppOrTypeApp (ExprApp _ _ _) = true
  isAppOrTypeApp (ExprTypeApp _ _ _) = true
  isAppOrTypeApp _ = false

monomorphizeBindLocal :: String -> InstantiationMap -> LocalInstMap -> Map Ident (Expr Ann) -> Bind Ann -> Bind Ann
monomorphizeBindLocal modName instMap localInstMap localDicts (NonRec b) = NonRec (monomorphizeBindingLocal modName instMap localDicts b)
monomorphizeBindLocal modName instMap localInstMap localDicts (Rec bs) = Rec (map (monomorphizeBindingLocal modName instMap localDicts) bs)

monomorphizeBindingLocal :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Binding Ann -> Binding Ann
monomorphizeBindingLocal modName instMap localDicts (Binding ann id e) = 
  let
    name = case id of Ident n -> n
    isGo = name == "go" || name == "$go"
    showHasVars = case _ of
      Just t -> show (hasTypeVariables t)
      Nothing -> "Nothing"
    _ = if isGo then trace ("\n\n[TRACE 1.6.3.1] monomorphizeBindingLocal for " <> name <> "\n  vars: " <> showHasVars (case ann of Ann a -> a).type) \_ -> unit else unit
  in
  Binding ann id (monomorphizeExpr modName instMap localInstMap localDicts e)

monomorphizeAlt :: String -> InstantiationMap -> Map Ident (Expr Ann) -> CaseAlternative Ann -> CaseAlternative Ann
monomorphizeAlt modName instMap localInstMap localDicts (CaseAlternative binders cg) = CaseAlternative binders (monomorphizeCaseGuard modName instMap localInstMap localDicts cg)

monomorphizeCaseGuard :: String -> InstantiationMap -> Map Ident (Expr Ann) -> CaseGuard Ann -> CaseGuard Ann
monomorphizeCaseGuard modName instMap localInstMap localDicts (Unconditional e) = Unconditional (monomorphizeExpr modName instMap localInstMap localDicts e)
monomorphizeCaseGuard modName instMap localInstMap localDicts (Guarded guards) = Guarded (map (monomorphizeGuard modName instMap localInstMap localDicts) guards)

monomorphizeGuard :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Guard Ann -> Guard Ann
monomorphizeGuard modName instMap localInstMap localDicts (Guard e1 e2) = Guard (monomorphizeExpr modName instMap localInstMap localDicts e1) (monomorphizeExpr modName instMap localInstMap localDicts e2)

monomorphizeProp :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Prop (Expr Ann) -> Prop (Expr Ann)
monomorphizeProp modName instMap localInstMap localDicts (Prop p e) = Prop p (monomorphizeExpr modName instMap localInstMap localDicts e)

rebuildSpine :: Ann -> Expr Ann -> Array SpineArg -> Expr Ann
rebuildSpine finalAnn f spine = Array.foldl applySpine f spine
  where
  applySpine acc (SpineApp arg) = ExprApp finalAnn acc arg
  applySpine acc (SpineTypeApp t) = ExprTypeApp finalAnn acc t

hasTypeVariables :: ExprType -> Boolean
hasTypeVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"
hasTypeVariables (Func args ret) = Array.any hasTypeVariables args || hasTypeVariables ret
hasTypeVariables (Array t) = hasTypeVariables t
hasTypeVariables (Record row) = hasTypeVariables row
hasTypeVariables (Row props tail) =
  let
    tailHas = case tail of
      Nothing -> false
      Just t -> hasTypeVariables t
  in
    Array.any (\(Tuple _ v) -> hasTypeVariables v) props || tailHas
hasTypeVariables (TypeApp c args) = hasTypeVariables c || Array.any hasTypeVariables args
hasTypeVariables (ForAll _ body) = hasTypeVariables body
hasTypeVariables (ConstrainedType constraints body) = Array.any (\(Tuple _ a) -> Array.any hasTypeVariables a) constraints || hasTypeVariables body
hasTypeVariables Int = false
hasTypeVariables String = false
hasTypeVariables Char = false
hasTypeVariables Number = false
hasTypeVariables Boolean = false
hasTypeVariables Unit = false
hasTypeVariables (TypeLevelString _) = false
hasTypeVariables (ADT _ _ args) = Array.any hasTypeVariables args
hasTypeVariables Any = true

stripTypeVariables :: ExprType -> ExprType
stripTypeVariables (ForAll _ t) = stripTypeVariables t
stripTypeVariables (ConstrainedType _ t) = stripTypeVariables t
stripTypeVariables t = t

resolveGlobals :: String -> Set Ident -> Expr Ann -> Expr Ann
resolveGlobals definerMod = go
  where
  go bound = case _ of
    ExprVar ann (Qualified Nothing ident) ->
      if Set.member ident bound then
        ExprVar ann (Qualified Nothing ident)
      else
        ExprVar ann (Qualified (Just (ModuleName definerMod)) ident)
    ExprVar ann q -> ExprVar ann q
    ExprAbs ann ident body ->
      ExprAbs ann ident (go (Set.insert ident bound) body)
    ExprApp ann f arg ->
      ExprApp ann (go bound f) (go bound arg)
    ExprLet ann binds body ->
      let
        bound' = foldl
          ( \acc -> case _ of
              NonRec (Binding _ ident _) -> Set.insert ident acc
              Rec bs -> foldl (\acc2 (Binding _ ident _) -> Set.insert ident acc2) acc bs
          )
          bound
          binds
        binds' = map
          ( \b -> case b of
              NonRec (Binding annB ident expr) -> NonRec (Binding annB ident (go bound' expr))
              Rec bs -> Rec (map (\(Binding annB ident expr) -> Binding annB ident (go bound' expr)) bs)
          )
          binds
      in
        ExprLet ann binds' (go bound' body)
    ExprConstructor ann t c idents ->
      ExprConstructor ann t c idents
    ExprTypeApp ann e t ->
      ExprTypeApp ann (go bound e) t
    ExprCase ann exprs alts ->
      ExprCase ann (map (go bound) exprs)
        ( map
            ( \(CaseAlternative binders guard) ->
                let
                  binders' = map (resolveGlobalsBinder definerMod) binders
                  bound' = foldl (\acc binder -> Set.union (binderIdents binder) acc) bound binders'

                  goGuard cgGuard = case cgGuard of
                    Unconditional e -> Unconditional (go bound' e)
                    Guarded guards -> Guarded (map (\(Guard g e) -> Guard (go bound' g) (go bound' e)) guards)

                in
                  CaseAlternative binders' (goGuard guard)
            )
            alts
        )
    ExprAccessor ann expr p -> ExprAccessor ann (go bound expr) p
    ExprUpdate ann expr updates -> ExprUpdate ann (go bound expr) (map (\(Prop p e) -> Prop p (go bound e)) updates)
    ExprLit ann lit -> ExprLit ann (map (go bound) lit)

resolveGlobalsBinder :: String -> Binder Ann -> Binder Ann
resolveGlobalsBinder definerMod = case _ of
  BinderConstructor ann t c binders ->
    let
      t' = case t of
        Qualified Nothing pn -> Qualified (Just (ModuleName definerMod)) pn
        _ -> t
      c' = case c of
        Qualified Nothing id -> Qualified (Just (ModuleName definerMod)) id
        _ -> c
    in
      BinderConstructor ann t' c' (map (resolveGlobalsBinder definerMod) binders)
  BinderNamed ann ident b -> BinderNamed ann ident (resolveGlobalsBinder definerMod b)
  BinderLit ann lit -> BinderLit ann (map (resolveGlobalsBinder definerMod) lit)
  b -> b

binderIdents :: forall a. Binder a -> Set Ident
binderIdents = case _ of
  BinderVar _ ident -> Set.singleton ident
  BinderConstructor _ _ _ binders -> foldl (\acc b -> Set.union (binderIdents b) acc) Set.empty binders
  BinderNamed _ ident binder -> Set.insert ident (binderIdents binder)
  BinderLit _ lit -> foldl (\acc b -> Set.union (binderIdents b) acc) Set.empty lit
  _ -> Set.empty

transitiveCollect :: Map String (Binding Ann) -> InstantiationMap -> InstantiationMap
transitiveCollect globalAstMap initialMap = loop initialMap
  where
  loop currentMap =
    let
      newMap = Array.foldl
        ( \acc1 (Tuple qualName typeMap) ->
            Array.foldl
              ( \acc2 (Tuple _ info) ->
                  let
                    definerMod = case String.split (Pattern ".") qualName of
                      parts -> String.joinWith "." (fromMaybe [] (Array.init parts))

                    genericExprOpt = Map.lookup qualName globalAstMap
                  in
                    case genericExprOpt of
                      Just (Binding _ _ expr) ->
                        if hasTypeVariables info.instType then acc2
                        else
                          let
                            stripForAlls = case _ of
                              ForAll _ b -> stripForAlls b
                              x -> x
                            astSubstFn t = 
                              let
                                printMap :: Map String ExprType -> String
                                printMap m = "{" <> String.joinWith ", " (map (\(Tuple k v) -> k <> ": " <> show (hasTypeVariables v)) (Map.toUnfoldable m :: Array (Tuple String ExprType))) <> "}"
                                _ = if String.contains (Pattern "filter") qualName || String.contains (Pattern "reverse") qualName then trace ("\n\n[TRACE 1.6.3.2] buildSpecializedBindings for " <> qualName <> "\ncollectAllLocalInsts :: Array SpineArg -> LocalInstMap -> Expr Ann -> LocalInstMap
collectAllLocalInsts spine acc expr = case expr of
  ExprApp ann f arg ->
    let
      acc1 = collectAllLocalInsts (Array.cons (SpineApp arg) spine) acc f
      acc2 = collectAllLocalInsts [] acc1 arg
    in
      acc2
  ExprTypeApp ann f t ->
    collectAllLocalInsts (Array.cons (SpineTypeApp t) spine) acc f
  ExprVar _ (Qualified Nothing id) ->
    let
      typeArgs = getSpineTypeArgs spine
      args = getSpineArgs spine
      
      genericType = case getExprAnn expr of Ann a -> fromMaybe Any a.type
      substFromTypeArgs = buildSubst genericType typeArgs
      
      unifySpine :: ExprType -> Array (Expr Ann) -> Map String ExprType -> Map String ExprType
      unifySpine _ [] s = s
      unifySpine (ForAll _ t) args' s = unifySpine t args' s
      unifySpine (ConstrainedType constraints t) args' s =
        let numConstraints = Array.length constraints
        in unifySpine t (Array.drop numConstraints args') s
      unifySpine (Func paramTypes ret) args' s =
        let
          numParams = Array.length paramTypes
          appliedArgs = Array.take numParams args'
          remainingArgs = Array.drop numParams args'
          s1 = foldl (\sAcc (Tuple paramType arg) ->
                 let actualType = case getExprAnn arg of Ann a -> fromMaybe Any a.type
                 in unify paramType actualType sAcc
               ) s (Array.zip paramTypes appliedArgs)
        in
          unifySpine ret remainingArgs s1
      unifySpine _ _ s = s
      
      finalSubst = unifySpine genericType args substFromTypeArgs
    in
      if not (Map.isEmpty finalSubst) then
        let
          stripForAlls = case _ t -> case t of
            ForAll _ b -> stripForAlls b
            x -> x
          instType = stripTypeVariables (substituteExprType finalSubst (stripForAlls genericType))
        in
          if (hasTypeVariables genericType) && (hasTypeVariables instType && instType == stripForAlls genericType) then
            acc
          else
            Map.insertWith (\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [finalSubst] } acc
      else
        Map.insertWith (\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [Map.empty] } acc

  ExprVar _ _ -> acc
  ExprLit _ lit -> foldl (\a e -> collectAllLocalInsts [] a e) acc lit
  ExprConstructor _ _ _ _ -> acc
  ExprAccessor _ e _ -> collectAllLocalInsts [] acc e
  ExprUpdate _ e props -> foldl (\a (Prop _ v) -> collectAllLocalInsts [] a v) (collectAllLocalInsts [] acc e) props
  ExprAbs _ _ e -> collectAllLocalInsts [] acc e
  ExprCase _ exprs alts ->
    foldl (\a (CaseAlternative _ cg) -> case cg of
      Unconditional e' -> collectAllLocalInsts [] a e'
      Guarded guards -> foldl (\a2 (Guard e1 e2) -> collectAllLocalInsts [] (collectAllLocalInsts [] a2 e1) e2) a guards
    ) (foldl (\a e -> collectAllLocalInsts [] a e) acc exprs) alts
  ExprLet _ binds e -> 
    foldl (\a b -> collectAllLocalBind a b) (collectAllLocalInsts spine acc e) binds

collectAllLocalBind :: LocalInstMap -> Bind Ann -> LocalInstMap
collectAllLocalBind acc = case _ of
  NonRec (Binding _ _ e) -> collectAllLocalInsts [] acc e
  Rec binds -> foldl (\a (Binding _ _ e) -> collectAllLocalInsts [] a e) acc binds

-- DUMMY TO REPLACE OLD CODE\n  info.subst: " <> printMap info.subst) \_ -> unit else unit
                              in substituteExprType info.subst (stripForAlls t)
                            exprWithDicts = applyStaticArgs info.dictArgs info.normalArgs expr
                            resolvedExpr = resolveGlobals definerMod Set.empty exprWithDicts
                          in
                            foldl
                              ( \acc3 caller ->
                                  let
                                    substitutedExpr = rewriteExpr globalAstMap Map.empty Map.empty astSubstFn resolvedExpr
                                    localInstMap = collectAllLocalInsts [] Map.empty substitutedExpr
                                    specializedExpr = monomorphizeExpr caller currentMap localInstMap Map.empty substitutedExpr
                                  in
                                    collectExpr globalAstMap caller acc3 specializedExpr
                              )
                              acc2
                              (Set.toUnfoldable info.callers :: Array String)
                      Nothing -> acc2
              )
              acc1
              (Map.toUnfoldable typeMap :: Array _)
        )
        currentMap
        (Map.toUnfoldable currentMap :: Array _)
    in
      let
        countCallers m = Array.foldl
          ( \acc (Tuple _ typeMap) ->
              acc + Array.foldl (\a (Tuple _ info) -> a + Set.size info.callers) 0 (Map.toUnfoldable typeMap :: Array _)
          )
          0
          (Map.toUnfoldable m :: Array _)
        callers1 = countCallers currentMap
        callers2 = countCallers newMap
      in
        if callers1 == callers2 then currentMap else loop newMap
