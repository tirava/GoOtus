/*
 * HomeWork-7: envdir utility like envdir
 * Created on 11.10.2019 21:51
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// EnvDirExec runs program with env from given directory
func EnvDirExec(out io.Writer, pathEnvDir, pathProgram string, inheritSysEnv bool) error {

	cmd := exec.Command(pathProgram)
	cmd.Stdout = out

	// replace system env with files env
	ef, err := getEnvFromFiles(pathEnvDir)
	if err != nil {
		return err
	}

	cmd.Env = replaceSystemEnvOnFilesEnv(os.Environ(), ef, inheritSysEnv)

	return cmd.Run()
}

func getEnvFromFiles(envDir string) ([]string, error) {

	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		return nil, err
	}

	envs := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		env, err := getLineFromFile(filepath.Join(envDir, file.Name()))
		if err != nil {
			return nil, err
		}
		envs = append(envs, fmt.Sprintf("%s=%s", file.Name(), env))
	}

	return envs, nil
}

func getLineFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text(), nil
}

func replaceSystemEnvOnFilesEnv(sysEnv, filesEnv []string, inheritSysEnv bool) []string {

	// return only dir envs
	if !inheritSysEnv {
		return filesEnv
	}

	hash := make(map[string]string)
	inter := make([]string, 0)

	// hash sys env
	for _, se := range sysEnv {
		env := strings.SplitN(se, "=", 2)
		hash[env[0]] = env[1]
	}

	// hash dir env
	for _, fe := range filesEnv {
		env := strings.SplitN(fe, "=", 2)
		hash[env[0]] = env[1]
	}

	// hash intercept -> slice
	for key, val := range hash {
		env := fmt.Sprintf("%s=%s", key, val)
		inter = append(inter, env)
	}

	sort.Strings(inter) // for best view

	return inter
}
