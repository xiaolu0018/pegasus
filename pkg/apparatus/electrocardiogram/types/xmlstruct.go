//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：xml相关结构体
//创建人：张志浩
//修改记录：若要修改请记录
//==================================================================
package types

import "encoding/xml"

type Value struct {
	Value_1       string         `xml:"value,attr"`
}

type Code struct {
	Code_1        string         `xml:"code,attr"`
}

type  Annotation struct {
	Code           Code          `xml:"code"`
	Value          Value         `xml:"value"`
}

type Component_1 struct {
	Annotation     Annotation    `xml:"annotation"`
}

type AnnotationSet struct {
	Component_1    []Component_1  `xml:"component"`
}

type SubjectOf struct {
	AnnotationSet   AnnotationSet  `xml:"annotationSet"`
}
type  Analysis struct {
	Gread           string       `xml:"gread"`
	Digcode         string       `xml:"digcode"`
	Miscode         string       `xml:"miscode"`
}

type Series struct {
	SubjectOf      SubjectOf     `xml:"subjectOf"`
	Analysis       Analysis      `xml:"analysis"`
}

type Component struct {
	Series         Series         `xml:"series"`
}

type Center struct {
	Value          string          `xml:"value,attr"`
}

type EffectiveTime struct {
	Center         Center          `xml:"center"`
}

type Result struct {
	XMLName        xml.Name
	Effective      EffectiveTime  `xml:"effectiveTime"`
	Component      Component      `xml:"component"`
}
