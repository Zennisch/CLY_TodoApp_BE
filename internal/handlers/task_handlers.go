package handlers

import (
	"CLY_TodoApp_BE/internal/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	tasks  []models.Task
	nextID int
	mutex  sync.RWMutex // An toàn cho xử lý đa luồng
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		tasks:  make([]models.Task, 0),
		nextID: 1,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest

	// Ánh xạ JSON vào struct CreateTaskRequest
	// Nếu có lỗi trong quá trình ánh xạ, trả về lỗi 400 Bad Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Khóa mutex để đảm bảo an toàn khi truy cập và sửa đổi danh sách tasks
	// Sử dụng defer để mở khóa mutex sau khi hoàn thành
	h.mutex.Lock()
	defer h.mutex.Unlock()

	task := models.Task{
		ID:          h.nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	h.tasks = append(h.tasks, task)
	h.nextID++

	response := models.CreateTaskResponse{
		Task: &task,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	// Khóa mutex để đảm bảo an toàn khi truy cập danh sách tasks trong khi không có luồng nào khác đang sửa đổi nó
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	// Chuyển đổi danh sách tasks thành danh sách con trỏ để trả về
	taskPointers := make([]*models.Task, len(h.tasks))
	for i := range h.tasks {
		taskPointers[i] = &h.tasks[i]
	}

	response := models.GetTasksResponse{
		Tasks: taskPointers,
	}

	c.JSON(http.StatusOK, response)
}
