package sifxml


type StudentScoreJudgementAgainstStandards []StudentScoreJudgementAgainstStandard

    type StudentScoreJudgementAgainstStandard struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      TermInfoRefId *string `xml:"TermInfoRefId" json:"TermInfoRefId"`
      LocalTermCode *LocalIdType `xml:"LocalTermCode" json:"LocalTermCode"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentStateProvinceId *StateProvinceIdType `xml:"StudentStateProvinceId" json:"StudentStateProvinceId"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId" json:"StudentLocalId"`
      YearLevel *YearLevelType `xml:"YearLevel" json:"YearLevel"`
      TeachingGroupRefId *string `xml:"TeachingGroupRefId" json:"TeachingGroupRefId"`
      ClassLocalId *string `xml:"ClassLocalId" json:"ClassLocalId"`
      StaffPersonalRefId *string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId" json:"StaffLocalId"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList" json:"LearningStandardList"`
      CurriculumCode *LocalIdType `xml:"CurriculumCode" json:"CurriculumCode"`
      CurriculumNodeCode *LocalIdType `xml:"CurriculumNodeCode" json:"CurriculumNodeCode"`
      Score *string `xml:"Score" json:"Score"`
      SpecialCircumstanceLocalCode *LocalIdType `xml:"SpecialCircumstanceLocalCode,omitempty" json:"SpecialCircumstanceLocalCode"`
      ManagedPathwayLocalCode *LocalIdType `xml:"ManagedPathwayLocalCode,omitempty" json:"ManagedPathwayLocalCode"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId" json:"SchoolLocalId"`
      SchoolCommonwealthId *string `xml:"SchoolCommonwealthId" json:"SchoolCommonwealthId"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    