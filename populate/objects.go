package populate

import (
	"../sifxml"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/au"
)

func create_address(state string) sifxml.AddressType {
	gofakeit.Seed(0)
	if state == "" {
		state = create_state()
	}
	postcode := postcode_seeded(state)

	ret := sifxml.AddressType{}
	ret.SetProperty("Type", "0123")
	ret.SetProperty("Role", "012A")
	ret.Street().SetProperty("StreetNumber", gofakeit.StreetNumber())
	ret.Street().SetProperty("StreetName", gofakeit.StreetName())
	ret.Street().SetProperty("StreetType", randomStringFromSlice([]string{"Avenue", "Boulevard", "Cove", "Court", "Crescent", "Drive", "Esplanade", "Lane", "Place", "Road", "Square", "Street", "Terrace", "Walk", "Way"}))
	ret.SetProperty("City", gofakeit.City())
	ret.SetProperty("StateProvince", state)
	ret.SetProperty("PostalCode", postcode)
	return ret
}

func addTeachers(t *sifxml.ScheduledTeacherListType, staff []*sifxml.StaffPersonal, start time.Time, finish time.Time) *sifxml.ScheduledTeacherListType {
	for _, s := range staff {
		t = t.AddNew()
		t.Last().SetProperty("StaffPersonalRefId", s.RefId().String())
		t.Last().SetProperty("StaffLocalId", s.LocalId().String())
		t.Last().SetProperty("StartTime", start.Format("15:04:05"))
		t.Last().SetProperty("FinishTime", finish.Format("15:04:05"))
		t.Last().SetProperty("Weighting", 1.0)
		if rand.Float64() > 0.6 {
			t.Last().SetProperty("Credit", threshold_rand_strings([]float64{0.75, 0.5, 0.25, 0}, []string{"Casual", "Extra", "In-Lieu", "Underload"}))
		}
		t.Last().SetProperty("Supervision", threshold_rand_strings([]float64{0.9, 0.8, 0}, []string{"MergedClass", "MinimalSupervision", "Normal"}))
	}
	return t
}

