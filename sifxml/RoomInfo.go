package sifxml



    type RoomInfos struct {
      roominfos `json:"RoomInfos"`
    }

    type roominfos struct {
      RoomInfo []roominfo `json:"RoomInfo"`
    }

    type RoomInfo struct {
      roominfo `xml:"RoomInfo" json:"RoomInfo"`
     }

     type roominfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      RoomNumber *String `xml:"RoomNumber" json:"RoomNumber"`
      StaffList *StaffListType `xml:"StaffList,omitempty" json:"StaffList,omitempty"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      Building *String `xml:"Building,omitempty" json:"Building,omitempty"`
      HomeroomNumber *String `xml:"HomeroomNumber,omitempty" json:"HomeroomNumber,omitempty"`
      Size *Float `xml:"Size,omitempty" json:"Size,omitempty"`
      Capacity *Int `xml:"Capacity,omitempty" json:"Capacity,omitempty"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
      RoomType *String `xml:"RoomType,omitempty" json:"RoomType,omitempty"`
      AvailableForTimetable *AUCodeSetsYesOrNoCategoryType `xml:"AvailableForTimetable,omitempty" json:"AvailableForTimetable,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

