This is plugin is merely a proof of concept to validate how a plugin can call an openstack
service and render its results in a new pane.

## Discussion

The plugin is takes a bit of a brute force approach in that it directly calls the `openstack` cli every time the
page is rendered, and does not make use of the caching layer that octant uses when dealing with kubernetes objects.

The output of `openstack endpoint list` is converted to markdown text and rendered with the markdown component, which
is kind of a hack.  It would have been better to use the table components, but I coulnd't get that to work.

One of the challenges to developing this plugin is that all log messages and print statements are somehow being surpressed
from the console output, making it very hard to debug, especially since I am new to `go`.


## Installation

Build the plugin

```sh
$ make install-plugins
```

Before starting (or restarting) the server, make sure to define the necessary openstack environment variables (e.g. `OS_CLOUD`)
such that you can successfully call the `openstack endpoint list` command from the shell.  After that is correctly working, (re)start the server.
