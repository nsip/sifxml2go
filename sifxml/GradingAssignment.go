package sifxml


type GradingAssignments []GradingAssignment

    type GradingAssignment struct {
  gradingassignment `xml:"GradingAssignment"`
}

type gradingassignment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId,omitempty"`
      StudentPersonalRefIdList *StudentsType `xml:"StudentPersonalRefIdList,omitempty" json:"StudentPersonalRefIdList,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      GradingCategory *String `xml:"GradingCategory,omitempty" json:"GradingCategory,omitempty"`
      Description *String `xml:"Description" json:"Description"`
      PointsPossible *Int `xml:"PointsPossible" json:"PointsPossible"`
      CreateDate *String `xml:"CreateDate,omitempty" json:"CreateDate,omitempty"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate,omitempty"`
      Weight *Float `xml:"Weight,omitempty" json:"Weight,omitempty"`
      MaxAttemptsAllowed *Int `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed,omitempty"`
      DetailedDescriptionURL *String `xml:"DetailedDescriptionURL,omitempty" json:"DetailedDescriptionURL,omitempty"`
      DetailedDescriptionBinary *String `xml:"DetailedDescriptionBinary,omitempty" json:"DetailedDescriptionBinary,omitempty"`
      AssessmentType *String `xml:"AssessmentType,omitempty" json:"AssessmentType,omitempty"`
      LevelAssessed *String `xml:"LevelAssessed,omitempty" json:"LevelAssessed,omitempty"`
      AssignmentPurpose *String `xml:"AssignmentPurpose,omitempty" json:"AssignmentPurpose,omitempty"`
      SubAssignmentList *AssignmentListType `xml:"SubAssignmentList,omitempty" json:"SubAssignmentList,omitempty"`
      RubricScoringGuide *GenericRubricType `xml:"RubricScoringGuide,omitempty" json:"RubricScoringGuide,omitempty"`
      PrerequisiteList *PrerequisitesType `xml:"PrerequisiteList,omitempty" json:"PrerequisiteList,omitempty"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

