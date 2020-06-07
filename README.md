# docksec

This is an authorization plugin made by itsfoss while keeping the KISS philosophy in mind. This lets you disable any docker command by just using its name.

# Why use an authorization plugin ?

An authorization plugin lets you disable specific commands based on your development environment, infrastructure, dev team etc. Why might you need that? Consider on a server you've deployed a wordpress instance under the user "Ross", and someone gets access to that user. He/She can then pretty easily damage the container beyond repair with `docker exec` (`$ docker exec -ti wordpress bash`). Because you can't use user mapping for every container, let's say it's disabled by default. If an attacker gets access to any of the user that's part of the docker group, consider your host gone. `$ docker run --rm -v  /bin:/begone alpine rm -r /begone/*` can just easily kill your development server. Usernamespace mapping doesn't even have to be disabled, the attacker can use `--userns host` to disable that for a specific container, and get on with killing all of your hardwork. Let's consider another situation, which probably isn't that fatal, to calm you down a little bit I guess. Consider a single server, that's hosting a database container. Instead of secrets, you've gone with env variables. The attacker can easily call inspect on the container and extract all of that information.  There are many other ways an attacker can damage your containers or even your host just because all the docker commands were accessible by anyone in the docker group on that system, be that the actual admin, or a student who's just trying to learn docker. 

In all the previous situations, disabling `docker run`, `docker inspect`, `docker exec` when in doubt, can increase the security by a large scale.

# How docksec works

Docksec works by mapping user configured commands to their respective api calls, internally creating a static list of http method & endpoints that the administrator considers not necessary to be enabled in a specific time period.

# Configurtion

# Using the handler

TODO

# Usng the json config file

Configuration to `docksec` can also be done via a json config file located in `/etc/docksec/main.json`. The file is not present there by default. So make sure you create it `# mkdir -pv /etc/docksec && touch /etc/docksec/main.json`. The plugin options are stored in a "plugin" object.

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

This field holds a single string, the command that you want to blacklist, without prepending with "docker". For example if you want to disable `docker inspect`, use the value "inspect" for the "cmd" field.

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

# Plans for future implementation

The following is the list containing the features that are getting considered to be benificial for `docksec`'s future. 

User based permissions
--------

UserA has permission to create a container, but userb doesn't, while they're on the same docker host. This ensures layered security & lesssens the vulneribility that rootless-docker poses.

Cluster implementation
-------

Like docker itself, `docksec`'s own cluster implementation, this will enable sharing a single config across multiple nodes, increasing administrator's  efficiency.

Container|Image|Network|Volume based permissions
--------

Along with permissions, based on specific users, this will enable permissions to docker commands based on on which container/image/network/volume the command is getting executed. This will add another layer of security. Consider there are two teams, and you want to use a single server for two (or more) different projects. You can therefore turn any access off for the elements of your project, and be rest assured that they won't be messed with. 


# Note

The program is still nowhere done. This roadmap is added for others inputs on this, an possibly enhance the list as well.
