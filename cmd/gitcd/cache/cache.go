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

package cache

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "github.com/coollog/gitcd/cmd/gitcd/repository"
  "os"
  "errors"
  "fmt"
)

// The cache stores the usages of certain repositories in order to find repositories by a shorter name.
// For example, using `coollog/gitcd` many times would mean that `gitcd` would resolve to `coollog/gitcd`.

/**
 * The YAML structure for the cache file.
 *
 * `apiVersion` is current 1.
 * `nameMap` maps from repo name to list of owners, in order of last access.
 *
 * Example:
 *
 * apiVersion: 1
 * nameMap:
 *   gitcd:
 *   - coollog
 *   bar:
 *   - foo
 *   - cat
 */
type RepoCache struct {
  ApiVersion int
  NameMap    map[string][]string
}

/** Bumps the repo to the top. */
func (r *RepoCache) Bump(repoToBump repository.Repository) {
  // Starts a new list with the repos.
  var newOwnerList []string
  newOwnerList = append(newOwnerList, repoToBump.Owner)

  // Collects current repos to map.
  if owners, ok := r.NameMap[repoToBump.Name]; ok {
    ownerMap := make(map[string]bool)
    ownerMap[repoToBump.Owner] = true

    for _, owner := range owners {
      if _, ok := ownerMap[owner]; !ok {
        newOwnerList = append(newOwnerList, owner)
        ownerMap[owner] = true
      }
    }
  }

  r.NameMap[repoToBump.Name] = newOwnerList
}

/** Gets the list of owners to try for the repoName, in the order in which to try them. */
func (r *RepoCache) FindOwners(repoName string) []string {
  if owners, ok := r.NameMap[repoName]; ok {
    return owners
  }
  return []string{}
}

/** Loads the gitcdFile into the RepoCache structure. */
func Load(gitcdFile string) (RepoCache, error) {
  if _, err := os.Stat(gitcdFile); os.IsNotExist(err) {
    return RepoCache{
      ApiVersion: 1,
      NameMap:    make(map[string][]string),
    }, nil
  }

  gitcdFileContents, err := ioutil.ReadFile(gitcdFile)
  if err != nil {
    return RepoCache{}, err
  }

  repoCache := RepoCache{}

  err = yaml.Unmarshal(gitcdFileContents, &repoCache)
  if err != nil {
    return RepoCache{}, err
  }

  if repoCache.ApiVersion != 1 {
    return RepoCache{}, errors.New(fmt.Sprintf(".gitcd file at `%s` has unknown apiVersion: %s", gitcdFile, repoCache.ApiVersion))
  }
  if repoCache.NameMap == nil {
    return RepoCache{}, errors.New(fmt.Sprintf(".gitcd file at `%s` has nil nameMap", gitcdFile))
  }

  return repoCache, nil
}

/** Saves the RepoCache into the gitcdFile. */
func Save(gitcdFile string, repoCache RepoCache) error {
  gitcdFileContents, err := yaml.Marshal(&repoCache)
  if err != nil {
    return err
  }

  return ioutil.WriteFile(gitcdFile, gitcdFileContents, 0644)
}
