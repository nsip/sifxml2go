package sifxml


type SectionInfos []SectionInfo

    type SectionInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolCourseInfoRefId *string `xml:"SchoolCourseInfoRefId" json:"SchoolCourseInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      TermInfoRefId *string `xml:"TermInfoRefId" json:"TermInfoRefId"`
      MediumOfInstruction *MediumOfInstructionType `xml:"MediumOfInstruction,omitempty" json:"MediumOfInstruction"`
      LanguageOfInstruction *LanguageOfInstructionType `xml:"LanguageOfInstruction,omitempty" json:"LanguageOfInstruction"`
      LocationOfInstruction *LocationOfInstructionType `xml:"LocationOfInstruction,omitempty" json:"LocationOfInstruction"`
      SummerSchool *string `xml:"SummerSchool,omitempty" json:"SummerSchool"`
      SchoolCourseInfoOverride *SchoolCourseInfoOverrideType `xml:"SchoolCourseInfoOverride,omitempty" json:"SchoolCourseInfoOverride"`
      CourseSectionCode *string `xml:"CourseSectionCode,omitempty" json:"CourseSectionCode"`
      SectionCode *string `xml:"SectionCode,omitempty" json:"SectionCode"`
      CountForAttendance *string `xml:"CountForAttendance,omitempty" json:"CountForAttendance"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    