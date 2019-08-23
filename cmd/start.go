/*
Copyright © 2019 Marcel Pohland <m.pohland@mwaysolutions.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		home = home + "/go/src/github.com/cosmos/gaia/"
		fmt.Println("YOUR HOME address: " + home)
		GitFetchCommand(home)
		GitCheckoutCommand(home)
		GitCheckoutVersionCommand(home)
		GitCheckoutCleanFDCommand(home)
		GitCheckoutCleanFXCommand(home)
		GoVersionCheck(home)
		MakeGoModCache(home)
		MakeInstall(home)
	},
}

func init() {

	//rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GitFetchCommand(home string) {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
func GitCheckoutCommand(home string) {
	cmd := exec.Command("git", "checkout", ".")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func GitCheckoutVersionCommand(home string) {
	cmd := exec.Command("git", "checkout", "v2.0.0")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func GitCheckoutCleanFDCommand(home string) {
	cmd := exec.Command("git", "clean", "-fd")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func GitCheckoutCleanFXCommand(home string) {
	cmd := exec.Command("git", "clean", "-fx")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func GoVersionCheck(home string) {
	cmd := exec.Command("go", "version")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GOVERSION: %q\n", out.String())
}

func MakeGoModCache(home string) {
	cmd := exec.Command("make", "go-mod-cache")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Log: %q\n", out.String())
}

func MakeInstall(home string) {
	cmd := exec.Command("make", "install")
	cmd.Dir = home
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Log: %q\n", out.String())
}