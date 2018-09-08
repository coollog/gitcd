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

import "testing"

func TestCanonicalize(t *testing.T) {
  expectedRepositories := []struct {
    repositoryString string
    expectedRepository Repository
    shouldError bool
  }{
    {"coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"coollog_underscore/gitcd_underscore-dash", Repository{"coollog_underscore", "gitcd_underscore-dash"}, false},
    {"github.com/coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"github.com/coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"http://github.com/coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"http://github.com/coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"https://github.com/coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"https://github.com/coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"ssh://github.com/coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"ssh://github.com/coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"git@github.com:coollog/gitcd", Repository{"coollog", "gitcd"}, false},
    {"git@github.com:coollog/gitcd.git", Repository{"coollog", "gitcd"}, false},
    {"notvalid", Repository{}, true},
    {"/not/valid", Repository{},true},
    {"not/valid/", Repository{},true},
    {"github.com:coollog/gitcd", Repository{},true},
  }

  for _, expectedRepository := range expectedRepositories {
    repo, err := Canonicalize(expectedRepository.repositoryString)

    if expectedRepository.shouldError {
      if err == nil {
        t.Errorf("Parse invalid repository `%s` should have errored", expectedRepository.repositoryString)
      }
      continue
    }

    if err != nil {
      t.Errorf("Parse valid repository `%s` errored: %s", expectedRepository.repositoryString, err.Error())
      continue
    }

    if repo != expectedRepository.expectedRepository {
      t.Errorf("Parse respository `%s` expected `%#v` but got `%#v`", expectedRepository.repositoryString, expectedRepository.expectedRepository, repo)
    }
  }
}
