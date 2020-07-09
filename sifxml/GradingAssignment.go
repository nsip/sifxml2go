package sifxml


type GradingAssignments []GradingAssignment

    type GradingAssignment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TeachingGroupRefId *string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      StudentPersonalRefIdList *StudentsType `xml:"StudentPersonalRefIdList,omitempty" json:"StudentPersonalRefIdList"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      GradingCategory *string `xml:"GradingCategory,omitempty" json:"GradingCategory"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      PointsPossible *int `xml:"PointsPossible,omitempty" json:"PointsPossible"`
      CreateDate *string `xml:"CreateDate,omitempty" json:"CreateDate"`
      DueDate *string `xml:"DueDate,omitempty" json:"DueDate"`
      Weight *float64 `xml:"Weight,omitempty" json:"Weight"`
      MaxAttemptsAllowed *int `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed"`
      DetailedDescriptionURL *string `xml:"DetailedDescriptionURL,omitempty" json:"DetailedDescriptionURL"`
      DetailedDescriptionBinary *string `xml:"DetailedDescriptionBinary,omitempty" json:"DetailedDescriptionBinary"`
      AssessmentType *string `xml:"AssessmentType,omitempty" json:"AssessmentType"`
      LevelAssessed *string `xml:"LevelAssessed,omitempty" json:"LevelAssessed"`
      AssignmentPurpose *string `xml:"AssignmentPurpose,omitempty" json:"AssignmentPurpose"`
      SubAssignmentList *AssignmentListType `xml:"SubAssignmentList,omitempty" json:"SubAssignmentList"`
      RubricScoringGuide *GenericRubricType `xml:"RubricScoringGuide,omitempty" json:"RubricScoringGuide"`
      PrerequisiteList *PrerequisitesType `xml:"PrerequisiteList,omitempty" json:"PrerequisiteList"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    