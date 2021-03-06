package postgres

import (
	"DBsummer/configDir"
	"context"
	"errors"
	"fmt"
	"log"
)

//go:generate dbx golang -d postgres -p postgres ../dbxSubdir/db.dbx .
//go:generate dbx schema -d postgres ../dbxSubdir/db.dbx ../dbxSubdir

const postgres = "postgres"

type Repository struct {
	db      *DB
	tx      *Tx
	methods Methods

	Log    *log.Logger
	Config *configDir.DBConfig

	tableNew     TableNewRepository
	subjects     SubjectsRepository
	groups       GroupsRepository
	students     StudentRepository
	teachers     TeachersRepository
	sheets       SheetsRepository
	runner_marks RunnerMarksRepository
	runner       RunnersRepository
	sheet_marks  SheetMarksRepository
	//tests TestsRepository
	//vidomosti VidomostiRepository
}

func New(log *log.Logger, config configDir.DBConfig) (*Repository, error) {
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s%s", postgres, config.UserName, config.Password, config.Host, config.Port, config.DBName, config.SSL)

	store, err := Open(postgres, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening database: %v", err)
	}
	log.Println(fmt.Sprintf("Connected to: %s %s", "db source", dbURL))

	serverDB := &Repository{
		Log:     log,
		Config:  &config,
		db:      store,
		methods: store,
	}

	serverDB.tableNew = TableNewRepository{serverDB}
	serverDB.subjects = SubjectsRepository{serverDB}
	serverDB.groups = GroupsRepository{serverDB}
	serverDB.students = StudentRepository{serverDB}
	serverDB.teachers = TeachersRepository{serverDB}
	serverDB.sheets = SheetsRepository{serverDB}
	serverDB.runner_marks = RunnerMarksRepository{serverDB}
	serverDB.runner = RunnersRepository{serverDB}
	serverDB.sheet_marks = SheetMarksRepository{serverDB}
	//serverDB.tests = TestsRepository{serverDB}
	//serverDB.vidomosti = VidomostiRepository{serverDB}

	return serverDB, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) TableNew() TableNewRepository {
	return r.tableNew
}

func (r *Repository) Subjects() SubjectsRepository {
	return r.subjects
}

func (r *Repository) Groups() GroupsRepository {
	return r.groups
}

func (r *Repository) Students() StudentRepository {
	return r.students
}

func (r *Repository) Teachers() TeachersRepository {
	return r.teachers
}

func (r *Repository) Sheets() SheetsRepository {
	return r.sheets
}

func (r *Repository) RunnerMarks() RunnerMarksRepository {
	return r.runner_marks
}

func (r *Repository) Runner() RunnersRepository {
	return r.runner
}

func (r *Repository) SheetMarks() SheetMarksRepository {
	return r.sheet_marks
}

func (r *Repository) BeginTx(ctx context.Context) (*Repository, error) {
	if r.db == nil {
		return nil, errors.New("db is not initialized")
	}

	tx, err := r.db.Open(ctx)
	if err != nil {
		return nil, err
	}

	ptx := *r

	ptx.tx = tx
	ptx.methods = tx

	return &ptx, nil
}

func (r *Repository) Commit() error {
	if r.tx == nil {
		return errors.New("begin transaction before commit it")
	}

	return r.tx.Commit()
}

func (r *Repository) Rollback() error {
	if r.tx == nil {
		return errors.New("begin transaction before rollback it")
	}

	return r.tx.Rollback()
}
