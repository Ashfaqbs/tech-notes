# Step 3 — Official Class Diagram Notations

**1. How to Draw a Class in UML:**  
A Class is shown as a **rectangle divided into 3 sections**:

| Part            | Purpose                          |
|:----------------|:----------------------------------|
| Top Section     | **Class Name** (bold, centered)   |
| Middle Section  | **Attributes** (properties/fields) |
| Bottom Section  | **Methods** (functions/behaviors)  |

**Example:** (Text version)  
```
+--------------------+
| Customer            |   ← Top: Class Name
+--------------------+
| - name: String      |   ← Middle: Attributes
| - email: String     |
+--------------------+
| + register(): void  |   ← Bottom: Methods
| + login(): boolean  |
+--------------------+
```

---

**2. Symbols Used in Attributes and Methods:**  
| Symbol | Meaning                     |
|:-------|:-----------------------------|
| `+`    | Public (accessible by anyone) |
| `-`    | Private (only inside the class) |
| `#`    | Protected (inside class and subclasses) |
| `~`    | Package-private (default in Java) |

**Examples:**  
- `- name: String` → private variable called `name` of type `String`.
- `+ login(): boolean` → public method `login` that returns `boolean`.

---

**3. Relationships (Associations) Between Classes:**  

| Relationship         | UML Symbol                     | Meaning                        |
|:---------------------|:-------------------------------|:-------------------------------|
| Association           | Straight Line                 | "Has a" relationship           |
| Inheritance (Generalization) | Line with hollow triangle | "Is a" relationship (parent-child) |
| Dependency           | Dotted Line with arrow         | "Uses temporarily" relationship |
| Aggregation          | Line with hollow diamond       | "Whole-Part" (but part can exist separately) |
| Composition          | Line with filled black diamond | "Whole-Part" (part dies with whole) |

---

**4. Let's Redo the Customer → Order Example (Official UML way):**

Drawing Structure (Imagine this cleanly):

```
+--------------------+                    +--------------------+
| Customer            | 1               * | Order               |
+--------------------+ ------------------ +--------------------+
| - name: String      |                   | - orderId: int       |
| - email: String     |                   | - orderDate: Date    |
+--------------------+                   +--------------------+
| + register(): void  |                   | + createOrder(): void |
| + login(): boolean  |                   | + cancelOrder(): void|
+--------------------+                   +--------------------+
```
**Line Between Them:**  
- A straight line connects Customer to Order.  
- Near Customer side: put `1` (meaning one customer)  
- Near Order side: put `*` (meaning many orders)

This means:  
→ **One Customer can have Many Orders**.

---

# Quick Memory Tip  
**Classes = Boxes**  
**Attributes = Variables inside box**  
**Methods = Functions inside box**  
**Lines = How boxes talk to each other**  
**Arrows/Triangles/Diamonds = Type of relationship**

---




Let’s break down each UML relationship type **step by step** with:

- A **simple real-world analogy**
- A **Java code example**
- A **UML diagram** (as text-based, but I can give a visual one if you want)

---

## ✅ 1. **Association** → `"Has-a"` relationship  
- **Meaning:** One class uses another class by *having it as a field*  
- **Real-world analogy:** A `Customer` **has an** `Address`

### Java Example:
```java
class Address {
    String city;
    String state;
}

class Customer {
    String name;
    Address address; // Association: Customer has an Address
}
```

### UML (Text Style):
```
+------------+       +-----------+
| Customer   | 1     1 | Address   |
+------------+-------+-----------+
| - name     |       | - city     |
| - address  |       | - state    |
+------------+       +-----------+
```

➡ Association is a **straight line** between classes.

---

## ✅ 2. **Inheritance (Generalization)** → `"Is-a"` relationship  
- **Meaning:** One class *inherits* another’s behavior  
- **Real-world analogy:** A `Dog` **is a** `Animal`

### Java Example:
```java
class Animal {
    void makeSound() {}
}

class Dog extends Animal {
    void bark() {}
}
```

### UML:
```
+--------+
| Animal |
+--------+
| +makeSound() |
+--------+
    ▲
    |
+--------+
| Dog    |
+--------+
| +bark()|
+--------+
```

➡ Inheritance uses a **solid line with a hollow triangle** pointing to the parent.

---

## ✅ 3. **Dependency** → `"Uses temporarily"`  
- **Meaning:** One class **depends** on another, but only **within a method**  
- **Real-world analogy:** A `ShoppingCart` **uses** `DiscountService` to apply discounts temporarily

### Java Example:
```java
class DiscountService {
    double getDiscount() { return 10.0; }
}

class ShoppingCart {
    void applyDiscount(DiscountService ds) {
        double d = ds.getDiscount();
    }
}
```

### UML:
```
+----------------+    . . . . . . . . .>
| ShoppingCart   | ------------------> | DiscountService |
+----------------+                     +-----------------+
```

➡ Dependency uses a **dotted arrow**. Think: “used *just for now*.”

---

## ✅ 4. **Aggregation** → `"Whole-Part"` (part can live independently)  
- **Meaning:** A class has another class as a field, but the part can **exist separately**  
- **Real-world analogy:** A `Team` has `Players` — even if the team is deleted, players still exist.

### Java Example:
```java
class Player {
    String name;
}

class Team {
    List<Player> players = new ArrayList<>(); // Aggregation
}
```

### UML:
```
+--------+ ◇------> +--------+
| Team   |          | Player |
+--------+          +--------+
```

➡ **Hollow diamond (◇)** at the “whole” side → **Aggregation**

---

## ✅ 5. **Composition** → `"Strong Whole-Part"` (part dies with whole)  
- **Meaning:** A class owns another class fully — the part **cannot exist alone**  
- **Real-world analogy:** A `House` contains `Rooms`. If the house is destroyed, rooms go too.

### Java Example:
```java
class Room {
    String type;
}

class House {
    private List<Room> rooms = new ArrayList<>();

    House() {
        rooms.add(new Room()); // created and owned here
    }
}
```

### UML:
```
+--------+ ◆------> +--------+
| House  |          | Room   |
+--------+          +--------+
```

➡ **Filled diamond (◆)** at the “whole” side → **Composition**

---

# 🔁 Quick Summary Table

| Type          | Lifetime Link           | UML Symbol        | Example             |
|---------------|-------------------------|--------------------|---------------------|
| Association   | Loose / General         | Straight Line      | Customer → Address  |
| Inheritance   | Permanent (parent-child) | ▲ Triangle Arrow   | Dog → Animal        |
| Dependency    | Temporary               | Dotted Arrow       | ShoppingCart → DiscountService |
| Aggregation   | Whole–Part (loose)      | ◇ Hollow Diamond   | Team → Players      |
| Composition   | Whole–Part (strong)     | ◆ Filled Diamond   | House → Rooms       |

---