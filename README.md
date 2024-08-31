# Jira CLI

Scripts to handle tickets in Jira from my terminal

### Why?

I got tired of creating a branch, start working on it and realize that I never moved the ticket in Jira, needing to go back to my browser, open Jira and drag&drop. I'd much rather do it from my terminal.

After a first try with curl scripts, I decided to give it a try to do it in GoLang.
Some inspiration came from [gh cli](https://github.com/cli/cli)


### Configs

This tool expects a file in `$XDG_CONFIG_HOME/jira-cli/config.conf` (or `~/.config/jira-cli/config.conf` if `$XDG_CONFIG_HOME` is not set)

``` launguage: conf
[auth]
    token = "<YOUR_TOKEN>"
[user]
    accountId = "<YOUR_ACCOUNT_ID>"
    email = "your@email.com"
```

You will need to fill this up manually, for now

- [How to get your API Token?]()
- How to get your Account ID? Go to your profile, the URL will be something like: `https://your-project.atlassian.net/jira/people/<YOUR_ACCOUNT_ID>` <- Copy and paste it
- How to get your email? 🤔 🤷
