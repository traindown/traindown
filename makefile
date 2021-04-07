build:
	./bin/speccer
	rsync -a ./script/ ./output/script
	rsync -a ./css/ ./output/css
	rsync -a ./img/ ./output/img
	tinystatic

serve:
	DOMAIN=localhost:1337 caddy run --watch
