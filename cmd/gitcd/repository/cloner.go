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
  "os"
  "path"
  "os/exec"
  "fmt"
  "log"
)

func Clone(absoluteGitcdHome string, repositoryString string, repository Repository) error {
  // Makes all the directories up to the owner directory.
  ownerDirectory := path.Join(absoluteGitcdHome, repository.Owner)
  err := os.MkdirAll(ownerDirectory, 0755)
  if err != nil {
    return err
  }

  // Tries to clone the original repositoryString first.
  err = command(exec.Command("git", "-C", ownerDirectory, "clone", repositoryString))
  if err == nil {
    return nil
  }

  // If that fails, then tries to construct a clone-able URL from repository.
  repositoryUrl := fmt.Sprintf("https://github.com/%s/%s", repository.Owner, repository.Name)
  log.Printf("Cloning repository `%s` failed, trying again with `%s`...\n", repositoryString, repositoryUrl)
  return command(exec.Command("git", "-C", ownerDirectory, "clone", repositoryUrl))
}

func command(cmd *exec.Cmd) error {
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}
