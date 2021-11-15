package sifxml


type StudentContactPersonals []StudentContactPersonal

    type StudentContactPersonal struct {
  studentcontactpersonal `xml:"StudentContactPersonal" json:"StudentContactPersonal"`
}

type studentcontactpersonal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList,omitempty"`
      PersonInfo *PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      EmploymentType *AUCodeSetsEmploymentTypeType `xml:"EmploymentType,omitempty" json:"EmploymentType,omitempty"`
      SchoolEducationalLevel *EducationalLevelType `xml:"SchoolEducationalLevel,omitempty" json:"SchoolEducationalLevel,omitempty"`
      NonSchoolEducation *AUCodeSetsNonSchoolEducationType `xml:"NonSchoolEducation,omitempty" json:"NonSchoolEducation,omitempty"`
      Employment *xsnormalizedString `xml:"Employment,omitempty" json:"Employment,omitempty"`
      Workplace *xsnormalizedString `xml:"Workplace,omitempty" json:"Workplace,omitempty"`
      WorkingWithChildrenCheck *WorkingWithChildrenCheckType `xml:"WorkingWithChildrenCheck,omitempty" json:"WorkingWithChildrenCheck,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    