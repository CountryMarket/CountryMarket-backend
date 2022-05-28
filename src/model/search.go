package model

import "fmt"

func (m *model) Search(key string) ([]Product, error) {
	var res []Product
	err := m.db.Model(Product{}).Where("title LIKE ?", fmt.Sprintf("%%%s%%", key)).Scan(&res).Error
	if err != nil {
		return []Product{}, err
	}
	return res, nil
}
