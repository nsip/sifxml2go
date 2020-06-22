package sifxml


    type WellbeingCharacteristic struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      WellbeingCharacteristicClassification string `xml:"WellbeingCharacteristicClassification"`
      WellbeingCharacteristicStartDate string `xml:"WellbeingCharacteristicStartDate"`
      WellbeingCharacteristicEndDate string `xml:"WellbeingCharacteristicEndDate"`
      WellbeingCharacteristicReviewDate string `xml:"WellbeingCharacteristicReviewDate"`
      WellbeingCharacteristicNotes string `xml:"WellbeingCharacteristicNotes"`
      WellbeingCharacteristicCategory string `xml:"WellbeingCharacteristicCategory"`
      WellbeingCharacteristicSubCategory string `xml:"WellbeingCharacteristicSubCategory"`
      LocalCharacteristicCode LocalIdType `xml:"LocalCharacteristicCode"`
      CharacteristicSeverity string `xml:"CharacteristicSeverity"`
      SymptomList SymptomListType `xml:"SymptomList"`
      DailyManagement string `xml:"DailyManagement"`
      EmergencyManagement string `xml:"EmergencyManagement"`
      EmergencyResponsePlan string `xml:"EmergencyResponsePlan"`
      Trigger string `xml:"Trigger"`
      ConfidentialFlag string `xml:"ConfidentialFlag"`
      Alert string `xml:"Alert"`
      MedicationList MedicationListType `xml:"MedicationList"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    