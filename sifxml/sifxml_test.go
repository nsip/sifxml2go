package sifxml

import (
        "encoding/xml"
        "github.com/clbanning/mxj"
        "reflect"
        "regexp"
        "strings"
        "testing"
)

var emptytag1 = regexp.MustCompile(`(?s:\s*<[^>/]+></[^>]+>\s*)`)
var emptytag2 = regexp.MustCompile(`(?s:\s*<[^>/]+/>\s*)`)
var emptytag3 = regexp.MustCompile(`(?s:\s+[^>='" ]+=(''|""))`)

func stripEmptyTags(s []byte) []byte {
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag2.ReplaceAll(s, []byte(""))
        arr := strings.SplitAfter(string(s), ">")
        for i, _ := range arr {
          if strings.HasPrefix(arr[i], "<") {
            // only works because we don't have mixed tags and text in SIF
            arr[i] = emptytag3.ReplaceAllString(arr[i], "")
          }
        }
        s = []byte(strings.Join(arr, ""))
        return s
}

func errcheck(t *testing.T, err error) {
        if err != nil {
                t.Fatalf("Error %s", err)
        }
}
func Test_test_ActivityExample1(t *testing.T) {
        a := Activity{}
        err := xml.Unmarshal([]byte(test_ActivityExample1), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ActivityExample1)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ActivityExample1 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ActivityExample1))), string(stripEmptyTags(output)))
        }
}
func Test_test_ActivityExample2(t *testing.T) {
        a := Activity{}
        err := xml.Unmarshal([]byte(test_ActivityExample2), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ActivityExample2)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ActivityExample2 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ActivityExample2))), string(stripEmptyTags(output)))
        }
}
func Test_test_AggregateCharacteristicInfo(t *testing.T) {
        a := AggregateCharacteristicInfo{}
        err := xml.Unmarshal([]byte(test_AggregateCharacteristicInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_AggregateCharacteristicInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_AggregateCharacteristicInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_AggregateCharacteristicInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_AggregateStatisticFact(t *testing.T) {
        a := AggregateStatisticFact{}
        err := xml.Unmarshal([]byte(test_AggregateStatisticFact), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_AggregateStatisticFact)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_AggregateStatisticFact not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_AggregateStatisticFact))), string(stripEmptyTags(output)))
        }
}
func Test_test_AggregateStatisticInfo(t *testing.T) {
        a := AggregateStatisticInfo{}
        err := xml.Unmarshal([]byte(test_AggregateStatisticInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_AggregateStatisticInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_AggregateStatisticInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_AggregateStatisticInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_CalendarDate(t *testing.T) {
        a := CalendarDate{}
        err := xml.Unmarshal([]byte(test_CalendarDate), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_CalendarDate)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_CalendarDate not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_CalendarDate))), string(stripEmptyTags(output)))
        }
}
func Test_test_CalendarSummary(t *testing.T) {
        a := CalendarSummary{}
        err := xml.Unmarshal([]byte(test_CalendarSummary), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_CalendarSummary)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_CalendarSummary not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_CalendarSummary))), string(stripEmptyTags(output)))
        }
}
func Test_test_EquipmentInfoExample(t *testing.T) {
        a := EquipmentInfo{}
        err := xml.Unmarshal([]byte(test_EquipmentInfoExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_EquipmentInfoExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_EquipmentInfoExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_EquipmentInfoExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_GradingAssignment(t *testing.T) {
        a := GradingAssignment{}
        err := xml.Unmarshal([]byte(test_GradingAssignment), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_GradingAssignment)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_GradingAssignment not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_GradingAssignment))), string(stripEmptyTags(output)))
        }
}
func Test_test_GradingAssignmentScore(t *testing.T) {
        a := GradingAssignmentScore{}
        err := xml.Unmarshal([]byte(test_GradingAssignmentScore), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_GradingAssignmentScore)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_GradingAssignmentScore not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_GradingAssignmentScore))), string(stripEmptyTags(output)))
        }
}
func Test_test_example0186(t *testing.T) {
        a := Identity{}
        err := xml.Unmarshal([]byte(test_example0186), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_example0186)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_example0186 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_example0186))), string(stripEmptyTags(output)))
        }
}
func Test_test_IdentitypublishedbyMicrosoftActiveDirectory(t *testing.T) {
        a := Identity{}
        err := xml.Unmarshal([]byte(test_IdentitypublishedbyMicrosoftActiveDirectory), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_IdentitypublishedbyMicrosoftActiveDirectory)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_IdentitypublishedbyMicrosoftActiveDirectory not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_IdentitypublishedbyMicrosoftActiveDirectory))), string(stripEmptyTags(output)))
        }
}
func Test_test_IdentitypublishedbyanOpenIDprovider(t *testing.T) {
        a := Identity{}
        err := xml.Unmarshal([]byte(test_IdentitypublishedbyanOpenIDprovider), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_IdentitypublishedbyanOpenIDprovider)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_IdentitypublishedbyanOpenIDprovider not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_IdentitypublishedbyanOpenIDprovider))), string(stripEmptyTags(output)))
        }
}
func Test_test_LEAInfo(t *testing.T) {
        a := LEAInfo{}
        err := xml.Unmarshal([]byte(test_LEAInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_LEAInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_LEAInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_LEAInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_LearningResource(t *testing.T) {
        a := LearningResource{}
        err := xml.Unmarshal([]byte(test_LearningResource), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_LearningResource)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_LearningResource not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_LearningResource))), string(stripEmptyTags(output)))
        }
}
func Test_test_LearningResourcePackage(t *testing.T) {
        a := LearningResourcePackage{}
        err := xml.Unmarshal([]byte(test_LearningResourcePackage), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_LearningResourcePackage)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_LearningResourcePackage not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_LearningResourcePackage))), string(stripEmptyTags(output)))
        }
}
func Test_test_LearningStandardItem(t *testing.T) {
        a := LearningStandardItem{}
        err := xml.Unmarshal([]byte(test_LearningStandardItem), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_LearningStandardItem)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_LearningStandardItem not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_LearningStandardItem))), string(stripEmptyTags(output)))
        }
}
func Test_test_MarkValueInfo(t *testing.T) {
        a := MarkValueInfo{}
        err := xml.Unmarshal([]byte(test_MarkValueInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_MarkValueInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_MarkValueInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_MarkValueInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_PersonPicture(t *testing.T) {
        a := PersonPicture{}
        err := xml.Unmarshal([]byte(test_PersonPicture), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_PersonPicture)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_PersonPicture not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_PersonPicture))), string(stripEmptyTags(output)))
        }
}
func Test_test_PersonalisedPlan(t *testing.T) {
        a := PersonalisedPlan{}
        err := xml.Unmarshal([]byte(test_PersonalisedPlan), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_PersonalisedPlan)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_PersonalisedPlan not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_PersonalisedPlan))), string(stripEmptyTags(output)))
        }
}
func Test_test_ResourceBookingExample(t *testing.T) {
        a := ResourceBooking{}
        err := xml.Unmarshal([]byte(test_ResourceBookingExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ResourceBookingExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ResourceBookingExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ResourceBookingExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_ResourceUsage(t *testing.T) {
        a := ResourceUsage{}
        err := xml.Unmarshal([]byte(test_ResourceUsage), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ResourceUsage)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ResourceUsage not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ResourceUsage))), string(stripEmptyTags(output)))
        }
}
func Test_test_RoomInfoExample(t *testing.T) {
        a := RoomInfo{}
        err := xml.Unmarshal([]byte(test_RoomInfoExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_RoomInfoExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_RoomInfoExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_RoomInfoExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_ScheduledActivityExample(t *testing.T) {
        a := ScheduledActivity{}
        err := xml.Unmarshal([]byte(test_ScheduledActivityExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ScheduledActivityExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ScheduledActivityExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ScheduledActivityExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_SchoolCourseInfoExample(t *testing.T) {
        a := SchoolCourseInfo{}
        err := xml.Unmarshal([]byte(test_SchoolCourseInfoExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SchoolCourseInfoExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SchoolCourseInfoExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SchoolCourseInfoExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_SchoolInfo(t *testing.T) {
        a := SchoolInfo{}
        err := xml.Unmarshal([]byte(test_SchoolInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SchoolInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SchoolInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SchoolInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_SchoolPrograms(t *testing.T) {
        a := SchoolPrograms{}
        err := xml.Unmarshal([]byte(test_SchoolPrograms), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SchoolPrograms)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SchoolPrograms not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SchoolPrograms))), string(stripEmptyTags(output)))
        }
}
func Test_test_SectionInfoExample1(t *testing.T) {
        a := SectionInfo{}
        err := xml.Unmarshal([]byte(test_SectionInfoExample1), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SectionInfoExample1)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SectionInfoExample1 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SectionInfoExample1))), string(stripEmptyTags(output)))
        }
}
func Test_test_SessionInfoExample(t *testing.T) {
        a := SessionInfo{}
        err := xml.Unmarshal([]byte(test_SessionInfoExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SessionInfoExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SessionInfoExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SessionInfoExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_StaffAssignment(t *testing.T) {
        a := StaffAssignment{}
        err := xml.Unmarshal([]byte(test_StaffAssignment), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StaffAssignment)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StaffAssignment not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StaffAssignment))), string(stripEmptyTags(output)))
        }
}
func Test_test_StaffPersonal(t *testing.T) {
        a := StaffPersonal{}
        err := xml.Unmarshal([]byte(test_StaffPersonal), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StaffPersonal)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StaffPersonal not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StaffPersonal))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentActivityInfo(t *testing.T) {
        a := StudentActivityInfo{}
        err := xml.Unmarshal([]byte(test_StudentActivityInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentActivityInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentActivityInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentActivityInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentActivityParticipation(t *testing.T) {
        a := StudentActivityParticipation{}
        err := xml.Unmarshal([]byte(test_StudentActivityParticipation), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentActivityParticipation)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentActivityParticipation not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentActivityParticipation))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentAttendanceSummary(t *testing.T) {
        a := StudentAttendanceSummary{}
        err := xml.Unmarshal([]byte(test_StudentAttendanceSummary), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentAttendanceSummary)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentAttendanceSummary not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentAttendanceSummary))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentAttendanceTimeListExample(t *testing.T) {
        a := StudentAttendanceTimeList{}
        err := xml.Unmarshal([]byte(test_StudentAttendanceTimeListExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentAttendanceTimeListExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentAttendanceTimeListExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentAttendanceTimeListExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentAttendanceTimeListExample2(t *testing.T) {
        a := StudentAttendanceTimeList{}
        err := xml.Unmarshal([]byte(test_StudentAttendanceTimeListExample2), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentAttendanceTimeListExample2)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentAttendanceTimeListExample2 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentAttendanceTimeListExample2))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentContactPersonalExample(t *testing.T) {
        a := StudentContactPersonal{}
        err := xml.Unmarshal([]byte(test_StudentContactPersonalExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentContactPersonalExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentContactPersonalExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentContactPersonalExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentContactRelationship(t *testing.T) {
        a := StudentContactRelationship{}
        err := xml.Unmarshal([]byte(test_StudentContactRelationship), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentContactRelationship)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentContactRelationship not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentContactRelationship))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentDailyAttendance(t *testing.T) {
        a := StudentDailyAttendance{}
        err := xml.Unmarshal([]byte(test_StudentDailyAttendance), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentDailyAttendance)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentDailyAttendance not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentDailyAttendance))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentGrade(t *testing.T) {
        a := StudentGrade{}
        err := xml.Unmarshal([]byte(test_StudentGrade), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentGrade)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentGrade not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentGrade))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentParticipation(t *testing.T) {
        a := StudentParticipation{}
        err := xml.Unmarshal([]byte(test_StudentParticipation), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentParticipation)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentParticipation not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentParticipation))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentPeriodAttendanceExample(t *testing.T) {
        a := StudentPeriodAttendance{}
        err := xml.Unmarshal([]byte(test_StudentPeriodAttendanceExample), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentPeriodAttendanceExample)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentPeriodAttendanceExample not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentPeriodAttendanceExample))), string(stripEmptyTags(output)))
        }
}
func Test_test_example0930(t *testing.T) {
        a := StudentPersonal{}
        err := xml.Unmarshal([]byte(test_example0930), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_example0930)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_example0930 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_example0930))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentSchoolEnrollment(t *testing.T) {
        a := StudentSchoolEnrollment{}
        err := xml.Unmarshal([]byte(test_StudentSchoolEnrollment), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentSchoolEnrollment)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentSchoolEnrollment not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentSchoolEnrollment))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously(t *testing.T) {
        a := StudentScoreJudgementAgainstStandard{}
        err := xml.Unmarshal([]byte(test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished(t *testing.T) {
        a := StudentScoreJudgementAgainstStandard{}
        err := xml.Unmarshal([]byte(test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentSectionEnrollment(t *testing.T) {
        a := StudentSectionEnrollment{}
        err := xml.Unmarshal([]byte(test_StudentSectionEnrollment), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentSectionEnrollment)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentSectionEnrollment not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentSectionEnrollment))), string(stripEmptyTags(output)))
        }
}
func Test_test_SystemRole(t *testing.T) {
        a := SystemRole{}
        err := xml.Unmarshal([]byte(test_SystemRole), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_SystemRole)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_SystemRole not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_SystemRole))), string(stripEmptyTags(output)))
        }
}
func Test_test_TeachingGroup(t *testing.T) {
        a := TeachingGroup{}
        err := xml.Unmarshal([]byte(test_TeachingGroup), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_TeachingGroup)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_TeachingGroup not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_TeachingGroup))), string(stripEmptyTags(output)))
        }
}
func Test_test_TermInfo(t *testing.T) {
        a := TermInfo{}
        err := xml.Unmarshal([]byte(test_TermInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_TermInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_TermInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_TermInfo))), string(stripEmptyTags(output)))
        }
}
func Test_test_TimeTable(t *testing.T) {
        a := TimeTable{}
        err := xml.Unmarshal([]byte(test_TimeTable), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_TimeTable)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_TimeTable not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_TimeTable))), string(stripEmptyTags(output)))
        }
}
func Test_test_TimeTableCell(t *testing.T) {
        a := TimeTableCell{}
        err := xml.Unmarshal([]byte(test_TimeTableCell), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_TimeTableCell)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_TimeTableCell not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_TimeTableCell))), string(stripEmptyTags(output)))
        }
}
func Test_test_TimeTableSubject(t *testing.T) {
        a := TimeTableSubject{}
        err := xml.Unmarshal([]byte(test_TimeTableSubject), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_TimeTableSubject)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_TimeTableSubject not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_TimeTableSubject))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingAlert(t *testing.T) {
        a := WellbeingAlert{}
        err := xml.Unmarshal([]byte(test_WellbeingAlert), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingAlert)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingAlert not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingAlert))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingAppeal(t *testing.T) {
        a := WellbeingAppeal{}
        err := xml.Unmarshal([]byte(test_WellbeingAppeal), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingAppeal)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingAppeal not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingAppeal))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingCharacteristic(t *testing.T) {
        a := WellbeingCharacteristic{}
        err := xml.Unmarshal([]byte(test_WellbeingCharacteristic), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingCharacteristic)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingCharacteristic not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingCharacteristic))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingEvent(t *testing.T) {
        a := WellbeingEvent{}
        err := xml.Unmarshal([]byte(test_WellbeingEvent), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingEvent)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingEvent not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingEvent))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingEvent2(t *testing.T) {
        a := WellbeingEvent{}
        err := xml.Unmarshal([]byte(test_WellbeingEvent2), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingEvent2)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingEvent2 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingEvent2))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingEvent3(t *testing.T) {
        a := WellbeingEvent{}
        err := xml.Unmarshal([]byte(test_WellbeingEvent3), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingEvent3)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingEvent3 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingEvent3))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingPersonLink(t *testing.T) {
        a := WellbeingPersonLink{}
        err := xml.Unmarshal([]byte(test_WellbeingPersonLink), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingPersonLink)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingPersonLink not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingPersonLink))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingPersonLink2(t *testing.T) {
        a := WellbeingPersonLink{}
        err := xml.Unmarshal([]byte(test_WellbeingPersonLink2), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingPersonLink2)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingPersonLink2 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingPersonLink2))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingPersonLink3(t *testing.T) {
        a := WellbeingPersonLink{}
        err := xml.Unmarshal([]byte(test_WellbeingPersonLink3), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingPersonLink3)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingPersonLink3 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingPersonLink3))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingPersonLink4(t *testing.T) {
        a := WellbeingPersonLink{}
        err := xml.Unmarshal([]byte(test_WellbeingPersonLink4), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingPersonLink4)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingPersonLink4 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingPersonLink4))), string(stripEmptyTags(output)))
        }
}
func Test_test_WellbeingResponse(t *testing.T) {
        a := WellbeingResponse{}
        err := xml.Unmarshal([]byte(test_WellbeingResponse), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_WellbeingResponse)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_WellbeingResponse not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_WellbeingResponse))), string(stripEmptyTags(output)))
        }
}
func Test_test_ChargedLocation(t *testing.T) {
        a := ChargedLocationInfo{}
        err := xml.Unmarshal([]byte(test_ChargedLocation), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_ChargedLocation)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_ChargedLocation not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_ChargedLocation))), string(stripEmptyTags(output)))
        }
}
func Test_test_CollectionRound(t *testing.T) {
        a := CollectionRound{}
        err := xml.Unmarshal([]byte(test_CollectionRound), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_CollectionRound)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_CollectionRound not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_CollectionRound))), string(stripEmptyTags(output)))
        }
}
func Test_test_CollectionStatus(t *testing.T) {
        a := CollectionStatus{}
        err := xml.Unmarshal([]byte(test_CollectionStatus), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_CollectionStatus)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_CollectionStatus not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_CollectionStatus))), string(stripEmptyTags(output)))
        }
}
func Test_test_Debtor(t *testing.T) {
        a := Debtor{}
        err := xml.Unmarshal([]byte(test_Debtor), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_Debtor)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_Debtor not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_Debtor))), string(stripEmptyTags(output)))
        }
}
func Test_test_FinancialAccount(t *testing.T) {
        a := FinancialAccount{}
        err := xml.Unmarshal([]byte(test_FinancialAccount), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_FinancialAccount)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_FinancialAccount not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_FinancialAccount))), string(stripEmptyTags(output)))
        }
}
func Test_test_FinancialQuestionnaireCollection(t *testing.T) {
        a := FinancialQuestionnaireCollection{}
        err := xml.Unmarshal([]byte(test_FinancialQuestionnaireCollection), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_FinancialQuestionnaireCollection)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_FinancialQuestionnaireCollection not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_FinancialQuestionnaireCollection))), string(stripEmptyTags(output)))
        }
}
func Test_test_Invoice(t *testing.T) {
        a := Invoice{}
        err := xml.Unmarshal([]byte(test_Invoice), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_Invoice)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_Invoice not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_Invoice))), string(stripEmptyTags(output)))
        }
}
func Test_test_Journal(t *testing.T) {
        a := Journal{}
        err := xml.Unmarshal([]byte(test_Journal), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_Journal)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_Journal not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_Journal))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPCodeFrameexample1(t *testing.T) {
        a := NAPCodeFrame{}
        err := xml.Unmarshal([]byte(test_NAPCodeFrameexample1), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPCodeFrameexample1)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPCodeFrameexample1 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPCodeFrameexample1))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPEventStudentLink(t *testing.T) {
        a := NAPEventStudentLink{}
        err := xml.Unmarshal([]byte(test_NAPEventStudentLink), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPEventStudentLink)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPEventStudentLink not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPEventStudentLink))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPStudentResponseSet(t *testing.T) {
        a := NAPStudentResponseSet{}
        err := xml.Unmarshal([]byte(test_NAPStudentResponseSet), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPStudentResponseSet)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPStudentResponseSet not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPStudentResponseSet))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPTest(t *testing.T) {
        a := NAPTest{}
        err := xml.Unmarshal([]byte(test_NAPTest), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPTest)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPTest not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPTest))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPTestItem(t *testing.T) {
        a := NAPTestItem{}
        err := xml.Unmarshal([]byte(test_NAPTestItem), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPTestItem)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPTestItem not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPTestItem))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPTestScoreSummary(t *testing.T) {
        a := NAPTestScoreSummary{}
        err := xml.Unmarshal([]byte(test_NAPTestScoreSummary), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPTestScoreSummary)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPTestScoreSummary not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPTestScoreSummary))), string(stripEmptyTags(output)))
        }
}
func Test_test_NAPTestlet(t *testing.T) {
        a := NAPTestlet{}
        err := xml.Unmarshal([]byte(test_NAPTestlet), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_NAPTestlet)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_NAPTestlet not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_NAPTestlet))), string(stripEmptyTags(output)))
        }
}
func Test_test_PaymentReceipt(t *testing.T) {
        a := PaymentReceipt{}
        err := xml.Unmarshal([]byte(test_PaymentReceipt), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_PaymentReceipt)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_PaymentReceipt not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_PaymentReceipt))), string(stripEmptyTags(output)))
        }
}
func Test_test_example0815(t *testing.T) {
        a := PurchaseOrder{}
        err := xml.Unmarshal([]byte(test_example0815), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_example0815)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_example0815 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_example0815))), string(stripEmptyTags(output)))
        }
}
func Test_test_StudentAttendanceCollectionExample2(t *testing.T) {
        a := StudentAttendanceCollection{}
        err := xml.Unmarshal([]byte(test_StudentAttendanceCollectionExample2), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_StudentAttendanceCollectionExample2)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_StudentAttendanceCollectionExample2 not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_StudentAttendanceCollectionExample2))), string(stripEmptyTags(output)))
        }
}
func Test_test_VendorInfo(t *testing.T) {
        a := VendorInfo{}
        err := xml.Unmarshal([]byte(test_VendorInfo), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte(test_VendorInfo)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("test_VendorInfo not DeepEqual:\nmv: %+v\nwant: %+v\n%s\n%s\n", mv, want, string(stripEmptyTags([]byte(test_VendorInfo))), string(stripEmptyTags(output)))
        }
}
