/*
@Time : 2022/5/27 上午11:21
@Author : tan
@File : dgraph
@Software: GoLand
*/
package db

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"visionvera/vfile_client/utils/config"
)

var DgraphClient *dgo.Dgraph

func DgraphInit() {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
	)
	dgraph, err := grpc.Dial(config.GetString("db.dgraph.host"), dialOpts...)

	if err != nil {
		panic(err)
	}
	DgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(dgraph))
}
