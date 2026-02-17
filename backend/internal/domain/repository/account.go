package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/kuchida1981/friii/internal/domain/entity"
)

type AccountRepository interface {
	SaveItem(item *entity.AccountItem) error
	FindItemByID(id uuid.UUID, at time.Time) (*entity.AccountItem, error)
	ListItem(at time.Time) ([]*entity.AccountItem, error)
	FindCategoryByID(id uuid.UUID) (*entity.AccountCategory, error)
	ListCategories() ([]*entity.AccountCategory, error)
}
