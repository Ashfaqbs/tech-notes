### 📦 Project Overview:

CRUD for managing `VisaUser` entities
Operations: Create, Read, Update, Delete

---

### ✅ Layers You Asked For:

| Layer                  | Responsibility                    |
| ---------------------- | --------------------------------- |
| Entity                 | Maps to DB table (JPA)            |
| DTO (Request/Response) | Shape incoming and outgoing JSON  |
| Repository             | DAO — interacts with the database |
| IService               | Interface defining business logic |
| ServiceImpl            | Implements business logic         |
| Controller             | Handles HTTP requests (API layer) |

---

### ⚙ Step 1: Define the `VisaUser` Entity

We’ll start with:

* `id` (UUID)
* `fullName`
* `email`
* `nationality`
* `visaStatus` (enum)
* `createdAt` (timestamp)

## Step 1: Define the Core Model

### Java Entity — `VisaUser.java`

```java
package com.example.visa.entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "visa_users")
public class VisaUser {

    @Id
    @GeneratedValue
    private UUID id;

    private String fullName;
    private String email;
    private String nationality;

    @Enumerated(EnumType.STRING)
    private VisaStatus visaStatus;

    private LocalDateTime createdAt;

    // Getters & Setters
}
```

---

###  Enum — `VisaStatus.java`

```java
package com.example.visa.entity;

public enum VisaStatus {
    PENDING,
    APPROVED,
    REJECTED
}
```

---

## Step 2: Define DTOs

###  Request DTO — `VisaUserRequestDto.java`

```java
package com.example.visa.dto;

public class VisaUserRequestDto {
    private String fullName;
    private String email;
    private String nationality;
    private String visaStatus; // keep as String for API input
}
```

### Response DTO — `VisaUserResponseDto.java`

```java
package com.example.visa.dto;

import java.time.LocalDateTime;
import java.util.UUID;

public class VisaUserResponseDto {
    private UUID id;
    private String fullName;
    private String email;
    private String nationality;
    private String visaStatus;
    private LocalDateTime createdAt;
}
```

---

##  UML Class Diagram (Text Format)

```
+------------------------+
|      VisaUser          |
+------------------------+
| - id: UUID             |
| - fullName: String     |
| - email: String        |
| - nationality: String  |
| - visaStatus: Enum     |
| - createdAt: DateTime  |
+------------------------+

+-----------------------------+       +------------------------------+
| VisaUserRequestDto          |       | VisaUserResponseDto          |
+-----------------------------+       +------------------------------+
| - fullName: String          |       | - id: UUID                   |
| - email: String             |       | - fullName: String           |
| - nationality: String       |       | - email: String              |
| - visaStatus: String        |       | - nationality: String        |
+-----------------------------+       | - visaStatus: String         |
                                      | - createdAt: LocalDateTime   |
                                      +------------------------------+

+------------------+
| VisaStatus Enum  |
+------------------+
| PENDING          |
| APPROVED         |
| REJECTED         |
+------------------+
```

##  Step 3: Repository Layer

### VisaUserRepository.java

```java
package com.example.visa.repository;

import com.example.visa.entity.VisaUser;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.UUID;

public interface VisaUserRepository extends JpaRepository<VisaUser, UUID> {
}
```

* This gives basic CRUD: `findById`, `save`, `deleteById`, `findAll`, etc.

---

##  Step 4: Service Interface

###  IVisaUserService.java

```java
package com.example.visa.service;

import com.example.visa.dto.VisaUserRequestDto;
import com.example.visa.dto.VisaUserResponseDto;
import java.util.List;
import java.util.UUID;

public interface IVisaUserService {
    VisaUserResponseDto createVisaUser(VisaUserRequestDto dto);
    VisaUserResponseDto getVisaUserById(UUID id);
    List<VisaUserResponseDto> getAllVisaUsers();
    VisaUserResponseDto updateVisaUser(UUID id, VisaUserRequestDto dto);
    void deleteVisaUser(UUID id);
}
```

---

##  Step 5: Service Implementation

###  VisaUserServiceImpl.java

