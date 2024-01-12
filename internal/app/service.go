package app

import (
	userHandler "myapp/internal/user/delivery/http/v1"
	taskHandler "myapp/internal/task/delivery/http/v1"

	userRepository "myapp/internal/user/repository"
	taskRepository "myapp/internal/task/repository"

	userUseCase "myapp/internal/user/usecase"
	taskUseCase "myapp/internal/task/usecase"
)

func (app *App) StartService() error {
	// define repository
	userRepo := userRepository.NewRepository(app.DB, app.Log)
	userRepository.AutoMigrate(app.DB) 

	taskRepo := taskRepository.NewRepository(app.DB, app.Log)
	taskRepository.AutoMigrate(app.DB) 

	// define usecase
	userUC := userUseCase.NewUseCase(userRepo, app.Log, app.Cfg)
	taskUC := taskUseCase.NewUseCase(taskRepo, app.Log, app.Cfg)

	// define controllers
	userCTRL := userHandler.NewHandlers(userHandler.HandlersDeps{
		UserUsecaseI: userUC,
	}, app.Log)

	taskCTRL := taskHandler.NewHandlers(taskHandler.HandlersDeps{
		TaskUsecaseI: taskUC,
	}, app.Log)

	version := app.Echo.Group("/api/v1/")

	userHandler.UserPrivateRoute(version, userCTRL, app.Cfg)
	taskHandler.UserPrivateRoute(version, taskCTRL, app.Cfg)

	return nil
}
