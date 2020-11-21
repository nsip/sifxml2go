package funcapi

//
// external type
//
type StatsCohortListType struct {
	statsCohortListType
}

//
// internal type
//
type statsCohortListType struct {
	StatsCohort []StatsCohortType `xml:"StatsCohort" json:"StatsCohort"`
}

//
// helper for user
//
func (t *StatsCohortListType) Append() *StatsCohortListType {
	if t.statsCohortListType.StatsCohort == nil {
		t.statsCohortListType.StatsCohort = make([]StatsCohortType, 0)
	}
	return t
}

//
// expose the internal list
//
func (t *StatsCohortListType) StatsCohort() *StatsCohortType {
	if t.statsCohortListType.StatsCohort == nil {
		t.statsCohortListType.StatsCohort = make([]StatsCohortType, 0)
	}
	sct := StatsCohortType{}
	t.statsCohortListType.StatsCohort = append(t.statsCohortListType.StatsCohort, sct)
	return &t.statsCohortListType.StatsCohort[len(t.statsCohortListType.StatsCohort)-1]
}

//
// access the most recently added member of the list
//
func (t *StatsCohortListType) Last() *StatsCohortType {
	if t.statsCohortListType.StatsCohort == nil {
		return t.StatsCohort()
	}
	return &t.statsCohortListType.StatsCohort[len(t.statsCohortListType.StatsCohort)-1]

}
