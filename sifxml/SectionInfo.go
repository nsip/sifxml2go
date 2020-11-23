package sifxml


type SectionInfos []SectionInfo

    type SectionInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolCourseInfoRefId *String `xml:"SchoolCourseInfoRefId" json:"SchoolCourseInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      MediumOfInstruction *MediumOfInstructionType `xml:"MediumOfInstruction,omitempty" json:"MediumOfInstruction"`
      LanguageOfInstruction *LanguageOfInstructionType `xml:"LanguageOfInstruction,omitempty" json:"LanguageOfInstruction"`
      LocationOfInstruction *LocationOfInstructionType `xml:"LocationOfInstruction,omitempty" json:"LocationOfInstruction"`
      SummerSchool *String `xml:"SummerSchool,omitempty" json:"SummerSchool"`
      SchoolCourseInfoOverride *SchoolCourseInfoOverrideType `xml:"SchoolCourseInfoOverride,omitempty" json:"SchoolCourseInfoOverride"`
      CourseSectionCode *String `xml:"CourseSectionCode,omitempty" json:"CourseSectionCode"`
      SectionCode *String `xml:"SectionCode,omitempty" json:"SectionCode"`
      CountForAttendance *String `xml:"CountForAttendance,omitempty" json:"CountForAttendance"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    