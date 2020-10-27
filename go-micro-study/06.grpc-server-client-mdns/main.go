package main

import (
	"context"
	"fmt"
	greeter "go-micro-study/06.grpc-server-client/proto"
	"log"
	"os"

	"github.com/micro/cli/v2"

	"github.com/micro/go-micro/v2"
)

const (
	serviceName = "greeter"
)

type Greeter struct {
}

func (g Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
func runClient(service micro.Service) {
	cliService := greeter.NewGreeterService(serviceName, service.Client())

	rsp, err := cliService.Hello(context.TODO(), &greeter.Request{Name: "sgfoot"})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print(rsp.GetGreeting())
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "hello world",
		}),
		micro.Flags(&cli.BoolFlag{
			Name:    "run_client",
			Aliases: []string{"cc"},
			Usage:   "Launch the client",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			if c.Bool("run_client") {
				log.Println("run_client")
				runClient(service)
				os.Exit(0)
			}
			return nil
		}),
	)
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
