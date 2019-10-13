/*
 * HomeWork-7: envdir utility like envdir
 * Created on 13.10.2019 12:55
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import "os"

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
