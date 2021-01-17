build:
	rsync -a ./script/ ./output/script
	rsync -a ./css/ ./output/css
	rsync -a ./img/ ./output/img
	tinystatic
