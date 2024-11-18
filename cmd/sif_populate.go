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
var stdn = flag.Bool("stdn", false, "Generate Student Data Transfer Note use case input data")
var studentcount = flag.Int("studentcount", 400, "Number of StudentPersonal objects to generate")
var staffcount = flag.Int("staffcount", 40, "Number of StaffPersonal objects to generate")
var schoolcount = flag.Int("schoolcount", 10, "Number of SchoolInfo objects to generate")
var vendorcount = flag.Int("vendorcount", 10, "Number of VendorInfo objects to generate")

func main() {
	flag.Parse()

	usecases := populate.MakeUsecases{Enrolment: *enrolment, Provisioning: *provisioning, DailyAttendance: *dailyattendance, Financial: *financial, Gradebook: *gradebook, StudentAttendanceTimeList: *studentattendancetimelist, TeacherJudgement: *teacherjudgement, Timetable: *timetable, Wellbeing: *wellbeing, AGCollections: *agcollections, StudentDataTransferNote: *stdn}
	counts := populate.MakeUsecaseCounts{Students: *studentcount, Staff: *staffcount, Schools: *schoolcount, Vendors: *vendorcount}
	log.Printf("%+v", usecases)
	log.Printf("%+v", counts)
	populate.MakeUsecaseObjects(usecases, counts, true)
	/*
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
	*/
}
