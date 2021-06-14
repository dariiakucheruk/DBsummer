package dataBaseDir

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
)

type Repository interface {
	Close() error
	BeginTx(ctx context.Context) (Transaction, error)
	TableNew() TableNewInterface
	Subjects() SubjectsInterface
	Groups() GroupsInterface
	Students() StudentsInterface
	Teachers() TeachersInterface
	Sheets() SheetsInterface
	RunnerMarks() RunnerMarksInterface
	Runners() RunnerInterface
	SheetMarks() SheetMarksInterface
}

type Transaction interface {
	Repository

	Commit() error
	Rollback() error
}

type TableNewInterface interface { //functions to be used for the table 'tablen'
	Get(ctx context.Context, id int) error
	Create(ctx context.Context) (id int, err error)
}

type SubjectsInterface interface { //functions to be used for the table 'subjects'
	Get(ctx context.Context, id int) (*structs.Subject, error)
	GetAll(ctx context.Context) ([]*structs.Subject, error)
	Create(ctx context.Context) (id int, err error)
	FindSubject(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error)
	AddSubject(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error)
	FindSubjectsWithYearParameter(ctx context.Context, year int) ([]*structs.SubjectName, error)
}

type GroupsInterface interface {
	FindGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error)
	AddGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error)
}

type StudentsInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context, id int) (*structs.Student, error)
	GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error)
	GetAllBorjniki(ctx context.Context) ([]*structs.AllStudInfo, error)
	GetStudentByPIB(ctx context.Context, fn string, ln string, mn string, year string) ([]*structs.StudentMarks, error)
	FindStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error)
	AddStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error)
}

type TeachersInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
	FindTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error)
	AddTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error)
	GetTeacherPassStatistics(ctx context.Context, passedOrNot string) ([]*structs.TeacherPassStatistics, error)
}

type SheetsInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
	GetSheetFromParameters(ctx context.Context, fn string, ln string, mn string, subj string, gr string, year string) ([]*structs.SheetByQuery, error)
	PostSheetToDataBase(ctx context.Context, sheet *pdfReading.ExtractedInformation, teacherId int, groupId int) (*pdfReading.ExtractedInformation, error)
}

type RunnerMarksInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type RunnerInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type SheetMarksInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
	PostSheetMarksToDataBase(ctx context.Context, sheetID int, studentId int, sheetMarks *pdfReading.StudInfoFromPDF) (int, error)
}
