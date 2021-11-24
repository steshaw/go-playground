module Main where

data Option t = None | Some t

-- foldoption ::
-- p
--
inspect :: Show a => Option a -> IO ()
inspect None = putStrLn "No"
inspect (Some a) = putStrLn ("Yes " ++ show a)

main :: IO ()
main = do
  let v = Some 1
  let n = None :: Option Char
  inspect v
  inspect n
