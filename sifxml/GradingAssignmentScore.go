package sifxml

import _ "github.com/creasty/defaults"


type GradingAssignmentScores []GradingAssignmentScore

    type GradingAssignmentScore struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentPersonalLocalId LocalIdType `xml:"StudentPersonalLocalId" json:"StudentPersonalLocalId"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      GradingAssignmentRefId string `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      DateGraded string `xml:"DateGraded,omitempty" json:"DateGraded"`
      ExpectedScore bool `xml:"ExpectedScore" json:"ExpectedScore"`
      ScorePoints int `xml:"ScorePoints" json:"ScorePoints"`
      ScorePercent float64 `xml:"ScorePercent" json:"ScorePercent"`
      ScoreLetter string `xml:"ScoreLetter" json:"ScoreLetter"`
      ScoreDescription string `xml:"ScoreDescription,omitempty" json:"ScoreDescription"`
      SubscoreList NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      TeacherJudgement string `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement"`
      MarkInfoRefId string `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      AssignmentScoreIteration string `xml:"AssignmentScoreIteration,omitempty" json:"AssignmentScoreIteration"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    