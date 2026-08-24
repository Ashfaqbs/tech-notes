# Java Version Features Cheatsheet 

## Java 9 (Sep 21, 2017)

Key Features:

1. Modular System (Project Jigsaw):

- Why Added: To modularize the JDK and make it more scalable and maintainable.
- How It Helps: Allows developers to create and use modules, improves application performance, and enhances security.

Example:
```
module com.example.myapp {
    requires java.base;
    exports com.example.myapp;
}
```

2. New HTTP Client:

- Why Added: To replace the old HttpURLConnection with a more modern and flexible API.
- How It Helps: Supports HTTP/2 and WebSocket, asynchronous requests, and cleaner API.

Example:

```

HttpClient client = HttpClient.newHttpClient();
HttpRequest request = HttpRequest.newBuilder()
    .uri(URI.create("https://example.com"))
    .build();
HttpResponse<String> response = client.send(request, BodyHandlers.ofString());
System.out.println(response.body());
```
3. Java 9 REPL (JShell):

- Why Added: To provide a read-eval-print loop for quickly testing and experimenting with Java code.
- How It Helps: Enhances productivity and learning experience.

Example:
```
jshell> int x = 10;
jshell> System.out.println(x + 2);
```


## Java 10 (Mar 20, 2018)

Key Features:

1. Local Variable Type Inference:

- Why Added: To reduce boilerplate code by allowing the var keyword for local variable declarations.
- How It Helps: Makes code more concise and readable.

Example:
```
var list = new ArrayList<String>();
list.add("example");
```

2. Garbage Collector Interface:

- Why Added: To provide a common interface for all garbage collectors.
- How It Helps: Simplifies the implementation and management of different garbage collectors.

Example:
```
This feature is internal and doesn't have direct API usage.
```

3. App Class-Data Sharing:

- Why Added: To improve startup and footprint by sharing class data across Java processes.
- How It Helps: Reduces memory usage and improves startup time.


## Java 11 (Sep 25, 2018) - LTS

Key Features:

1. HTTP Client (Standard):

- Why Added: To stabilize the new HTTP Client API introduced in Java 9.
- How It Helps: Offers a more powerful and flexible HTTP client.

Example:

```
HttpClient client = HttpClient.newHttpClient();
HttpRequest request = HttpRequest.newBuilder()
    .uri(URI.create("https://example.com"))
    .build();
HttpResponse<String> response = client.send(request, BodyHandlers.ofString());
System.out.println(response.body());
```

2. Dynamic Class-File Constants:

- Why Added: To allow dynamic constants in the class file.
- How It Helps: Supports more efficient class file structures and reduces duplication.

3. Epsilon GC (No-Op Garbage Collector):

- Why Added: For low-latency and high-throughput applications where GC pauses are undesirable.
- How It Helps: Offers predictable performance by not performing any garbage collection.


## Java 12 (Mar 19, 2019)

Key Features:

1. Switch Expressions (Preview):

- Why Added: To simplify switch statements and make them more expressive.
- How It Helps: Reduces boilerplate code and adds more flexibility to switch cases.

Example:
```
int numLetters = switch (day) {
    case MONDAY, FRIDAY, SUNDAY -> 6;
    case TUESDAY -> 7;
    case THURSDAY, SATURDAY -> 8;
    case WEDNESDAY -> 9;
    default -> throw new IllegalStateException("Unexpected value: " + day);
};
```

2. Teeing Collectors:

- Why Added: To allow collecting elements in multiple ways and combining the results.
- How It Helps: Simplifies collecting and combining data.

Example:

```
Collector<T, ?, R> collector = Collectors.teeing(
    Collectors.summingInt(T::getValue),
    Collectors.counting(),
    (sum, count) -> new Result(sum, count)
);
```

## Java 13 (Sep 17, 2019)

Key Features:

1. Text Blocks (Preview):

- Why Added: To simplify the creation of multiline strings.
- How It Helps: Reduces the need for escape sequences and improves readability.

Example:
```
String json = """
    {
        "name": "John",
        "age": 30,
        "city": "New York"
    }
    """;
```

2. Enhancements in Switch Expressions:

- Why Added: To further enhance the switch expression feature.
- How It Helps: Improves code readability and reduces errors.

Example:

```
    int result = switch (input) {
    case 1 -> 10;
    case 2 -> 20;
    default -> 0;
};
```
3. ZGC Improvements:

