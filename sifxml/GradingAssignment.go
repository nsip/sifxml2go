package sifxml


    type GradingAssignment struct {
        RefId RefIdType `xml:"RefId,attr"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId"`
      StudentPersonalRefIdList StudentsType `xml:"StudentPersonalRefIdList"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      GradingCategory string `xml:"GradingCategory"`
      Description string `xml:"Description"`
      PointsPossible string `xml:"PointsPossible"`
      CreateDate string `xml:"CreateDate"`
      DueDate string `xml:"DueDate"`
      Weight string `xml:"Weight"`
      MaxAttemptsAllowed string `xml:"MaxAttemptsAllowed"`
      DetailedDescriptionURL string `xml:"DetailedDescriptionURL"`
      DetailedDescriptionBinary string `xml:"DetailedDescriptionBinary"`
      AssessmentType string `xml:"AssessmentType"`
      LevelAssessed string `xml:"LevelAssessed"`
      AssignmentPurpose string `xml:"AssignmentPurpose"`
      SubAssignmentList AssignmentListType `xml:"SubAssignmentList"`
      RubricScoringGuide GenericRubricType `xml:"RubricScoringGuide"`
      PrerequisiteList PrerequisitesType `xml:"PrerequisiteList"`
      LearningStandardList LearningStandardListType `xml:"LearningStandardList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    