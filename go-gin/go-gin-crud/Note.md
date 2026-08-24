
## 1. Folder + module setup

Inside the existing repo:

```bash
cd C:\tmp\projects\golang-references

mkdir go-gin-crud
cd go-gin-crud

go mod init example.com/go-gin-crud
```

Then add dependencies:

```bash
go get github.com/gin-gonic/gin
go get github.com/lib/pq
```

---

## 2. `main.go` — Gin + Postgres CRUD

Create `go-gin-crud/main.go`:

```go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// ApplicationUser mirrors labschema.application_user
type ApplicationUser struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // in real systems this should be a hash, never plain text
	Role     string `json:"role"`
	Username string `json:"username"`
}

// Request body for create/update (ID comes from DB / URL)
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Username string `json:"username" binding:"required"`
}

var db *sql.DB

func main() {
	// Connect to Postgres
	var err error
	connStr := "host=localhost port=9991 user=postgres password=admin dbname=mainschema sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("error pinging DB:", err)
	}

	log.Println("Connected to Postgres on port 9991, DB mainschema")

	// Set up Gin router
	router := gin.Default()

	// Simple health/status
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// CRUD routes
	router.GET("/users", listUsersHandler)
	router.GET("/users/:id", getUserHandler)
	router.POST("/users", createUserHandler)
	router.PUT("/users/:id", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)

	// Start server
	addr := ":8080"
	log.Println("Starting Gin server on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("server error:", err)
	}
}

// GET /users
func listUsersHandler(c *gin.Context) {
	rows, err := db.Query(`
		SELECT id, email, "password", "role", username
		FROM labschema.application_user
	`)
	if err != nil {
		log.Println("listUsersHandler: query error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []ApplicationUser
	for rows.Next() {
		var u ApplicationUser
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.Username); err != nil {
			log.Println("listUsersHandler: scan error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan user"})
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Println("listUsersHandler: rows error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to iterate users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var u ApplicationUser
	err = db.QueryRow(`
		SELECT id, email, "password", "role", username
		FROM labschema.application_user
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.Username)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		log.Println("getUserHandler: query error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, u)
}

// POST /users
func createUserHandler(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	var id int64
	err := db.QueryRow(`
		INSERT INTO labschema.application_user (email, "password", "role", username)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, req.Email, req.Password, req.Role, req.Username).Scan(&id)

	if err != nil {
		log.Println("createUserHandler: insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	user := ApplicationUser{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		Username: req.Username,
	}

	c.JSON(http.StatusCreated, user)
}

// PUT /users/:id
func updateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := db.Exec(`
		UPDATE labschema.application_user
		SET email = $1,
		    "password" = $2,
		    "role" = $3,
		    username = $4
		WHERE id = $5
	`, req.Email, req.Password, req.Role, req.Username, id)

	if err != nil {
		log.Println("updateUserHandler: update error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("updateUserHandler: rowsAffected error:", err)
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	user := ApplicationUser{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		Username: req.Username,
	}

	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func deleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, err := db.Exec(`
		DELETE FROM labschema.application_user
		WHERE id = $1
	`, id)
	if err != nil {
		log.Println("deleteUserHandler: delete error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("deleteUserHandler: rowsAffected error:", err)
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
```

---

## 3. How to run and test

From `go-gin-crud`:

```bash
go run main.go
```

Server listens on **`http://localhost:8080`**.

Example curl calls:

List all users (maps to `SELECT id, email, "password", "role", username FROM labschema.application_user`):

```bash
curl http://localhost:8080/users
```

Get a single user:

```bash
curl http://localhost:8080/users/1
```

Create a user:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"admin123","role":"USER","username":"testuser"}'
```

Update a user:

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"email":"new@example.com","password":"newpass","role":"ADMIN","username":"newname"}'
```

Delete a user:

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

## 4. Security notes (worth keeping in back pocket)

* Plain-text passwords in the table and JSON responses are **not** safe for any real system.
* In a realistic setup:

  * `password` would be a hash (bcrypt/argon2).
  * API responses would not return the password field at all.
  * Role/permissions would control who can call these endpoints.

