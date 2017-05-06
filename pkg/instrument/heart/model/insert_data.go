//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：把xml数据插入到数据库表中
//创建人：张志浩
//修改记录：若要修改请记录 郭世江修改 修改日期2017-4-27
//==================================================================
package model

import (
	"fmt"

	"bjdaos/pegasus/pkg/instrument/heart/types"
)

//将数据插入到表中
func InsertXmlDataToDB(queryCode string, result types.Result) int {
	qCode := len(queryCode)
	//测心电时扫的是腕表号
	if qCode == 8 {
		stmt, _ := DB.Prepare(`INSERT INTO heart_data
		(watch_mark,examination_no,heart_rate,time_pd_rr,time_pd_pr,qrs_front,pd_qrs,pd_qt,pd_qtc,wave_rs,wave_rv5,wave_rv6,gread,discode,miscode,confirm,annotation,result)
		 VALUES
		 ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)`)
		stmt.Exec(queryCode,
			"",
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[0].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[1].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[2].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[3].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[4].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[5].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[6].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[7].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[8].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[9].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Gread),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Digcode),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Miscode),
			"", "", fmt.Sprintf("%s", result.Component.Series.Analysis.Digcode))
	}
	//测心电扫的是体检单号
	if qCode == 16 {
		stmt, _ := DB.Prepare(`INSERT INTO heart_data
		(watch_mark, examination_no, heart_rate, time_pd_rr, time_pd_pr,qrs_front ,pd_qrs ,pd_qt, pd_qtc ,wave_rs,wave_rv5 ,wave_rv6, gread ,discode ,miscode,confirm,annotation,result)
		 VALUES
		 ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)`)
		stmt.Exec("",
			queryCode,
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[0].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[1].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[2].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[3].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[4].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[5].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[6].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[7].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[8].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.SubjectOf.AnnotationSet.Component_1[9].Annotation.Value.Value_1),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Gread),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Digcode),
			fmt.Sprintf("%s", result.Component.Series.Analysis.Miscode),
			"", "",
			fmt.Sprintf("%s", result.Component.Series.Analysis.Digcode))
	}
	return types.InsertDataDBSucc
}

//INSERT INTO heart_data( "watch_mark", "examination_no", "heart_rate", "time_pd_rr", "time_pd_pr", "qrs_front", "pd_qrs", "pd_qt", "pd_qtc", "wave_rs", "wave_rv5", "wave_rv6", "heart_analysis", "gread", "discode", "miscode", "confirm", "annotation", "result")
//VALUES ('111', '0001160001405', '60', '1000', '177', '56', '110', '439', '439', '2.30', '1.24', '1.36', '', '正常范围的心电图', '101:正常范围 ', '1-0-0 ', '16', '17', '20123762');
