// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnmage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/target"
)

// EndToEnd namespace
type EndToEnd mg.Namespace

var (
	testDatabaseName = "ttn_lorawan_test"
	databaseURI      = "postgresql://root@localhost:26257/" + testDatabaseName + "?sslmode=disable"
)

func getTestURL() string {
	nodeEnv := os.Getenv("NODE_ENV")
	port := "1885"
	if nodeEnv == "development" {
		port = "8080"
	}
	return "http://localhost:" + port
}

func (endToEnd EndToEnd) cypress() (func(args ...string) error, error) {
	return execFromNodeBin("cypress", true, "--config-file", "./config/cypress.json", "--config", "baseUrl="+getTestURL())
}

func (endToEnd EndToEnd) prepareDB() error {
	isCI := os.Getenv("CI") == "true"
	dumpExists := pathExists("./.cache/sqldump.sql")
	changed, err := target.Dir("./.cache/sqldump.sql", "./pkg/identityserver/store")
	if err != nil {
		return err
	}
	if (changed && !isCI) || !dumpExists {
		os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
		mg.SerialDeps(Dev.DBErase, Dev.InitStack, EndToEnd.DBDump)
	} else {
		mg.SerialDeps(EndToEnd.DBRestore)
	}
	return nil
}

func (endToEnd EndToEnd) prepareBuild() {
	mg.SerialDeps(Js.Deps, Js.BuildIfNecessary)
}

// Prepare prepares the server for running end to end tests.
func (endToEnd EndToEnd) Prepare() {
	mg.Deps(endToEnd.prepareDB, endToEnd.prepareBuild)
}

// DBDump performs a database dump to the .cache folder.
func (EndToEnd) DBDump() error {
	if mg.Verbose() {
		fmt.Printf("Execute database dump\n")
	}
	_ = os.Mkdir(".cache", 0755)
	return execDockerComposeWithOutput(filepath.Join(".cache", "sqldump.sql"), "exec", "-T", "cockroach", "./cockroach", "dump", testDatabaseName, "--insecure")
}

// DBRestore restores the database using a previously generated dump.
func (EndToEnd) DBRestore() error {
	mg.Deps(Js.Deps, Dev.DBStart)
	if mg.Verbose() {
		fmt.Printf("Restore database from dump")
	}
	n, err := node()
	if err != nil {
		return err
	}
	return n("./.mage/restore-db-dump.js", "--db", testDatabaseName)
}

// Cypress runs the Cypress end-to-end tests.
func (endToEnd EndToEnd) Cypress() error {
	mg.Deps(endToEnd.WaitUntilReady)
	if mg.Verbose() {
		fmt.Println("Running Cypress E2E Tests")
	}
	cypress, err := endToEnd.cypress()
	if err != nil {
		return err
	}
	return cypress("run")
}

// CypressOpen runs the Cypress end-to-end tests in interactive mode.
func (endToEnd EndToEnd) CypressOpen() error {
	if mg.Verbose() {
		fmt.Println("Running Cypress in interactive mode")
	}
	cypress, err := endToEnd.cypress()
	if err != nil {
		return err
	}

	return cypress("open")
}

// StartTestStack starts TTS in end-to-end test configuration.
func (endToEnd EndToEnd) StartTestStack() error {
	mg.Deps(Js.BuildIfNecessary)
	if mg.Verbose() {
		fmt.Println("Starting stack in end-to-end test configuration")
	}
	os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
	_, err := outputGo("run", "./cmd/ttn-lw-stack", "start")
	return err
}

// WaitUntilReady waits until the web endpoints become available. For CI use.
func (endToEnd EndToEnd) WaitUntilReady() error {
	if mg.Verbose() {
		fmt.Println("Waiting for the stack to be ready...")
	}
	for i := 0; i < 100; i++ {
		resp, _ := http.Get(getTestURL() + "/oauth")
		if resp != nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return errors.New("Could not connect to server")
}

// RunFrontend starts the frontend end-to-end tests from scratch.
func (endToEnd EndToEnd) RunFrontend() {
	if mg.Verbose() {
		fmt.Println("Running end-to-end frontend based tests")
	}
	mg.Deps(EndToEnd.Cypress)
}

// Run starts the end-to-end tests from scratch.
func (endToEnd EndToEnd) Run() {
	if mg.Verbose() {
		fmt.Println("Running end-to-end tests")
	}
	mg.Deps(EndToEnd.RunFrontend)
}
