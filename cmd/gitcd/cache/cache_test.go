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
  "testing"
  "github.com/coollog/gitcd/cmd/gitcd/repository"
  "reflect"
)

func TestBump(t *testing.T) {
  repoCache := RepoCache {
    apiVersion: 1,
    nameMap: map[string][]string{
      `gitcd`: {`coollog`, `imposter`},
      `bar`: {`foo`},
    },
  }

  repoCache.testBump(`owner`, `name`, []string{`owner`}, t)
  repoCache.testBump(`foo2`, `bar`, []string{`foo2`, `foo`}, t)
  repoCache.testBump(`foo`, `bar`, []string{`foo`, `foo2`}, t)
  repoCache.testBump(`fake`, `gitcd`, []string{`fake`, `coollog`, `imposter`}, t)
  repoCache.testBump(`coollog`, `gitcd`, []string{`coollog`, `fake`, `imposter`}, t)
}

func (repoCache *RepoCache) testBump(owner string, name string, expectedOwnerList []string, t *testing.T) {
  repoToBump := repository.Repository{
    Owner: owner,
    Name: name,
  }
  repoCache.Bump(repoToBump)
  ownerList, ok := repoCache.nameMap[name]
  if !ok {
    t.Errorf("Bump did not work: %#v\n", repoToBump)
  }
  if !reflect.DeepEqual(ownerList, expectedOwnerList) {
    t.Errorf("Owner list expected `%#v`, but got `%#v`", expectedOwnerList, ownerList)
  }
}
