all: slideshow primitives echosrv

slideshow:
	~/src/nest/script/n2x slideshow.nest > slideshow.html

primitives: primitives.go
	6g primitives.go && 6l -o primitives primitives.6

echosrv: echosrv.go
	6g echosrv.go && 6l -o echosrv echosrv.6
