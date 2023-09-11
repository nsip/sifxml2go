package populate

import (
	"log"

	"github.com/nsip/sifxml2go/sifxml"
)

// All the SIF objects generated by MakeUsecaseObjects(), bundled up as a single object.
type UseCaseObjects struct {
	Schools             *sifxml.SchoolInfos
	Students            *sifxml.StudentPersonals
	Staff               *sifxml.StaffPersonals
	Enrolments          *sifxml.StudentSchoolEnrollments
	Assignments         *sifxml.StaffAssignments
	TeachingGroups      *sifxml.TeachingGroups
	TimeTables          *sifxml.TimeTables
	TimeTableSubjects   *sifxml.TimeTableSubjects
	TimeTableCells      *sifxml.TimeTableCells
	Rooms               *sifxml.RoomInfos
	Contacts            *sifxml.StudentContactPersonals
	Relationships       *sifxml.StudentContactRelationships
	FinancialAccounts   *sifxml.FinancialAccounts
	Vendors             *sifxml.VendorInfos
	Debtors             *sifxml.Debtors
	ChargedLocations    *sifxml.ChargedLocationInfos
	Terms               *sifxml.TermInfos
	CalendarSummarys    *sifxml.CalendarSummarys
	CalendarDates       *sifxml.CalendarDates
	SessionInfos        *sifxml.SessionInfos
	SchoolCourses       *sifxml.SchoolCourseInfos
	ScheduledActivities *sifxml.ScheduledActivitys
	CollectionRounds    *sifxml.CollectionRounds
	CollectionStatuses  *sifxml.CollectionStatuss
}

// The counts of different SIF objects to be generated by MakeUsecaseObjects(), for each object that it makes sense
// to specify a count for as an indepedent variable. Each count is a count of objects to be generated
// per school (except for Schools itself).
type MakeUsecaseCounts struct {
	Students          int
	Staff             int
	Schools           int
	Rooms             int
	Vendors           int
	ChargedLocations  int
	FinancialAccounts int
}

// The different HITS use cases that MakeUsecaseObjects() is to generate objects for. These use cases are documented
// in https://github.com/nsip/usecases/tree/master/docs, replicated in HITS http://hits.nsip.edu.au/dashboard/usecase.html
type MakeUsecases struct {
	DailyAttendance           bool
	Financial                 bool
	Enrolment                 bool
	Gradebook                 bool
	StudentAttendanceTimeList bool
	TeacherJudgement          bool
	Timetable                 bool
	Wellbeing                 bool
	Provisioning              bool
	AGCollections             bool
}

/*
The design of use cases is going to be different to the Perl, which retrieved objects from the database.
Use case creators will by default create all the objects they need, but can optionally be fed objects to use.
*/

func initUseCaseObjects() UseCaseObjects {
	return UseCaseObjects{
		Schools:             sifxml.NewSchoolInfos(),
		Students:            sifxml.NewStudentPersonals(),
		Staff:               sifxml.NewStaffPersonals(),
		Enrolments:          sifxml.NewStudentSchoolEnrollments(),
		Assignments:         sifxml.NewStaffAssignments(),
		TeachingGroups:      sifxml.NewTeachingGroups(),
		TimeTables:          sifxml.NewTimeTables(),
		TimeTableSubjects:   sifxml.NewTimeTableSubjects(),
		TimeTableCells:      sifxml.NewTimeTableCells(),
		Rooms:               sifxml.NewRoomInfos(),
		Contacts:            sifxml.NewStudentContactPersonals(),
		Relationships:       sifxml.NewStudentContactRelationships(),
		FinancialAccounts:   sifxml.NewFinancialAccounts(),
		Vendors:             sifxml.NewVendorInfos(),
		Debtors:             sifxml.NewDebtors(),
		ChargedLocations:    sifxml.NewChargedLocationInfos(),
		Terms:               sifxml.NewTermInfos(),
		CalendarSummarys:    sifxml.NewCalendarSummarys(),
		CalendarDates:       sifxml.NewCalendarDates(),
		SessionInfos:        sifxml.NewSessionInfos(),
		SchoolCourses:       sifxml.NewSchoolCourseInfos(),
		ScheduledActivities: sifxml.NewScheduledActivitys(),
		CollectionRounds:    sifxml.NewCollectionRounds(),
		CollectionStatuses:  sifxml.NewCollectionStatuss(),
	}
}

