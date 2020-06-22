package sifxml


    type PersonPicture struct {
        RefId RefIdType `xml:"RefId,attr"`
      ParentObjectRefId PersonPicture_ParentObjectRefId `xml:"ParentObjectRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      PictureSource PersonPicture_PictureSource `xml:"PictureSource"`
      OKToPublish string `xml:"OKToPublish"`
      PublishingPermissionList PublishingPermissionListType `xml:"PublishingPermissionList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type PersonPicture_ParentObjectRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type PersonPicture_PictureSource struct {
      Type string `xml:"Type,attr"`
      Value URIOrBinaryType `xml:",chardata"`
}
