// +build ledger_device

/*******************************************************************************
*   (c) 2018 ZondaX GmbH
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
********************************************************************************/

package ledger_go

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zondax/hid"
)

func Test_ThereAreDevices(t *testing.T) {
	devices := hid.Enumerate(0, 0)
	assert.NotEqual(t, 0, len(devices))
}

/*
func Test_ListDevices(t *testing.T) {
	ListDevices()
}
*/

func Test_FindLedger(t *testing.T) {
	ledger, err := FindLedger()
	if err != nil {
		fmt.Println("\n*********************************")
		fmt.Println("Did you enter the password??")
		fmt.Println("*********************************")
		t.Fatalf("Error: %s", err.Error())
	}
	assert.NotNil(t, ledger)
	ledger.Close()
}
