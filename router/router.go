package router

import (
	"coredemo/controller"
	"coredemo/framework"
)

// 注册路由规则
func RegisterRouter(core *framework.Core) {
	// HTTP方法+静态路由匹配
	core.Get("/user/login", controller.UserLoginController)

	// 批量通用路由匹配
	subjectApi := core.Group("/subject")
	{
		// 实现动态路由
		subjectApi.Delete("/:id", controller.SubjectDelController)
		subjectApi.Put("/:id", controller.SubjectUpdateController)
		subjectApi.Get("/:id", controller.SubjectGetController)
		subjectApi.Get("/list/all", controller.SubjectListController)
	}
}
