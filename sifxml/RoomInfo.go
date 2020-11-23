package sifxml


type RoomInfos []RoomInfo

    type RoomInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      RoomNumber *String `xml:"RoomNumber" json:"RoomNumber"`
      StaffList *StaffListType `xml:"StaffList,omitempty" json:"StaffList"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      Building *String `xml:"Building,omitempty" json:"Building"`
      HomeroomNumber *String `xml:"HomeroomNumber,omitempty" json:"HomeroomNumber"`
      Size *Float `xml:"Size,omitempty" json:"Size"`
      Capacity *Int `xml:"Capacity,omitempty" json:"Capacity"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      RoomType *String `xml:"RoomType,omitempty" json:"RoomType"`
      AvailableForTimetable *AUCodeSetsYesOrNoCategoryType `xml:"AvailableForTimetable,omitempty" json:"AvailableForTimetable"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    