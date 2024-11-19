/*
 Copyright 2021 Qiniu Cloud (qiniu.com)
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ar0c/goc/v2/pkg/client/rest"
	"github.com/ar0c/goc/v2/pkg/log"
	"github.com/olekukonko/tablewriter"
)

const (
	DISCONNECT   = 1 << iota
	RPCCONNECT   = 1 << iota
	WATCHCONNECT = 1 << iota
)

func ListAgents(host string, ids []string, wide, isJson bool) {
	gocClient := rest.NewV2Client(host)

	agents, err := gocClient.Agent().Get(ids)

	if err != nil {
		log.Fatalf("cannot get agent list from goc server: %v", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	if isJson {
		goto asJson
	}

	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("   ") // pad with 3 blank spaces
	table.SetNoWhiteSpace(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	if wide {
		table.SetHeader([]string{"ID", "STATUS", "REMOTEIP", "HOSTNAME", "PID", "CMD", "EXTRA"})
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	} else {
		table.SetHeader([]string{"ID", "STATUS", "REMOTEIP", "CMD"})
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	}
asJson:
	for _, agent := range agents {
		var status string
		if agent.Status == DISCONNECT {
			status = "DISCONNECT"
		} else if agent.Status&(RPCCONNECT|WATCHCONNECT) > 0 {
			status = "CONNECT"
		}
		agent.StatusStr = status
		if !isJson {
			if wide {
				table.Append([]string{agent.Id, status, agent.RemoteIP, agent.Hostname, agent.Pid, agent.CmdLine, agent.Extra})
			} else {
				preLen := len(agent.Id) + len(agent.RemoteIP) + 9
				table.Append([]string{agent.Id, status, agent.RemoteIP, getSimpleCmdLine(preLen, agent.CmdLine)})
			}
		}
	}
	if !isJson {
		table.Render()
	} else {
		b, _ := json.Marshal(agents)
		fmt.Fprint(os.Stdout, string(b))
	}
}

func DeleteAgents(host string, ids []string) {
	gocClient := rest.NewV2Client(host)

	err := gocClient.Agent().Delete(ids)

	if err != nil {
		log.Fatalf("cannot delete agents from goc server: %v", err)
	}
}
