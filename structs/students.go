package structs

type Student struct {
	StudentCipher    string
	FirstName        string
	LastName         string
	MiddleName       string
	RecordBookNumber string
}

type StudentPIB struct {
	StudentCipher string
	Pib           string
	RecordNumber  string
}

type StudentMarks struct {
	SubjectName  string
	MarkTogether int
	EctsMark     string
	TeacherPIB   string
}

type StudentAllMarks struct {
	SubjectName     string
	SheetMark       int
	SheetID         int
	RunnerMark      string
	RunnerID        string
	Semester        string
	EducationalYear string
}
