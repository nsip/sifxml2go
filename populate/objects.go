package populate

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/nsip/sifxml2go/sifxml"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/au"
)

// Populate an address. Used in Create_SchoolInfo(), Create_StudentContactPersonal(), Create_ChargedLocationInfo().
//
// If state (StateProvince) is not provided, it is selected at random from valid Australian states.
//
// * PostalCode is random based on state.
//
// * Type and Role are set at "0123" (Mailing Address) and "012A" (Home Address).
//
// * StreetNumber, StreetName, City are randomised.
//
// * StreetType is a random selection from a list of typical Australian street types.
func Create_Address(state string) sifxml.AddressType {
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
	ret.Street().SetProperty("StreetType",
		randomStringFromSlice([]string{"Avenue", "Boulevard", "Cove", "Court", "Crescent", "Drive",
			"Esplanade", "Lane", "Place", "Road", "Square", "Street", "Terrace", "Walk", "Way"}))
	ret.SetProperty("City", gofakeit.City())
	ret.SetProperty("StateProvince", state)
	ret.SetProperty("PostalCode", postcode)
	return ret
}

// Add a teacher to a list of scheduled teachers (SIF type sifxml.ScheduledTeacherListType,
// appearing in ScheduledActivty or TimeTableCell),
// with a starting time and finish time for their scheduled teaching.
//
// * Weighting is set at 1.0.
//
// * Supervision is set with weighted probability: "Normal" (0.8), "MergedClass" (0.1), "MinimalSupervision" (0.1).
//
// * Credit is only set with probability 0.4. If it is set, its value is set randomly as one of "Casual", "Extra", "In-Lieu", "Underload".
func AddToScheduledTeacherList(t *sifxml.ScheduledTeacherListType, staff []*sifxml.StaffPersonal,
	start time.Time, finish time.Time) *sifxml.ScheduledTeacherListType {
	for _, s := range staff {
		t = t.AddNew()
		t.Last().SetProperty("StaffPersonalRefId", s.RefId().String())
		t.Last().SetProperty("StaffLocalId", s.LocalId().String())
		t.Last().SetProperty("StartTime", start.Format("15:04:05"))
		t.Last().SetProperty("FinishTime", finish.Format("15:04:05"))
		t.Last().SetProperty("Weighting", 1.0)
		if rand.Float64() > 0.6 {
			t.Last().SetProperty("Credit", threshold_rand_strings([]float64{0.75, 0.5, 0.25, 0},
				[]string{"Casual", "Extra", "In-Lieu", "Underload"}))
		}
		t.Last().SetProperty("Supervision", threshold_rand_strings([]float64{0.9, 0.8, 0},
			[]string{"MergedClass", "MinimalSupervision", "Normal"}))
	}
	return t
}

