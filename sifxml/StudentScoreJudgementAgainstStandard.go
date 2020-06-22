package sifxml


    type StudentScoreJudgementAgainstStandard struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      TermInfoRefId string `xml:"TermInfoRefId"`
      LocalTermCode LocalIdType `xml:"LocalTermCode"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      StudentStateProvinceId StateProvinceIdType `xml:"StudentStateProvinceId"`
      StudentLocalId LocalIdType `xml:"StudentLocalId"`
      YearLevel YearLevelType `xml:"YearLevel"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId"`
      ClassLocalId string `xml:"ClassLocalId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      LearningStandardList LearningStandardListType `xml:"LearningStandardList"`
      CurriculumCode LocalIdType `xml:"CurriculumCode"`
      CurriculumNodeCode LocalIdType `xml:"CurriculumNodeCode"`
      Score string `xml:"Score"`
      SpecialCircumstanceLocalCode LocalIdType `xml:"SpecialCircumstanceLocalCode"`
      ManagedPathwayLocalCode LocalIdType `xml:"ManagedPathwayLocalCode"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      SchoolCommonwealthId string `xml:"SchoolCommonwealthId"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    