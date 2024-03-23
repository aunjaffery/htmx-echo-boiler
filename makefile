run: templ_gen
	go run main.go

templ_gen: tw_gen
	templ generate

tw_gen:
	npx tailwindcss -i ./public/input.css -o ./public/output.css
