package headers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	zshSampleScriptFilename = `test_script.sh`

	// zshTemplate is a template used to maintain consistent
	// behavior across different posix compatible scripts.
	//
	// The replacement variables, ir the order that they
	// appear in the string, are:
	//
	//  header
	//  script setup
	//  script debug code
	//  script main() and exit code
	zshTemplate = `%s
%s
%s
%s
`

	// zshHeader is the format string for an executable file
	// header for posix based scripts run by zsh on macOS.
	//
	// It is likely that the most portable scripts will
	// operate fine when run by sh, bash, and most posix
	// based shells on macOS and likely any *nix.
	//
	// The replacement variables, ir the order that they
	// appear in the string, are:
	//
	//  name, description
	//  copyright information
	//  license information
	//  github repository url
	zshHeader = `#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034

#*-------------------------------> zsh utilities for macOS
    %s - %s

    %s %s

    %s
`
	zshScriptSetup = `
#*-------------------------------> script setup
    . $(which ansi_colors)          #* terminal colors and utils
    SET_DEBUG=${SET_DEBUG:-0}       #* set to 1 for verbose testing
    SCRIPT_NAME="${0##*/}"	        #* name of this script
    SCRIPT_PATH="${0%/*}"		    #* path of this script
    PID=$$						    #* PID of this process
`
	zshScriptDebug = `
#*-------------------------------> debug info
    _debug_tests() {
        if (( SET_DEBUG == 1 )); then
            warn "Debug Mode Details for ${CANARY}${SCRIPT_NAME}"
            green "Debug Mode: $SET_DEBUG"
        fi
        }
`
	zshMainSearchString = `#*-------------------------------> script main`
	zshMain             = `%s
    main() {
        ### main script begins here
        %s
    }

#*-------------------------------> main entry point
    _debug_tests
	main "$@"

`
	commentPrefix = `#*-------------------------------> `
)

// ZshMakeTemplate returns a complete posix compatible script.
// A template is used to maintain consistent behavior
//
// The replacement variables, ir the order that they
// appear in the string, are:
//
//  header
//  script setup
//  script debug code
//  script main() and exit code
func ZshMakeTemplate(file string) string {
	return fmt.Sprintf(zshTemplate, zshMakeHeader(), zshScriptSetup, zshScriptDebug, zshMakeMain(file))
}

// cutHeadertoSubstring removes sub and anything prior from s. If sub is not present in s, s is returned unchanged.
func cutHeadertoSubstring(s string, sub string) string {
	if i := strings.Index(s, sub); i > 0 {
		return s[:(i + len(sub))]
	}
	return s
}

// zshMakeMain removes all text before `zshMainSearchString` if
// it exists in the script, otherwise returns the entire script.
func zshMakeMain(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		if err == os.ErrNotExist {
			log.Fatalf("script file %s not found", file)
		}
		log.Fatalf("error opening file %s", file)
	}

	scriptBody := string(f)

	// todo - add option to leave out '_debug_tests' section
	return fmt.Sprintf(
		zshMain,
		zshMainSearchString,
		cutHeadertoSubstring(scriptBody, zshMainSearchString),
	)
}

// The replacement variables, ir the order that they
// appear in the string, are:
//
//  name, description
//  copyright information
//  license information
//  github repository url
func zshMakeHeader() string {
	fmt.Sprintf(zshHeader)
}

func SampleScript(zshSampleScriptFilename) {

}
