package cmd



import (
	"fmt"
	"github.com/spf13/viper"
	"go-web/internal/api/employee"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool

var addCmd = &cobra.Command{
	Use: "add",
	Short: "添加员工",
	Run: func(cmd *cobra.Command, args []string) {
		employee.New().List(args)
		if true {
			fmt.Println("About to greet friends from Earth...")
		}
	},
}

var deleteCmd = &cobra.Command{
	Use: "del",
	Short: "删除员工",
	Run: func(cmd *cobra.Command, args []string) {
		employee.New().List(args)
		if true {
			fmt.Println("About to greet friends from Earth...")
		}
	},
}


var searchCmd = &cobra.Command{
	Use: "search",
	Short: "按照id搜索员工",
	Run: func(cmd *cobra.Command, args []string) {
		employee.New().List(args)
		if true {
			fmt.Println("About to greet friends from Earth...")
		}
	},
}


var modifyCmd = &cobra.Command{
	Use: "modify",
	Short: "按照id修改员工信息",
	Run: func(cmd *cobra.Command, args []string) {
		employee.New().List(args)
		if true {
			fmt.Println("About to greet friends from Earth...")
		}
	},
}


var listCmd = &cobra.Command{
	Use: "list",
	Short: "查看员工列表",
	Run: func(cmd *cobra.Command, args []string) {
		employee.New().List(args)
		if true {
			fmt.Println("About to greet friends from Earth...")
		}
	},
}



func init() {

	//先初始化控制器

	addCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")

	addCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", addCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", addCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	addCmd.AddCommand(deleteCmd)
	addCmd.AddCommand(searchCmd)
	addCmd.AddCommand(modifyCmd)
	addCmd.AddCommand(listCmd)
	addCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	addCmd.PersistentFlags().String("lang", "en", "language to use")


}


func Execute() {
	if err := addCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}