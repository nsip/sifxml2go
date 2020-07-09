package sifxml


type CollectionRounds []CollectionRound

    type CollectionRound struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AGCollection *string `xml:"AGCollection,omitempty" json:"AGCollection"`
      CollectionYear *SchoolYearType `xml:"CollectionYear,omitempty" json:"CollectionYear"`
      AGRoundList *AGRoundListType `xml:"AGRoundList,omitempty" json:"AGRoundList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    