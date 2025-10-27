package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// CategoryHandler 分类管理接口
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取分类列表或单个分类
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个分类
			category, err := dao.Imp.GetCategoryByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = category
			}
		} else {
			// 获取所有分类
			categories, err := dao.Imp.GetCategories()
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = categories
			}
		}

	case http.MethodPost:
		// 创建分类
		var category model.CategoryModel
		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 自动生成ID
			if category.ID == "" {
				category.ID = "cat-" + time.Now().Format("20060102150405")
			}
			// 设置默认值
			if category.Status == 0 {
				category.Status = 1
			}
			if category.Sort == 0 {
				category.Sort = 999
			}
			
			if err := dao.Imp.CreateCategory(&category); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Code = 0
				res.Data = category
			}
		}

	case http.MethodPut:
		// 更新分类
		var category model.CategoryModel
		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.UpdateCategory(&category); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = category
			}
		}

	case http.MethodDelete:
		// 删除分类
		if id := r.URL.Query().Get("id"); id != "" {
			if err := dao.Imp.DeleteCategory(id); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "删除成功"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少分类ID参数"
		}

	default:
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	// 返回JSON响应
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// ParentCategoryHandler 获取父分类接口
func ParentCategoryHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodGet {
		res.Code = -1
		res.ErrorMsg = "只支持GET方法"
	} else {
		categories, err := dao.Imp.GetParentCategories()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = categories
		}
	}

	// 返回JSON响应
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
