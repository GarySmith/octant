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
	"strconv"

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
		IsModule: true,
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
			{
				Title:    "Images",
				Path:     request.GeneratePath("images"),
				IconName: "folder",
			},
			{
				Title:    "Servers",
				Path:     request.GeneratePath("servers"),
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
}

type image struct {
	Name        string
	Status      string
	Visibility  string
	Protected   bool
	DiskFormat  string `json:"Disk Format"`
	Size        int32
}

type server struct {
	Status      string
	Name        string
	ImageName   string `json:"Image Name"`
	Flavor      string `json:"Flavor Name"`
	Networks    string
	PowerState  string `json:"Power State"`
    AvailZone   string `json:"Availability Zone"`
}


func endpointList(request *service.Request) (component.ContentResponse, error) {

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

    cols := component.NewTableCols("ServiceType", "URL", "Region", "Enabled", "Interface", "ServiceName")
    table := component.NewTable("Endpoints", "placeholder", cols)
    for _, elem := range res {
        row := component.TableRow{}
        row["ServiceType"] = component.NewText(elem.ServiceType)
        row["URL"] = component.NewText(elem.URL)
        row["Region"] = component.NewText(elem.Region)
        row["Enabled"] = component.NewText(strconv.FormatBool(elem.Enabled))
        row["Interface"] = component.NewText(elem.Interface)
        row["ServiceName"] = component.NewText(elem.ServiceName)
        table.Add(row)
    }

    contentResponse := component.NewContentResponse(component.TitleFromString("Endpoints"))
    contentResponse.Add(table)
    return *contentResponse, nil
}

func imageList(request *service.Request) (component.ContentResponse, error) {

    // Requires that you have appropriate OS_* variables defined in the environment so that 'openstack image list' succeeds
    out, err := exec.Command("openstack", "image", "list", "--long", "-f", "json").Output()
    if err != nil {
        log.Fatal(err)
    }

    // Convert output into an array slice of image structs.
    res := []image{}
    if err := json.Unmarshal(out, &res); err != nil {
        empty := component.ContentResponse{}
        return empty, err
    }

    // Sort the table by name
    sort.Slice(res, func(i, j int) bool {
        return res[i].Name < res[j].Name
    })

    cols := component.NewTableCols("Name", "Status", "Visibility", "Protected", "DiskFormat", "Size")
    table := component.NewTable("Images", "placeholder", cols)
    for _, elem := range res {
        row := component.TableRow{}
        row["Name"] = component.NewText(elem.Name)
        row["Status"] = component.NewText(elem.Status)
        row["Visibility"] = component.NewText(elem.Visibility)
        row["Protected"] = component.NewText(strconv.FormatBool(elem.Protected))
        row["DiskFormat"] = component.NewText(elem.DiskFormat)
        row["Size"] = component.NewText(fmt.Sprint(elem.Size))
        table.Add(row)
    }

    contentResponse := component.NewContentResponse(component.TitleFromString("Images"))
    contentResponse.Add(table)
    return *contentResponse, nil
}

func serverList(request *service.Request) (component.ContentResponse, error) {

    // Requires that you have appropriate OS_* variables defined in the environment so that 'openstack server list' succeeds
    out, err := exec.Command("openstack", "server", "list", "--all", "--long", "-f", "json").Output()
    if err != nil {
        log.Fatal(err)
    }

    // Convert output into an array slice of server structs.
    res := []server{}
    if err := json.Unmarshal(out, &res); err != nil {
        empty := component.ContentResponse{}
        return empty, err
    }

    // Sort the table by name
    sort.Slice(res, func(i, j int) bool {
        return res[i].Name < res[j].Name
    })

    cols := component.NewTableCols("Name", "Status", "Image Name", "Flavor", "Networks", "Power State", "Availability Zone")
    table := component.NewTable("Servers", "placeholder", cols)
    for _, elem := range res {
        row := component.TableRow{}
        row["Name"] = component.NewText(elem.Name)
        row["Status"] = component.NewText(elem.Status)
        row["Image Name"] = component.NewText(elem.ImageName)
        row["Flavor"] = component.NewText(elem.Flavor)
        row["Networks"] = component.NewText(elem.Networks)
        row["Power State"] = component.NewText(elem.PowerState)
        row["Availability Zone"] = component.NewText(elem.AvailZone)
        table.Add(row)
    }

    contentResponse := component.NewContentResponse(component.TitleFromString("Servers"))
    contentResponse.Add(table)
    return *contentResponse, nil
}

// initRoutes routes for this plugin. In this example, there is a global catch all route
// that will return the content for every single path.
func initRoutes(router *service.Router) {

	router.HandleFunc("/endpoints", endpointList)
	router.HandleFunc("/images", imageList)
	router.HandleFunc("/servers", serverList)
}
