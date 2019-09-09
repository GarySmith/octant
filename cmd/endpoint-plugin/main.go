/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
    "sort"
    "strings"

    "github.com/vmware/octant/pkg/navigation"
    "github.com/vmware/octant/pkg/plugin"
    "github.com/vmware/octant/pkg/plugin/service"
    "github.com/vmware/octant/pkg/view/component"
)

var pluginName = "endpoint-plugin"

// This is a sample plugin showing the features of Octant's plugin API.
func main() {
    // Remove the prefix from the go logger since Octant will print logs with timestamps.
    log.SetPrefix("")

    // Tell Octant to call this plugin has its own new menu item
    capabilities := &plugin.Capabilities{
        IsModule:              true,
    }

    // Set up what should happen when Octant calls this plugin.
    options := []service.PluginOption{
        service.WithNavigation(handleNavigation, initRoutes),
    }

    // Use the plugin service helper to register this plugin.
    p, err := service.Register(pluginName, "EndpointPlugin", capabilities, options...)
    if err != nil {
        log.Fatal(err)
    }

    // The plugin can log and the log messages will show up in Octant.
    log.Printf("%s is starting", pluginName)
    p.Serve()
}


// handleNavigation creates a navigation tree for this plugin. Navigation is dynamic and will
// be called frequently from Octant. Navigation is a tree of `Navigation` structs.
// The plugin can use whatever paths it likes since these paths can be namespaced to the
// the plugin.
func handleNavigation(request *service.NavigationRequest) (navigation.Navigation, error) {
    return navigation.Navigation{
        Title: "OpenStack",
        Path:  request.GeneratePath(),
        Children: []navigation.Navigation{
            {
                Title:    "Endpoints",
                Path:     request.GeneratePath("endpoints"),
                IconName: "folder",
            },
        },
        IconName: "cloud",
    }, nil
}

type endpoint struct {
    ServiceType string `json:"Service Type"`
    URL         string
    Region      string
    Enabled     bool
    Interface   string
    ServiceName string `json:"Service Name"`
    ID          string
}

// initRoutes routes for this plugin. In this example, there is a global catch all route
// that will return the content for every single path.
func initRoutes(router *service.Router) {

    router.HandleFunc("/*", func(request *service.Request) (component.ContentResponse, error) {

        // Requires that you have appropriate OS_* variables defined in the environment so that 'openstack endpoint list' succeeds
        out, err := exec.Command("openstack", "endpoint", "list", "-f", "json").Output()
        if err != nil {
            log.Fatal(err)
        }

        // Convert output into an array slice of endpoint structs.
        res := []endpoint{}
        if err := json.Unmarshal(out, &res); err != nil {
            empty := component.ContentResponse{}
            return empty, err
        }

        // Sort the table by service type and interface
        sort.Slice(res, func(i, j int) bool {
            if res[i].ServiceType < res[j].ServiceType {
                return true
            } else if res[i].ServiceType > res[j].ServiceType {
                return false
            }
            return res[i].Interface < res[j].Interface
        })

        var sb strings.Builder
        sb.WriteString("| ServiceType | URL | Region | Enabled | Interface | ServiceName |\n")
        sb.WriteString("|-------------|-----|--------|---------|-----------|-------------|\n")
        /*
        cols := component.NewTableCols("ServiceType", "URL", "Region", "Enabled", "Interface", "ServiceName", "ID")
        table := component.NewTable("My Table", "placeholder", cols)
        */
        for _, elem := range res {
            sb.WriteString(fmt.Sprintf("| %s | %s | %s | %t | %s | %s |\n",
                elem.ServiceType,
                elem.URL,
                elem.Region,
                elem.Enabled,
                elem.Interface,
                elem.ServiceName))

            /*
            row := component.TableRow{}
            row["ServiceType"] = component.NewText(elem.ServiceType)
            row["URL"]         = component.NewText(elem.URL)
            row["Region"]      = component.NewText(elem.Region)
            row["Enabled"]     = component.NewText(strconv.FormatBool(elem.Enabled))
            row["Interface"]   = component.NewText(elem.Interface)
            row["ServiceName"] = component.NewText(elem.ServiceName)
            row["ID"]          = component.NewText(elem.ID)
            table.Add(row)
            */
        }

        card := component.NewCard("Endpoint List")

        // Cannot get the stupid table to render
        // card.SetBody(table)
        card.SetBody(component.NewMarkdownText(sb.String()))

        cardList := component.NewCardList("Tab 1")
        cardList.AddCard(*card)
        cardList.SetAccessor("tab1")

        contentResponse := component.NewContentResponse(component.TitleFromString("Endpoints"))
        contentResponse.Add(cardList)

        return *contentResponse, nil
    })
}
