package sifxml


type TimeTableSubjects []TimeTableSubject

    type TimeTableSubject struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId" json:"SubjectLocalId"`
      AcademicYear *YearLevelType `xml:"AcademicYear" json:"AcademicYear"`
      AcademicYearRange *YearRangeType `xml:"AcademicYearRange" json:"AcademicYearRange"`
      CourseLocalId *LocalIdType `xml:"CourseLocalId,omitempty" json:"CourseLocalId"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      Faculty *string `xml:"Faculty,omitempty" json:"Faculty"`
      SubjectShortName *string `xml:"SubjectShortName,omitempty" json:"SubjectShortName"`
      SubjectLongName *string `xml:"SubjectLongName" json:"SubjectLongName"`
      SubjectType *string `xml:"SubjectType,omitempty" json:"SubjectType"`
      ProposedMaxClassSize *float64 `xml:"ProposedMaxClassSize" json:"ProposedMaxClassSize"`
      ProposedMinClassSize *float64 `xml:"ProposedMinClassSize" json:"ProposedMinClassSize"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      Semester *int `xml:"Semester" json:"Semester"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    