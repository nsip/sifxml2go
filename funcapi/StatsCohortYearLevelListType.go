package funcapi

//
// exposed object
//
type StatsCohortYearLevelListType struct {
	statsCohortYearLevelListType
}

//
// internal object for serialisation
//
type statsCohortYearLevelListType struct {
	StatsCohortYearLevel []StatsCohortYearLevelType `xml:"StatsCohortYearLevel" json:"StatsCohortYearLevel"`
}

//
// if the type of the list is a complex object
// append creates a member of that type, initilaises the list
// if necessary
//
func (s *StatsCohortYearLevelListType) Append() *StatsCohortYearLevelType {
	if s.statsCohortYearLevelListType.StatsCohortYearLevel == nil {
		s.statsCohortYearLevelListType.StatsCohortYearLevel = make([]StatsCohortYearLevelType, 0)
	}
	scylt := StatsCohortYearLevelType{}
	s.statsCohortYearLevelListType.StatsCohortYearLevel = append(s.statsCohortYearLevelListType.StatsCohortYearLevel, scylt)
	return &s.statsCohortYearLevelListType.StatsCohortYearLevel[len(s.statsCohortYearLevelListType.StatsCohortYearLevel)-1]
}

//
// allows reference to most recently added memebr of list,
// allows setting of (for example) another object that may be part of the
// member
//
func (s *StatsCohortYearLevelListType) Last() *StatsCohortYearLevelType {
	if s.statsCohortYearLevelListType.StatsCohortYearLevel == nil {
		return s.Append()
	}
	return &s.statsCohortYearLevelListType.StatsCohortYearLevel[len(s.statsCohortYearLevelListType.StatsCohortYearLevel)-1]
}
