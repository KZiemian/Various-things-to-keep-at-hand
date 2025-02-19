-- Comments

{-
Comments
-}

-- main = putStrLn "Hello, Haskell!"

import Data.List
import System.IO

-- Int
-- maxInt = maxBound :: Int
-- minInt = minBound :: Int

-- Integer
-- Float
-- Double
-- bigFloat1 = 3.99999999999 + 0.00000000005
-- bigFloat2 = 3.999999999999 + 0.00000000005

-- Bool
-- Char
-- Tuple

-- always5 :: Int
-- always5 = 5

-- sumOfNums = sum [1..1000]

-- addEx = 5 + 5
-- subEx = 5 - 3
-- mulEx = 2 *5
-- divEx = 7 / 3

-- modEx = mod 5 4
-- modEx1 = 5 `mod` 4
-- negNumEx = 7 + (-4)

-- num1 = 9 :: Int

-- sqrtof9 = sqrt (fromIntegral num1)

-- piVal = pi
-- ePow9 = exp 9
-- logOf9 = log 9

-- squared9 = 9 ** 2
-- truncateVal = truncate 9.999
-- roundVal = round 9.999
-- ceilingVar = ceiling 9.999
-- floorVal = floor 9.999

-- trueAndFalse = True && False
-- trueOrFalse = True || False
-- notTrue = not(True)

primeNumbers = [3, 5, 7, 11]

-- primeNumbers1 = primeNumbers ++ [13, 17, 19, 23]

-- favNums = 2 : 7 : 21 : 66 : []

-- multList = [[3, 5, 7], [11, 13, 17]]

morePrimes2 = 2 : primeNumbers
lenPrime = length morePrimes2

revPrime = reverse morePrimes2
isListEmpty = null morePrimes2

secondPrime = morePrimes2 !! 1

firstPrime = head morePrimes2
lastPrime = last morePrimes2

primeInit = init morePrimes2

-- main = putStrLn firstPrime
-- main = putStrLn lastPrime

-- first3Primes = take 3 morePrimes2

-- removedPrimes = drop 3 morePrimes2

-- is7InList = 7 `elem` morePrimes2
-- is3InList = elem 3 morePrimes2

-- maxPrime = maximum morePrimes2
-- minPrime = minimum morePrimes2

-- newList = [2, 3, 5]

-- prodPrimes = product newList

-- zeroToTen = [0..10]
-- evenList = [2,4..20]

letterList = ['A', 'C'..'Z']

infinitPow10 = [10, 20..]

many2s = take 10 (repeat 2)
