package main

import (
	"fmt"
	"github.com/d3v-friends/go-grpc/fn/fnYaml"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path"
)

type protocYaml struct {
	ProtoPath string   `yaml:"protoPath"`
	OutPath   string   `yaml:"outPath"`
	PackageNm string   `yaml:"packageNm"`
	Include   []string `yaml:"include"`
	Import    []string `yaml:"import"`
}

func newProtocCmd() (res *cobra.Command) {
	res = &cobra.Command{
		Use: "protoc",
	}

	yamlPath := "protoc.yaml"
	yamlPathFlag := "config"

	res.Flags().StringVar(&yamlPath, yamlPathFlag, yamlPath, "--config protoc.yaml")

	res.Run = func(cmd *cobra.Command, args []string) {
		var err error
		if yamlPath, err = cmd.Flags().GetString(yamlPathFlag); err != nil {
			panic(err)
		}

		var pwd string
		if pwd, err = os.Getwd(); err != nil {
			panic(err)
		}
		yamlPath = path.Join(pwd, yamlPath)

		yaml := &protocYaml{}
		if err = fnYaml.Open(yamlPath, yaml); err != nil {
			panic(err)
		}

		outPath := path.Join(pwd, yaml.OutPath)
		if err = os.MkdirAll(outPath, os.ModePerm); err != nil {
			panic(err)
		}

		yaml.generate("go")
		yaml.generate("go-grpc")

		fmt.Printf("generated!\n")
	}

	return
}

func (x *protocYaml) generate(opt string) {
	var err error
	var pwd string
	if pwd, err = os.Getwd(); err != nil {
		panic(err)
	}

	outPath := path.Join(pwd, x.OutPath)
	protoPath := path.Join(pwd, x.ProtoPath)

	for _, protoNm := range x.Include {
		strCmd := make([]string, 0)
		strCmd = append(strCmd,
			fmt.Sprintf("--proto_path=%s", protoPath),
			fmt.Sprintf("--%s_out=%s", opt, outPath),
			fmt.Sprintf("--%s_opt=M%s=/%s", opt, protoNm, x.PackageNm),
		)

		importFileLs := make([]string, 0)
		for _, importProtoNm := range x.Import {
			strCmd = append(strCmd,
				fmt.Sprintf("--%s_opt=M%s=/%s", opt, importProtoNm, x.PackageNm),
			)

			importFileLs = append(importFileLs, importProtoNm)
		}

		strCmd = append(strCmd, protoNm)
		strCmd = append(strCmd, importFileLs...)

		cmd := exec.Command("protoc", strCmd...)
		log.Println(fmt.Sprintf("cmd: %s", cmd.String()))
		if err = cmd.Run(); err != nil {
			panic(err)
		}
	}
}
