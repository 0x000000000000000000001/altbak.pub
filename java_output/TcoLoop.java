public class TcoLoop extends RuntimeException {
    public String loopId;
    public Object[] args;
    public TcoLoop(String loopId, Object[] args) {
        this.loopId = loopId;
        this.args = args;
    }
    @Override
    public synchronized Throwable fillInStackTrace() { return this; }
}
