package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/work/SRM/api/V1/controllers"
)

func AdminRoutes(app *fiber.App, c *controllers.Api) {
	routes := app.Group("/admin")

	//get
	routes.Get("/get-admin-list/", c.GetAminList)
	routes.Get("/get-programmer-list/", c.GetProgrammerList)
	routes.Get("/get-project-list/", c.GetProjectsList)
	routes.Get("/get-my-projects/", c.GetMyProjects)
	routes.Get("/get-user/:user_id/", c.GetUser)
	routes.Get("/get-project/:project_id/", c.GetProject)
	routes.Get("/get-attendance/:programmer_id/", c.GetAttendance)

	//post
	routes.Post("/create-admin/", c.CreateAdmin)
	routes.Post("/create-programmer/", c.CreateProgrammer)
	routes.Post("/create-project/", c.CreateProject)
	routes.Post("/add-people-project/", c.AddPeopleProject)

	//put
	routes.Put("/edit-admin/", c.EditAdmin)
	routes.Put("/edit-programmer/", c.EditProgrammer)
	routes.Put("/edit-project/", c.EditProject)
	routes.Put("/update-project-status/", c.UpdateProjectStatus)

	//delete
	routes.Delete("/delete-admin/:admin_id/", c.DeleteAdmin)
	routes.Delete("/delete-programmer/:programmer_id/", c.DeleteProgrammer)
	routes.Delete("/delete-project/:project_id/", c.DeleteProject)
}
