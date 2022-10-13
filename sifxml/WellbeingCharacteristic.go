package sifxml


type WellbeingCharacteristics []WellbeingCharacteristic

    type WellbeingCharacteristic struct {
  wellbeingcharacteristic `xml:"WellbeingCharacteristic" json:"WellbeingCharacteristic"`
}

type wellbeingcharacteristic struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      WellbeingCharacteristicClassification *AUCodeSetsWellbeingCharacteristicClassificationType `xml:"WellbeingCharacteristicClassification,omitempty" json:"WellbeingCharacteristicClassification,omitempty"`
      WellbeingCharacteristicStartDate *String `xml:"WellbeingCharacteristicStartDate,omitempty" json:"WellbeingCharacteristicStartDate,omitempty"`
      WellbeingCharacteristicEndDate *String `xml:"WellbeingCharacteristicEndDate,omitempty" json:"WellbeingCharacteristicEndDate,omitempty"`
      WellbeingCharacteristicReviewDate *String `xml:"WellbeingCharacteristicReviewDate,omitempty" json:"WellbeingCharacteristicReviewDate,omitempty"`
      WellbeingCharacteristicNotes *String `xml:"WellbeingCharacteristicNotes,omitempty" json:"WellbeingCharacteristicNotes,omitempty"`
      WellbeingCharacteristicCategory *String `xml:"WellbeingCharacteristicCategory,omitempty" json:"WellbeingCharacteristicCategory,omitempty"`
      WellbeingCharacteristicSubCategory *String `xml:"WellbeingCharacteristicSubCategory,omitempty" json:"WellbeingCharacteristicSubCategory,omitempty"`
      LocalCharacteristicCode *LocalIdType `xml:"LocalCharacteristicCode,omitempty" json:"LocalCharacteristicCode,omitempty"`
      CharacteristicSeverity *String `xml:"CharacteristicSeverity,omitempty" json:"CharacteristicSeverity,omitempty"`
      SymptomList *SymptomListType `xml:"SymptomList,omitempty" json:"SymptomList,omitempty"`
      DailyManagement *String `xml:"DailyManagement,omitempty" json:"DailyManagement,omitempty"`
      EmergencyManagement *String `xml:"EmergencyManagement,omitempty" json:"EmergencyManagement,omitempty"`
      PreferredHospital *String `xml:"PreferredHospital,omitempty" json:"PreferredHospital,omitempty"`
      EmergencyResponsePlan *String `xml:"EmergencyResponsePlan,omitempty" json:"EmergencyResponsePlan,omitempty"`
      Trigger *String `xml:"Trigger,omitempty" json:"Trigger,omitempty"`
      ConfidentialFlag *AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag,omitempty"`
      Alert *AUCodeSetsYesOrNoCategoryType `xml:"Alert,omitempty" json:"Alert,omitempty"`
      MedicationList *MedicationListType `xml:"MedicationList,omitempty" json:"MedicationList,omitempty"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

