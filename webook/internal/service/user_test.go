package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"signup_issue/webook/internal/domain"
	"signup_issue/webook/internal/repository"
	repomocks "signup_issue/webook/internal/repository/mocks"
	"testing"
)

func TestUserService_Signup(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(*gomock.Controller) repository.UserRepository
		user    *domain.User
		wantErr error
	}{
		{
			name: "注册成功！",
			user: &domain.User{
				Id:       123,
				Email:    "l0slakers@gmail.com",
				Password: "Abcd#1234",
			},
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(nil)
				return userRepo
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := NewUserService(tc.mock(ctrl))
			err := svc.Signup(context.Background(), tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
