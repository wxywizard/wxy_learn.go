package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	cmd := &cobra.Command{
		Use:   "healthCheck",
		Short: "体脂计算器",
		Long:  "体脂计算器，并根据身高、体重、性别给出建议。。。。。",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name is {}", name)
			fmt.Println("sex is ", sex)
			fmt.Println("tall is ", tall)
			fmt.Println("weight is ", weight)
			fmt.Println("age is ", age)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	if err := cmd.Execute(); err == nil {
		return
	}

}
