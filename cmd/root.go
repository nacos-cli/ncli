package cmd

import (
	"os"

	"github.com/nacos-cli/ncli/svc"
	"github.com/spf13/cobra"
)

const (
	AppName    = "ncli"
	AppVersion = "0.1.0"
)

var GServer = new(svc.Server)
var GLogin = new(svc.LoginPost)
var GLoginInfo = new(svc.LoginResult)

var FNamespaceId string

var FIniFile string

var FSchema string
var FHost string
var FPort uint16
var FContext string
var FUsername string
var FPassword string

var rootCmd = &cobra.Command{
	Use:   "ncli",
	Short: "ncli is a nacos cli client",
	Long:  `ncli is a nacos cli client`,
}

func parseServerFlag(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&FSchema, "schema", "s", "http", "nacos server schema")

	cmd.PersistentFlags().StringVarP(&FHost, "host", "H", "127.0.0.1", "nacos server ip")

	cmd.PersistentFlags().Uint16VarP(&FPort, "port", "P", 8848, "nacos server port")

	cmd.PersistentFlags().StringVarP(&FContext, "context", "c", "/nacos", "nacos server context path")

}

func parseLoginFlag(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&FUsername, "username", "u", "nacos", "nacos server auth username")
	cmd.PersistentFlags().StringVarP(&FPassword, "password", "p", "nacos", "nacos server auth password")

}

func parseNamespaceFlag(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&FNamespaceId, "namespaceId", "n", "public", "nacos namespace id")

}

func parseArg() {

	GServer.Schema = FSchema
	GServer.Host = FHost
	GServer.Port = FPort
	GServer.Context = FContext
	GLogin.Username = FUsername
	GLogin.Password = FPassword

}

func RootExecute() {

	err := rootCmd.Execute()
	if err != nil {
		_, _ = os.Stderr.WriteString("Failed, caused by:" + err.Error())
		os.Exit(1)
	}

}
