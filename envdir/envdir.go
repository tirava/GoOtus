/*
 * HomeWork-7: envdir utility like envdir
 * Created on 11.10.2019 21:51
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"log"
	"os/exec"
)

// EnvDirExec runs program with env from given directory
func EnvDirExec(pathProgram, pathEnvDir string) error {

	cmd := exec.Command(pathProgram)

	// get env from files in dir
	cmd.Env = []string{"QQQ=qqq", "VVV=vvv"}

	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Get output:\n%s", out)
	}

	return nil
}
