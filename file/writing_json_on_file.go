package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type activitiesStats struct {
	AcDupQuestText    []string     `json:"activities_duplicated_questions_text"`
	AcDupQuestCode    []string     `json:"activities_duplicated_questions_code"`
	AcDupQuestQttText int          `json:"activities_duplicated_questions_text_quantity"`
	AcDupQuestQttCode int          `json:"activities_duplicated_questions_code_quantity"`
	TotalAc           int          `json:"total_activities"`
	MappedAc          int          `json:"mapped_activities"`
	TotalQuest        int          `json:"total_questions"`
	QuestionsMean     int          `json:"questions_per_activities"`
	FaqText           []faqTextObj `json:"faq_text"`
	FaqCode           []faqCodeObj `json:"faq_code"`
	ErrorAc           []string     `json:"error_while_fetching_activities"`
	AddrTypes         []string     `json:"address_types"`
	AddrPointTypes    []string     `json:"address_point_types"`
}

type activity struct {
	Code       string     `json:"code"`
	Modalities []modality `json:"modalities"`
	Name       string     `json:"name"`
}

type modality struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Questions []question `json:"questions"`
}

type question struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type faqTextObj struct {
	Text     string `json:"text"`
	Quantity int    `json:"quantity"`
}

type faqCodeObj struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

func main() {
	q1 := question{
		Code: "1",
		Text: "t1",
	}
	q2 := question{
		Code: "2",
		Text: "t2",
	}
	q3 := question{
		Code: "3",
		Text: "t2",
	}
	q4 := question{
		Code: "3",
		Text: "t4",
	}
	m1 := modality{
		Code:      "codeModality1",
		Name:      "nameModality1",
		Questions: []question{q1, q2, q3, q4},
	}
	m2 := modality{
		Code:      "codeModality2",
		Name:      "nameModality2",
		Questions: []question{q1, q2, q3, q4},
	}
	ac1 := activity{
		Code:       "codeActivity1",
		Modalities: []modality{m1},
		Name:       "nameActivity1",
	}
	ac2 := activity{
		Code:       "codeActivity2",
		Modalities: []modality{m2},
		Name:       "nameActivity2",
	}
	acStats := activitiesStats{}

	xac := []activity{ac1, ac2}
	faqText := map[string]int{}
	faqCode := map[string]int{}
	for _, ac := range xac {
		acStats.MappedAc++
		acStats.TotalAc++
		alreadyCountedAcText := false
		alreadyCountedAcCode := false

		for _, m := range ac.Modalities {
			DupQuestionsText := map[string]int{}
			DupQuestionsCode := map[string]int{}

			for _, q := range m.Questions {
				acStats.TotalQuest++
				q.Text = strings.ToUpper(q.Text)
				q.Code = strings.ToUpper(q.Code)
				// text FAQ
				{
					faqText[q.Text]++
				}
				// text duplicated
				{
					DupQuestionsText[q.Text]++
					if DupQuestionsText[q.Text] > 1 && !alreadyCountedAcText {
						acStats.AcDupQuestText = append(acStats.AcDupQuestText, ac.Code)
						alreadyCountedAcText = true
					}
				}
				// code FAQ
				{
					faqCode[q.Code]++
				}
				// code duplicated
				{
					DupQuestionsCode[q.Code]++
					if DupQuestionsCode[q.Code] > 1 && !alreadyCountedAcCode {
						acStats.AcDupQuestCode = append(acStats.AcDupQuestCode, ac.Code)
						alreadyCountedAcCode = true
					}
				}
			}
		}
	}
	for qt, qtt := range faqText {
		f := faqTextObj{
			Text:     qt,
			Quantity: qtt,
		}
		acStats.FaqText = append(acStats.FaqText, f)
	}
	for qc, qtt := range faqCode {
		f := faqCodeObj{
			Code:     qc,
			Quantity: qtt,
		}
		acStats.FaqCode = append(acStats.FaqCode, f)
	}
	acStats.AcDupQuestQttText = len(acStats.AcDupQuestText)
	acStats.AcDupQuestQttCode = len(acStats.AcDupQuestCode)
	acStats.QuestionsMean = acStats.TotalQuest / acStats.TotalAc
	sort.Slice(acStats.FaqText, func(i, j int) bool { return acStats.FaqText[i].Quantity > acStats.FaqText[j].Quantity })
	sort.Slice(acStats.FaqCode, func(i, j int) bool { return acStats.FaqCode[i].Quantity > acStats.FaqCode[j].Quantity })
	xb, err := json.MarshalIndent(acStats, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(xb))

	err = ioutil.WriteFile("file/stats.json", xb, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
