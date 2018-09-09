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

package home

import (
  "os"
  "github.com/mitchellh/go-homedir"
  "path"
  "path/filepath"
    "errors"
  "fmt"
)

/** Environment variable defining the home directory for gitcd. */
const GitcdHomeEnvvar = `GITCD_HOME`

/** Name of the .gitcd file. */
const GitcdFilename = `.gitcd`

/** Gets the gitcd home directory. */
func GitcdHome() (string, error) {
  gitcdHome := os.Getenv(GitcdHomeEnvvar)
  if len(gitcdHome) > 0 {
    return absolute(gitcdHome)
  }
  return defaultGitcdHome()
}
func defaultGitcdHome() (string, error) {
  userHome, err := homedir.Dir()
  if err != nil {
    return ``, err
  }
  return path.Join(userHome, `gitcd`), nil
}

/** Gets the .gitcd file. */
func GitcdFile() (string, error) {
  gitcdHome, err := GitcdHome()
  if err != nil {
    return ``, err
  }
  return path.Join(gitcdHome, GitcdFilename), nil
}

/** Gets the absolute gitcd home. */
func absolute(gitcdHome string) (string, error){
  absoluteGitcdHome, err := filepath.Abs(gitcdHome)
  if err != nil {
    return ``, errors.New(fmt.Sprintf("Cannot resolve directory `%s`: %s", gitcdHome, err.Error()))
  }
  return absoluteGitcdHome, nil
}