- Why Added: To improve the Z Garbage Collector's performance and scalability.
- How It Helps: Provides better pause time and memory management.


## Java 14 (Mar 17, 2020)

Key Features:

1. Pattern Matching for instanceof (Preview):

- Why Added: To simplify the type-checking and casting.
- How It Helps: Reduces boilerplate code and improves readability.

Example:

```
if (obj instanceof String s) {
    System.out.println(s.toLowerCase());
}
```

2. Records (Preview):

- Why Added: To provide a compact syntax for declaring data classes.
- How It Helps: Reduces boilerplate code for immutable data carriers.

Example:
```
public record Point(int x, int y) {}
```

3. JFR Event Streaming:

- Why Added: To enable continuous monitoring and diagnostics.
- How It Helps: Provides real-time visibility into application performance.

Example:
```
This feature is more about configuration and usage rather than code.
```
## Java 15 (Sep 15, 2020)

Key Features:

1. Sealed Classes (Preview):

- Why Added: To restrict which classes can extend or implement a given class or interface.
- How It Helps: Enhances encapsulation and security.

Example:

```
public sealed class Shape
    permits Circle, Square, Rectangle {}

public final class Circle extends Shape {}
public final class Square extends Shape {}
public final class Rectangle extends Shape {}
```
2. Hidden Classes:

- Why Added: To support frameworks that generate classes dynamically.
- How It Helps: Improves security and reduces memory footprint.

3. EdDSA (Edwards-Curve Digital Signature Algorithm):

- Why Added: To support new digital signature algorithms.
- How It Helps: Enhances cryptographic capabilities.

## Java 16 (Mar 16, 2021)

Key Features:

1. Records (Standard):

- Why Added: To stabilize the record feature introduced in Java 14.
- How It Helps: Provides a standard way to create immutable data classes.

Example:
```
public record Person(String name, int age) {}
```
2. Pattern Matching for instanceof (Standard):

- Why Added: To stabilize the pattern matching feature introduced in Java 14.
- How It Helps: Simplifies type-checking and casting.

Example:
```
if (obj instanceof String s) {
    System.out.println(s.toUpperCase());
}
```
3. Vector API (Incubator):

- Why Added: To provide a mechanism for expressing vector computations.
- How It Helps: Enables high-performance computing.

Example:
```
Vector<Integer> vector = Vector.of(1, 2, 3, 4);

```



## Java 17 (Sep 14, 2021) - LTS

Key Features:

1. Pattern Matching for switch (Preview):
- Why Added: To allow pattern matching in switch statements.
- How It Helps: Enhances the expressiveness and readability of switch statements.

 Example:
```
switch (obj) {
    case Integer i -> System.out.println("Integer: " + i
```


2. Sealed Classes (Standard):

- Why Added: To stabilize the sealed classes feature introduced in Java 15.
- How It Helps: Restricts which classes can extend or implement a given class or interface, enhancing security and encapsulation.

Example:
```
public sealed class Shape
    permits Circle, Square, Rectangle {}

public final class Circle extends Shape {}
public final class Square extends Shape {}
public final class Rectangle extends Shape {}
```

3. Foreign Function & Memory API (Incubator):

- Why Added: To provide a mechanism to interoperate with code and data outside of the JVM.
- How It Helps: Allows Java programs to call native libraries and manage memory outside of the Java heap.

Example:

```
try (MemorySegment segment = MemorySegment.allocateNative(100)) {
    MemoryAddress address = segment.baseAddress();
    // Use the address to interact with native memory
}
```

## Java 18 (Mar 22, 2022)

Key Features:

1. UTF-8 by Default:

- Why Added: To standardize character encoding.
- How It Helps: Ensures consistency across different environments and simplifies internationalization.

Example:
```
No direct code example needed, but it affects file operations and string handling.
```
2. Simple Web Server:

- Why Added: To provide a simple, out-of-the-box web server.
- How It Helps: Useful for prototyping, testing, and serving static content.

Example:
```
var server = SimpleFileServer.createFileServer(
    new InetSocketAddress(8000),
    Path.of("public"),
    SimpleFileServer.OutputLevel.VERBOSE
);
server.start();
```
3. Code Snippets in API Documentation:

