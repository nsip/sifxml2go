package sifxml


type StudentContactRelationships []StudentContactRelationship

    type StudentContactRelationship struct {
  studentcontactrelationship `xml:"StudentContactRelationship"`
}

type studentcontactrelationship struct {
        StudentContactRelationshipRefId *String `xml:"StudentContactRelationshipRefId,attr" json:"StudentContactRelationshipRefId"`
      StudentPersonalRefId *RefIdType `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentContactPersonalRefId *RefIdType `xml:"StudentContactPersonalRefId" json:"StudentContactPersonalRefId"`
      Relationship *RelationshipType `xml:"Relationship" json:"Relationship"`
      ParentRelationshipStatus *String `xml:"ParentRelationshipStatus,omitempty" json:"ParentRelationshipStatus,omitempty"`
      HouseholdList *HouseholdListType `xml:"HouseholdList,omitempty" json:"HouseholdList,omitempty"`
      ContactFlags *ContactFlagsType `xml:"ContactFlags,omitempty" json:"ContactFlags,omitempty"`
      MainlySpeaksEnglishAtHome *AUCodeSetsYesOrNoCategoryType `xml:"MainlySpeaksEnglishAtHome,omitempty" json:"MainlySpeaksEnglishAtHome,omitempty"`
      ContactSequence *Int `xml:"ContactSequence,omitempty" json:"ContactSequence,omitempty"`
      ContactSequenceSource *AUCodeSetsSourceCodeTypeType `xml:"ContactSequenceSource,omitempty" json:"ContactSequenceSource,omitempty"`
      ContactMethod *AUCodeSetsContactMethodType `xml:"ContactMethod,omitempty" json:"ContactMethod,omitempty"`
      FeePercentage *StudentContactFeePercentageType `xml:"FeePercentage,omitempty" json:"FeePercentage,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

