package repository

import (
    "database/sql"
    "server/domain/model"
)

type UserRepository interface {
    Insert(DB *sql.DB, userID, name, email string) error
    GetByUserID(DB *sql.DB, userID string) ([]*domain.User, error)
}