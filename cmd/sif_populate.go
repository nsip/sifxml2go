package main

import (
	"flag"
	"log"

	"github.com/nsip/sifxml2go/populate"
	//"./sifxml"
)

var enrolment = flag.Bool("enrolment", false, "Generate enrolment use case input data")
var provisioning = flag.Bool("provisioning", false, "Generate provisioning use case input data")
var dailyattendance = flag.Bool("dailyattendance", false, "Generate daily attendance use case input data")
var financial = flag.Bool("financial", false, "Generate financial use case input data")
var gradebook = flag.Bool("gradebook", false, "Generate gradebook use case input data")
var studentattendancetimelist = flag.Bool("studentattendancetimelist", false, "Generate student attendance time list use case input data")
var teacherjudgement = flag.Bool("teacherjudgement", false, "Generate teacher judgement use case input data")
var timetable = flag.Bool("timetable", false, "Generate timetable use case input data")
var wellbeing = flag.Bool("wellbeing", false, "Generate wellbeing use case input data")
var agcollections = flag.Bool("agcollections", false, "Generate AG collections use case input data")
var studentcount = flag.Int("studentcount", 400, "Number of StudentPersonal objects to generate")
var staffcount = flag.Int("staffcount", 40, "Number of StaffPersonal objects to generate")
var schoolcount = flag.Int("schoolcount", 10, "Number of SchoolInfo objects to generate")
var vendorcount = flag.Int("vendorcount", 10, "Number of VendorInfo objects to generate")

func main() {
	flag.Parse()

	var err error
	/*
		students := populate.Create_StudentPersonals(100000, populate.Schooltype2Yearlevels("Pri/Sec"))
		err = populate.PrintXML(students[3])
		populate.Errcheck(err)

		school := populate.Create_SchoolInfo("Pri/Sec")
		err = populate.PrintXML(school)
		populate.Errcheck(err)
		schools := make([]*sifxml.SchoolInfo, 0)
		schools = append(schools, school)

		sse := populate.Create_StudentSchoolEnrollments(students, school)
		err = populate.PrintXML(sse[3])
		populate.Errcheck(err)
		err = populate.PrintXML(students[3])
		populate.Errcheck(err)

		staff := populate.Create_StaffPersonals(10)
		err = populate.PrintXML(staff[3])
		populate.Errcheck(err)

		sa := populate.Create_StaffAssignment(staff[3], school)
		err = populate.PrintXML(sa)
		populate.Errcheck(err)
		err = populate.PrintXML(staff[3])
		populate.Errcheck(err)

		scp, scr := populate.Create_StudentContactPersonalAndRelationship(students)
		err = populate.PrintXML(scp[3])
		populate.Errcheck(err)
		err = populate.PrintXML(scr[3])
		populate.Errcheck(err)

		rooms := populate.Create_RoomInfos(10, school)
		err = populate.PrintXML(rooms[3])
		populate.Errcheck(err)

		tt := populate.Create_TimeTable(school)
		err = populate.PrintXML(tt)
		populate.Errcheck(err)

		tts := populate.Create_TimeTableSubjects(school)
		err = populate.PrintXML(tts[3])
		populate.Errcheck(err)

		t := populate.Create_TeachingGroup(school, students, staff, tts[3])
		err = populate.PrintXML(t)
		populate.Errcheck(err)

		fa := populate.Create_FinancialAccount(nil, nil)
		err = populate.PrintXML(fa)
		populate.Errcheck(err)

		cl := populate.Create_ChargedLocationInfo(nil, school)
		err = populate.PrintXML(cl)
		populate.Errcheck(err)

		cl2 := populate.Create_ChargedLocationInfo(cl, nil)
		err = populate.PrintXML(cl2)
		populate.Errcheck(err)

		cl3 := populate.Create_ChargedLocationInfos(3, schools)
		err = populate.PrintXML(cl3[2])
		populate.Errcheck(err)

		v := populate.Create_VendorInfos(10)
		err = populate.PrintXML(v[3])
		populate.Errcheck(err)

		act := populate.Create_ScheduledActivity(school, tt, nil, tts[3], students, staff, make([]*sifxml.TeachingGroup, 0), rooms)
		err = populate.PrintXML(act)
		populate.Errcheck(err)

		ga := populate.Create_GradingAssignment(school, t, students)
		err = populate.PrintXML(ga)
		populate.Errcheck(err)

		gs := populate.Create_GradingAssignmentScores(ga, school, t, students, staff[3])
		err = populate.PrintXML(gs[3])
		populate.Errcheck(err)

		debt := populate.Create_Debtors(students, staff, scp, v)
		err = populate.PrintXML(debt[3])
		populate.Errcheck(err)

		cal := populate.Create_CalendarSummary(school)
		err = populate.PrintXML(cal)
		populate.Errcheck(err)

		day := populate.Create_CalendarDates(cal, school)
		err = populate.PrintXML(day[3])
		populate.Errcheck(err)

		cell := populate.Create_TimeTableCell("1", "1", "Teaching", school, tt, tts[3], t, rooms[3], rooms, staff[3], staff)
		err = populate.PrintXML(cell)
		populate.Errcheck(err)
	*/
	log.Printf("%+v", *enrolment)
	ret := populate.MakeUsecaseObjects(populate.MakeUsecases{Enrolment: *enrolment, Provisioning: *provisioning, DailyAttendance: *dailyattendance, Financial: *financial, Gradebook: *gradebook, StudentAttendanceTimeList: *studentattendancetimelist, TeacherJudgement: *teacherjudgement, Timetable: *timetable, Wellbeing: *wellbeing, AGCollections: *agcollections}, populate.MakeUsecaseCounts{Students: *studentcount, Staff: *staffcount, Schools: *schoolcount, Vendors: *vendorcount})
	err = populate.PrintXML(ret.Schools)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Students)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Enrolments)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Staff)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Assignments)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.TimeTables)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.TimeTableSubjects)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.TimeTableCells)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.TeachingGroups)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Contacts)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Relationships)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Vendors)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Debtors)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.ChargedLocations)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.FinancialAccounts)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.Terms)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.CalendarSummarys)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.CalendarDates)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.SessionInfos)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.SchoolCourses)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.ScheduledActivities)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.CollectionRounds)
	populate.Errcheck(err)
	err = populate.PrintXML(ret.CollectionStatuses)
	populate.Errcheck(err)
}
