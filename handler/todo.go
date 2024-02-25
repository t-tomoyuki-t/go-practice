package handler

import (
	"github.com/gin-gonic/gin"
	"go-practice/domain/object"
	"go-practice/usecase"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	u *usecase.TodoUseCase
}

func NewTodoHandler(u *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{u}
}

func (th *TodoHandler) GetTodoList(c *gin.Context) {
	l, err := th.u.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"todos": l})
}

func (th *TodoHandler) GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t, err := th.u.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"todo": t})
}

func (th *TodoHandler) Store(c *gin.Context) {
	var t object.Todo
	c.BindJSON(&t)
	err := th.u.Store(&t)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, nil)
}

func (th *TodoHandler) Update(c *gin.Context) {
	var t object.Todo
	c.BindUri(&t)
	c.BindJSON(&t)
	err := th.u.Update(&t)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (th *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := th.u.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}
