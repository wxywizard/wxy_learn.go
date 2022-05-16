package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	f, err := excelize.OpenFile("/Users/will/Downloads/经营分析资料/标杆门店/tm_cre_appraisal_new.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet6", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("sheet2")
	sql := make([]string, 0)
	fmt.Println(len(rows))
	for i, row := range rows[1:] {
		//for _, colCell := range row {
		//	fmt.Print(colCell, "\t")
		//}
		str := generatorSql(row)
		fmt.Println("index===", i)
		sql = append(sql, str)
	}
	writeFileWithAppend("/Users/will/Downloads/tm_cre_appraisal_new_data.sql", sql)
}

func writeFileWithAppend(filePath string, sql []string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// b := make([]byte, 10)
	// var n int
	for _, s := range sql {
		_, err = file.WriteString(s)
	}

}

func generatorSql(row []string) (str string) {
	fmt.Println(len(row))
	str = "insert into dw_cre.ta_cre_benchmark_store_appraisal" +
		"(apsl_dim, kpi_type_code, kpi_type_name, opr_stage_type," +
		"product_line_type, kpi_code, kpi_name, kpi_definition, apsl_rule1_desc," +
		"apsl_rule1_value, apsl_rule1_point, apsl_rule2_desc, apsl_rule2_value," +
		"apsl_rule2_point, apsl_rule3_desc, apsl_rule3_value, apsl_rule3_point," +
		"apsl_rule4_desc, apsl_rule4_value, apsl_rule4_point, apsl_rule5_desc," +
		"apsl_rule5_value, apsl_rule5_point, is_del, apsl_type,unit)" +
		" values " +
		"('" + row[0] + "','" + row[1] + "','" + row[2] + "','" + row[3] + "','" + row[4] + "','" + row[5] + "'," +
		"'" + row[6] + "','" + row[7] + "','" + row[8] + "','" + row[9] + "','" + row[10] + "','" + row[11] + "'," +
		"'" + row[12] + "','" + row[13] + "','" + row[14] + "','" + row[15] + "','" + row[16] + "','" + row[17] + "'," +
		"'" + row[18] + "','" + row[19] + "','" + row[20] + "','" + row[21] + "','" + row[22] + "','" + row[23] + "'," +
		"'" + row[24] + "','" + row[25] + "');\n"
	return

}
