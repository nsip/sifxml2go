package sifxml

import _ "github.com/creasty/defaults"


type StudentContactRelationships []StudentContactRelationship

    type StudentContactRelationship struct {
        StudentContactRelationshipRefId string `xml:"StudentContactRelationshipRefId,attr" json:"StudentContactRelationshipRefId"`
      StudentPersonalRefId RefIdType `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentContactPersonalRefId RefIdType `xml:"StudentContactPersonalRefId" json:"StudentContactPersonalRefId"`
      Relationship RelationshipType `xml:"Relationship" json:"Relationship"`
      ParentRelationshipStatus string `xml:"ParentRelationshipStatus,omitempty" json:"ParentRelationshipStatus"`
      HouseholdList HouseholdListType `xml:"HouseholdList,omitempty" json:"HouseholdList"`
      ContactFlags ContactFlagsType `xml:"ContactFlags" json:"ContactFlags"`
      MainlySpeaksEnglishAtHome AUCodeSetsYesOrNoCategoryType `xml:"MainlySpeaksEnglishAtHome,omitempty" json:"MainlySpeaksEnglishAtHome"`
      ContactSequence int `default:"-2147483648" xml:"ContactSequence" json:"ContactSequence"`
      ContactSequenceSource AUCodeSetsSourceCodeTypeType `xml:"ContactSequenceSource,omitempty" json:"ContactSequenceSource"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    