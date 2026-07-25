-module(control_extend@foreign).
-export([arrayExtend/2]).

arrayExtend(F, Xs) ->
    XsL = array:to_list(Xs),
    array:map(fun (I, _) -> 
        F(array:from_list(lists:nthtail(I, XsL)))
    end, Xs).