func copyTeachingGroupPeriodFromCell(group *sifxml.TeachingGroup, cell *sifxml.TimeTableCell, start_times map[string]map[string]string) *sifxml.TeachingGroup {
	group.TeachingGroupPeriodList().AddNew()
	group.TeachingGroupPeriodList().Last().SetProperty("TimeTableCellRefId", cell.RefId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("RoomNumber", cell.RoomNumber().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StaffLocalId", cell.StaffLocalId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("DayId", cell.DayId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("PeriodId", cell.PeriodId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StartTime", cell.RefId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("CellType", cell.CellType().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StartTime", start_times[cell.DayId().String()][cell.PeriodId().String()])
	return group
}

/* if yearlevel = "", create any year level */
func Create_StudentPersonal(yearlevel string) *sifxml.StudentPersonal {
	if yearlevel == "" {
		yearlevel = strconv.Itoa(rand.Intn(12) + 1)
	}
	gofakeit.Seed(0)
	person := gofakeit.Person()
	middlename := gofakeit.FirstName()
	hasParent1 := rand.Float64() < 0.8
	hasParent2 := rand.Float64() < 0.8
	if !hasParent1 {
		hasParent2 = true
	}
	birthyr, birthyr_err := birth_year(yearlevel)

	ret := sifxml.StudentPersonal{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.PersonInfo().Name().SetProperty("Type", "LGL")
	ret.PersonInfo().Name().SetProperty("FamilyName", person.LastName)
	ret.PersonInfo().Name().SetProperty("GivenName", person.FirstName)
	ret.PersonInfo().Name().SetProperty("MiddleName", middlename)
	ret.PersonInfo().Demographics().SetProperty("Sex", sex_seeded(person.Gender))
	if birthyr_err == nil {
		ret.PersonInfo().Demographics().SetProperty("BirthDate", birthyr)
	}
	ret.PersonInfo().Demographics().SetProperty("IndigenousStatus", threshold_rand_strings([]float64{0.2, 0.15, 0.1, 0.5, 0}, []string{"4", "1", "2", "3", "9"}))
	ret.PersonInfo().Demographics().SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
	if hasParent1 {
		ret.MostRecent().SetProperty("Parent1Language", "1201")
		ret.MostRecent().SetProperty("Parent1EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent().SetProperty("Parent1SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent().SetProperty("Parent1NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if hasParent2 {
		ret.MostRecent().SetProperty("Parent2Language", "1201")
		ret.MostRecent().SetProperty("Parent2EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent().SetProperty("Parent2SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent().SetProperty("Parent2NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	ret.MostRecent().YearLevel().SetProperty("Code", yearlevel)
	if out, ok := sifxml.StudentPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_StudentPersonals(count int, yearlevels []string) []*sifxml.StudentPersonal {
	ret := sifxml.StudentPersonalSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StudentPersonal(randomStringFromSlice(yearlevels)))
	}
	return ret
}

/* schooltype is AUCodeSetsSchoolLevelType; if "", will be created */
func Create_SchoolInfo(schooltype string) *sifxml.SchoolInfo {
	gofakeit.Seed(0)
	local_id := strconv.Itoa(seq_gen("localId"))
	state := create_state()
	if schooltype == "" {
		schooltype = randomStringFromSlice(sifxml.AUCodeSetsSchoolLevelType_values)
	}

	ret := sifxml.SchoolInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", local_id)
	ret.SetProperty("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 9999)+1))
	ret.SetProperty("CommonwealthId", strconv.Itoa(random_seq_gen("schoolCommonwealthId", 9999)+1))
	ret.SetProperty("SchoolName", school_name(schooltype))
	ret.Campus().SetProperty("ParentSchoolId", local_id)
	ret.Campus().SetProperty("SchoolCampusId", strconv.Itoa(rand.Intn(4)+1))
	ret.Campus().SetProperty("AdminStatus", threshold_rand_strings([]float64{0.8, 0}, []string{"N", "Y"}))
	ret.Campus().SetProperty("CampusType", schooltype)
	ret.SetProperty("SchoolSector", "Gov")
	ret.SetProperty("OperationalStatus", "O")
	ret.SetProperty("IndependentSchool", "N")
	ret.SetProperty("SchoolType", schooltype)
	ret.SetProperty("ARIA", 1.0)
	ret.SetProperty("Entity_Open", "1990-01-01")
	ret.AddressList().Append(create_address(state))
	if out, ok := sifxml.SchoolInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SchoolInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_SchoolInfos(count int) []*sifxml.SchoolInfo {
	ret := sifxml.SchoolInfoSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_SchoolInfo(""))
	}
	return ret
}

func Create_StudentSchoolEnrollment(student *sifxml.StudentPersonal, school *sifxml.SchoolInfo) *sifxml.StudentSchoolEnrollment {
	ret := sifxml.StudentSchoolEnrollment{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("MembershipType", "01")
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("TimeFrame", "C")
	ret.SetProperty("FTE", 1.0)
	ret.SetProperty("EntryDate", this_year()+"-01-25")

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if student != nil {
		ret.SetProperty("StudentPersonalRefId", student.RefId().String())
	}

	if student != nil && !student.MostRecent_IsNil() && !student.MostRecent().YearLevel_IsNil() {
		ret.YearLevel().SetProperty("Code", student.MostRecent().YearLevel().Code().String())
	} else {
		ret.YearLevel().SetProperty("Code", strconv.Itoa(rand.Intn(12)+1))
	}
	if school != nil && student != nil {
		student.PersonInfo().EmailList().Last().SetProperty("Value", create_email(student.PersonInfo().Name().GivenName().String(), student.PersonInfo().Name().MiddleName().String(), student.PersonInfo().Name().FamilyName().String(), create_school_email_domain(school)))
	}

	if out, ok := sifxml.StudentSchoolEnrollmentPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentSchoolEnrollment: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_StudentSchoolEnrollments(students []*sifxml.StudentPersonal, school *sifxml.SchoolInfo) []*sifxml.StudentSchoolEnrollment {
	ret := sifxml.StudentSchoolEnrollmentSlice()
	for _, s := range students {
		ret = append(ret, Create_StudentSchoolEnrollment(s, school))
	}
	return ret
}

func Create_StaffPersonal() *sifxml.StaffPersonal {
	gofakeit.Seed(0)
	person := gofakeit.Person()
	middlename := gofakeit.FirstName()

	ret := sifxml.StaffPersonal{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.SetProperty("EmploymentStatus", "A")
	ret.OtherIdList().AddNew()
	ret.OtherIdList().Last().SetProperty("Type", "DET_USER_ID")
	ret.OtherIdList().Last().SetProperty("Value", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.OtherIdList().AddNew()
	ret.OtherIdList().Last().SetProperty("Type", threshold_rand_strings([]float64{0.1, 0}, []string{"pep", "cep"}))
	ret.OtherIdList().Last().SetProperty("Value", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.PersonInfo().Name().SetProperty("Type", "LGL")
	ret.PersonInfo().Name().SetProperty("FamilyName", person.LastName)
	ret.PersonInfo().Name().SetProperty("GivenName", person.FirstName)
	ret.PersonInfo().Name().SetProperty("PreferredGivenName", person.FirstName)
	ret.PersonInfo().Name().SetProperty("MiddleName", middlename)
	ret.PersonInfo().Name().SetProperty("Title", create_salutation(person.Gender))
	ret.PersonInfo().Demographics().SetProperty("Sex", sex_seeded(person.Gender))
	ret.PersonInfo().Demographics().SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo().Demographics().SetProperty("BirthDate", random_date(strconv.Itoa(time.Now().Year()-65)+"-01-01", strconv.Itoa(time.Now().Year()-25)+"-12-31"))
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
	if out, ok := sifxml.StaffPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StaffPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_StaffPersonals(count int) []*sifxml.StaffPersonal {
	ret := sifxml.StaffPersonalSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StaffPersonal())
	}
	return ret
}

func Create_StaffAssignment(staff *sifxml.StaffPersonal, school *sifxml.SchoolInfo) *sifxml.StaffAssignment {
	ret := sifxml.StaffAssignment{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("PrimaryAssignment", "Y")
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("JobStartDate", "1990-01-01")
	ret.SetProperty("JobFunction", "teacher")
	ret.StaffActivity().SetProperty("Code", randomStringFromSlice(sifxml.AUCodeSetsStaffActivityType_values))

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if staff != nil {
		ret.SetProperty("StaffPersonalRefId", staff.RefId().String())
	}
	if school != nil && staff != nil {
		staff.PersonInfo().EmailList().Last().SetProperty("Value", create_email(staff.PersonInfo().Name().GivenName().String(), staff.PersonInfo().Name().MiddleName().String(), staff.PersonInfo().Name().FamilyName().String(), create_school_email_domain(school)))
	}
	if out, ok := sifxml.StaffAssignmentPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StaffAssignment: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_StaffAssignments(staff []*sifxml.StaffPersonal, school *sifxml.SchoolInfo) []*sifxml.StaffAssignment {
	ret := sifxml.StaffAssignmentSlice()
	for _, s := range staff {
		ret = append(ret, Create_StaffAssignment(s, school))
	}
	return ret
}

func Create_StudentContactPersonal(student *sifxml.StudentPersonal, ordinal int) *sifxml.StudentContactPersonal {
	if student != nil && student.MostRecent().Parent1Language_IsNil() && ordinal == 1 {
		return nil
	}
	if student != nil && student.MostRecent().Parent2Language_IsNil() && ordinal == 2 {
		return nil
	}
	state := create_state()
	if student != nil && !student.PersonInfo().AddressList_IsNil() {
		state = student.PersonInfo().AddressList().Last().StateProvince().String()
	}

	gofakeit.Seed(0)
	person := gofakeit.Person()
	middlename := gofakeit.FirstName()

	ret := sifxml.StudentContactPersonal{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.PersonInfo().Name().SetProperty("Type", "LGL")
	ret.PersonInfo().Name().SetProperty("FamilyName", person.LastName)
	ret.PersonInfo().Name().SetProperty("GivenName", person.FirstName)
	ret.PersonInfo().Name().SetProperty("PreferredGivenName", person.FirstName)
	ret.PersonInfo().Name().SetProperty("PreferredFamilyName", person.LastName)
	ret.PersonInfo().Name().SetProperty("MiddleName", gofakeit.FirstName())
	ret.PersonInfo().Name().SetProperty("Title", create_salutation(person.Gender))
	ret.PersonInfo().Demographics().SetProperty("Sex", sex_seeded(person.Gender))
	ret.PersonInfo().Demographics().SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo().Demographics().LanguageList().AddNew()
	ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("Code", "1201")
	ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("LanguageType", "1")
	if rand.Float64() < 0.2 {
		ret.PersonInfo().Demographics().LanguageList().AddNew()
		ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("Code", randomStringFromSlice([]string{"0002", "7101", "2401", "2201", "5203", "4202"}))
		ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("LanguageType", "2")
	}
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, create_commercial_email_domain()))
	ret.PersonInfo().PhoneNumberList().AddNew()
	ret.PersonInfo().PhoneNumberList().Last().SetProperty("Type", "0096")
	ret.PersonInfo().PhoneNumberList().Last().SetProperty("Number", create_phone_number(nil))
	ret.PersonInfo().AddressList().Append(create_address(state))
	if student != nil && ordinal == 1 {
		ret.SetProperty("EmploymentType", student.MostRecent().Parent1EmploymentType().String())
		ret.SetProperty("SchoolEducationalLevel", student.MostRecent().Parent1SchoolEducationLevel().String())
		ret.SetProperty("NonSchoolEducation", student.MostRecent().Parent1NonSchoolEducation().String())
	} else if student != nil && ordinal == 2 {
		ret.SetProperty("EmploymentType", student.MostRecent().Parent2EmploymentType().String())
		ret.SetProperty("SchoolEducationalLevel", student.MostRecent().Parent2SchoolEducationLevel().String())
		ret.SetProperty("NonSchoolEducation", student.MostRecent().Parent2NonSchoolEducation().String())
	} else {
		ret.SetProperty("EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.SetProperty("SchoolEducationalLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.SetProperty("NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if student != nil && rand.Float64() < 0.8 {
		ret.PersonInfo().Name().SetProperty("FamilyName", student.PersonInfo().Name().FamilyName().String())
		ret.PersonInfo().Name().SetProperty("PreferredFamilyName", student.PersonInfo().Name().FamilyName().String())
	}
	if out, ok := sifxml.StudentContactPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentContactPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_StudentContactRelationship(student *sifxml.StudentPersonal, contact *sifxml.StudentContactPersonal) *sifxml.StudentContactRelationship {
	ret := sifxml.StudentContactRelationship{}
	ret.SetProperty("StudentContactRelationshipRefId", create_GUID())
	ret.Relationship().SetProperty("Code", threshold_rand_strings([]float64{0.26, 0.24, 0.22, 0.2, 0.18, 0.16, 0.14, 0.12, 0.10, 0.08, 0.04, 0.02, 0}, []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}))
	ret.ContactFlags().SetProperty("ParentLegalGuardian", threshold_rand_strings([]float64{0.8, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("PickupRights", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("LivesWith", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("AccessToRecords", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("ReceivesAssessmentReport", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("EmergencyContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("HasCustody", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("DisciplinaryContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("AttendanceContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("PrimaryCareProvider", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FeesBilling", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FeesAccess", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FamilyMail", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("InterventionOrder", threshold_rand_strings([]float64{0.1, 0}, []string{"N", "Y"}))

	if student != nil {
		ret.SetProperty("StudentPersonalRefId", student.RefId().String())
	}
	if contact != nil {
		ret.SetProperty("StudentContactPersonalRefId", contact.RefId().String())
	}
	if out, ok := sifxml.StudentContactRelationshipPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentContactRelationship: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* no siblings */
func Create_StudentContactPersonalAndRelationship(students []*sifxml.StudentPersonal) ([]*sifxml.StudentContactPersonal, []*sifxml.StudentContactRelationship) {
	contacts := sifxml.StudentContactPersonalSlice()
	relationships := sifxml.StudentContactRelationshipSlice()
	for _, s := range students {
		p1 := Create_StudentContactPersonal(s, 1)
		if p1 != nil {
			contacts = append(contacts, p1)
			relationships = append(relationships, Create_StudentContactRelationship(s, p1))
		}
		p2 := Create_StudentContactPersonal(s, 2)
		if p2 != nil {
			contacts = append(contacts, p2)
			relationships = append(relationships, Create_StudentContactRelationship(s, p2))
		}
	}
	return contacts, relationships
}

func Create_RoomInfo(school *sifxml.SchoolInfo) *sifxml.RoomInfo {
	return Create_RoomInfoWithStaff(school, []*sifxml.StaffPersonal{})
}

func Create_RoomInfoWithStaff(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal) *sifxml.RoomInfo {
	room_number := random_seq_gen("roomNumber", 1000)

	ret := sifxml.RoomInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("RoomNumber", strconv.Itoa(room_number))
	ret.SetProperty("Description", fmt.Sprintf("Room %d", room_number))
	ret.SetProperty("Capacity", rand.Intn(50)+10)
	ret.SetProperty("Size", float64(rand.Intn(5)+2))
	ret.SetProperty("RoomType", randomStringFromSlice([]string{"Classroom", "Classroom", "Classroom", "Classroom", "Classroom", "Classroom", "Art", "Basketball court"}))
	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if len(staff) > 0 {
		for _, s1 := range staff {
			ret.StaffList().AppendString(s1.RefId().String())
		}
	}
	if out, ok := sifxml.RoomInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to RoomInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_RoomInfos(count int, school *sifxml.SchoolInfo) []*sifxml.RoomInfo {
	ret := sifxml.RoomInfoSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_RoomInfo(school))
	}
	return ret
}

func Create_TeachingGroup(school *sifxml.SchoolInfo, students []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal, timetablesubject *sifxml.TimeTableSubject) *sifxml.TeachingGroup {

	ret := sifxml.TeachingGroup{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("Semester", 1)
	ret.SetProperty("MinClassSize", 20)
	ret.SetProperty("MaxClassSize", 40)

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
		ret.SetProperty("SchoolLocalId", school.LocalId().String())
	}

	shortname := teachingSubject()
	if timetablesubject == nil {
	} else {
		ret.SetProperty("TimeTableSubjectRefId", timetablesubject.RefId().String())
		ret.SetProperty("TimeTableSubjectLocalId", timetablesubject.SubjectLocalId().String())
		shortname = timetablesubject.SubjectShortName().String()
	}
	ret.SetProperty("ShortName", shortname)
	ret.SetProperty("LongName", teachingSubjectLongName(shortname))
	ret.SetProperty("KeyLearningArea", teachingGroupKLA(shortname))

	if len(students) > 0 {
		for _, s := range students {
			ret.StudentList().AddNew()
			ret.StudentList().Last().SetProperty("StudentPersonalRefId", s.RefId().String())
			ret.StudentList().Last().SetProperty("StudentLocalId", s.LocalId().String())
			ret.StudentList().Last().SetProperty("Name", s.PersonInfo().Name().Clone())
		}
	}
	if len(staff) > 0 {
		for _, s := range staff {
			ret.TeacherList().AddNew()
			ret.TeacherList().Last().SetProperty("StaffPersonalRefId", s.RefId().String())
			ret.TeacherList().Last().SetProperty("StaffLocalId", s.LocalId().String())
			ret.TeacherList().Last().SetProperty("Name", s.PersonInfo().Name().Clone())
		}
	}
	if out, ok := sifxml.TeachingGroupPointer(ret); !ok {
		log.Fatalf("Could not create pointer to TeachingGroup: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_FinancialAccount(parent *sifxml.FinancialAccount, location *sifxml.ChargedLocationInfo) *sifxml.FinancialAccount {
	gofakeit.Seed(0)
	ret := sifxml.FinancialAccount{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("CreationDate", random_date("2012-01-01", "2015-12-31"))
	ret.SetProperty("CreationTime", gofakeit.Date().Format("15:04:05"))
	ret.SetProperty("AccountNumber", strconv.Itoa(random_seq_gen("financial_account_number", 99999999)+1))
	ret.SetProperty("Name", gofakeit.Name())
	ret.SetProperty("ClassType", randomStringFromSlice([]string{"Asset", "Liability", "Revenue", "Expense", "Other"}))

	if parent != nil {
		ret.SetProperty("ParentAccountRefId", parent.RefId().String())
	}
	if location == nil {
		ret.SetProperty("Name", gofakeit.Name())
	} else {
		ret.SetProperty("ChargedLocationInfoRefId", location.RefId().String())
		ret.SetProperty("Name", location.Name().String())
	}
	if out, ok := sifxml.FinancialAccountPointer(ret); !ok {
		log.Fatalf("Could not create pointer to FinacialAccount: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* parent and location are either empty, or match count in size */
func Create_FinancialAccountsBase(count int, parent []*sifxml.FinancialAccount, location []*sifxml.ChargedLocationInfo) []*sifxml.FinancialAccount {
	if len(parent) != 0 && len(location) != 0 {
		if len(parent) != count {
			log.Fatalf("Mismatch in count of financial accounts (%d) and count of parent financial accounts (%d)", count, len(parent))
		}
		if len(location) != count {
			log.Fatalf("Mismatch in count of financial accounts (%d) and count of charged locations (%d)", count, len(location))
		}
	}
	ret := sifxml.FinancialAccountSlice()
	var p *sifxml.FinancialAccount = nil
	var l *sifxml.ChargedLocationInfo = nil
	for i := 0; i < count; i++ {
		if len(parent) > 0 {
			p = parent[i]
		}
		if len(location) > 0 {
			l = location[i]
		}
		ret = append(ret, Create_FinancialAccount(p, l))
	}
	return ret
}

/* half the accounts are children of other accounts. A third of all parent accounts are associated with a charged location; all children of that account are associated with the same charged location */
func Create_FinancialAccounts(count int, location []*sifxml.ChargedLocationInfo) []*sifxml.FinancialAccount {
	parent_account_count := count/2 + 1
	parent_charge_locations := sifxml.ChargedLocationInfoSlice()
	for i := 0; i < parent_account_count; i++ {
		if rand.Float32() > 0.333 {
			parent_charge_locations = append(parent_charge_locations, location[rand.Intn(len(location))])
		} else {
			parent_charge_locations = append(parent_charge_locations, nil)
		}
	}
	parent_accounts := Create_FinancialAccountsBase(parent_account_count, sifxml.FinancialAccountSlice(), parent_charge_locations)

	child_account_count := count - parent_account_count
	child_charge_locations := sifxml.ChargedLocationInfoSlice()
	child_parent_accounts := sifxml.FinancialAccountSlice()
	for i := 0; i < child_account_count; i++ {
		parent_id := rand.Intn(len(parent_accounts))
		child_parent_accounts = append(child_parent_accounts, parent_accounts[parent_id])
		child_charge_locations = append(child_charge_locations, parent_charge_locations[parent_id])
	}
	child_accounts := Create_FinancialAccountsBase(child_account_count, child_parent_accounts, child_charge_locations)
	return append(parent_accounts, child_accounts...)
}

func Create_ChargedLocationInfo(parent *sifxml.ChargedLocationInfo, school *sifxml.SchoolInfo) *sifxml.ChargedLocationInfo {
	state := create_state()

	ret := sifxml.ChargedLocationInfo{}
	ret.SetProperty("RefId", create_GUID())
	if parent != nil {
		ret.SetProperty("ParentChargedLocationInfoRefId", parent.RefId().String())
	}
	if school == nil {
		gofakeit.Seed(0)
		locationtype := randomStringFromSlice([]string{"HR", "Professional Development", "Accounting", "Management", "Cleaning"})
		ret.SetProperty("SiteCategory", "NonSchool")
		ret.SetProperty("LocationType", locationtype)
		ret.SetProperty("Name", gofakeit.Company()+" "+locationtype)
		ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
		ret.AddressList().Append(create_address(state))
		ret.PhoneNumberList().AddNew()
		ret.PhoneNumberList().Last().SetProperty("Type", "0096")
		ret.PhoneNumberList().Last().SetProperty("Number", create_phone_number(&state))
	} else {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
		ret.SetProperty("SiteCategory", "School")
		ret.SetProperty("LocationType", "School")
		ret.SetProperty("Name", school.SchoolName().String())
		ret.SetProperty("LocalId", school.LocalId().String())
		ret.SetProperty("StateProvinceId", school.StateProvinceId().String())
		ret.SetProperty("AddressList", school.AddressList().Clone())
		// Gotcha: accessing a value creates it, so we get empty value in source if we try to clone a nil value
		if !school.PhoneNumberList_IsNil() {
			ret.SetProperty("PhoneNumberList", school.PhoneNumberList().Clone())
		}
	}
	if out, ok := sifxml.ChargedLocationInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to ChargedLocationInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_ChargedLocationInfos(count int, school []*sifxml.SchoolInfo) []*sifxml.ChargedLocationInfo {
	ret := sifxml.ChargedLocationInfoSlice()
	for i := 0; i < count && i < len(school); i++ {
		ret = append(ret, Create_ChargedLocationInfo(nil, school[i]))
	}
	nonschoolcount := count - len(school)
	if nonschoolcount > 0 {
		for i := 0; i < nonschoolcount/2; i++ {
			ret = append(ret, Create_ChargedLocationInfo(nil, nil))
		}
		for i := nonschoolcount / 2; i < nonschoolcount; i++ {
			ret = append(ret, Create_ChargedLocationInfo(ret[len(school)+rand.Intn(nonschoolcount/2)], nil))
		}
	}
	return ret
}

func Create_VendorInfo() *sifxml.VendorInfo {
	gofakeit.Seed(0)
	person := gofakeit.Person()
	companytype := randomStringFromSlice([]string{"Company", "Pty Ltd", "Ltd", "Pty", "Inc"})
	companyname := gofakeit.Company()
	middlename := gofakeit.FirstName()
	emaildomain := strings.ToLower(companyname) + "." + randomStringFromSlice([]string{"com.au", "com", "com.au", "org.au"})

	ret := sifxml.VendorInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("Name", companyname+" "+companytype)
	ret.SetProperty("CustomerId", strconv.Itoa(random_seq_gen("customerid", 99999)+1000))
	ret.SetProperty("ABN", strconv.Itoa(random_seq_gen("abn", 99999999999)+1000000000))
	ret.SetProperty("RegisteredForGST", "Y")
	ret.SetProperty("PaymentTerms", "15 days")
	ret.SetProperty("BPay", strconv.Itoa(random_seq_gen("bpay", 999999)+10000))
	ret.SetProperty("BSB", strconv.Itoa(random_seq_gen("bsb", 999999)+10000))
	ret.SetProperty("AccountNumber", strconv.Itoa(random_seq_gen("account_number", 999999)+10000))
	ret.SetProperty("AccountName", companyname)

	ret.ContactInfo().Name().SetProperty("Type", "LGL")
	ret.ContactInfo().Name().SetProperty("FamilyName", person.LastName)
	ret.ContactInfo().Name().SetProperty("GivenName", person.FirstName)
	ret.ContactInfo().SetProperty("PositionTitle", "Sales")
	ret.ContactInfo().SetProperty("Role", "Sales")

	ret.ContactInfo().EmailList().AddNew()
	ret.ContactInfo().EmailList().Last().SetProperty("Type", "01")
	ret.ContactInfo().EmailList().Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, emaildomain))
	ret.ContactInfo().PhoneNumberList().AddNew()
	ret.ContactInfo().PhoneNumberList().Last().SetProperty("Type", "0096")
	ret.ContactInfo().PhoneNumberList().Last().SetProperty("Number", create_phone_number(nil))

	if out, ok := sifxml.VendorInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to VendorInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_VendorInfos(count int) []*sifxml.VendorInfo {
	ret := sifxml.VendorInfoSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_VendorInfo())
	}
	return ret
}

func Create_ScheduledActivity(school *sifxml.SchoolInfo, timetable *sifxml.TimeTable, timetablecell *sifxml.TimeTableCell, timetablesubject *sifxml.TimeTableSubject, students []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal, teachinggroups []*sifxml.TeachingGroup, rooms []*sifxml.RoomInfo) *sifxml.ScheduledActivity {
	gofakeit.Seed(0)
	starttime := gofakeit.Date()
	finishtime := starttime.Add(time.Hour * time.Duration(rand.Intn(3)))

	ret := sifxml.ScheduledActivity{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("ActivityDate", random_date(this_year()+"-01-01", this_year()+"-12-31"))
	ret.SetProperty("StartTime", starttime.Format("15:04:05"))
	ret.SetProperty("FinishTime", finishtime.Format("15:04:05"))
	ret.SetProperty("CellType", "Excursion")
	ret.SetProperty("ActivityType", "Excursion")
	ret.SetProperty("Location", "Zoo")
	ret.SetProperty("ActivityName", "Zoo Excursion")

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if timetable != nil {
		ret.SetProperty("TimeTableRefId", timetable.RefId().String())
	}
	if timetablecell == nil {
	} else {
		ret.SetProperty("TimeTableCellRefId", timetablecell.RefId().String())
	}
	if timetablesubject != nil {
		ret.SetProperty("TimeTableSubjectRefId", timetablesubject.RefId().String())
	}
	if len(students) > 0 {
		for _, s := range rooms {
			ret.StudentList().AppendString(s.RefId().String())
		}
	}
	if len(rooms) > 0 {
		for _, s := range rooms {
			ret.RoomList().AppendString(s.RefId().String())
		}
	}
	if len(teachinggroups) > 0 {
		for _, s := range teachinggroups {
			ret.TeachingGroupList().AppendString(s.RefId().String())
		}
	}
	if len(staff) > 0 {
		addTeachers(ret.TeacherList(), staff, starttime, finishtime)
	}
	if out, ok := sifxml.ScheduledActivityPointer(ret); !ok {
		log.Fatalf("Could not create pointer to ScheduledActivity: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_CalendarSummary(school *sifxml.SchoolInfo) *sifxml.CalendarSummary {
	ret := sifxml.CalendarSummary{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("DaysInSession", 67+67+68+75)
	ret.SetProperty("StartDate", this_year()+"-01-28")
	ret.SetProperty("EndDate", this_year()+"-12-19")

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if out, ok := sifxml.CalendarSummaryPointer(ret); !ok {
		log.Fatalf("Could not create pointer to CalendarSummary: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_GradingAssignment(school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup, students []*sifxml.StudentPersonal) *sifxml.GradingAssignment {
	gofakeit.Seed(0)
	createdate := random_date(this_year()+"-01-01", this_year()+"-11-01")
	duedate_time, _ := time.Parse("2006-01-02", createdate)
	duedate := duedate_time.Add(time.Hour * time.Duration(24*rand.Intn(30))).Format("2006-01-02")
	description := gofakeit.LetterN(8)

	ret := sifxml.GradingAssignment{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("GradingCategory", randomStringFromSlice([]string{"quiz", "essay", "project"}))
	ret.SetProperty("PointsPossible", 10)
	ret.SetProperty("CreateDate", createdate)
	ret.SetProperty("DueDate", duedate)
	ret.SetProperty("Weight", float64(rand.Intn(4)+2))
	ret.SetProperty("MaxAttemptsAllowed", 5)
	ret.SetProperty("Description", description)
	ret.SetProperty("DetailedDescriptionURL", "http://www.example.com/"+description)

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if teachinggroup != nil {
		ret.SetProperty("TeachingGroupRefId", teachinggroup.RefId().String())
	}
	if len(students) > 0 {
		for _, s := range students {
			ret.StudentPersonalRefIdList().AppendString(s.RefId().String())
		}
	}
	if out, ok := sifxml.GradingAssignmentPointer(ret); !ok {
		log.Fatalf("Could not create pointer to GradingAssignment: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_GradingAssignments(count int, school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup, students []*sifxml.StudentPersonal) []*sifxml.GradingAssignment {
	ret := sifxml.GradingAssignmentSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_GradingAssignment(school, teachinggroup, students))
	}
	return ret
}

func Create_GradingAssignmentScore(assignment *sifxml.GradingAssignment, school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup, student *sifxml.StudentPersonal, staff *sifxml.StaffPersonal) *sifxml.GradingAssignmentScore {
	gofakeit.Seed(0)
	ret := sifxml.GradingAssignmentScore{}
	ret.SetProperty("RefId", create_GUID())

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if staff != nil {
		ret.SetProperty("StaffPersonalRefId", staff.RefId().String())
	}
	if student != nil {
		ret.SetProperty("StudentPersonalRefId", student.RefId().String())
		ret.SetProperty("StudentPersonalLocalId", student.LocalId().String())
	}
	if teachinggroup != nil {
		ret.SetProperty("TeachingGroupRefId", teachinggroup.RefId().String())
	}

	if assignment != nil {
		ret.SetProperty("GradingAssignmentRefId", assignment.RefId().String())
		pointspossible := assignment.PointsPossible().Int()
		ret.SetProperty("ScorePoints", rand.Intn(pointspossible))
		ret.SetProperty("ScoreDescription", gofakeit.LoremIpsumSentence(10))
		duedate := assignment.DueDate().String()
		dategraded_time, err := time.Parse("2006-01-02", duedate)
		if err == nil {
			dategraded := dategraded_time.Add(time.Hour * time.Duration(24*rand.Intn(7))).Format("2006-01-02")
			ret.SetProperty("DateGraded", dategraded)
		}
	}
	if out, ok := sifxml.GradingAssignmentScorePointer(ret); !ok {
		log.Fatalf("Could not create pointer to GradingAssignmentScore: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_GradingAssignmentScores(assignment *sifxml.GradingAssignment, school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup, students []*sifxml.StudentPersonal, staff *sifxml.StaffPersonal) []*sifxml.GradingAssignmentScore {
	ret := sifxml.GradingAssignmentScoreSlice()
	for _, s := range students {
		ret = append(ret, Create_GradingAssignmentScore(assignment, school, teachinggroup, s, staff))
	}
	return ret
}

/* assume only one of the arguments is populated */
func Create_Debtor(student *sifxml.StudentPersonal, staff *sifxml.StaffPersonal, contact *sifxml.StudentContactPersonal, vendor *sifxml.VendorInfo) *sifxml.Debtor {
	ret := sifxml.Debtor{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))

	if student != nil {
		ret.BilledEntity().SetProperty("SIF_RefObject", "StudentPersonal")
		ret.BilledEntity().SetProperty("Value", student.RefId().String())
	} else if staff != nil {
		ret.BilledEntity().SetProperty("SIF_RefObject", "StaffPersonal")
		ret.BilledEntity().SetProperty("Value", staff.RefId().String())
	} else if contact != nil {
		ret.BilledEntity().SetProperty("SIF_RefObject", "StudenContactPersonal")
		ret.BilledEntity().SetProperty("Value", contact.RefId().String())
	} else if vendor != nil {
		ret.BilledEntity().SetProperty("SIF_RefObject", "VendorInfo")
		ret.BilledEntity().SetProperty("Value", vendor.RefId().String())
	} else {
		log.Fatalf("No billing entity passed to Create_Debtor()")
	}
	if out, ok := sifxml.DebtorPointer(ret); !ok {
		log.Fatalf("Could not create pointer to Debtor: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_Debtors(student []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal, contact []*sifxml.StudentContactPersonal, vendor []*sifxml.VendorInfo) []*sifxml.Debtor {
	ret := sifxml.DebtorSlice()
	for _, s := range student {
		ret = append(ret, Create_Debtor(s, nil, nil, nil))
	}
	for _, s := range staff {
		ret = append(ret, Create_Debtor(nil, s, nil, nil))
	}
	for _, s := range contact {
		ret = append(ret, Create_Debtor(nil, nil, s, nil))
	}
	for _, s := range vendor {
		ret = append(ret, Create_Debtor(nil, nil, nil, s))
	}
	return ret
}

/* if calendar_date_number = -1, ignore */
func Create_CalendarDate(calendar *sifxml.CalendarSummary, school *sifxml.SchoolInfo, date time.Time, studentholiday bool, publicholiday bool, calendar_date_number int) *sifxml.CalendarDate {
	ret := sifxml.CalendarDate{}
	ret.SetProperty("CalendarDateRefId", create_GUID())
	ret.SetProperty("Date", date.Format("2006-01-02"))

	if !studentholiday {
		ret.CalendarDateType().SetProperty("Code", "INST")
	} else if publicholiday {
		ret.CalendarDateType().SetProperty("Code", "0846")
	} else {
		ret.CalendarDateType().SetProperty("Code", "0845")
	}

	if studentholiday {
		ret.StudentAttendance().SetProperty("CountsTowardAttendance", "No")
		ret.StudentAttendance().SetProperty("AttendanceValue", 0.0)
	} else {
		ret.StudentAttendance().SetProperty("CountsTowardAttendance", "Yes")
		ret.StudentAttendance().SetProperty("AttendanceValue", 1.0)
	}
	if publicholiday {
		ret.TeacherAttendance().SetProperty("CountsTowardAttendance", "No")
		ret.TeacherAttendance().SetProperty("AttendanceValue", 0.0)
	} else {
		ret.TeacherAttendance().SetProperty("CountsTowardAttendance", "Yes")
		ret.TeacherAttendance().SetProperty("AttendanceValue", 1.0)
	}
	if publicholiday {
		ret.AdministratorAttendance().SetProperty("CountsTowardAttendance", "No")
		ret.AdministratorAttendance().SetProperty("AttendanceValue", 0.0)
	} else {
		ret.AdministratorAttendance().SetProperty("CountsTowardAttendance", "Yes")
		ret.AdministratorAttendance().SetProperty("AttendanceValue", 1.0)
	}

	if calendar_date_number != -1 {
		ret.SetProperty("CalendarDateNumber", calendar_date_number)
	}
	if calendar != nil {
		ret.SetProperty("CalendarSummaryRefId", calendar.RefId().String())
		ret.SetProperty("SchoolYear", calendar.SchoolYear().String())
	}
	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	}
	if out, ok := sifxml.CalendarDatePointer(ret); !ok {
		log.Fatalf("Could not create pointer to CalendarDate: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_CalendarDates(calendar *sifxml.CalendarSummary, school *sifxml.SchoolInfo) []*sifxml.CalendarDate {
	ret := sifxml.CalendarDateSlice()
	local_c := cal.NewBusinessCalendar()
	blank_c := cal.NewBusinessCalendar()
	local_c.AddHoliday(au.HolidaysVIC...)
	start, _ := time.Parse("2006-01-02", calendar.StartDate().String())
	end, _ := time.Parse("2006-01-02", calendar.EndDate().String())
	calendar_date_number := 1
	for date := start; date.Before(end); date = blank_c.NextWorkdayStart(date.Add(time.Hour * 10)) {
		student_holiday := !local_c.IsWorkday(date)
		if calendar_date_number > 45 && calendar_date_number < 55 ||
			calendar_date_number > 105 && calendar_date_number < 115 ||
			calendar_date_number > 105 && calendar_date_number < 115 ||
			calendar_date_number > 155 && calendar_date_number < 165 {
			student_holiday = true
		}
		ret = append(ret, Create_CalendarDate(calendar, school, date, student_holiday, !local_c.IsWorkday(date), calendar_date_number))
		calendar_date_number++
	}
	return ret
}

func Create_TimeTable(school *sifxml.SchoolInfo) *sifxml.TimeTable {
	ret := sifxml.TimeTable{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("Title", "Timetable "+ret.RefId().String())
	ret.SetProperty("DaysPerCycle", 10)
	ret.SetProperty("PeriodsPerDay", 7)
	ret.SetProperty("TeachingPeriodsPerDay", 6)

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	periods := periods()
	for i := 0; i < 10; i++ {
		ret.TimeTableDayList().AddNew()
		ret.TimeTableDayList().Last().SetProperty("DayTitle", days[i%5])
		ret.TimeTableDayList().Last().SetProperty("DayId", strconv.Itoa(i+1))
		for j := 0; j < 7; j++ {
			ret.TimeTableDayList().Last().TimeTablePeriodList().AddNew()
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("PeriodTitle", periods[j])
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("PeriodId", strconv.Itoa(j+1))
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("StartTime", periodStart(j+1).Format("15:04:05"))
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("EndTime", periodEnd(j+1).Format("15:04:05"))
		}
	}

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
		ret.SetProperty("SchoolLocalId", school.LocalId().String())
		ret.SetProperty("SchoolName", school.SchoolName().String())
	}
	if out, ok := sifxml.TimeTablePointer(ret); !ok {
		log.Fatalf("Could not create pointer to TimeTable: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* assuming that subject is the same enum value across year levels, and that TTS can be grouped on SubjectShortName = subject */
func Create_TimeTableSubject(school *sifxml.SchoolInfo, course *sifxml.SchoolCourseInfo, subject string, acyear string, acyear_end string, semester int) *sifxml.TimeTableSubject {
	code := strconv.Itoa(random_seq_gen("timetablesubjects", 900) + 100)

	ret := sifxml.TimeTableSubject{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolYear", this_year())

	ret.SetProperty("SubjectShortName", subject)
	ret.SetProperty("SubjectLongName", teachingSubjectLongName(subject))
	ret.SetProperty("SubjectType", randomStringFromSlice([]string{"Core", "Elective", "?"}))
	if acyear_end == "" {
		ret.SetProperty("SubjectLocalId", fmt.Sprintf("%s Y%s %s", teachingSubjectLongName(subject), acyear, code))
	} else {
		ret.SetProperty("SubjectLocalId", fmt.Sprintf("%s Y%s-%s %s", teachingSubjectLongName(subject), acyear, acyear_end, code))
	}
	ret.SetProperty("ProposedMinClassSize", float64(rand.Intn(15)+5))
	ret.SetProperty("ProposedMaxClassSize", ret.ProposedMinClassSize().Float())
	ret.SetProperty("Semester", semester)

	ret.OtherCodeList().AddNew()
	ret.OtherCodeList().Last().SetProperty("Codeset", "mycodeset")
	ret.OtherCodeList().Last().SetProperty("Value", strconv.Itoa(random_seq_gen("other_timetablesubjects", 900)+100))

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
		ret.SetProperty("SchoolLocalId", school.LocalId().String())
	}
	if course != nil {
		ret.SetProperty("SchoolCourseInfoRefId", course.RefId().String())
		ret.SetProperty("SchoolCourseLocalId", course.CourseCode().String())
	}

	if acyear_end == "" {
		ret.AcademicYear().SetProperty("Code", acyear)
	} else {
		ret.AcademicYearRange().Start().SetProperty("Code", acyear)
		ret.AcademicYearRange().End().SetProperty("Code", acyear_end)
	}

	if out, ok := sifxml.TimeTableSubjectPointer(ret); !ok {
		log.Fatalf("Could not create pointer to TimeTableSubject: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* assume terms are consecutive; offer each subject each term */
func Create_TimeTableSubjects(school *sifxml.SchoolInfo, subjects []string, terms []*sifxml.TermInfo) []*sifxml.TimeTableSubject {
	random_seq_gen_reset("timetablesubjects")
	random_seq_gen_reset("other_timetablesubjects")
	if len(subjects) == 0 {
		subjects = all_teachingSubjects()
	}
	ret := sifxml.TimeTableSubjectSlice()
	for term_no, _ := range terms {
		for _, s := range subjects {
			for _, y := range Schooltype2Yearlevels(school.SchoolType().String()) {
				ret = append(ret, Create_TimeTableSubject(school, nil, s, y, "", term_no+1))
			}
		}
	}
	return ret
}

func Create_TermInfo(school *sifxml.SchoolInfo, semester int) *sifxml.TermInfo {
	ret := sifxml.TermInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("StartDate", term_start_date(this_year(), semester))
	ret.SetProperty("EndDate", term_end_date(this_year(), semester))
	ret.SetProperty("TermCode", fmt.Sprintf("Term %d", semester))
	ret.SetProperty("TermSpan", "0828")
	ret.SetProperty("MarkingTerm", "Y")
	ret.SetProperty("SchedulingTerm", "Y")
	ret.SetProperty("AttendanceTerm", "Y")

	if out, ok := sifxml.TermInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to TermInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* 2 a year */
func Create_TermInfos(school *sifxml.SchoolInfo) []*sifxml.TermInfo {
	ret := sifxml.TermInfoSlice()
	for semester := 1; semester <= 2; semester++ {
		ret = append(ret, Create_TermInfo(school, semester))
	}
	return ret
}

func Create_TimeTableCell(day string, period string, celltype string, school *sifxml.SchoolInfo, timetable *sifxml.TimeTable, subject *sifxml.TimeTableSubject, teachinggroup *sifxml.TeachingGroup, room *sifxml.RoomInfo, rooms []*sifxml.RoomInfo, staff *sifxml.StaffPersonal, teachers []*sifxml.StaffPersonal) *sifxml.TimeTableCell {
	ret := sifxml.TimeTableCell{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("DayId", day)
	ret.SetProperty("PeriodId", period)
	ret.SetProperty("CellType", celltype)

	if school != nil {
		ret.SetProperty("SchoolInfoRefId", school.RefId().String())
		ret.SetProperty("SchoolLocalId", school.LocalId().String())
	}
	if timetable != nil {
		ret.SetProperty("TimeTableRefId", timetable.RefId().String())
		ret.SetProperty("TimeTableLocalId", timetable.LocalId().String())
	}
	if subject != nil {
		ret.SetProperty("TimeTableSubjectRefId", subject.RefId().String())
		ret.SetProperty("SubjectLocalId", subject.SubjectLocalId().String())
	}
	if teachinggroup != nil {
		ret.SetProperty("TeachingGroupRefId", teachinggroup.RefId().String())
		ret.SetProperty("TeachingGroupLocalId", teachinggroup.LocalId().String())
	}
	if room != nil {
		ret.SetProperty("RoomInfoRefId", room.RefId().String())
		ret.SetProperty("RoomNumber", room.RoomNumber().String())
	}
	if len(rooms) > 0 {
		for _, s := range rooms {
			ret.RoomList().AppendString(s.RefId().String())
		}
	}
	if staff != nil {
		ret.SetProperty("StaffPersonalRefId", staff.RefId().String())
		ret.SetProperty("StaffLocalId", staff.LocalId().String())
	}
	if len(teachers) > 0 {
		period_int, _ := strconv.Atoi(period)
		addTeachers(ret.TeacherList(), teachers, periodStart(period_int), periodEnd(period_int))
	}

	if out, ok := sifxml.TimeTableCellPointer(ret); !ok {
		log.Fatalf("Could not create pointer to TimeTableCell: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_SessionInfo(c *sifxml.TimeTableCell, date string) *sifxml.SessionInfo {
	periodid, _ := strconv.Atoi(c.PeriodId().String())

	ret := sifxml.SessionInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolInfoRefId", c.SchoolInfoRefId().String())
	ret.SetProperty("TimeTableCellRefId", c.RefId().String())
	ret.SetProperty("SchoolYear", date[0:4])
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("TimeTableSubjectLocalId", c.SubjectLocalId().String())
	ret.SetProperty("TeachingGroupLocalId", c.TeachingGroupLocalId().String())
	ret.SetProperty("SchoolLocalId", c.SchoolLocalId().String())
	ret.SetProperty("StaffPersonalLocalId", c.StaffLocalId().String())
	ret.SetProperty("RoomNumber", c.RoomNumber().String())
	ret.SetProperty("DayId", c.DayId().String())
	ret.SetProperty("PeriodId", c.PeriodId().String())
	ret.SetProperty("SessionDate", date)
	ret.SetProperty("StartTime", periodStart(periodid).Format("15:04:05"))
	ret.SetProperty("FinishTime", periodEnd(periodid).Format("15:04:05"))
	ret.SetProperty("RollMarked", "Y")

	if out, ok := sifxml.SessionInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SessionInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* presuppose weekly recurrence. presuppose full year not terms for timetable cell */
func Create_SessionInfos(cells []*sifxml.TimeTableCell, dates []*sifxml.CalendarDate) []*sifxml.SessionInfo {
	calendar := make(map[string]*sifxml.CalendarDate)
	for _, d := range dates {
		calendar[d.Date().String()] = d
	}
	ret := sifxml.SessionInfoSlice()
	for _, c := range cells {
		yr, _ := strconv.Atoi(this_year())
		t := time.Date(yr, time.Month(1), 0, 0, 0, 0, 0, time.UTC)
		dayid, _ := strconv.Atoi(c.DayId().String())
		day := (8-int(t.Weekday()))%7 + dayid
		end := time.Date(yr, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
		for t := time.Date(yr, time.Month(1), day, 0, 0, 0, 0, time.UTC); t.Before(end); t = t.Add(time.Hour * 24 * 7) {
			date := t.Format("2006-01-02")
			if _, ok := calendar[date]; ok {
				if calendar[date].StudentAttendance().CountsTowardAttendance().String() == "Yes" {
					ret = append(ret, Create_SessionInfo(c, date))
				}
			}
		}
	}
	return ret
}

/* either period event (cell) or excursion for teaching group */
func Create_ScheduledActivityBasic(date string, c *sifxml.TimeTableCell, tg *sifxml.TeachingGroup) *sifxml.ScheduledActivity {
	ret := sifxml.ScheduledActivity{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("ActivityDate", date)
	if c != nil {
		ret.SetProperty("SchoolInfoRefId", c.SchoolInfoRefId().String())
		ret.SetProperty("TimeTableCellRefId", c.RefId().String())
		ret.SetProperty("DayId", c.DayId().String())
		ret.SetProperty("PeriodId", c.PeriodId().String())
		ret.SetProperty("TimeTableRefId", c.TimeTableRefId().String())
		periodid, _ := strconv.Atoi(c.PeriodId().String())
		ret.SetProperty("StartTime", periodStart(periodid).Format("15:04:05"))
		ret.SetProperty("FinishTime", periodEnd(periodid).Format("15:04:05"))
		ret.SetProperty("CellType", c.CellType().String())
		if !c.TeacherList_IsNil() {
			ret.SetProperty("TeacherList", c.TeacherList().Clone())
		}
		if !c.RoomList_IsNil() {
			ret.SetProperty("RoomList", c.RoomList().Clone())
		}
	} else if tg != nil {
		ret.SetProperty("SchoolInfoRefId", tg.SchoolInfoRefId().String())
		ret.SetProperty("StartTime", "09:00:00")
		ret.SetProperty("FinishTime", "15:00:00")
		ret.SetProperty("Location", "Zoo")
		ret.SetProperty("ActivityType", "Excursion")
		ret.SetProperty("ActivityName", "Zoo Excursion")
		ret.TeachingGroupList().AppendString(tg.RefId().String())
	}

	if out, ok := sifxml.ScheduledActivityPointer(ret); !ok {
		log.Fatalf("Could not create pointer to ScheduledActivity: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* presuppose weekly recurrence. presuppose full year not terms for timetable cell.
Create one scheduled activity for each session, and one excursion for each teaching group*/
func Create_ScheduledActivities(cells []*sifxml.TimeTableCell, dates []*sifxml.CalendarDate, tg []*sifxml.TeachingGroup) []*sifxml.ScheduledActivity {
	calendar := make(map[string]*sifxml.CalendarDate)
	for _, d := range dates {
		calendar[d.Date().String()] = d
	}
	ret := sifxml.ScheduledActivitySlice()
	for _, c := range cells {
		yr, _ := strconv.Atoi(this_year())
		t := time.Date(yr, time.Month(1), 0, 0, 0, 0, 0, time.UTC)
		dayid, _ := strconv.Atoi(c.DayId().String())
		day := (8-int(t.Weekday()))%7 + dayid
		end := time.Date(yr, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
		for t := time.Date(yr, time.Month(1), day, 0, 0, 0, 0, time.UTC); t.Before(end); t = t.Add(time.Hour * 24 * 7) {
			date := t.Format("2006-01-02")
			if _, ok := calendar[date]; ok {
				if calendar[date].StudentAttendance().CountsTowardAttendance().String() == "Yes" {
					ret = append(ret, Create_ScheduledActivityBasic(date, c, nil))
				}
			}
		}
	}
	for _, g := range tg {
		for c := dates[rand.Intn(len(dates))]; c.StudentAttendance().CountsTowardAttendance().String() == "No"; c = dates[rand.Intn(len(dates))] {
			date := c.Date().String()
			ret = append(ret, Create_ScheduledActivityBasic(date, nil, g))
		}
	}
	return ret
}

/* assume terms are consecutive, and that TimeTableSubject is incremental number of semester.
Updates SchoolCourseInfoRefId on TimeTableSubject*/
func Create_SchoolCourseInfo(s *sifxml.TimeTableSubject, terms []*sifxml.TermInfo) *sifxml.SchoolCourseInfo {
	ret := sifxml.SchoolCourseInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolInfoRefId", s.SchoolInfoRefId().String())
	ret.SetProperty("SchoolLocalId", s.SchoolLocalId().String())
	ret.SetProperty("SchoolYear", s.SchoolYear().String())
	ret.SetProperty("CourseCode", fmt.Sprintf("%05d", random_seq_gen("course_code", 99999)+1))
	ret.SetProperty("CourseTitle", s.SubjectLocalId().String())
	ret.SetProperty("TermInfoRefId", terms[s.Semester().Int()-1].RefId().String())

	s.SetProperty("SchoolCourseInfoRefId", ret.RefId().String())

	if out, ok := sifxml.SchoolCourseInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SessionInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

func Create_SchoolCourseInfos(subjects []*sifxml.TimeTableSubject, terms []*sifxml.TermInfo) []*sifxml.SchoolCourseInfo {
	ret := sifxml.SchoolCourseInfoSlice()
	for _, s := range subjects {
		ret = append(ret, Create_SchoolCourseInfo(s, terms))
	}
	return ret
}
