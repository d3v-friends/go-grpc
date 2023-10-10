package main

import (
	"fmt"
	"github.com/d3v-friends/go-grpc/fn/fnYaml"
	"github.com/d3v-friends/go-pure/fnFile"
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func main() {
	var err error
	cmd := &cobra.Command{
		Use: "go-grpc",
	}

	cmd.AddCommand(newProtocCmd())

	if err = cmd.Execute(); err != nil {
		panic(err)
	}
}

type protocYaml struct {
	ProtoPath fnFile.Path   `yaml:"protoPath"`
	OutPath   fnFile.Path   `yaml:"outPath"`
	PackageNm string        `yaml:"packageNm"`
	Include   []fnFile.Path `yaml:"include"`
	Import    []fnFile.Path `yaml:"import"`
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

		var fp = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
		fp.Join(yamlPath)
		yamlPath = fp.String()

		var yaml = &protocYaml{}
		if err = fnYaml.Open(yamlPath, yaml); err != nil {
			panic(err)
		}

		var outPath string
		if outPath, err = yaml.OutPath.Path(); err != nil {
			panic(err)
		}

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
	var protoPath = fnPanic.OnValue(x.ProtoPath.Path())
	var outPath = fnPanic.OnValue(x.OutPath.Path())

	for _, protoNm := range x.Include {
		strCmd := make([]string, 0)
		strCmd = append(strCmd,
			fmt.Sprintf("--proto_path=%s", protoPath),
			fmt.Sprintf("--%s_out=%s", opt, outPath),
			fmt.Sprintf("--%s_opt=M%s=.\\%s", opt, protoNm, x.PackageNm),
		)

		importFileLs := make([]string, 0)
		for _, importProtoNm := range x.Import {
			strCmd = append(strCmd,
				fmt.Sprintf("--%s_opt=M%s=.\\%s", opt, importProtoNm.String(), x.PackageNm),
			)

			importFileLs = append(importFileLs, importProtoNm.String())
		}

		strCmd = append(strCmd, protoNm.String())
		strCmd = append(strCmd, importFileLs...)

		cmd := exec.Command("protoc", strCmd...)
		log.Println(fmt.Sprintf("cmd: %s", cmd.String()))
		fnPanic.On(cmd.Run())
	}
}
