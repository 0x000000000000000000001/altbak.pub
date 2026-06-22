-module(effect_ref@foreign).
-behavior(gen_server).
-export([modifyImpl/2, new/1, newWithSelf/1, read/1, write/2]).
-export([code_change/3, handle_call/3, handle_cast/2, handle_info/2, init/1, terminate/2]).
-purs_ignore_exports([code_change/3, handle_call/3, handle_cast/2, handle_info/2, init/1, terminate/2]).

modifyImpl(F, Ref) ->
    fun () ->
        gen_server:call(Ref, {modify, F})
    end.

new(Value) ->
    fun () ->
        {ok, Pid} = gen_server:start(?MODULE, Value, []),
        Pid
    end.

newWithSelf(MkValue) ->
    fun () ->
        {ok, Pid} = gen_server:start(?MODULE, undefined, []),
        Value = MkValue(Pid),
        gen_server:call(Pid, {write, Value}),
        Pid
    end.

read(Ref) ->
    fun () ->
        gen_server:call(Ref, read)
    end.

write(Value, Ref) ->
    fun () ->
        gen_server:call(Ref, {write, Value})
    end.

%% gen_server callbacks

code_change(_OldVersion, State, _Extra) -> {ok, State}.

init(Value) ->
    {ok, Value}.

handle_call({modify, F}, _From, Value) ->
    #{ state := NewState, value := NewValue } = F(Value),
    {reply, NewValue, NewState};
handle_call(read, _From, State) -> {reply, State, State};
handle_call({write, State}, _From, _State) -> {reply, unit, State}.

handle_cast(_Request, State) -> {noreply, State}.

handle_info(_Info, State) -> {noreply, State}.

terminate(_Reason, _State) -> ok.
