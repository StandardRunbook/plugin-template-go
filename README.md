# plugin-template-go

This plugin template for go is for anyone that wants to make their own plugins. By no means is it necessary to use this template, only that this plugin interface must be honored: [plugin-interface](https://github.com/StandardRunbook/plugin-interface)

There's a script in `app/script/run.sh` that can be configured and this template will automatically run the script. You can also configure the cleanup script within the same `run.sh` script.

In the v0 version, scripts are embedded directly into the go binary to make deployment really easy (no volume mounts). There are caveats to this:

- no secure information should be put into the scripts
- this means use a secrets manager or address a credentials file like `.aws/credentials`
- plugin binaries will be slightly larger

We intend on making it easier to configure how the scripts are loaded with [hypothecary](https://github.com/StandardRunbook/hypothecary).
