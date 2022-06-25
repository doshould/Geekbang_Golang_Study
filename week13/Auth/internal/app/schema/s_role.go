package schema

// Role 角色对象
type Role struct {
	ID        int    	`json:"id"`                                    // 唯一标识
	Name      string    `json:"name" binding:"required"`               // 角色名称
	//CreatedAt time.Time `json:"created_at"`                            // 创建时间
	//UpdatedAt time.Time `json:"updated_at"`                            // 更新时间
	//RoleMenus RoleMenus `json:"role_menus" binding:"required,gt=0"`    // 角色菜单列表
}

// RoleQueryParam 查询条件
type RoleQueryParam struct {
	PaginationParam
	IDs        []int 	`form:"-"`          // 唯一标识列表
	AppKey     string   `form:"appKey"`     // 应用标识
	Name       string   `form:"name"`       // 角色名称
	AccountKey string   `form:"accountKey"`	// 账号标识
}

// RoleQueryOptions 查询可选参数项
type RoleQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// RoleQueryResult 查询结果
type RoleQueryResult struct {
	Data       Roles
	PageResult *PaginationResult
}

// Roles 角色对象列表
type Roles []*Role

// ToNames 获取角色名称列表
func (a Roles) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Roles) ToMap() map[int]*Role {
	m := make(map[int]*Role)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// ----------------------------------------RoleResource--------------------------------------

// RoleResource 角色资源对象
type RoleResource struct {
	ID       	int `json:"id"`                           	// 唯一标识
	RoleID   	int `json:"roleID" binding:"required"`   	// 角色ID
	ResourceID  int `json:"resourceID" binding:"required"`	// 资源ID
}

// RoleResourceQueryParam 查询条件
type RoleResourceQueryParam struct {
	PaginationParam
	RoleID  string   // 角色ID
	RoleIDs []string // 角色ID列表
}

// RoleResourceQueryOptions 查询可选参数项
type RoleResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// RoleResourceQueryResult 查询结果
type RoleResourceQueryResult struct {
	Data       RoleResources
	PageResult *PaginationResult
}

// RoleResources 角色资源列表
type RoleResources []*RoleResource

// ToRoleIDMap 转换为角色ID映射
func (a RoleResources) ToRoleIDMap() map[int]RoleResources {
	m := make(map[int]RoleResources)
	for _, item := range a {
		m[item.RoleID] = append(m[item.RoleID], item)
	}
	return m
}

// ToResourceIDs 转换为资源ID列表
func (a RoleResources) ToResourceIDs() []int {
	idList := make([]int, len(a))
	m := make(map[int]struct{})
	for i, item := range a {
		if _, ok := m[item.ResourceID]; ok {
			continue
		}
		idList[i] = item.ResourceID
		m[item.ResourceID] = struct{}{}
	}
	return idList
}

//// ToMap 转换为map
//func (a RoleMenus) ToMap() map[string]*RoleMenu {
//	m := make(map[string]*RoleMenu)
//	for _, item := range a {
//		m[item.MenuID+"-"+item.ActionID] = item
//	}
//	return m
//}
//
//
//// ToMenuIDs 转换为菜单ID列表
//func (a RoleMenus) ToMenuIDs() []string {
//	var idList []string
//	m := make(map[string]struct{})
//
//	for _, item := range a {
//		if _, ok := m[item.MenuID]; ok {
//			continue
//		}
//		idList = append(idList, item.MenuID)
//		m[item.MenuID] = struct{}{}
//	}
//
//	return idList
//}
//
