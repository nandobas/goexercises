package services

import "testing"

func TestBridgeCircle(t *testing.T){
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(3)
	circle.Draw()
}

func TestBridgeRaster(t *testing.T){
	raster := RasterRenderer{}
	circle := NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(3)
	circle.Draw()
}
