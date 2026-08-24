# **📘 Creating a Simple SDK in Java Using Spring Boot and Maven**
This guide documents how we created a simple SDK that wraps `RestTemplate` and provided it as a dependency to a client project.

---

## **1️⃣ Creating the SDK Project**
We first created a separate Maven project for the SDK.

### **🔹 Step 1: Generate a New Maven Project**
Run the following command to create a new Maven project:

```sh
mvn archetype:generate -DgroupId=com.ashfaq.sdk -DartifactId=simple-sdk -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
```

- `groupId=com.ashfaq.sdk` → Group identifier for the SDK package.
- `artifactId=simple-sdk` → Name of the SDK.
- `archetypeArtifactId=maven-archetype-quickstart` → Generates a basic Maven project.

### **🔹 Step 2: Modify `pom.xml` for SDK**
We updated `simple-sdk/pom.xml` to add Spring Web support.

#### **Modifications in `pom.xml`:**
```xml
<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.ashfaq.sdk</groupId>
  <artifactId>simple-sdk</artifactId>
  <version>1.0.0</version>
  <packaging>jar</packaging>

  <properties>
    <maven.compiler.source>21</maven.compiler.source>
    <maven.compiler.target>21</maven.compiler.target>
  </properties>

  <dependencies>
    <dependency>
      <groupId>org.springframework</groupId>
      <artifactId>spring-web</artifactId>
      <version>6.0.0</version>
    </dependency>
  </dependencies>
</project>
```
- Added `spring-web` dependency for `RestTemplate`.
- Set `packaging` to `jar` since SDKs are libraries.

---

## **2️⃣ Implementing SDK Code**
We created a class called `SdkClient.java` that wraps `RestTemplate`.

### **🔹 Step 3: Create `SdkClient.java`**
**File Location:** `src/main/java/com/ashfaq/sdk/SdkClient.java`

```java
package com.ashfaq.sdk;

import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestTemplate;

public class SdkClient {
    private final RestTemplate restTemplate;

    public SdkClient() {
        this.restTemplate = new RestTemplate();
    }

    public ApiResponse getApiResponse(String url) {
        ResponseEntity<ApiResponse> response = restTemplate.getForEntity(url, ApiResponse.class);
        return response.getBody();
    }
}
```

- **`SdkClient` class**: Acts as the main entry point for our SDK.
- **`getApiResponse(String url)`**: Calls the given API URL using `RestTemplate`.

---

## **3️⃣ Building and Installing the SDK**
Once the SDK code was complete, we needed to package it as a dependency.

### **🔹 Step 4: Install the SDK in the Local Repository**
Inside the `simple-sdk` project directory, run:

```sh
mvn clean install
```

- This builds the project and installs it in the local Maven repository (`~/.m2/repository`).
- Now, the SDK can be used in other projects as a dependency.

---

## **4️⃣ Integrating SDK into a Client Project**
Now, we created a separate **Spring Boot Client Project** that will use the SDK.

### **🔹 Step 5: Create a New Client Project**
Run this command to create the client project:

```sh
mvn archetype:generate -DgroupId=com.ashfaq.client -DartifactId=sdk-client -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
```

### **🔹 Step 6: Add SDK Dependency in `sdk-client/pom.xml`**
```xml
<dependency>
    <groupId>com.ashfaq.sdk</groupId>
    <artifactId>simple-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

- This tells Maven to use the locally installed SDK.
- If the SDK were published to a remote repository, it would be pulled from there.

---

## **5️⃣ Using the SDK in Client Code**
### **🔹 Step 7: Create a Client Class (`App.java`)**
**File Location:** `src/main/java/com/ashfaq/client/App.java`

```java
package com.ashfaq.client;

import com.ashfaq.sdk.SdkClient;

public class App {
    public static void main(String[] args) {
        SdkClient sdkClient = new SdkClient();
        String response = sdkClient.getApiResponse("https://jsonplaceholder.typicode.com/todos/1");
        System.out.println("Response: " + response);
    }
}
```
- **Creates an instance of `SdkClient`**.
- **Calls `getApiResponse(url)`** to fetch data from an external API.

---

## **6️⃣ Running the Client**
After setting up the client, we ran the project:

```sh
mvn clean install
mvn exec:java -Dexec.mainClass="com.ashfaq.client.App"
```

### **✅ Expected Output:**
If successful, it should print the response from the API.

Example:
```json
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

---

## **🎯 Summary of Steps**
1️⃣ **Created SDK Project**
   - Used `mvn archetype:generate` to create `simple-sdk`.
   - Modified `pom.xml` to include `spring-web`.

2️⃣ **Implemented `SdkClient.java`**
   - Wrapped `RestTemplate` for API calls.

3️⃣ **Installed SDK Locally**
   - Ran `mvn clean install` to install SDK.

4️⃣ **Created Client Project**
   - Used `mvn archetype:generate` to create `sdk-client`.
   - Added SDK as a dependency in `pom.xml`.

5️⃣ **Used SDK in Client**
   - Created `App.java` to call SDK methods.

6️⃣ **Ran and Verified the Client**
   - Used `mvn exec:java` to execute the client.

---

## **Next Steps**
- **Publish SDK to a Remote Repository** (e.g., Maven Central, Nexus, JFrog) so others can use it.
- **Add More Methods to SDK** (e.g., `POST`, `PUT`, `DELETE` methods).
- **Improve Error Handling** in SDK (log errors, retry failed requests, etc.).

---
