
# How to configure uWelcome

uWelcome has a default built-in look, but it's actually made to be highly customizable.
If you're managing a custom system, it might interest you.

## Where to put your config

You can create a custom config file at `/etc/uwelcome/config.json` (system-wide) or `~/.config/uwelcome/config.json` (user-specific).

## Translations ?!

uWelcome supports translations for its interface, as well as for some command descriptions and link names.
They have specific names / codes that are used to get translated strings.

Any other option not listed won't be translated.

## Breaking down the configuration file

Here's a breakdown of the config file options - there's also the example folder if you want to see concrete use cases.

## Commands

This option allows you to define a list of commands to display in the banner.

Here are the unique codes you can use to get translated strings for command descriptions : 

- `cmd_list`: "List of available commands"
- `cli_pkg`: "Manage command line packages"
- `term_bling`: "Enable terminal bling"
- `banner_toggle`: "Toggle this banner on/off" (there are built-in commands for this)
- `sys_info`: "View system info"
- `man_upd`: "Manually update the system"

```json
{
  "commands": [
    {
      "cmd": "uwelcome toggle",
      "desc": "banner_toggle"
    },
    {
      "cmd": "fastfetch",
      "desc": "sys_info"
    },
    {
      "cmd": "brew help",
      "desc": "cli_pkg"
    },
    {
      "cmd": "cowsay",
      "desc": "Display a cow saying something"
    }
  ]
}
```

## Links

This option allows to add custom links to the banner.

There are unique names you can use to get a translated name for the link :

- `website` : "Website"
- `issues` : "Report an issue"
- `docs` : "Documentation"
- `discuss` : "Discuss"
- `discord` : "Discord"
- `matrix` : "Matrix"
- `bluesky` : "Bluesky"
- `mastodon` : "Mastodon"
- `donate` : "Donate"

```json
{
  "links": [
    {
      "name": "issues",
      "url": "https://issues.bazzite.gg/"
    },
    {
      "name": "docs",
      "url": "https://docs.bazzite.gg/"
    },
    {
      "name": "discord",
      "url": "https://discord.gg/bazzite"
    },
    {
      "name": "bluesky",
      "url": "https://bluesky.bazzite.gg/"
    },
    {
      "name": "discuss",
      "url": "https://github.com/ublue-os/aurora/discussions"
    },
    {
      "name": "Custom Link (it won't get translated)",
      "url": "https://www.innersloth.com/games/among-us/"
    }
  ]
}
```

## Prefix and Suffix

These options allow to customize the prefix and suffix of the welcome message.

```json
{
  "prefix": "> ",
  "suffix": " !"
}
```

Example:

```
`> Welcome to Bluefin !`
```

## Motd

The `motd` option allows to add messages to the banner !
They will get chosen at random and can both be straight up messages or the result of commands.

```json
{
  "motd": {
    "messages" : [
      "This is a custom tip by yours truly ! :D",
      "This is another custom tip (they won't get translated)"
    ],
    "commands" : [
      "umotd"
    ]
  }
}
```

## Use Accent Color

This option allows the banner to use the accent color of the system.

> Note: It's only available for the GNOME desktop as it relies on `dconf`.

```json
{
  "use-accent-color": true
}
```