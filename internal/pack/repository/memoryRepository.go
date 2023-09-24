package repository

import "repartners/internal/pack"

type MemoryRepository struct{}

func NewMemoryRepo() MemoryRepository {
	return MemoryRepository{}
}

func (repo MemoryRepository) GetAll() ([]pack.PackSize, error) {
	packs := []int{250, 500, 1000, 2000, 5000}

	packList := make([]pack.PackSize, len(packs), len(packs))

	for key, packSize := range packs {
		packList[key] = pack.NewPack(packSize)
	}

	return packList, nil
}
