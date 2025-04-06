//nolint:all
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32              `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
