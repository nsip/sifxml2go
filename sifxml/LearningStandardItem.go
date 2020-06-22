package sifxml


    type LearningStandardItem struct {
        RefId RefIdType `xml:"RefId,attr"`
      Resources LResourcesType `xml:"Resources"`
      StandardSettingBody StandardsSettingBodyType `xml:"StandardSettingBody"`
      StandardHierarchyLevel StandardHierarchyLevelType `xml:"StandardHierarchyLevel"`
      PredecessorItems LearningStandardsType `xml:"PredecessorItems"`
      StatementCodes StatementCodesType `xml:"StatementCodes"`
      Statements StatementsType `xml:"Statements"`
      YearLevels YearLevelsType `xml:"YearLevels"`
      ACStrandSubjectArea ACStrandSubjectAreaType `xml:"ACStrandSubjectArea"`
      StandardIdentifier StandardIdentifierType `xml:"StandardIdentifier"`
      LearningStandardDocumentRefId string `xml:"LearningStandardDocumentRefId"`
      RelatedLearningStandardItems RelatedLearningStandardItemRefIdListType `xml:"RelatedLearningStandardItems"`
      Level4 string `xml:"Level4"`
      Level5 string `xml:"Level5"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    