module Main where

data Option t = None | Some t

foldOption :: (() -> r) -> (t -> r) -> Option t -> r
foldOption onNone _onSome None = onNone ()
foldOption _onNone onSome (Some a) = onSome a

inspect :: Show a => Option a -> IO ()
inspect = foldOption onNone onSome
  where
    onNone () = putStrLn "No"
    onSome a = putStrLn ("Yes " ++ show a)

main :: IO ()
main = do
  let v = Some 1
  let n = None :: Option Char
  inspect v
  inspect n
