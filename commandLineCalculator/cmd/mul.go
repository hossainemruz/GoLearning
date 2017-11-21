// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"strconv"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "Multiply tow numbers",
	Long: `It takes input two integer as arguments then output result of multiplication of them`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<2{
			fmt.Println("Require two arguments")
		}else {
			a,err1:=strconv.Atoi(args[0])
			b,err2:=strconv.Atoi(args[1])
			if err1!=nil || err2!=nil{
				fmt.Println("Require two integer arguments")
			}else{
				fmt.Printf("%v * %v = %v\n",a,b,a*b)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(mulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
