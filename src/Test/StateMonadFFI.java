    private record StateResult<A>(A value, int state) {}

    @FunctionalInterface
    private interface State<A> {
        StateResult<A> run(int initial);
    }

    private static <A, B> State<B> bindState(State<A> action, java.util.function.Function<A, State<B>> next) {
        return initial -> {
            StateResult<A> first = action.run(initial);
            return next.apply(first.value()).run(first.state());
        };
    }

    private static <A> State<A> pureState(A value) {
        return initial -> new StateResult<>(value, initial);
    }

    private static State<Integer> getState() {
        return initial -> new StateResult<>(initial, initial);
    }

    private static State<Void> putState(int value) {
        return ignored -> new StateResult<>(null, value);
    }

    private static State<Void> modifyState(java.util.function.IntUnaryOperator update) {
        return bindState(getState(), current -> putState(update.applyAsInt(current)));
    }

    private static State<Void> chainModifications(int depth) {
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
