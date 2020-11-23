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
	//ret.Street = ret.Street.New()
	ret.StreetRead().SetProperty("StreetNumber", gofakeit.StreetNumber())
	ret.Street.SetProperty("StreetName", gofakeit.StreetName())
	ret.Street.SetProperty("StreetType", randomStringFromSlice([]string{"Avenue", "Boulevard", "Cove", "Court", "Crescent", "Drive", "Esplanade", "Lane", "Place", "Road", "Square", "Street", "Terrace", "Walk", "Way"}))
	ret.SetProperty("City", gofakeit.City())
	ret.SetProperty("StateProvince", state)
	ret.SetProperty("PostalCode", postcode)
	return ret
}

func addTeachers(t *sifxml.ScheduledTeacherListType, staff []*sifxml.StaffPersonal, start time.Time, finish time.Time) *sifxml.ScheduledTeacherListType {
	for _, s := range staff {
		t = t.AddNew()
		t.Last().CopyString("StaffPersonalRefId", s.RefId)
		t.Last().CopyString("StaffLocalId", s.LocalId)
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
	//ret.PersonInfo = ret.PersonInfo.New()
	//ret.PersonInfoRead().Name = ret.PersonInfo.Name.New()
	ret.PersonInfoRead().NameRead().SetProperty("Type", "LGL")
	ret.PersonInfo.Name.SetProperty("FamilyName", person.LastName)
	ret.PersonInfo.Name.SetProperty("GivenName", person.FirstName)
	ret.PersonInfo.Name.SetProperty("MiddleName", middlename)
	//ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.DemographicsRead().SetProperty("Sex", sex_seeded(person.Gender))
	if birthyr_err == nil {
		ret.PersonInfo.Demographics.SetProperty("BirthDate", birthyr)
	}
	ret.PersonInfo.Demographics.SetProperty("IndigenousStatus", threshold_rand_strings([]float64{0.2, 0.15, 0.1, 0.5, 0}, []string{"4", "1", "2", "3", "9"}))
	ret.PersonInfo.Demographics.SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().SetProperty("Type", "01")
	ret.PersonInfo.EmailList.Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
	//ret.MostRecent = ret.MostRecent.New()
	if hasParent1 {
		ret.MostRecentRead().SetProperty("Parent1Language", "1201")
		ret.MostRecent.SetProperty("Parent1EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent.SetProperty("Parent1SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent.SetProperty("Parent1NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if hasParent2 {
		ret.MostRecentRead().SetProperty("Parent2Language", "1201")
		ret.MostRecent.SetProperty("Parent2EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent.SetProperty("Parent2SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent.SetProperty("Parent2NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	//ret.MostRecent.YearLevel = ret.MostRecent.YearLevel.New()
	ret.MostRecent.YearLevelRead().SetProperty("Code", yearlevel)
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
	//ret.Campus = ret.Campus.New()
	ret.CampusRead().SetProperty("ParentSchoolId", local_id)
	ret.Campus.SetProperty("SchoolCampusId", strconv.Itoa(rand.Intn(4)+1))
	ret.Campus.SetProperty("AdminStatus", threshold_rand_strings([]float64{0.8, 0}, []string{"N", "Y"}))
	ret.Campus.SetProperty("CampusType", schooltype)
	ret.SetProperty("SchoolSector", "Gov")
	ret.SetProperty("OperationalStatus", "O")
	ret.SetProperty("IndependentSchool", "N")
	ret.SetProperty("SchoolType", schooltype)
	ret.SetProperty("ARIA", 1.0)
	ret.SetProperty("Entity_Open", "1990-01-01")
	ret.AddressList = ret.AddressList.Append(create_address(state))
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
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if student != nil {
		ret.CopyString("StudentPersonalRefId", student.RefId)
	}

	//ret.YearLevel = ret.YearLevel.New()
	if student != nil && student.MostRecent != nil && student.MostRecent.YearLevel != nil {
		ret.YearLevelRead().CopyString("Code", student.MostRecent.YearLevel.Code)
	} else {
		ret.YearLevelRead().SetProperty("Code", strconv.Itoa(rand.Intn(12)+1))
	}
	if school != nil && student != nil {
		student.PersonInfo.EmailList.Last().SetProperty("Value", create_email(student.PersonInfo.Name.GivenName.String(), student.PersonInfo.Name.MiddleName.String(), student.PersonInfo.Name.FamilyName.String(), create_school_email_domain(school)))
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
	ret.OtherIdList = ret.OtherIdList.AddNew()
	ret.OtherIdList.Last().SetProperty("Type", "DET_USER_ID")
	ret.OtherIdList.Last().SetProperty("Value", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.OtherIdList = ret.OtherIdList.AddNew()
	ret.OtherIdList.Last().SetProperty("Type", threshold_rand_strings([]float64{0.1, 0}, []string{"pep", "cep"}))
	ret.OtherIdList.Last().SetProperty("Value", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	//ret.PersonInfo = ret.PersonInfo.New()
	//ret.PersonInfo.Name = ret.PersonInfo.Name.New()
	ret.PersonInfoRead().NameRead().SetProperty("Type", "LGL")
	ret.PersonInfo.Name.SetProperty("FamilyName", person.LastName)
	ret.PersonInfo.Name.SetProperty("GivenName", person.FirstName)
	ret.PersonInfo.Name.SetProperty("PreferredGivenName", person.FirstName)
	ret.PersonInfo.Name.SetProperty("MiddleName", middlename)
	ret.PersonInfo.Name.SetProperty("Title", create_salutation(person.Gender))
	//ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.DemographicsRead().SetProperty("Sex", sex_seeded(person.Gender))
	ret.PersonInfo.Demographics.SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo.Demographics.SetProperty("BirthDate", random_date(strconv.Itoa(time.Now().Year()-65)+"-01-01", strconv.Itoa(time.Now().Year()-25)+"-12-31"))
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().SetProperty("Type", "01")
	ret.PersonInfo.EmailList.Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
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
	//ret.StaffActivity = ret.StaffActivity.New()
	ret.StaffActivityRead().SetProperty("Code", randomStringFromSlice(sifxml.AUCodeSetsStaffActivityType_values))

	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if staff != nil {
		ret.CopyString("StaffPersonalRefId", staff.RefId)
	}
	if school != nil && staff != nil {
		staff.PersonInfo.EmailList.Last().SetProperty("Value", create_email(staff.PersonInfo.Name.GivenName.String(), staff.PersonInfo.Name.MiddleName.String(), staff.PersonInfo.Name.FamilyName.String(), create_school_email_domain(school)))
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
	if student != nil && student.MostRecent.Parent1Language == nil && ordinal == 1 {
		return nil
	}
	if student != nil && student.MostRecent.Parent2Language == nil && ordinal == 2 {
		return nil
	}
	state := create_state()
	if student != nil && student.PersonInfo.AddressList != nil {
		state = fmt.Sprint(student.PersonInfo.AddressList.Last().StateProvince)
	}

	gofakeit.Seed(0)
	person := gofakeit.Person()
	middlename := gofakeit.FirstName()

	ret := sifxml.StudentContactPersonal{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	//ret.PersonInfo = ret.PersonInfo.New()
	//ret.PersonInfo.Name = ret.PersonInfo.Name.New()
	ret.PersonInfoRead().NameRead().SetProperty("Type", "LGL")
	ret.PersonInfo.Name.SetProperty("FamilyName", person.LastName)
	ret.PersonInfo.Name.SetProperty("GivenName", person.FirstName)
	ret.PersonInfo.Name.SetProperty("PreferredGivenName", person.FirstName)
	ret.PersonInfo.Name.SetProperty("PreferredFamilyName", person.LastName)
	ret.PersonInfo.Name.SetProperty("MiddleName", gofakeit.FirstName())
	ret.PersonInfo.Name.SetProperty("Title", create_salutation(person.Gender))
	//ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.DemographicsRead().SetProperty("Sex", sex_seeded(person.Gender))
	ret.PersonInfo.Demographics.SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo.Demographics.LanguageList = ret.PersonInfo.Demographics.LanguageList.AddNew()
	ret.PersonInfo.Demographics.LanguageList.Last().SetProperty("Code", "1201")
	ret.PersonInfo.Demographics.LanguageList.Last().SetProperty("LanguageType", "1")
	if rand.Float64() < 0.2 {
		ret.PersonInfo.Demographics.LanguageList = ret.PersonInfo.Demographics.LanguageList.AddNew()
		ret.PersonInfo.Demographics.LanguageList.Last().SetProperty("Code", randomStringFromSlice([]string{"0002", "7101", "2401", "2201", "5203", "4202"}))
		ret.PersonInfo.Demographics.LanguageList.Last().SetProperty("LanguageType", "2")
	}
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().SetProperty("Type", "01")
	ret.PersonInfo.EmailList.Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, create_commercial_email_domain()))
	ret.PersonInfo.PhoneNumberList = ret.PersonInfo.PhoneNumberList.AddNew()
	ret.PersonInfo.PhoneNumberList.Last().SetProperty("Type", "0096")
	ret.PersonInfo.PhoneNumberList.Last().SetProperty("Number", create_phone_number(nil))
	ret.PersonInfo.AddressList = ret.PersonInfo.AddressList.Append(create_address(state))
	if student != nil && ordinal == 1 {
		ret.CopyString("EmploymentType", student.MostRecent.Parent1EmploymentType)
		ret.CopyString("SchoolEducationalLevel", student.MostRecent.Parent1SchoolEducationLevel)
		ret.CopyString("NonSchoolEducation", student.MostRecent.Parent1NonSchoolEducation)
	} else if student != nil && ordinal == 2 {
		ret.CopyString("EmploymentType", student.MostRecent.Parent2EmploymentType)
		ret.CopyString("SchoolEducationalLevel", student.MostRecent.Parent2SchoolEducationLevel)
		ret.CopyString("NonSchoolEducation", student.MostRecent.Parent2NonSchoolEducation)
	} else {
		ret.SetProperty("EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.SetProperty("SchoolEducationalLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.SetProperty("NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if student != nil && rand.Float64() < 0.8 {
		ret.PersonInfo.Name.CopyString("FamilyName", student.PersonInfo.Name.FamilyName)
		ret.PersonInfo.Name.CopyString("PreferredFamilyName", student.PersonInfo.Name.FamilyName)
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
	//ret.Relationship = ret.Relationship.New()
	ret.RelationshipRead().SetProperty("Code", threshold_rand_strings([]float64{0.26, 0.24, 0.22, 0.2, 0.18, 0.16, 0.14, 0.12, 0.10, 0.08, 0.04, 0.02, 0}, []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}))
	//ret.ContactFlags = ret.ContactFlags.New()
	ret.ContactFlagsRead().SetProperty("ParentLegalGuardian", threshold_rand_strings([]float64{0.8, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("PickupRights", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("LivesWith", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("AccessToRecords", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("ReceivesAssessmentReport", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("EmergencyContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("HasCustody", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("DisciplinaryContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("AttendanceContact", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("PrimaryCareProvider", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("FeesBilling", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("FeesAccess", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("FamilyMail", threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags.SetProperty("InterventionOrder", threshold_rand_strings([]float64{0.1, 0}, []string{"N", "Y"}))

	if student != nil {
		ret.CopyString("StudentPersonalRefId", student.RefId)
	}
	if contact != nil {
		ret.CopyString("StudentContactPersonalRefId", contact.RefId)
	}
	if out, ok := sifxml.StudentContactRelationshipPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentContactRelationship: %+v", ret)
		return nil
	} else {
		return out
	}
}

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
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if len(staff) > 1 {
		//ret.StaffList = ret.StaffList.New()
		for _, s1 := range staff {
			ret.StaffList = ret.StaffListRead().AppendString(s1.RefId)
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

func Create_TeachingGroup(school *sifxml.SchoolInfo, students []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal) *sifxml.TeachingGroup {
	shortname := randomStringFromSlice([]string{"MAT", "ENG", "PHYS", "BIO", "CHEM", "COMP", "VIS", "ECON", "HIST"})

	ret := sifxml.TeachingGroup{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("ShortName", shortname)
	ret.SetProperty("LongName", teachingGroupLongName(shortname))
	ret.SetProperty("KeyLearningArea", teachingGroupKLA(shortname))
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("Semester", 1)
	ret.SetProperty("MinClassSize", 20)
	ret.SetProperty("MaxClassSize", 40)

	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
		ret.CopyString("SchoolLocalId", school.LocalId)
	}
	if len(students) > 1 {
		//ret.StudentList = ret.StudentList.New()
		for _, s := range students {
			ret.StudentList = ret.StudentListRead().AddNew()
			ret.StudentList.Last().CopyString("StudentPersonalRefId", s.RefId)
			ret.StudentList.Last().CopyString("StudentLocalId", s.LocalId)
			ret.StudentList.Last().CopyClone("Name", s.PersonInfo.Name)
		}
	}
	if len(staff) > 1 {
		//ret.TeacherList = ret.TeacherList.New()
		for _, s := range staff {
			ret.TeacherList = ret.TeacherListRead().AddNew()
			ret.TeacherList.Last().CopyString("StaffPersonalRefId", s.RefId)
			ret.TeacherList.Last().CopyString("StaffLocalId", s.LocalId)
			ret.TeacherList.Last().CopyClone("Name", s.PersonInfo.Name)
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
		ret.CopyString("ParentAccountRefId", parent.RefId)
	}
	if location != nil {
		ret.CopyString("ChargedLocationInfoRefId", location.RefId)
	}
	if out, ok := sifxml.FinancialAccountPointer(ret); !ok {
		log.Fatalf("Could not create pointer to FinacialAccount: %+v", ret)
		return nil
	} else {
		return out
	}
}

/* parent and location are either empty, or match count in size */
func Create_FinancialAccounts(count int, parent []*sifxml.FinancialAccount, location []*sifxml.ChargedLocationInfo) []*sifxml.FinancialAccount {
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

func Create_ChargedLocationInfo(parent *sifxml.ChargedLocationInfo, school *sifxml.SchoolInfo) *sifxml.ChargedLocationInfo {
	state := create_state()

	ret := sifxml.ChargedLocationInfo{}
	ret.SetProperty("RefId", create_GUID())
	if parent != nil {
		ret.CopyString("ParentChargedLocationInfoRefId", parent.RefId)
	}
	if school == nil {
		gofakeit.Seed(0)
		locationtype := randomStringFromSlice([]string{"HR", "Professional Development", "Accounting", "Management", "Cleaning"})
		ret.SetProperty("SiteCategory", "NonSchool")
		ret.SetProperty("LocationType", locationtype)
		ret.SetProperty("Name", gofakeit.Company()+" "+locationtype)
		ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
		ret.AddressList = ret.AddressList.Append(create_address(state))
		ret.PhoneNumberList = ret.PhoneNumberList.AddNew()
		ret.PhoneNumberList.Last().SetProperty("Type", "0096")
		ret.PhoneNumberList.Last().SetProperty("Number", create_phone_number(&state))
	} else {
		ret.CopyString("SchoolInfoRefId", school.RefId)
		ret.SetProperty("SiteCategory", "School")
		ret.SetProperty("LocationType", "School")
		ret.CopyString("Name", school.SchoolName)
		ret.CopyString("LocalId", school.LocalId)
		ret.CopyString("StateProvinceId", school.StateProvinceId)
		ret.CopyClone("AddressList", school.AddressList)
		ret.CopyClone("PhoneNumberList", school.PhoneNumberList)
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

	//ret.ContactInfo = ret.ContactInfo.New()
	//ret.ContactInfo.Name = ret.ContactInfo.Name.New()
	ret.ContactInfoRead().NameRead().SetProperty("Type", "LGL")
	ret.ContactInfo.Name.SetProperty("FamilyName", person.LastName)
	ret.ContactInfo.Name.SetProperty("GivenName", person.FirstName)
	ret.ContactInfo.SetProperty("PositionTitle", "Sales")
	ret.ContactInfo.SetProperty("Role", "Sales")

	ret.ContactInfo.EmailList = ret.ContactInfo.EmailList.AddNew()
	ret.ContactInfo.EmailList.Last().SetProperty("Type", "01")
	ret.ContactInfo.EmailList.Last().SetProperty("Value", create_email(person.FirstName, middlename, person.LastName, emaildomain))
	ret.ContactInfo.PhoneNumberList = ret.ContactInfo.PhoneNumberList.AddNew()
	ret.ContactInfo.PhoneNumberList.Last().SetProperty("Type", "0096")
	ret.ContactInfo.PhoneNumberList.Last().SetProperty("Number", create_phone_number(nil))

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
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if timetable != nil {
		ret.CopyString("TimeTableRefId", timetable.RefId)
	}
	if timetablecell == nil {
	} else {
		ret.CopyString("TimeTableCellRefId", timetablecell.RefId)
	}
	if timetablesubject != nil {
		ret.CopyString("TimeTableSubjectRefId", timetablesubject.RefId)
	}
	if len(students) > 0 {
		//ret.StudentList = ret.StudentList.New()
		for _, s := range rooms {
			ret.StudentList = ret.StudentListRead().AppendString(s.RefId)
		}
	}
	if len(rooms) > 0 {
		//ret.RoomList = ret.RoomList.New()
		for _, s := range rooms {
			ret.RoomList = ret.RoomListRead().AppendString(s.RefId)
		}
	}
	if len(teachinggroups) > 0 {
		//ret.TeachingGroupList = ret.TeachingGroupList.New()
		for _, s := range teachinggroups {
			ret.TeachingGroupList = ret.TeachingGroupListRead().AppendString(s.RefId)
		}
	}
	if len(staff) > 0 {
		//ret.TeacherList = ret.TeacherList.New()
		ret.TeacherList = addTeachers(ret.TeacherListRead(), staff, starttime, finishtime)
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
		ret.CopyString("SchoolInfoRefId", school.RefId)
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
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if teachinggroup != nil {
		ret.CopyString("TeachingGroupRefId", teachinggroup.RefId)
	}
	if len(students) > 0 {
		//ret.StudentPersonalRefIdList = ret.StudentPersonalRefIdList.New()
		for _, s := range students {
			ret.StudentPersonalRefIdList = ret.StudentPersonalRefIdListRead().AppendString(s.RefId)
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
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if staff != nil {
		ret.CopyString("StaffPersonalRefId", staff.RefId)
	}
	if student != nil {
		ret.CopyString("StudentPersonalRefId", student.RefId)
		ret.CopyString("StudentPersonalLocalId", student.LocalId)
	}
	if teachinggroup != nil {
		ret.CopyString("TeachingGroupRefId", teachinggroup.RefId)
	}

	if assignment != nil {
		ret.CopyString("GradingAssignmentRefId", assignment.RefId)
		pointspossible := assignment.PointsPossible.Int()
		ret.SetProperty("ScorePoints", rand.Intn(pointspossible))
		ret.SetProperty("ScoreDescription", gofakeit.LoremIpsumSentence(10))
		duedate := assignment.DueDate.String()
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

	//ret.BilledEntity = ret.BilledEntity.New()
	if student != nil {
		ret.BilledEntityRead().SetProperty("SIF_RefObject", "StudentPersonal")
		ret.BilledEntity.CopyString("Value", student.RefId)
	} else if staff != nil {
		ret.BilledEntityRead().SetProperty("SIF_RefObject", "StaffPersonal")
		ret.BilledEntity.CopyString("Value", staff.RefId)
	} else if contact != nil {
		ret.BilledEntityRead().SetProperty("SIF_RefObject", "StudenContactPersonal")
		ret.BilledEntity.CopyString("Value", contact.RefId)
	} else if vendor != nil {
		ret.BilledEntityRead().SetProperty("SIF_RefObject", "VendorInfo")
		ret.BilledEntity.CopyString("Value", vendor.RefId)
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

	//ret.CalendarDateType = ret.CalendarDateType.New()
	if !studentholiday {
		ret.CalendarDateTypeRead().SetProperty("Code", "INST")
	} else if publicholiday {
		ret.CalendarDateTypeRead().SetProperty("Code", "0846")
	} else {
		ret.CalendarDateTypeRead().SetProperty("Code", "0845")
	}

	//ret.StudentAttendance = ret.StudentAttendance.New()
	if studentholiday {
		ret.StudentAttendanceRead().SetProperty("CountsTowardAttendance", "No")
		ret.StudentAttendanceRead().SetProperty("AttendanceValue", 0.0)
	} else {
		ret.StudentAttendanceRead().SetProperty("CountsTowardAttendance", "Yes")
		ret.StudentAttendance.SetProperty("AttendanceValue", 1.0)
	}
	//ret.TeacherAttendance = ret.TeacherAttendance.New()
	if publicholiday {
		ret.TeacherAttendanceRead().SetProperty("CountsTowardAttendance", "No")
		ret.TeacherAttendance.SetProperty("AttendanceValue", 0.0)
	} else {
		ret.TeacherAttendanceRead().SetProperty("CountsTowardAttendance", "Yes")
		ret.TeacherAttendance.SetProperty("AttendanceValue", 1.0)
	}
	//ret.AdministratorAttendance = ret.AdministratorAttendance.New()
	if publicholiday {
		ret.AdministratorAttendanceRead().SetProperty("CountsTowardAttendance", "No")
		ret.AdministratorAttendance.SetProperty("AttendanceValue", 0.0)
	} else {
		ret.AdministratorAttendanceRead().SetProperty("CountsTowardAttendance", "Yes")
		ret.AdministratorAttendance.SetProperty("AttendanceValue", 1.0)
	}

	if calendar_date_number != -1 {
		ret.SetProperty("CalendarDateNumber", calendar_date_number)
	}
	if calendar != nil {
		ret.CopyString("CalendarSummaryRefId", calendar.RefId)
		ret.CopyString("SchoolYear", calendar.SchoolYear)
	}
	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
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
	start, _ := time.Parse("2006-01-02", calendar.StartDate.String())
	end, _ := time.Parse("2006-01-02", calendar.EndDate.String())
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
