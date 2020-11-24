package sifxml


type CollectionRounds []CollectionRound

    type CollectionRound struct {
  collectionround `xml:"CollectionRound" json:"CollectionRound"`
}

type collectionround struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AGCollection *AUCodeSetsAGCollectionType `xml:"AGCollection,omitempty" json:"AGCollection,omitempty"`
      CollectionYear *SchoolYearType `xml:"CollectionYear" json:"CollectionYear"`
      AGRoundList *AGRoundListType `xml:"AGRoundList" json:"AGRoundList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    