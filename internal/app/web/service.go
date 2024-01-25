package web

import (
	taskHttpService "myapp/internal/task/delivery/http/v1"
	workerHttpService "myapp/internal/worker/delivery/http/v1"

	taskRepository "myapp/internal/task/repository"
	workerRepository "myapp/internal/worker/repository"

	taskUseCase "myapp/internal/task/usecase"
	workerUseCase "myapp/internal/worker/usecase"
)

func (app *App) StartService() error {
	// define repository
	taskRepo := taskRepository.NewRepository(app.DB, app.Log)
	taskRepository.AutoMigrate(app.DB)

	workerRepo := workerRepository.NewRepository(app.DB, app.Log)
	workerRepository.AutoMigrate(app.DB)

	// define usecase
	taskUseCase := taskUseCase.NewUseCase(taskRepo, app.Log, app.Cfg)
	workerUseCase := workerUseCase.NewUseCase(workerRepo, app.Log, app.Cfg)

	// define controllers
	taskHttpSvc := taskHttpService.NewService(taskHttpService.ServiceDeps{
		TaskUsecaseI: taskUseCase,
	}, app.Log)

	workerHttpSvc := workerHttpService.NewService(workerHttpService.ServiceDeps{
		WorkerUsecaseI: workerUseCase,
	}, app.Log)

	version := app.Echo.Group("/api/v1/")
	taskHttpService.UserPrivateRoute(version, taskHttpSvc, app.Cfg)
	workerHttpService.UserPrivateRoute(version, workerHttpSvc, app.Cfg)

	return nil
}
