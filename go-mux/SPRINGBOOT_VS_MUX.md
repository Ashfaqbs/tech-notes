# Spring Boot vs Go + Mux - Feature Comparison

If you know Spring Boot, this guide shows how to do the same things in Go + Mux.

## Quick Comparison Table

| Feature | Spring Boot | Go + Mux |
|---------|-------------|----------|
| **Web Framework** | Spring MVC (`@RestController`) | Gorilla Mux (`mux.Router`) |
| **Dependency Injection** | `@Autowired`, Spring Container | Manual, constructor injection |
| **Configuration** | `application.yml`, `@ConfigurationProperties` | Environment variables, config structs |
| **Database (JPA)** | Spring Data JPA, `@Entity`, `@Repository` | `database/sql`, GORM, sqlx |
| **Database Migrations** | Flyway, Liquibase | golang-migrate, Goose |
| **Validation** | `@Valid`, Jakarta Validation | `go-playground/validator` |
| **Cron Jobs** | `@Scheduled`, `@EnableScheduling` | `robfig/cron` |
| **Caching** | `@Cacheable`, Spring Cache | Redis, `go-redis/redis` |
| **Security/Auth** | Spring Security | Custom middleware, `golang-jwt/jwt` |
| **Testing** | JUnit, `@SpringBootTest`, MockMvc | `testing`, `testify`, `httptest` |
| **Logging** | SLF4J, Logback | `log`, `zap`, `logrus` |
| **API Docs** | Springdoc OpenAPI (`@Swagger`) | `swaggo/swag` |
| **Graceful Shutdown** | `server.shutdown=graceful` | Signal handling with `os/signal` |

---

## 1. REST API Endpoints

### Spring Boot
```java
@RestController
@RequestMapping("/api/users")
public class UserController {
    
    @Autowired
    private UserService userService;
    
    @GetMapping("/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        User user = userService.findById(id);
        return ResponseEntity.ok(user);
    }
    
    @PostMapping
    public ResponseEntity<User> createUser(@RequestBody @Valid User user) {
        User created = userService.save(user);
        return ResponseEntity.status(201).body(created);
    }
}
```

### Go + Mux
```go
type UserController struct {
    service *UserService  // Manual dependency injection
}

func NewUserController(service *UserService) *UserController {
    return &UserController{service: service}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    user, err := uc.service.FindByID(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    JSONResponse(w, user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    created, err := uc.service.Save(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    JSONResponse(w, created)
}

// Register routes
uc := NewUserController(NewUserService())
r.HandleFunc("/api/users/{id}", uc.GetUser).Methods("GET")
r.HandleFunc("/api/users", uc.CreateUser).Methods("POST")
```

**Key Differences:**
- No annotations in Go - routes registered explicitly
- Manual dependency injection (no `@Autowired`)
- More verbose but clearer what's happening

---

## 2. Configuration (application.yml equivalent)

### Spring Boot
```yaml
# application.yml
server:
  port: 8000
spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/mydb
    username: user
    password: password
  jpa:
    hibernate:
      ddl-auto: validate
app:
  jwt-secret: my-secret-key
  cache-ttl: 3600
```

```java
@Configuration
@ConfigurationProperties(prefix = "app")
public class AppConfig {
    private String jwtSecret;
    private int cacheTtl;
    // getters/setters
}
```

### Go + Mux
```go
// config.go
type Config struct {
    Port       string
    DatabaseURL string `env:"DATABASE_URL"`
    JWTSecret  string `env:"JWT_SECRET"`
    CacheTTL   int    `env:"CACHE_TTL"`
}

func LoadConfig() *Config {
    _ = godotenv.Load()  // Load .env file
    
    return &Config{
        Port:        getEnv("PORT", "8000"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost:5432/mydb"),
        JWTSecret:   getEnv("JWT_SECRET", "change-me"),
        CacheTTL:    getEnvInt("CACHE_TTL", 3600),
    }
}

// Usage
config := LoadConfig()
fmt.Println(config.DatabaseURL)
```

**.env file:**
```bash
PORT=8000
DATABASE_URL=postgres://user:password@localhost:5432/mydb
JWT_SECRET=my-secret-key
CACHE_TTL=3600
```

**Key Differences:**
- No YAML in Go (use `.env` or JSON/ TOML files)
- No `@ConfigurationProperties` - use struct + env vars
- Libraries: `joho/godotenv`, `caarlos0/env`

---

## 3. Database Connection (JPA equivalent)

### Spring Boot
```java
@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    private String name;
    private String email;
    // getters/setters
}

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByEmail(String email);
}
```

### Go + Mux with GORM (JPA-like ORM)
```go
// model.go
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:255;not null"`
    Email string `gorm:"size:255;uniqueIndex;not null"`
}

