package store

import (
    "errors"
	"context"

    "github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

var ErrNotFound = errors.New("not found")

type Store struct {
    Pool *pgxpool.Pool
}

func (s *Store) MemberByEmail(ctx context.Context, email string) (Member, error) {
    var m Member
    err := s.Pool.QueryRow(ctx,
        `SELECT id, email, display_name, role, password_hash
         FROM members WHERE lower(email) = $1`, email).
        Scan(&m.ID, &m.Email, &m.DisplayName, &m.Role, &m.PasswordHash)
    if errors.Is(err, pgx.ErrNoRows) {
        return Member{}, ErrNotFound
    }
    return m, err
}