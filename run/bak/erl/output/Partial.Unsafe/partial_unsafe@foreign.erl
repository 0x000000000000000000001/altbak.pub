% module Partial.Unsafe
-module(partial_unsafe@foreign).
-export(['_unsafePartial'/1]).

'_unsafePartial'(F) -> F(unit).
