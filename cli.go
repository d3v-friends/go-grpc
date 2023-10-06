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
	"runtime"
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

		var fp = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
		fp.Join(yamlPath)
		yamlPath = fp.String()

		var yaml = &protocYaml{}
		if err = fnYaml.Open(yamlPath, yaml); err != nil {
			panic(err)
		}

		fp = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
		fp.Join(yaml.OutPath)

		var outPath = fp.String()
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
	var outPathBuilder = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
	outPathBuilder.Join(x.OutPath)
	var outPath = outPathBuilder.String()

	var protoPathBuilder = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
	protoPathBuilder.Join(x.ProtoPath)
	var protoPath = protoPathBuilder.String()

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

		if runtime.GOOS == "windows" {
			strCmd = append(strCmd, "--experimental_allow_proto3_optional")
		}

		strCmd = append(strCmd, protoNm)
		strCmd = append(strCmd, importFileLs...)

		cmd := exec.Command("protoc", strCmd...)
		log.Println(fmt.Sprintf("cmd: %s", cmd.String()))
		fnPanic.On(cmd.Run())
	}
}
