package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type AggregateStatisticFacts []AggregateStatisticFact

type ObjectNameType string
type ServiceNameType string
type ObjectType string
type ReportDataObjectType string
type URIOrBinaryType string
type GUIDType string
type MsgIdType GUIDType
type RefIdType GUIDType
type IdRefType RefIdType
type VersionType string
type VersionWithWildcardsType string
type DefinedProtocolsType string
type ExtendedContentType string
type SelectedContentType string
type LocalIdType string

type YearLevelsType struct {
	YearLevel []YearLevelType `xml:"YearLevel" json:"YearLevel"`
}

type YearLevelType struct {
	Code *AUCodeSetsYearLevelCodeType `xml:"Code" json:"Code"`
}

type AUCodeSetsYearLevelCodeType string

var AUCodeSetsYearLevelCodeType_values = []string{"0", "1", "10", "11", "12", "13", "2", "3", "4", "5", "6", "7", "8", "9", "K", "P", "K3", "K4", "CC", "PS", "UG", "11MINUS", "12PLUS", "UGJunSec", "UGPri", "UGSec", "UGSnrSec"}

type AggregateStatisticFact struct {
	RefId                       *RefIdType                    `xml:"RefId,attr" json:"RefId"`
	AggregateStatisticInfoRefId *string                       `xml:"AggregateStatisticInfoRefId" json:"AggregateStatisticInfoRefId"`
	Characteristics             *CharacteristicsType          `xml:"Characteristics" json:"Characteristics"`
	Excluded                    *string                       `xml:"Excluded,omitempty" json:"Excluded"`
	Value                       *float64                      `xml:"Value" json:"Value"`
	StatsCohortYearLevelList    *StatsCohortYearLevelListType `xml:"StatsCohortYearLevelList,omitempty" json:"StatsCohortYearLevelList"`
}

type StatsCohortYearLevelListType struct {
	StatsCohortYearLevel []StatsCohortYearLevelType `xml:"StatsCohortYearLevel" json:"StatsCohortYearLevel"`
}

type StatsCohortYearLevelType struct {
	CohortYearLevel *YearLevelType       `xml:"CohortYearLevel" json:"CohortYearLevel"`
	StatsCohortList *StatsCohortListType `xml:"StatsCohortList" json:"StatsCohortList"`
}

type StatsCohortListType struct {
	StatsCohort []StatsCohortType `xml:"StatsCohort" json:"StatsCohort"`
}

