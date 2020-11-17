package populate

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func Errcheck(err error) {
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

func PrintXML(record interface{}) error {
	output, err := xml.MarshalIndent(record, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
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

func birth_year(yearlevel int) string {
	minyear := time.Now().Year() - 5
	gofakeit.Seed(0)
	return gofakeit.DateRange(time.Date(minyear-yearlevel-1, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(minyear-yearlevel, time.December, 31, 0, 0, 0, 0, time.UTC)).Format("2006-01-02")
}

// assumes probabilities ranked in descending order from 1.0 to 0, and that cardinality of
// probabilities is one less than that of options
func threshold_rand_strings(probabilities []float64, options []string) string {
	r := rand.Float64()
	if len(probabilities) != len(options)-1 {
		log.Fatalf("threshold_rand_strings: The cardinality of %+v is not one less than the cardinality of %+v", probabilities, options)
	}
	for i, prob := range probabilities {
		if r > prob {
			return options[i]
		}
	}
	return options[len(options)-1]
}

func school_name() string {
	gofakeit.Seed(0)
	return gofakeit.City() + " " + randomStringFromSlice([]string{"Academy", "Grammar", "College"})
}

func create_email(given string, middle string, family string) string {
	domain := randomStringFromSlice([]string{"mail.vic.edu.au", "people.vic.edu.au", "vic.edu.au", "dashboard.vic.edu.au", "distance.vic.edu.au"})
	return family + "." + given + "." + string(middle[0]) + "@" + domain
}

func create_phone_number(state *string) string {
	return fmt.Sprintf("04%08d", rand.Intn(100000000))
}

func teachingGroupLongName(shortname string) string {
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

func teachingGroupKLA(shortname string) string {
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
