-module(record_unsafe@foreign).
-export([unsafeHas/2, unsafeGet/2, unsafeSet/3, unsafeDelete/2]).

unsafeHas(Label, Rec) -> maps:is_key(binary_to_atom(Label, utf8), Rec).
unsafeGet(Label, Rec) -> maps:get(binary_to_atom(Label, utf8), Rec).
unsafeSet(Label, Value, Rec) -> maps:put(binary_to_atom(Label, utf8), Value, Rec).
unsafeDelete(Label, Rec) -> maps:remove(binary_to_atom(Label, utf8), Rec).
