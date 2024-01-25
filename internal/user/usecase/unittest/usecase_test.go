package usecase

import (
	"context"
	"myapp/config"
	"myapp/internal/user/dtos"
	"myapp/internal/user/entity"
	"myapp/internal/user/usecase"
	errorCode "myapp/pkg/errorCode"
	"myapp/pkg/otel/zerolog"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cfg, _ = config.LoadConfig("auth")

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

// Mocked implementation for method in the repository.
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) CreateOneUser(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func TestCreateUserUseCase_Execute(t *testing.T) {

	ctx := context.Background()

	// Arrange
	mockRepo := new(MockUserRepository)
	createUserUseCase := usecase.NewUseCase(mockRepo, zerolog.NewZeroLog(ctx, os.Stdout),
		config.Config{
			Authentication: config.AuthenticationConfig{
				Key:       cfg.Authentication.Key,
				SecretKey: cfg.Authentication.SecretKey,
				SaltKey:   cfg.Authentication.SaltKey,
			},
		},
	)

	// Define mock data for the test
	mockUser := dtos.CreateUserRequest{
		FullName:    "fullname",
		PhoneNumber: "081236548974",
		Email:       "martinyonatann@testing.com",
		Password:    "testingPwd",
	}

	// Expectation: The repository's GetUserByEmail, CreateOneUser method should be called with the mockUser.
	mockRepo.On("GetUserByEmail", ctx, mockUser.Email).Return(entity.User{}, nil)
	mockRepo.On("CreateOneUser", ctx, mock.Anything).Return(entity.User{}, nil)

	// Act
	_, _, err := createUserUseCase.Create(ctx, mockUser)

	// Assert
	assert.NoError(t, err, "Unexpected error during user creation")

	// Assert that the CreateUser method was called with the expected user.
	mockRepo.AssertExpectations(t)
}

func TestCreateUserUseCase_ExecuteExistEmail(t *testing.T) {

	ctx := context.Background()

	// Arrange
	mockRepo := new(MockUserRepository)
	createUserUseCase := usecase.NewUseCase(mockRepo, zerolog.NewZeroLog(ctx, os.Stdout),
		config.Config{
			Authentication: config.AuthenticationConfig{
				Key:       cfg.Authentication.Key,
				SecretKey: cfg.Authentication.SecretKey,
				SaltKey:   cfg.Authentication.SaltKey,
			},
		},
	)

	// Define mock data for the test
	mockUser := dtos.CreateUserRequest{
		FullName:    "fullname",
		PhoneNumber: "081236548974",
		Email:       "martinyonatann@testing.com",
		Password:    "testingPwd",
	}

	// Expectation: The repository's GetUserByEmail method should be called with the mockUser's email.
	mockRepo.On("GetUserByEmail", ctx, mockUser.Email).Return(entity.User{UserID: 1, Fullname: mockUser.FullName}, nil).Once()

	// Expectation: The repository's CreateOneUser method should not be called.
	mockRepo.AssertNotCalled(t, "CreateOneUser", ctx, mock.Anything)

	// Act
	_, httpCode, err := createUserUseCase.Create(ctx, mockUser)

	// Assert
	assert.Equal(t, http.StatusConflict, httpCode, "Expected conflict status for existing user")
	assert.Equal(t, errorCode.ErrEmailAlreadyExist, err, "Expected ErrEmailAlreadyExist error")

	// Assert that the CreateUser method was called with the expected user.
	mockRepo.AssertExpectations(t)
}
