package sifxml


type StudentSectionEnrollments []StudentSectionEnrollment

    type StudentSectionEnrollment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SectionInfoRefId *String `xml:"SectionInfoRefId" json:"SectionInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      EntryDate *String `xml:"EntryDate,omitempty" json:"EntryDate"`
      ExitDate *String `xml:"ExitDate,omitempty" json:"ExitDate"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    