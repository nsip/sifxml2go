package populate

import (
	"../sifxml"
	"fmt"
	"math/rand"
	"strconv"
)

type TG_usecase struct {
	Schools           []*sifxml.SchoolInfo
	Students          []*sifxml.StudentPersonal
	Staff             []*sifxml.StaffPersonal
	Enrolments        []*sifxml.StudentSchoolEnrollment
	Assignments       []*sifxml.StaffAssignment
	TeachingGroups    []*sifxml.TeachingGroup
	TimeTables        []*sifxml.TimeTable
	TimeTableSubjects []*sifxml.TimeTableSubject
	TimeTableCells    []*sifxml.TimeTableCell
	Rooms             []*sifxml.RoomInfo
}

/*
The design of use cases is going to be different to the Perl, which retrieved objects from the database.
Use case creators will by default create all the objects they need, but can optionally be fed objects to use.
*/

/* counts of staff and students per school */
func MakeTeachingGroupsBare(studentcount int, staffcount int, schoolcount int) TG_usecase {
	ret := TG_usecase{
		Schools:           sifxml.SchoolInfoSlice(),
		Students:          sifxml.StudentPersonalSlice(),
		Staff:             sifxml.StaffPersonalSlice(),
		Enrolments:        sifxml.StudentSchoolEnrollmentSlice(),
		Assignments:       sifxml.StaffAssignmentSlice(),
		TeachingGroups:    sifxml.TeachingGroupSlice(),
		TimeTables:        sifxml.TimeTableSlice(),
		TimeTableSubjects: sifxml.TimeTableSubjectSlice(),
		TimeTableCells:    sifxml.TimeTableCellSlice(),
		Rooms:             sifxml.RoomInfoSlice(),
	}
	for i := 0; i < schoolcount; i++ {
		school := Create_SchoolInfo("Pri/Sec")
		staff := Create_StaffPersonals(staffcount)
		assignments := Create_StaffAssignments(staff, school)
		students := Create_StudentPersonals(studentcount, Schooltype2Yearlevels(school.SchoolType().String()))
		enrolments := Create_StudentSchoolEnrollments(students, school)
		timetable := Create_TimeTable(school)
		subjects := all_teachingSubjects()
		tts := Create_TimeTableSubjects(school, subjects)
		rooms := Create_RoomInfos(100, school)

		tg := MakeTeachingGroups(school, staff, students, tts)
		cells := MakeTimeTableCells(school, timetable, tg, staff, rooms, tts)

		ret.Schools = append(ret.Schools, school)
		ret.Students = append(ret.Students, students...)
		ret.Staff = append(ret.Staff, staff...)
		ret.Rooms = append(ret.Rooms, rooms...)
		ret.Enrolments = append(ret.Enrolments, enrolments...)
		ret.Assignments = append(ret.Assignments, assignments...)
		ret.TimeTables = append(ret.TimeTables, timetable)
		ret.TeachingGroups = append(ret.TeachingGroups, tg...)
		ret.TimeTableSubjects = append(ret.TimeTableSubjects, tts...)
		ret.TimeTableCells = append(ret.TimeTableCells, cells...)
	}
	return ret
}

func primaryOrSecondary(studentsperyr map[string][]*sifxml.StudentPersonal) (primary bool, secondary bool, snr_secondary bool) {
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
func primaryOrSecondaryStaff(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal, primary bool, secondary bool, snr_secondary bool, subjects []string) ([]*sifxml.StaffPersonal, []*sifxml.StaffPersonal, map[string][]*sifxml.StaffPersonal) {
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
				idx := random_seq_gen(fmt.Sprintf("%s-%s", school.RefId().String(), s), len(secondarystaff))
				subjectstaff[s] = append(subjectstaff[s], secondarystaff[idx])
			}
		}
	}
	return primarystaff, secondarystaff, subjectstaff
}

/* split up time table subjects by year level and subject matter */
func timeTableSubjects2yr2subj(tts []*sifxml.TimeTableSubject) (map[string]map[string]*sifxml.TimeTableSubject, []string) {
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
func makePrimaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal, primarystaff []*sifxml.StaffPersonal) []*sifxml.TeachingGroup {
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
func makeSecondaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal, secondarystaff []*sifxml.StaffPersonal) []*sifxml.TeachingGroup {
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
func makeSnrSecondaryTeachingGroups(school *sifxml.SchoolInfo, studentsperyr map[string][]*sifxml.StudentPersonal, subjectstaff map[string][]*sifxml.StaffPersonal, subjects []string, tts2subj map[string]map[string]*sifxml.TimeTableSubject) []*sifxml.TeachingGroup {
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
				tg.SetProperty("KeyLearningArea", teachingGroupKLA(s))
				tg.SetProperty("CurriculumLevel", y)
				ret = append(ret, tg)
			}
		}
	}
	return ret
}

