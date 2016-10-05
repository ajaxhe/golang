package db

import (
    "github.com/ajaxhe/thelittlegobook/ch04/shopping/models"
)

func LoadItem(id int) *models.Item {
    return &models.Item{
        Price: 9.001,
    }
}
