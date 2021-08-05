package builder

type StudentPackage struct {
	stuCourse *StuCourse
	stuClass  *StuClass
	stuPlan   *StuPlan
}

func (s *StudentPackage) GetStuPlan() *StuPlan {
	return s.stuPlan
}

func (s *StudentPackage) SetStuPlan(stuPlan *StuPlan) {
	s.stuPlan = stuPlan
}

func (s *StudentPackage) GetStuCourse() *StuCourse {
	return s.stuCourse
}

func (s *StudentPackage) SetStuCourse(stuCourse *StuCourse) {
	s.stuCourse = stuCourse
}

func (s *StudentPackage) GetStuClass() *StuClass {
	return s.stuClass
}

func (s *StudentPackage) SetStuClass(stuClass *StuClass) {
	s.stuClass = stuClass
}

type Builder interface {
	BuildStuCourse()
	BuildStuClass()
	BuildStuPlan()
	CreateStudentPackage() *StudentPackage
}

type LiveCourseBuilder struct {
	studentPackage *StudentPackage
}

func NewLiveCourseBuilder() *LiveCourseBuilder {
	return &LiveCourseBuilder{
		studentPackage: &StudentPackage{},
	}
}

func (s LiveCourseBuilder) BuildStuCourse() {
	s.studentPackage.SetStuCourse(NewStuCourse())
}

func (s LiveCourseBuilder) BuildStuClass() {
	s.studentPackage.SetStuClass(NewStuClass())
}

func (s LiveCourseBuilder) BuildStuPlan() {
	s.studentPackage.SetStuPlan(NewStuPlan())
}

func (s LiveCourseBuilder) CreateStudentPackage() *StudentPackage {
	return s.studentPackage
}

type RecordCourseBuilder struct {
	studentPackage *StudentPackage
}

func (r RecordCourseBuilder) BuildStuCourse() {
	r.studentPackage.SetStuCourse(NewStuCourse())
}

func (r RecordCourseBuilder) BuildStuClass() {
}

func (r RecordCourseBuilder) BuildStuPlan() {
}

func (r RecordCourseBuilder) CreateStudentPackage() *StudentPackage {
	return r.studentPackage
}

func NewRecordCourseBuilder() *RecordCourseBuilder {
	return &RecordCourseBuilder{
		studentPackage: &StudentPackage{},
	}
}

type StuCourse struct {
}

func NewStuCourse() *StuCourse {
	return &StuCourse{}
}

type StuClass struct {
}

func NewStuClass() *StuClass {
	return &StuClass{}
}

type StuPlan struct {
}

func NewStuPlan() *StuPlan {
	return &StuPlan{}
}

type Director struct {
	StudentPackageBuilder Builder
}

func NewDirector(studentPackageBuilder Builder) *Director {
	return &Director{StudentPackageBuilder: studentPackageBuilder}
}

func (s *Director) construct() *StudentPackage {
	s.StudentPackageBuilder.BuildStuCourse()
	s.StudentPackageBuilder.BuildStuClass()
	s.StudentPackageBuilder.BuildStuPlan()
	return s.StudentPackageBuilder.CreateStudentPackage()
}

type Client struct {
}

func (c *Client) exec() {
	liveCourseBuilder := NewLiveCourseBuilder()
	recordCourseBuilder := NewRecordCourseBuilder()
	c.deliveryStudentPackage(liveCourseBuilder)
	c.deliveryStudentPackage(recordCourseBuilder)
}

func (c *Client) deliveryStudentPackage(builder Builder) {
	director := NewDirector(builder)
	studentPackage := director.construct()
	studentPackage.GetStuCourse()
	studentPackage.GetStuClass()
	studentPackage.GetStuPlan()
}
