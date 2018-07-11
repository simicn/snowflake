package main

import (
	"net"
	"net/http"
	"os"
	pb "snowflake/proto"

	cli "gopkg.in/urfave/cli.v2"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		log.Info(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	app := &cli.App{
		Name: "snowflake",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "listen",
				Value: ":10000",
				Usage: "listening address:port",
				EnvVars: []string{"LISTEN"},
			},
			&cli.StringSliceFlag{
				Name:  "etcd-hosts",
				Value: cli.NewStringSlice("http://127.0.0.1:2379"),
				Usage: "etcd hosts",
				EnvVars: []string{"ETCD_HOSTS"},
			},
			&cli.IntFlag{
				Name:  "machine-id",
				Value: 0,
				Usage: "snowflake machine id, 0-1023",
				EnvVars: []string{"MACHINE_ID"},
			},
			&cli.StringFlag{
				Name:  "pk-root",
				Value: "/seqs",
				Usage: "path for auto increment primary keys",
				EnvVars: []string{"PK_ROOT"},
			},
			&cli.StringFlag{
				Name:  "uuid-key",
				Value: "/seqs/snowflake-uuid",
				Usage: "uuid main key",
				EnvVars: []string{"UUID_KEY"},
			},
		},
		Action: func(c *cli.Context) error {
			log.Println("listen:", c.String("listen"))
			log.Println("etcd-hosts:", c.StringSlice("etcd-hosts"))
			log.Println("machine-id:", c.Int("machine-id"))
			log.Println("pk-root:", c.String("pk-root"))
			log.Println("uuid-key:", c.String("uuid-key"))
			// 监听
			lis, err := net.Listen("tcp", c.String("listen"))
			if err != nil {
				log.Fatalln(err)
			}
			log.Info("listening on ", lis.Addr())

			// 注册服务
			s := grpc.NewServer()
			ins := &server{}
			ins.init(c)
			pb.RegisterSnowflakeServiceServer(s, ins)

			// 开始服务
			return s.Serve(lis)
		},
	}
	app.Run(os.Args)

}