- Why Added: To enhance Javadoc with executable snippets.
- How It Helps: Improves documentation quality and accuracy.

Example:
```
/**
 * {@snippet :
 * System.out.println("Hello, World!");
 * }
 */
public class Example {}
```

## Java 19 (Sep 20, 2022)

Key Features:

1. Record Patterns (Preview):

- Why Added: To enhance pattern matching with records.
- How It Helps: Simplifies data extraction from records.

Example:

```
record Point(int x, int y) {}
var point = new Point(1, 2);
if (point instanceof Point(int x, int y)) {
    System.out.println("x: " + x + ", y: " + y);
}
```

2. Virtual Threads (Preview):

- Why Added: To introduce lightweight threads.
- How It Helps: Improves the scalability of concurrent applications.

Example:
```
try (var scope = newVirtualThreadPerTaskExecutor()) {
    scope.submit(() -> System.out.println("Hello from virtual thread"));
}
```
3. Structured Concurrency (Incubator):

- Why Added: To simplify multithreaded programming.
- How It Helps: Makes concurrent code easier to read and maintain.

Example:
```
try (var scope = new StructuredTaskScope.ShutdownOnFailure()) {
    Future<String> future = scope.fork(() -> "Result");
    scope.join();  // join all tasks
    scope.throwIfFailed();  // propagate errors
    System.out.println(future.resultNow());
}
```

## Java 20 (Mar 21, 2023)

Key Features:

1. Scoped Values (Incubator):

- Why Added: To provide a better alternative to thread-local variables.
- How It Helps: Makes it easier to share immutable data across threads.

Example:
```
public static final ScopedValue<String> MY_VALUE = ScopedValue.newInstance();

ScopedValue.where(MY_VALUE, "Scoped data").run(() -> {
    System.out.println(MY_VALUE.get());
});
```

2. Virtual Threads (Preview):

- Why Added: Continued work on virtual threads.
- How It Helps: Further stabilizes and improves the feature for scalable concurrency.

Example:
```
var thread = Thread.ofVirtual().start(() -> System.out.println("Running in a virtual thread"));
thread.join();
```

## Java 21 (Sep 19, 2023) - LTS

Key Features:

1. Sequenced Collections:

- Why Added: To provide collections with predictable iteration order.
- How It Helps: Simplifies code that relies on iteration order.

Example:

```
SequencedCollection<String> collection = SequencedCollections.create();
collection.add("one");
collection.add("two");
collection.forEach(System.out::println);
```
2. Record Patterns:

- Why Added: To finalize the record patterns feature.
- How It Helps: Simplifies pattern matching with records.

Example:

```
record Point(int x, int y) {}
var point = new Point(1, 2);
if (point instanceof Point(int x, int y)) {
    System.out.println("x: " + x + ", y: " + y);
}
```
3. Virtual Threads:

- Why Added: To stabilize virtual threads.
- How It Helps: Makes concurrent programming more scalable.

Example:
```
var thread = Thread.ofVirtual().start(() -> System.out.println("Running in a virtual thread"));
thread.join();
```
4. Key Encapsulation Mechanism API:

- Why Added: To support cryptographic key encapsulation.
- How It Helps: Enhances security by simplifying key exchange.

Example:
```
KeyGenerator keyGen = KeyGenerator.getInstance("AES");
keyGen.init(256);
SecretKey secretKey = keyGen.generateKey();
```

## Java 22 (Mar 19, 2024)

Key Features:

1. Region Pinning for G1:

- Why Added: To improve the performance of the G1 garbage collector.
- How It Helps: Reduces pause times and enhances performance for applications with large heaps.

2. Foreign Function & Memory API (Standard):

- Why Added: To stabilize the API for interoperation with native code and memory.
- How It Helps: Provides a reliable way to manage memory and call native functions from Java.

Example:
```
try (MemorySegment segment = MemorySegment.allocateNative(100)) {
    MemoryAddress address = segment.baseAddress();
    // Use the address to interact with native memory
}
```
3. Unnamed Var & Patterns:

- Why Added: To support unnamed variables and pattern matching.
- How It Helps: Simplifies code by reducing the need for explicit variable names.

Example:

```
var list = List.of(1, 2, 3);
list.forEach(_ -> System.out.println(_));
```




Here are summaries for **Java 23** and **Java 24** in the same style, with key features + “why / how it helps” + examples.

