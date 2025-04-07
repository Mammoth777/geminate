package router

import (
	"github.com/gin-gonic/gin"

	"api-flow/handlers"
	"api-flow/services"
)

// SetupRouter 配置API路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 创建服务实例
	workflowService := services.NewWorkflowService()
	nodeService := services.NewNodeService()
	nodeExecutionService := services.NewNodeExecutionService(nodeService)

	// 创建处理器实例
	workflowHandler := handlers.NewWorkflowHandler(workflowService)
	nodeHandler := handlers.NewNodeHandler(nodeService, nodeExecutionService)

	// 定义API路由
	api := r.Group("/api")
	{
		// 工作流路由
		workflows := api.Group("/workflows")
		{	
			workflows.POST("/save", workflowHandler.CreateWithNodes)
			workflows.GET("", workflowHandler.List)
			workflows.GET("/:id", workflowHandler.Get)
			workflows.PUT("/:id", workflowHandler.Update)
			workflows.DELETE("/:id", workflowHandler.Delete)
		}

		// 节点路由
		nodes := api.Group("/nodes")
		{
			nodes.GET("/:id", nodeHandler.Get)
			// TODO: 此处可能要走workflow整体的保存/更新逻辑
			// nodes.PUT("/:id", nodeHandler.Update)
			// nodes.DELETE("/:id", nodeHandler.Delete)
			nodes.POST("/:id/execute", nodeHandler.Execute) // 执行节点
		}

		// 节点类型路由
		api.GET("/node-types", nodeHandler.GetNodeTypes)
	}

	return r
}