// repository.go
type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (*User, error) {
    var user User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
    var user User
    err := r.db.Where("email = ?", email).First(&user).Error
    return &user, err
}

func (r *UserRepository) Save(user *User) error {
    return r.db.Save(user).Error
}

// main.go - setup
db, _ := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
db.AutoMigrate(&User{})  // Like ddl-auto: update
userRepo := NewUserRepository(db)
```

**Key Differences:**
- GORM is similar to JPA but lighter
- No interfaces like Spring Data - write methods explicitly
- More control, less magic

**Alternative: Pure SQL (like JdbcTemplate)**
```go
import "database/sql"

db, _ := sql.Open("postgres", config.DatabaseURL)

func (r *UserRepository) FindByID(id int) (*User, error) {
    row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    return &user, err
}
```

---

## 4. Database Migrations (Flyway/Liquibase)

### Spring Boot (Flyway)
```sql
-- V1__create_users.sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);
```

Auto-runs on startup with `spring.flyway.enabled=true`.

### Go + Mux (golang-migrate)
```bash
# Install CLI
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration
migrate create -ext sql -dir migrations -seq create_users_table
```

```sql
-- migrations/000001_create_users_table.up.sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

-- migrations/000001_create_users_table.down.sql
DROP TABLE IF EXISTS users;
```

```go
// main.go - run migrations on startup
import "github.com/golang-migrate/migrate/v4"
_ "github.com/golang-migrate/migrate/v4/database/postgres"
_ "github.com/golang-migrate/migrate/v4/source/file"

m, _ := migrate.New("file://migrations", config.DatabaseURL)
m.Up()  // Run all pending migrations
```

**Key Differences:**
- Same concept (versioned SQL files)
- Go requires explicit call to run migrations
- CLI tool for creating migrations

---

## 5. Validation (Jakarta Validation equivalent)

### Spring Boot
```java
public class CreateUserRequest {
    @NotBlank
    private String name;
    
    @Email
    @NotBlank
    private String email;
    
    @Min(18)
    private int age;
}

@PostMapping
public ResponseEntity<?> create(@RequestBody @Valid CreateUserRequest req) {
    // Validated automatically
}
```

### Go + Mux (go-playground/validator)
```go
import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=3,max=50"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,min=18"`
}

var validate = validator.New()

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    if err := validate.Struct(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Valid - proceed
}
```

**Key Differences:**
- Similar validation tags (`validate:` vs `@Valid`)
- Go requires manual call to validator
- Same validation rules available (required, email, min, max, etc.)

---

## 6. Cron Jobs (@Scheduled equivalent)

### Spring Boot
```java
@Configuration
@EnableScheduling
public class ScheduledTasks {
    
    @Scheduled(fixedRate = 60000)  // Every 60 seconds
    public void cleanupOldUsers() {
        log.info("Cleaning up old users...");
    }
    
    @Scheduled(cron = "0 0 2 * * *")  // Every day at 2 AM
    public void dailyReport() {
        log.info("Generating daily report...");
    }
}
```

### Go + Mux (robfig/cron)
```go
import "github.com/robfig/cron/v3"

func main() {
    c := cron.New()
    
    // Every 60 seconds
    c.AddFunc("@every 60s", func() {
        log.Println("Cleaning up old users...")
        cleanupOldUsers()
    })
    
    // Every day at 2 AM
    c.AddFunc("0 2 * * *", func() {
        log.Println("Generating daily report...")
        dailyReport()
    })
    
    c.Start()
    
    // ... rest of server setup
}

func cleanupOldUsers() {
    // Your logic here
}

func dailyReport() {
    // Your logic here
}
```

**Cron Expression Comparison:**

| Spring Boot | Go (robfig/cron) | Description |
|-------------|------------------|-------------|
| `@Scheduled(fixedRate = 60000)` | `@every 60s` | Every 60 seconds |
| `cron = "0 0 2 * * *"` | `0 2 * * *` | Daily at 2 AM |
| `cron = "*/5 * * * *"` | `*/5 * * * *` | Every 5 minutes |

**Key Differences:**
- Similar cron syntax
- Go requires explicit scheduler setup
- No `@EnableScheduling` - just create and start

---

## 7. Caching (@Cacheable equivalent)

### Spring Boot
```java
@Service
public class UserService {
    
    @Cacheable(value = "users", key = "#id")
    public User findById(Long id) {
        return userRepository.findById(id).orElse(null);
    }
    
    @CacheEvict(value = "users", key = "#user.id")
    public void update(User user) {
        userRepository.save(user);
    }
}
```

