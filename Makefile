all: slideshow demo

demo:
	6g primitives.go && 6l -o demo primitives.6

slideshow:
	~/src/nest/script/n2x slideshow.nest > slideshow.html
