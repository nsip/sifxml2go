package sifxml


type StudentContactRelationships []StudentContactRelationship

    type StudentContactRelationship struct {
        StudentContactRelationshipRefId *String `xml:"StudentContactRelationshipRefId,attr" json:"StudentContactRelationshipRefId"`
      StudentPersonalRefId *RefIdType `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentContactPersonalRefId *RefIdType `xml:"StudentContactPersonalRefId" json:"StudentContactPersonalRefId"`
      Relationship *RelationshipType `xml:"Relationship" json:"Relationship"`
      ParentRelationshipStatus *String `xml:"ParentRelationshipStatus,omitempty" json:"ParentRelationshipStatus"`
      HouseholdList *HouseholdListType `xml:"HouseholdList,omitempty" json:"HouseholdList"`
      ContactFlags *ContactFlagsType `xml:"ContactFlags,omitempty" json:"ContactFlags"`
      MainlySpeaksEnglishAtHome *AUCodeSetsYesOrNoCategoryType `xml:"MainlySpeaksEnglishAtHome,omitempty" json:"MainlySpeaksEnglishAtHome"`
      ContactSequence *Int `xml:"ContactSequence,omitempty" json:"ContactSequence"`
      ContactSequenceSource *AUCodeSetsSourceCodeTypeType `xml:"ContactSequenceSource,omitempty" json:"ContactSequenceSource"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    