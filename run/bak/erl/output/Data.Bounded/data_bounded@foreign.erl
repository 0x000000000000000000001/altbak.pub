-module(data_bounded@foreign).
-export([topChar/0, bottomChar/0]).

topChar() -> 16#10FFFF.
bottomChar() -> 0.