### Go + Mux (Redis with go-redis)
```go
import "github.com/go-redis/redis/v8"

type UserService struct {
    db    *gorm.DB
    redis *redis.Client
}

func (s *UserService) FindByID(ctx context.Context, id uint) (*User, error) {
    // Try cache first
    key := fmt.Sprintf("user:%d", id)
    cached, err := s.redis.Get(ctx, key).Result()
    
    if err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // Cache miss - query database
    var user User
    err = s.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    data, _ := json.Marshal(user)
    s.redis.Set(ctx, key, data, 1*time.Hour)
    
    return &user, nil
}

func (s *UserService) Update(ctx context.Context, user *User) error {
    err := s.db.Save(user).Error
    if err != nil {
        return err
    }
    
    // Evict cache
    key := fmt.Sprintf("user:%d", user.ID)
    s.redis.Del(ctx, key)
    
    return nil
}
```

**Key Differences:**
- Spring has annotation-based caching (magic)
- Go requires explicit cache logic (clear but more code)
- Same Redis backend works for both

---

## 8. Security/Authentication

### Spring Boot (Spring Security)
```java
@Configuration
@EnableWebSecurity
public class SecurityConfig {
    
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/api/public/**").permitAll()
                .requestMatchers("/api/admin/**").hasRole("ADMIN")
                .anyRequest().authenticated()
            )
            .httpBasic(Customizer.withDefaults())
            .sessionManagement(sess -> sess.sessionCreationPolicy(SessionCreationPolicy.STATELESS));
        return http.build();
    }
}
```

### Go + Mux (Custom Middleware)
```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        
        // Parse JWT
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Add claims to context
        claims := token.Claims.(jwt.MapClaims)
        ctx := context.WithValue(r.Context(), "userID", claims["user_id"])
        ctx = context.WithValue(ctx, "role", claims["role"])
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Role-based authorization
func RequireRole(role string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userRole := r.Context().Value("role").(string)
            if userRole != role {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}

// Apply to routes
r.HandleFunc("/api/public/health", HealthHandler)  // No auth
r.Use(AuthMiddleware)                              // All below require auth
r.HandleFunc("/api/admin/users", AdminHandler).Methods(RequireRole("admin"))
```

**Key Differences:**
- Spring Security = powerful but complex configuration
- Go middleware = simple, explicit, composable
- Same concepts (authentication, authorization, roles)

---

## 9. Testing

### Spring Boot
```java
@SpringBootTest
@AutoConfigureMockMvc
class UserControllerTest {
    
    @Autowired
    private MockMvc mockMvc;
    
    @MockBean
    private UserService userService;
    
    @Test
    void shouldReturnUser() throws Exception {
        when(userService.findById(1L)).thenReturn(new User(1L, "John"));
        
        mockMvc.perform(get("/api/users/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.name").value("John"));
    }
}
```

### Go + Mux
```go
import (
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
    // Mock service
    mockService := &MockUserService{}
    mockService.On("FindByID", uint(1)).Return(&User{ID: 1, Name: "John"}, nil)
    
    // Create handler
    uc := NewUserController(mockService)
    
    // Create test request
    req := httptest.NewRequest("GET", "/api/users/1", nil)
    w := httptest.NewRecorder()
    
    // Call handler
    uc.GetUser(w, req)
    
    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "John")
}
```

**Key Differences:**
- Similar testing concepts (mocks, assertions)
- Go uses `httptest` instead of MockMvc
- `testify` library for assertions (like AssertJ)

---

## 10. Logging

### Spring Boot
```java
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@Service
public class UserService {
    private static final Logger log = LoggerFactory.getLogger(UserService.class);
    
    public User findById(Long id) {
        log.info("Finding user by id: {}", id);
        log.debug("Executing query...");
        return userRepository.findById(id).orElse(null);
    }
}
```

### Go + Mux (zap logger)
```go
import "go.uber.org/zap"

type UserService struct {
    db   *gorm.DB
    log  *zap.Logger
}

func NewUserService(log *zap.Logger) *UserService {
    return &UserService{log: log}
}

func (s *UserService) FindByID(id uint) (*User, error) {
    s.log.Info("Finding user by id", zap.Uint("id", id))
    s.log.Debug("Executing query...")
    
    var user User
    err := s.db.First(&user, id).Error
    return &user, err
}

// Setup in main()
logger, _ := zap.NewProduction()
defer logger.Sync()
```

**Key Differences:**
- Similar log levels (INFO, DEBUG, WARN, ERROR)
- Go requires explicit logger injection
- `zap` and `logrus` are popular Go logging libraries

---

## 11. API Documentation (OpenAPI/Swagger)

