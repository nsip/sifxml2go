package sifxml


type GradingAssignments []GradingAssignment

    type GradingAssignment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      StudentPersonalRefIdList *StudentsType `xml:"StudentPersonalRefIdList,omitempty" json:"StudentPersonalRefIdList"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      GradingCategory *String `xml:"GradingCategory,omitempty" json:"GradingCategory"`
      Description *String `xml:"Description" json:"Description"`
      PointsPossible *Int `xml:"PointsPossible" json:"PointsPossible"`
      CreateDate *String `xml:"CreateDate,omitempty" json:"CreateDate"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate"`
      Weight *Float `xml:"Weight,omitempty" json:"Weight"`
      MaxAttemptsAllowed *Int `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed"`
      DetailedDescriptionURL *String `xml:"DetailedDescriptionURL,omitempty" json:"DetailedDescriptionURL"`
      DetailedDescriptionBinary *String `xml:"DetailedDescriptionBinary,omitempty" json:"DetailedDescriptionBinary"`
      AssessmentType *String `xml:"AssessmentType,omitempty" json:"AssessmentType"`
      LevelAssessed *String `xml:"LevelAssessed,omitempty" json:"LevelAssessed"`
      AssignmentPurpose *String `xml:"AssignmentPurpose,omitempty" json:"AssignmentPurpose"`
      SubAssignmentList *AssignmentListType `xml:"SubAssignmentList,omitempty" json:"SubAssignmentList"`
      RubricScoringGuide *GenericRubricType `xml:"RubricScoringGuide,omitempty" json:"RubricScoringGuide"`
      PrerequisiteList *PrerequisitesType `xml:"PrerequisiteList,omitempty" json:"PrerequisiteList"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    