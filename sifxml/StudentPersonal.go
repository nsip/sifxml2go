package sifxml

import _ "github.com/creasty/defaults"


type StudentPersonals []StudentPersonal

    type StudentPersonal struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      AlertMessages AlertMessagesType `xml:"AlertMessages,omitempty" json:"AlertMessages"`
      MedicalAlertMessages MedicalAlertMessagesType `xml:"MedicalAlertMessages,omitempty" json:"MedicalAlertMessages"`
      LocalId LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ElectronicIdList ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      OtherIdList OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      PersonInfo PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      ProjectedGraduationYear ProjectedGraduationYearType `xml:"ProjectedGraduationYear,omitempty" json:"ProjectedGraduationYear"`
      OnTimeGraduationYear OnTimeGraduationYearType `xml:"OnTimeGraduationYear,omitempty" json:"OnTimeGraduationYear"`
      GraduationDate GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate"`
      MostRecent StudentMostRecentContainerType `xml:"MostRecent,omitempty" json:"MostRecent"`
      AcceptableUsePolicy AUCodeSetsYesOrNoCategoryType `xml:"AcceptableUsePolicy,omitempty" json:"AcceptableUsePolicy"`
      GiftedTalented AUCodeSetsYesOrNoCategoryType `xml:"GiftedTalented,omitempty" json:"GiftedTalented"`
      EconomicDisadvantage AUCodeSetsYesOrNoCategoryType `xml:"EconomicDisadvantage,omitempty" json:"EconomicDisadvantage"`
      ESL AUCodeSetsYesOrNoCategoryType `xml:"ESL,omitempty" json:"ESL"`
      ESLDateAssessed string `xml:"ESLDateAssessed,omitempty" json:"ESLDateAssessed"`
      YoungCarersRole AUCodeSetsYesOrNoCategoryType `xml:"YoungCarersRole,omitempty" json:"YoungCarersRole"`
      Disability AUCodeSetsYesOrNoCategoryType `xml:"Disability,omitempty" json:"Disability"`
      IntegrationAide AUCodeSetsYesOrNoCategoryType `xml:"IntegrationAide,omitempty" json:"IntegrationAide"`
      EducationSupport AUCodeSetsYesOrNoCategoryType `xml:"EducationSupport,omitempty" json:"EducationSupport"`
      HomeSchooledStudent AUCodeSetsYesOrNoCategoryType `xml:"HomeSchooledStudent,omitempty" json:"HomeSchooledStudent"`
      Sensitive AUCodeSetsYesOrNoCategoryType `xml:"Sensitive,omitempty" json:"Sensitive"`
      OfflineDelivery AUCodeSetsYesOrNoCategoryType `xml:"OfflineDelivery,omitempty" json:"OfflineDelivery"`
      ESLSupport AUCodeSetsYesOrNoCategoryType `xml:"ESLSupport,omitempty" json:"ESLSupport"`
      PrePrimaryEducation string `xml:"PrePrimaryEducation,omitempty" json:"PrePrimaryEducation"`
      PrePrimaryEducationHours AUCodeSetsPrePrimaryHoursType `xml:"PrePrimaryEducationHours,omitempty" json:"PrePrimaryEducationHours"`
      FirstAUSchoolEnrollment string `xml:"FirstAUSchoolEnrollment,omitempty" json:"FirstAUSchoolEnrollment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    