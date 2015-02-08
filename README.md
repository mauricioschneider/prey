# goprey
Experimental client for preyproject.com written in golang, exclusively for learning purposes. At first, the script will have a very basic set of features, but I'll be working my way up to a more robust solution as I get better with golang.

MVP should:
- Get status of a device, that is, if whether the device is missing or not
- Execute the following actions: alarm, alert, lock
- Send reports in case the device is missing

This functionality will work on a on-demand basis, that is, explicity running the script (`go run prey.go`).

Next steps after MVP:
- Use intervals to check for status and instructions
- Daemonize
- Allow the use of plugins, and "pluginize" the control-panel.
