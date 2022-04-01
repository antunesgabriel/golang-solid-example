package srp

import "database/sql"

type Movie struct {
	name     string
	duration int32
}

type MovieRepository struct {
	db *sql.DB
}

func (m *Movie) Name() string {
	return m.name
}

func (m *Movie) SetName(name string) {
	m.name = name
}

func (m *Movie) Duration() int32 {
	return m.duration
}

func (m *Movie) SetDuration(duration int32) {
	m.duration = duration
}

func (repo *MovieRepository) Create(movie *Movie) error {
	_, err := repo.db.Exec("INSERT INTO movies (name, duration) VALUES($1, $2)", movie.name, movie.duration)

	return err
}

func (repo *MovieRepository) Delete(movie *Movie) error {
	_, err := repo.db.Exec("DELETE FROM movies WHERE name = $1", movie.name)

	return err
}
