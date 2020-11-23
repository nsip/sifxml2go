package sifxml


type LearningStandardDocuments []LearningStandardDocument

    type LearningStandardDocument struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *String `xml:"Title" json:"Title"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      RichDescription *AbstractContentElementType `xml:"RichDescription,omitempty" json:"RichDescription"`
      Source *String `xml:"Source" json:"Source"`
      Organizations *OrganizationsType `xml:"Organizations" json:"Organizations"`
      Authors *AuthorsType `xml:"Authors,omitempty" json:"Authors"`
      OrganizationContactPoint *String `xml:"OrganizationContactPoint,omitempty" json:"OrganizationContactPoint"`
      SubjectAreas *ACStrandAreaListType `xml:"SubjectAreas" json:"SubjectAreas"`
      DocumentStatus *String `xml:"DocumentStatus" json:"DocumentStatus"`
      DocumentDate *String `xml:"DocumentDate,omitempty" json:"DocumentDate"`
      LocalAdoptionDate *String `xml:"LocalAdoptionDate,omitempty" json:"LocalAdoptionDate"`
      LocalArchiveDate *String `xml:"LocalArchiveDate,omitempty" json:"LocalArchiveDate"`
      EndOfLifeDate *String `xml:"EndOfLifeDate,omitempty" json:"EndOfLifeDate"`
      Copyright *CopyRightContainerType `xml:"Copyright,omitempty" json:"Copyright"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      RepositoryDate *String `xml:"RepositoryDate,omitempty" json:"RepositoryDate"`
      LearningStandardItemRefId *String `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      RelatedLearningStandards *LearningStandardsDocumentType `xml:"RelatedLearningStandards,omitempty" json:"RelatedLearningStandards"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    