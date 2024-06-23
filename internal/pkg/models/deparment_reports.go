package models

import (
	"fmt"
	"strings"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
)

type ReportOne struct {
	FullName            string `json:"ФИО"`
	SupervisorName      string `json:"Научный руководитель"`
	AttestationMark     int32  `json:"Научно-исследовательская работа (оценка должна соответствовать ведомости)"`
	Progressiveness     int32  `json:"Процент выполнения диссертационного исследования"`
	ScientificWorkCount string `json:"Количество научных работ"`
	Qualified           string `json:"Отметка об аттестации"`
	ActualYear          int32  `json:"Курс обучения"`
}

type ScientificWorkCount struct {
	WAC    int `json:"ВАК"`
	Scopus int `json:"Scopus"`
	Rinc   int `json:"РИНЦ"`
	Wos    int `json:"WOS"`
}

func (s *ScientificWorkCount) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("ВАК: %d\n", s.WAC))
	builder.WriteString(fmt.Sprintf("Scopus: %d\n", s.Scopus))
	builder.WriteString(fmt.Sprintf("РИНЦ: %d\n", s.Rinc))
	builder.WriteString(fmt.Sprintf("WOS: %d", s.Wos))
	return builder.String()
}

type StudentInfoForReportOne struct {
	Student         model.Students
	SupervisorName  string
	AttestationMark int32
	Progressiveness int32
	Conferences     []model.Conferences
	Publications    []model.Publications
}

type ReportTwo struct {
	FullName            string `json:"ФИО"`
	SupervisorName      string `json:"Научный руководитель"`
	DissertationTitle   string `json:"Тема диссертационного исследования"`
	SupervisorMark      int32  `json:"Оценка научного руководителя за НИР"`
	Progressiveness     int32  `json:"Процент выполнения диссертационного исследования"`
	TeachingLoad        int32  `json:"Текущая педагогическая нагрузка (совокупность за все семестры)"`
	ScientificWorkCount string `json:"Количество научных публикаций"`
	ActualYear          int32  `json:"Курс обучения"`
	Comment             string `json:"Примечания"`
}

type StudentInfoForReportTwo struct {
	Student           model.Students
	SupervisorName    string
	SupervisorMark    int32
	Progressiveness   int32
	Publications      []model.Publications
	ClassroomLoad     []model.ClassroomLoad
	DissertationTitle model.DissertationTitles
}