/* presupposed for a single school. Using CurriculumLevel to track year level of group */
func MakeTeachingGroups(school *sifxml.SchoolInfo, staff []*sifxml.StaffPersonal, students []*sifxml.StudentPersonal, tts []*sifxml.TimeTableSubject) []*sifxml.TeachingGroup {
	yrs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	ret := sifxml.TeachingGroupSlice()

	/* split up students */
	studentsperyr := make(map[string][]*sifxml.StudentPersonal)
	for _, y := range yrs {
		studentsperyr[y] = sifxml.StudentPersonalSlice()
	}
	for _, s := range students {
		studentsperyr[s.MostRecent().YearLevel().Code().String()] = append(studentsperyr[s.MostRecent().YearLevel().Code().String()], s)
	}
	tts2subj, subjects := timeTableSubjects2yr2subj(tts)
	primary, secondary, snr_secondary := primaryOrSecondary(studentsperyr)
	primarystaff, secondarystaff, subjectstaff := primaryOrSecondaryStaff(school, staff, primary, secondary, snr_secondary, subjects)

	if primary {
		ret = append(ret, makePrimaryTeachingGroups(school, studentsperyr, primarystaff)...)
	}
	if secondary {
		ret = append(ret, makeSecondaryTeachingGroups(school, studentsperyr, secondarystaff)...)
	}
	if snr_secondary {
		ret = append(ret, makeSnrSecondaryTeachingGroups(school, studentsperyr, subjectstaff, subjects, tts2subj)...)
	}

	return ret
}

/* No attempt whatsoever to avoid room clashes. Using CurriculumLevel to work out year level of group.
Period IDs are 1-3, 5-7 (skip lunch), day IDs are 1-5. assuming single teacher and single room per teaching group.*/
func MakeTimeTableCells(school *sifxml.SchoolInfo, timetable *sifxml.TimeTable, tg []*sifxml.TeachingGroup, staff []*sifxml.StaffPersonal, rooms []*sifxml.RoomInfo, tts []*sifxml.TimeTableSubject) []*sifxml.TimeTableCell {
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
		dayid := timetable.TimeTableDayList().Index(i).DayId().String()
		start_times[dayid] = make(map[string]string)
		for j := 0; j < timetable.TimeTableDayList().Index(i).TimeTablePeriodList().Len(); j++ {
			periodid := timetable.TimeTableDayList().Index(i).TimeTablePeriodList().Index(j).PeriodId().String()
			start_times[dayid][periodid] = timetable.TimeTableDayList().Index(i).TimeTablePeriodList().Index(j).StartTime().String()
		}
	}

	ret := sifxml.TimeTableCellSlice()
	for _, g := range tg {
		yrlvl := g.CurriculumLevel().String()
		if g.TimeTableSubjectRefId_IsNil() {
			/* assign 2 cells a week of each subject available for that year level, in the same room */
			room := rooms[rand.Intn(len(rooms))]
			for _, subj := range tts2subj[yrlvl] {
				for i := 0; i < 2; i++ {
					dayid := randomStringFromSlice(dayids)
					periodid := randomStringFromSlice(periodids)
					for ; seen[dayid][periodid]; periodid = randomStringFromSlice(periodids) {
					}
					cell := Create_TimeTableCell(dayid, periodid, "Teaching", school, timetable, subj, g, room, sifxml.RoomInfoSlice(), staff_map[g.TeacherList().Index(0).StaffPersonalRefId().String()], sifxml.StaffPersonalSlice())
					ret = append(ret, cell)
					g = copyTeachingGroupPeriodFromCell(g, cell, start_times)
				}
			}
		} else {
			/* assign 2 cells a week for each teaching group */
			for i := 0; i < 2; i++ {
				dayid := randomStringFromSlice(dayids)
				periodid := randomStringFromSlice(periodids)
				for ; seen[dayid][periodid]; periodid = randomStringFromSlice(periodids) {
				}
				room := rooms[rand.Intn(len(rooms))]
				cell := Create_TimeTableCell(dayid, periodid, "Teaching", school, timetable, tts_map[g.TimeTableSubjectRefId().String()], g, room, sifxml.RoomInfoSlice(), staff_map[g.TeacherList().Index(0).StaffPersonalRefId().String()], sifxml.StaffPersonalSlice())
				ret = append(ret, cell)
				g = copyTeachingGroupPeriodFromCell(g, cell, start_times)
			}
		}
	}

	return ret
}
