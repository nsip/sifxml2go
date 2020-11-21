package funcapi

//
// external type
//
type StatsCohortYearLevelType struct {
	statsCohortYearLevelType
}

//
// internal type
//
type statsCohortYearLevelType struct {
	CohortYearLevel *YearLevelType       `xml:"CohortYearLevel" json:"CohortYearLevel"`
	StatsCohortList *StatsCohortListType `xml:"StatsCohortList" json:"StatsCohortList"`
}

//
// expose member object, create if needed
//
func (s *StatsCohortYearLevelType) CohortYearLevel() *YearLevelType {
	if s.statsCohortYearLevelType.CohortYearLevel == nil {
		s.statsCohortYearLevelType.CohortYearLevel = &YearLevelType{}
	}
	return s.statsCohortYearLevelType.CohortYearLevel
}

//
// expose member object, create if needed
//
func (t *StatsCohortYearLevelType) StatsCohortList() *StatsCohortListType {
	if t.statsCohortYearLevelType.StatsCohortList == nil {
		t.statsCohortYearLevelType.StatsCohortList = &StatsCohortListType{}
	}
	return t.statsCohortYearLevelType.StatsCohortList
}
