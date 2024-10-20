package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetProjects() ([]*Project, error)
	GetProjectById(int) (*Project, error)
	CreateProject(*Project) error
	UpdateProject(int) error
	DeleteProject(int) error
}

type PostgresStore struct {
	db *sql.DB
}

// TODO: add to .env file
const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "featureflagdb"
)

func NewPostgressStore() (*PostgresStore, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {

		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createProjectTable()
}

func (s *PostgresStore) createProjectTable() error {
	query := `create table if not exists project (
		id serial primary key,
		description varchar(50),
		created_at timestamp,
		is_active boolean
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetProjects() ([]*Project, error) {
	rows, err := s.db.Query("select * from project")
	if err != nil {
		return nil, err
	}

	projects := []*Project{}

	for rows.Next() {
		project := &Project{}
		err := rows.Scan(
			&project.ID,
			&project.Description,
			&project.CreatedAt,
			&project.IsActive,
		)

		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}
	return projects, nil
}

func (s *PostgresStore) GetProjectById(id int) (*Project, error) {
	return nil, nil
}

func (s *PostgresStore) CreateProject(project *Project) error {
	query := `insert into 
	project (description, created_at, is_active) 
	values ($1, $2, $3)`

	_, err := s.db.Query(
		query,
		project.Description,
		project.CreatedAt,
		project.IsActive,
	)
	if err != nil {
		return nil
	}

	return nil
}

func (s *PostgresStore) UpdateProject(id int) error {
	return nil
}

func (s *PostgresStore) DeleteProject(id int) error {
	return nil
}
