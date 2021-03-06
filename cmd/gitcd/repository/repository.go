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

package repository

import (
  "regexp"
    "errors"
  "path"
  "os"
)

type Repository struct {
  Owner string
  Name  string
}

/** Repository with a directory path. */
type ResolvedRepository struct {
  Repository Repository
  Directory  string
}

func (rr *ResolvedRepository) Exists() bool {
  _, err := os.Stat(rr.Directory)
  return !os.IsNotExist(err)
}

var PrefixProtocol = regexp.MustCompile(`(((git|ssh|http(s)?)://)?github\.com/)`)
var PrefixGit = regexp.MustCompile(`(git@github\.com:)`)
var Prefix = regexp.MustCompile(`(` + PrefixProtocol.String() + `|` + PrefixGit.String() + `)?`)
var RepositoryPart = regexp.MustCompile(`[\w-_]+`)
var RepositoryRegex = regexp.MustCompile(`^` + Prefix.String() + `(?P<owner>` + RepositoryPart.String() + `)/(?P<name>` + RepositoryPart.String() + `)(\.git)?$`)

/**
 * Canonicalizes the repositoryString into a Repository.
 *
 * For example:
 *   coollog/gitcd -> (Owner: coollog, Name: gitcd)
 *   github.com/coollog/gitcd -> (Owner: coollog, Name: gitcd)
 *   https://github.com/coollog/gitcd -> (Owner: coollog, Name: gitcd)
 */
func Canonicalize(repositoryString string) (Repository, error) {
  if !RepositoryRegex.MatchString(repositoryString) {
    return Repository{}, errors.New(`repository not valid`)
  }

  // Extracts the owner and name from repositoryString.
  repositoryMatches := matchNamedGroups(repositoryString)
  return Repository{repositoryMatches[`owner`], repositoryMatches[`name`]}, nil
}

/** Resolves the repo directory under the absoluteGitcdHome. */
func Resolve(gitcdHome string, repo Repository) ResolvedRepository {
  return ResolvedRepository{
    repo,
    path.Join(gitcdHome, repo.Owner, repo.Name),
  }
}

/**
 * Matches the named groups in RepositoryRegex and returns a map from the named groups to their matched values.
 */
func matchNamedGroups(repositoryString string) map[string]string {
  matchList := RepositoryRegex.FindStringSubmatch(repositoryString)

  matchMap := make(map[string]string)
  for i, name := range RepositoryRegex.SubexpNames() {
    if len(name) > 0 {
      matchMap[name] = matchList[i]
    }
  }

  return matchMap
}
