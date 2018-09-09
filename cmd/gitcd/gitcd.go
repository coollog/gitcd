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
      "strings"
  "github.com/coollog/gitcd/cmd/gitcd/cache"
  "github.com/coollog/gitcd/cmd/gitcd/home"
  "io/ioutil"
  "path"
)

const Usage = `Quickly navigate to your GitHub repositories.

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
    err := showClonedRepositories()
    if err != nil {
      log.Fatal(err)
    }
  }
  os.Exit(1)
}

/** Prints the repo directory matching the repositoryString query. Returns true if successful; false if not. */
func gitcd(repositoryString string) bool {
  // Gets the gitcd home directory.
  gitcdHome, err := home.GitcdHome()
  if err != nil {
    log.Fatal(err)
    return false
  }

  // If the respository string is just one part, then try to guess the full repository.
  if !strings.ContainsAny(repositoryString, `/`) {
    repoName := repositoryString

    // Loads the .gitcd file.
    gitcdFile, err := home.GitcdFile()
    if err != nil {
      log.Fatal(err)
      return false
    }
    repoCache, err := cache.Load(gitcdFile)
    if err != nil {
      log.Fatal(err)
      return false
    }

    // Tries to find owners for repoName.
    owners := repoCache.FindOwners(repoName)
    for _, owner := range owners {
      resolvedRepository := repository.Resolve(gitcdHome, repository.Repository{Owner: owner, Name: repoName})
      if !resolvedRepository.Exists() {
        continue
      }
      fmt.Println(resolvedRepository.Directory)
      return true
    }

    log.Printf("No known matching repositories with name `%s`", repoName)
    showClonedRepositories()
    return false
  }

  // Parses the repository string into a canonicalized form.
  canonicalRepository, err := repository.Canonicalize(repositoryString)
  if err != nil {
    log.Fatal(err)
    return false
  }

  // Checks if the repository exists.
  resolvedRepository := repository.Resolve(gitcdHome, canonicalRepository)
  if !resolvedRepository.Exists() {
    // Repository doesn't exist, clone it.
    err := repository.Clone(gitcdHome, repositoryString, canonicalRepository)
    if err != nil {
      log.Fatalf("Could not clone repository `%s`:%s", repositoryString, err.Error())
      return false
    }
  }

  // Bumps the repo to the top in the .gitcd file.
  gitcdFile, err := home.GitcdFile()
  if err != nil {
    log.Printf("Could not resolve .gitcd file: %s\n", err.Error())

  } else {
    // Loads the .gitcd file.
    repoCache, err := cache.Load(gitcdFile)
    if err != nil {
      log.Printf("Could not load .gitcd file: %s\n", err.Error())
    } else {
      // Saves the .gitcd file with the repo bumped to top.
      repoCache.Bump(resolvedRepository.Repository)
      err := cache.Save(gitcdFile, repoCache)
      if err != nil {
        log.Printf("Could not save .gitcd file: %s\n", err.Error())
      }
    }
  }

  // Prints the repo directory.
  fmt.Println(resolvedRepository.Directory)
  return true
}

func showUsage() {
  fmt.Println(Usage)
}

/** Shows all the cloned repos. */
func showClonedRepositories() error {
  gitcdHome, err := home.GitcdHome()
  if err != nil {
    return err
  }

  if _, err := os.Stat(gitcdHome); os.IsNotExist(err) {
    return nil
  }

  var clonedRepos []repository.Repository
  // Lists all the owner directories.
  fileInfos, err := ioutil.ReadDir(gitcdHome)
  if err != nil {
    return err
  }
  for _, fileInfo := range fileInfos {
    if fileInfo.Mode().IsDir() {
      repoOwner := fileInfo.Name()

      // Lists all the owner/name directories.
      fileInfos, err := ioutil.ReadDir(path.Join(gitcdHome, fileInfo.Name()))
      if err != nil {
        return err
      }
      for _, fileInfo := range fileInfos {
        if fileInfo.Mode().IsDir() {
          repoName := fileInfo.Name()
          clonedRepos = append(clonedRepos, repository.Repository{Owner: repoOwner, Name: repoName})
        }
      }
    }
  }

  if len(clonedRepos) > 0 {
    fmt.Println()
    fmt.Println("Cloned repositories:")
    for _, repo := range clonedRepos {
      fmt.Printf("\t%s/%s\n", repo.Owner, repo.Name)
    }
  }

  return nil
}
