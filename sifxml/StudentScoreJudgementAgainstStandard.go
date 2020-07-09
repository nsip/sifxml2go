package sifxml


type StudentScoreJudgementAgainstStandards []StudentScoreJudgementAgainstStandard

    type StudentScoreJudgementAgainstStandard struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      TermInfoRefId *string `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      LocalTermCode *LocalIdType `xml:"LocalTermCode,omitempty" json:"LocalTermCode"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentStateProvinceId *StateProvinceIdType `xml:"StudentStateProvinceId,omitempty" json:"StudentStateProvinceId"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId,omitempty" json:"StudentLocalId"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      TeachingGroupRefId *string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      ClassLocalId *string `xml:"ClassLocalId,omitempty" json:"ClassLocalId"`
      StaffPersonalRefId *string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList"`
      CurriculumCode *LocalIdType `xml:"CurriculumCode,omitempty" json:"CurriculumCode"`
      CurriculumNodeCode *LocalIdType `xml:"CurriculumNodeCode,omitempty" json:"CurriculumNodeCode"`
      Score *string `xml:"Score,omitempty" json:"Score"`
      SpecialCircumstanceLocalCode *LocalIdType `xml:"SpecialCircumstanceLocalCode,omitempty" json:"SpecialCircumstanceLocalCode"`
      ManagedPathwayLocalCode *LocalIdType `xml:"ManagedPathwayLocalCode,omitempty" json:"ManagedPathwayLocalCode"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolCommonwealthId *string `xml:"SchoolCommonwealthId,omitempty" json:"SchoolCommonwealthId"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    