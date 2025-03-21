package routes

import(
"github.com/gofiber/fiber/v2"
"github.com/lokesh2201013/Usermanagement/controllers"
"github.com/lokesh2201013/Usermanagement/middleware"
)

func AuthRoutes(app *fiber.App){
   app.Post("/signup", controllers.Register)
   app.Post("/login", controllers.Login)

   app.Use(middleware.AuthMiddleware())
    
    app.Post("/createuser", middleware.AdminOnly(controllers.Createusers))
	app.Get("/getallusers", middleware.AdminOnly(controllers.GetAllusers))
	app.Get("/getuser/:id", middleware.AdminOnly(controllers.Getuser))
	app.Put("/updateuser/:id", middleware.AdminOnly(controllers.Updateusers))
	app.Delete("/deleteuser/:id", middleware.AdminOnly(controllers.Deleteuser))
	
}