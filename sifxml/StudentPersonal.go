package sifxml



    type StudentPersonals struct {
      studentpersonals `json:"StudentPersonals"`
    }

    type studentpersonals struct {
      StudentPersonal []studentpersonal `json:"StudentPersonal"`
    }

    type StudentPersonal struct {
      studentpersonal `xml:"StudentPersonal" json:"StudentPersonal"`
     }

     type studentpersonal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AlertMessages *AlertMessagesType `xml:"AlertMessages,omitempty" json:"AlertMessages,omitempty"`
      MedicalAlertMessages *MedicalAlertMessagesType `xml:"MedicalAlertMessages,omitempty" json:"MedicalAlertMessages,omitempty"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId,omitempty"`
      NationalUniqueStudentIdentifier *String `xml:"NationalUniqueStudentIdentifier,omitempty" json:"NationalUniqueStudentIdentifier,omitempty"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList,omitempty"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList,omitempty"`
      PersonInfo *PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      ProjectedGraduationYear *ProjectedGraduationYearType `xml:"ProjectedGraduationYear,omitempty" json:"ProjectedGraduationYear,omitempty"`
      OnTimeGraduationYear *OnTimeGraduationYearType `xml:"OnTimeGraduationYear,omitempty" json:"OnTimeGraduationYear,omitempty"`
      GraduationDate *GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate,omitempty"`
      MostRecent *StudentMostRecentContainerType `xml:"MostRecent,omitempty" json:"MostRecent,omitempty"`
      AcceptableUsePolicy *AUCodeSetsYesOrNoCategoryType `xml:"AcceptableUsePolicy,omitempty" json:"AcceptableUsePolicy,omitempty"`
      GiftedTalented *AUCodeSetsYesOrNoCategoryType `xml:"GiftedTalented,omitempty" json:"GiftedTalented,omitempty"`
      EconomicDisadvantage *AUCodeSetsYesOrNoCategoryType `xml:"EconomicDisadvantage,omitempty" json:"EconomicDisadvantage,omitempty"`
      ESL *AUCodeSetsYesOrNoCategoryType `xml:"ESL,omitempty" json:"ESL,omitempty"`
      ESLDateAssessed *String `xml:"ESLDateAssessed,omitempty" json:"ESLDateAssessed,omitempty"`
      YoungCarersRole *AUCodeSetsYesOrNoCategoryType `xml:"YoungCarersRole,omitempty" json:"YoungCarersRole,omitempty"`
      Disability *AUCodeSetsYesOrNoCategoryType `xml:"Disability,omitempty" json:"Disability,omitempty"`
      CategoryOfDisability *AUCodeSetsNCCDDisabilityType `xml:"CategoryOfDisability,omitempty" json:"CategoryOfDisability,omitempty"`
      IntegrationAide *AUCodeSetsYesOrNoCategoryType `xml:"IntegrationAide,omitempty" json:"IntegrationAide,omitempty"`
      EducationSupport *AUCodeSetsYesOrNoCategoryType `xml:"EducationSupport,omitempty" json:"EducationSupport,omitempty"`
      HomeSchooledStudent *AUCodeSetsYesOrNoCategoryType `xml:"HomeSchooledStudent,omitempty" json:"HomeSchooledStudent,omitempty"`
      IndependentStudent *AUCodeSetsYesOrNoCategoryType `xml:"IndependentStudent,omitempty" json:"IndependentStudent,omitempty"`
      Sensitive *AUCodeSetsYesOrNoCategoryType `xml:"Sensitive,omitempty" json:"Sensitive,omitempty"`
      OfflineDelivery *AUCodeSetsYesOrNoCategoryType `xml:"OfflineDelivery,omitempty" json:"OfflineDelivery,omitempty"`
      ESLSupport *AUCodeSetsYesOrNoCategoryType `xml:"ESLSupport,omitempty" json:"ESLSupport,omitempty"`
      PrePrimaryEducation *String `xml:"PrePrimaryEducation,omitempty" json:"PrePrimaryEducation,omitempty"`
      PrePrimaryEducationHours *AUCodeSetsPrePrimaryHoursType `xml:"PrePrimaryEducationHours,omitempty" json:"PrePrimaryEducationHours,omitempty"`
      FirstAUSchoolEnrollment *String `xml:"FirstAUSchoolEnrollment,omitempty" json:"FirstAUSchoolEnrollment,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

