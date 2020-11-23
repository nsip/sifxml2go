package sifxml


type GradingAssignmentScores []GradingAssignmentScore

    type GradingAssignmentScore struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentPersonalLocalId *LocalIdType `xml:"StudentPersonalLocalId" json:"StudentPersonalLocalId"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      GradingAssignmentRefId *String `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      DateGraded *String `xml:"DateGraded,omitempty" json:"DateGraded"`
      ExpectedScore *Bool `xml:"ExpectedScore,omitempty" json:"ExpectedScore"`
      ScorePoints *Int `xml:"ScorePoints,omitempty" json:"ScorePoints"`
      ScorePercent *Float `xml:"ScorePercent,omitempty" json:"ScorePercent"`
      ScoreLetter *String `xml:"ScoreLetter,omitempty" json:"ScoreLetter"`
      ScoreDescription *String `xml:"ScoreDescription,omitempty" json:"ScoreDescription"`
      SubscoreList *NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      TeacherJudgement *String `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement"`
      MarkInfoRefId *String `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      AssignmentScoreIteration *String `xml:"AssignmentScoreIteration,omitempty" json:"AssignmentScoreIteration"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    