---

## Java 23 (Sep 17, 2024)

**Key Features:**

1. **Primitive Types in Patterns, `instanceof`, and `switch` (Preview, JEP 455)**

   * *Why Added:* To make type checking & pattern matching more uniform, allowing primitives as well as reference types. Reduces boilerplate and avoids unneeded boxing/unboxing.
   * *How It Helps:* You can write `if (x instanceof int i)` etc., or match primitive values in `switch` directly. Improves clarity and safety.

   **Example:**

   ```java
   int value = 42;
   if (value instanceof int i && i > 0) {
       System.out.println("Positive int: " + i);
   }

   double d = 3.14;
   String description = switch (d) {
       case int i -> "Integer-like: " + i;
       case double dd -> "Double: " + dd;
       default -> "Other primitive";
   };
   System.out.println(description);
   ```

2. **Markdown Documentation Comments (JEP 467)**

   * *Why Added:* To let developers write documentation comments in Markdown, which is more readable / familiar than raw HTML + `@`-tags for many.
   * *How It Helps:* Easier to write and read docs; supports code blocks, lists, links, etc.

   **Example:**

   ```java
   /// This method computes factorial of a number
   /// 
   /// * `n` : the number whose factorial is needed  
   /// * returns: factorial of `n`
   public static long factorial(int n) {
       return (n <= 1) ? 1 : n * factorial(n - 1);
   }
   ```

3. **Module Import Declarations (Preview, JEP 476)**

   * *Why Added:* To simplify imports when many packages from a module are used. Rather than multiple `import package.Class`, you can import the whole module.
   * *How It Helps:* Reduces verbosity when using many classes from a module.

   **Example:**

   ```java
   import module java.util;

   public class MyApp {
       public void process() {
           List<String> list = List.of("a", "b", "c");
           Set<Integer> set = Set.of(1, 2, 3);
           // No need to import java.util.List, java.util.Set separately
       }
   }
   ```

4. **Implicitly Declared Classes and Instance Main Methods (Third Preview, JEP 477)**

   * *Why Added:* To allow simpler “script-style” programs, where you don’t need explicit public class / static main signature. For beginners and small tools.
   * *How It Helps:* You can write just `void main()` (or with minimal boilerplate) and run small programs more cleanly.

   **Example:**

   ```java
   void main() {
       println("Hello, Java 23!");
   }
   ```

5. **Flexible Constructor Bodies (Second Preview, JEP 482)**

   * *Why Added:* Previously, `super(...)` or `this(...)` had to be the first statement in constructor. This restricted argument validation / field initialization before calling parent.
   * *How It Helps:* Let you put some code before calling `super(...)`, for example input checks, setting fields that do not depend on the instance yet, etc.

   **Example:**

   ```java
   class Base {
       Base(String key) { /*...*/ }
   }

   class Derived extends Base {
       String key;

       Derived(String key) {
           if (key == null || key.isEmpty()) {
               throw new IllegalArgumentException("key required");
           }
           this.key = key.toUpperCase();
           super(key);  // allowed now after preliminary checks / assignments
       }
   }
   ```

6. **Stream Gatherers (Second Preview, JEP 473)**

   * *Why Added:* To make easier aggregation / collection of streams in flexible ways.
   * *How It Helps:* Simplifies writing operations like “grouping”, “partitioning”, or other custom collection logic from streams.

   **Example:**

   ```java
   var result = Stream.of(1,2,3,4,5,6)
       .gathering(Collectors.groupingBy(n -> n % 2 == 0 ? "even" : "odd"))
       .get();
   System.out.println(result);  // {odd=[1,3,5], even=[2,4,6]}
   ```

7. **Vector API Enhancements (Incubator, JEP 469)**

   * *Why Added:* For performance, especially numeric / data-parallel operations. Vector API gives access to SIMD (single instruction multiple data) operations in a portable way.
   * *How It Helps:* More efficient loops / array operations using vectorization.

   **Example:**

   ```java
   int[] a = {1,2,3,4};
   int[] b = {5,6,7,8};
   int[] c = new int[a.length];

   VectorSpecies<Integer> species = IntVector.SPECIES_PREFERRED;
   for (int i = 0; i < a.length; i += species.length()) {
       var va = IntVector.fromArray(species, a, i);
       var vb = IntVector.fromArray(species, b, i);
       var vc = va.add(vb);
       vc.intoArray(c, i);
   }
   // c now [6,8,10,12]
   ```

