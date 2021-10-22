module Main where

data Option t = None | Some t

data FoldAlg t r = FoldAlg {
  faOnNone :: () -> r,
  faOnSome :: t -> r
  }

foldOption :: FoldAlg t r -> Option t -> r
foldOption (FoldAlg onNone _onSome) None = onNone ()
foldOption (FoldAlg _onNone onSome) (Some a) = onSome a

inspect :: Show a => Option a -> IO ()
inspect = foldOption (FoldAlg onNone onSome)
  where
    onNone () = putStrLn "No"
    onSome a = putStrLn ("Yes " ++ show a)

main :: IO ()
main = do
  let v = Some 1
  let n = None :: Option Char
  inspect v
  inspect n
