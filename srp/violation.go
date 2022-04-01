//go:build violation
// +build violation

package srp

import "database/sql"

type Movie struct {
	name     string
	duration int32
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

func (m *Movie) Create() error {
	db, err := sql.Open("foo", "buzz")

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO movies (name, duration) VALUES($1, $2)", m.name, m.duration)

	return err
}

func (m *Movie) Delete() error {
	db, err := sql.Open("foo", "buzz")

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM movies WHERE name = $1", m.name)

	return err
}
