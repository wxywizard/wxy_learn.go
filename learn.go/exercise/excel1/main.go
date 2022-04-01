package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	f, err := excelize.OpenFile("/Users/will/Downloads/科目映射表.xlsx")
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
			if i < 106 {
				gw := generatorGW(row)
				sql = append(sql, gw)
			}
			if i < 46 {
				bc := generatorBC(row)
				sql = append(sql, bc)
			}
			if i < 40 {
				wy := generatorWY(row)
				sql = append(sql, wy)
			}
		}
		fmt.Println(len(sql))
		writeFileWithAppend("/Users/will/Downloads/tc_mapping_account_6s_act_man_data.sql", sql)
	}

}

func generatorGW(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1," +
		"segment_code2,segment_code3) VALUES ('" + row[2] + "'," +
		"'" + row[3] + "','" + row[4] + "','" + row[0] + "','" + row[1] + "','" + row[5] + "','" + row[6] + "','" + row[7] + "');\n"
	return str
}

func generatorBC(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1) VALUES ('" + row[11] + "'," +
		"'" + row[12] + "','" + row[13] + "','" + row[9] + "','" + row[10] + "','" + row[14] + "');\n"
	return str
}

func generatorWY(row []string) string {
	str := "insert into atm_dw.tc_mapping_account_6s_act_man" +
		"(rpt_type,account_seq,account_item_code,account_code,account_name,segment_code1," +
		"segment_code2,segment_code3) VALUES ('" + row[18] + "'," +
		"'" + row[19] + "','" + row[20] + "','" + row[16] + "','" + row[17] + "','" + row[21] + "','" + row[22] + "','" + row[23] + "');\n"
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
