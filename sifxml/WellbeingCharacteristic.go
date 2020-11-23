package sifxml


type WellbeingCharacteristics []WellbeingCharacteristic

    type WellbeingCharacteristic struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      WellbeingCharacteristicClassification *AUCodeSetsWellbeingCharacteristicClassificationType `xml:"WellbeingCharacteristicClassification,omitempty" json:"WellbeingCharacteristicClassification"`
      WellbeingCharacteristicStartDate *String `xml:"WellbeingCharacteristicStartDate,omitempty" json:"WellbeingCharacteristicStartDate"`
      WellbeingCharacteristicEndDate *String `xml:"WellbeingCharacteristicEndDate,omitempty" json:"WellbeingCharacteristicEndDate"`
      WellbeingCharacteristicReviewDate *String `xml:"WellbeingCharacteristicReviewDate,omitempty" json:"WellbeingCharacteristicReviewDate"`
      WellbeingCharacteristicNotes *String `xml:"WellbeingCharacteristicNotes,omitempty" json:"WellbeingCharacteristicNotes"`
      WellbeingCharacteristicCategory *String `xml:"WellbeingCharacteristicCategory,omitempty" json:"WellbeingCharacteristicCategory"`
      WellbeingCharacteristicSubCategory *String `xml:"WellbeingCharacteristicSubCategory,omitempty" json:"WellbeingCharacteristicSubCategory"`
      LocalCharacteristicCode *LocalIdType `xml:"LocalCharacteristicCode,omitempty" json:"LocalCharacteristicCode"`
      CharacteristicSeverity *String `xml:"CharacteristicSeverity,omitempty" json:"CharacteristicSeverity"`
      SymptomList *SymptomListType `xml:"SymptomList,omitempty" json:"SymptomList"`
      DailyManagement *String `xml:"DailyManagement,omitempty" json:"DailyManagement"`
      EmergencyManagement *String `xml:"EmergencyManagement,omitempty" json:"EmergencyManagement"`
      EmergencyResponsePlan *String `xml:"EmergencyResponsePlan,omitempty" json:"EmergencyResponsePlan"`
      Trigger *String `xml:"Trigger,omitempty" json:"Trigger"`
      ConfidentialFlag *AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag"`
      Alert *AUCodeSetsYesOrNoCategoryType `xml:"Alert,omitempty" json:"Alert"`
      MedicationList *MedicationListType `xml:"MedicationList,omitempty" json:"MedicationList"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    