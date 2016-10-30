### Atlas Map Generation Service ###

This repo provides an intranetwork (my home network) REST service for use with <a href="https://github.com/subtlepseudonym/colloid">colloid</a>

#### Requirements ####

+ Go
	- <a href="https://golang.org/doc/install">Install Golang</a>

#### Project Files ####

+ main/main.go
	- On startup, prints the intranetwork ip and port that it's listening on
	- TODO: swagger for API endpoints (and update of /)
+ atlas/atlas.go
	- First attempt at plate tectonics style map generation
	- TODO: group higher resolution, non-gaussian map into plates