8. **ZGC: Generational Mode by Default (JEP 474)**

   * *Why Added:* To reduce GC pauses and improve memory reclamation in large heaps.
   * *How It Helps:* Younger garbage gets collected more frequently, reducing long pause times.

---

## Java 24 (Mar 18, 2025)

**Key Features:**

1. **Pattern Matching: Primitive Types in Patterns, `instanceof`, and `switch` (Second Preview, JEP 488)**

   * *Why Added:* Builds on Java 23’s preview to refine and stabilize use of primitive types in pattern matching.
   * *How It Helps:* As in Java 23, but with more polish / better support, fewer rough edges.

   **Example:**

   ```java
   Object obj = 10;
   if (obj instanceof int i && i % 2 == 0) {
       System.out.println("Even int: " + i);
   }

   long status = 2L;
   String msg = switch (status) {
       case 0L -> "Zero";
       case 1L -> "One";
       case long l -> "Other long: " + l;
   };
   System.out.println(msg);
   ```

2. **Flexible Constructor Bodies (Third Preview, JEP 492)**

   * *Why Added:* Further iteration of allowing statements before `super(...)` / `this(...)`. Java 23 had preview; Java 24 refines.
   * *How It Helps:* More reliable checks, better initialization patterns.

   **Example:**

   ```java
   class Parent {
       Parent(int x) { /* ... */ }
   }

   class Child extends Parent {
       int x;

       Child(int x) {
           if (x < 0) {
               throw new IllegalArgumentException("x must be non-negative");
           }
           this.x = x * 2;
           super(x);
       }
   }
   ```

3. **Module Import Declarations (Second Preview, JEP 494)**

   * *Why Added:* Continue simplifying modules & imports so modules exporting many packages are easier to consume.
   * *How It Helps:* Less boilerplate when using many classes from a module.

   **Example:**

   ```java
   import module java.util;

   public class Demo {
       public static void main(String[] args) {
           List<String> list = List.of("a", "b", "c");
           Set<Integer> set = Set.of(1,2,3);
           list.forEach(System.out::println);
       }
   }
   ```

4. **Simple Source Files and Instance Main Methods (Fourth Preview, JEP 495)**

   * *Why Added:* To further simplify writing small programs, making Java more approachable for scripting, rapid prototyping. 
   * *How It Helps:* You can write minimal-file Java code (no named class), and use instance `main` methods.

   **Example:**

   ```java
   void main() {
       println("Hello from simple source file!");
   }
   ```

5. **Scoped Values (Fourth Preview, JEP 487)**

   * *Why Added:* To improve how context (like user info, request tracing, etc.) is passed through code without heavy use of `ThreadLocal`. More control, safer.
   * *How It Helps:* Scoped values allow carrying contextual state through calls and threads with defined lifetimes. 

   **Example:**

   ```java
   ScopedValue<String> CONTEXT = ScopedValue.newInstance();

   void handleRequest() {
       ScopedValue.where(CONTEXT, "user123")
           .run(() -> process());
   }

   void process() {
       String user = CONTEXT.get();
       println("Handling for: " + user);
   }
   ```

6. **Vector API (Ninth Incubator, JEP 489)**

   * *Why Added:* Continue improving performance, especially for data-heavy, numeric, or vectorizable workloads. 
   * *How It Helps:* More operations, better performance, hardware utilization.

   **Example:**

   ```java
   float[] a = {1.0f, 2.0f, 3.0f, 4.0f};
   float[] b = {0.5f, 1.5f, 2.5f, 3.5f};
   float[] c = new float[a.length];

   var species = FloatVector.SPECIES_PREFERRED;
   for (int i = 0; i < a.length; i += species.length()) {
       var va = FloatVector.fromArray(species, a, i);
       var vb = FloatVector.fromArray(species, b, i);
       va.mul(vb).intoArray(c, i);
   }
   // c now has products element-wise
   ```

