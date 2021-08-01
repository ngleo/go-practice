package main

import (
    "log"
    "os"
    "fmt"
    "net"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "Website lookup"
    app.Usage = "Query IPs, CNAMEs, MX records and name servers"

    myFlags := []cli.Flag{
        cli.StringFlag {
            Name: "host",
            Value: "google.co.uk",
        },
    }

    app.Commands = []cli.Command {
        {
            Name: "ns",
            Usage: "Looks for name servers for a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                ns, err := net.LookupNS(c.String("host"))
                if err != nil {
                    return err
                }
                for i:= 0; i < len(ns); i++ {
                    fmt.Println(ns[i].Host)
                }
                return nil
            },
        },
        {
            Name: "ip",
            Usage: "Looks for IP address of a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                ip, err := net.LookupIP(c.String("host"))
                if err != nil {
                    return err
                }
                for i:= 0; i < len(ip); i++ {
                    fmt.Println(ip[i])
                }
                return nil
            },
        },
        {
            Name: "cn",
            Usage: "Looks for CNAME records of a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                cn, err := net.LookupCNAME(c.String("host"))
                if err != nil {
                    return err
                }         
                fmt.Println(cn)
                return nil
            },
        },
        {
            Name: "mx",
            Usage: "Looks for MX records of a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                mx, err := net.LookupMX(c.String("host"))
                if err != nil {
                    return err
                }         
                for i := 0; i < len(mx); i++ {
                    fmt.Println(mx[i].Host, mx[i].Pref)
                }
                return nil
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
