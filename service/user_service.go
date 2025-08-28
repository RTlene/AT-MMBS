package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// UserHandler 用户管理接口
func UserHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取用户列表或单个用户
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个用户
			user, err := dao.Imp.GetUserByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = user
			}
		} else {
			// 获取用户列表
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
			if page <= 0 {
				page = 1
			}
			if pageSize <= 0 {
				pageSize = 10
			}

			users, total, err := dao.Imp.GetUsers(page, pageSize)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = map[string]interface{}{
					"users": users,
					"total": total,
					"page":  page,
					"size":  pageSize,
				}
			}
		}

	case http.MethodPost:
		// 创建用户
		var user model.UserModel
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 自动生成ID
			if user.ID == "" {
				user.ID = "user-" + time.Now().Format("20060102150405")
			}
			// 设置默认值
			if user.Role == "" {
				user.Role = "user"
			}
			// 对密码进行MD5加密
			user.Password = HashPassword(user.Password)
			if err := dao.Imp.CreateUser(&user); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Code = 0
				res.Data = user
			}
		}

	case http.MethodPut:
		// 更新用户
		var user model.UserModel
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.UpdateUser(&user); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = user
			}
		}

	case http.MethodDelete:
		// 删除用户
		if id := r.URL.Query().Get("id"); id != "" {
			if err := dao.Imp.DeleteUser(id); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "删除成功"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少用户ID参数"
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
