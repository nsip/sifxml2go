package sifxml


type RoomInfos []RoomInfo

    type RoomInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      RoomNumber *string `xml:"RoomNumber" json:"RoomNumber"`
      StaffList *StaffListType `xml:"StaffList,omitempty" json:"StaffList"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      Building *string `xml:"Building,omitempty" json:"Building"`
      HomeroomNumber *string `xml:"HomeroomNumber,omitempty" json:"HomeroomNumber"`
      Size *float64 `xml:"Size" json:"Size"`
      Capacity *int `xml:"Capacity" json:"Capacity"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      RoomType *string `xml:"RoomType,omitempty" json:"RoomType"`
      AvailableForTimetable *AUCodeSetsYesOrNoCategoryType `xml:"AvailableForTimetable,omitempty" json:"AvailableForTimetable"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    