package cmd

import (
	"fmt"
	"github.com/nacos-cli/ncli/svc"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

var FGroup string
var FDataId string
var FApp string
var FDescription string
var FTags string
var FType string
var FContent string
var FCfgFile string

var cfgCmd = &cobra.Command{
	Use:   "config",
	Short: "manage config",
	Long:  "\n  Manage Nacos config.",
}

var cfgAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add config",
	Long:  "\n  Add Nacos config.",

	Run: func(cmd *cobra.Command, args []string) {

		parseArg()

		login, err := svc.Login(GServer, GLogin)
		if err != nil {
			_, _ = os.Stderr.WriteString("Login failed, " + err.Error())
			return
		}

		cfg := new(svc.ConfigPost)
		cfg.NamespaceId = FNamespaceId
		cfg.Group = FGroup
		cfg.DataId = FDataId
		cfg.Content = FContent
		cfg.Type = FType
		cfg.App = FApp
		cfg.Description = FDescription
		cfg.Tags = FTags

		if len(FCfgFile) > 0 {
			cont, ext, err := readConfigFile(FCfgFile)
			if err != nil {
				_, _ = os.Stderr.WriteString("Read config from file failed, " + err.Error())
				return
			}
			if cfg.Type == "" && len(ext) > 0 {
				cfg.Type = ext
			}
			cfg.Content = cont
		}

		_, err = svc.ConfigCreateOrUpdate(login, cfg)
		if err != nil {
			_, _ = os.Stderr.WriteString("Create config/update failed, " + err.Error())
			return
		}
		fmt.Printf("Config '%s' created/updated\n", cfg.DataId)
	},
}

var cfgGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get config",
	Long:  "\n  Get Nacos config.",

	Run: func(cmd *cobra.Command, args []string) {

		parseArg()

		login, err := svc.Login(GServer, GLogin)
		if err != nil {
			_, _ = os.Stderr.WriteString("Login failed, " + err.Error())
			return
		}

		s := new(svc.ConfigQuery)
		s.NamespaceId = FNamespaceId
		s.Group = FGroup
		s.DataId = FDataId
		cfg, err := svc.ConfigGet(login, s)
		if err != nil {
			_, _ = os.Stderr.WriteString("Get config failed, " + err.Error())
			return
		}
		println(cfg.Content)
	},
}

func init() {
	parseServerFlag(cfgCmd)
	parseLoginFlag(cfgCmd)
	parseNamespaceFlag(cfgCmd)

	parseCfg(cfgCmd)

	parseOpCfg(cfgAddCmd)

	cfgCmd.AddCommand(cfgAddCmd)
	cfgCmd.AddCommand(cfgGetCmd)

	rootCmd.AddCommand(cfgCmd)

}

func readConfigFile(file string) (cont string, ext string, err error) {
	empty := ""
	extension := path.Ext(file)
	if len(extension) > 0 {
		ext = strings.TrimLeft(extension, ".")
	} else {
		ext = empty
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return empty, ext, err
	}
	return string(bytes), ext, nil
}

func parseCfg(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&FGroup, "group", "g", "DEFAULT_GROUP", "config group")
	cmd.PersistentFlags().StringVarP(&FDataId, "dataId", "d", "", "config data id")

}

func parseOpCfg(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&FType, "type", "t", "", "config type, optional if 'from' is specified")
	cmd.PersistentFlags().StringVarP(&FContent, "content", "c", "", "config content, optional if 'from' is specified")
	cmd.PersistentFlags().StringVarP(&FCfgFile, "from", "f", "", "post config from file, automatically override config content/type")
	cmd.PersistentFlags().StringVarP(&FApp, "app", "a", "", "optional, config application")
	cmd.PersistentFlags().StringVarP(&FDescription, "desc", "D", "", "optional, config description (optional)")
	cmd.PersistentFlags().StringVarP(&FTags, "tags", "T", "", "optional, comma-delimited config tags")

}
