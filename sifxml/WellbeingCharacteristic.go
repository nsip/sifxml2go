package sifxml


type WellbeingCharacteristics []WellbeingCharacteristic

    type WellbeingCharacteristic struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      WellbeingCharacteristicClassification *AUCodeSetsWellbeingCharacteristicClassificationType `xml:"WellbeingCharacteristicClassification" json:"WellbeingCharacteristicClassification"`
      WellbeingCharacteristicStartDate *string `xml:"WellbeingCharacteristicStartDate,omitempty" json:"WellbeingCharacteristicStartDate"`
      WellbeingCharacteristicEndDate *string `xml:"WellbeingCharacteristicEndDate,omitempty" json:"WellbeingCharacteristicEndDate"`
      WellbeingCharacteristicReviewDate *string `xml:"WellbeingCharacteristicReviewDate,omitempty" json:"WellbeingCharacteristicReviewDate"`
      WellbeingCharacteristicNotes *string `xml:"WellbeingCharacteristicNotes,omitempty" json:"WellbeingCharacteristicNotes"`
      WellbeingCharacteristicCategory *string `xml:"WellbeingCharacteristicCategory,omitempty" json:"WellbeingCharacteristicCategory"`
      WellbeingCharacteristicSubCategory *string `xml:"WellbeingCharacteristicSubCategory,omitempty" json:"WellbeingCharacteristicSubCategory"`
      LocalCharacteristicCode *LocalIdType `xml:"LocalCharacteristicCode,omitempty" json:"LocalCharacteristicCode"`
      CharacteristicSeverity *string `xml:"CharacteristicSeverity,omitempty" json:"CharacteristicSeverity"`
      SymptomList *SymptomListType `xml:"SymptomList,omitempty" json:"SymptomList"`
      DailyManagement *string `xml:"DailyManagement,omitempty" json:"DailyManagement"`
      EmergencyManagement *string `xml:"EmergencyManagement,omitempty" json:"EmergencyManagement"`
      EmergencyResponsePlan *string `xml:"EmergencyResponsePlan,omitempty" json:"EmergencyResponsePlan"`
      Trigger *string `xml:"Trigger,omitempty" json:"Trigger"`
      ConfidentialFlag *AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag"`
      Alert *AUCodeSetsYesOrNoCategoryType `xml:"Alert,omitempty" json:"Alert"`
      MedicationList *MedicationListType `xml:"MedicationList,omitempty" json:"MedicationList"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    