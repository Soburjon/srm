package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/work/SRM/api/V1/controllers"
)

func ProgrammerRoutes(app *fiber.App, c *controllers.Api) {
	routes := app.Group("/programmer")

	//get
	routes.Get("/get-project-tasks/:project_id/", c.GetProjectTasks)
	routes.Get("/get-task/:task_id/", c.GetTask)
	routes.Get("/get-my-tasks/:project_id/", c.GetMyTasks)
	routes.Get("/get-commits/:task_id/", c.GetCommits)
	routes.Get("/users-in-project/:project_id/", c.UsersInProject)
	routes.Get("/get-my-projects/", c.MyProjects)
	routes.Get("/get-project/:project_id/", c.GetProjects)

	//post
	routes.Post("/create-task/", c.CreateTask)
	routes.Post("/create-commit/", c.CreateCommit)
	routes.Post("/create-attendance/:type/", c.CreateAttendance)
	routes.Post("/programmer/add-people-project/", c.AddPeoplesProject)

	//put
	routes.Put("/edit-task/", c.EditTask)
	routes.Put("/update-task-status/", c.UpdateTaskStatus)
	routes.Put("/edit-commit/", c.EditCommit)

	//delete
	routes.Delete("/delete-task/:task_id/", c.DeleteTask)
	routes.Delete("/delete-commit/:created_at/", c.DeleteCommit)

}