func appendUseCaseObjects(ret UseCaseObjects, add UseCaseObjects) UseCaseObjects {
	ret.Schools.Append(add.Schools.ToSlice()...)
	ret.Students.Append(add.Students.ToSlice()...)
	ret.Staff.Append(add.Staff.ToSlice()...)
	ret.Rooms.Append(add.Rooms.ToSlice()...)
	ret.Enrolments.Append(add.Enrolments.ToSlice()...)
	ret.Assignments.Append(add.Assignments.ToSlice()...)
	ret.TimeTables.Append(add.TimeTables.ToSlice()...)
	ret.TeachingGroups.Append(add.TeachingGroups.ToSlice()...)
	ret.TimeTableSubjects.Append(add.TimeTableSubjects.ToSlice()...)
	ret.TimeTableCells.Append(add.TimeTableCells.ToSlice()...)
	ret.Contacts.Append(add.Contacts.ToSlice()...)
	ret.Relationships.Append(add.Relationships.ToSlice()...)
	ret.FinancialAccounts.Append(add.FinancialAccounts.ToSlice()...)
	ret.Vendors.Append(add.Vendors.ToSlice()...)
	ret.Debtors.Append(add.Debtors.ToSlice()...)
	ret.ChargedLocations.Append(add.ChargedLocations.ToSlice()...)
	ret.Terms.Append(add.Terms.ToSlice()...)
	ret.CalendarSummarys.Append(add.CalendarSummarys.ToSlice()...)
	ret.CalendarDates.Append(add.CalendarDates.ToSlice()...)
	ret.SessionInfos.Append(add.SessionInfos.ToSlice()...)
	ret.SchoolCourses.Append(add.SchoolCourses.ToSlice()...)
	ret.ScheduledActivities.Append(add.ScheduledActivities.ToSlice()...)
	ret.CollectionRounds.Append(add.CollectionRounds.ToSlice()...)
	ret.CollectionStatuses.Append(add.CollectionStatuses.ToSlice()...)
	return ret
}

// The default counts of objects that MakeUsecaseObjects() is to generate, if they are not overtly
// specified in the call to MakeUsecaseObjects(). As noted, all counts but Schools are counts per school. The defaults are:
//
// * SchoolInfo: 1
//
// * StaffPersonal: 50
//
// * StudentPersonal: 500
//
// * RoomInfo: 100
//
// * VendorInfo: 20
//
// * ChargedLocationInfo: 10
//
// * FinancialAccount: 100
func Initcounts(counts MakeUsecaseCounts) MakeUsecaseCounts {
	if counts.Schools == 0 {
		counts.Schools = 1
	}
	if counts.Staff == 0 {
		counts.Staff = 50
	}
	if counts.Students == 0 {
		counts.Students = 500
	}
	if counts.Rooms == 0 {
		counts.Rooms = 100
	}
	if counts.Vendors == 0 {
		counts.Vendors = 20
	}
	if counts.ChargedLocations == 0 {
		counts.ChargedLocations = 10
	}
	if counts.FinancialAccounts == 0 {
		counts.FinancialAccounts = 100
	}
	return counts
}

