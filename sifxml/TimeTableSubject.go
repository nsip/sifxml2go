package sifxml


type TimeTableSubjects []TimeTableSubject

    type TimeTableSubject struct {
  timetablesubject `xml:"TimeTableSubject"`
}

type timetablesubject struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId" json:"SubjectLocalId"`
      AcademicYear *YearLevelType `xml:"AcademicYear,omitempty" json:"AcademicYear,omitempty"`
      AcademicYearRange *YearRangeType `xml:"AcademicYearRange,omitempty" json:"AcademicYearRange,omitempty"`
      CourseLocalId *LocalIdType `xml:"CourseLocalId,omitempty" json:"CourseLocalId,omitempty"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId,omitempty"`
      Faculty *String `xml:"Faculty,omitempty" json:"Faculty,omitempty"`
      SubjectShortName *String `xml:"SubjectShortName,omitempty" json:"SubjectShortName,omitempty"`
      SubjectLongName *String `xml:"SubjectLongName" json:"SubjectLongName"`
      SubjectType *String `xml:"SubjectType,omitempty" json:"SubjectType,omitempty"`
      ProposedMaxClassSize *Float `xml:"ProposedMaxClassSize,omitempty" json:"ProposedMaxClassSize,omitempty"`
      ProposedMinClassSize *Float `xml:"ProposedMinClassSize,omitempty" json:"ProposedMinClassSize,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      Semester *Int `xml:"Semester,omitempty" json:"Semester,omitempty"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

