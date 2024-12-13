package models

type ParkingLot struct {
	Spaces       []*Space
	Cars         []*Car
}

func NewParkingLot() *ParkingLot {
	parkingLot := &ParkingLot{
		Spaces:       initializeSpaces(),
		Cars:         initializeCars(),
	}
	return parkingLot
}

func initializeSpaces() []*Space {
	coordinates := [][2]float64{
		{378, 46}, {439, 46}, {501, 46}, {750, 14}, {750, 76},
		{750, 136}, {750, 198}, {750, 260}, {750, 321}, {750, 383},
		{750, 444}, {750, 505}, {458, 545}, {520, 545}, {582, 545},
	}
	rotations := []float64{
		0, 0, 0, 90, 90,
		90, 90, 90, 90, 90,
		90, 90, 180, 180, 180,
	}

	var spaces []*Space
	for i := 0; i < 15; i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		rotation := rotations[i]
		space := NewSpace(i, x, y, 0.16, rotation)
		spaces = append(spaces, space)
	}
	return spaces
}

func initializeCars() []*Car{
	var cars []*Car

	for i := 0; i < 15; i++ {
		car := NewCar(i, 170.0, -1.0, 1.0, 2.0)
		cars = append(cars, car)
	}
	
	return cars
}

func (p *ParkingLot) Update() {
	for _, car := range p.Cars {
		car.Update()
	}
}
