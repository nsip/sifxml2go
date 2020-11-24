package sifxml


type LearningResources []LearningResource

    type LearningResource struct {
  learningresource `xml:"LearningResource" json:"LearningResource"`
}

type learningresource struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Name *String `xml:"Name" json:"Name"`
      Author *String `xml:"Author,omitempty" json:"Author,omitempty"`
      Contacts *ContactsType `xml:"Contacts,omitempty" json:"Contacts,omitempty"`
      Location *LearningResource_Location
      Status *String `xml:"Status,omitempty" json:"Status,omitempty"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      SubjectAreas *ACStrandAreaListType `xml:"SubjectAreas,omitempty" json:"SubjectAreas,omitempty"`
      MediaTypes *MediaTypesType `xml:"MediaTypes,omitempty" json:"MediaTypes,omitempty"`
      UseAgreement *String `xml:"UseAgreement,omitempty" json:"UseAgreement,omitempty"`
      AgreementDate *String `xml:"AgreementDate,omitempty" json:"AgreementDate,omitempty"`
      Approvals *ApprovalsType `xml:"Approvals,omitempty" json:"Approvals,omitempty"`
      Evaluations *EvaluationsType `xml:"Evaluations,omitempty" json:"Evaluations,omitempty"`
      Components *ComponentsType `xml:"Components" json:"Components"`
      LearningStandards *LearningStandardsType `xml:"LearningStandards,omitempty" json:"LearningStandards,omitempty"`
      LearningResourcePackageRefId *String `xml:"LearningResourcePackageRefId,omitempty" json:"LearningResourcePackageRefId,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type LearningResource_Location struct {
  learningresource_location `xml:"LearningResource_Location" json:"LearningResource_Location"`
}

type learningresource_location struct {
      ReferenceType *String `xml:"ReferenceType,attr" json:"ReferenceType"`
      Value *String `xml:",chardata" json:"value"`
}
