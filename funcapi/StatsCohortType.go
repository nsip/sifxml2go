package funcapi

import "fmt"

//
// externally visible object
//
type StatsCohortType struct {
	statsCohortType
}

//
// internal object
//
type statsCohortType struct {
	StatsCohortId                           *LocalIdType `xml:"StatsCohortId" json:"StatsCohortId"`
	StatsIndigenousStudentType              *string      `xml:"StatsIndigenousStudentType" json:"StatsIndigenousStudentType"`
	CohortGender                            *string      `xml:"CohortGender" json:"CohortGender"`
	DaysInReferencePeriod                   *int         `xml:"DaysInReferencePeriod" json:"DaysInReferencePeriod"`
	PossibleSchoolDays                      *int         `xml:"PossibleSchoolDays" json:"PossibleSchoolDays"`
	AttendanceDays                          *float64     `xml:"AttendanceDays" json:"AttendanceDays"`
	AttendanceLess90Percent                 *int         `xml:"AttendanceLess90Percent" json:"AttendanceLess90Percent"`
	AttendanceGTE90Percent                  *int         `xml:"AttendanceGTE90Percent" json:"AttendanceGTE90Percent"`
	PossibleSchoolDaysGT90PercentAttendance *int         `xml:"PossibleSchoolDaysGT90PercentAttendance" json:"PossibleSchoolDaysGT90PercentAttendance"`
}

//
// allow user to set object properies, handle type conversions etc. for user
//
func (t *StatsCohortType) SetProperty(key string, value interface{}) *StatsCohortType {

	switch key {
	case "StatsCohortId":
		if v, ok := LocalIdPointer(value); ok {
			t.StatsCohortId = v
		}
	case "StatsIndigenousStudentType":
		if v, ok := StringPointer(value); ok {
			t.StatsIndigenousStudentType = v
		}
	case "CohortGender":
		if v, ok := StringPointer(value); ok {
			t.CohortGender = v
		}
	case "DaysInReferencePeriod":
		if v, ok := IntPointer(value); ok {
			t.DaysInReferencePeriod = v
		}
	case "PossibleSchoolDays":
		if v, ok := IntPointer(value); ok {
			t.PossibleSchoolDays = v
		}
	case "AttendanceDays":
		if v, ok := FloatPointer(value); ok {
			t.AttendanceDays = v
		}
	case "AttendanceLess90Percent":
		if v, ok := IntPointer(value); ok {
			t.AttendanceLess90Percent = v
		}
	case "AttendanceGTE90Percent":
		if v, ok := IntPointer(value); ok {
			t.AttendanceGTE90Percent = v
		}
	case "PossibleSchoolDaysGT90PercentAttendance":
		if v, ok := IntPointer(value); ok {
			t.PossibleSchoolDaysGT90PercentAttendance = v
		}
	default:
		fmt.Printf("Warning: %T has no property: %s\n", t, key)
	}

	return t
}

//
// lets user set many properties at once
//
func (t *StatsCohortType) SetProperties(props ...Prop) *StatsCohortType {
	for _, p := range props {
		t.SetProperty(p.Key, p.Value)
	}
	return t
}
