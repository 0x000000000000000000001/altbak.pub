    private record StateResult<S, A>(A value, S state) {}

    @FunctionalInterface
    private interface State<S, A> {
        StateResult<S, A> run(S initial);
    }

    private static <S, A, B> State<S, B> bindState(State<S, A> action, java.util.function.Function<A, State<S, B>> next) {
        return initial -> {
            StateResult<S, A> first = action.run(initial);
            return next.apply(first.value()).run(first.state());
        };
    }

    private static <S, A> State<S, A> pureState(A value) {
        return initial -> new StateResult<>(value, initial);
    }

    private static <S> State<S, S> getState() {
        return initial -> new StateResult<>(initial, initial);
    }

    private static <S> State<S, Void> putState(S value) {
        return ignored -> new StateResult<>(null, value);
    }

    private static <S> State<S, Void> modifyState(java.util.function.Function<S, S> update) {
        State<S, S> get = getState();
        return bindState(get, current -> putState(update.apply(current)));
    }

    private static State<Integer, Void> chainModifications(int depth) {
        if (depth == 0) return pureState(null);
        return bindState(modifyState(value -> value + 1), ignored -> chainModifications(depth - 1));
    }

    public static final java.util.function.Function<Object, Object> runStateMonadFFI = Bench.nativeBenchmark(input -> {
        // The shared FFI wrapper supplies depth 60; the PureScript test supplies 20 repetitions.
        int depth = (Integer) input;
        int result = 0;
        for (int i = 0; i < 20; i++) {
            result += chainModifications(depth).run(0).state();
        }
        return result;
    });
