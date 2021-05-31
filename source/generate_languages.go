// +build ignore

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	tr        *translate.Translate
	awsKey    string
	awsSecret string
)

type Text struct {
	Key string
	En  string
	Fr  string
	De  string
	Ru  string
	Sp  string
	Jp  string
	Cn  string
	Ko  string
	It  string
}

func main() {
	fmt.Println("RUNNING: ./source/generate_languages.go")
	awsKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if awsKey == "" || awsSecret == "" {
		fmt.Println("AWS_ACCESS_KEY_ID or AWS_SECRET_ACCESS_KEY not set")
		os.Exit(0)
		return
	}

	InitAWS()

	file, _ := os.Open("../frontend/src/languages/data.csv")

	defer file.Close()
	c := csv.NewReader(file)

	var translations []*Text
	line := 0
	for {
		// Read each record from csv
		record, err := c.Read()
		if err == io.EOF {
			break
		}
		if line == 0 {
			line++
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		key := record[0]
		english := record[1]

		translated := TranslateAll(key, english)

		translations = append(translations, translated)

		fmt.Printf("%s | English: %s | French: %s | German: %s | Russian: %s\n", translated.Key, translated.En, translated.Fr, translated.De, translated.Ru)
		line++

		time.Sleep(250 * time.Millisecond)
	}

	//CreateGo(translations)

	CreateJS("english", translations)
	CreateJS("russian", translations)
	CreateJS("french", translations)
	CreateJS("german", translations)
	CreateJS("spanish", translations)
	CreateJS("japanese", translations)
	CreateJS("chinese", translations)
	CreateJS("italian", translations)
	CreateJS("korean", translations)
}

func Translate(val, language string) string {
	input := &translate.TextInput{
		SourceLanguageCode: aws.String("en"),
		TargetLanguageCode: aws.String(language),
		Text:               aws.String(val),
	}
	req, out := tr.TextRequest(input)
	if err := req.Send(); err != nil {
		panic(req.Error)
	}
	return *out.TranslatedText
}

func TranslateAll(key, en string) *Text {
	return &Text{
		Key: key,
		En:  en,
		Fr:  Translate(en, "fr"),
		De:  Translate(en, "de"),
		Ru:  Translate(en, "ru"),
		Sp:  Translate(en, "es"),
		Jp:  Translate(en, "ja"),
		Cn:  Translate(en, "zh"),
		Ko:  Translate(en, "ko"),
		It:  Translate(en, "it"),
	}
}

func (t *Text) String(lang string) string {
	switch lang {
	case "english":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.En)
	case "russian":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Ru)
	case "spanish":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Sp)
	case "german":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.De)
	case "french":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Fr)
	case "japanese":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Jp)
	case "chinese":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Cn)
	case "korean":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.Ko)
	case "italian":
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.It)
	default:
		return fmt.Sprintf(`    %s: "%s"`, t.Key, t.En)
	}
}

func GoLang(trs []*Text) string {
	var allvars []string
	languages := []string{"english", "russian"}
	for _, language := range languages {
		allvars = append(allvars, language+" := make(map[string]string)")
		for _, t := range trs {
			allvars = append(allvars, GoLine(language, t))
		}
		allvars = append(allvars, "\nLanguage[\""+language+"\"] = "+language)
	}
	return strings.Join(allvars, "\n")
}

func GoLine(lang string, t *Text) string {
	switch lang {
	case "english":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.En)
	case "russian":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Ru)
	case "spanish":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Sp)
	case "german":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.De)
	case "french":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Fr)
	case "japanese":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Jp)
	case "chinese":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Cn)
	case "korean":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.Ko)
	case "italian":
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.It)
	default:
		return fmt.Sprintf(`		%s["%s"] = "%s"`, lang, t.Key, t.En)
	}
}

func CreateGo(trs []*Text) {

	data := `package utils

var Language map[string]map[string]string

func init() {
	Language = make(map[string]map[string]string)
	` + GoLang(trs) + `
}
`

	ioutil.WriteFile("../utils/languages.go", []byte(data), os.ModePerm)
}

func CreateJS(name string, trs []*Text) {

	data := "const " + name + " = {\n"

	var allvars []string
	for _, v := range trs {
		allvars = append(allvars, v.String(name))
	}

	data += strings.Join(allvars, ",\n")

	data += "\n}\n\nexport default " + name

	ioutil.WriteFile("../frontend/src/languages/"+name+".js", []byte(data), os.ModePerm)

}

func InitAWS() {
	creds := credentials.NewStaticCredentials(awsKey, awsSecret, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: creds,
	})
	if err != nil {
		panic(err)
	}
	tr = translate.New(sess)
}
