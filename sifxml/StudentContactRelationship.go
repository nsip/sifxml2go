package sifxml


    type StudentContactRelationship struct {
        StudentContactRelationshipRefId string `xml:"StudentContactRelationshipRefId,attr"`
      StudentPersonalRefId RefIdType `xml:"StudentPersonalRefId"`
      StudentContactPersonalRefId RefIdType `xml:"StudentContactPersonalRefId"`
      Relationship RelationshipType `xml:"Relationship"`
      ParentRelationshipStatus string `xml:"ParentRelationshipStatus"`
      HouseholdList HouseholdListType `xml:"HouseholdList"`
      ContactFlags ContactFlagsType `xml:"ContactFlags"`
      MainlySpeaksEnglishAtHome string `xml:"MainlySpeaksEnglishAtHome"`
      ContactSequence string `xml:"ContactSequence"`
      ContactSequenceSource string `xml:"ContactSequenceSource"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    