package impl_test

import (
	"context"
	"testing"

	"github.com/xmtlzzz/vblog/apps/token"
)

var ctx = context.Background()
var accessToken = "e58547c6-ad1d-4d98-b696-a4b88c11c2f1"
var refreshToken = "6bfa8c22-75ee-42fe-a189-2a35d1988a6a"

func TestIssueToken(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = token.IssueTokenRequest{
		Username: "sz",
		Password: "123456",
	}
	ins, err := token.GetService().IssueToken(ctx, it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestValidateToken(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = token.ValidateTokenRequest{
		AccessToken: accessToken,
	}
	ins, err := token.GetService().ValidateToken(ctx, &it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestRevolkToken(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = token.RevolkTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	ins, err := token.GetService().RevolkToken(ctx, it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestRefreshToken(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = token.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}
	ins, err := token.GetService().RefreshToken(ctx, it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
