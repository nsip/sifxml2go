package sifxml


    type CollectionRound struct {
        RefId RefIdType `xml:"RefId,attr"`
      AGCollection string `xml:"AGCollection"`
      CollectionYear SchoolYearType `xml:"CollectionYear"`
      AGRoundList AGRoundListType `xml:"AGRoundList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    