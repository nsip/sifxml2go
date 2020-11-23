package sifxml


type StudentGrades []StudentGrade

    type StudentGrade struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      TeachingGroupShortName *String `xml:"TeachingGroupShortName,omitempty" json:"TeachingGroupShortName"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Markers *StudentGradeMarkersListType `xml:"Markers,omitempty" json:"Markers"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      LearningArea *ACStrandSubjectAreaType `xml:"LearningArea,omitempty" json:"LearningArea"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList"`
      GradingScoreList *GradingScoreListType `xml:"GradingScoreList,omitempty" json:"GradingScoreList"`
      Grade *GradeType `xml:"Grade,omitempty" json:"Grade"`
      TeacherJudgement *String `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    