func copyTeachingGroupPeriodFromCell(group *sifxml.TeachingGroup, cell *sifxml.TimeTableCell,
	start_times map[string]map[string]string) *sifxml.TeachingGroup {
	group.TeachingGroupPeriodList().AddNew()
	group.TeachingGroupPeriodList().Last().SetProperty("TimeTableCellRefId", cell.RefId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("RoomNumber", cell.RoomNumber().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StaffLocalId", cell.StaffLocalId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("DayId", cell.DayId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("PeriodId", cell.PeriodId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StartTime", cell.RefId().String())
	group.TeachingGroupPeriodList().Last().SetProperty("CellType", cell.CellType().String())
	group.TeachingGroupPeriodList().Last().SetProperty("StartTime",
		start_times[cell.DayId().String()][cell.PeriodId().String()])
	return group
}

// Create a StudentPersonal record, with the student enrolled in the given year level. If no year level is provided,
// generate a random value.
//
// The student has a parent 1 with probability 0.8, and has a parent 2 with probability 0.8, or if parent 1 is not provided.
//
// Only a legal name (Type: "LGL") is provided for the student.
//
// * FamilyName, GivenName, Sex are randomised as a bundle (https://github.com/brianvoe/gofakeit). Sex is binary.
//
// * MiddleName is randomised separately (gofakeit does not support middle names), and may not be of the same gender as the GivenName.
//
// * Email is only a single entry, of type "01" (Primary), and is generated using the first name, the middle initial, the surname, and the domain "example.edu.au". (Overridden in Create_StudentSchoolEnrollment().)
//
// * LocalId is a sequence number shared between all objects.
//
// * StateProvinceId is a unique random number between 1 and 100000000.
//
// * BirthDate is a random date consistent with the year level and the current year.
//
// * IndigenousStatus is set with weighted probability: "4" = "Neither Aboriginal nor Torres Strait Islander origin" (0.8), "1" = "Aboriginal but not Torres Strait Islander origin" (0.05), "2" = "Torres Strait Islander but not Aboriginal origin" (0.05), "3" = "Both Aboriginal and Torres Strait Islander origin" (0.05), "9" = "Not stated/unknown" (0.05).
//
// * CountryOfBirth is fixed at "1101" (Australia).
//
// * Parent1Language and Parent2Language are fixed at "1201" (English).
//
// * Parent1EmploymentType, Parent1SchoolEducationLevel, Parent1NonSchoolEducation, Parent2EmploymentType, Parent2SchoolEducationLevel, Parent2NonSchoolEducation are set randomly to a valid value.
//
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
	ret.PersonInfo().Demographics().SetProperty("IndigenousStatus",
		threshold_rand_strings([]float64{0.2, 0.15, 0.1, 0.05, 0}, []string{"4", "1", "2", "3", "9"}))
	ret.PersonInfo().Demographics().SetProperty("CountryOfBirth", "1101")
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value",
		create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
	if hasParent1 {
		ret.MostRecent().SetProperty("Parent1Language", "1201")
		ret.MostRecent().SetProperty("Parent1EmploymentType",
			randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent().SetProperty("Parent1SchoolEducationLevel",
			randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent().SetProperty("Parent1NonSchoolEducation",
			randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if hasParent2 {
		ret.MostRecent().SetProperty("Parent2Language", "1201")
		ret.MostRecent().SetProperty("Parent2EmploymentType",
			randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.MostRecent().SetProperty("Parent2SchoolEducationLevel",
			randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.MostRecent().SetProperty("Parent2NonSchoolEducation",
			randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	ret.MostRecent().YearLevel().SetProperty("Code", yearlevel)
	if out, ok := sifxml.StudentPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create count StudentPersonal records; each student is assigned a random member of yearlevels as their year level.
func Create_StudentPersonals(count int, yearlevels []string) []*sifxml.StudentPersonal {
	ret := sifxml.StudentPersonalSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StudentPersonal(randomStringFromSlice(yearlevels)))
	}
	return ret
}

// Create a school with its school of SchoolType schooltype. If schooltype is empty, a random value will be selected (drawn from the codeset AUCodeSetsSchoolLevelType).
//
// Each school has a campus type populated. The SchoolCampusId is a random number between 1 and 4. The ParentSchoolId is the same as the school LocalId. The CampusType is the same as the SchoolType. The AdminStatus is set with weighted probability: "N" (0.2), "Y" (0.8).
//
// * LocalId is a sequence number shared between all objects.
//
// * StateProvinceId is a unique random number between 1 and 10000.
//
// * CommonwealthId is a unique random number between 1 and 10000.
//
// * SchoolName is a random name, suffixed with a descriptor consistent with schooltype.
//
// * SchoolSector is set to "Gov".
//
// * OperationalStatus is set to "O".
//
// * IndependentSchool is set to "N".
//
// * ARIA is set to 1.0.
//
// * Entity_Open is set to "1990-01-01".
//
// * AddressList is populated by a single entry from Create_Address().
//
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
	ret.AddressList().Append(Create_Address(state))
	if out, ok := sifxml.SchoolInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SchoolInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Creates count SchoolInfo objects, with random school types.
func Create_SchoolInfos(count int) []*sifxml.SchoolInfo {
	ret := sifxml.SchoolInfoSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_SchoolInfo(""))
	}
	return ret
}

// Create a StudentSchoolEnrollment object, linking StudentPersonal object student to SchoolInfo object school.
//
// Updates student: the student email is set to a domain value based on the school name.
//
// * MembershipType is set to "01".
//
// * SchoolYear is set to the current year.
//
// * TimeFrame is set to "C".
//
// * FTE is set to 1.0.
//
// * EntryDate is set to 25 January of the current year.
//
// * If the student year level is set, YearLevel is set to the same value; otherwise it is set to a random value between 1 and 12.
//
func Create_StudentSchoolEnrollment(student *sifxml.StudentPersonal,
	school *sifxml.SchoolInfo) *sifxml.StudentSchoolEnrollment {
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
		student.PersonInfo().EmailList().Last().SetProperty("Value",
			create_email(student.PersonInfo().Name().GivenName().String(),
				student.PersonInfo().Name().MiddleName().String(),
				student.PersonInfo().Name().FamilyName().String(), create_school_email_domain(school)))
	}

	if out, ok := sifxml.StudentSchoolEnrollmentPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentSchoolEnrollment: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Creates one StudentSchoolEnrollment for each student in students, linking it to school.
func Create_StudentSchoolEnrollments(students []*sifxml.StudentPersonal,
	school *sifxml.SchoolInfo) []*sifxml.StudentSchoolEnrollment {
	ret := sifxml.StudentSchoolEnrollmentSlice()
	for _, s := range students {
		ret = append(ret, Create_StudentSchoolEnrollment(s, school))
	}
	return ret
}

// Create a StaffPersonal record.
//
// Only a legal name (Type: "LGL") is provided for the staff member.
//
// * FamilyName, GivenName, Sex are randomised as a bundle (https://github.com/brianvoe/gofakeit). Sex is binary.
//
// * MiddleName is randomised separately (gofakeit does not support middle names), and may not be of the same gender as the GivenName.
//
// * PreferredGivenName is set to be the same as GivenName.
//
// * Title is randomised, based on the staff sex.
//
// * The staff is assigned two OtherId identifiers: One of type "DET_USER_ID",  a unique random number between 1 and 100000000,
// and one of type either "pep" (probability 0.9) or "cep" (probability 0.1), also a unique random number between 1 and 100000000.
//
// * Email is only a single entry, of type "01" (Primary), and is generated using the first name, the middle initial, the surname, and the domain "example.edu.au". (Overridden in Create_StaffAssignment().)
//
// * LocalId is a sequence number shared between all objects.
//
// * StateProvinceId is a unique random number between 1 and 100000000.
//
// * EmploymentStatus is set to "A".
//
// * BirthDate is a random date betweem 25 and 65 years before the current year.
//
// * CountryOfBirth is fixed at "1101" (Australia).
//
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
	ret.PersonInfo().Demographics().SetProperty("BirthDate",
		random_date(strconv.Itoa(time.Now().Year()-65)+"-01-01", strconv.Itoa(time.Now().Year()-25)+"-12-31"))
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value",
		create_email(person.FirstName, middlename, person.LastName, "example.edu.au"))
	if out, ok := sifxml.StaffPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StaffPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create count StaffPersonal objects.
func Create_StaffPersonals(count int) []*sifxml.StaffPersonal {
	ret := sifxml.StaffPersonalSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_StaffPersonal())
	}
	return ret
}

// Create a StaffAssignment object, linking StaffPersonal object staff to SchoolInfo object school.
//
// Updates student: the staff email is set to a domain value based on the school name.
//
// * PrimaryAssignment is set to "Y".
//
// * SchoolYear is set to the current year.
//
// * JobStartDate is set to "1990-01-01".
//
// * JobFunction is set to "teacher"
//
// * StaffActivity is set to a random valid value.
//
// * If the student year level is set, YearLevel is set to the same value; otherwise it is set to a random value between 1 and 12.
//
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
		staff.PersonInfo().EmailList().Last().SetProperty("Value",
			create_email(staff.PersonInfo().Name().GivenName().String(),
				staff.PersonInfo().Name().MiddleName().String(),
				staff.PersonInfo().Name().FamilyName().String(), create_school_email_domain(school)))
	}
	if out, ok := sifxml.StaffAssignmentPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StaffAssignment: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Creates one StaffAssignment for each staff member in staff, linking it to school.
func Create_StaffAssignments(staff []*sifxml.StaffPersonal, school *sifxml.SchoolInfo) []*sifxml.StaffAssignment {
	ret := sifxml.StaffAssignmentSlice()
	for _, s := range staff {
		ret = append(ret, Create_StaffAssignment(s, school))
	}
	return ret
}

// Create a StudentContactPersonal record for a student contact corresponding to StudentPersonal object student, with the Parent1/Parent2
// status set by ordinal.
//
// If ordinal == 1, the EmploymentType, SchoolEducationalLevel and NonSchoolEducation of the contact are set to match the Parent1 values
// under student. If ordinal == 2, the EmploymentType, SchoolEducationalLevel and NonSchoolEducation of the contact are set to match the Parent2 values
// under student. If ordinal is neither, the EmploymentType, SchoolEducationalLevel and NonSchoolEducation of the contact are set to random valid values.
//
// Only a legal name (Type: "LGL") is provided for the contact.
//
// * FamilyName, GivenName, Sex are randomised as a bundle (https://github.com/brianvoe/gofakeit). Sex is binary. With probability 0.8, the family name
// is updated to match that of the student.
//
// * MiddleName is randomised separately (gofakeit does not support middle names), and may not be of the same gender as the GivenName.
//
// * PreferredGivenName is set to the same as GivenName.
//
// * PreferredFamilyName is set to the same as FamilyName.
//
// * Title is randomised, based on the contact sex.
//
// * EmailList is a single entry, and is of type "01" (Primary), and is generated using the first name, the middle initial, the surname, and a random commercial domain.
//
// * LocalId is a sequence number shared between all objects.
//
// * CountryOfBirth is fixed at "1101" (Australia).
//
// * AddressList is a single entry set by Create_Address(). If the student has an address, the address is set to be in the same state.
//
// * Phonelist is a single entry of Type "0096" (Main Telephone Number), and is a mobile phone number.
//
// * LanguageList has one entry of Type "1" (Main Language Spoken At Home), set to "1201" (English).
//
// With probability 0.2, a second language is added to LanguageList, of Type "2" (Main Language Other Than English Spoken at Home),
// set to a random selection of "0002", "7101", "2401", "2201", "5203", "4202" (Not Stated, Cantonese, Italian, Greek, Hindi, Arabic).
//
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
		ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("Code",
			randomStringFromSlice([]string{"0002", "7101", "2401", "2201", "5203", "4202"}))
		ret.PersonInfo().Demographics().LanguageList().Last().SetProperty("LanguageType", "2")
	}
	ret.PersonInfo().EmailList().AddNew()
	ret.PersonInfo().EmailList().Last().SetProperty("Type", "01")
	ret.PersonInfo().EmailList().Last().SetProperty("Value",
		create_email(person.FirstName, middlename, person.LastName, create_commercial_email_domain()))
	ret.PersonInfo().PhoneNumberList().AddNew()
	ret.PersonInfo().PhoneNumberList().Last().SetProperty("Type", "0096")
	ret.PersonInfo().PhoneNumberList().Last().SetProperty("Number", create_phone_number(nil))
	ret.PersonInfo().AddressList().Append(Create_Address(state))
	if student != nil && ordinal == 1 {
		ret.SetProperty("EmploymentType", student.MostRecent().Parent1EmploymentType().String())
		ret.SetProperty("SchoolEducationalLevel", student.MostRecent().Parent1SchoolEducationLevel().String())
		ret.SetProperty("NonSchoolEducation", student.MostRecent().Parent1NonSchoolEducation().String())
	} else if student != nil && ordinal == 2 {
		ret.SetProperty("EmploymentType", student.MostRecent().Parent2EmploymentType().String())
		ret.SetProperty("SchoolEducationalLevel", student.MostRecent().Parent2SchoolEducationLevel().String())
		ret.SetProperty("NonSchoolEducation", student.MostRecent().Parent2NonSchoolEducation().String())
	} else {
		ret.SetProperty("EmploymentType",
			randomStringFromSlice(sifxml.AUCodeSetsEmploymentTypeType_values))
		ret.SetProperty("SchoolEducationalLevel",
			randomStringFromSlice(sifxml.AUCodeSetsSchoolEducationLevelTypeType_values))
		ret.SetProperty("NonSchoolEducation",
			randomStringFromSlice(sifxml.AUCodeSetsNonSchoolEducationType_values))
	}
	if student != nil && rand.Float64() < 0.8 {
		ret.PersonInfo().Name().SetProperty("FamilyName", student.PersonInfo().Name().FamilyName().String())
		ret.PersonInfo().Name().SetProperty("PreferredFamilyName",
			student.PersonInfo().Name().FamilyName().String())
	}
	if out, ok := sifxml.StudentContactPersonalPointer(ret); !ok {
		log.Fatalf("Could not create pointer to StudentContactPersonal: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create a StudentContactRelationship record joining StudentPersonal object student to StudentContactPersonal object contact.
//
// * Relationship is set with weighted probability: "01" (Parent) (0.26), and 0.02 for each other value between "02" and "13" inclusive.
//
// * Each flag in ContactFlags is set to "Y" with probability 0.9, else "N". Exceptions: "ParentLegalGuardian" is set to "Y" with probability 0.8; "InterventionOrder" is set to "N" with probability 0.9.
func Create_StudentContactRelationship(student *sifxml.StudentPersonal, contact *sifxml.StudentContactPersonal) *sifxml.StudentContactRelationship {
	ret := sifxml.StudentContactRelationship{}
	ret.SetProperty("StudentContactRelationshipRefId", create_GUID())
	ret.Relationship().SetProperty("Code", threshold_rand_strings(
		[]float64{0.26, 0.24, 0.22, 0.2, 0.18, 0.16, 0.14, 0.12, 0.10, 0.08, 0.04, 0.02, 0},
		[]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}))
	ret.ContactFlags().SetProperty("ParentLegalGuardian",
		threshold_rand_strings([]float64{0.8, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("PickupRights",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("LivesWith",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("AccessToRecords",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("ReceivesAssessmentReport",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("EmergencyContact",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("HasCustody",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("DisciplinaryContact",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("AttendanceContact",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("PrimaryCareProvider",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FeesBilling",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FeesAccess",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("FamilyMail",
		threshold_rand_strings([]float64{0.9, 0}, []string{"N", "Y"}))
	ret.ContactFlags().SetProperty("InterventionOrder",
		threshold_rand_strings([]float64{0.1, 0}, []string{"N", "Y"}))

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

// Create one or two StudentContactPersonal objects and StudentContactRelationship objects for each student in students. Whether one or two contacts and relationships
// are created depends on the encoding of Parent1 and Parent2 attributes in the student object. Each student has distinct parents: there is no modelling of siblings.
func Create_StudentContactPersonalAndRelationship(students []*sifxml.StudentPersonal) ([]*sifxml.StudentContactPersonal,
	[]*sifxml.StudentContactRelationship) {
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

// Create a RoomInfo object linked to SchoolInfo object school, with no associated staff.
func Create_RoomInfo(school *sifxml.SchoolInfo) *sifxml.RoomInfo {
	return Create_RoomInfoWithStaff(school, []*sifxml.StaffPersonal{})
}

// Create a RoomInfo object linked to SchoolInfo object school, with associated StaffPersonal objects given in staff (populating StaffList).
//
// * LocalId is a sequence number shared between all objects.
//
// * RoomNumber is a unique random number between 1 and 999.
//
// * Description is "Room" plus RoomNumber.
//
// * Capacity is a random integer between 10 and 60.
//
// * Size is a random integer between  2 and 7.
//
// * RoomType is set with weighted probability: "Classroom" (0.8), "Art" (0.1), "Basketball court" (0.1)
func Create_RoomInfoWithStaff(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal) *sifxml.RoomInfo {
	room_number := random_seq_gen("roomNumber", 998) + 1

	ret := sifxml.RoomInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("RoomNumber", strconv.Itoa(room_number))
	ret.SetProperty("Description", fmt.Sprintf("Room %d", room_number))
	ret.SetProperty("Capacity", rand.Intn(50)+10)
	ret.SetProperty("Size", float64(rand.Intn(5)+2))
	ret.SetProperty("RoomType", threshold_rand_strings([]float64{0.2, 0.1, 0},
		[]string{"Classroom", "Art", "Basketball court"}))
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

// Create count RoomInfo objects, linked to SchoolInfo object school.
func Create_RoomInfos(count int, school *sifxml.SchoolInfo) []*sifxml.RoomInfo {
	ret := sifxml.RoomInfoSlice()
	for i := 0; i < count; i++ {
		random_seq_gen_reset("roomNumber")
		ret = append(ret, Create_RoomInfo(school))
	}
	return ret
}

// Create a TeachingGroup object, linked to StudentPersonal objects students (StudentList), StaffPersonal objects staff (TeacherList), and TimeTableSubject timetablesubject.
//
// * LocalId is a sequence number shared between all objects.
//
// * SchoolYear is the current year.
//
// * Semester is set to 1.
//
// * MinClassSize is set to 20.
//
// * MaxClassSize is set to 40.
//
// * ShortName is set to timetablesubject SubjectShortName if present; otherwise it is a random choice of subject out of the defaults on offer (All_teachingSubjects())
//
// * LongName and KeyLearningArea are the full name and key learning area associated with ShortName (TeachingGroupKLA()).
//
func Create_TeachingGroup(school *sifxml.SchoolInfo, students []*sifxml.StudentPersonal,
	staff []*sifxml.StaffPersonal, timetablesubject *sifxml.TimeTableSubject) *sifxml.TeachingGroup {
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
	ret.SetProperty("KeyLearningArea", TeachingGroupKLA(shortname))

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

// Create a FinancialAccount object, linked to parent FinancialAccount object parent (if it is not nil),
// and ChargedLocationInfo object location (if it is not nil)
//
// * LocalId is a sequence number shared between all objects.
//
// * CreationDate is a random date between 2012-01-01 and 2015-12-31.
//
// * CreationTime is a random time.
//
// * AccountNumber is a unique random number between 1 and 100000000
//
// * ClassType is randomly set to one of "Asset", "Liability", "Revenue", "Expense", "Other".
//
// * Name is set to a random name. If location is not nil, it is instead set to the same Name as Location.Name.
//
func Create_FinancialAccount(parent *sifxml.FinancialAccount, location *sifxml.ChargedLocationInfo) *sifxml.FinancialAccount {
	gofakeit.Seed(0)
	ret := sifxml.FinancialAccount{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
	ret.SetProperty("CreationDate", random_date("2012-01-01", "2015-12-31"))
	ret.SetProperty("CreationTime", gofakeit.Date().Format("15:04:05"))
	ret.SetProperty("AccountNumber", strconv.Itoa(random_seq_gen("financial_account_number", 99999999)+1))
	ret.SetProperty("ClassType",
		randomStringFromSlice([]string{"Asset", "Liability", "Revenue", "Expense", "Other"}))

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

func create_FinancialAccountsBase(count int, parent []*sifxml.FinancialAccount, location []*sifxml.ChargedLocationInfo) []*sifxml.FinancialAccount {
	if len(parent) != 0 && len(location) != 0 {
		if len(parent) != count {
			log.Fatalf("Mismatch in count of financial accounts (%d) and count of parent financial accounts (%d)",
				count, len(parent))
		}
		if len(location) != count {
			log.Fatalf("Mismatch in count of financial accounts (%d) and count of charged locations (%d)",
				count, len(location))
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

// Create count FinancialAccount objects, some of which shall be linked to ChargedLocationInfo objects randomly drawn
// from locations.
//
// Half the accounts to be generated are children of other accounts generated in the same function, chosen randomly.
// A third of all parent accounts are randomly associated with a charged location;
// all children of that account are associated with the same charged location.
//
func Create_FinancialAccounts(count int, locations []*sifxml.ChargedLocationInfo) []*sifxml.FinancialAccount {
	parent_account_count := count/2 + 1
	parent_charge_locations := sifxml.ChargedLocationInfoSlice()
	for i := 0; i < parent_account_count; i++ {
		if rand.Float32() > 0.333 {
			parent_charge_locations = append(parent_charge_locations, locations[rand.Intn(len(locations))])
		} else {
			parent_charge_locations = append(parent_charge_locations, nil)
		}
	}
	parent_accounts := create_FinancialAccountsBase(parent_account_count, sifxml.FinancialAccountSlice(),
		parent_charge_locations)

	child_account_count := count - parent_account_count
	child_charge_locations := sifxml.ChargedLocationInfoSlice()
	child_parent_accounts := sifxml.FinancialAccountSlice()
	for i := 0; i < child_account_count; i++ {
		parent_id := rand.Intn(len(parent_accounts))
		child_parent_accounts = append(child_parent_accounts, parent_accounts[parent_id])
		child_charge_locations = append(child_charge_locations, parent_charge_locations[parent_id])
	}
	child_accounts := create_FinancialAccountsBase(child_account_count, child_parent_accounts,
		child_charge_locations)
	return append(parent_accounts, child_accounts...)
}

// Create a ChargedLocationInfo object, linked to parent ChargedLocationInfo object parent (if it is not nil),
// and SchoolInfo object school (if it is not nil)
//
// If school is present:
//
// * SiteCategory is set to "School"
//
// * LocationType is set to "School"
//
// * Name, LocalId, StateProvinceId, AddressList are copied from school.
//
// * PhoneNumberList is copied from school if present.
//
// If school is not present:
//
// * SiteCategory is set to "NonSchool"
//
// * LocationType is set to a random choice out of "HR", "Professional Development", "Accounting",
// "Management", "Cleaning"
//
// * Name is set to a radndom company name followed by LocationType.
//
// * LocalId is a sequence number shared between all objects.
//
// * AddressList is a single entry set by Create_Address(), in a random state.
//
// * PhoneNumberList is a single entry of Type "0096" (Main Telephone Number), and is a phone number
// with a landline corresponding to the address state.
//
func Create_ChargedLocationInfo(parent *sifxml.ChargedLocationInfo,
	school *sifxml.SchoolInfo) *sifxml.ChargedLocationInfo {
	state := create_state()

	ret := sifxml.ChargedLocationInfo{}
	ret.SetProperty("RefId", create_GUID())
	if parent != nil {
		ret.SetProperty("ParentChargedLocationInfoRefId", parent.RefId().String())
	}
	if school == nil {
		gofakeit.Seed(0)
		locationtype := randomStringFromSlice([]string{"HR", "Professional Development", "Accounting",
			"Management", "Cleaning"})
		ret.SetProperty("SiteCategory", "NonSchool")
		ret.SetProperty("LocationType", locationtype)
		ret.SetProperty("Name", gofakeit.Company()+" "+locationtype)
		ret.SetProperty("LocalId", strconv.Itoa(seq_gen("localId")))
		ret.AddressList().Append(Create_Address(state))
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

// Create count ChargedLocationInfo objects. The slice schools gives SchoolInfo objects that are to be represented
// in the generated ChargedLocationInfo objects, one per school. If count is larger than the length of schools,
// the remaining objects are non-school charged locations. Half of the remaining objects generated are parents
// of the other half, assigned randomly.
//
func Create_ChargedLocationInfos(count int, schools []*sifxml.SchoolInfo) []*sifxml.ChargedLocationInfo {
	ret := sifxml.ChargedLocationInfoSlice()
	for i := 0; i < count && i < len(schools); i++ {
		ret = append(ret, Create_ChargedLocationInfo(nil, schools[i]))
	}
	nonschoolcount := count - len(schools)
	if nonschoolcount > 0 {
		for i := 0; i < nonschoolcount/2; i++ {
			ret = append(ret, Create_ChargedLocationInfo(nil, nil))
		}
		for i := nonschoolcount / 2; i < nonschoolcount; i++ {
			ret = append(ret, Create_ChargedLocationInfo(ret[len(schools)+rand.Intn(nonschoolcount/2)], nil))
		}
	}
	return ret
}

// Create a VendorInfo object.
//
// * LocalId is a sequence number shared between all objects.
//
// * Name is a random company name, followed by a random corporate suffix ("Company", "Pty Ltd", "Ltd", "Pty", "Inc").
//
// * CustomerId is a unique random number between 1000 and 100000.
//
// * ABN is a random number between 1000000000 and 100000000000.
//
// * RegisteredForGST is set to "Y".
//
// * PaymentTerms is set to "15 days".
//
// * AccountNumber is set to a unique random number between 10000 and 1000000.
//
// * AccountName is set to be the same as Name.
//
// * In ContactInfo, FamilyName and GivenName are randomised as a bundle (https://github.com/brianvoe/gofakeit).
// PositionTitle is set to "Sales", and Role is set to "Sales".
//
// * EmailList is a single entry, and is of type "01" (Primary), and is generated using the first name, the middle initial,
// the surname, and a random commercial domain.
//
// * PhoneNumberList is a single entry of Type "0096" (Main Telephone Number), and is a phone number
// with a landline corresponding to the address state.
//
func Create_VendorInfo() *sifxml.VendorInfo {
	gofakeit.Seed(0)
	person := gofakeit.Person()
	companytype := randomStringFromSlice([]string{"Company", "Pty Ltd", "Ltd", "Pty", "Inc"})
	companyname := gofakeit.Company()
	middlename := gofakeit.FirstName()
	emaildomain := strings.ToLower(companyname) + "." +
		randomStringFromSlice([]string{"com.au", "com", "com.au", "org.au"})

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
	ret.SetProperty("AccountName", ret.Name().String())

	ret.ContactInfo().Name().SetProperty("Type", "LGL")
	ret.ContactInfo().Name().SetProperty("FamilyName", person.LastName)
	ret.ContactInfo().Name().SetProperty("GivenName", person.FirstName)
	ret.ContactInfo().SetProperty("PositionTitle", "Sales")
	ret.ContactInfo().SetProperty("Role", "Sales")

	ret.ContactInfo().EmailList().AddNew()
	ret.ContactInfo().EmailList().Last().SetProperty("Type", "01")
	ret.ContactInfo().EmailList().Last().SetProperty("Value",
		create_email(person.FirstName, middlename, person.LastName, emaildomain))
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

// Create count VendorInfo objects.
func Create_VendorInfos(count int) []*sifxml.VendorInfo {
	ret := sifxml.VendorInfoSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_VendorInfo())
	}
	return ret
}

// Create a ScheduledActivity object, linked to SchoolInfo object school, TimeTable object timetable,
// TimeTableCell object timetablecell, TimeTableSubject object timetablesubject, slice of StudentPersonal
// objects students, slice of StaffPersonal objects staff, Slice of TeachingGroup objects teachinggroups,
// and slice of RoomInfo objects rooms. This object is hard-coded to generate excursions, and is not the
// object used for use case object generation.
//
// * ActivityDate is set to a random date in the current year.
//
// * StartTime is set to a random time.
//
// * FinishTime is set to three hours after StartTime.
//
// * CellType is set to "Excursion".
//
// * ActivityType is set to "Excursion".
//
// * Location is set to "Zoo".
//
// * ActivityName is set to "Zoo Excursion".
//
// * TeacherList is set through AddToScheduledTeacherList().
//
func Create_ScheduledActivity(school *sifxml.SchoolInfo, timetable *sifxml.TimeTable,
	timetablecell *sifxml.TimeTableCell, timetablesubject *sifxml.TimeTableSubject,
	students []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal, teachinggroups []*sifxml.TeachingGroup,
	rooms []*sifxml.RoomInfo) *sifxml.ScheduledActivity {
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
		AddToScheduledTeacherList(ret.TeacherList(), staff, starttime, finishtime)
	}
	if out, ok := sifxml.ScheduledActivityPointer(ret); !ok {
		log.Fatalf("Could not create pointer to ScheduledActivity: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create a CalendarSummary object, linked to SchoolInfo object school.
//
// * SchoolYear is set to the current year.
//
// * LocalId is a sequence number shared between all objects.
//
// * DaysInSession is set to 277.
//
// * StartDate is set to January 28 of the current year. No attempt is made to set it to a weekday.
//
// * EndDate is set to December 19 of the current year. No attempt is made to set it to a weekday.
//
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

// Create a GradingAssignment object, linked to SchoolInfo object school, TeachingGroup object teachinggroup,
// and slice of StudentPersonal objects students.
//
// * GradingCategory is randomly selected from "quiz", "essay", "project".
//
// * PointsPossible is set at 10.
//
// * CreateDate is set at a random date between February 1 and November 1 in the current year.
//
// * DueDate is set to 30 days after CreateDate.
//
// * Description is set to a randomised string of eight letters.
//
// * Weight is set to a random integer between 2 and 5 inclusive.
//
// * MaxAttemptsAllowed is set to 5.
//
// * DetailedDescriptionURL is set to "http://www.example.com/" followed by Description.
//
func Create_GradingAssignment(school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup,
	students []*sifxml.StudentPersonal) *sifxml.GradingAssignment {
	gofakeit.Seed(0)
	createdate := random_date(this_year()+"-02-01", this_year()+"-11-01")
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

// Create count GradingAssignment objects, each linked to SchoolInfo object school, TeachingGroup object teachinggroup,
// and slice of StudentPersonal objects students.
//
func Create_GradingAssignments(count int, school *sifxml.SchoolInfo, teachinggroup *sifxml.TeachingGroup,
	students []*sifxml.StudentPersonal) []*sifxml.GradingAssignment {
	ret := sifxml.GradingAssignmentSlice()
	for i := 0; i < count; i++ {
		ret = append(ret, Create_GradingAssignment(school, teachinggroup, students))
	}
	return ret
}

// Create a GradingAssignmentScore object, linked to GradingAssignment object assignment, SchoolInfo object school,
// TeachingGroup object teachinggroup, StudentPersonal object student, and StaffPersonalObject staff.
//
// * ScorePoints is set to a random number between 0 and GradingAssignment.PointsPossible.
//
// * ScoreDescription is set to a randomised sentence of 10 words.
//
// * DateGraded is set to a date between 0 and 6 days inclusive after GradingAssignment.DueDate.
//
func Create_GradingAssignmentScore(assignment *sifxml.GradingAssignment, school *sifxml.SchoolInfo,
	teachinggroup *sifxml.TeachingGroup, student *sifxml.StudentPersonal,
	staff *sifxml.StaffPersonal) *sifxml.GradingAssignmentScore {
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

// Create a GradingAssignmentScore object for each student in the slice of StudentPersonal objects students,
// each of them linked to GradingAssignment object assignment, SchoolInfo object school,
// TeachingGroup object teachinggroup, and StaffPersonalObject staff.
//
func Create_GradingAssignmentScores(assignment *sifxml.GradingAssignment, school *sifxml.SchoolInfo,
	teachinggroup *sifxml.TeachingGroup, students []*sifxml.StudentPersonal,
	staff *sifxml.StaffPersonal) []*sifxml.GradingAssignmentScore {
	ret := sifxml.GradingAssignmentScoreSlice()
	for _, s := range students {
		ret = append(ret, Create_GradingAssignmentScore(assignment, school, teachinggroup, s, staff))
	}
	return ret
}

// Create a Debtor object, associated with one of StudentPersonal object student, StaffPersonal object staff,
// StudentContactPersonal object contact, or VendorInfo object vendor. The options are mutually exclusive,
// and all other arguments are expected to be nil.
//
// * LocalId is a sequence number shared between all objects.
//
func Create_Debtor(student *sifxml.StudentPersonal, staff *sifxml.StaffPersonal,
	contact *sifxml.StudentContactPersonal, vendor *sifxml.VendorInfo) *sifxml.Debtor {
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

// Create a Debtor object corresponding to each object in each of the slices of objects passed to the function.
func Create_Debtors(student []*sifxml.StudentPersonal, staff []*sifxml.StaffPersonal,
	contact []*sifxml.StudentContactPersonal, vendor []*sifxml.VendorInfo) []*sifxml.Debtor {
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

// Create a CalendarDate object linked to CalendarSummary object calendar and SchoolInfo object school,
// for the given date (date). Indicate whether the date is a student holiday (studentholiday),
// and/or a public holiday (publicholiday), and provide the ordinal number of the day in the student calendar
// (calendar_date_number; to be ignored if given the value -1).
//
// * Date is set to date.
//
// * SchoolYear is set to calendar.SchoolYear.
//
// * CalendarDateType is set to the enum corresponding to the holiday status of the day:
// "0846" if publicholiday, else "0845" if studentholiday, else "INST".
//
// * StudentAttendance.CountsTowardAttendance is set to No only if studentholiday is set.
//
// * TeacherAttendance.CountsTowardAttendance is set to No only if publicholiday is set.
//
// * AdministratorAttendance.CountsTowardAttendance is set to No only if publicholiday is set.
//
// * All of StudentAttendance.AttendanceValue, TeacherAttendance.AttendanceValue, AdministratorAttendance.AttendanceValue
// are set to either 1.0 or 0.0, depending on whether they count towards that attendance.
//
func Create_CalendarDate(calendar *sifxml.CalendarSummary, school *sifxml.SchoolInfo, date time.Time,
	studentholiday bool, publicholiday bool, calendar_date_number int) *sifxml.CalendarDate {
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

// Create a range of CalendarDate objects linked to CalendarSummary object calendar and SchoolInfo object school.
//
// A calendar date is generated for each weekday between calendar.StartDate and calendar.EndDate. It is deemed
// a public holiday if it is a public holiday in Victoria. It is deemed a student hoiliday if it is a public
// holiday, or else if it is between days 45 and 55, 105 and 115, and 155 and 165 exclusive of the calendar.
//
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
			calendar_date_number > 155 && calendar_date_number < 165 {
			student_holiday = true
		}
		ret = append(ret, Create_CalendarDate(calendar, school, date, student_holiday,
			!local_c.IsWorkday(date), calendar_date_number))
		calendar_date_number++
	}
	return ret
}

// Create a TimeTable object linked to SchoolInfo object school.
//
// * LocalId is a sequence number shared between all objects.
//
// * SchoolYear is the current year.
//
// * Title is "Timetable" plus the timetable RefId.
//
// * DaysPerCycle is set as 10.
//
// * PeriodsPerDay is set as 7.
//
// * TeachingPeriodsPerDay is set as 6.
//
// * TimeTableDayList/TimeTableDay/DayTitle cycles through "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" twice.
//
// * TimeTableDayList/TimeTableDay/TimeTablePeriodList/TimeTablePeriod/PeriodTitle cycles through the list in Periods().
//
// * TimeTableDayList/TimeTableDay/TimeTablePeriodList/TimeTablePeriod/PeriodStart is set through PeriodStart().
//
// * TimeTableDayList/TimeTableDay/TimeTablePeriodList/TimeTablePeriod/PeriodEnd is set through PeriodEnd().
//
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
	periods := Periods()
	for i := 0; i < 10; i++ {
		ret.TimeTableDayList().AddNew()
		ret.TimeTableDayList().Last().SetProperty("DayTitle", days[i%5])
		ret.TimeTableDayList().Last().SetProperty("DayId", strconv.Itoa(i+1))
		for j := 0; j < 7; j++ {
			ret.TimeTableDayList().Last().TimeTablePeriodList().AddNew()
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("PeriodTitle",
				periods[j])
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("PeriodId",
				strconv.Itoa(j+1))
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("StartTime",
				PeriodStart(j+1).Format("15:04:05"))
			ret.TimeTableDayList().Last().TimeTablePeriodList().Last().SetProperty("EndTime",
				PeriodEnd(j+1).Format("15:04:05"))
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

// Create a TimeTableSubject object linking to SchoolInfo object school, SchoolCourseInfo object course,
// with parameters giving the subject matter (subject), the year level (acyear), the end year level if
// the subject is offered across multiple year levels (acyear_end; "" if it does not apply), and the semester number.
//
// Code assumes that subject has the same enum value across all year levels.
// Code further assumes that time table subjects across year labels can be grouped on SubjectShortName as a
// label for the subject matter.
//
// * SchoolYear is the current year.
//
// * SubjectShortName is subject.
//
// * SubjectLongName is a long version of subject (looked up locally).
//
// * SubjectType is a random selection of one of "Core", "Elective", "?".
//
// * SubjectLocalId is formed from concatenating SubjectLongName, acyear (dash acyear_end, if present), and
// a unique random number between 100 and 1000.
//
// * ProposedMinClassSize is set to a random number between 5 and 20.
//
// * ProposedMaxClassSize is set to ProposedMinClassSize.
//
// * Semester is set to semester.
//
// * OtherCodeList has a single entry, with Codeset set to "mycodeset", and Value set to a unique random number
// between 100 and 1000.
//
// * AcademicYear is set to acyear if acyear_end is empty.
//
// * AcademicYearRange.Start and AcademicYearRange.End are set to acyear and acyear_end, if acyear_end is not empty.
//
func Create_TimeTableSubject(school *sifxml.SchoolInfo, course *sifxml.SchoolCourseInfo, subject string,
	acyear string, acyear_end string, semester int) *sifxml.TimeTableSubject {
	code := strconv.Itoa(random_seq_gen("timetablesubjects", 900) + 100)

	ret := sifxml.TimeTableSubject{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolYear", this_year())

	ret.SetProperty("SubjectShortName", subject)
	ret.SetProperty("SubjectLongName", teachingSubjectLongName(subject))
	ret.SetProperty("SubjectType", randomStringFromSlice([]string{"Core", "Elective", "?"}))
	if acyear_end == "" {
		ret.SetProperty("SubjectLocalId", fmt.Sprintf("%s Y%s %s", teachingSubjectLongName(subject), acyear,
			code))
	} else {
		ret.SetProperty("SubjectLocalId", fmt.Sprintf("%s Y%s-%s %s", teachingSubjectLongName(subject),
			acyear, acyear_end, code))
	}
	ret.SetProperty("ProposedMinClassSize", float64(rand.Intn(15)+5))
	ret.SetProperty("ProposedMaxClassSize", ret.ProposedMinClassSize().Float())
	ret.SetProperty("Semester", semester)

	ret.OtherCodeList().AddNew()
	ret.OtherCodeList().Last().SetProperty("Codeset", "mycodeset")
	ret.OtherCodeList().Last().SetProperty("Value",
		strconv.Itoa(random_seq_gen("other_timetablesubjects", 900)+100))

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

// Create TimeTableSubject objects linked to SchoolInfo object school. One object is created for
// each subject in the list given, for each year level that the school offers, and for each of the
// terms in the TermInfo slice terms.
// Presupposes that terms are in consecutive chronological order, and are semesters.
//
// There is no implementation of differentiating timetable subject offerings by year level.
//
func Create_TimeTableSubjects(school *sifxml.SchoolInfo, subjects []string,
	terms []*sifxml.TermInfo) []*sifxml.TimeTableSubject {
	random_seq_gen_reset("timetablesubjects")
	random_seq_gen_reset("other_timetablesubjects")
	if len(subjects) == 0 {
		subjects = All_teachingSubjects()
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

// Create a TermInfo object linked to the SchoolInfo object school, and representing the given semester.
//
// * SchoolYear is set to the current year.
//
// * StartDate is set to the start date of the given semester, as given in Term_start_date().
//
// * EndtDate is set to the start date of the given semester, as given in Term_end_date().
//
// * TermCode is set to "Term" followed by the semester number.
//
// * TermSpan is set to "0828" (Semester).
//
// * MarkingTerm, SchedulingTerm, and AttendanceTerm are all set to "Y".
//
func Create_TermInfo(school *sifxml.SchoolInfo, semester int) *sifxml.TermInfo {
	ret := sifxml.TermInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolInfoRefId", school.RefId().String())
	ret.SetProperty("SchoolYear", this_year())
	ret.SetProperty("StartDate", Term_start_date(this_year(), semester))
	ret.SetProperty("EndDate", Term_end_date(this_year(), semester))
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

// Create TermInfo objects linked to SchoolInfo object school. Two objects are created for the year, each representing a semester.
func Create_TermInfos(school *sifxml.SchoolInfo) []*sifxml.TermInfo {
	ret := sifxml.TermInfoSlice()
	for semester := 1; semester <= 2; semester++ {
		ret = append(ret, Create_TermInfo(school, semester))
	}
	return ret
}

// Create a TimeTableCell object linked to the SchoolInfo object school, the TimeTable object timetable,
// the TimeTableSubject object subject, the TeachingGroup object teachinggroup, the RoomInfo object room
// (as primary room), the slice of RoomInfo objects rooms, the StaffPersonal object staff (as primary staff),
// and the slice of StaffPersonal objects teachers. The created object has DayId day, PeriodId period, and Celltype celltype.
//
// * TeacherList is set through AddToScheduledTeacherList(), with period starts and ends set through PeriodStart() and PeriodEnd().
//
func Create_TimeTableCell(day string, period string, celltype string, school *sifxml.SchoolInfo,
	timetable *sifxml.TimeTable, subject *sifxml.TimeTableSubject, teachinggroup *sifxml.TeachingGroup,
	room *sifxml.RoomInfo, rooms []*sifxml.RoomInfo, staff *sifxml.StaffPersonal,
	teachers []*sifxml.StaffPersonal) *sifxml.TimeTableCell {
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
		AddToScheduledTeacherList(ret.TeacherList(), teachers, PeriodStart(period_int), PeriodEnd(period_int))
	}

	if out, ok := sifxml.TimeTableCellPointer(ret); !ok {
		log.Fatalf("Could not create pointer to TimeTableCell: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create a SessionInfo object for TimeTableCell c, and the given date.
//
// * LocalId is a sequence number shared between all objects.
//
// * SchoolYear is set to the year component of date.
//
// * SessionDate is set to date.
//
// * StartTime is set through PeriodStart().
//
// * EndTime is set through PeriodEnd().
//
// * RollMarked is set to "Y".
//
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
	ret.SetProperty("StartTime", PeriodStart(periodid).Format("15:04:05"))
	ret.SetProperty("FinishTime", PeriodEnd(periodid).Format("15:04:05"))
	ret.SetProperty("RollMarked", "Y")

	if out, ok := sifxml.SessionInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SessionInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create a sequence of SessionInfo objects, one for each TimeTableCell object in cells, and for each
// CalendarDate object in dates which has a StudentAttendance.CountsTowardAttendance value of "Y".
//
// Presupposes that the calendar dates are in the current year. Presupposes that all calendar dates are on weekdays.
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

// Create a ScheduledActivity object, based on the contents of TimeTableCell object c, and TeachingGroup object tg,
// for the given date (ActivityDate).
//
// If c is not nil (so the ScheduledActivity reflects a recurring teaching event):
//
// * The TimeTableCell-based attributes TimeTableCellRefId, DayId, PeriodId, TimeTableRefId, CellType, TeacherList,
// RoomList are copied across.
//
// * StartTime and FinishTime are set based on PeriodId, through PeriodStart(), PeriodEnd().
//
// Else, if tg is not nil (so the ScheduledActivity is a non-recurring event for a teaching group):
//
// * StartTime is set as "09:00:00".
//
// * EndTime is set as "15:00:00".
//
// * Location is set to "Zoo".
//
// * ActivityType is set to "Excursion".
//
// * ActivityName is set to "Zoo Excursion".
//
func Create_ScheduledActivityBasic(date string, c *sifxml.TimeTableCell,
	tg *sifxml.TeachingGroup) *sifxml.ScheduledActivity {
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
		ret.SetProperty("StartTime", PeriodStart(periodid).Format("15:04:05"))
		ret.SetProperty("FinishTime", PeriodEnd(periodid).Format("15:04:05"))
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

// Create ScheduledActivity objects corresponding to the slice of TimeTableCell objects cells,
// the slice of CalendarDate objects dates, and the slice of TeachingGroup objects tg.
// Uses Create_ScheduledActivityBasic().
// Creates one scheduled activity for each cell and each date which has StudentAttendance.CountsTowardAttendance as "Yes".
// Also creates one scheduled activity as an excursion for each teaching group, on a random day
// in the calendar with StudentAttendance.CountsTowardAttendance as "Yes", before the last 10 days of dates.
//
// Presupposes that the calendar dates are in the current year. Presupposes that all calendar dates are on weekdays.
//
//
func Create_ScheduledActivities(cells []*sifxml.TimeTableCell, dates []*sifxml.CalendarDate,
	tg []*sifxml.TeachingGroup) []*sifxml.ScheduledActivity {
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
		for c := dates[rand.Intn(len(dates)-10)]; c.StudentAttendance().CountsTowardAttendance().String() == "No"; c = dates[rand.Intn(len(dates))] {
			date := c.Date().String()
			ret = append(ret, Create_ScheduledActivityBasic(date, nil, g))
		}
	}
	return ret
}

// Creates a SchoolCourseInfo object, linking to TimeTableSubject object subject and the slice of TermInfo object terms.
// Presupposes that terms are in consecutive chronological order, and are semesters.
//
// Updates subject: sets subject.SchoolCourseInfoRefId to be the RefId of the generated object.
//
// * CourseCode is set to a unique random number between 1 and 10000.
//
// * SchoolYear is copied across from subject.
//
// * CourseTitle is copied across from subject.SubjectLocalId.
//
// * TermInfoRefId is identified as the RefId of the nth element of terms, where n is subject.Semester.
//
func Create_SchoolCourseInfo(subject *sifxml.TimeTableSubject, terms []*sifxml.TermInfo) *sifxml.SchoolCourseInfo {
	ret := sifxml.SchoolCourseInfo{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("SchoolInfoRefId", subject.SchoolInfoRefId().String())
	ret.SetProperty("SchoolLocalId", subject.SchoolLocalId().String())
	ret.SetProperty("SchoolYear", subject.SchoolYear().String())
	ret.SetProperty("CourseCode", fmt.Sprintf("%05d", random_seq_gen("course_code", 99999)+1))
	ret.SetProperty("CourseTitle", subject.SubjectLocalId().String())
	ret.SetProperty("TermInfoRefId", terms[subject.Semester().Int()-1].RefId().String())

	subject.SetProperty("SchoolCourseInfoRefId", ret.RefId().String())

	if out, ok := sifxml.SchoolCourseInfoPointer(ret); !ok {
		log.Fatalf("Could not create pointer to SessionInfo: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Creates a SchoolCourseInfo object for each subject in TimeTableSubject, linking to the slice of TermInfo object terms.
// Presupposes that terms are in consecutive chronological order, and are semesters.
//
func Create_SchoolCourseInfos(subjects []*sifxml.TimeTableSubject, terms []*sifxml.TermInfo) []*sifxml.SchoolCourseInfo {
	ret := sifxml.SchoolCourseInfoSlice()
	for _, s := range subjects {
		ret = append(ret, Create_SchoolCourseInfo(s, terms))
	}
	return ret
}

func classesPerWeek(yrlvl string, subject string) int {
	return 2
}

func primaryOrSecondary(studentsperyr map[string][]*sifxml.StudentPersonal) (primary bool, secondary bool,
	snr_secondary bool) {
	primary = false
	secondary = false
	snr_secondary = false

	for _, y := range []string{"1", "2", "3", "4", "5", "6"} {
		if len(studentsperyr[y]) > 0 {
			primary = true
		}
	}
	for _, y := range []string{"7", "8", "9"} {
		if len(studentsperyr[y]) > 0 {
			secondary = true
		}
	}
	for _, y := range []string{"10", "11", "12"} {
		if len(studentsperyr[y]) > 0 {
			snr_secondary = true
		}
	}
	return primary, secondary, snr_secondary
}

/* split up staff */
func primaryOrSecondaryStaff(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal, primary bool, secondary bool,
	snr_secondary bool, subjects []string) ([]*sifxml.StaffPersonal, []*sifxml.StaffPersonal,
	map[string][]*sifxml.StaffPersonal) {
	staffcount := len(staff)
	primarystaff := sifxml.StaffPersonalSlice()
	secondarystaff := sifxml.StaffPersonalSlice()
	if primary && (secondary || snr_secondary) {
		for i := 0; i < staffcount/2; i++ {
			primarystaff = append(primarystaff, staff[i])
		}
		for i := staffcount / 2; i < staffcount; i++ {
			secondarystaff = append(secondarystaff, staff[i])
		}
	} else if primary {
		primarystaff = staff
	} else {
		secondarystaff = staff
	}

	subjectstaff := make(map[string][]*sifxml.StaffPersonal)
	staffpersubject_count := (len(secondarystaff) / len(subjects))
	if secondary || snr_secondary {
		if staffpersubject_count < 1 {
			staffpersubject_count = staffcount / 2
		}
		if staffpersubject_count < 1 {
			staffpersubject_count = 1
		}
		for _, s := range subjects {
			subjectstaff[s] = sifxml.StaffPersonalSlice()
			for j := 0; j < staffpersubject_count; j++ {
				idx := random_seq_gen(fmt.Sprintf("%s-%s", school.RefId().String(), s),
					len(secondarystaff))
				subjectstaff[s] = append(subjectstaff[s], secondarystaff[idx])
			}
		}
	}
	return primarystaff, secondarystaff, subjectstaff
}

/* split up time table subjects by year level and subject matter */
func timeTableSubjects2yr2subj(tts []*sifxml.TimeTableSubject) (map[string]map[string]*sifxml.TimeTableSubject,
	[]string) {
	tts2subj := make(map[string]map[string]*sifxml.TimeTableSubject)
	subjects_map := make(map[string]bool)
	for _, x := range tts {
		if !x.SubjectShortName_IsNil() {
			s := x.SubjectShortName().String()
			subjects_map[s] = true
			if !x.AcademicYear_IsNil() {
				y := x.AcademicYear().Code().String()
				if _, ok := tts2subj[y]; !ok {
					tts2subj[y] = make(map[string]*sifxml.TimeTableSubject)
				}
				if _, ok := tts2subj[y][s]; !ok {
					tts2subj[y][s] = x
				}
			} else if !x.AcademicYearRange_IsNil() {
				/* assume numeric year levels */
				start, _ := strconv.Atoi(x.AcademicYearRange().Start().Code().String())
				end, _ := strconv.Atoi(x.AcademicYearRange().End().Code().String())
				for y_int := start; y_int <= end; y_int++ {
					y := fmt.Sprint(y_int)
					if _, ok := tts2subj[y]; !ok {
						tts2subj[y] = make(map[string]*sifxml.TimeTableSubject)
					}
					if _, ok := tts2subj[y][s]; !ok {
						tts2subj[y][s] = x
					}
				}
			}
		}
	}
	subjects := make([]string, 0)
	for k, _ := range subjects_map {
		subjects = append(subjects, k)
	}
	return tts2subj, subjects
}

func classsize(yrlvl string) int {
	return 20
}

/* 1 teacher per group out of primary staff, students only in 1 group, no subject distinctions */
func makePrimaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal,
	primarystaff []*sifxml.StaffPersonal) []*sifxml.TeachingGroup {
	ret := sifxml.TeachingGroupSlice()
	for _, y := range []string{"1", "2", "3", "4", "5", "6"} {
		for studentidx := 0; studentidx*classsize(y) < len(studentsperyr[y]); studentidx++ {
			class_students := sifxml.StudentPersonalSlice()
			for i := studentidx * classsize(y); i < len(studentsperyr[y]) && i < studentidx*classsize(y)+classsize(y); i++ {
				class_students = append(class_students, studentsperyr[y][i])
			}
			class_teachers := sifxml.StaffPersonalSlice()
			class_teachers = append(class_teachers, primarystaff[rand.Intn(len(primarystaff))])
			tg := Create_TeachingGroup(school, class_students, class_teachers, nil)
			tg.SetProperty("ShortName", y+string(65+studentidx))
			tg.SetProperty("LongName", y+string(65+studentidx))
			tg.Unset("KeyLearningArea")
			tg.SetProperty("CurriculumLevel", y)
			ret = append(ret, tg)
		}
	}
	return ret
}

/* 1 teacher per group out of secondary staff, students only in 1 group, no subject distinctions */
func makeSecondaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal,
	secondarystaff []*sifxml.StaffPersonal) []*sifxml.TeachingGroup {
	ret := sifxml.TeachingGroupSlice()
	for _, y := range []string{"7", "8", "9"} {
		for studentidx := 0; studentidx*classsize(y) < len(studentsperyr[y]); studentidx++ {
			class_students := sifxml.StudentPersonalSlice()
			for i := studentidx * classsize(y); i < len(studentsperyr[y]) && i < studentidx*classsize(y)+classsize(y); i++ {
				class_students = append(class_students, studentsperyr[y][i])
			}
			class_teachers := sifxml.StaffPersonalSlice()
			class_teachers = append(class_teachers, secondarystaff[rand.Intn(len(secondarystaff))])
			tg := Create_TeachingGroup(school, class_students, class_teachers, nil)
			tg.SetProperty("ShortName", y+string(65+studentidx))
			tg.SetProperty("LongName", y+string(65+studentidx))
			tg.Unset("KeyLearningArea")
			tg.SetProperty("CurriculumLevel", y)
			ret = append(ret, tg)
		}
	}
	return ret
}

/* 1 teacher per group, students in random 4 groups, groups specific to subjects */
func makeSnrSecondaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal,
	subjectstaff map[string][]*sifxml.StaffPersonal, subjects []string,
	tts2subj map[string]map[string]*sifxml.TimeTableSubject) []*sifxml.TeachingGroup {
	ret := sifxml.TeachingGroupSlice()
	for _, y := range []string{"10", "11", "12"} {
		student2subjects := make(map[int][]string)
		for s := range studentsperyr[y] {
			random_seq_gen_reset(fmt.Sprintf("student-%d-subjects", s))
			student2subjects[s] = make([]string, 0)
			for i := 0; i < 4; i++ {
				idx := random_seq_gen(fmt.Sprintf("student-%d-subjects", s), len(subjects))
				student2subjects[s] = append(student2subjects[s], subjects[idx])
			}
		}
		subject2students := make(map[string][]*sifxml.StudentPersonal)
		for _, s := range subjects {
			subject2students[s] = sifxml.StudentPersonalSlice()
		}
		for k, v := range student2subjects {
			for _, s := range v {
				subject2students[s] = append(subject2students[s], studentsperyr[y][k])
			}
		}

		for _, s := range subjects {
			for studentidx := 0; studentidx*classsize(y) < len(subject2students[s]); studentidx++ {
				class_students := sifxml.StudentPersonalSlice()
				for i := studentidx * classsize(y); i < len(subject2students[s]) && i < studentidx*classsize(y)+classsize(y); i++ {
					class_students = append(class_students, subject2students[s][i])
				}
				class_teachers := sifxml.StaffPersonalSlice()
				class_teachers = append(class_teachers, subjectstaff[s][rand.Intn(len(subjectstaff[s]))])
				tg := Create_TeachingGroup(school, class_students, class_teachers, tts2subj[y][s])
				tg.SetProperty("ShortName", s+" "+y+string(65+studentidx))
				tg.SetProperty("LongName", teachingSubjectLongName(s)+" "+y+string(65+studentidx))
				tg.SetProperty("KeyLearningArea", TeachingGroupKLA(s))
				tg.SetProperty("CurriculumLevel", y)
				ret = append(ret, tg)
			}
		}
	}
	return ret
}

/* presupposed for a single school. Using CurriculumLevel to track year level of group */

// Create TeachingGroup objects linking to SchoolInfo object school, drawing on preexisting objects in staff, students, and subjects.
//
// * Create teaching groups for all 12 year levels.
//
// * Split the student population into the year levels they are assigned to.
//
// * Split staff into primary school teachers (half) and secondary school teachers (half).
//
// * There are len(secondarystaff)/len(subjects) secondary staff assigned to each subject, meaning that on average each
// teacher is assigned to one subject in the best case. Assignment of secondary staff to subjects is random.
//
// * Teaching groups have a membership of at most 20 students. Students are assigned in slices of 20 to each group; the remainder
// is assigned to another group.
//
// * CurriculumLevel is set to the year level of the group.
//
// * Primary school students (years 1 to 6) are assigned to a single teaching group. No KeyLearningArea or TimeTableSubject
// is assigned to the group. A random single primary teacher is assigned to the group. The ShortName and LongName of the group
// are both of format "6A".
//
// * Junior secondary school students (years 7 to 9) are assigned to a single teaching group. No KeyLearningArea or TimeTableSubject
// is assigned to the group. A random single secondary teacher is assigned to the group. The ShortName and LongName of the group
// are both of format "7A".
//
// * Senior secondary school students (years 10 to 12) are each assigned to four random subjects. For each subject in each year level,
// teaching groups are created to accommodate the assgined students. A random single secondary teacher assigned to that subject
// is assigned to each group. The KeyLearningArea corresponding to the subject is assigned to the group (TeachingGroupKLA()).
// The ShortName and LongName are of format "MAT 12C" and "Mathematics 12C" respectively.
func Create_TeachingGroups(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal,
	students []*sifxml.StudentPersonal, subjects []*sifxml.TimeTableSubject) []*sifxml.TeachingGroup {
	yrs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	ret := sifxml.TeachingGroupSlice()

	/* split up students */
	studentsperyr := make(map[string][]*sifxml.StudentPersonal)
	for _, y := range yrs {
		studentsperyr[y] = sifxml.StudentPersonalSlice()
	}
	for _, s := range students {
		studentsperyr[s.MostRecent().YearLevel().Code().String()] =
			append(studentsperyr[s.MostRecent().YearLevel().Code().String()], s)
	}
	tts2subj, subjectnames := timeTableSubjects2yr2subj(subjects)
	primary, secondary, snr_secondary := primaryOrSecondary(studentsperyr)
	primarystaff, secondarystaff, subjectstaff := primaryOrSecondaryStaff(school, staff, primary, secondary,
		snr_secondary, subjectnames)

	if primary {
		ret = append(ret, makePrimaryTeachingGroups(school, studentsperyr, primarystaff)...)
	}
	if secondary {
		ret = append(ret, makeSecondaryTeachingGroups(school, studentsperyr, secondarystaff)...)
	}
	if snr_secondary {
		ret = append(ret, makeSnrSecondaryTeachingGroups(school, studentsperyr, subjectstaff, subjectnames,
			tts2subj)...)
	}

	return ret
}

// Create TimeTableCell objects linking to SchoolInfo object school, TimeTable object timetable,
// slice of StaffPersonal objects staff, slice of RoomInfo objects rooms, and slice of TimeTableSubject objects tts.
//
// Teaching groups without a timetable subject (i.e. Primary, Junior Secondary) are assigned two cells a week
// for each subject in tts that is available for that year level, in the same randomly selected room.
//
// Teaching groups with a timetable subject (i.e. Senior Secondary) are assigned two cells a week, each teaching group
// in a separate randomly selected room.
//
// The two groups for each subject and teaching group are on randomly selected weekdays and periods, but are not at the same time.
// Each timetable cell is assigned the first teacher listed for the group.
//
// No attempt whatsoever to avoid room clashes. The year level of each teaching group is determined from its
// CurriculumLevel. Currently all subjects in the program are available at all year levels.
//
func Create_TimeTableCells(school *sifxml.SchoolInfo, timetable *sifxml.TimeTable, tg []*sifxml.TeachingGroup,
	staff []*sifxml.StaffPersonal, rooms []*sifxml.RoomInfo, tts []*sifxml.TimeTableSubject) []*sifxml.TimeTableCell {
	staff_map := make(map[string]*sifxml.StaffPersonal)
	for _, s := range staff {
		staff_map[s.RefId().String()] = s
	}
	tts_map := make(map[string]*sifxml.TimeTableSubject)
	for _, s := range tts {
		tts_map[s.RefId().String()] = s
	}

	tts2subj, _ := timeTableSubjects2yr2subj(tts)
	dayids := []string{"1", "2", "3", "4", "5"}
	periodids := []string{"1", "2", "3", "5", "6", "7"}
	seen := make(map[string]map[string]bool)

	start_times := make(map[string]map[string]string)
	for i := 0; i < timetable.TimeTableDayList().Len(); i++ {
		e, _ := timetable.TimeTableDayList().Index(i)
		dayid := e.DayId().String()
		start_times[dayid] = make(map[string]string)
		for j := 0; j < e.TimeTablePeriodList().Len(); j++ {
			e1, _ := e.TimeTablePeriodList().Index(j)
			periodid := e1.PeriodId().String()
			start_times[dayid][periodid] = e1.StartTime().String()
		}
	}

	ret := sifxml.TimeTableCellSlice()
	for _, g := range tg {
		yrlvl := g.CurriculumLevel().String()
		if g.TimeTableSubjectRefId_IsNil() {
			/* assign 2 cells a week of each subject available for that year level, in the same room */
			room := rooms[rand.Intn(len(rooms))]
			for _, subj := range tts2subj[yrlvl] {
				for i := 0; i < classesPerWeek(yrlvl, subj.SubjectShortName().String()); i++ {
					dayid := randomStringFromSlice(dayids)
					periodid := randomStringFromSlice(periodids)
					for ; seen[dayid][periodid]; periodid = randomStringFromSlice(periodids) {
					}
					t0, _ := g.TeacherList().Index(0)
					cell := Create_TimeTableCell(dayid, periodid, "Teaching", school, timetable,
						subj, g, room, sifxml.RoomInfoSlice(),
						staff_map[t0.StaffPersonalRefId().String()], sifxml.StaffPersonalSlice())
					ret = append(ret, cell)
					g = copyTeachingGroupPeriodFromCell(g, cell, start_times)
				}
			}
		} else {
			/* assign 2 cells a week for each teaching group */
			for i := 0; i < classesPerWeek(yrlvl, tts_map[g.TimeTableSubjectRefId().String()].SubjectShortName().String()); i++ {
				dayid := randomStringFromSlice(dayids)
				periodid := randomStringFromSlice(periodids)
				for ; seen[dayid][periodid]; periodid = randomStringFromSlice(periodids) {
				}
				room := rooms[rand.Intn(len(rooms))]
				t0, _ := g.TeacherList().Index(0)
				cell := Create_TimeTableCell(dayid, periodid, "Teaching", school, timetable,
					tts_map[g.TimeTableSubjectRefId().String()], g, room, sifxml.RoomInfoSlice(),
					staff_map[t0.StaffPersonalRefId().String()], sifxml.StaffPersonalSlice())
				ret = append(ret, cell)
				g = copyTeachingGroupPeriodFromCell(g, cell, start_times)
			}
		}
	}

	return ret
}

// Creates a CollectionRound object for a given AG Collection.
// Presupposes that there are two rounds for each collection per year.
//
// * AGCollection is set to the standard abbreviation for the collection.
//
// * CollectionYear is set to the current year.
//
// * AGRoundList/AGRound/RoundCode is set to the collection abbreviation plus the year plus the round number.
//
// * AGRoundList/AGRound/RoundName is set to the expansion of the collection abbreviation plus the year plus the round number.
//
// * In Round 1, AGRoundList/AGRound/StartDate is set to March 1.
// In Round 2, it is set to September 1.
//
// * In Round 1, AGRoundList/AGRound/DueDate is set to March 30.
// In Round 2, it is set to a September 30.
//
// * In Round 1, AGRoundList/AGRound/EndDate is set to March 30.
// In Round 2, it is set to September 30.
//
func Create_CollectionRound(collection string) *sifxml.CollectionRound {

	ret := sifxml.CollectionRound{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("AGCollection", collection)
	ret.SetProperty("CollectionYear", this_year())
	for i := 0; i < 2; i++ {
		ret.AGRoundList().AddNew()
		ret.AGRoundList().Last().SetProperty("RoundCode", CollectionRoundCode(collection, this_year(), i+1))
		ret.AGRoundList().Last().SetProperty("RoundName", fmt.Sprintf("%s %s:%02d", CollectionCode2Name(collection), this_year(), i+1))
		ret.AGRoundList().Last().SetProperty("StartDate", fmt.Sprintf("%s-%02d-01", this_year(), i*6+3))
		ret.AGRoundList().Last().SetProperty("DueDate", fmt.Sprintf("%s-%02d-30", this_year(), i*6+3))
		ret.AGRoundList().Last().SetProperty("EndDate", fmt.Sprintf("%s-%02d-30", this_year(), i*6+3))
	}

	if out, ok := sifxml.CollectionRoundPointer(ret); !ok {
		log.Fatalf("Could not create pointer to CollectionRound: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create CollectionRound objects for the currently supported AG Collections.
//
func Create_CollectionRounds() []*sifxml.CollectionRound {
	ret := sifxml.CollectionRoundSlice()
	col := All_AGCollections()
	for i := 0; i < len(col); i++ {
		ret = append(ret, Create_CollectionRound(col[i]))
	}
	return ret
}

// Create a CollectionStatus object, given the standard abbreviation for the collection, and the round number.
// Presupposes that there are two rounds for each collection per year. The CollectionStatus is not linked
// to any Collection objects submitted to HITS, as it is statically populated.
//
// * ReportingAuthority is set to "Middleton Primary School Reporting Authority".
//
// * ReportingAuthoritySystem is set to "Reporting Authority System B"
//
// * ReportingAuthorityCommonwealthId is set to "1234".
//
// * SubmittedBy is set to "John Smith".
//
// * SubmissionTimestamp is set to the 30th of either March or September, depending on the round number.
//
// * AGCollection is set to the collection value.
//
// * CollectionYear is set to the current year.
//
// * RoundCode is set to the collection abbreviation plus the current year plus the round number.
//
// * ReportingObjectResponseList/ReportingObjectResponse/HTTPStatusCode is set one of 201 (75%), 422 (15%) or 500 (10%).
//
// * ReportingObjectResponseList/ReportingObjectResponse/ErrorText is set to the HTTP Status text corresponding to HTTPStatusCode.
//
// * ReportingObjectResponseList/ReportingObjectResponse/CommonwealthId is set to "101".
//
// * ReportingObjectResponseList/ReportingObjectResponse/EntityName is set to "Middleton Primary School".
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGSubmissionStatusCode is set to a random value of the valid status codes: "Cancelled", "Declaration Pending", "Declared", "Exempt", "Finalised", "In Error", "In Progress", "In Review", "Not Started", "Processing", "Reopened".
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGRuleList is given three rule entries.
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGRuleList/AGRule/AGRuleCode is set to "WR-nnn", for nnn is 001 to 003.
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGRuleList/AGRule/AGRuleComment is set to "Cannot be ignored because Components do not add up to Total - Fix"
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGRuleList/AGRule/AGRuleResponse is set to "Rejected".
//
// * ReportingObjectResponseList/ReportingObjectResponse/AGRuleList/AGRule/AGRuleStatus is set to "Fail".
//
func Create_CollectionStatus(collection string, round int, school *sifxml.SchoolInfo) *sifxml.CollectionStatus {

	ret := sifxml.CollectionStatus{}
	ret.SetProperty("RefId", create_GUID())
	ret.SetProperty("AGCollection", collection)
	ret.SetProperty("CollectionYear", this_year())
	ret.SetProperty("ReportingAuthority", "Systemic School Reporting Authority")
	ret.SetProperty("ReportingAuthoritySystem", "Systemic Reporting Authority System B")
	ret.SetProperty("ReportingAuthorityCommonwealthId", "1234")
	ret.SetProperty("SubmittedBy", "John Smith")
	ret.SetProperty("SubmissionTimestamp", fmt.Sprintf("%s-%02d-30T09:00:00", this_year(), (round-1)*6+3))
	ret.SetProperty("RoundCode", CollectionRoundCode(collection, this_year(), round))
	ret.AGReportingObjectResponseList().AddNew()
	/* We do not point to submissions in our test data, the submissions are made by the user */
	ret.AGReportingObjectResponseList().Last().SetProperty("CommonwealthId", school.CommonwealthId().String())
	ret.AGReportingObjectResponseList().Last().SetProperty("EntityName", school.SchoolName().String())

	status := threshold_rand_strings([]float64{0.5, 0.45, 0.4, 0.35, 0.3, 0.25, 0.2, 0.15, 0.1, 0.05, 0},
		[]string{"In Error", "Cancelled", "Declaration Pending", "Declared", "Exempt", "Finalised", "In Progress", "In Review", "Not Started", "Processing", "Reopened"})
	statuscode := "201"
	if status == "In Error" {
		statuscode = threshold_rand_strings([]float64{0.5, 0}, []string{"422", "500"})
	}
	ret.AGReportingObjectResponseList().Last().SetProperty("HTTPStatusCode", statuscode)
	ret.AGReportingObjectResponseList().Last().SetProperty("ErrorText", HTTPStatus2Text(statuscode))
	ret.AGReportingObjectResponseList().Last().SetProperty("AGSubmissionStatusCode", status)

	if status == "In Error" {
		for i := 1; i <= 3; i++ {
			ret.AGReportingObjectResponseList().Last().AGRuleList().AddNew()
			ret.AGReportingObjectResponseList().Last().AGRuleList().Last().SetProperty("AGRuleCode", fmt.Sprintf("WR-%03d", i))
			ret.AGReportingObjectResponseList().Last().AGRuleList().Last().SetProperty("AGRuleComment", "Cannot be ignored because Components do not add up to Total - Fix")
			ret.AGReportingObjectResponseList().Last().AGRuleList().Last().SetProperty("AGRuleResponse", "Rejected")
			ret.AGReportingObjectResponseList().Last().AGRuleList().Last().SetProperty("AGRuleStatus", "Fail")
		}
	}

	if out, ok := sifxml.CollectionStatusPointer(ret); !ok {
		log.Fatalf("Could not create pointer to CollectionStatus: %+v", ret)
		return nil
	} else {
		return out
	}
}

// Create CollectionStatus objects for the currently supported AG Collections, and for each round
// created in CollectionRounds.
// Presupposes that there are two rounds for each collection per year.
//
func Create_CollectionStatuses(school *sifxml.SchoolInfo) []*sifxml.CollectionStatus {
	ret := sifxml.CollectionStatusSlice()
	col := All_AGCollections()
	for i := 0; i < len(col); i++ {
		for round := 1; round <= 2; round++ {
			ret = append(ret, Create_CollectionStatus(col[i%len(col)], round, school))
		}
	}
	return ret

}
