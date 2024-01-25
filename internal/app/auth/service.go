package auth

import (
	pb "myapp/proto/pb"

	userGrpcService "myapp/internal/user/delivery/grpc"

	userHttpService "myapp/internal/user/delivery/http/v1"

	userRepository "myapp/internal/user/repository"

	userUseCase "myapp/internal/user/usecase"
)

func (app *App) StartService() error {
	// define repository
	userRepo := userRepository.NewRepository(app.DB, app.Log)
	userRepository.AutoMigrate(app.DB)

	// define usecase
	userUseCase := userUseCase.NewUseCase(userRepo, app.Log, app.Cfg)

	// define controllers
	userHttpSvc := userHttpService.NewService(userHttpService.ServiceDeps{
		UserUsecaseI: userUseCase,
	}, app.Log)

	version := app.Echo.Group("/api/v1/")
	userHttpService.UserPrivateRoute(version, userHttpSvc, app.Cfg)

	pb.RegisterUserServiceServer(app.GRPC, userGrpcService.NewUserService(userGrpcService.ServiceDeps{UserUsecaseI: userUseCase}, app.Log))

	return nil
}
