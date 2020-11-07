package sifxml

import _ "github.com/creasty/defaults"


type StaffAssignments []StaffAssignment

    type StaffAssignment struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      PrimaryAssignment AUCodeSetsYesOrNoCategoryType `xml:"PrimaryAssignment" json:"PrimaryAssignment"`
      JobStartDate string `xml:"JobStartDate,omitempty" json:"JobStartDate"`
      JobEndDate string `xml:"JobEndDate,omitempty" json:"JobEndDate"`
      JobFTE float64 `default:"-2147483648" xml:"JobFTE" json:"JobFTE"`
      JobFunction string `xml:"JobFunction,omitempty" json:"JobFunction"`
      EmploymentStatus AUCodeSetsStaffStatusType `xml:"EmploymentStatus,omitempty" json:"EmploymentStatus"`
      StaffSubjectList StaffSubjectListType `xml:"StaffSubjectList,omitempty" json:"StaffSubjectList"`
      StaffActivity StaffActivityExtensionType `xml:"StaffActivity,omitempty" json:"StaffActivity"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      CasualReliefTeacher AUCodeSetsYesOrNoCategoryType `xml:"CasualReliefTeacher,omitempty" json:"CasualReliefTeacher"`
      Homegroup string `xml:"Homegroup,omitempty" json:"Homegroup"`
      House string `xml:"House,omitempty" json:"House"`
      CalendarSummaryList CalendarSummaryListType `xml:"CalendarSummaryList,omitempty" json:"CalendarSummaryList"`
      PreviousSchoolName string `xml:"PreviousSchoolName,omitempty" json:"PreviousSchoolName"`
      AvailableForTimetable AUCodeSetsYesOrNoCategoryType `xml:"AvailableForTimetable,omitempty" json:"AvailableForTimetable"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    