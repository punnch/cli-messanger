package rooms_postgres_repository

import core_postgres_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool"

type RoomsRepository struct {
	pool core_postgres_pool.Pool
}

func NewRoomsRepository(
	pool core_postgres_pool.Pool,
) *RoomsRepository {
	return &RoomsRepository{
		pool: pool,
	}
}
