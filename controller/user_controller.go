package controller

import "coredemo/framework"

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, UserLoginController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, SubjectGetController")
	return nil
}
