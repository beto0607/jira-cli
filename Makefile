
build:
	go build -o dist/jira-cli
	

install:
	cp dist/jira-cli ~/.local/bin/jira-cli

ifeq ("", "$(wildcard ~/.config/jira-cli/config.conf)")
	cp example.conf ~/.config/jira-cli/config.conf
endif

