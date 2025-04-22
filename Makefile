.PHONY: templ-gen
templ-gen:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: tailwind-gen
tailwind-gen:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: dev
dev:
	go run ./main/main.go
