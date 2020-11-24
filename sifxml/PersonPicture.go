package sifxml


type PersonPictures []PersonPicture

    type PersonPicture struct {
  personpicture `xml:"PersonPicture" json:"PersonPicture"`
}

type personpicture struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ParentObjectRefId *PersonPicture_ParentObjectRefId `xml:"ParentObjectRefId" json:"ParentObjectRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      PictureSource *PersonPicture_PictureSource `xml:"PictureSource" json:"PictureSource"`
      OKToPublish *AUCodeSetsYesOrNoCategoryType `xml:"OKToPublish,omitempty" json:"OKToPublish,omitempty"`
      PublishingPermissionList *PublishingPermissionListType `xml:"PublishingPermissionList,omitempty" json:"PublishingPermissionList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type PersonPicture_ParentObjectRefId struct {
  personpicture_parentobjectrefid `xml:"PersonPicture_ParentObjectRefId" json:"PersonPicture_ParentObjectRefId"`
}

type personpicture_parentobjectrefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type PersonPicture_PictureSource struct {
  personpicture_picturesource `xml:"PersonPicture_PictureSource" json:"PersonPicture_PictureSource"`
}

type personpicture_picturesource struct {
      Type *AUCodeSetsPictureSourceType `xml:"Type,attr" json:"Type"`
      Value *URIOrBinaryType `xml:",chardata" json:"value"`
}
