package populate

import (
	"../sifxml"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

func create_address(state string) sifxml.AddressType {
	gofakeit.Seed(0)
	if state == "" {
		state = create_state()
	}
	postcode := postcode_seeded(state)

	ret := sifxml.AddressType{}
	ret.Set("Type", "0123")
	ret.Set("Role", "012A")
	ret.Street = ret.Street.New()
	ret.Street.Set("StreetNumber", gofakeit.StreetNumber())
	ret.Street.Set("StreetName", gofakeit.StreetName())
	ret.Street.Set("StreetType", randomStringFromSlice([]string{"Avenue", "Boulevard", "Cove", "Court", "Crescent", "Drive", "Esplanade", "Lane", "Place", "Road", "Square", "Street", "Terrace", "Walk", "Way"}))
	ret.Set("City", gofakeit.City())
	ret.Set("StateProvince", state)
	ret.Set("PostalCode", postcode)
	return ret
}

func Create_StudentPersonal() *sifxml.StudentPersonal {
	yearlevel := rand.Intn(12) + 1
	gofakeit.Seed(0)
	person := gofakeit.Person()
	middlename := gofakeit.FirstName()
	hasParent1 := rand.Float64() < 0.8
	hasParent2 := rand.Float64() < 0.8
	if !hasParent1 {
		hasParent2 = true
	}

	ret := sifxml.StudentPersonal{}
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.Set("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.PersonInfo = ret.PersonInfo.New()
	ret.PersonInfo.Name = ret.PersonInfo.Name.New()
	ret.PersonInfo.Name.Set("Type", "LGL")
	ret.PersonInfo.Name.Set("FamilyName", person.LastName)
	ret.PersonInfo.Name.Set("GivenName", person.FirstName)
	ret.PersonInfo.Name.Set("MiddleName", middlename)
	ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.Demographics.Set("Sex", sex_seeded(person.Gender))
	ret.PersonInfo.Demographics.Set("BirthDate", birth_year(yearlevel))
	ret.PersonInfo.Demographics.Set("IndigenousStatus", randomStringFromSlice(sifxml.AUCodeSetsIndigenousStatusType_values))
	ret.PersonInfo.Demographics.Set("CountryOfBirth", "1101")
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().Set("Type", "01")
	ret.PersonInfo.EmailList.Last().Set("Value", create_email(person.FirstName, middlename, person.LastName))
	ret.MostRecent = ret.MostRecent.New()
	if hasParent1 {
		ret.MostRecent.Set("Parent1Language", "1201")
		ret.MostRecent.Set("Parent1EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent.Set("Parent1SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent.Set("Parent1NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if hasParent2 {
		ret.MostRecent.Set("Parent2Language", "1201")
		ret.MostRecent.Set("Parent2EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent.Set("Parent2SchoolEducationLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent.Set("Parent2NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	ret.MostRecent.YearLevel = ret.MostRecent.YearLevel.New()
	ret.MostRecent.YearLevel.Set("Code", strconv.Itoa(yearlevel))
	return sifxml.StudentPersonalCreate(ret)
}

func Create_StudentPersonals(count int) []*sifxml.StudentPersonal {
	ret := make([]*sifxml.StudentPersonal, 0)
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StudentPersonal())
	}
	return ret
}

func Create_SchoolInfo() *sifxml.SchoolInfo {
	local_id := strconv.Itoa(seq_gen("localId"))
	school_type := randomStringFromSlice(sifxml.AUCodeSetsSchoolLevelType_values)
	state := create_state()
	gofakeit.Seed(0)

	ret := sifxml.SchoolInfo{}
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", local_id)
	ret.Set("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 9999)+1))
	ret.Set("CommonwealthId", strconv.Itoa(random_seq_gen("schoolCommonwealthId", 9999)+1))
	ret.Set("SchoolName", school_name())
	ret.Campus = ret.Campus.New()
	ret.Campus.Set("ParentSchoolId", local_id)
	ret.Campus.Set("SchoolCampusId", strconv.Itoa(rand.Intn(4)+1))
	ret.Campus.Set("AdminStatus", threshold_rand_strings([]float64{0.8}, []string{"N", "Y"}))
	ret.Campus.Set("CampusType", school_type)
	ret.Set("SchoolSector", "Gov")
	ret.Set("OperationalStatus", "O")
	ret.Set("IndependentSchool", "N")
	ret.Set("SchoolType", school_type)
	ret.Set("ARIA", 1.0)
	ret.Set("Entity_Open", "1990-01-01")
	ret.AddressList = ret.AddressList.Append(create_address(state))
	return sifxml.SchoolInfoCreate(ret)
}

func Create_SchoolInfos(count int) []*sifxml.SchoolInfo {
	ret := make([]*sifxml.SchoolInfo, 0)
	for i := 0; i < count; i++ {
		ret = append(ret, Create_SchoolInfo())
	}
	return ret
}

func Create_StudentSchoolEnrollment(student *sifxml.StudentPersonal, school *sifxml.SchoolInfo) *sifxml.StudentSchoolEnrollment {
	ret := sifxml.StudentSchoolEnrollment{}
	ret.Set("RefId", create_GUID())
	ret.Set("MembershipType", "01")
	ret.Set("SchoolYear", this_year())
	ret.Set("TimeFrame", "C")
	ret.Set("FTE", 1.0)
	ret.Set("EntryDate", this_year()+"-01-25")

	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if student != nil {
		ret.CopyString("StudentPersonalRefId", student.RefId)
	}

	ret.YearLevel = ret.YearLevel.New()
	if student != nil && student.MostRecent != nil && student.MostRecent.YearLevel != nil {
		ret.YearLevel.CopyString("Code", student.MostRecent.YearLevel.Code)
	} else {
		ret.YearLevel.Set("Code", strconv.Itoa(rand.Intn(12)+1))
	}

	return sifxml.StudentSchoolEnrollmentCreate(ret)
}

func Create_StudentSchoolEnrollments(students []*sifxml.StudentPersonal, school *sifxml.SchoolInfo) []*sifxml.StudentSchoolEnrollment {
	ret := make([]*sifxml.StudentSchoolEnrollment, 0)
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
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.Set("StateProvinceId", strconv.Itoa(random_seq_gen("stateProvinceId", 99999999)+1))
	ret.Set("EmploymentStatus", "A")
	ret.PersonInfo = ret.PersonInfo.New()
	ret.PersonInfo.Name = ret.PersonInfo.Name.New()
	ret.PersonInfo.Name.Set("Type", "LGL")
	ret.PersonInfo.Name.Set("FamilyName", person.LastName)
	ret.PersonInfo.Name.Set("GivenName", person.FirstName)
	ret.PersonInfo.Name.Set("PreferredGivenName", person.FirstName)
	ret.PersonInfo.Name.Set("MiddleName", gofakeit.FirstName())
	ret.PersonInfo.Name.Set("Title", create_salutation(person.Gender))
	ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.Demographics.Set("Sex", sex_seeded(person.Gender))
	ret.PersonInfo.Demographics.Set("CountryOfBirth", "1101")
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().Set("Type", "01")
	ret.PersonInfo.EmailList.Last().Set("Value", create_email(person.FirstName, middlename, person.LastName))
	return sifxml.StaffPersonalCreate(ret)
}

func Create_StaffPersonals(count int) []*sifxml.StaffPersonal {
	ret := make([]*sifxml.StaffPersonal, 0)
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StaffPersonal())
	}
	return ret
}

func Create_StaffAssignment(staff *sifxml.StaffPersonal, school *sifxml.SchoolInfo) *sifxml.StaffAssignment {
	ret := sifxml.StaffAssignment{}
	ret.Set("RefId", create_GUID())
	ret.Set("PrimaryAssignment", "Y")
	ret.Set("SchoolYear", this_year())
	ret.Set("JobStartDate", "1990-01-01")
	ret.Set("JobFunction", "teacher")
	ret.StaffActivity = ret.StaffActivity.New()
	ret.StaffActivity.Set("Code", randomStringFromSlice(sifxml.AUCodeSetsStaffActivityType_values))

	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if staff != nil {
		ret.CopyString("StaffPersonalRefId", staff.RefId)
	}
	return sifxml.StaffAssignmentCreate(ret)
}

func Create_StaffAssignments(staff []*sifxml.StaffPersonal, school *sifxml.SchoolInfo) []*sifxml.StaffAssignment {
	ret := make([]*sifxml.StaffAssignment, 0)
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
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.PersonInfo = ret.PersonInfo.New()
	ret.PersonInfo.Name = ret.PersonInfo.Name.New()
	ret.PersonInfo.Name.Set("Type", "LGL")
	ret.PersonInfo.Name.Set("FamilyName", person.LastName)
	ret.PersonInfo.Name.Set("GivenName", person.FirstName)
	ret.PersonInfo.Name.Set("PreferredGivenName", person.FirstName)
	ret.PersonInfo.Name.Set("PreferredFamilyName", person.LastName)
	ret.PersonInfo.Name.Set("MiddleName", gofakeit.FirstName())
	ret.PersonInfo.Name.Set("Title", create_salutation(person.Gender))
	ret.PersonInfo.Demographics = ret.PersonInfo.Demographics.New()
	ret.PersonInfo.Demographics.Set("Sex", sex_seeded(person.Gender))
	ret.PersonInfo.Demographics.Set("CountryOfBirth", "1101")
	ret.PersonInfo.Demographics.LanguageList = ret.PersonInfo.Demographics.LanguageList.AddNew()
	ret.PersonInfo.Demographics.LanguageList.Last().Set("Code", "1201")
	ret.PersonInfo.Demographics.LanguageList.Last().Set("LanguageType", "1")
	if rand.Float64() < 0.2 {
		ret.PersonInfo.Demographics.LanguageList = ret.PersonInfo.Demographics.LanguageList.AddNew()
		ret.PersonInfo.Demographics.LanguageList.Last().Set("Code", randomStringFromSlice([]string{"0002", "7101", "2401", "2201", "5203", "4202"}))
		ret.PersonInfo.Demographics.LanguageList.Last().Set("LanguageType", "2")
	}
	ret.PersonInfo.EmailList = ret.PersonInfo.EmailList.AddNew()
	ret.PersonInfo.EmailList.Last().Set("Type", "01")
	ret.PersonInfo.EmailList.Last().Set("Value", create_email(person.FirstName, middlename, person.LastName))
	ret.PersonInfo.PhoneNumberList = ret.PersonInfo.PhoneNumberList.AddNew()
	ret.PersonInfo.PhoneNumberList.Last().Set("Type", "0096")
	ret.PersonInfo.PhoneNumberList.Last().Set("Number", create_phone_number(nil))
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
		ret.Set("EmploymentType", randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.Set("SchoolEducationalLevel", randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.Set("NonSchoolEducation", randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if student != nil && rand.Float64() < 0.8 {
		ret.PersonInfo.Name.CopyString("FamilyName", student.PersonInfo.Name.FamilyName)
		ret.PersonInfo.Name.CopyString("PreferredFamilyName", student.PersonInfo.Name.FamilyName)
	}
	return sifxml.StudentContactPersonalCreate(ret)
}

func Create_StudentContactRelationship(student *sifxml.StudentPersonal, contact *sifxml.StudentContactPersonal) *sifxml.StudentContactRelationship {
	ret := sifxml.StudentContactRelationship{}
	ret.Set("StudentContactRelationshipRefId", create_GUID())
	ret.Relationship = ret.Relationship.New()
	ret.Relationship.Set("Code", threshold_rand_strings([]float64{0.24, 0.22, 0.2, 0.18, 0.16, 0.14, 0.12, 0.10, 0.08, 0.04, 0.02, 0}, []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}))
	ret.ContactFlags = ret.ContactFlags.New()
	ret.ContactFlags.Set("ParentLegalGuardian", threshold_rand_strings([]float64{0.8}, []string{"N", "Y"}))
	ret.ContactFlags.Set("PickupRights", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("LivesWith", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("AccessToRecords", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("ReceivesAssessmentReport", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("EmergencyContact", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("HasCustody", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("DisciplinaryContact", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("AttendanceContact", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("PrimaryCareProvider", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("FeesBilling", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("FeesAccess", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("FamilyMail", threshold_rand_strings([]float64{0.9}, []string{"N", "Y"}))
	ret.ContactFlags.Set("InterventionOrder", threshold_rand_strings([]float64{0.1}, []string{"N", "Y"}))

	if student != nil {
		ret.CopyString("StudentPersonalRefId", student.RefId)
	}
	if contact != nil {
		ret.CopyString("StudentContactPersonalRefId", contact.RefId)
	}
	return sifxml.StudentContactRelationshipCreate(ret)
}

func Create_StudentContactPersonalAndRelationship(students []*sifxml.StudentPersonal) ([]sifxml.StudentContactPersonal, []sifxml.StudentContactRelationship) {
	contacts := make([]sifxml.StudentContactPersonal, 0)
	relationships := make([]sifxml.StudentContactRelationship, 0)
	for _, s := range students {
		p1 := Create_StudentContactPersonal(s, 1)
		if p1 != nil {
			contacts = append(contacts, *p1)
			relationships = append(relationships, *Create_StudentContactRelationship(s, p1))
		}
		p2 := Create_StudentContactPersonal(s, 2)
		if p2 != nil {
			contacts = append(contacts, *p2)
			relationships = append(relationships, *Create_StudentContactRelationship(s, p2))
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
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.Set("RoomNumber", strconv.Itoa(room_number))
	ret.Set("Description", fmt.Sprintf("Room %d", room_number))
	ret.Set("Capacity", rand.Intn(50)+10)
	ret.Set("Size", float64(rand.Intn(5)+2))
	ret.Set("RoomType", randomStringFromSlice([]string{"Classroom", "Classroom", "Classroom", "Classroom", "Classroom", "Classroom", "Art", "Basketball court"}))
	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
	}
	if len(staff) > 1 {
		ret.StaffList = ret.StaffList.New()
		for _, s1 := range staff {
			ret.StaffList = ret.StaffList.AppendString(s1.RefId)
		}
	}
	return sifxml.RoomInfoCreate(ret)
}

func Create_RoomInfos(count int, school *sifxml.SchoolInfo) []*sifxml.RoomInfo {
	ret := make([]*sifxml.RoomInfo, 0)
	for i := 0; i < count; i++ {
		ret = append(ret, Create_RoomInfo(school))
	}
	return ret
}

func Create_TeachingGroup(school *sifxml.SchoolInfo, students []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal) *sifxml.TeachingGroup {
	shortname := randomStringFromSlice([]string{"MAT", "ENG", "PHYS", "BIO", "CHEM", "COMP", "VIS", "ECON", "HIST"})

	ret := sifxml.TeachingGroup{}
	ret.Set("RefId", create_GUID())
	ret.Set("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.Set("ShortName", shortname)
	ret.Set("LongName", teachingGroupLongName(shortname))
	ret.Set("KeyLearningArea", teachingGroupKLA(shortname))
	ret.Set("SchoolYear", this_year())
	ret.Set("Semester", 1)
	ret.Set("MinClassSize", 20)
	ret.Set("MaxClassSize", 40)

	if school != nil {
		ret.CopyString("SchoolInfoRefId", school.RefId)
		ret.CopyString("SchoolLocalId", school.LocalId)
	}
	if len(students) > 1 {
		ret.StudentList = ret.StudentList.New()
		for _, s := range students {
			ret.StudentList = ret.StudentList.AddNew()
			ret.StudentList.Last().CopyString("StudentPersonalRefId", s.RefId)
			ret.StudentList.Last().CopyString("StudentLocalId", s.LocalId)
			ret.StudentList.Last().Set("Name", s.PersonInfo.Name.Clone())
		}
	}
	if len(staff) > 1 {
		ret.TeacherList = ret.TeacherList.New()
		for _, s := range staff {
			ret.TeacherList = ret.TeacherList.AddNew()
			ret.TeacherList.Last().CopyString("StaffPersonalRefId", s.RefId)
			ret.TeacherList.Last().CopyString("StaffLocalId", s.LocalId)
			ret.TeacherList.Last().Set("Name", s.PersonInfo.Name.Clone())
		}
	}
	return sifxml.TeachingGroupCreate(ret)
}
