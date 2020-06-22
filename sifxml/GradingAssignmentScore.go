package sifxml


    type GradingAssignmentScore struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      StudentPersonalLocalId LocalIdType `xml:"StudentPersonalLocalId"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      GradingAssignmentRefId string `xml:"GradingAssignmentRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      DateGraded string `xml:"DateGraded"`
      ExpectedScore string `xml:"ExpectedScore"`
      ScorePoints string `xml:"ScorePoints"`
      ScorePercent string `xml:"ScorePercent"`
      ScoreLetter string `xml:"ScoreLetter"`
      ScoreDescription string `xml:"ScoreDescription"`
      SubscoreList NAPSubscoreListType `xml:"SubscoreList"`
      TeacherJudgement string `xml:"TeacherJudgement"`
      MarkInfoRefId string `xml:"MarkInfoRefId"`
      AssignmentScoreIteration string `xml:"AssignmentScoreIteration"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    