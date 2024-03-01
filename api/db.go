package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}
	if err := s.createTasksTable(); err != nil {
		return nil, err
	}
	return s.db, nil
}

func (s *MySQLStorage) createProjectsTable() error {
	log.Println("creating projects table")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
		    id int unsigned not null auto_increment,
		    name varchar(255) not null,
		    createAt timestamp not null default current_timestamp,
		    
			primary key (id)
		) engine=InnoDB default charset=utf8;
	`)
	return err
}

func (s *MySQLStorage) createTasksTable() error {
	log.Println("creating tasks table")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
		    id int unsigned not null auto_increment,
		    name varchar(255) not null,
		    status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') not null default 'TODO',
		    projectId int unsigned not null,
		    assignedToID int unsigned not null,		    
		    createAt timestamp not null default current_timestamp,
		    
			primary key (id),
		    foreign key (assignedToID) references users(id),
		    foreign key (projectId) references projects(id)
		) engine=InnoDB default charset=utf8;
	`)
	return err
}

func (s *MySQLStorage) createUsersTable() error {
	log.Println("creating users table")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		    id int unsigned not null auto_increment,
		    firstname varchar(255) not null,
		    lastname varchar(255) not null,
			password varchar(255) not null,		    
		    createAt timestamp not null default current_timestamp,
		    
			primary key (id)
		) engine=InnoDB default charset=utf8;
	`)
	return err
}
