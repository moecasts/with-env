package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func LoadEnv(files ...string) {
	for _, envfile := range files {
		absPath, _ := GetAbsPath(envfile)
		if _, err := os.Stat(absPath); !os.IsNotExist(err) {
			_ = godotenv.Overload(absPath)
		}
	}
}

func GetAbsPath(path string) (string, error) {
	relativePath := path

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("get home dir failed: %v", home)
			return "", err
		}

		relativePath = strings.Replace(path, "~", home, 1)
	}

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("get abs path failed: %v", err)
		return "", err
	}

	return absPath, nil
}

func WithEnvAction(ctx *cli.Context) error {
	LoadEnv(ctx.StringSlice("env")...)

	c := exec.Command(ctx.Args().First(), ctx.Args().Tail()...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%q\n", err)
		os.Exit(1)
	}

	return nil
}
