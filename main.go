package main

import (
	"database/sql"
	// "github.com/aianman4823/todo-app-api/handler"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db, err = sql.Open("mysql", "admin:password@tcp(127.0.0.1:3306)/todos")
)

func main() {
	// db, err := sql.Open("mysql", "admin:password@tcp(127.0.0.1:3306)/todos")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer db.Close()
	cmd := `CREATE TABLE IF NOT EXISTS todos(
		id INT AUTO_INCREMENT PRIMARY KEY,    
		task VARCHAR(50) NOT NULL);`
	_, err := db.Exec(cmd)

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Method: GET
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/todos", GetAll)
	e.GET("/todos/:id", SelectTodos)

	// Method: POST
	e.POST("/todos", InsertTodo)

	// Method: PUT
	e.PUT("todo/:id", UpdateTodo)

	// Method: DELETE
	e.DELETE("/todos/:id", DeleteTodo)

	// localhost:1323で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// Todo: 型
type Todo struct {
	ID   int    `db:"id" json:"id"`
	Task string `db:"task" json:"task"`
}

// ResponseTodos
type ResponseTodos struct {
	Todos []ResponseTodo `json:"todos"`
}

// ResponseTodo
type ResponseTodo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// InsertTodo: POST用
func InsertTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	_, err := db.Exec(`INSERT INTO todos 
	(task) VALUES (?)`, todo.Task)

	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// GetAll
func GetAll(c echo.Context) error {
	var todos []ResponseTodo
	cmd := `SELECT * FROM todos`
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var todo ResponseTodo
		err := rows.Scan(&todo.ID, &todo.Task)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}
	res := ResponseTodos{}
	res.Todos = todos
	return c.JSON(http.StatusOK, res)
}

// SelectTodos
func SelectTodos(c echo.Context) error {
	var todo ResponseTodo
	id := c.Param("id")
	cmd := `SELECT * FROM todos WHERE id = ?`
	row := db.QueryRow(cmd, id)

	err = row.Scan(&todo.ID, &todo.Task)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Rowがない")
		} else {
			log.Println(err)
		}
	}
	return c.JSON(http.StatusOK, todo)
}

// UpdateTodo
func UpdateTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	id := c.Param("id")
	// attrsMap := map[string]interface{}{"id":todo.ID,"task":todo.Task}
	cmd := `UPDATE todos SET task = ? WHERE id = ?`
	_, err := db.Exec(cmd, todo.Task, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// DeleteTodo
func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cmd := "DELETE FROM todos WHERE id = ?"
	_, err = db.Exec(cmd, id)
	if err != nil {
		log.Fatal(err)
	}
	return c.NoContent(http.StatusOK)
}
