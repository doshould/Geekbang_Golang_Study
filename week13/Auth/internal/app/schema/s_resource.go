package schema

import (
	"github.com/LyricTian/gin-admin/v6/pkg/util"
)

// Resource 资源对象
type Resource struct {
	ID         	int      	`json:"id"`                                         // 主键ID
	AppKey   	string      `json:"appKey" binding:"required"`                 	// 应用标识
	Name    	string      `json:"name"`                                     	// 资源名称
	Type   		string    	`json:"type"`                                   	// 资源类型
	Feature    	string      `json:"feature"`                                   	// 特征值
	Method   	string    	`json:"method"`                                   	// restful api请求类型
	Pid   		int      	`json:"pid"`                                  		// 父级资源ID
	Sequence   	int         `json:"sequence"`                                   // 排序值
}

func (a *Resource) String() string {
	return util.JSONMarshalToString(a)
}

// ResourceQueryParam 查询条件
type ResourceQueryParam struct {
	PaginationParam
	IDs     	[]int 		`form:"-"`
	AppKey     	string   	`form:"appKey"`
	Name     	string   	`form:"name"`
	Feature 	string  	`form:"feature"`
	Type 		string  	`form:"type"`
}

// ResourceQueryOptions 查询可选参数项
type ResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// ResourceQueryResult 查询结果
type ResourceQueryResult struct {
	Data       Resources
	PageResult *PaginationResult
}

// Resources 资源列表
type Resources []*Resource

func (a Resources) Len() int {
	return len(a)
}

func (a Resources) Less(i, j int) bool {
	return a[i].Sequence > a[j].Sequence
}

