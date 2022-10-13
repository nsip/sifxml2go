package sifxml


type StudentDataTransferNotes []StudentDataTransferNote

    type StudentDataTransferNote struct {
  studentdatatransfernote `xml:"StudentDataTransferNote" json:"StudentDataTransferNote"`
}

type studentdatatransfernote struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Name *NameOfRecordType `xml:"Name" json:"Name"`
      Gender *AUCodeSetsSexCodeType `xml:"Gender,omitempty" json:"Gender,omitempty"`
      BirthDate *BirthDateType `xml:"BirthDate" json:"BirthDate"`
      PlaceOfBirth *String `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth,omitempty"`
      StateOfBirth *StateProvinceType `xml:"StateOfBirth,omitempty" json:"StateOfBirth,omitempty"`
      CountryOfBirth *CountryType `xml:"CountryOfBirth" json:"CountryOfBirth"`
      CountriesOfCitizenship *CountryListType `xml:"CountriesOfCitizenship,omitempty" json:"CountriesOfCitizenship,omitempty"`
      ArrivalSchool *ArrivalSchoolType `xml:"ArrivalSchool,omitempty" json:"ArrivalSchool,omitempty"`
      DepartureSchool *DepartureSchoolType `xml:"DepartureSchool,omitempty" json:"DepartureSchool,omitempty"`
      PreviousSchoolList *PreviousSchoolListType `xml:"PreviousSchoolList,omitempty" json:"PreviousSchoolList,omitempty"`
      NAPLANScoreList *NAPLANScoreWithYearsListType `xml:"NAPLANScoreList,omitempty" json:"NAPLANScoreList,omitempty"`
      NCCDList *NCCDListType `xml:"NCCDList,omitempty" json:"NCCDList,omitempty"`
      FollowupRequest *Bool `xml:"FollowupRequest,omitempty" json:"FollowupRequest,omitempty"`
      ChildSubjectToOrders *Bool `xml:"ChildSubjectToOrders,omitempty" json:"ChildSubjectToOrders,omitempty"`
      Attendance *Bool `xml:"Attendance,omitempty" json:"Attendance,omitempty"`
      NationalUniqueStudentIdentifier *String `xml:"NationalUniqueStudentIdentifier,omitempty" json:"NationalUniqueStudentIdentifier,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel" json:"YearLevel"`
      IndigenousStatus *AUCodeSetsIndigenousStatusType `xml:"IndigenousStatus,omitempty" json:"IndigenousStatus,omitempty"`
      LBOTE *AUCodeSetsYesOrNoCategoryType `xml:"LBOTE,omitempty" json:"LBOTE,omitempty"`
      VisaStatus *VisaSubClassType `xml:"VisaStatus,omitempty" json:"VisaStatus,omitempty"`
      OtherNames *OtherNamesType `xml:"OtherNames,omitempty" json:"OtherNames,omitempty"`
      EducationalAssessmentList *EducationalAssessmentListType `xml:"EducationalAssessmentList,omitempty" json:"EducationalAssessmentList,omitempty"`
      StudentGradeList *STDNGradeListType `xml:"StudentGradeList,omitempty" json:"StudentGradeList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

