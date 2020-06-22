package sifxml


    type SectionInfo struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolCourseInfoRefId string `xml:"SchoolCourseInfoRefId"`
      LocalId LocalIdType `xml:"LocalId"`
      Description string `xml:"Description"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      TermInfoRefId string `xml:"TermInfoRefId"`
      MediumOfInstruction MediumOfInstructionType `xml:"MediumOfInstruction"`
      LanguageOfInstruction LanguageOfInstructionType `xml:"LanguageOfInstruction"`
      LocationOfInstruction LocationOfInstructionType `xml:"LocationOfInstruction"`
      SummerSchool string `xml:"SummerSchool"`
      SchoolCourseInfoOverride SchoolCourseInfoOverrideType `xml:"SchoolCourseInfoOverride"`
      CourseSectionCode string `xml:"CourseSectionCode"`
      SectionCode string `xml:"SectionCode"`
      CountForAttendance string `xml:"CountForAttendance"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    