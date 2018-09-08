/*
 * Copyright 2018 Google LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
  "os"
  "fmt"
  "github.com/coollog/gitcd/cmd/gitcd/repository"
  "log"
  "path"
  "path/filepath"
)

const USAGE = `Quickly navigate to your GitHub repositories.

Usage:
  1) Add this function to your bash profile (~/.bashrc or ~/.bash_profile):
    
    gcd() { gitcd "$@" && cd ` + "`" + `gitcd "$@"` + "`" + `; }

  2) gcd [repository] - goes to the directory for that repository

Repositories live under $GITCD_HOME.

If the repository does not exist, clones the repository.
`

func main() {
  switch len(os.Args) {
  case 2:
    repositoryString := os.Args[1]
  if gitcd(repositoryString) {
    os.Exit(0)
    }

  default:
  showUsage()
  }
  os.Exit(1)
}

func gitcd(repositoryString string) bool {
  // Gets the gitcd home directory.
  gitcdHome := getGitcdHome()

  // Parses the repository string into a canonicalized form.
  canonicalRepository, err := repository.Canonicalize(repositoryString)
  if err != nil {
  log.Fatal(err.Error())
  return false
  }

  // Checks if the repository exists.
  repositoryDirectory := path.Join(gitcdHome, canonicalRepository.Owner, canonicalRepository.Name)
  if _, err := os.Stat(repositoryDirectory); os.IsNotExist(err) {
    // Repository doesn't exist, clone it.
    log.Fatal(`Cannot clone repository - unimplemented`)
    return false
  }

  absoluteRepositoryDirectory, err := filepath.Abs(repositoryDirectory)
  if err != nil {
    log.Fatalf("Cannot resolve directory `%s`: %s", repositoryDirectory, err.Error())
    return false
  }

  fmt.Println(absoluteRepositoryDirectory)
  return true
}

func showUsage() {
  fmt.Println(USAGE)
}
