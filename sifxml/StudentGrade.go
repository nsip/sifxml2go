package sifxml


    type StudentGrade struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      Homegroup string `xml:"Homegroup"`
      YearLevel YearLevelType `xml:"YearLevel"`
      TeachingGroupShortName string `xml:"TeachingGroupShortName"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      Markers StudentGradeMarkersListType `xml:"Markers"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      TermInfoRefId string `xml:"TermInfoRefId"`
      Description string `xml:"Description"`
      LearningArea ACStrandSubjectAreaType `xml:"LearningArea"`
      LearningStandardList LearningStandardListType `xml:"LearningStandardList"`
      GradingScoreList GradingScoreListType `xml:"GradingScoreList"`
      Grade GradeType `xml:"Grade"`
      TeacherJudgement string `xml:"TeacherJudgement"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    