package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	f, err := excelize.OpenFile("/Users/will/Downloads/经营分析资料/科目映射表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	generateSql(f)
}

func generateSql(f *excelize.File) {
	rows, err := f.GetRows("Sheet3")
	sql := make([]string, 0)
	if err == nil {
		rowsR := rows[2:]
		for i, row := range rowsR {
			if i < 108 {
				gw := generatorGW(row)
				sql = append(sql, gw)
			}
			if i < 47 {
				bc := generatorBC(row)
				sql = append(sql, bc)
			}
			if i < 43 {
				wy := generatorWY(row)
				sql = append(sql, wy)
			}
			if i < 41 {
				wy := generatorZT(row)
				sql = append(sql, wy)
			}
		}
		fmt.Println(len(sql))
		writeFileWithAppend("/Users/will/Downloads/tc_mapping_account_6s_act_man_data_new1.sql", sql)
	}

}

func generatorGW(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1," +
		"segment_code2,segment_code3,sbu_code) VALUES ('" + row[2] + "'," +
		"" + row[3] + ",'" + row[4] + "','" + row[0] + "','" + row[1] + "','" + row[5] + "','" + row[6] + "','" + row[7] + "','" + row[8] + "');\n"
	return str
}

func generatorBC(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1,sbu_code) VALUES ('" + row[12] + "'," +
		"" + row[13] + ",'" + row[14] + "','" + row[10] + "','" + row[11] + "','" + row[15] + "','" + row[16] + "');\n"
	return str
}

func generatorWY(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1," +
		"segment_code2,segment_code3,sbu_code) VALUES ('" + row[20] + "'," +
		"" + row[21] + ",'" + row[22] + "','" + row[18] + "','" + row[19] + "','" + row[23] + "','" + row[24] + "','" + row[25] + "','" + row[26] + "');\n"
	return str
}

func generatorZT(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1," +
		"sbu_code) VALUES ('" + row[30] + "'," +
		"" + row[31] + ",'" + row[32] + "','" + row[28] + "','" + row[29] + "','" + row[33] + "','" + row[34] + "');\n"
	return str
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
