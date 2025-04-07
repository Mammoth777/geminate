package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-flow/dto"
	"api-flow/models"
	"api-flow/services"
)

// WorkflowHandler 处理流程相关API
type WorkflowHandler struct {
	workflowService *services.WorkflowService
}

// NewWorkflowHandler 创建流程处理器实例
func NewWorkflowHandler(workflowService *services.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{
		workflowService: workflowService,
	}
}

// Get 获取单个流程
func (h *WorkflowHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID参数"})
		return
	}

	workflow, err := h.workflowService.GetWorkflowWithNodes(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workflow)
}

// List 获取流程列表
func (h *WorkflowHandler) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	workflows, count, err := h.workflowService.GetAllWorkflows(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": count,
		"data":  workflows,
	})
}

// Update 更新流程
func (h *WorkflowHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID参数"})
		return
	}

	var workflow models.Workflow
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.workflowService.UpdateWorkflow(uint(id), &workflow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "流程更新成功"})
}

// Delete 删除流程
func (h *WorkflowHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID参数"})
		return
	}

	if err := h.workflowService.DeleteWorkflow(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "流程删除成功"})
}

// CreateWithNodes 创建工作流及其节点
func (h *WorkflowHandler) CreateWithNodes(c *gin.Context) {
	var workflowDTO dto.WorkflowDTO
	if err := c.ShouldBindJSON(&workflowDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.workflowService.SaveWorkflowWithNodes(&workflowDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

