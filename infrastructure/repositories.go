package infrastructure

import "github.com/1azar/QRChan/domain"

func NewMemoryDBRepo() *MemoryDBRepo {
	return &MemoryDBRepo{
		Db: make(map[int64]domain.QRSettings),
	}
}

type MemoryDBRepo struct {
	Db map[int64]domain.QRSettings
}

func (m *MemoryDBRepo) Store(settings domain.QRSettings) error {
	m.Db[settings.ID] = settings
	return nil
}

func (m *MemoryDBRepo) FindById(id int64) (domain.QRSettings, error) {
	return m.Db[id], nil
}
