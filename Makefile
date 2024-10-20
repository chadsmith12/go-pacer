live/tailwind:
	npx tailwindcss -i ./app/assets/css/tailwind.css -o ./wwwroot/styles/styles.css --watch

live/templ:
	templ generate --watch --proxy="http://localhost:4500" --open-browser=false

live/app:
	air \
		--build.cmd "go build -o ./tmp/app ./cmd/web" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

live/web:
	air \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "app/assets/*, wwwroot" \
		--build.include_ext "js,css"

live:
	make -j2 live/templ live/app live/tailwind live/web
