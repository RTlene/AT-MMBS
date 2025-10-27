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

// ProductHandler 商品管理接口
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.Method {
	case http.MethodGet:
		// 获取商品列表或单个商品
		if id := r.URL.Query().Get("id"); id != "" {
			// 获取单个商品
			product, err := dao.Imp.GetProductByID(id)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = product
			}
		} else {
			// 获取商品列表
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
			category := r.URL.Query().Get("category")
			
			if page <= 0 {
				page = 1
			}
			if pageSize <= 0 {
				pageSize = 10
			}

			products, total, err := dao.Imp.GetProducts(page, pageSize, category)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = map[string]interface{}{
					"products": products,
					"total":    total,
					"page":     page,
					"size":     pageSize,
				}
			}
		}

	case http.MethodPost:
		// 创建商品
		var product model.ProductModel
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 自动生成ID
			if product.ID == "" {
				product.ID = "prod-" + time.Now().Format("20060102150405")
			}
			// 设置默认值
			if product.Status == 0 {
				product.Status = 1
			}
			if product.Kucun == 0 {
				product.Kucun = 0
			}
			if product.Price == 0 {
				product.Price = 0.01
			}
			// 确保Tp字段有默认值
			if product.Tp == "" {
				product.Tp = "[]"
			}
			// 设置时间字段
			now := time.Now()
			if product.CreateTime.IsZero() {
				product.CreateTime = now
			}
			if product.UpdateTime.IsZero() {
				product.UpdateTime = now
			}
			
			// 添加调试日志
			fmt.Printf("准备保存商品，Tp字段值: %+v, 类型: %T\n", product.Tp, product.Tp)
			
			if err := dao.Imp.CreateProduct(&product); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Code = 0
				res.Data = product
			}
		}

	case http.MethodPut:
		// 更新商品
		var product model.ProductModel
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 设置更新时间
			product.UpdateTime = time.Now()
			
			if err := dao.Imp.UpdateProduct(&product); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = product
			}
		}

	case http.MethodDelete:
		// 删除商品
		if id := r.URL.Query().Get("id"); id != "" {
			if err := dao.Imp.DeleteProduct(id); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = "删除成功"
			}
		} else {
			res.Code = -1
			res.ErrorMsg = "缺少商品ID参数"
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

// ProductStatusHandler 商品上下架接口
func ProductStatusHandler(w http.ResponseWriter, r *http.Request) {
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
			if req.Status != 0 && req.Status != 1 {
				res.Code = -1
				res.ErrorMsg = "状态值无效，0表示下架，1表示上架"
			} else {
				if err := dao.Imp.UpdateProductStatus(req.ID, req.Status); err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
				} else {
					// 获取更新后的商品信息
					updatedProduct, err := dao.Imp.GetProductByID(req.ID)
					if err != nil {
						res.Code = -1
						res.ErrorMsg = "状态更新成功，但获取更新后信息失败"
					} else {
						res.Data = updatedProduct
					}
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
