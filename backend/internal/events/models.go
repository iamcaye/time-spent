package events

type Event struct {
	ID         int
	Name       string
	IdUser     int    `db:"id_user"`
	IdCategory int    `db:"id_category"`
	CreatedAt  string `db:"created_at"`
}
