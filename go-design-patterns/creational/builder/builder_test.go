package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 but they were %d", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 but they were %d", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Car structure must be 'Car' but it was %s", car.Structure)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.Build()
	bike.Seats = 1

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 but they were %d", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Bike structure must be 'Bike' but it was %s", bike.Structure)
	}
}
