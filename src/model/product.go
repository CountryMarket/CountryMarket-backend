package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
)

type ProductTab struct {
	Name     string
	Products string
	gorm.Model
}
type ProductHome struct {
	Ids string
}

func (m *model) ProductGetTabList() ([]int, []string, error) {
	var tabs []ProductTab
	err := m.db.Model(&ProductTab{}).Scan(&tabs).Error
	if err != nil {
		return []int{}, []string{}, err
	}
	var ids []int
	var names []string
	for _, v := range tabs {
		ids = append(ids, (int)(v.ID))
		names = append(names, v.Name)
	}
	return ids, names, nil
}
func (m *model) ProductGetTabProducts(tabId int) ([]Product, error) {
	var productTab ProductTab
	err := m.db.Model(&ProductTab{}).Where("id = ?", tabId).Take(&productTab).Error
	if err != nil {
		return []Product{}, err
	}
	productsStringArray := strings.Split(productTab.Products, ", ")
	var products []Product
	for _, v := range productsStringArray {
		tmpInt, err := strconv.Atoi(v)
		if err != nil {
			return []Product{}, err
		}

		product := Product{}
		err = m.db.Model(&Product{}).Where("id = ?", tmpInt).Take(&product).Error
		if err != nil {
			return []Product{}, err
		}

		products = append(products, product)
	}
	return products, nil
}
func (m *model) ProductModifyTabProducts(tabId int, productTab ProductTab) error {
	return m.db.Model(&ProductTab{}).Where("id = ?", tabId).Updates(productTab).Error
}
func (m *model) ProductAddTabProducts(productTab ProductTab) error {
	return m.db.Model(&ProductTab{}).Create(&productTab).Error
}
func (m *model) ProductDeleteTabProducts(tabId int) error {
	return m.db.Delete(&ProductTab{}, tabId).Error
}
func (m *model) ProductGetHomeTab(from, length int) ([]Product, error) {
	var productHome ProductHome
	err := m.db.Model(&ProductHome{}).Take(&productHome).Error
	if err != nil {
		return []Product{}, err
	}
	idsStr := strings.Split(productHome.Ids, ",")
	var products []Product
	for _, v := range idsStr {
		id, err := strconv.Atoi(v)
		if err != nil {
			return []Product{}, err
		}
		product, err := m.ShopGetProduct(id)
		if err != nil {
			return []Product{}, err
		}
		products = append(products, product)
	}
	if from < len(products) && from+length-1 <= len(products)-1 { // 全在前面
		return products[from : from+length], nil
	}
	if from >= len(products) { // 全在后面
		var fp []Product
		err = m.db.Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: "RAND()", WithoutParentheses: true},
		}).Model(&Product{}).Limit(length).Offset(from - len(products)).Scan(&fp).Error
		return fp, nil
	}
	var fp []Product // 两边都有
	err = m.db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RAND()", WithoutParentheses: true},
	}).Model(&Product{}).Limit(length - len(products) + from).Order("id DESC").Scan(&fp).Error
	if err != nil {
		return []Product{}, err
	}
	return append(products[from:], fp...), nil
}
