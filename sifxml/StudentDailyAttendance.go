package sifxml



    type StudentDailyAttendances struct {
      studentdailyattendances `json:"StudentDailyAttendances"`
    }

    type studentdailyattendances struct {
      StudentDailyAttendance []studentdailyattendance `json:"StudentDailyAttendance"`
    }

    type StudentDailyAttendance struct {
      studentdailyattendance `xml:"StudentDailyAttendance" json:"StudentDailyAttendance"`
     }

     type studentdailyattendance struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      DayValue *AUCodeSetsDayValueCodeType `xml:"DayValue,omitempty" json:"DayValue,omitempty"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn,omitempty"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut,omitempty"`
      AbsenceValue *Float `xml:"AbsenceValue,omitempty" json:"AbsenceValue,omitempty"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

