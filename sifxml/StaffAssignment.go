package sifxml


type StaffAssignments []StaffAssignment

    type StaffAssignment struct {
  staffassignment `xml:"StaffAssignment" json:"StaffAssignment"`
}

type staffassignment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      PrimaryAssignment *AUCodeSetsYesOrNoCategoryType `xml:"PrimaryAssignment" json:"PrimaryAssignment"`
      JobStartDate *String `xml:"JobStartDate,omitempty" json:"JobStartDate,omitempty"`
      JobEndDate *String `xml:"JobEndDate,omitempty" json:"JobEndDate,omitempty"`
      JobFTE *Float `xml:"JobFTE,omitempty" json:"JobFTE,omitempty"`
      JobFunction *String `xml:"JobFunction,omitempty" json:"JobFunction,omitempty"`
      EmploymentStatus *AUCodeSetsStaffStatusType `xml:"EmploymentStatus,omitempty" json:"EmploymentStatus,omitempty"`
      StaffSubjectList *StaffSubjectListType `xml:"StaffSubjectList,omitempty" json:"StaffSubjectList,omitempty"`
      StaffActivity *StaffActivityExtensionType `xml:"StaffActivity,omitempty" json:"StaffActivity,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      CasualReliefTeacher *AUCodeSetsYesOrNoCategoryType `xml:"CasualReliefTeacher,omitempty" json:"CasualReliefTeacher,omitempty"`
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup,omitempty"`
      House *String `xml:"House,omitempty" json:"House,omitempty"`
      CalendarSummaryList *CalendarSummaryListType `xml:"CalendarSummaryList,omitempty" json:"CalendarSummaryList,omitempty"`
      PreviousSchoolName *String `xml:"PreviousSchoolName,omitempty" json:"PreviousSchoolName,omitempty"`
      AvailableForTimetable *AUCodeSetsYesOrNoCategoryType `xml:"AvailableForTimetable,omitempty" json:"AvailableForTimetable,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    