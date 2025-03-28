import Data.List
import System.IO

-- many3s = replicate 10 3

-- cycleList = take 10 (cycle [1, 2, 3, 4, 5])

-- listTimes2 = [x * 2 | x <- [1..10]]

-- listTimes3 = [x * 3 | x <- [1..10], x * 3 <= 20]

-- divisibleBy9And13 = [x | x <- [1..500], x `mod` 13 == 0 && x `mod` 9 == 0]

-- divisibleBy9And13 = [x | x <- [1..500], mod x 13 == 0 && mod x 9 == 0]

-- divisibleBy9And13 = [x | x <- [1..500], mod x 13 == 0, mod x 9 == 0]

-- sortedList = sort [9, 1, 8, 3, 4, 7, 6]

-- sumOfLists = zipWith (+) [1, 2, 3, 4, 5] [6, 7, 8, 9, 10]

-- primeNumbers = [2, 3, 5, 7, 11, 13, 17, 19, 23]

-- listBiggerThen5 = filter (> 5) primeNumbers

-- evensUpTo20 = takeWhile (<= 20) [2, 4..]

-- multOfList = foldl (*) 1 [2, 3, 4, 5]

-- multOfList1 = foldr (*) 1 [2, 3, 4, 5]

pow3List = [3^n | n <- [1..10]]

-- 23:44
-- https://www.youtube.com/watch?v=02_H3LjqMr8
