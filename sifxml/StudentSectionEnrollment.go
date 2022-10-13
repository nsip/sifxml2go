package sifxml


type StudentSectionEnrollments []StudentSectionEnrollment

    type StudentSectionEnrollment struct {
  studentsectionenrollment `xml:"StudentSectionEnrollment" json:"StudentSectionEnrollment"`
}

type studentsectionenrollment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SectionInfoRefId *String `xml:"SectionInfoRefId" json:"SectionInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      EntryDate *String `xml:"EntryDate,omitempty" json:"EntryDate,omitempty"`
      ExitDate *String `xml:"ExitDate,omitempty" json:"ExitDate,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

