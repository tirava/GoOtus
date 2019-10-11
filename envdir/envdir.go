/*
 * HomeWork-7: envdir utility like envdir
 * Created on 11.10.2019 21:51
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
)

// EnvDirExec runs program with env from given directory
func EnvDirExec(pathProgram, pathEnvDir string) error {

	cmd := exec.Command(pathProgram)

	// replace system env with files env
	ef, err := getEnvFromFiles(pathEnvDir)
	if err != nil {
		return err
	}
	cmd.Env = replaceSystemEnvOnFilesEnv(os.Environ(), ef)

	// run and print env
	out, err := cmd.Output()
	if err != nil {
		return err
	} else {
		fmt.Printf("%s", out)
	}

	return nil
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
		env, err := getLineFromFile(path.Join(envDir, file.Name()))
		if err != nil {
			return nil, err
		}
		envs = append(envs, file.Name()+"="+env)
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

func replaceSystemEnvOnFilesEnv(sysEnv, filesEnv []string) []string {
	if !inheritEnv {
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

	// hash -> slice
	for key, val := range hash {
		env := fmt.Sprintf("%s=%s", key, val)
		inter = append(inter, env)
	}

	sort.Strings(inter)

	return inter
}
