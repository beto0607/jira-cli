
build:
	go build -o dist/jira-cli

install:
	cp dist/jira-cli ~/.local/bin/jira-cli

ifeq ("", "$(wildcard ${XDG_CONFIG_HOME}/jira-cli/config.conf)")
	mkdir -p ${XDG_CONFIG_HOME}/jira-cli/
	cp example.conf ${XDG_CONFIG_HOME}/jira-cli/config.conf
endif

