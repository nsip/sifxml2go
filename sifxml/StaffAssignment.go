package sifxml


    type StaffAssignment struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      Description string `xml:"Description"`
      PrimaryAssignment string `xml:"PrimaryAssignment"`
      JobStartDate string `xml:"JobStartDate"`
      JobEndDate string `xml:"JobEndDate"`
      JobFTE string `xml:"JobFTE"`
      JobFunction string `xml:"JobFunction"`
      EmploymentStatus string `xml:"EmploymentStatus"`
      StaffSubjectList StaffSubjectListType `xml:"StaffSubjectList"`
      StaffActivity StaffActivityExtensionType `xml:"StaffActivity"`
      YearLevels YearLevelsType `xml:"YearLevels"`
      CasualReliefTeacher string `xml:"CasualReliefTeacher"`
      Homegroup string `xml:"Homegroup"`
      House string `xml:"House"`
      CalendarSummaryList CalendarSummaryListType `xml:"CalendarSummaryList"`
      PreviousSchoolName string `xml:"PreviousSchoolName"`
      AvailableForTimetable string `xml:"AvailableForTimetable"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    