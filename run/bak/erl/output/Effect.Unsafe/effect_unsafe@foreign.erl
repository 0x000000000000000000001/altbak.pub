-module(effect_unsafe@foreign).
-export([unsafePerformEffect/1]).

unsafePerformEffect(F) -> F().
