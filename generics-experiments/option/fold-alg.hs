{-# LANGUAGE RankNTypes #-}

module Main where

data FoldAlg t r = FoldAlg {
  faOnNone :: () -> r,
  faOnSome :: t -> r
  }

newtype Option t = Option {
  unOption :: forall r. (() -> r) -> (t -> r) -> r
  }

none :: () -> Option t
none () = Option $ \onNone _onSome -> onNone ()

some :: t -> Option t
some a = Option $ \_onNone onSome -> onSome a

optionAlg :: FoldAlg t (Option t)
optionAlg = FoldAlg none some

inspect :: Show t => Option t -> IO ()
inspect option = unOption option onNone onSome
  where
    onNone () = putStrLn "No"
    onSome a = putStrLn ("Yes " ++ show a)

main :: IO ()
main = do
  let v = some 1
  let n = none () :: Option Char
  inspect v
  inspect n
