package sifxml



    type LearningStandardItems struct {
      learningstandarditems `json:"LearningStandardItems"`
    }

    type learningstandarditems struct {
      LearningStandardItem []learningstandarditem `json:"LearningStandardItem"`
    }

    type LearningStandardItem struct {
      learningstandarditem `xml:"LearningStandardItem" json:"LearningStandardItem"`
     }

     type learningstandarditem struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Resources *LResourcesType `xml:"Resources,omitempty" json:"Resources,omitempty"`
      StandardSettingBody *StandardsSettingBodyType `xml:"StandardSettingBody,omitempty" json:"StandardSettingBody,omitempty"`
      StandardHierarchyLevel *StandardHierarchyLevelType `xml:"StandardHierarchyLevel" json:"StandardHierarchyLevel"`
      PredecessorItems *LearningStandardsType `xml:"PredecessorItems,omitempty" json:"PredecessorItems,omitempty"`
      StatementCodes *StatementCodesType `xml:"StatementCodes,omitempty" json:"StatementCodes,omitempty"`
      Statements *StatementsType `xml:"Statements,omitempty" json:"Statements,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels" json:"YearLevels"`
      ACStrandSubjectArea *ACStrandSubjectAreaType `xml:"ACStrandSubjectArea,omitempty" json:"ACStrandSubjectArea,omitempty"`
      StandardIdentifier *StandardIdentifierType `xml:"StandardIdentifier,omitempty" json:"StandardIdentifier,omitempty"`
      LearningStandardDocumentRefId *String `xml:"LearningStandardDocumentRefId" json:"LearningStandardDocumentRefId"`
      RelatedLearningStandardItems *RelatedLearningStandardItemRefIdListType `xml:"RelatedLearningStandardItems,omitempty" json:"RelatedLearningStandardItems,omitempty"`
      Level4 *String `xml:"Level4,omitempty" json:"Level4,omitempty"`
      Level5 *String `xml:"Level5,omitempty" json:"Level5,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

