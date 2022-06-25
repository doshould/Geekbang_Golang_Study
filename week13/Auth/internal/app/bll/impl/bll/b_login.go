package bll

import (
	"context"
	"github.com/LyricTian/gin-admin/v6/internal/app/bll"
	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/auth"
	"github.com/LyricTian/gin-admin/v6/pkg/errors"
	"github.com/LyricTian/gin-admin/v6/pkg/util"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
)

var _ bll.ILogin = (*Login)(nil)

// LoginSet 注入Login
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"), wire.Bind(new(bll.ILogin), new(*Login)))

// Login 登录管理
type Login struct {
	Enforcer      		*casbin.SyncedEnforcer
	Auth            	auth.Auther
	AppModel       		model.IApp
	AccountModel       	model.IAccount
	AccountRoleModel   	model.IAccountRole
	RoleModel       	model.IRole
	RoleResourceModel   model.IRoleResource
	ResourceModel       model.IResource
}

// Verify 登录验证
func (a *Login) Verify(ctx context.Context, appKey, username, password string) (*schema.Account, error) {
	// 检查是否是超级用户
	root := schema.GetRootAccount()
	if root.Username == username && root.Password == password {
		return root, nil
	}

	appResult, err := a.AppModel.Query(ctx, schema.AppQueryParam{
		PaginationParam: 	schema.PaginationParam{OnlyCount: true},
		AppKey: appKey,
	})
	if err != nil {
		return nil, err
	} else if appResult.PageResult.Total == 0 {
		return nil, errors.New403Response(errors.ErrCodeAppNotExist, "该应用尚未接入")
	}

	result, err := a.AccountModel.Query(ctx, schema.AccountQueryParam{
		AppKey: appKey,
		Username: username,
		Password: util.MD5HashString(password),
	})

	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.New403Response(errors.ErrCodeLoginFailed, "登录失败，账号或密码错误")
	}

	item := result.Data[0]

	return item, nil
}

// GenerateToken 生成令牌
func (a *Login) GenerateToken(ctx context.Context, accountKey string, info interface{}) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := a.Auth.GenerateToken(ctx, accountKey, info)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		Token: tokenInfo.GetAccessToken(),
		ExpiresAt: tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

// Authenticate 鉴权
func (a *Login) Authenticate(ctx context.Context, appKey, accountKey, resourceType, feature, method string) (bool, error) {
	b, err := a.Enforcer.Enforce(accountKey, appKey, resourceType, feature, method)

	if err != nil {
		return false, err
	}

	return b, nil
}

//// GetCaptcha 获取图形验证码信息
//func (a *Login) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
//	captchaID := captcha.NewLen(length)
//	item := &schema.LoginCaptcha{
//		CaptchaID: captchaID,
//	}
//	return item, nil
//}
//
//// ResCaptcha 生成并响应图形验证码
//func (a *Login) ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error {
//	err := captcha.WriteImage(w, captchaID, width, height)
//	if err != nil {
//		if err == captcha.ErrNotFound {
//			return errors.ErrNotFound
//		}
//		return errors.WithStack(err)
//	}
//
//	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
//	w.Header().Set("Pragma", "no-cache")
//	w.Header().Set("Expires", "0")
//	w.Header().Set("Content-Type", "image/png")
//	return nil
//}

