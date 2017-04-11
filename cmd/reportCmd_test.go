//Copyright 2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ottenwbe/golook/routing"
)

var _ = Describe("The query command", func() {

	BeforeEach(func() {
		golookIfce = routing.NewMockedRouter()
	})

	It("triggers a query for reported files, when system is set by using the keyword 'this'", func() {
		system = QUERY_THIS
		queryCmd.Run(nil, nil)
		Expect(golookIfce.(*routing.MockedLookController).Visited).To(BeTrue())
	})
})