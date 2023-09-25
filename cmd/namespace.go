package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nacos-cli/ncli/svc"
	"github.com/spf13/cobra"
)

var nsCmd = &cobra.Command{
	Use:   "namespace",
	Short: "manage Nacos namespace",
	Long:  `add/update Nacos namespace.`,
}

var FNamespaceName string
var FNamespaceDesc string

var nsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add Nacos namespace",
	Long:  `add Nacos namespace.`,

	Run: func(cmd *cobra.Command, args []string) {

		parseArg()

		login, err := svc.Login(GServer, GLogin)
		if err != nil {
			os.Stderr.WriteString("Login failed, " + err.Error())
			return
		}

		_, err = svc.Create(login, FNamespaceId, FNamespaceName, FNamespaceDesc)
		if err != nil {
			os.Stderr.WriteString("Create namespace failed, " + err.Error())
			return
		}
		fmt.Printf("Namespace '%s' created/exists\n", FNamespaceId)

	},
}

var nsExistCmd = &cobra.Command{
	Use:   "exist",
	Short: "check if Nacos namespace exist",
	Long:  `check Nacos namespace.`,

	Run: func(cmd *cobra.Command, args []string) {

		parseArg()

		login, err := svc.Login(GServer, GLogin)
		if err != nil {
			os.Stderr.WriteString("Login failed, " + err.Error())
			return
		}

		exist, err := svc.Exist(login, FNamespaceId)
		if err != nil {
			os.Stderr.WriteString("Check namespace failed, " + err.Error())
			return
		}
		fmt.Printf("%s\n", strconv.FormatBool(exist))

	},
}

func init() {

	parseServerFlag(nsCmd)
	parseLoginFlag(nsCmd)
	parseNamespaceFlag(nsCmd)

	nsAddCmd.PersistentFlags().StringVarP(&FNamespaceName, "name", "N", "", "namespace name")
	nsAddCmd.PersistentFlags().StringVarP(&FNamespaceDesc, "desc", "D", "", "namespace description")

	nsCmd.AddCommand(nsAddCmd)
	nsCmd.AddCommand(nsExistCmd)

	rootCmd.AddCommand(nsCmd)

}
