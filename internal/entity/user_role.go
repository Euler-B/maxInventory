package entity

type USER_ROLE struct {
	ID     int64 `db:"id"`
	UserID int64 `db:"user_id"`
	RoleID int64 `db:"role_id"`
}
