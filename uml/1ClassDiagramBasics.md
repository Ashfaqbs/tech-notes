# Class Diagram Basics

**1. What is a Class Diagram?**  
A **Class Diagram** is a **static (unchanging)** picture that shows:
- **Classes** (like blueprints of objects)
- **Attributes** (data each object holds)
- **Methods** (actions each object can do)
- **Relationships** (how classes are connected)

In simple words:  
It’s like a **map of our system’s building blocks** — what things exist and how they are linked.

---

**2. Why Class Diagrams Matter:**  
- **Before coding:** We can *plan* our entities (e.g., User, Product, Order).  
- **While coding:** We can *refer* to ensure the structure is consistent.  
- **Teamwork:** Everyone knows which class does what — less confusion.
- **Scalability:** Easy to extend if we see clean relationships.

---

**3. Key Parts of a Class Diagram:**  
| Element         | Meaning                                    |
|:----------------|:-------------------------------------------|
| Rectangle       | A class (box divided into 3 parts)         |
| Top part        | Class name (bold)                         |
| Middle part     | Attributes (variables/data members)        |
| Bottom part     | Methods (functions/behaviors)              |
| Lines           | Connections (associations between classes) |
| Arrows          | Directions like inheritance or dependencies |

---

**4. Minimal Example:**

**Scenario:**  
We are designing a basic **Customer-Order** system.

Classes:
- **Customer**
  - Attributes: name, email
  - Methods: register(), login()

- **Order**
  - Attributes: orderId, orderDate
  - Methods: createOrder(), cancelOrder()

**Relationship:**  
- A Customer **places** an Order.
- (One-to-many: One Customer -> many Orders)

**Rough Visualization in Mind:**  
```
+------------------+
| Customer         |
+------------------+
| - name           |
| - email          |
+------------------+
| + register()     |
| + login()        |
+------------------+

   (connects to)

+------------------+
| Order            |
+------------------+
| - orderId        |
| - orderDate      |
+------------------+
| + createOrder()  |
| + cancelOrder()  |
+------------------+
```
(There will be a line between them, with a "1" near Customer and "*" near Order.)

---