### Spring Boot (Springdoc)
```java
@Operation(summary = "Get user by ID")
@ApiResponses(value = {
    @ApiResponse(responseCode = "200", description = "User found"),
    @ApiResponse(responseCode = "404", description = "User not found")
})
@GetMapping("/{id}")
public ResponseEntity<User> getUser(@PathVariable Long id) {
    // ...
}
```

Auto-generates `/swagger-ui.html`.

### Go + Mux (swaggo/swag)
```go
// GetUser godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id} [get]
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    // ...
}

// Generate docs
// swag init -g main.go
// Visit /swagger/index.html
```

**Key Differences:**
- Spring: annotations on methods
- Go: comments above handlers (parsed by swag)
- Both generate Swagger UI

---

## 12. Graceful Shutdown

### Spring Boot
```yaml
# application.yml
server:
  shutdown: graceful
spring:
  lifecycle:
    timeout-per-shutdown-phase: 30s
```

### Go + Mux
```go
import (
    "context"
    "os/signal"
    "syscall"
)

func main() {
    srv := &http.Server{
        Addr:         ":8000",
        Handler:      r,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
    
    // Listen for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    
    // Graceful shutdown with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
```

**Key Differences:**
- Spring: YAML configuration
- Go: Explicit signal handling code
- Same result (graceful shutdown with timeout)

---

## Libraries Cheat Sheet

| Purpose | Spring Boot | Go + Mux Equivalent |
|---------|-------------|---------------------|
| **Web Framework** | Spring MVC | `gorilla/mux`, `chi`, `gin` |
| **ORM** | Hibernate/JPA | `gorm`, `sqlx` |
| **DB Migration** | Flyway, Liquibase | `golang-migrate/migrate`, `pressly/goose` |
| **Validation** | Jakarta Validation | `go-playground/validator` |
| **Cron** | `@Scheduled` | `robfig/cron` |
| **Cache** | Spring Cache + Redis | `go-redis/redis` |
| **Security** | Spring Security | Custom middleware + `golang-jwt/jwt` |
| **Logging** | SLF4J + Logback | `zap`, `logrus` |
| **Config** | `@ConfigurationProperties` | `caarlos0/env`, `joho/godotenv` |
| **Testing** | JUnit, MockMvc | `testing`, `testify`, `httptest` |
| **API Docs** | Springdoc OpenAPI | `swaggo/swag` |
| **HTTP Client** | RestTemplate, WebClient | `net/http`, `go-resty/resty` |

---

## What Go Doesn't Have (Out of the Box)

| Feature | Spring Boot | Go + Mux |
|---------|-------------|----------|
| Dependency Injection | ✅ Built-in (Spring Container) | ❌ Manual or use `google/wire` |
| AOP (Aspects) | ✅ `@Aspect` | ❌ Not available |
| Auto-configuration | ✅ Spring Boot starters | ❌ Manual setup |
| Embedded Server | ✅ Tomcat embedded | ❌ Use `net/http` or add `embed` |
| Actuator/Health | ✅ `/actuator/health` | ❌ Build your own endpoint |
| Profiles | ✅ `@Profile` | ❌ Use env vars |

---

## What Go Does Better

| Aspect | Why Go Wins |
|--------|-------------|
| **Startup Time** | milliseconds vs seconds |
| **Memory Usage** | ~10MB vs ~200MB+ |
| **Binary Size** | ~10MB single binary vs fat JAR |
| **Deployment** | Single binary, no JVM needed |
| **Performance** | Native code, no GC pauses like JVM |
| **Simplicity** | Less magic, more explicit |

---

## Migration Strategy

If you're moving from Spring Boot to Go:

1. **Start with new microservices** in Go, keep existing Spring Boot apps
2. **Map patterns**: Controller → Handler, Service → Service, Repository → Repository
3. **Use similar structure**: Keep familiar layering (handlers → services → repositories)
4. **Leverage your knowledge**: REST, middleware, DI concepts all transfer
5. **Embrace simplicity**: Go has less magic - more explicit but clearer

---

## Summary

| Aspect | Spring Boot | Go + Mux |
|--------|-------------|----------|
| **Philosophy** | Convention over configuration | Simplicity and explicitness |
| **Learning Curve** | Steep (many concepts) | Gentle (fewer concepts) |
| **Boilerplate** | Annotations everywhere | More manual code |
| **Magic** | A lot (auto-config, AOP) | Almost none |
| **Performance** | Good (JIT compiled) | Excellent (native compiled) |
| **Ecosystem** | Mature, extensive | Growing, pragmatic |

**Bottom Line:** If you know Spring Boot, you already understand 80% of what Go web development is about. The patterns are the same - REST APIs, middleware, DI, databases. Go just makes you write it out explicitly instead of using annotations.
