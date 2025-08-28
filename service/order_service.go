package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// OrderHandler 订单管理接口
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取订单列表或单个订单
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个订单
			order, err := dao.Imp.GetOrderByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = order
			}
		} else if orderNo := r.URL.Query().Get("order_no"); orderNo != "" {
			// 根据订单号获取订单
			order, err := dao.Imp.GetOrderByOrderNo(orderNo)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = order
			}
		} else {
			// 获取订单列表
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
			memberID := r.URL.Query().Get("member_id")
			
			if page <= 0 {
				page = 1
			}
			if pageSize <= 0 {
				pageSize = 10
			}

			orders, total, err := dao.Imp.GetOrders(page, pageSize, memberID)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = map[string]interface{}{
					"orders": orders,
					"total":  total,
					"page":   page,
					"size":   pageSize,
				}
			}
		}

	case http.MethodPost:
		// 创建订单
		var order model.OrderModel
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			if err := dao.Imp.CreateOrder(&order); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = order
			}
		}

	case http.MethodPut:
		// 更新订单
		var order model.OrderModel
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 检查订单是否可以修改
			if err := validateOrderUpdate(&order); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				if err := dao.Imp.UpdateOrder(&order); err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
				} else {
					res.Data = order
				}
			}
		}

	case http.MethodDelete:
		// 废除订单（软删除）
		if id := r.URL.Query().Get("id"); id != "" {
			// 这里可以实现软删除逻辑，将订单状态改为已取消
			if err := dao.Imp.UpdateOrderStatus(id, 4); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "订单已废除"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少订单ID参数"
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

// OrderStatusHandler 订单状态更新接口
func OrderStatusHandler(w http.ResponseWriter, r *http.Request) {
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
				res.ErrorMsg = "状态值无效，0-4分别表示：待确认、已确认、已发货、已完成、已取消"
			} else {
				if err := dao.Imp.UpdateOrderStatus(req.ID, req.Status); err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
				} else {
					statusText := getOrderStatusText(req.Status)
					res.Data = fmt.Sprintf("订单状态更新为：%s", statusText)
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

// validateOrderUpdate 验证订单是否可以修改
func validateOrderUpdate(order *model.OrderModel) error {
	// 获取原订单信息
	originalOrder, err := dao.Imp.GetOrderByID(order.ID)
	if err != nil {
		return err
	}

	// 已支付的订单不能修改订购商品
	if originalOrder.PaymentStatus == 1 && order.Products != nil {
		return fmt.Errorf("已支付的订单不能修改订购商品")
	}

	// 订单号、下单会员、支付状态不能修改
	if order.OrderNo != originalOrder.OrderNo {
		return fmt.Errorf("订单号不能修改")
	}
	if order.MemberID != originalOrder.MemberID {
		return fmt.Errorf("下单会员不能修改")
	}
	if order.PaymentStatus != originalOrder.PaymentStatus {
		return fmt.Errorf("支付状态不能修改")
	}

	return nil
}

// getOrderStatusText 获取订单状态文本
func getOrderStatusText(status int) string {
	statusMap := map[int]string{
		0: "待确认",
		1: "已确认",
		2: "已发货",
		3: "已完成",
		4: "已取消",
	}
	return statusMap[status]
}
