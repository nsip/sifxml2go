package sifxml


type StudentScoreJudgementAgainstStandards []StudentScoreJudgementAgainstStandard

    type StudentScoreJudgementAgainstStandard struct {
  studentscorejudgementagainststandard `xml:"StudentScoreJudgementAgainstStandard" json:"StudentScoreJudgementAgainstStandard"`
}

type studentscorejudgementagainststandard struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId,omitempty"`
      LocalTermCode *LocalIdType `xml:"LocalTermCode,omitempty" json:"LocalTermCode,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      StudentStateProvinceId *StateProvinceIdType `xml:"StudentStateProvinceId,omitempty" json:"StudentStateProvinceId,omitempty"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId,omitempty" json:"StudentLocalId,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel" json:"YearLevel"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId,omitempty"`
      ClassLocalId *String `xml:"ClassLocalId,omitempty" json:"ClassLocalId,omitempty"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId,omitempty"`
      LearningStandardList *LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList,omitempty"`
      CurriculumCode *LocalIdType `xml:"CurriculumCode,omitempty" json:"CurriculumCode,omitempty"`
      CurriculumNodeCode *LocalIdType `xml:"CurriculumNodeCode,omitempty" json:"CurriculumNodeCode,omitempty"`
      Score *String `xml:"Score" json:"Score"`
      SpecialCircumstanceLocalCode *LocalIdType `xml:"SpecialCircumstanceLocalCode,omitempty" json:"SpecialCircumstanceLocalCode,omitempty"`
      ManagedPathwayLocalCode *LocalIdType `xml:"ManagedPathwayLocalCode,omitempty" json:"ManagedPathwayLocalCode,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolCommonwealthId *String `xml:"SchoolCommonwealthId,omitempty" json:"SchoolCommonwealthId,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    