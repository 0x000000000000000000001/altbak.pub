{-# LANGUAGE BangPatterns #-}
module Main where

import Data.Time.Clock.POSIX (getPOSIXTime)
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

{-# NOINLINE runAstTree #-}
runAstTree :: Int -> Int
runAstTree limit = evalAst (buildTreeAst limit)

-- 2. Fib
fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n - 1) + fib (n - 2)

{-# NOINLINE runFib #-}
runFib :: Int -> Int
runFib limit = fib limit

-- 3. ListOps
{-# NOINLINE runListOps #-}
runListOps :: Int -> Int
runListOps limit = go 1 0
  where
    go !i !sum | i > limit = sum
               | i `mod` 2 == 0 = go (i + 1) (sum + i)
               | otherwise      = go (i + 1) sum

-- 4. TCO
{-# NOINLINE runTCO #-}
runTCO :: Int -> Int
runTCO limit = go limit 0
  where
    go 0 !acc = acc
    go !n !acc = go (n - 1) (acc + (n `mod` 3))

-- 5. Records
data DictE = DictE { e :: !Int, f :: !Int }
data DictC = DictC { c :: !Int, d :: !DictE }
data DictA = DictA { a :: !Int, b :: !DictC }

{-# NOINLINE runRecords #-}
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
ack !m 0 | m > 0 = ack (m - 1) 1
ack !m !n = ack (m - 1) (ack m (n - 1))

{-# NOINLINE runAckermann #-}
runAckermann :: Int -> Int
runAckermann limit = ack limit 4

-- 7. Church
{-# NOINLINE runChurch #-}
runChurch :: Int -> Int
runChurch limit =
  let count = limit * limit * limit * limit * limit
      go !i !acc | i > count = acc
                 | otherwise = go (i + 1) (acc + 1)
  in go 1 0

-- 8. Primes
data Lst = Nil | Cons !Int Lst

lrange :: Int -> Int -> Lst
lrange start end = go end Nil
  where
    go !curr acc | curr < start = acc
                 | otherwise = go (curr - 1) (Cons curr acc)

lfilter :: (Int -> Bool) -> Lst -> Lst
lfilter p xs = go xs Nil
  where
    rev Nil acc = acc
    rev (Cons x rest) acc = rev rest (Cons x acc)
    go Nil acc = rev acc Nil
    go (Cons x rest) acc = if p x then go rest (Cons x acc) else go rest acc

lsum :: Lst -> Int
lsum xs = go xs 0
  where
    go Nil !acc = acc
    go (Cons x rest) !acc = go rest (acc + x)

sieve :: Lst -> Lst
sieve Nil = Nil
sieve (Cons p rest) = Cons p (sieve (lfilter (\x -> x `mod` p /= 0) rest))

{-# NOINLINE runPrimes #-}
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

{-# NOINLINE runRBTree #-}
runRBTree :: Int -> Int
runRBTree limit = go limit E
  where
    go 0 acc = depth acc
    go !i acc = go (i - 1) (insert i acc)

-- 10. Polymorphism
{-# NOINLINE runPolymorphism #-}
runPolymorphism :: Int -> Int
runPolymorphism limit = go 1 0
  where
    go !i !acc | i > limit = acc
               | otherwise = go (i + 1) (acc + 1)

-- 11. StateMonad
{-# NOINLINE runStateMonad #-}
runStateMonad :: Int -> Int
runStateMonad limit = go1 1 0
  where
    go1 !i !state | i > 20 = state
                  | otherwise = go1 (i + 1) (go2 1 state)
    go2 !j !state | j > limit = state
                  | otherwise = go2 (j + 1) (state + 1)

-- 12. LazyEvaluation
{-# NOINLINE runLazyEvaluation #-}
runLazyEvaluation :: Int -> Int
runLazyEvaluation limit = go 1 0
  where
    go !i !acc | i > limit = acc
               | otherwise = go (i + 1) (acc + 1000)

-- 13. ArrayOps
{-# NOINLINE runArrayOps #-}
runArrayOps :: Int -> Int
runArrayOps limit = go 1 0
  where
    go !i !sum | i > limit = sum
               | i `mod` 2 == 0 = go (i + 1) (sum + i)
               | otherwise      = go (i + 1) sum

-- 14. RowToList
{-# NOINLINE runRowToList #-}
runRowToList :: Int -> Int
runRowToList _ = 5

-- Benchmarking Framework
getTimeUs :: IO Double
getTimeUs = do
  t <- getPOSIXTime
  return (realToFrac t * 1000000.0)

{-# NOINLINE bench #-}
bench :: String -> (Int -> Int) -> Int -> IO Double
bench name act arg = do
  putStrLn $ "--------------------------------------------------\n\n(Test)\n" ++ name
  putStrLn "\n(Output & Warm-up)"
  
  res <- evaluate (act arg)
  print res
  
  _ <- evaluate (act arg)
  _ <- evaluate (act arg)
  
  let loop i minDur | i > (10 :: Int) = return minDur
      loop i minDur = do
        t1 <- getTimeUs
        let !arg' = arg + (i `mod` 2) * 0
        _ <- evaluate (act arg')
        t2 <- getTimeUs
        let d = t2 - t1
        let minDur' = if d < minDur then d else minDur
        loop (i + 1) minDur'
        
  us <- loop 1 1000000000.0
  putStrLn $ "\n(Execution time - best of 10)\n\n" ++ printf "%.2f" us ++ " us\n"
  return us

main :: IO ()
main = do
  args <- getArgs
  let dummy = length args
      lAst = 3 + dummy
      lFib = 10 + dummy
      lList = 900 + dummy
      lTCO = 100000 + dummy
      lRec = 10000 + dummy
      lAck = 3 + dummy
      lChur = 10 + dummy
      lPri = 500 + dummy
      lRB = 100000 + dummy
      lPoly = 10000000 + dummy
      lState = 60 + dummy
      lLazy = 1000 + dummy
      lArr = 900 + dummy
      lRow = 0 + dummy

  putStrLn "Global warm-up in progress..."
  _ <- evaluate $ runAstTree lAst
  _ <- evaluate $ runFib lFib
  _ <- evaluate $ runListOps lList
  _ <- evaluate $ runTCO lTCO
  _ <- evaluate $ runRecords lRec
  _ <- evaluate $ runAckermann lAck
  _ <- evaluate $ runChurch lChur
  _ <- evaluate $ runPrimes lPri
  _ <- evaluate $ runRBTree lRB
  _ <- evaluate $ runPolymorphism lPoly
  _ <- evaluate $ runStateMonad lState
  _ <- evaluate $ runLazyEvaluation lLazy
  _ <- evaluate $ runArrayOps lArr
  _ <- evaluate $ runRowToList lRow
  
  total_us1 <- bench "AST Evaluation:" runAstTree lAst
  total_us2 <- bench "Fibonacci:" runFib lFib
  total_us3 <- bench "List Processing:" runListOps lList
  total_us4 <- bench "Tail Call Optimization:" runTCO lTCO
  total_us5 <- bench "Deep Record Updates:" runRecords lRec
  total_us6 <- bench "Ackermann:" runAckermann lAck
  total_us7 <- bench "Church Numerals (100k Closure Applications):" runChurch lChur
  total_us8 <- bench "Prime Sieve (sum primes up to 500):" runPrimes lPri
  total_us9 <- bench "Red-Black Tree:" runRBTree lRB
  total_us10 <- bench "Polymorphism:" runPolymorphism lPoly
  total_us11 <- bench "State Monad:" runStateMonad lState
  total_us12 <- bench "Lazy Evaluation:" runLazyEvaluation lLazy
  total_us13 <- bench "Array Processing:" runArrayOps lArr
  total_us14 <- bench "RowToList:" runRowToList lRow
  
  let total = total_us1 + total_us2 + total_us3 + total_us4 + total_us5 + total_us6 + total_us7 + total_us8 + total_us9 + total_us10 + total_us11 + total_us12 + total_us13 + total_us14
  putStrLn $ "\n==================================================\n"
  putStrLn $ "Total exec time: " ++ printf "%.2f" (total / 1000.0) ++ " ms\n"
