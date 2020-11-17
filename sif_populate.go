package main

import (
	"./populate"
)

func main() {
	var err error
	/*
		students := populate.Create_StudentPersonals(10)
		err = populate.PrintXML(students[3])
		populate.Errcheck(err)

		school := populate.Create_SchoolInfo()
		err = populate.PrintXML(school)
		populate.Errcheck(err)

		sse := populate.Create_StudentSchoolEnrollments(students, school)
		err = populate.PrintXML(sse[3])
		populate.Errcheck(err)

		staff := populate.Create_StaffPersonals(10)
		err = populate.PrintXML(staff[3])
		populate.Errcheck(err)

		sa := populate.Create_StaffAssignment(staff[3], school)
		err = populate.PrintXML(sa)
		populate.Errcheck(err)

		scp, scr := populate.Create_StudentContactPersonalAndRelationship(students)
		err = populate.PrintXML(scp[3])
		populate.Errcheck(err)
		err = populate.PrintXML(scr[3])
		populate.Errcheck(err)

		room := populate.Create_RoomInfo(school)
		err = populate.PrintXML(room)
		populate.Errcheck(err)

		t := populate.Create_TeachingGroup(school, students, staff)
		err = populate.PrintXML(t)
		populate.Errcheck(err)
	*/
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
}
