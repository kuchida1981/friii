package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/kuchida1981/friii/internal/domain/entity"
)

type PartnerRepository interface {
	Save(partner *entity.Partner) error
	FindByID(id uuid.UUID, at time.Time) (*entity.Partner, error)
	List(at time.Time) ([]*entity.Partner, error)
}
