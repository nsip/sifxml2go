package sifxml


type GradingAssignmentScores []GradingAssignmentScore

    type GradingAssignmentScore struct {
  gradingassignmentscore `xml:"GradingAssignmentScore"`
}

type gradingassignmentscore struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      StudentPersonalLocalId *LocalIdType `xml:"StudentPersonalLocalId" json:"StudentPersonalLocalId"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      GradingAssignmentRefId *String `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      DateGraded *String `xml:"DateGraded,omitempty" json:"DateGraded,omitempty"`
      ExpectedScore *Bool `xml:"ExpectedScore,omitempty" json:"ExpectedScore,omitempty"`
      ScorePoints *Int `xml:"ScorePoints,omitempty" json:"ScorePoints,omitempty"`
      ScorePercent *Float `xml:"ScorePercent,omitempty" json:"ScorePercent,omitempty"`
      ScoreLetter *String `xml:"ScoreLetter,omitempty" json:"ScoreLetter,omitempty"`
      ScoreDescription *String `xml:"ScoreDescription,omitempty" json:"ScoreDescription,omitempty"`
      SubscoreList *NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList,omitempty"`
      TeacherJudgement *String `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement,omitempty"`
      MarkInfoRefId *String `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId,omitempty"`
      AssignmentScoreIteration *String `xml:"AssignmentScoreIteration,omitempty" json:"AssignmentScoreIteration,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