7. **Key Derivation Function API (Preview, JEP 478)**

   * *Why Added:* to standardize cryptographic key derivation functions (KDFs) (e.g. HKDF, PBKDF2) within JDK’s core. Enhances security and interoperability.
   * *How It Helps:* You no longer need third-party libs for basic KDF needs; using built-in APIs improves security posture.

   **Example:**

   ```java
   KeyDerivationFunction hkdf = KeyDerivationFunction.getInstance("HKDF-SHA256");
   HKDFParameters params = HKDFParameters.builder()
       .withSalt(saltBytes)
       .withInputKeyingMaterial(ikmBytes)
       .withInfo(infoBytes)
       .build();
   SecretKey derived = hkdf.deriveKey("AES", params, 32);
   ```

8. **Compact Object Headers (Experimental, JEP 450 / Lilliput)**

   * *Why Added:* To reduce the memory overhead of object headers on 64-bit platforms. Improves memory footprint.
   * *How It Helps:* Less memory consumed per object, potentially more objects fit in cache, better performance in memory-sensitive apps. 

   **Example:**
   (No code example needed – this is JVM / runtime behavior; usage is implicit once enabled.)

---


## Java 25 (Sep 16, 2025)

**Key Features:**

1. **Primitive Types in Patterns, `instanceof`, and `switch` (Third Preview, JEP 507)**

   * *Why Added:* To let pattern matching handle primitive types (like `int`, `double`, etc.) as well as reference types, reducing boilerplate and avoiding unnecessary boxing/unboxing.
   * *How It Helps:* Enables cleaner code when checking types or matching values; more expressive `switch` / `instanceof` usage.

   **Example:**

   ```java
   Object o = 5;

   if (o instanceof int i && i > 0) {
       System.out.println("Positive int: " + i);
   }

   switch (o) {
       case double d -> System.out.println("Double: " + d);
       case char c    -> System.out.println("Char: " + c);
       case int i     -> System.out.println("Int: " + i);
       default        -> System.out.println("Other type");
   }
   ```

2. **Module Import Declarations (Final, JEP 511)**

   * *Why Added:* To simplify using entire modules without having to import each package/class individually. Helps modular readability.
   * *How It Helps:* Less boilerplate in code that uses many classes from the same module; clearer intent.

   **Example:**

   ```java
   import module java.base;

   public class Demo {
       public static void main(String[] args) {
           // Now classes from java.base module (e.g., String, System) are available without explicit import
           String s = "Hello";
           System.out.println(s.toUpperCase());
       }
   }
   ```

3. **Compact Source Files and Instance Main Methods (Final, JEP 512)**

   * *Why Added:* To make simple programs easier to write (for learning, scripting, quick tasks) by reducing boilerplate elements like `public class`, `static`, etc.
   * *How It Helps:* Faster prototyping, less ceremony for small applications.

   **Example:**

   ```java
   void main() {
       println("Hello from compact Java 25!");
   }
   ```

4. **Flexible Constructor Bodies (Final, JEP 513)**

   * *Why Added:* Previously, you had to call `super(...)` or `this(...)` as the first statement in constructors. Now code (safe computations, validations, field initializations) can happen before that.
   * *How It Helps:* Enables earlier argument validation, avoids needing static helpers or duplicating code just to prepare arguments.

   **Example:**

   ```java
   class Parent {
       Parent(String name) { /* ... */ }
   }

   class Child extends Parent {
       String processed;

       Child(String name) {
           if (name == null || name.isBlank()) {
               throw new IllegalArgumentException("Name required");
           }
           processed = name.trim().toUpperCase();
           super(processed);
       }
   }
   ```

5. **Scoped Values (Final, JEP 506)**

   * *Why Added:* To provide a safer, more explicit alternative to `ThreadLocal` for passing immutable contextual data (like user info, request tracing) across threads/call stacks.
   * *How It Helps:* Defined scope: the value is bound only in certain code regions; less risk of memory leaks or global mutable state; better integration with virtual threads etc.

   **Example:**

   ```java
   import java.lang.ScopedValue;
   import java.util.concurrent.Executors;

   public class ContextExample {
       static final ScopedValue<String> USER = ScopedValue.newInstance();

       public static void main(String[] args) {
           var exec = Executors.newVirtualThreadPerTaskExecutor();
           ScopedValue.where(USER, "Alice", () -> {
               exec.submit(() -> {
                   System.out.println("Inside virtual thread: user=" + USER.get());
               });
           });
           exec.close();
       }
   }
   ```

