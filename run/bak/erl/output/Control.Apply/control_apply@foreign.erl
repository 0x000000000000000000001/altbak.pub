-module(control_apply@foreign).
-export([arrayApply/2]).

arrayApply(Fs, Xs) ->
  Arr = array:new(array:size(Fs)*array:size(Xs)),
  {_, Result} = array:foldl(
    fun (_, F, AccN) ->
      array:foldl(
        fun (_, X, {N, Acc1}) ->
          {N+1, array:set(N, F(X), Acc1)}
        end,
        AccN,
        Xs
      )
    end,
    {0, Arr},
    Fs
  ),
  Result.