type StatsCohortType struct {
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

type CharacteristicsType struct {
	AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId" json:"AggregateCharacteristicInfoRefId"`
}

func errcheck(err error) {
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

// https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go/30716481#30716481

// Note that we actually didn't allocate anything, the Go compiler did that when we returned the address of the function argument. The Go compiler performs escape analysis, and allocate local variables on the heap (instead of the stack) if they may escape the function. For details, see Is returning a slice of a local array in a Go function safe?
func IntCreate(x int) *int {
	return &x
}

func FloatCreate(x float64) *float64 {
	return &x
}

func StringCreate(x string) *string {
	return &x
}

func BoolCreate(x bool) *bool {
	return &x
}

func CharacteristicsTypeCreate(x CharacteristicsType) *CharacteristicsType {
	return &x
}

func StatsCohortYearLevelTypeCreate(x StatsCohortYearLevelType) *StatsCohortYearLevelType {
	return &x
}

func StatsCohortTypeCreate(x StatsCohortType) *StatsCohortType {
	return &x
}

func StatsCohortListTypeCreate(x StatsCohortListType) *StatsCohortListType {
	return &x
}

func StatsCohortYearLevelListTypeCreate(x StatsCohortYearLevelListType) *StatsCohortYearLevelListType {
	return &x
}

func YearLevelTypeCreate(x YearLevelType) *YearLevelType {
	return &x
}

// Need this templated
func (t *CharacteristicsType) Add(v string) *CharacteristicsType {
	if t == nil {
		t = CharacteristicsTypeCreate(CharacteristicsType{})
	}
	if t.AggregateCharacteristicInfoRefId == nil {
		t.AggregateCharacteristicInfoRefId = make([]string, 0)
	}
	t.AggregateCharacteristicInfoRefId = append(t.AggregateCharacteristicInfoRefId, v)
	return t
}

func (t *StatsCohortYearLevelListType) Add(v StatsCohortYearLevelType) *StatsCohortYearLevelListType {
	if t == nil {
		t = StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
	}
	if t.StatsCohortYearLevel == nil {
		t.StatsCohortYearLevel = make([]StatsCohortYearLevelType, 0)
	}
	t.StatsCohortYearLevel = append(t.StatsCohortYearLevel, v)
	return t
}

func (t *StatsCohortYearLevelListType) AddNew() *StatsCohortYearLevelListType {
	if t == nil {
		t = StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
	}
	if t.StatsCohortYearLevel == nil {
		t.StatsCohortYearLevel = make([]StatsCohortYearLevelType, 0)
	}
	t.StatsCohortYearLevel = append(t.StatsCohortYearLevel, StatsCohortYearLevelType{})
	return t
}

func (t *StatsCohortYearLevelListType) Last() *StatsCohortYearLevelType {
	return &(t.StatsCohortYearLevel[len(t.StatsCohortYearLevel)-1])
}

func (t *StatsCohortListType) Last() *StatsCohortType {
	return &(t.StatsCohort[len(t.StatsCohort)-1])
}

func (t *StatsCohortListType) Add(v StatsCohortType) {
	if t.StatsCohort == nil {
		t.StatsCohort = make([]StatsCohortType, 0)
	}
	t.StatsCohort = append(t.StatsCohort, v)
}

func (t *StatsCohortListType) AddNew() *StatsCohortListType {
	if t == nil {
		t = StatsCohortListTypeCreate(StatsCohortListType{})
	}
	if t.StatsCohort == nil {
		t.StatsCohort = make([]StatsCohortType, 0)
	}
	t.StatsCohort = append(t.StatsCohort, StatsCohortType{})
	return t
}

func setAggregateStatisticFactEmpty() AggregateStatisticFact {
	new := AggregateStatisticFact{}
	new.RefId = ((*RefIdType)(StringCreate("")))
	new.AggregateStatisticInfoRefId = StringCreate("")
	new.Excluded = StringCreate("")
	new.Value = FloatCreate(0)
	arr := CharacteristicsType{}
	arr.Add("0B972D86-44EE-4D79-94F5-FF1C3131772D")
	arr.Add("0B972D86-44EE-4D79-94F5-FF1C3131772E")
	new.Characteristics = &arr
	return new
}

func (n *AggregateStatisticFact) Set(key string, value interface{}) {
	switch key {
	case "RefId":
		n.RefId = ((*RefIdType)(StringCreate(value.(string))))
	case "AggregateStatisticInfoRefId":
		n.AggregateStatisticInfoRefId = StringCreate(value.(string))
	case "Exclude":
		n.Excluded = StringCreate(value.(string))
	case "Value":
		n.Value = FloatCreate(value.(float64))
	case "Characteristics":
		n.Characteristics = CharacteristicsTypeCreate(value.(CharacteristicsType))
	case "StatsCohortYearLevelList":
		n.StatsCohortYearLevelList = StatsCohortYearLevelListTypeCreate(value.(StatsCohortYearLevelListType))
	}
}

func (n *StatsCohortType) Set(key string, value interface{}) *StatsCohortType {
	if n == nil {
		n = StatsCohortTypeCreate(StatsCohortType{})
	}
	switch key {
	case "StatsCohortId":
		n.StatsCohortId = ((*LocalIdType)(StringCreate(value.(string))))
	case "StatsIndigenousStudentType":
		n.StatsIndigenousStudentType = StringCreate(value.(string))
	case "CohortGender":
		n.CohortGender = StringCreate(value.(string))
	case "DaysInReferencePeriod":
		n.DaysInReferencePeriod = IntCreate(value.(int))
	case "PossibleSchoolDays":
		n.PossibleSchoolDays = IntCreate(value.(int))
	case "AttendanceDays":
		n.AttendanceDays = FloatCreate(value.(float64))
	case "AttendanceLess90Percent":
		n.AttendanceLess90Percent = IntCreate(value.(int))
	case "AttendanceGTE90Percent":
		n.AttendanceGTE90Percent = IntCreate(value.(int))
	case "PossibleSchoolDaysGT90PercentAttendance":
		n.PossibleSchoolDaysGT90PercentAttendance = IntCreate(value.(int))
	}
	return n
}

func (n *StatsCohortYearLevelType) Set(key string, value interface{}) {
	switch key {
	case "CohortYearLevel":
		n.CohortYearLevel = YearLevelTypeCreate(value.(YearLevelType))
	case "StatsCohortList":
		n.StatsCohortList = StatsCohortListTypeCreate(value.(StatsCohortListType))
	}
}

func (n *YearLevelType) Set(key string, value interface{}) *YearLevelType {
	if n == nil {
		n = YearLevelTypeCreate(YearLevelType{})
	}
	switch key {
	case "Code":
		n.Code = ((*AUCodeSetsYearLevelCodeType)(StringCreate(value.(string))))
	}
	return n
}

func (n *YearLevelType) New() *YearLevelType {
	return YearLevelTypeCreate(YearLevelType{})
}

func main() {
	var err error
	new := AggregateStatisticFact{}
	output, err := xml.MarshalIndent(new, "", "  ")
	errcheck(err)
	fmt.Println(string(output))

	new = setAggregateStatisticFactEmpty()
	output, err = xml.MarshalIndent(new, "", "  ")
	errcheck(err)
	fmt.Println(string(output))

	new = AggregateStatisticFact{}
	new.Set("RefId", "ABC")
	new.Set("AggregateStatisticInfoRefId", "DEF")
	new.Set("Exclude", "GHI")
	new.Set("Value", 3.45)
	new.Characteristics = new.Characteristics.Add("XYZ")
	new.Characteristics = new.Characteristics.Add("UVW")

	s := StatsCohortType{}
	s.Set("StatsCohortId", "1")
	s.Set("DaysInReferencePeriod", 5)
	s_arr := StatsCohortListType{}
	s_arr.Add(s)
	y := YearLevelType{}
	y.Set("Code", "12")
	s1 := StatsCohortYearLevelType{}
	s1.Set("CohortYearLevel", y)
	s1.Set("StatsCohortList", s_arr)
	new.StatsCohortYearLevelList = new.StatsCohortYearLevelList.Add(s1)

	new.StatsCohortYearLevelList = new.StatsCohortYearLevelList.AddNew()
	// NEED x = x.Set if skipping setting of container struct (CohortYearLevel), which by default is nil
	new.StatsCohortYearLevelList.Last().CohortYearLevel = new.StatsCohortYearLevelList.Last().CohortYearLevel.Set("Code", "11")
	new.StatsCohortYearLevelList.Last().StatsCohortList = new.StatsCohortYearLevelList.Last().StatsCohortList.AddNew()
	// new.StatsCohortYearLevelList.Last().StatsCohortList.Last().Set("StatsCohortId", "22")
	a := new.StatsCohortYearLevelList.Last().StatsCohortList.Last()
	a = a.Set("StatsCohortId", "22")

	new.StatsCohortYearLevelList = new.StatsCohortYearLevelList.AddNew()
	new.StatsCohortYearLevelList.Last().CohortYearLevel = new.StatsCohortYearLevelList.Last().CohortYearLevel.New()
	new.StatsCohortYearLevelList.Last().CohortYearLevel.Set("Code", "13")

	output, err = xml.MarshalIndent(new, "", "  ")
	errcheck(err)
	fmt.Println(string(output))
}
