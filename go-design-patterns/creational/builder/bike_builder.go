package builder

type BikeBuilder struct {
	vehicle VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.vehicle.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.vehicle.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.vehicle.Structure = "Bike"
	return b
}

func (b *BikeBuilder) Build() VehicleProduct {
	return b.vehicle
}
