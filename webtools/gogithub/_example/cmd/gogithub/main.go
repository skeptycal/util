// Copyright 2015 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The basicauth command demonstrates using the github.BasicAuthTransport,
// including handling two-factor authentication. This won't currently work for
// accounts that use SMS to receive one-time passwords.
//
// Deprecation Notice: GitHub will discontinue password authentication to the API.
// You must now authenticate to the GitHub API with an API token, such as an OAuth access token,
// GitHub App installation access token, or personal access token, depending on what you need to do with the token.
// Password authentication to the API will be removed on November 13, 2020.
// See the tokenauth example for details.
package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/google/go-github/v32/github"
	gh "github.com/skeptycal/gogithub"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.Info("GitHub client example:\n")

	ctx := gh.DefaultContext

	ts, err := gh.GetEnvToken()
	if err == nil {
		log.Info(ts)
	}

	client, err := gh.NewGitHubClient(ctx)
	if err != nil {
		fmt.Printf("client creation failed %v", err)
		return
	}

	user, _, err := client.Users.Get(ctx, "")

	fmt.Printf("\n%v\n", github.Stringify(user))
}

func sampleLogOutput() {
	log.Info("Some info. Earth is not flat.")
	log.Warning("This is a warning")
	log.Error("Not fatal. An error. Won't stop execution")
	log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	log.Panic("Do not panic")
}