// Given a list of HITS use cases to fulfil, and a list of counts of objects to generate (for those counts which are independent
// variables), generate all the SIF objects required to fulfil those use case requirements, ensuring that the objects
// are interrelated as appropriate. Objects are generated one school at a time, and the objects for each school have no
// connection to the objects generated for other schools.
//
// All objects are generated using Create_Xs where appropriate, to guarantee that they are interrelated.
// The only parametrisations local to this function are:
//
// * All schools generated are of School Type "Pri/Sec", i.e. they are combined schools
//
// * Timetable subjects are drawn from All_teachingSubjects()
//
// If appendmode is true, all generated objects are returned in one big object.
// If appendmode is false, generated objects are printed as they are generated.
func MakeUsecaseObjects(usecases MakeUsecases, counts MakeUsecaseCounts, appendmode bool) UseCaseObjects {
	ret := initUseCaseObjects()
	counts = Initcounts(counts)

	for i := 0; i < counts.Schools; i++ {
		if i%2 == 0 && i > 0 {
			log.Printf("Processing school %d of %d...\n", i, counts.Schools)
		}
		school := Create_SchoolInfo("Pri/Sec")
		add := initUseCaseObjects()
		add.Schools.Append(school)

		if usecases.Provisioning || usecases.Gradebook || usecases.StudentAttendanceTimeList || usecases.Financial || usecases.TeacherJudgement || usecases.Timetable || usecases.Wellbeing || usecases.DailyAttendance || usecases.Enrolment {
			add.Students = Create_StudentPersonals(counts.Students, Schooltype2Yearlevels(school.SchoolType().String()))
			add.Enrolments = Create_StudentSchoolEnrollments(add.Students, school)
		}
		if usecases.Provisioning || usecases.Gradebook || usecases.StudentAttendanceTimeList || usecases.Financial || usecases.TeacherJudgement || usecases.Timetable || usecases.Wellbeing {
			add.Staff = Create_StaffPersonals(counts.Staff)
			add.Assignments = Create_StaffAssignments(add.Staff, school)
		}
		if usecases.Financial || usecases.Wellbeing || usecases.Enrolment {
			add.Contacts, add.Relationships = Create_StudentContactPersonalAndRelationship(add.Students)
		}
		if usecases.Financial {
			add.Vendors = Create_VendorInfos(counts.Vendors)
			add.Debtors = Create_Debtors(add.Students, add.Staff, add.Contacts, add.Vendors)
			add.ChargedLocations = Create_ChargedLocationInfos(counts.ChargedLocations, add.Schools)
			add.FinancialAccounts = Create_FinancialAccounts(counts.FinancialAccounts, add.ChargedLocations)
		}
		if usecases.Provisioning || usecases.Gradebook || usecases.StudentAttendanceTimeList || usecases.TeacherJudgement || usecases.Timetable || usecases.Wellbeing {
			add.Terms = Create_TermInfos(school)
			add.TimeTableSubjects = Create_TimeTableSubjects(school, All_teachingSubjects(), add.Terms)
		}
		if usecases.Provisioning || usecases.Gradebook || usecases.StudentAttendanceTimeList || usecases.TeacherJudgement || usecases.Wellbeing {

			add.TeachingGroups = Create_TeachingGroups(school, add.Staff, add.Students, add.TimeTableSubjects)
		}
		if usecases.StudentAttendanceTimeList || usecases.Timetable || usecases.Wellbeing {
			add.Rooms = Create_RoomInfos(counts.Rooms, school)
		}
		if usecases.StudentAttendanceTimeList || usecases.Wellbeing {
			t := Create_TimeTable(school)
			add.TimeTables.Append(t)
			add.TimeTableCells = Create_TimeTableCells(school, t, add.TeachingGroups, add.Staff, add.Rooms, add.TimeTableSubjects)
		}
		if usecases.StudentAttendanceTimeList || usecases.DailyAttendance {
			s := Create_CalendarSummary(school)
			add.CalendarSummarys.Append(s)
			add.CalendarDates.Append(Create_CalendarDates(s, school).ToSlice()...)
		}
		if usecases.StudentAttendanceTimeList || usecases.Wellbeing {
			add.ScheduledActivities = Create_ScheduledActivities(add.TimeTableCells, add.CalendarDates, add.TeachingGroups)
		}
		if usecases.StudentAttendanceTimeList {
			add.SessionInfos = Create_SessionInfos(add.TimeTableCells, add.CalendarDates)
		}
		if usecases.Timetable {
			add.SchoolCourses = Create_SchoolCourseInfos(add.TimeTableSubjects, add.Terms)
		}
		if usecases.AGCollections {
			add.CollectionStatuses = Create_CollectionStatuses(school)
		}
		if appendmode {
			ret = appendUseCaseObjects(ret, add)
		} else {
			/*printAllJSON(add)*/
			printAll(add)
		}
	}
	if usecases.AGCollections {
		// rounds apply across all schools
		log.Printf("Processing AG Collections")
		add := initUseCaseObjects()
		add.CollectionRounds = Create_CollectionRounds()
		if appendmode {
			ret = appendUseCaseObjects(ret, add)
		} else {
			printAll(add)
		}
	}
	return ret
}

