package adapter

import (
	"context"
	"fmt"
	"github.com/LyricTian/gin-admin/v6/internal/app/model"
	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
	"github.com/LyricTian/gin-admin/v6/pkg/logger"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

// CasbinAdapterSet 注入CasbinAdapter
var CasbinAdapterSet = wire.NewSet(wire.Struct(new(CasbinAdapter), "*"), wire.Bind(new(persist.Adapter), new(*CasbinAdapter)))

// CasbinAdapter casbin适配器
type CasbinAdapter struct {
	AccountModel      	model.IAccount
	RoleModel        	model.IRole
	ResourceModel 		model.IResource
	AccountRoleModel  	model.IAccountRole
	RoleResourceModel   model.IRoleResource
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := a.loadRolePolicy(ctx, model)
	if err != nil {
		logger.Errorf(ctx, "Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(ctx, model)
	if err != nil {
		logger.Errorf(ctx, "Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// 加载角色资源策略(p,roleID,appKey,resourceType,feature,method)
func (a *CasbinAdapter) loadRolePolicy(ctx context.Context, m casbinModel.Model) error {
	roleResult, err := a.RoleModel.Query(ctx, schema.RoleQueryParam{
		//Status: 1,
	})

	if err != nil {
		return err
	} else if len(roleResult.Data) == 0 {
		return nil
	}

	roleResourceResult, err := a.RoleResourceModel.Query(ctx, schema.RoleResourceQueryParam{})
	if err != nil {
		return err
	}
	mRoleResources := roleResourceResult.Data.ToRoleIDMap()

	resourceResult, err := a.ResourceModel.Query(ctx, schema.ResourceQueryParam{})
	if err != nil {
		return err
	}
	mResources := resourceResult.Data.ToMap()

	for _, item := range roleResult.Data {
		mcache := make(map[string]struct{})
		if rms, ok := mRoleResources[item.ID]; ok {
			for _, resourceID := range rms.ToResourceIDs() {
				if mrs, ok := mResources[resourceID]; ok {
					flag := mrs.AppKey + "/" + mrs.Type + "/" + mrs.Feature + "/" + mrs.Method
					if _, ok := mcache[flag]; ok {
						continue
					}
					mcache[flag] = struct{}{}
					line := fmt.Sprintf("p,%v,%v,%v,%v,%v", item.ID, mrs.AppKey, mrs.Type, mrs.Feature, mrs.Method)
					persist.LoadPolicyLine(line, m)
				}
			}
		}
	}

	return nil
}

// 加载用户角色策略(g,accountKey,roleID)
func (a *CasbinAdapter) loadUserPolicy(ctx context.Context, m casbinModel.Model) error {
	userResult, err := a.AccountModel.Query(ctx, schema.AccountQueryParam{})

	if err != nil {
		return err
	} else if len(userResult.Data) > 0 {
		userRoleResult, err := a.AccountRoleModel.Query(ctx, schema.AccountRoleQueryParam{})
		if err != nil {
			return err
		}

		mUserRoles := userRoleResult.Data.ToAccountKeyMap()
		for _, uitem := range userResult.Data {
			if urs, ok := mUserRoles[uitem.AccountKey]; ok {
				for _, ur := range urs {
					line := fmt.Sprintf("g,%v,%v", ur.AccountKey, ur.RoleID)
					persist.LoadPolicyLine(line, m)
				}
			}
		}
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
