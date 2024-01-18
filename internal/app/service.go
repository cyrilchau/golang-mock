package app

import (
	pb "myapp/proto/pb"

	userGrpcService "myapp/internal/user/delivery/grpc"

	taskHttpService "myapp/internal/task/delivery/http/v1"
	userHttpService "myapp/internal/user/delivery/http/v1"

	taskRepository "myapp/internal/task/repository"
	userRepository "myapp/internal/user/repository"

	taskUseCase "myapp/internal/task/usecase"
	userUseCase "myapp/internal/user/usecase"
)

func (app *App) StartService() error {
	// define repository
	userRepo := userRepository.NewRepository(app.DB, app.Log)
	userRepository.AutoMigrate(app.DB)

	taskRepo := taskRepository.NewRepository(app.DB, app.Log)
	taskRepository.AutoMigrate(app.DB)

	// define usecase
	userUseCase := userUseCase.NewUseCase(userRepo, app.Log, app.Cfg)
	taskUseCase := taskUseCase.NewUseCase(taskRepo, app.Log, app.Cfg)

	// define controllers
	userHttpSvc := userHttpService.NewService(userHttpService.ServiceDeps{
		UserUsecaseI: userUseCase,
	}, app.Log)

	taskHttpSvc := taskHttpService.NewService(taskHttpService.ServiceDeps{
		TaskUsecaseI: taskUseCase,
	}, app.Log)

	version := app.Echo.Group("/api/v1/")
	userHttpService.UserPrivateRoute(version, userHttpSvc, app.Cfg)
	taskHttpService.UserPrivateRoute(version, taskHttpSvc, app.Cfg)

	pb.RegisterUserServiceServer(app.GRPC, userGrpcService.NewUserService(userGrpcService.ServiceDeps{UserUsecaseI: userUseCase}, app.Log))

	return nil
}
