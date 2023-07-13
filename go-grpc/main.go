package main

import "github.com/spf13/cobra"

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
