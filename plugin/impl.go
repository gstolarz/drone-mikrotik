package plugin

import (
    "crypto/tls"
    "fmt"
    "strconv"
    "strings"

    "github.com/urfave/cli/v2"
    "gopkg.in/routeros.v2"
)

const (
    defaultApiPort    = 8728
    defaultApiTlsPort = 8729
)

// Settings for the plugin.
type Settings struct {
    Address  string
    TLS      bool
    Insecure bool

    Username string
    Password string

    Script cli.StringSlice
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
    // Validation of the settings.
    if len(p.settings.Address) == 0 {
        return fmt.Errorf("no MikroTik address provided")
    }

    if len(p.settings.Username) == 0 {
        return fmt.Errorf("no MikroTik username provided")
    }

    if len(p.settings.Password) == 0 {
        return fmt.Errorf("no MikroTik password provided")
    }

    if strings.Index(p.settings.Address, ":") == -1 {
        if p.settings.TLS {
            p.settings.Address += ":" + strconv.Itoa(defaultApiTlsPort)
        } else {
            p.settings.Address += ":" + strconv.Itoa(defaultApiPort)
        }
    }

    return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
    var client *routeros.Client
    var err error

    fmt.Printf("Connecting to %s\n", p.settings.Address)

    if p.settings.TLS {
        client, err = routeros.DialTLS(p.settings.Address, p.settings.Username, p.settings.Password, &tls.Config{
            InsecureSkipVerify: p.settings.Insecure,
        })
    } else {
        client, err = routeros.Dial(p.settings.Address, p.settings.Username, p.settings.Password)
    }
    if err != nil {
        return fmt.Errorf("error while connecting to %s: %w", p.settings.Address, err)
    }
    defer client.Close()

    for _, script := range p.settings.Script.Value() {
        fmt.Printf("Executing: %s\n", script)
        reply, err := client.RunArgs(strings.Split(script, " "))
        if err != nil {
            return fmt.Errorf("error while executing command: %w", err)
        }

        headers := false
        for _, re := range reply.Re {
            if !headers {
                for i, el := range re.List {
                    fmt.Print(el.Key)
                    if i < len(re.List)-1 {
                        fmt.Print("\t")
                    }
                }
                fmt.Println()
                headers = true
            }

            for i, el := range re.List {
                fmt.Print(el.Value)
                if i < len(re.List)-1 {
                    fmt.Print("\t")
                }
            }
            fmt.Println()
        }
    }

    return nil
}
