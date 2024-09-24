package main
import ( "encoding/json"
"fmt")

type Todo struct {
//TODO
}

var todos []Todo

func getTodos(/*TODO*/){
	//TODO
}
func createTodo(/*TODO*/) {
	//TODO
}

func getTodo(/*TODO*/) {
	//TODO
}

func updateTodo(/*TODO*/) {
	/*TODO*/

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Todo not found"})
}