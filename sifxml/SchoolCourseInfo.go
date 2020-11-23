package sifxml


type SchoolCourseInfos []SchoolCourseInfo

    type SchoolCourseInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      TermInfoRefId *String `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      CourseCode *String `xml:"CourseCode" json:"CourseCode"`
      StateCourseCode *String `xml:"StateCourseCode,omitempty" json:"StateCourseCode"`
      DistrictCourseCode *String `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode"`
      SubjectAreaList *SubjectAreaListType `xml:"SubjectAreaList,omitempty" json:"SubjectAreaList"`
      CourseTitle *String `xml:"CourseTitle" json:"CourseTitle"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      InstructionalLevel *String `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel"`
      CourseCredits *String `xml:"CourseCredits,omitempty" json:"CourseCredits"`
      CoreAcademicCourse *AUCodeSetsYesOrNoCategoryType `xml:"CoreAcademicCourse,omitempty" json:"CoreAcademicCourse"`
      GraduationRequirement *AUCodeSetsYesOrNoCategoryType `xml:"GraduationRequirement,omitempty" json:"GraduationRequirement"`
      Department *String `xml:"Department,omitempty" json:"Department"`
      CourseContent *String `xml:"CourseContent,omitempty" json:"CourseContent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    