func printAll(ret UseCaseObjects) {
	var err error

	err = PrintXML(ret.Schools)
	Errcheck(err)
	err = PrintXML(ret.Students)
	Errcheck(err)
	err = PrintXML(ret.Enrolments)
	Errcheck(err)
	err = PrintXML(ret.Staff)
	Errcheck(err)
	err = PrintXML(ret.Assignments)
	Errcheck(err)
	err = PrintXML(ret.TimeTables)
	Errcheck(err)
	err = PrintXML(ret.TimeTableSubjects)
	Errcheck(err)
	err = PrintXML(ret.TimeTableCells)
	Errcheck(err)
	err = PrintXML(ret.TeachingGroups)
	Errcheck(err)
	err = PrintXML(ret.Contacts)
	Errcheck(err)
	err = PrintXML(ret.Relationships)
	Errcheck(err)
	err = PrintXML(ret.Vendors)
	Errcheck(err)
	err = PrintXML(ret.Debtors)
	Errcheck(err)
	err = PrintXML(ret.ChargedLocations)
	Errcheck(err)
	err = PrintXML(ret.FinancialAccounts)
	Errcheck(err)
	err = PrintXML(ret.Terms)
	Errcheck(err)
	err = PrintXML(ret.CalendarSummarys)
	Errcheck(err)
	err = PrintXML(ret.CalendarDates)
	Errcheck(err)
	err = PrintXML(ret.SessionInfos)
	Errcheck(err)
	err = PrintXML(ret.SchoolCourses)
	Errcheck(err)
	err = PrintXML(ret.ScheduledActivities)
	Errcheck(err)
	err = PrintXML(ret.CollectionRounds)
	Errcheck(err)
	err = PrintXML(ret.CollectionStatuses)
	Errcheck(err)
}

func printAllJSON(ret UseCaseObjects) {
	var err error
	err = PrintJSON(ret)
	Errcheck(err)
	/*
	   err = PrintJSON(ret.Schools)
	   Errcheck(err)
	   err = PrintJSON(ret.Students)
	   Errcheck(err)
	   err = PrintJSON(ret.Enrolments)
	   Errcheck(err)
	   err = PrintJSON(ret.Staff)
	   Errcheck(err)
	   err = PrintJSON(ret.Assignments)
	   Errcheck(err)
	   err = PrintJSON(ret.TimeTables)
	   Errcheck(err)
	   err = PrintJSON(ret.TimeTableSubjects)
	   Errcheck(err)
	   err = PrintJSON(ret.TimeTableCells)
	   Errcheck(err)
	   err = PrintJSON(ret.TeachingGroups)
	   Errcheck(err)
	   err = PrintJSON(ret.Contacts)
	   Errcheck(err)
	   err = PrintJSON(ret.Relationships)
	   Errcheck(err)
	   err = PrintJSON(ret.Vendors)
	   Errcheck(err)
	   err = PrintJSON(ret.Debtors)
	   Errcheck(err)
	   err = PrintJSON(ret.ChargedLocations)
	   Errcheck(err)
	   err = PrintJSON(ret.FinancialAccounts)
	   Errcheck(err)
	   err = PrintJSON(ret.Terms)
	   Errcheck(err)
	   err = PrintJSON(ret.CalendarSummarys)
	   Errcheck(err)
	   err = PrintJSON(ret.CalendarDates)
	   Errcheck(err)
	   err = PrintJSON(ret.SessionInfos)
	   Errcheck(err)
	   err = PrintJSON(ret.SchoolCourses)
	   Errcheck(err)
	   err = PrintJSON(ret.ScheduledActivities)
	   Errcheck(err)
	   err = PrintJSON(ret.CollectionRounds)
	   Errcheck(err)
	   err = PrintJSON(ret.CollectionStatuses)
	   Errcheck(err)
	*/
}
