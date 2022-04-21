package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	f, err := excelize.OpenFile("/Users/will/Downloads/经营分析资料/组织映射表.xlsx")
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
	rows, err := f.GetRows("sheet1")
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
	writeFileWithAppend("/Users/will/Downloads/td_cre_6s_erp_project_map_data.sql", sql)
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
	if len(row) > 4 {
		str = "insert into atm_dw.td_cre_6s_erp_project_map " +
			"(bi_project_code,bi_project_name,staging_code,staging_name," +
			"f6s_project_code_overall,f6s_project_name_overall,f6s_project_code_child," +
			"f6s_project_name_child,subsidiary_business,budget_project_code,budget_project_name," +
			"flag) values ('" + row[0] + "','" + row[1] + "','" + row[2] + "','" + row[3] + "','" + row[4] + "','" +
			row[5] + "','" + row[6] + "','" + row[7] + "','" + row[8] + "','" + row[9] + "','" + row[10] + "',false);\n"
	} else {
		str = "insert into atm_dw.td_cre_6s_erp_project_map " +
			"(bi_project_code,bi_project_name,staging_code,staging_name," +
			"f6s_project_code_overall,f6s_project_name_overall,f6s_project_code_child," +
			"f6s_project_name_child,subsidiary_business,budget_project_code,budget_project_name," +
			"flag) values ('" + row[0] + "','" + row[1] + "','" + row[2] + "','" + row[3] + "',null,null,null,null,null,null,null,false);\n"
	}
	return

}
