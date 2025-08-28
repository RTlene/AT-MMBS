package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// MemberLevelHandler 会员等级管理接口
func MemberLevelHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取会员等级列表或单个等级
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个等级
			level, err := dao.Imp.GetMemberLevelByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = level
			}
		} else {
			// 获取所有等级
			levels, err := dao.Imp.GetMemberLevels()
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = levels
			}
		}

	case http.MethodPost:
		// 创建会员等级
		var level model.MemberLevelModel
		if err := json.NewDecoder(r.Body).Decode(&level); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.CreateMemberLevel(&level); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = level
			}
		}

	case http.MethodPut:
		// 更新会员等级
		var level model.MemberLevelModel
		if err := json.NewDecoder(r.Body).Decode(&level); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.UpdateMemberLevel(&level); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = level
			}
		}

	case http.MethodDelete:
		// 删除会员等级
		if id := r.URL.Query().Get("id"); id != "" {
			if err := dao.Imp.DeleteMemberLevel(id); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "删除成功"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少等级ID参数"
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