6. **Compact Object Headers (JEP 519)**

   * *Why Added:* To reduce the size of object metadata (object headers) on 64-bit JVMs. Lower memory usage for small objects, better cache behaviour.
   * *How It Helps:* Improves overall memory efficiency, especially in applications with many small objects. 

   **Example:**
   *No code change needed.* Enable via JVM option:

   ```
   -XX:+UseCompactObjectHeaders
   ```

7. **JDK Flight Recorder (JFR) Enhancements: CPU-Time Profiling, Cooperative Sampling, Method Timing & Tracing (JEPs 509, 518, 520)**

   * *Why Added:* To improve observability and diagnostics: better precision, less overhead, more meaningful data.
   * *How It Helps:* Helps pinpoint performance bottlenecks, understand thread behaviour, track which methods use CPU, etc.

   **Example:**

   ```shell
   # Launch with JFR options (example)
   java -XX:StartFlightRecording=filename=recording.jfr,dumponexit=true \
        -XX:+UseCompressedOops MyApp
   # Later inspect with tools
   ```

8. **Generational Shenandoah Garbage Collector (JEP 521)**

   * *Why Added:* To bring generational GC (young/old generation divisions) to Shenandoah, improving throughput and reducing pause times under many workloads.
   * *How It Helps:* Better memory reclamation behaviour, more predictable performance in large‐heap environments.

   **Example:**
   *No code example needed.* Use appropriate JVM flags to enable and tune Shenandoah with generational mode.

9. **Key Derivation Function API (Final, JEP 510)**

   * *Why Added:* To standardize support for deriving cryptographic keys from shared secrets or inputs (e.g. HKDF, PBKDF2) within the JDK.
   * *How It Helps:* Reduces dependence on third-party libraries for common crypto tasks; safer and more consistent practices.

   **Example:**

   ```java
   import java.security.SecureRandom;
   import javax.crypto.SecretKey;

   byte[] salt = new byte[16];
   new SecureRandom().nextBytes(salt);
   byte[] ikm = "shared secret".getBytes();

   var hkdf = KeyDerivationFunction.getInstance("HKDF-SHA256");
   var params = HKDFParameters.builder()
       .withSalt(salt)
       .withInputKeyingMaterial(ikm)
       .withInfo("info".getBytes())
       .build();
   SecretKey key = hkdf.deriveKey("AES", params, 32);
   ```

10. **Stable Values (Preview, JEP 502)**

    * *Why Added:* To provide a mechanism for values that are immutable once set, but can be initialized lazily (i.e. deferred), including across threads. Offers more flexibility than `final` fields.
    * *How It Helps:* Helps startup performance (delay heavy initialization until needed), supports safe concurrency, optimisations by JVM.

    **Example:**

    ```java
    import java.util.concurrent.atomic.AtomicReference;

    // Pseudocode style; exact API evolving

    var stable = StableValue.newInstance();  // preview API
    stable.setIfAbsent(() -> computeHeavyResource());  // thread-safe deferred init
    var r = stable.get();  // returns the computed value
    ```

11. **Vector API (Incubator, JEP 508, Tenth Incubator Round)**

    * *Why Added:* To continue improving performance for data-parallel numeric tasks by allowing developers to use SIMD (Single Instruction Multiple Data) via Java vectors. 
    * *How It Helps:* Faster vectorized operations for compute-intensive workloads (machine learning, signal processing, image manipulation, etc.).

    **Example:**

    ```java
    import jdk.incubator.vector.FloatVector;
    import jdk.incubator.vector.VectorSpecies;

    public class VecDemo {
        static final VectorSpecies<Float> SPECIES = FloatVector.SPECIES_PREFERRED;

        static void doubleArray(float[] input, float[] output) {
            int i = 0;
            for (; i <= input.length - SPECIES.length(); i += SPECIES.length()) {
                var v = FloatVector.fromArray(SPECIES, input, i);
                var r = v.mul(2.0f);
                r.intoArray(output, i);
            }
            for (; i < input.length; i++) {
                output[i] = input[i] * 2.0f;
            }
        }
    }
    ```

12. **Removal of 32-bit x86 Port**

    * *Why Added:* The maintenance cost for 32-bit x86 (on various operating systems) no longer justifies the usage, especially since most hardware supports 64-bit.
    * *How It Helps:* Simplifies support, reduces maintenance burden, allows focusing optimizations and testing on 64-bit platforms. 

---
