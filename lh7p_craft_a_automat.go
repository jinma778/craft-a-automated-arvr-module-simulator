package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulator represents an automated AR/VR module simulator
type Simulator struct {
	ModuleName string
	Modules     map[string]int
}

// NewSimulator creates a new Simulator instance
func NewSimulator(moduleName string) *Simulator {
	return &Simulator{
		ModuleName: moduleName,
		Modules:    make(map[string]int),
	}
}

// AddModule adds a new module to the simulator
func (s *Simulator) AddModule(moduleName string, quantity int) {
	s.Modules[moduleName] = quantity
}

// RemoveModule removes a module from the simulator
func (s *Simulator) RemoveModule(moduleName string) {
	delete(s.Modules, moduleName)
}

// GetModuleQuantity returns the quantity of a specific module
func (s *Simulator) GetModuleQuantity(moduleName string) int {
	return s.Modules[moduleName]
}

// SimulateARVRModule simulates an AR/VR module
func (s *Simulator) SimulateARVRModule(moduleName string) {
	fmt.Printf("Simulating %s module...\n", moduleName)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	fmt.Printf("%s module simulation complete!\n", moduleName)
}

func main() {
	simulator := NewSimulator("LH7P Simulator")
	simulator.AddModule("AR Lens", 5)
	simulator.AddModule("VR Headset", 3)
	simulator.AddModule("Sensor Array", 2)

	fmt.Println("Simulator modules:")
	for moduleName, quantity := range simulator.Modules {
		fmt.Printf("- %s: %d\n", moduleName, quantity)
	}

	simulator.SimulateARVRModule("AR Lens")
	simulator.SimulateARVRModule("VR Headset")
	simulator.SimulateARVRModule("Sensor Array")

	simulator.RemoveModule("Sensor Array")
	fmt.Println("Simulator modules after removal:")
	for moduleName, quantity := range simulator.Modules {
		fmt.Printf("- %s: %d\n", moduleName, quantity)
	}
}