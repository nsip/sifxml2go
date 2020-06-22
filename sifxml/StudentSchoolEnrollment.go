package sifxml


    type StudentSchoolEnrollment struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      MembershipType string `xml:"MembershipType"`
      TimeFrame string `xml:"TimeFrame"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      EntryDate string `xml:"EntryDate"`
      EntryType StudentEntryContainerType `xml:"EntryType"`
      YearLevel YearLevelType `xml:"YearLevel"`
      Homeroom StudentSchoolEnrollment_Homeroom `xml:"Homeroom"`
      Advisor StudentSchoolEnrollment_Advisor `xml:"Advisor"`
      Counselor StudentSchoolEnrollment_Counselor `xml:"Counselor"`
      Homegroup string `xml:"Homegroup"`
      ACARASchoolId LocalIdType `xml:"ACARASchoolId"`
      ClassCode string `xml:"ClassCode"`
      TestLevel YearLevelType `xml:"TestLevel"`
      ReportingSchool string `xml:"ReportingSchool"`
      House string `xml:"House"`
      IndividualLearningPlan string `xml:"IndividualLearningPlan"`
      Calendar StudentSchoolEnrollment_Calendar `xml:"Calendar"`
      ExitDate string `xml:"ExitDate"`
      ExitStatus StudentExitStatusContainerType `xml:"ExitStatus"`
      ExitType StudentExitContainerType `xml:"ExitType"`
      FTE string `xml:"FTE"`
      FTPTStatus string `xml:"FTPTStatus"`
      FFPOS string `xml:"FFPOS"`
      CatchmentStatus CatchmentStatusContainerType `xml:"CatchmentStatus"`
      RecordClosureReason string `xml:"RecordClosureReason"`
      PromotionInfo PromotionInfoType `xml:"PromotionInfo"`
      PreviousSchool LocalIdType `xml:"PreviousSchool"`
      PreviousSchoolName string `xml:"PreviousSchoolName"`
      DestinationSchool LocalIdType `xml:"DestinationSchool"`
      DestinationSchoolName string `xml:"DestinationSchoolName"`
      StudentSubjectChoiceList StudentSubjectChoiceListType `xml:"StudentSubjectChoiceList"`
      StartedAtSchoolDate string `xml:"StartedAtSchoolDate"`
      StudentGroupList StudentGroupListType `xml:"StudentGroupList"`
      PublishingPermissionList PublishingPermissionListType `xml:"PublishingPermissionList"`
      DisabilityLevelOfAdjustment string `xml:"DisabilityLevelOfAdjustment"`
      DisabilityCategory string `xml:"DisabilityCategory"`
      CensusAge string `xml:"CensusAge"`
      DistanceEducationStudent string `xml:"DistanceEducationStudent"`
      BoardingStatus string `xml:"BoardingStatus"`
      InternationalStudent string `xml:"InternationalStudent"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type StudentSchoolEnrollment_Homeroom struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type StudentSchoolEnrollment_Advisor struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type StudentSchoolEnrollment_Counselor struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type StudentSchoolEnrollment_Calendar struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
