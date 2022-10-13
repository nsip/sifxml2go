package sifxml


type SchoolCourseInfos []SchoolCourseInfo

    type SchoolCourseInfo struct {
  schoolcourseinfo `xml:"SchoolCourseInfo" json:"SchoolCourseInfo"`
}

type schoolcourseinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId,omitempty"`
      CourseCode *String `xml:"CourseCode" json:"CourseCode"`
      StateCourseCode *String `xml:"StateCourseCode,omitempty" json:"StateCourseCode,omitempty"`
      DistrictCourseCode *String `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode,omitempty"`
      SubjectAreaList *SubjectAreaListType `xml:"SubjectAreaList,omitempty" json:"SubjectAreaList,omitempty"`
      CourseTitle *String `xml:"CourseTitle" json:"CourseTitle"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      InstructionalLevel *String `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel,omitempty"`
      CourseCredits *String `xml:"CourseCredits,omitempty" json:"CourseCredits,omitempty"`
      CoreAcademicCourse *AUCodeSetsYesOrNoCategoryType `xml:"CoreAcademicCourse,omitempty" json:"CoreAcademicCourse,omitempty"`
      GraduationRequirement *AUCodeSetsYesOrNoCategoryType `xml:"GraduationRequirement,omitempty" json:"GraduationRequirement,omitempty"`
      Department *String `xml:"Department,omitempty" json:"Department,omitempty"`
      CourseContent *String `xml:"CourseContent,omitempty" json:"CourseContent,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

