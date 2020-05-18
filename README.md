# docksec
This is an authorization plugin made keeping simplicity in mind. This lets you dissable any docker command by just naming that in a a json configuration file.

# Configurtion

Configuration to `docksec` is done via a json config file located in `/etc/docksec/main.json`. The file is not present there by default. So make sure you create it `# mkdir -pv /etc/docksec && touch /etc/docksec/main.json`. The plugin options are stored in a "plugin" object.

```
{
    "plugin": {
        ...
    }
}
```

Inside, Create a "commands" array, this will store the commands you're going to blacklist (For a list of available commands please refer to the release page).

```
{
    "plugins": {
        "commands": [
            {
                ...
            },
            {
                ...
            }
        ]
    }
}
```

The commands array holds multiple objects. Each object has the fields "cmd", "dmsg" & "allow".

cmd
---

This field holds a single string, the command that you want to blacklist, without prepending with "docker". For example if you ant to disable `docker inspect`, use the value "inspect" for "cmd" field.

dmsg
-----

This field  contains a string, a message that'll be printed upon denied access.

allow
-----

This enables the handler (TODO) to disable and enable a command temporarily without removing the whole object from the config file.

Outside the commands array, there's another field named "description". It takes a boolean value, which tells `docksec` whether to show a `access denied` message or not.

## Sample config

```
{
    "plugin": {
        "description": true,
        "commands": [
            {
                "cmd": "ps",
                "dmsg": "Access denied",
                "allow": false
            }
        ]
    }
}
```

After you're done with the config file, restart docksec daemon `# systemctl restart docksec`