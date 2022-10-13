package sifxml


type LearningStandardDocuments []LearningStandardDocument

    type LearningStandardDocument struct {
  learningstandarddocument `xml:"LearningStandardDocument" json:"LearningStandardDocument"`
}

type learningstandarddocument struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *String `xml:"Title" json:"Title"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      RichDescription *AbstractContentElementType `xml:"RichDescription,omitempty" json:"RichDescription,omitempty"`
      Source *String `xml:"Source" json:"Source"`
      Organizations *OrganizationsType `xml:"Organizations" json:"Organizations"`
      Authors *AuthorsType `xml:"Authors,omitempty" json:"Authors,omitempty"`
      OrganizationContactPoint *String `xml:"OrganizationContactPoint,omitempty" json:"OrganizationContactPoint,omitempty"`
      SubjectAreas *ACStrandAreaListType `xml:"SubjectAreas" json:"SubjectAreas"`
      DocumentStatus *String `xml:"DocumentStatus" json:"DocumentStatus"`
      DocumentDate *String `xml:"DocumentDate,omitempty" json:"DocumentDate,omitempty"`
      LocalAdoptionDate *String `xml:"LocalAdoptionDate,omitempty" json:"LocalAdoptionDate,omitempty"`
      LocalArchiveDate *String `xml:"LocalArchiveDate,omitempty" json:"LocalArchiveDate,omitempty"`
      EndOfLifeDate *String `xml:"EndOfLifeDate,omitempty" json:"EndOfLifeDate,omitempty"`
      Copyright *CopyRightContainerType `xml:"Copyright,omitempty" json:"Copyright,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      RepositoryDate *String `xml:"RepositoryDate,omitempty" json:"RepositoryDate,omitempty"`
      LearningStandardItemRefId *String `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      RelatedLearningStandards *LearningStandardsDocumentType `xml:"RelatedLearningStandards,omitempty" json:"RelatedLearningStandards,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

