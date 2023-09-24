package repository

import "repartners/internal/pack"

// Repository - an interface representing the repository to fetch the various pack sizes. Currently, it is hardcoded, but other storages can be used like database or config. Simply need to implement this interface.
type Repository interface {
	GetAll() ([]pack.PackSize, error)
}
