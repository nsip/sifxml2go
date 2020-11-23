package main

import (
	"./populate"
	"./sifxml"
)

func main() {
	var err error
	students := populate.Create_StudentPersonals(10, populate.Schooltype2Yearlevels("Pri/Sec"))
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

	t := populate.Create_TeachingGroup(school, students, staff)
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

	act := populate.Create_ScheduledActivity(school, nil, nil, nil, students, staff, make([]*sifxml.TeachingGroup, 0), rooms)
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

	/*
		ret := populate.MakeTeachingGroupsBare(400, 20, 10)
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
		err = populate.PrintXML(ret.TeachingGroups)
		populate.Errcheck(err)
	*/
}
