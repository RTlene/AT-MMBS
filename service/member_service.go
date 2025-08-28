package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// MemberHandler 会员管理接口
func MemberHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取会员列表或单个会员
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个会员
			member, err := dao.Imp.GetMemberByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = member
			}
		} else {
			// 获取会员列表
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
			
			if page <= 0 {
				page = 1
			}
			if pageSize <= 0 {
				pageSize = 10
			}

			members, total, err := dao.Imp.GetMembers(page, pageSize)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = map[string]interface{}{
					"members": members,
					"total":   total,
					"page":    page,
					"size":    pageSize,
				}
			}
		}

	case http.MethodPost:
		// 创建会员
		var member model.MemberModel
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.CreateMember(&member); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = member
			}
		}

	case http.MethodPut:
		// 更新会员
		var member model.MemberModel
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.UpdateMember(&member); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = member
			}
		}

	case http.MethodDelete:
		// 删除会员
		if id := r.URL.Query().Get("id"); id != "" {
			if err := dao.Imp.DeleteMember(id); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "删除成功"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少会员ID参数"
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

// MemberNetworkHandler 会员关系网接口
func MemberNetworkHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodGet {
		res.Code = -1
		res.ErrorMsg = "只支持GET方法"
	} else {
		if id := r.URL.Query().Get("id"); id != "" {
			network, err := dao.Imp.GetMemberNetwork(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = network
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少会员ID参数"
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

// MemberDistributionHandler 会员分销信息接口
func MemberDistributionHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodGet {
		res.Code = -1
		res.ErrorMsg = "只支持GET方法"
	} else {
		if id := r.URL.Query().Get("id"); id != "" {
			distribution, err := dao.Imp.GetMemberDistribution(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = distribution
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少会员ID参数"
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

// MemberConsumptionHandler 会员消费记录接口
func MemberConsumptionHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodGet {
		res.Code = -1
		res.ErrorMsg = "只支持GET方法"
	} else {
		if id := r.URL.Query().Get("id"); id != "" {
			orders, err := dao.Imp.GetMemberConsumption(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = orders
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少会员ID参数"
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
