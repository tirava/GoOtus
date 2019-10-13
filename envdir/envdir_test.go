/*
 * HomeWork-7: envdir utility like envdir
 * Created on 12.10.2019 12:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
)

// short name of the test program to run with params
const EXECNAME = "envdir"

var testCasesErrors = []struct {
	envDir      string
	envVars     []string
	fileAttr    os.FileMode
	inherit     bool
	description string
}{
	{
		"fakeDir",
		[]string{},
		0,
		false,
		"read fake directory",
	},
	{
		"realDir",
		[]string{"222=333"},
		0200,
		false,
		"real directory but fail read envs (*nix OS only)",
	},
}

var testCases = []struct {
	envDir       string
	envVars      []string
	expectedVars []string
	inherit      bool
	description  string
}{
	{
		"TestDir",
		[]string{"QQQ=AAA", "WWW=SSS"},
		[]string{"QQQ=AAA", "WWW=SSS"},
		false,
		"no inherit 1",
	},
	{
		"TestDir1",
		[]string{"EEE=DDD", "RRR=FFF", "ttt=ggg"},
		[]string{"EEE=DDD", "RRR=FFF", "ttt=ggg"},
		false,
		"no inherit 2",
	},
	{
		"TestDir2",
		[]string{"ZZZ=zzz\nyyy\nxxx"},
		[]string{"ZZZ=zzz"},
		false,
		"no inherit 3, multiline",
	},
	{
		"TestDir3",
		[]string{"111=222", "YYY=yyy"},
		[]string{"111=222", "YYY=yyy"},
		true,
		"inherit system env",
	},
	{
		"emptyDir",
		[]string{"EEE"},
		[]string{},
		false,
		"empty dir instead env file",
	},
}

func TestEnvDirExecErrors(t *testing.T) {

	execFile := getExecFile()

	for _, test := range testCasesErrors {

		// skip *nix specific tests
		if test.fileAttr != 0 && runtime.GOOS == "windows" {
			t.Logf("SKIPPED TestEnvDirExecErrors - %s", test.description)
			continue
		}

		cleanEnvDir(test.envDir)
		generateEnvDir(test.envDir, test.envVars, test.fileAttr)

		out := new(strings.Builder)
		err := EnvDirExec(out, test.envDir, execFile, test.inherit)
		if err == nil {
			t.Errorf("FAIL '%s' - TestEnvDirExecErrors() returns 'no error', expected error.",
				test.description)
			continue
		}
		t.Logf("PASS TestEnvDirExecErrors - %s", test.description)

		// make clean if not need results
		cleanEnvDir(test.envDir)
	}
}

func TestEnvDirExec(t *testing.T) {

	execFile := getExecFile()

	for _, test := range testCases {

		cleanEnvDir(test.envDir)
		generateEnvDir(test.envDir, test.envVars)

		out := new(strings.Builder)
		err := EnvDirExec(out, test.envDir, execFile, test.inherit)
		if err != nil {
			t.Fatalf("FAIL '%s' - TestEnvDirExec() returns error\n %s\nexpected no error.",
				test.description, err)
		}

		result := strings.Split(out.String(), "\n")
		result = result[:len(result)-1] // delete last '\n'
		if test.inherit {
			test.expectedVars = append(test.expectedVars, os.Environ()...)
			sort.Strings(test.expectedVars)
		} else {
			if runtime.GOOS == "windows" && len(result) > 0 {
				if strings.HasPrefix(result[len(result)-1], "SYSTEMROOT=") {
					result = result[:len(result)-1] // delete nested SYSTEMROOT in Windows 7
				}
			}
		}
		if !reflect.DeepEqual(result, test.expectedVars) {
			t.Errorf("FAIL '%s' - TestEnvDirExec() - result:\n%s\nexpected:\n%s\n",
				test.description, result, test.expectedVars)
			continue
		}

		t.Logf("PASS TestEnvDirExec - %s", test.description)

		// make clean if not need results
		cleanEnvDir(test.envDir)
	}
}

func getExecFile() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Can't get current test directory!", err)
	}
	return filepath.Join(dir, EXECNAME)
}

func generateEnvDir(envDir string, envVars []string, fileAttr ...os.FileMode) {
	if envDir == "fakeDir" {
		return
	}
	err := os.Mkdir(envDir, 0777)
	if err != nil {
		log.Fatalln("Can't create test directory!", err)
	}
	for _, ev := range envVars {
		fileName := strings.SplitN(ev, "=", 2)
		if len(fileName) == 1 { // nested directory
			err := os.MkdirAll(filepath.Join(envDir, fileName[0]), 0777)
			if err != nil {
				log.Fatalln("Can't create empty directory!", err)
			}
			continue
		}
		file, err := os.Create(filepath.Join(envDir, fileName[0]))
		if err != nil {
			log.Fatalln("Can't create test file!", err)
		}
		_, err = file.Write([]byte(fileName[1]))
		if err != nil {
			log.Fatalln("Can't write test data to file!", err)
		}
		err = file.Close()
		if err != nil {
			log.Fatalln("Can't close test file!", err)
		}
		if len(fileAttr) > 0 {
			err := os.Chmod(filepath.Join(envDir, fileName[0]), fileAttr[0])
			if err != nil {
				log.Fatalln("Can't set file mode!", err)
			}
		}
	}
}

func cleanEnvDir(envDir string) {
	err := os.RemoveAll(envDir)
	if err != nil {
		log.Fatalln("Can't delete test directory!", err)
	}
}
