package sifxml

import _ "github.com/creasty/defaults"


type PersonPictures []PersonPicture

    type PersonPicture struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      ParentObjectRefId PersonPicture_ParentObjectRefId `xml:"ParentObjectRefId" json:"ParentObjectRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      PictureSource PersonPicture_PictureSource `xml:"PictureSource" json:"PictureSource"`
      OKToPublish AUCodeSetsYesOrNoCategoryType `xml:"OKToPublish,omitempty" json:"OKToPublish"`
      PublishingPermissionList PublishingPermissionListType `xml:"PublishingPermissionList,omitempty" json:"PublishingPermissionList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type PersonPicture_ParentObjectRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value string `xml:",chardata" json:"value"`
}
type PersonPicture_PictureSource struct {
      Type AUCodeSetsPictureSourceType `xml:"Type,attr" json:"Type"`
      Value URIOrBinaryType `xml:",chardata" json:"value"`
}
