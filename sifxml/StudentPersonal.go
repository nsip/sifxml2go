package sifxml


    type StudentPersonal struct {
        RefId RefIdType `xml:"RefId,attr"`
      AlertMessages AlertMessagesType `xml:"AlertMessages"`
      MedicalAlertMessages MedicalAlertMessagesType `xml:"MedicalAlertMessages"`
      LocalId LocalIdType `xml:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId"`
      ElectronicIdList ElectronicIdListType `xml:"ElectronicIdList"`
      OtherIdList OtherIdListType `xml:"OtherIdList"`
      PersonInfo PersonInfoType `xml:"PersonInfo"`
      ProjectedGraduationYear ProjectedGraduationYearType `xml:"ProjectedGraduationYear"`
      OnTimeGraduationYear OnTimeGraduationYearType `xml:"OnTimeGraduationYear"`
      GraduationDate GraduationDateType `xml:"GraduationDate"`
      MostRecent StudentMostRecentContainerType `xml:"MostRecent"`
      AcceptableUsePolicy string `xml:"AcceptableUsePolicy"`
      GiftedTalented string `xml:"GiftedTalented"`
      EconomicDisadvantage string `xml:"EconomicDisadvantage"`
      ESL string `xml:"ESL"`
      ESLDateAssessed string `xml:"ESLDateAssessed"`
      YoungCarersRole string `xml:"YoungCarersRole"`
      Disability string `xml:"Disability"`
      IntegrationAide string `xml:"IntegrationAide"`
      EducationSupport string `xml:"EducationSupport"`
      HomeSchooledStudent string `xml:"HomeSchooledStudent"`
      Sensitive string `xml:"Sensitive"`
      OfflineDelivery string `xml:"OfflineDelivery"`
      ESLSupport string `xml:"ESLSupport"`
      PrePrimaryEducation string `xml:"PrePrimaryEducation"`
      PrePrimaryEducationHours string `xml:"PrePrimaryEducationHours"`
      FirstAUSchoolEnrollment string `xml:"FirstAUSchoolEnrollment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    