{-# LANGUAGE BangPatterns #-}
{-# LANGUAGE DataKinds #-}
{-# LANGUAGE FlexibleInstances #-}
{-# LANGUAGE GADTs #-}
{-# LANGUAGE KindSignatures #-}
{-# LANGUAGE ScopedTypeVariables #-}
{-# LANGUAGE TypeApplications #-}
{-# LANGUAGE TypeOperators #-}
module Main where

import GHC.Clock (getMonotonicTimeNSec)
import Data.IORef (IORef, newIORef, readIORef, writeIORef)
import Data.Word (Word64)
import qualified Data.Array.Unboxed as Array
import qualified Data.List as List
import Data.Kind (Type)
import Data.Proxy (Proxy(..))
import GHC.TypeLits (Symbol)
import Control.Monad (forM_, replicateM, replicateM_)
import Control.Exception (evaluate)
import Text.Printf
import System.Environment (getArgs)

-- 1. AstTree
data Expr = Val !Int | Add !Expr !Expr | Mul !Expr !Expr | Sub !Expr !Expr

evalAst :: Expr -> Int
evalAst (Val v) = v
evalAst (Add l r) = evalAst l + evalAst r
evalAst (Mul l r) = evalAst l * evalAst r
evalAst (Sub l r) = evalAst l - evalAst r

buildTreeAst :: Int -> Expr
buildTreeAst 0 = Val 1
buildTreeAst n = Add (Mul (Val n) (buildTreeAst (n - 1))) (Sub (buildTreeAst (n - 1)) (Val 1))

runAstTree :: Int -> Int
runAstTree limit = evalAst (buildTreeAst limit)

-- 2. Fib
fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n - 1) + fib (n - 2)

runFib :: Int -> Int
runFib limit = fib limit

-- 3. ListOps
-- Strict constructor fields preserve PureScript's eager, immutable lists.
-- Keep range, filtering (in reverse order), and folding as separate operations.
data Lst a = Nil | Cons !a !(Lst a)

lrange :: Int -> Int -> Lst Int
lrange start end = go end Nil
  where
    go !curr !acc | curr < start = acc
                   | otherwise = go (curr - 1) (Cons curr acc)

filterEvens :: Lst Int -> Lst Int
filterEvens lst = go lst Nil
  where
    go Nil !acc = acc
    go (Cons x xs) !acc
      | x `mod` 2 == 0 = go xs (Cons x acc)
      | otherwise = go xs acc

lfoldl :: (b -> a -> b) -> b -> Lst a -> b
lfoldl _ !acc Nil = acc
lfoldl f !acc (Cons x xs) = lfoldl f (f acc x) xs

runListOps :: Int -> Int
runListOps limit = lfoldl (+) 0 (filterEvens (lrange 1 limit))

-- 4. TCO
runTCO :: Int -> Int
runTCO limit = go limit 0
  where
    go 0 !acc = acc
    go !n !acc = go (n - 1) (acc + (n `mod` 3))

-- 5. Records
data DictE = DictE { e :: !Int, f :: !Int }
data DictC = DictC { c :: !Int, d :: !DictE }
data DictA = DictA { a :: !Int, b :: !DictC }

runRecords :: Int -> Int
runRecords limit = go limit (DictA 0 (DictC 0 (DictE 0 0)))
  where
    go 0 !r = f (d (b r))
    go !n !r =
      let oldD = d (b r)
          newD = oldD { e = e oldD + 3, f = f oldD + (n `mod` 5) }
          oldC = b r
          newC = oldC { c = c oldC + 2, d = newD }
      in go (n - 1) r { a = a r + 1, b = newC }

-- 6. Ackermann
ack :: Int -> Int -> Int
ack 0 !n = n + 1
ack !m 0 = ack (m - 1) 1
ack !m !n = ack (m - 1) (ack m (n - 1))

runAckermann :: Int -> Int
runAckermann limit = ack limit 4

-- 7. Church
-- These are the actual higher-order numerals from Test.Church, not integer
-- arithmetic or a hand-written count of closure applications.
type Church a = (a -> a) -> a -> a

zeroC :: Church a
zeroC _ !x = x

succC :: Church a -> Church a
succC !n !f !x = f (n f x)

addC :: Church a -> Church a -> Church a
addC !m !n !f !x = m f (n f x)

mulC :: Church a -> Church a -> Church a
mulC !m !n !f !x = m (n f) x

fromInt :: Int -> Church Int
fromInt 0 = zeroC
fromInt n = let !previous = fromInt (n - 1) in succC previous

toInt :: Church Int -> Int
toInt n = n (\x -> x + 1) 0

c10 :: Int -> Church Int
c10 = fromInt

c100 :: Int -> Church Int
c100 n =
  let !left = c10 n
      !right = c10 n
  in mulC left right

c10k :: Int -> Church Int
c10k n =
  let !left = c100 n
      !right = c100 n
  in mulC left right

c100k :: Int -> Church Int
c100k n =
  let !left = c10k n
      !right = c10 n
  in mulC left right

runChurch :: Int -> Int
runChurch limit = toInt (c100k limit)

-- 8. Primes
lreverse :: Lst a -> Lst a
lreverse lst = go lst Nil
  where
    go Nil !acc = acc
    go (Cons x xs) !acc = go xs (Cons x acc)

lfilter :: (a -> Bool) -> Lst a -> Lst a
lfilter p lst = go lst Nil
  where
    go Nil !acc = lreverse acc
    go (Cons x xs) !acc
      | p x = go xs (Cons x acc)
      | otherwise = go xs acc

lsum :: Lst Int -> Int
lsum xs = go xs 0
  where
    go Nil !acc = acc
    go (Cons x rest) !acc = go rest (acc + x)

sieve :: Lst Int -> Lst Int
sieve Nil = Nil
sieve (Cons p rest) = Cons p (sieve (lfilter (\x -> x `mod` p /= 0) rest))

runPrimes :: Int -> Int
runPrimes limit = lsum (sieve (lrange 2 limit))

-- 9. RBTree
data Color = R | B
data Tree = E | T !Color !Tree !Int !Tree

balance :: Color -> Tree -> Int -> Tree -> Tree
balance B (T R (T R a x b) y cc) z d = T R (T B a x b) y (T B cc z d)
balance B (T R a x (T R b y cc)) z d = T R (T B a x b) y (T B cc z d)
balance B a x (T R (T R b y cc) z d) = T R (T B a x b) y (T B cc z d)
balance B a x (T R b y (T R cc z d)) = T R (T B a x b) y (T B cc z d)
balance c l v r = T c l v r

ins :: Int -> Tree -> Tree
ins x E = T R E x E
ins x (T c l y r)
  | x < y     = balance c (ins x l) y r
  | x > y     = balance c l y (ins x r)
  | otherwise = T c l y r

insert :: Int -> Tree -> Tree
insert x t = case ins x t of
  T _ l y r -> T B l y r
  E -> E

depth :: Tree -> Int
depth E = 0
depth (T _ l _ r) =
  let ld = depth l
      rd = depth r
  in if ld > rd then 1 + ld else 1 + rd

runRBTree :: Int -> Int
runRBTree limit = go limit E
  where
    go 0 !acc = depth acc
    go !i !acc = go (i - 1) (insert i acc)

-- 10. Polymorphism
-- GHC may specialize this class dictionary, just as an optimizing PureScript
-- backend may; the source retains the generic loop and instance dispatch.
class Monoidish a where
  mempty_ :: a
  mappend_ :: a -> a -> a

instance Monoidish Int where
  mempty_ = 1
  mappend_ x y = x + y

polyLoop :: Monoidish a => Int -> a -> a
polyLoop nInit accInit = go nInit accInit
  where
    go 0 !acc = acc
    go !n !acc = go (n - 1) (mappend_ acc mempty_)

runPolymorphism :: Int -> Int
runPolymorphism limit = polyLoop limit 0

-- 11. StateMonad
-- Match the custom State closure implementation, including a fresh state of
-- zero for each 60-deep run. This is not a mutable counter or a flat loop.
data StateResult s a = StateResult { stateVal :: !a, stateValue :: !s }
newtype State s a = State (s -> StateResult s a)

runState :: State s a -> s -> StateResult s a
runState (State f) !s = f s

bindState :: State s a -> (a -> State s b) -> State s b
bindState (State f) !g = State $ \s ->
  case f s of
    StateResult value state ->
      case g value of
        State next -> next state

pureState :: a -> State s a
pureState !a = State $ \s -> StateResult a s

get :: State s s
get = State $ \s -> StateResult s s

put :: s -> State s ()
put !s = State $ \_ -> StateResult () s

modify :: (s -> s) -> State s ()
modify !f = bindState get $ \s -> put (f s)

chainModifications :: Int -> State Int ()
chainModifications 0 = pureState ()
chainModifications n = bindState (modify (\x -> x + 1)) $ \_ -> chainModifications (n - 1)

stateManyTimes :: Int -> Int -> Int
stateManyTimes 0 !acc = acc
stateManyTimes n !acc =
  stateManyTimes (n - 1) (acc + stateValue (runState (chainModifications 60) 0))

runStateMonad :: Int -> Int
runStateMonad limit = stateManyTimes limit 0

-- 12. LazyEvaluation
-- Explicit Unit -> a closures match PureScript's non-memoizing Lazy newtype.
-- Do not replace these with Haskell lazy values or a sum of the chain lengths.
newtype Lazy a = Lazy (() -> a)

defer :: (() -> a) -> Lazy a
defer !f = Lazy f

force :: Lazy a -> a
force (Lazy f) = f ()

buildThunks :: Int -> Lazy Int -> Lazy Int
buildThunks 0 !acc = acc
buildThunks n !acc = buildThunks (n - 1) (defer $ \_ -> force acc + 1)

lazyManyTimes :: Int -> Int -> Int
lazyManyTimes 0 !acc = acc
lazyManyTimes n !acc =
  lazyManyTimes (n - 1) (acc + force (buildThunks 1000 (defer $ \_ -> 0)))

runLazyEvaluation :: Int -> Int
runLazyEvaluation limit = lazyManyTimes limit 0

-- 13. ArrayOps
-- The installed toolchain has array rather than vector. Materialize immutable
-- integer arrays before and after filtering, then use the library fold.
-- UArray gives eager integer elements, matching PureScript array evaluation.
arrayRange :: Int -> Int -> Array.UArray Int Int
arrayRange start end =
  let step = if start <= end then 1 else -1
  in Array.listArray (0, abs (end - start)) [start, start + step .. end]

arrayFilterEvens :: Array.UArray Int Int -> Array.UArray Int Int
arrayFilterEvens arr =
  let values = filter (\x -> x `mod` 2 == 0) (Array.elems arr)
  in Array.listArray (0, length values - 1) values

runArrayOps :: Int -> Int
runArrayOps limit = List.foldl' (+) 0 (Array.elems (arrayFilterEvens (arrayRange 1 limit)))

-- 14. RowToList
-- Haskell has no built-in PureScript-style row conversion. A typed,
-- heterogeneous record carries its row as a type-level list of (label, type)
-- pairs; the class follows that list recursively, exactly as RecordKeys does.
data Record (row :: [(Symbol, Type)]) where
  RNil :: Record '[]
  (:&) :: !a -> !(Record tail) -> Record ('(label, a) ': tail)
infixr 5 :&

class RecordKeys (row :: [(Symbol, Type)]) where
  keysImpl :: Proxy row -> Int

instance RecordKeys '[] where
  keysImpl _ = 0

instance RecordKeys tail => RecordKeys ('(label, a) ': tail) where
  keysImpl _ = 1 + keysImpl (Proxy @tail)

keys :: forall row. RecordKeys row => Record row -> Int
keys !_ = keysImpl (Proxy @row)

sampleRecord :: Record '[ '("a", Int), '("b", String), '("c", Bool), '("d", Double), '("e", String)]
sampleRecord = 1 :& "two" :& True :& 4.0 :& "five" :& RNil

runRowToList :: Int -> Int
runRowToList !_ = keys sampleRecord

-- Benchmarking Framework
-- Read the argument and consume the forced result on every invocation. Keeping
-- this IO boundary out of line prevents sharing a pure result across a batch.
{-# NOINLINE runOnce #-}
runOnce :: (Int -> Int) -> IORef Int -> IORef Int -> IO Int
runOnce act input output = do
  arg <- readIORef input
  result <- evaluate (act arg)
  writeIORef output result
  return result

measureBatch :: Int -> (Int -> Int) -> IORef Int -> IORef Int -> IO Word64
measureBatch count act input output = do
  start <- getMonotonicTimeNSec
  let loop 0 = return ()
      loop !remaining = do
        _ <- runOnce act input output
        loop (remaining - 1)
  loop count
  end <- getMonotonicTimeNSec
  return (end - start)

calibrate :: (Int -> Int) -> IORef Int -> IORef Int -> IO Int
calibrate act input output = go 1
  where
    go !count = do
      elapsed <- measureBatch count act input output
      if elapsed >= 10000000 || count >= 16777216
        then return count
        else go (count * 2)

{-# NOINLINE bench #-}
bench :: String -> (Int -> Int) -> Int -> IO Double
bench name act arg = do
  input <- newIORef arg
  output <- newIORef 0
  putStrLn $ "--------------------------------------------------\n\n(Test)\n" ++ name
  putStrLn "\n(Output & Warm-up)"
  result <- runOnce act input output
  print result
  replicateM_ 2 (runOnce act input output)

  count <- calibrate act input output
  durations <- replicateM 10 (measureBatch count act input output)
  let us = fromIntegral (minimum durations) / fromIntegral count / 1000.0
  putStrLn $ "\n(Execution time - best of 10)\n\n" ++ printf "%.6f" us ++ " μs\n"
  putStrLn $ "Batch iterations: " ++ show count ++ "\n"
  return us

-- Keep the validation keys, labels and inputs together so --check-only and
-- measurement execute the same kernels. Command-line flags never alter inputs.
cases :: [(String, String, Int -> Int, Int)]
cases =
  [ ("AstTree", "AST Evaluation:", runAstTree, 3)
  , ("Fib", "Fibonacci:", runFib, 10)
  , ("ListOps", "List Processing (900 elements):", runListOps, 900)
  , ("TCO", "Tail Call Optimization (100k calls):", runTCO, 100000)
  , ("Records", "Deep Record Updates (10k iterations):", runRecords, 10000)
  , ("Ackermann", "Ackermann (3, 4):", runAckermann, 3)
  , ("Church", "Church Numerals (100k Closure Applications):", runChurch, 10)
  , ("Primes", "Prime Sieve (sum primes up to 500):", runPrimes, 500)
  , ("RBTree", "Red-Black Tree (100k Worst-Case Insertions):", runRBTree, 100000)
  , ("Polymorphism", "Polymorphism (10M Type Class Dict Lookups):", runPolymorphism, 10000000)
  , ("StateMonad", "State Monad (1.2k Binds, 60 Stack Depth):", runStateMonad, 20)
  , ("LazyEvaluation", "Lazy Evaluation (1M Thunks Forced, 1k Depth):", runLazyEvaluation, 1000)
  , ("ArrayOps", "Array Processing (900 elements):", runArrayOps, 900)
  , ("RowToList", "RowToList (Keys Count):", runRowToList, 10000)
  ]

main :: IO ()
main = do
  args <- getArgs
  case args of
    ["--check-only"] -> forM_ cases $ \(key, _, act, arg) -> do
      value <- evaluate (act arg)
      putStrLn (key ++ "=" ++ show value)
    [] -> do
      putStrLn "Global warm-up in progress..."
      output <- newIORef 0
      replicateM_ 3 $ forM_ cases $ \(_, _, act, arg) -> do
        input <- newIORef arg
        _ <- runOnce act input output
        return ()
      total <- sum <$> mapM (\(_, name, act, arg) -> bench name act arg) cases
      putStrLn "\n==================================================\n"
      putStrLn $ "Total exec time: " ++ printf "%.6f" (total / 1000.0) ++ " ms\n"
    _ -> fail "Usage: bench_haskell [--check-only]"