//// DestroyToken 销毁令牌
//func (a *Login) DestroyToken(ctx context.Context, tokenString string) error {
//	err := a.Auth.DestroyToken(ctx, tokenString)
//	if err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//func (a *Login) checkAndGetUser(ctx context.Context, userID string) (*schema.Account, error) {
//	user, err := a.UserModel.Get(ctx, userID)
//	if err != nil {
//		return nil, err
//	} else if user == nil {
//		return nil, errors.ErrInvalidUser
//	}
//	return user, nil
//}
//
//// GetLoginInfo 获取当前用户登录信息
//func (a *Login) GetLoginInfo(ctx context.Context, userID string) (*schema.UserLoginInfo, error) {
//	if isRoot := schema.CheckIsRootUser(ctx, userID); isRoot {
//		root := schema.GetRootUser()
//		loginInfo := &schema.UserLoginInfo{
//			UserName: root.Username,
//		}
//		return loginInfo, nil
//	}
//
//	user, err := a.checkAndGetUser(ctx, userID)
//	if err != nil {
//		return nil, err
//	}
//
//	info := &schema.UserLoginInfo{
//		UserID:   user.AccountKey,
//		UserName: user.Username,
//	}
//
//	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
//		UserID: userID,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	if roleIDs := userRoleResult.Data.ToRoleIDs(); len(roleIDs) > 0 {
//		roleResult, err := a.RoleModel.Query(ctx, schema.RoleQueryParam{
//			IDs:    roleIDs,
//			Status: 1,
//		})
//		if err != nil {
//			return nil, err
//		}
//		info.Roles = roleResult.Data
//	}
//
//	return info, nil
//}
//
//// QueryUserMenuTree 查询当前用户的权限菜单树
//func (a *Login) QueryUserMenuTree(ctx context.Context, userID string) (schema.MenuTrees, error) {
//	isRoot := schema.CheckIsRootUser(ctx, userID)
//	// 如果是root用户，则查询所有显示的菜单树
//	if isRoot {
//		result, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
//			Status: 1,
//		}, schema.MenuQueryOptions{
//			OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
//		})
//		if err != nil {
//			return nil, err
//		}
//
//		menuActionResult, err := a.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{})
//		if err != nil {
//			return nil, err
//		}
//		return result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree(), nil
//	}
//
//	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
//		UserID: userID,
//	})
//	if err != nil {
//		return nil, err
//	} else if len(userRoleResult.Data) == 0 {
//		return nil, errors.ErrNoPerm
//	}
//
//	roleMenuResult, err := a.RoleMenuModel.Query(ctx, schema.RoleMenuQueryParam{
//		RoleIDs: userRoleResult.Data.ToRoleIDs(),
//	})
//	if err != nil {
//		return nil, err
//	} else if len(roleMenuResult.Data) == 0 {
//		return nil, errors.ErrNoPerm
//	}
//
//	menuResult, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
//		IDs:    roleMenuResult.Data.ToMenuIDs(),
//		Status: 1,
//	})
//	if err != nil {
//		return nil, err
//	} else if len(menuResult.Data) == 0 {
//		return nil, errors.ErrNoPerm
//	}
//
//	mData := menuResult.Data.ToMap()
//	var qIDs []string
//	for _, pid := range menuResult.Data.SplitParentIDs() {
//		if _, ok := mData[pid]; !ok {
//			qIDs = append(qIDs, pid)
//		}
//	}
//
//	if len(qIDs) > 0 {
//		pmenuResult, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
//			IDs: menuResult.Data.SplitParentIDs(),
//		})
//		if err != nil {
//			return nil, err
//		}
//		menuResult.Data = append(menuResult.Data, pmenuResult.Data...)
//	}
//
//	sort.Sort(menuResult.Data)
//	menuActionResult, err := a.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{
//		IDs: roleMenuResult.Data.ToActionIDs(),
//	})
//	if err != nil {
//		return nil, err
//	}
//	return menuResult.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree(), nil
//}
//
//// UpdatePassword 更新当前用户登录密码
//func (a *Login) UpdatePassword(ctx context.Context, userID string, params schema.UpdatePasswordParam) error {
//	if schema.CheckIsRootUser(ctx, userID) {
//		return errors.New400Response("root用户不允许更新密码")
//	}
//
//	user, err := a.checkAndGetUser(ctx, userID)
//	if err != nil {
//		return err
//	} else if util.SHA1HashString(params.OldPassword) != user.Password {
//		return errors.New400Response("旧密码不正确")
//	}
//
//	params.NewPassword = util.SHA1HashString(params.NewPassword)
//	return a.UserModel.UpdatePassword(ctx, userID, params.NewPassword)
//}
