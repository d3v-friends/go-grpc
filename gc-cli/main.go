package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/d3v-friends/go-tools/fnError"
	"github.com/d3v-friends/go-tools/fnPanic"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	var cmd = &cobra.Command{
		Use: "go-grpc",
	}

	var configPath = "grpc.yml"
	var configPathFlag = "config"
	cmd.Flags().StringVar(&configPath, configPathFlag, configPath, "--config grpc.yml")

	cmd.Run = func(c *cobra.Command, args []string) {
		// 01. load file
		configPath = fnPanic.Value(c.Flags().GetString(configPathFlag))
		configPath = parseLocalAbsolute(configPath)

		var config = &ConfigFormat{}
		var ext = filepath.Ext(configPath)
		switch ext {
		case ".yaml", ".yml":
			fnPanic.On(parseYAML(configPath, config))
		case ".json":
			fnPanic.On(parseJSON(configPath, config))
		default:
			log.Printf("not support config file type: %s", ext)
			return
		}

		// 02. generate proto file

		fnPanic.On(config.generate(ProtocOptGo))
		fnPanic.On(config.generate(ProtocOptGoGrpc))
	}

	fnPanic.On(cmd.Execute())
}

type ConfigFormat struct {
	Proto   string   `yaml:"proto" json:"proto"`
	Out     string   `yaml:"out" json:"out"`
	Package string   `yaml:"package" json:"package"`
	Include []string `yaml:"include" json:"include"`
	Import  []string `yaml:"import" json:"import"`
}

type ProtocOpt string

func (x ProtocOpt) String() string {
	return string(x)
}

const (
	ProtocOptGo     ProtocOpt = "go"
	ProtocOptGoGrpc ProtocOpt = "go-grpc"
)

func (x ConfigFormat) generate(opt ProtocOpt) (err error) {
	var rootPath = parseLocalAbsolute(x.Proto)
	var outPath = parseLocalAbsolute(x.Out)

	for _, protoNm := range x.Include {
		var strCmd = make([]string, 0)
		strCmd = append(strCmd,
			fmt.Sprintf("--proto_path=%s", rootPath),
			fmt.Sprintf("--%s_out=%s", opt.String(), outPath),
			fmt.Sprintf("--%s_opt=M%s=/%s", opt.String(), protoNm, x.Package),
		)

		var importFiles = make([]string, 0)
		for _, importProtoNm := range x.Import {
			strCmd = append(
				strCmd,
				fmt.Sprintf("--%s_opt=M%s=/%s", opt, importProtoNm, x.Package),
			)

			importFiles = append(importFiles, importProtoNm)
		}

		strCmd = append(
			strCmd,
			protoNm,
		)

		strCmd = append(
			strCmd,
			importFiles...,
		)

		var cmd = exec.Command("protoc", strCmd...)
		log.Println(fmt.Sprintf("cmd: %s", cmd.String()))
		fnPanic.On(cmd.Run())
	}

	return
}

func parseLocalAbsolute(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	var wd = fnPanic.Value(os.Getwd())
	return filepath.Join(wd, path)
}

func parseYAML(fp string, model any) (err error) {
	var file *os.File
	if file, err = readFile(fp); err != nil {
		return
	}

	defer func() {
		fnPanic.On(file.Close())
	}()

	if err = yaml.
		NewDecoder(file).
		Decode(model); err != nil {
		return
	}

	return
}

func readFile(fp string) (file *os.File, err error) {
	if !hasFile(fp) {
		err = fnError.NewF("not found proto file")
		return
	}

	return os.Open(fp)
}

func hasFile(fp string) bool {
	var stat, err = os.Stat(fp)
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}

func parseJSON(fp string, model any) (err error) {
	var file *os.File
	if file, err = readFile(fp); err != nil {
		return
	}

	defer func() {
		fnPanic.On(file.Close())
	}()

	if err = json.
		NewDecoder(file).
		Decode(model); err != nil {
		return
	}

	return
}
