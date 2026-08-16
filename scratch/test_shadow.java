public class test_shadow {
  public static void main(String[] args) {
    Object x = 1;
    java.util.function.Supplier<Object> s = new java.util.function.Supplier<Object>() {
      public Object get() {
        Object x = 2;
        return x;
      }
    };
    System.out.println(s.get());
  }
}
