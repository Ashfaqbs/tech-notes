# **📘 Understanding SDKs: What, Why, and How to Design One**  

This document explains what an SDK is, why we need it, how to design one, and the differences between an API and an SDK with straightforward examples.  

---

## **1️⃣ What is an SDK?**
An **SDK (Software Development Kit)** is a set of tools, libraries, documentation, and sample code that helps developers integrate with or extend a specific service, platform, or framework.  

Think of it as a **pre-packaged toolkit** that makes it easier to interact with an external system.

### **🛠 Example of an SDK**
- **Google Maps SDK** → Instead of making raw API calls, we can use Google’s SDK to easily embed maps in a mobile or web app.  
- **AWS SDK** → Helps developers interact with AWS services without manually handling authentication and HTTP requests.  

---
## **2️⃣ Why Do We Need an SDK?**
Without an SDK, developers would have to:
✅ Manually handle **HTTP requests and responses**  
✅ Manage **authentication, error handling, and retries**  
✅ Write **boilerplate code** to interact with external services  

### **🆚 API vs SDK**
| Feature  | API | SDK |
|----------|----|----|
| **Definition** | An interface for communication between two systems. | A toolkit that simplifies interaction with APIs. |
| **What It Provides** | Endpoints, request & response formats. | Prebuilt functions, authentication handling, API wrappers. |
| **Ease of Use** | Developers must manually handle requests. | Simplifies interaction by providing ready-to-use functions. |
| **Example** | REST API for Payment Processing. | Stripe SDK that provides `processPayment(amount, currency)` instead of raw HTTP calls. |

---
## **3️⃣ How to Design an SDK**
### **🔹 Step 1: Define the Purpose of the SDK**
Before writing code, ask:  
- What problem does this SDK solve?  
- Which APIs or services does it interact with?  
- What functionality should be exposed?

✅ **Example:**  
We are building an SDK for interacting with a **PostgREST API**, where users can:  
- Retrieve data from a table  
- Insert new records  
- Update existing records  
- Delete records  

---
### **🔹 Step 2: Design the SDK Structure**
An SDK usually consists of:
- **Core Client (`SdkClient`)** → Main class that provides methods like `getData()`, `postData()`, etc.  
- **Models (`ApiResponse.java`)** → Represent data structures returned by the API.  
- **Utilities (`HttpClient.java`)** → Handles authentication, request execution, and error handling.  

✅ **Example SDK Structure**:
```
simple-sdk/
 ├── src/main/java/com/ashfaq/sdk/
 │   ├── SdkClient.java   # Main entry point of SDK
 │   ├── ApiResponse.java # DTO for API responses
 │   ├── HttpClient.java  # Handles requests
 ├── pom.xml              # Project dependencies
 ├── README.md            # Documentation
```

---
### **🔹 Step 3: Implement Core SDK Functions**
Instead of requiring users to make direct API calls, wrap them inside meaningful functions.

#### **Example Without an SDK (Direct API Call)**
A user manually makes an API request:
```java
RestTemplate restTemplate = new RestTemplate();
ResponseEntity<String> response = restTemplate.getForEntity("http://localhost:3000/employee/2", String.class);
System.out.println(response.getBody());
```

#### **Example With an SDK**
The SDK provides a simple function like:
```java
SdkClient sdkClient = new SdkClient();
Employee employee = sdkClient.getEmployeeById(2);
```
✅ **Users don’t need to know how the API works**—they just call the function!

---
## **4️⃣ Flow of an SDK Call**
Let’s take an example where a user wants to fetch an employee record.

### **Flow:**
1️⃣ **User Calls SDK Method:**  
```java
SdkClient sdk = new SdkClient();
Employee employee = sdk.getEmployeeById(2);
```
⬇️  
2️⃣ **SDK Constructs the URL & Calls API:**  
`http://localhost:3000/employee?id=eq.2`  
⬇️  
3️⃣ **SDK Handles Response & Converts JSON to DTO:**  
```json
{
  "id": 2,
  "emp_code": 102,
  "name": "Jane Smith"
}
```
⬇️  
4️⃣ **SDK Returns Parsed Object to the User:**  
```java
System.out.println(employee.getName()); // Prints: Jane Smith
```
✅ The user **never deals with raw HTTP requests**—everything is abstracted!

---
## **5️⃣ Example: Full CRUD SDK Implementation**
### **📌 `SdkClient.java` (SDK Class)**
```java
package com.ashfaq.sdk;

import org.springframework.http.*;
import org.springframework.web.client.RestTemplate;

import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class SdkClient {
    private final String baseUrl;
    private final RestTemplate restTemplate;

    public SdkClient() {
        this.baseUrl = "http://localhost:3000";
        this.restTemplate = new RestTemplate();
    }

    public Employee getEmployeeById(int id) {
        String url = baseUrl + "/employee?id=eq." + id;
        ResponseEntity<Employee[]> response = restTemplate.getForEntity(url, Employee[].class);
        return (response.getBody() != null && response.getBody().length > 0) ? response.getBody()[0] : null;
    }

    public List<Employee> findByParams(Map<String, String> params) {
        String queryString = params.entrySet().stream()
            .map(entry -> entry.getKey() + "=eq." + entry.getValue())
            .collect(Collectors.joining("&"));

        String url = baseUrl + "/employee?" + queryString;
        ResponseEntity<List<Employee>> response = restTemplate.exchange(url, HttpMethod.GET, null, new ParameterizedTypeReference<List<Employee>>() {});
        return response.getBody();
    }

    public Employee createEmployee(Employee employee) {
        String url = baseUrl + "/employee";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<Employee> request = new HttpEntity<>(employee, headers);
        return restTemplate.postForObject(url, request, Employee.class);
    }
}
```

### **📌 `Employee.java` (DTO Model)**
```java
package com.ashfaq.sdk;

public class Employee {
    private int id;
    private int emp_code;
    private String name;

    // Getters & Setters
}
```

---
## **6️⃣ Final Thoughts**
### **When Should we Build an SDK?**
- **When multiple projects need to interact with the same API**.
- **When reducing complexity for developers** using our service.
- **When adding extra functionality** (caching, authentication, retries, etc.).

### **When NOT to Build an SDK?**
- If our API is simple and doesn’t require abstraction.
- If using an existing SDK (e.g., AWS SDK) is sufficient.

---
## **🎯 Summary**
✅ **SDK makes interacting with APIs easier** by abstracting complexity.  
✅ **Key differences**: APIs provide endpoints, SDKs provide ready-to-use functions.  
✅ **Design involves creating wrappers around API calls** and exposing simple methods.  
✅ **Flow:** User calls SDK → SDK calls API → SDK returns a simplified response.  

---
## **Next Steps**
🚀 **Extend SDK**: Add `updateEmployee()` and `deleteEmployee()`.  
📌 **Publish SDK**: Deploy it to a Maven repository for broader use.  
🔐 **Security**: Add authentication (e.g., API Keys, OAuth).  
