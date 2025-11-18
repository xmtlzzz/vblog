package impl_test

import (
	"context"
	"testing"
	"time"

	"github.com/xmtlzzz/vblog/apps/user"
	"github.com/xmtlzzz/vblog/apps/user/impl"
)

var ctx = context.Background()

func TestRegistry(t *testing.T) {
	t.Setenv("workdir", "D:\\Desktop\\code\\Go\\vblog")
	request := user.NewRegistryRequest()
	request.Username = "xmtlz"
	request.Password = "123456"
	ins, err := impl.UserService.Registry(ctx, request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeUser(t *testing.T) {
	var describe user.DescribeUserRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	describe.DescribeBy = user.Describe_By_UserName
	describe.Value = "sz"
	ins, err := impl.UserService.DescribeUser(ctx, describe)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateUserStatus(t *testing.T) {
	var status user.UpdateUserStatusRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	status.Username = "sz"
	status.BlockAT = time.Now()
	status.BlockReason = "error login"
	ins, err := impl.UserService.UpdateUserStatus(ctx, status)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateUserProfile(t *testing.T) {
	var profile user.UpdateProfileRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	profile.Username = "xmtlz"
	profile.Email = "xmtlzloveasuka@gmail.com"
	profile.NickName = "bigchanzi"
	ins, err := impl.UserService.UpdateUserProfile(ctx, profile)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateUserPassword(t *testing.T) {
	var password user.UpdatePasswordRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	password.Username = "xmtlz"
	password.OldPassword = "7777777"
	password.NewPassword = "123123"
	err := impl.UserService.UpdateUserPassword(ctx, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(password)
}

func TestResetUserPassword(t *testing.T) {
	var password user.ResetPasswordRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	password.Username = "xmtlz"
	password.NewPassword = ""
	err := impl.UserService.ResetUserPassword(ctx, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(password)
}

func TestUnRegistry(t *testing.T) {
	var unregistry user.UnRegistryRequest
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	unregistry.Username = "xmtlz"
	err := impl.UserService.UnRegistry(ctx, unregistry)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(unregistry)
}
