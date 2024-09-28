package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
