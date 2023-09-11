package populate

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"
	"github.com/nsip/sifxml2go/sifxml"
)

// Abort on error.
func Errcheck(err error) {
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

// Marshal to XML, and print to output.
func PrintXML(record interface{}) error {
	output, err := xml.MarshalIndent(record, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

// Marshal to JSON, and print to output.
func PrintJSON(record interface{}) error {
	b, err := json.Marshal(record)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	json.Indent(&out, b, " ", "\t")
	out.WriteTo(os.Stdout)

	return nil
}

func create_GUID() string {
	return strings.ToUpper(uuid.New().String())
}

func randomStringFromSlice(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func this_year() string {
	return strconv.Itoa(time.Now().Year())
}

func sex_seeded(sex string) string {
	r := rand.Float64()
	if r < 0.1 {
		return "9"
	}
	if sex == "female" {
		return "2"
	}
	if sex == "male" {
		return "1"
	}
	return "9"
}

func create_salutation(sex string) string {
	if sex == "female" {
		return randomStringFromSlice([]string{"Mrs", "Dr", "Ms", "Miss"})
	} else if sex == "male" {
		return randomStringFromSlice([]string{"Mr", "Dr", "Mr", "Mr"})
	} else {
		return randomStringFromSlice([]string{"Mx", "Dr", "Mx", "Mx"})
	}
}

func create_state() string {
	return randomStringFromSlice([]string{"ACT", "NT", "QLD", "VIC", "WA", "NSW", "TAS", "SA"})
}

func postcode_seeded(state string) string {
	var ret int
	switch state {
	case "ACT":
		ret = rand.Intn(19) + 2600
	case "NT":
		ret = rand.Intn(100) + 800
	case "QLD":
		ret = rand.Intn(1000) + 4000
	case "NSW":
		ret = rand.Intn(600) + 2000
	case "VIC":
		ret = rand.Intn(1000) + 3000
	case "TAS":
		ret = rand.Intn(800) + 7000
	case "SA":
		ret = rand.Intn(800) + 5000
	case "WA":
		ret = rand.Intn(798) + 6000
	}
	return fmt.Sprintf("%04d", ret)
}

var sequences map[string]int = make(map[string]int)

func seq_gen(key string) int {
	if _, ok := sequences[key]; !ok {
		sequences[key] = 0
	}
	sequences[key] += 1
	return sequences[key]
}

func seq_gen_reset(key string) {
	sequences[key] = 0
}

var random_sequences map[string]map[int]bool = make(map[string]map[int]bool)

func random_seq_gen(key string, n int) int {
	var ok bool
	if _, ok = random_sequences[key]; !ok {
		random_sequences[key] = make(map[int]bool)
	}
	ret := rand.Intn(n)
	_, ok = random_sequences[key][ret]
	if ok {
		found := false
		for i := 0; i < n; i++ {
			if _, ok2 := random_sequences[key][i]; !ok2 {
				found = true
			}
		}
		if !found {
			log.Fatalf("no available slots remaining in random int sequence %s, max length %d", key, n)
		}
	}
	for ok {
		ret = rand.Intn(n)
		_, ok = random_sequences[key][ret]
	}
	random_sequences[key][ret] = true
	return ret
}

func random_seq_gen_reset(key string) {
	random_sequences[key] = make(map[int]bool)
}

func birth_year(yearlevel_str string) (string, error) {
	minyear := time.Now().Year() - 5
	gofakeit.Seed(0)
	yearlevel := yearlevel2yr(yearlevel_str)
	return gofakeit.DateRange(time.Date(minyear-yearlevel-1, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(minyear-yearlevel, time.December, 31, 0, 0, 0, 0, time.UTC)).Format("2006-01-02"), nil
}

// assumes probabilities ranked in descending order from 1.0 to 0 (final probability 0), and that cardinality of
// probabilities is same as that of options
func threshold_rand_strings(probabilities []float64, options []string) string {
	r := rand.Float64()
	if len(probabilities) != len(options) {
		log.Fatalf("threshold_rand_strings: The cardinality of %+v is not same as the cardinality of %+v", probabilities, options)
	}
	for i, prob := range probabilities {
		if r > prob {
			return options[i]
		}
	}
	return options[len(options)-1]
}

func school_name(schooltype string) string {
	gofakeit.Seed(0)
	school := randomStringFromSlice([]string{"Academy", "Grammar", "College"})
	switch schooltype {
	case "Camp":
		school = "Camp"
	case "Commty":
		school = "Community College"
	case "EarlyCh":
		school = "Childcare Centre"
	case "JunPri":
		school = "Primary School"
	case "Kgarten":
		school = "Kindergarten"
	case "Kind":
		school = "Kindergarten"
	case "Lang":
		school = "Language School"
	case "MCH":
		school = "Maternal Child Health Centre"
	case "Middle":
		school = "Middle School"
	case "PreSch":
		school = "Preschool"
	case "Pri/Sec":
		school = "College"
	case "Prim":
		school = "Primary School"
	case "Sec":
		school = "Secondary College"
	case "Senior":
		school = "Secondary College"
	case "Spec/P-12":
		school = "Special School"
	case "Spec/Pri":
		school = "Special School"
	case "Spec/Sec":
		school = "Special School"
	case "Special":
		school = "Special School"
	case "Supp":
		school = "Support Centre"
	}
	return gofakeit.City() + " " + school
}

func create_email(given string, middle string, family string, domain string) string {
	return family + "." + given + "." + string((middle)[0]) + "@" + domain
}

func create_school_email_domain(school *sifxml.SchoolInfo) string {
	domain := randomStringFromSlice([]string{"mail.vic.edu.au", "people.vic.edu.au", "vic.edu.au", "dashboard.vic.edu.au", "distance.vic.edu.au"})
	state := school.AddressList().Last().StateProvince().String()
	return strings.Replace(domain, "vic", strings.ToLower(state), 1)
}

func create_commercial_email_domain() string {
	gofakeit.Seed(0)
	return gofakeit.DomainName()
}

func create_phone_number(state *string) string {
	if state == nil {
		return "04"
	}
	var areacode string
	switch *state {
	case "ACT":
		areacode = "02"
	case "NSW":
		areacode = "02"
	case "VIC":
		areacode = "03"
	case "TAS":
		areacode = "03"
	case "QLD":
		areacode = "07"
	case "WA":
		areacode = "08"
	case "SA":
		areacode = "08"
	case "NT":
		areacode = "08"
	default:
		areacode = "04"
	}
	return fmt.Sprintf("%s%08d", areacode, rand.Intn(100000000))
}

func random_date(from string, to string) string {
	gofakeit.Seed(0)
	t1, err := time.Parse("2006-01-02", from)
	if err != nil {
		log.Fatalf("%s is in an illegal date format", from)
	}
	t2, err := time.Parse("2006-01-02", to)
	if err != nil {
		log.Fatalf("%s is in an illegal date format", to)
	}
	return gofakeit.DateRange(t1, t2).Format("2006-01-02")
}

// Default selection of teaching subjects: "MAT", "ENG", "PHYS", "BIO", "CHEM", "COMP", "VIS", "ECON", "HIST".
func All_teachingSubjects() []string {
	return ([]string{"MAT", "ENG", "PHYS", "BIO", "CHEM", "COMP", "VIS", "ECON", "HIST"})
}
func teachingSubject() string {
	return randomStringFromSlice(All_teachingSubjects())
}

func teachingSubjectLongName(shortname string) string {
	switch shortname {
	case "MAT":
		return "Mathematics"
	case "ENG":
		return "English"
	case "PHYS":
		return "Physics"
	case "BIO":
		return "Biology"
	case "CHEM":
		return "Chemistry"
	case "COMP":
		return "Computer Science"
	case "VIS":
		return "Visual Design"
	case "ECON":
		return "Economics"
	case "HIST":
		return "History"
	}
	return "???"
}

func TeachingGroupKLA(shortname string) string {
	switch shortname {
	case "MAT":
		return "M"
	case "ENG":
		return "E"
	case "PHYS":
		return "S"
	case "BIO":
		return "S"
	case "CHEM":
		return "S"
	case "COMP":
		return "I"
	case "VIS":
		return "D"
	case "ECON":
		return "B"
	case "HIST":
		return "H"
	}
	return "???"
}

func yearlevel2yr(yr string) int {
	switch yr {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10":
		return 10
	case "11":
		return 11
	case "12":
		return 12
	case "13":
		return 13
	case "11MINUS":
		return 6
	case "12PLUS":
		return 7
	case "P":
		return 0
	case "K":
		return -1
	case "K4":
		return -1
	case "K3":
		return -2
	case "CC":
		return -2
	case "PS":
		return -2
	case "UG":
		return rand.Intn(12) + 1
	case "UGPri":
		return rand.Intn(6) + 1
	case "UGSec":
		return rand.Intn(6) + 7
	case "UGJunSec":
		return rand.Intn(4) + 7
	case "UGSnrSec":
		return rand.Intn(2) + 11
	}
	return 1
}

// Map the School Type schooltype to a slice of year levels that a school of that type offers.
func Schooltype2Yearlevels(schooltype string) []string {
	switch schooltype {
	case "EarlyCh":
		return []string{"CC"}
	case "JunPri":
		return []string{"1", "2", "3"}
	case "Kgarten":
		return []string{"K"}
	case "Kind":
		return []string{"PS", "K"}
	case "Middle":
		return []string{"5", "6", "7", "8"}
	case "PreSch":
		return []string{"PS"}
	case "Pri/Sec":
		return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	case "Spec/P-12":
		return []string{"P", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	case "Spec/Sec":
		return []string{"7", "8", "9", "10", "11", "12"}
	case "Prim":
		return []string{"1", "2", "3", "4", "5", "6"}
	case "Sec":
		return []string{"7", "8", "9", "10", "11", "12"}
	case "Senior":
		return []string{"11", "12"}
	}
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
}

// The permitted periods for a school day.
func Periods() []string {
	return []string{"9AM", "10AM", "11AM", "12PM", "1PM", "2PM", "3PM"}
}

func periodId2Hour(id int) int {
	return 8 + id
}

// Set the time that the period with PeriodID id starts; e.g. period 1 maps to 9:05.
func PeriodStart(id int) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05.000Z", fmt.Sprintf("2000-01-01T%02d:05:00.000Z", periodId2Hour(id)))
	return t
}

// Set the time that the period with PeriodID id ends; e.g. period 1 maps to 9:55.
func PeriodEnd(id int) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05.000Z", fmt.Sprintf("2000-01-01T%02d:55:00.000Z", periodId2Hour(id)))
	return t
}

// Start date for the given semester in the given year. Set to February 1 and July 1.
func Term_start_date(year string, semester int) string {
	if semester == 1 {
		return fmt.Sprintf("%s-02-01", year)
	} else {
		return fmt.Sprintf("%s-07-01", year)
	}
}

// End date for the given semester in the given year. Set to May 30 and November 30.
func Term_end_date(year string, semester int) string {
	if semester == 1 {
		return fmt.Sprintf("%s-05-30", year)
	} else {
		return fmt.Sprintf("%s-11-30", year)
	}
}

func All_AGCollections() []string {
	return ([]string{"COI", "FQ", "SES", "STATS"})
}

func CollectionCode2Name(code string) string {
	switch code {
	case "COI":
		return "Non-Government School Census"
	case "FQ":
		return "Financial Questionnaire"
	case "SES":
		return "Address Collection"
	case "STATS":
		return "Student Attendance"
	}
	return "XXX"
}

func CollectionRoundCode(collection string, year string, round int) string {
	return fmt.Sprintf("%s %s-%02d", collection, year, round)
}

func HTTPStatus2Text(code string) string {
	switch code {
	case "201":
		return "Created"
	case "422":
		return "Unprocessable Entity"
	case "500":
		return "Internal Server Error"
	}
	return "XXX"
}