func (a Resources) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToMap 转换为键值映射
func (a Resources) ToMap() map[int]*Resource {
	m := make(map[int]*Resource)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

//// SplitParentIDs 拆分父级路径的唯一标识列表
//func (a Resources) SplitParentIDs() []string {
//	idList := make([]string, 0, len(a))
//	mIDList := make(map[string]struct{})
//
//	for _, item := range a {
//		if _, ok := mIDList[item.ID]; ok || item.Feature == "" {
//			continue
//		}
//
//		for _, pp := range strings.Split(item.Feature, "/") {
//			if _, ok := mIDList[pp]; ok {
//				continue
//			}
//			idList = append(idList, pp)
//			mIDList[pp] = struct{}{}
//		}
//	}
//
//	return idList
//}

//// ToTree 转换为菜单树
//func (a Resources) ToTree() MenuTrees {
//	list := make(MenuTrees, len(a))
//	for i, item := range a {
//		list[i] = &MenuTree{
//			ID:         item.ID,
//			Name:       item.Name,
//			Sequence:   item.Sequence,
//		}
//	}
//	return list.ToTree()
//}

//// FillMenuAction 填充菜单动作列表
//func (a Resources) FillMenuAction(mActions map[string]MenuActions) Resources {
//	for _, item := range a {
//		if v, ok := mActions[item.ID]; ok {
//			item.Actions = v
//		}
//	}
//	return a
//}
//
//// ----------------------------------------MenuTree--------------------------------------
//
//// MenuTree 菜单树
//type MenuTree struct {
//	ID         string      `yaml:"-" json:"id"`                                  // 唯一标识
//	Name       string      `yaml:"name" json:"name"`                             // 菜单名称
//	Icon       string      `yaml:"icon" json:"icon"`                             // 菜单图标
//	Router     string      `yaml:"router,omitempty" json:"router"`               // 访问路由
//	ParentID   string      `yaml:"-" json:"parent_id"`                           // 父级ID
//	ParentPath string      `yaml:"-" json:"parent_path"`                         // 父级路径
//	Sequence   int         `yaml:"sequence" json:"sequence"`                     // 排序值
//	ShowStatus int         `yaml:"-" json:"show_status"`                         // 显示状态(1:显示 2:隐藏)
//	Status     int         `yaml:"-" json:"status"`                              // 状态(1:启用 2:禁用)
//	Actions    MenuActions `yaml:"actions,omitempty" json:"actions"`             // 动作列表
//	Children   *MenuTrees  `yaml:"children,omitempty" json:"children,omitempty"` // 子级树
//}
//
//// MenuTrees 菜单树列表
//type MenuTrees []*MenuTree
//
//// ToTree 转换为树形结构
//func (a MenuTrees) ToTree() MenuTrees {
//	mi := make(map[string]*MenuTree)
//	for _, item := range a {
//		mi[item.ID] = item
//	}
//
//	var list MenuTrees
//	for _, item := range a {
//		if item.ParentID == "" {
//			list = append(list, item)
//			continue
//		}
//		if pitem, ok := mi[item.ParentID]; ok {
//			if pitem.Children == nil {
//				children := MenuTrees{item}
//				pitem.Children = &children
//				continue
//			}
//			*pitem.Children = append(*pitem.Children, item)
//		}
//	}
//	return list
//}

// ----------------------------------------MenuAction--------------------------------------

//// MenuAction 菜单动作对象
//type MenuAction struct {
//	ID        string              `yaml:"-" json:"id"`                          // 唯一标识
//	MenuID    string              `yaml:"-" binding:"required" json:"menu_id"`  // 菜单ID
//	Code      string              `yaml:"code" binding:"required" json:"code"`  // 动作编号
//	Name      string              `yaml:"name" binding:"required" json:"name"`  // 动作名称
//	Resources MenuActionResources `yaml:"resources,omitempty" json:"resources"` // 资源列表
//}

//// MenuActionQueryParam 查询条件
//type MenuActionQueryParam struct {
//	PaginationParam
//	MenuID string   // 菜单ID
//	IDs    []string // 唯一标识列表
//}
//
//// MenuActionQueryOptions 查询可选参数项
//type MenuActionQueryOptions struct {
//	OrderFields []*OrderField // 排序字段
//}
//
//// MenuActionQueryResult 查询结果
//type MenuActionQueryResult struct {
//	Data       MenuActions
//	PageResult *PaginationResult
//}
//
//// MenuActions 菜单动作管理列表
//type MenuActions []*MenuAction
//
//// ToMap 转换为map
//func (a MenuActions) ToMap() map[string]*MenuAction {
//	m := make(map[string]*MenuAction)
//	for _, item := range a {
//		m[item.Code] = item
//	}
//	return m
//}
//
//// FillResources 填充资源数据
//func (a MenuActions) FillResources(mResources map[string]MenuActionResources) {
//	for i, item := range a {
//		a[i].Resources = mResources[item.ID]
//	}
//}
//
//// ToMenuIDMap 转换为菜单ID映射
//func (a MenuActions) ToMenuIDMap() map[string]MenuActions {
//	m := make(map[string]MenuActions)
//	for _, item := range a {
//		m[item.MenuID] = append(m[item.MenuID], item)
//	}
//	return m
//}
//
//// ----------------------------------------MenuActionResource--------------------------------------
//
//// MenuActionResource 菜单动作关联资源对象
//type MenuActionResource struct {
//	ID       string `yaml:"-" json:"id"`                             // 唯一标识
//	ActionID string `yaml:"-" json:"action_id"`                      // 菜单动作ID
//	Method   string `yaml:"method" binding:"required" json:"method"` // 资源请求方式(支持正则)
//	Path     string `yaml:"path" binding:"required" json:"path"`     // 资源请求路径（支持/:id匹配）
//}
//
//// MenuActionResourceQueryParam 查询条件
//type MenuActionResourceQueryParam struct {
//	PaginationParam
//	MenuID  string   // 菜单ID
//	MenuIDs []string // 菜单ID列表
//}
//
//// MenuActionResourceQueryOptions 查询可选参数项
//type MenuActionResourceQueryOptions struct {
//	OrderFields []*OrderField // 排序字段
//}
//
//// MenuActionResourceQueryResult 查询结果
//type MenuActionResourceQueryResult struct {
//	Data       MenuActionResources
//	PageResult *PaginationResult
//}
//
//// MenuActionResources 菜单动作关联资源管理列表
//type MenuActionResources []*MenuActionResource
//
//// ToMap 转换为map
//func (a MenuActionResources) ToMap() map[string]*MenuActionResource {
//	m := make(map[string]*MenuActionResource)
//	for _, item := range a {
//		m[item.Method+item.Path] = item
//	}
//	return m
//}
//
//// ToActionIDMap 转换为动作ID映射
//func (a MenuActionResources) ToActionIDMap() map[string]MenuActionResources {
//	m := make(map[string]MenuActionResources)
//	for _, item := range a {
//		m[item.ActionID] = append(m[item.ActionID], item)
//	}
//	return m
//}
