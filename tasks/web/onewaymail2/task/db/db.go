package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var Instance *pgxpool.Pool
