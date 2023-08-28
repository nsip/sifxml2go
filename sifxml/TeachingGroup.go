package sifxml


type TeachingGroups []TeachingGroup

    type TeachingGroup struct {
  teachinggroup `xml:"TeachingGroup"`
}

type teachinggroup struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      ShortName *String `xml:"ShortName" json:"ShortName"`
      LongName *String `xml:"LongName,omitempty" json:"LongName,omitempty"`
      GroupType *String `xml:"GroupType,omitempty" json:"GroupType,omitempty"`
      Set *String `xml:"Set,omitempty" json:"Set,omitempty"`
      Block *String `xml:"Block,omitempty" json:"Block,omitempty"`
      CurriculumLevel *String `xml:"CurriculumLevel,omitempty" json:"CurriculumLevel,omitempty"`
      SchoolInfoRefId *RefIdType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId,omitempty"`
      SchoolCourseLocalId *LocalIdType `xml:"SchoolCourseLocalId,omitempty" json:"SchoolCourseLocalId,omitempty"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      TimeTableSubjectLocalId *LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId,omitempty"`
      KeyLearningArea *AUCodeSetsACStrandType `xml:"KeyLearningArea,omitempty" json:"KeyLearningArea,omitempty"`
      Semester *Int `xml:"Semester,omitempty" json:"Semester,omitempty"`
      StudentList *StudentListType `xml:"StudentList,omitempty" json:"StudentList,omitempty"`
      TeacherList *TeacherListType `xml:"TeacherList,omitempty" json:"TeacherList,omitempty"`
      MinClassSize *Int `xml:"MinClassSize,omitempty" json:"MinClassSize,omitempty"`
      MaxClassSize *Int `xml:"MaxClassSize,omitempty" json:"MaxClassSize,omitempty"`
      TeachingGroupPeriodList *TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList,omitempty" json:"TeachingGroupPeriodList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