```java
package com.example.visa.service.impl;

import com.example.visa.dto.VisaUserRequestDto;
import com.example.visa.dto.VisaUserResponseDto;
import com.example.visa.entity.VisaStatus;
import com.example.visa.entity.VisaUser;
import com.example.visa.repository.VisaUserRepository;
import com.example.visa.service.IVisaUserService;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.*;
import java.util.stream.Collectors;

@Service
public class VisaUserServiceImpl implements IVisaUserService {

    private final VisaUserRepository repo;

    public VisaUserServiceImpl(VisaUserRepository repo) {
        this.repo = repo;
    }

    @Override
    public VisaUserResponseDto createVisaUser(VisaUserRequestDto dto) {
        VisaUser user = new VisaUser();
        user.setFullName(dto.getFullName());
        user.setEmail(dto.getEmail());
        user.setNationality(dto.getNationality());
        user.setVisaStatus(VisaStatus.valueOf(dto.getVisaStatus().toUpperCase()));
        user.setCreatedAt(LocalDateTime.now());

        return toResponse(repo.save(user));
    }

    @Override
    public VisaUserResponseDto getVisaUserById(UUID id) {
        return toResponse(repo.findById(id)
            .orElseThrow(() -> new RuntimeException("User not found")));
    }

    @Override
    public List<VisaUserResponseDto> getAllVisaUsers() {
        return repo.findAll()
            .stream()
            .map(this::toResponse)
            .collect(Collectors.toList());
    }

    @Override
    public VisaUserResponseDto updateVisaUser(UUID id, VisaUserRequestDto dto) {
        VisaUser user = repo.findById(id)
            .orElseThrow(() -> new RuntimeException("User not found"));

        user.setFullName(dto.getFullName());
        user.setEmail(dto.getEmail());
        user.setNationality(dto.getNationality());
        user.setVisaStatus(VisaStatus.valueOf(dto.getVisaStatus().toUpperCase()));

        return toResponse(repo.save(user));
    }

    @Override
    public void deleteVisaUser(UUID id) {
        repo.deleteById(id);
    }

    private VisaUserResponseDto toResponse(VisaUser user) {
        VisaUserResponseDto dto = new VisaUserResponseDto();
        dto.setId(user.getId());
        dto.setFullName(user.getFullName());
        dto.setEmail(user.getEmail());
        dto.setNationality(user.getNationality());
        dto.setVisaStatus(user.getVisaStatus().name());
        dto.setCreatedAt(user.getCreatedAt());
        return dto;
    }
}
```

---

##  Step 6: Controller Layer

###  VisaUserController.java

```java
package com.example.visa.controller;

import com.example.visa.dto.VisaUserRequestDto;
import com.example.visa.dto.VisaUserResponseDto;
import com.example.visa.service.IVisaUserService;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/visa-users")
public class VisaUserController {

    private final IVisaUserService service;

    public VisaUserController(IVisaUserService service) {
        this.service = service;
    }

    @PostMapping
    public VisaUserResponseDto create(@RequestBody VisaUserRequestDto dto) {
        return service.createVisaUser(dto);
    }

    @GetMapping("/{id}")
    public VisaUserResponseDto getById(@PathVariable UUID id) {
        return service.getVisaUserById(id);
    }

    @GetMapping
    public List<VisaUserResponseDto> getAll() {
        return service.getAllVisaUsers();
    }

    @PutMapping("/{id}")
    public VisaUserResponseDto update(@PathVariable UUID id, @RequestBody VisaUserRequestDto dto) {
        return service.updateVisaUser(id, dto);
    }

    @DeleteMapping("/{id}")
    public void delete(@PathVariable UUID id) {
        service.deleteVisaUser(id);
    }
}
```

---

##  UML Layer View (Simplified Class Diagram)

```
Controller
    |
    v
IVisaUserService <-------- VisaUserServiceImpl
    |
    v
VisaUserRepository <------ JPA DB
    ↑
    |
VisaUser (Entity)
    ↑
    |
DTOs (Request <-> Response)
```