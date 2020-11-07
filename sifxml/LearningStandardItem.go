package sifxml

import _ "github.com/creasty/defaults"


type LearningStandardItems []LearningStandardItem

    type LearningStandardItem struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      Resources LResourcesType `xml:"Resources,omitempty" json:"Resources"`
      StandardSettingBody StandardsSettingBodyType `xml:"StandardSettingBody,omitempty" json:"StandardSettingBody"`
      StandardHierarchyLevel StandardHierarchyLevelType `xml:"StandardHierarchyLevel" json:"StandardHierarchyLevel"`
      PredecessorItems LearningStandardsType `xml:"PredecessorItems,omitempty" json:"PredecessorItems"`
      StatementCodes StatementCodesType `xml:"StatementCodes,omitempty" json:"StatementCodes"`
      Statements StatementsType `xml:"Statements,omitempty" json:"Statements"`
      YearLevels YearLevelsType `xml:"YearLevels" json:"YearLevels"`
      ACStrandSubjectArea ACStrandSubjectAreaType `xml:"ACStrandSubjectArea,omitempty" json:"ACStrandSubjectArea"`
      StandardIdentifier StandardIdentifierType `xml:"StandardIdentifier,omitempty" json:"StandardIdentifier"`
      LearningStandardDocumentRefId string `xml:"LearningStandardDocumentRefId" json:"LearningStandardDocumentRefId"`
      RelatedLearningStandardItems RelatedLearningStandardItemRefIdListType `xml:"RelatedLearningStandardItems,omitempty" json:"RelatedLearningStandardItems"`
      Level4 string `xml:"Level4,omitempty" json:"Level4"`
      Level5 string `xml:"Level5,omitempty" json:"Level5"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    