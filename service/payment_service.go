package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// PaymentHandler 支付管理接口
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取支付记录列表或单个记录
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个支付记录
			payment, err := dao.Imp.GetPaymentByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = payment
			}
		} else if orderID := r.URL.Query().Get("order_id"); orderID != "" {
			// 根据订单ID获取支付记录
			payment, err := dao.Imp.GetPaymentByOrderID(orderID)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = payment
			}
		} else {
			// 获取支付记录列表
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
			memberID := r.URL.Query().Get("member_id")
			
			if page <= 0 {
				page = 1
			}
			if pageSize <= 0 {
				pageSize = 10
			}

			payments, total, err := dao.Imp.GetPayments(page, pageSize, memberID)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = map[string]interface{}{
					"payments": payments,
					"total":    total,
					"page":     page,
					"size":     pageSize,
				}
			}
		}

	case http.MethodPost:
		// 创建支付记录
		var payment model.PaymentModel
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.CreatePayment(&payment); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = payment
			}
		}

	case http.MethodPut:
		// 更新支付记录
		var payment model.PaymentModel
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.UpdatePayment(&payment); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = payment
			}
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

// PaymentStatusHandler 支付状态更新接口
func PaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodPut {
		res.Code = -1
		res.ErrorMsg = "只支持PUT方法"
	} else {
		var req struct {
			ID     string `json:"id"`
			Status int    `json:"status"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if req.Status < 0 || req.Status > 4 {
				res.Code = -1
				res.ErrorMsg = "状态值无效，0-4分别表示：待支付、支付中、支付成功、支付失败、已退款"
			} else {
				if err := dao.Imp.UpdatePaymentStatus(req.ID, req.Status); err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
				} else {
					statusText := getPaymentStatusText(req.Status)
					res.Data = fmt.Sprintf("支付状态更新为：%s", statusText)
				}
			}
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

// getPaymentStatusText 获取支付状态文本
func getPaymentStatusText(status int) string {
	statusMap := map[int]string{
		0: "待支付",
		1: "支付中",
		2: "支付成功",
		3: "支付失败",
		4: "已退款",
	}
	return statusMap[status]
}
