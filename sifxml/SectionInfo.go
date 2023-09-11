package sifxml



    type SectionInfos struct {
      sectioninfos `json:"SectionInfos"`
    }

    type sectioninfos struct {
      SectionInfo []sectioninfo `json:"SectionInfo"`
    }

    type SectionInfo struct {
      sectioninfo `xml:"SectionInfo" json:"SectionInfo"`
     }

     type sectioninfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolCourseInfoRefId *String `xml:"SchoolCourseInfoRefId" json:"SchoolCourseInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId,omitempty"`
      MediumOfInstruction *MediumOfInstructionType `xml:"MediumOfInstruction,omitempty" json:"MediumOfInstruction,omitempty"`
      LanguageOfInstruction *LanguageOfInstructionType `xml:"LanguageOfInstruction,omitempty" json:"LanguageOfInstruction,omitempty"`
      LocationOfInstruction *LocationOfInstructionType `xml:"LocationOfInstruction,omitempty" json:"LocationOfInstruction,omitempty"`
      SummerSchool *String `xml:"SummerSchool,omitempty" json:"SummerSchool,omitempty"`
      SchoolCourseInfoOverride *SchoolCourseInfoOverrideType `xml:"SchoolCourseInfoOverride,omitempty" json:"SchoolCourseInfoOverride,omitempty"`
      CourseSectionCode *String `xml:"CourseSectionCode,omitempty" json:"CourseSectionCode,omitempty"`
      SectionCode *String `xml:"SectionCode,omitempty" json:"SectionCode,omitempty"`
      CountForAttendance *String `xml:"CountForAttendance,omitempty" json:"CountForAttendance,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

