install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/1024pix/metabase@latest/steampipe-plugin-metabase.plugin *.go

dev:
	go build -o  ~/.steampipe/plugins/local/metabase@latest/steampipe-plugin-metabase.plugin *.go
