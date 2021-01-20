/*
 * MinIO Cloud Storage, (C) 2017-2020 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bvs

import (
	"github.com/minio/cli"
	minio "github.com/minio/minio/cmd"

	"github.com/minio/minio/cmd/logger"
)

func init() {
	const bvsGatewayTemplate = `NAME:
	{{.HelpName}} - {{.Usage}}
  
  USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[FLAGS]{{end}} [ENDPOINT]
  {{if .VisibleFlags}}
  FLAGS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
  ENDPOINT:
	bvs server endpoint.
  
  EXAMPLES:
	1. Start sgw gateway server for BVS backend
	   {{.Prompt}} {{.EnvVarSetCommand}} MINIO_ROOT_USER{{.AssignmentOperator}}accesskey
	   {{.Prompt}} {{.EnvVarSetCommand}} MINIO_ROOT_PASSWORD{{.AssignmentOperator}}secretkey
	   {{.Prompt}} {{.HelpName}}
  `

	minio.RegisterGatewayCommand(cli.Command{
		Name:               minio.BVSBackendGateway,
		Usage:              "Storage Service Gateway",
		Action:             bvsGatewayMain,
		CustomHelpTemplate: bvsGatewayTemplate,
		HideHelpCommand:    true,
	})
}

// Handler for 'minio gateway bvs' command line.
func bvsGatewayMain(ctx *cli.Context) {
	args := ctx.Args()
	if !ctx.Args().Present() {
		return
	}

	serverAddr := ctx.GlobalString("address")
	if serverAddr == "" || serverAddr == ":"+minio.GlobalMinioDefaultPort {
		serverAddr = ctx.String("address")
	}
	// Validate gateway arguments.
	logger.FatalIf(minio.ValidateGatewayArguments(serverAddr, args.First()), "Invalid argument")

	// Start the gateway..
	minio.StartGateway(ctx, &BVS{args.First()})
}

// BVS implements Gateway.
type BVS struct {
	host string // scheme + ip:port
}
