package sifxml


type StudentGrades []StudentGrade

    type StudentGrade struct {
  studentgrade `xml:"StudentGrade" json:"StudentGrade"`
}

type studentgrade struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel,omitempty"`
      TeachingGroupShortName *String `xml:"TeachingGroupShortName,omitempty" json:"TeachingGroupShortName,omitempty"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId,omitempty"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      Markers *StudentGradeMarkersListType `xml:"Markers,omitempty" json:"Markers,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId,omitempty"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      LearningArea *ACStrandSubjectAreaType `xml:"LearningArea,omitempty" json:"LearningArea,omitempty"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList,omitempty"`
      GradingScoreList *GradingScoreListType `xml:"GradingScoreList,omitempty" json:"GradingScoreList,omitempty"`
      Grade *GradeType `xml:"Grade,omitempty" json:"Grade,omitempty"`
      TeacherJudgement *String `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    