package populate

import (
	"../sifxml"
	"fmt"
	//"log"
	"math/rand"
)

type TG_usecase struct {
	Schools        []*sifxml.SchoolInfo
	Students       []*sifxml.StudentPersonal
	Staff          []*sifxml.StaffPersonal
	Enrolments     []*sifxml.StudentSchoolEnrollment
	Assignments    []*sifxml.StaffAssignment
	TeachingGroups []*sifxml.TeachingGroup
}

/*
The design of use cases is going to be different to the Perl, which retrieved objects from the database.
Use case creators will by default create all the objects they need, but can optionally be fed objects to use.
*/

/* counts of staff and students per school */
func MakeTeachingGroupsBare(studentcount int, staffcount int, schoolcount int) TG_usecase {
	ret := TG_usecase{
		Schools:        sifxml.SchoolInfoSlice(),
		Students:       sifxml.StudentPersonalSlice(),
		Staff:          sifxml.StaffPersonalSlice(),
		Enrolments:     sifxml.StudentSchoolEnrollmentSlice(),
		Assignments:    sifxml.StaffAssignmentSlice(),
		TeachingGroups: sifxml.TeachingGroupSlice(),
	}
	for i := 0; i < schoolcount; i++ {
		school := Create_SchoolInfo("Pri/Sec")
		staff := Create_StaffPersonals(staffcount)
		assignments := Create_StaffAssignments(staff, school)
		students := Create_StudentPersonals(studentcount, Schooltype2Yearlevels("Pri/Sec"))
		enrolments := Create_StudentSchoolEnrollments(students, school)

		in := MakeTeachingGroups(school, staff, assignments, students, enrolments)

		ret.Schools = append(ret.Schools, in.Schools...)
		ret.Students = append(ret.Students, in.Students...)
		ret.Staff = append(ret.Staff, in.Staff...)
		ret.Enrolments = append(ret.Enrolments, in.Enrolments...)
		ret.Assignments = append(ret.Assignments, in.Assignments...)
		ret.TeachingGroups = append(ret.TeachingGroups, in.TeachingGroups...)
	}
	return ret
}

/* presupposed for a single school */
func MakeTeachingGroups(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal, assignments []*sifxml.StaffAssignment, students []*sifxml.StudentPersonal, enrolments []*sifxml.StudentSchoolEnrollment) TG_usecase {
	ret := TG_usecase{
		Schools:        sifxml.SchoolInfoSlice(),
		Students:       students,
		Staff:          staff,
		Enrolments:     enrolments,
		Assignments:    assignments,
		TeachingGroups: sifxml.TeachingGroupSlice(),
	}
	ret.Schools = append(ret.Schools, school)
	yrs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	subjects := []string{"MAT", "ENG", "PHYS", "BIO", "CHEM", "COMP", "VIS", "ECON", "HIST"}

	/* split up staff */
	staffcount := len(staff)
	primarystaff := sifxml.StaffPersonalSlice()
	for i := 0; i < staffcount/2; i++ {
		primarystaff = append(primarystaff, staff[i])
	}
	secondarystaff := sifxml.StaffPersonalSlice()
	for i := staffcount / 2; i < staffcount; i++ {
		secondarystaff = append(secondarystaff, staff[i])
	}
	subjectstaff := make(map[string][]*sifxml.StaffPersonal)
	staffpersubject_count := (len(secondarystaff) / len(subjects))
	if staffpersubject_count < 1 {
		staffpersubject_count = staffcount / 2
	}
	if staffpersubject_count < 1 {
		staffpersubject_count = 1
	}
	for _, s := range subjects {
		subjectstaff[s] = sifxml.StaffPersonalSlice()
		for j := 0; j < staffpersubject_count; j++ {
			idx := random_seq_gen(fmt.Sprintf("%s-%s", school.RefId.String(), s), len(secondarystaff))
			subjectstaff[s] = append(subjectstaff[s], secondarystaff[idx])
		}
	}

	/* split up students */
	studentsperyr := make(map[string][]*sifxml.StudentPersonal)
	for _, y := range yrs {
		studentsperyr[y] = sifxml.StudentPersonalSlice()
	}
	for _, s := range students {
		studentsperyr[s.MostRecent.YearLevel.Code.String()] = append(studentsperyr[s.MostRecent.YearLevel.Code.String()], s)
	}

	/* 1 teacher per group out of primary staff, students only in 1 group, no subject distinctions */
	for _, y := range []string{"1", "2", "3", "4", "5", "6"} {
		for studentidx := 0; studentidx*20 < len(studentsperyr[y]); studentidx++ {
			class_students := sifxml.StudentPersonalSlice()
			for i := studentidx * 20; i < len(studentsperyr[y]) && i < studentidx*20+20; i++ {
				class_students = append(class_students, studentsperyr[y][i])
			}
			class_teachers := sifxml.StaffPersonalSlice()
			class_teachers = append(class_teachers, primarystaff[rand.Intn(len(primarystaff))])
			tg := Create_TeachingGroup(school, class_students, class_teachers)
			tg.Set("ShortName", y+string(65+studentidx))
			tg.Set("LongName", y+string(65+studentidx))
			tg.Unset("KeyLearningArea")
			ret.TeachingGroups = append(ret.TeachingGroups, tg)
		}
	}

	/* 1 teacher per group out of secondary staff, students only in 1 group, no subject distinctions */
	for _, y := range []string{"7", "8", "9"} {
		for studentidx := 0; studentidx*20 < len(studentsperyr[y]); studentidx++ {
			class_students := sifxml.StudentPersonalSlice()
			for i := studentidx * 20; i < len(studentsperyr[y]) && i < studentidx*20+20; i++ {
				class_students = append(class_students, studentsperyr[y][i])
			}
			class_teachers := sifxml.StaffPersonalSlice()
			class_teachers = append(class_teachers, secondarystaff[rand.Intn(len(secondarystaff))])
			tg := Create_TeachingGroup(school, class_students, class_teachers)
			tg.Set("ShortName", y+string(65+studentidx))
			tg.Set("LongName", y+string(65+studentidx))
			tg.Unset("KeyLearningArea")
			ret.TeachingGroups = append(ret.TeachingGroups, tg)
		}
	}

	/* 1 teacher per group, students in random 4 groups, groups specific to subjects */
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
			for studentidx := 0; studentidx*20 < len(subject2students[s]); studentidx++ {
				class_students := sifxml.StudentPersonalSlice()
				for i := studentidx * 20; i < len(subject2students[s]) && i < studentidx*20+20; i++ {
					class_students = append(class_students, subject2students[s][i])
				}
				class_teachers := sifxml.StaffPersonalSlice()
				class_teachers = append(class_teachers, subjectstaff[s][rand.Intn(len(subjectstaff[s]))])
				tg := Create_TeachingGroup(school, class_students, class_teachers)
				tg.Set("ShortName", s+" "+y+string(65+studentidx))
				tg.Set("LongName", teachingGroupLongName(s)+" "+y+string(65+studentidx))
				tg.Set("KeyLearningArea", teachingGroupKLA(s))
				ret.TeachingGroups = append(ret.TeachingGroups, tg)
			}
		}
	}

	return ret
}
