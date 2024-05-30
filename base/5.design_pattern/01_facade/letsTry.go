package main

import "fmt"

/*
	假设我们正在开发一个计算机系统监控程序，该程序可以监控 CPU 使用率、内存占用和磁盘空间等信息。
	为了获取这些信息，我们需要与多个底层系统组件进行交互，比如与 CPU 子系统、内存子系统和磁盘子系统等。
*/

/*
	首先，我们定义各个子系统的接口，表示各个子系统所提供的功能：
*/
// CPU represents the CPU subsystem.
type CPU interface {
	GetUsage() float64
}

// Memory represents the Memory subsystem.
type Memory interface {
	GetUsage() float64
}

// Disk represents the Disk subsystem.
type Disk interface {
	GetFreeSpace() float64
}

/*
	然后，我们实现各个子系统的具体类，分别表示 CPU 子系统、内存子系统和磁盘子系统
*/
// CPUSubsystem represents the CPU subsystem.
type CPUSubsystem struct{}

// GetUsage returns the CPU usage percentage.
func (c *CPUSubsystem) GetUsage() float64 {
	// Simulate getting CPU usage from the actual system.
	return 25.0
}

// MemorySubsystem represents the Memory subsystem.
type MemorySubsystem struct{}

// GetUsage returns the memory usage percentage.
func (m *MemorySubsystem) GetUsage() float64 {
	// Simulate getting memory usage from the actual system.
	return 60.0
}

// DiskSubsystem represents the Disk subsystem.
type DiskSubsystem struct{}

// GetFreeSpace returns the free disk space in GB.
func (d *DiskSubsystem) GetFreeSpace() float64 {
	// Simulate getting free disk space from the actual system.
	return 500.0
}

/*
	接下来，我们创建一个外观类SystemMonitor，它封装了各个子系统的操作，并提供一个简化的接口给客户端：
*/
// SystemMonitor is the facade class that simplifies interactions with subsystems.
type SystemMonitor struct {
	cpu    CPU
	memory Memory
	disk   Disk
}

// NewSystemMonitor creates a new SystemMonitor.
func NewSystemMonitor() *SystemMonitor {
	return &SystemMonitor{
		cpu:    &CPUSubsystem{},
		memory: &MemorySubsystem{},
		disk:   &DiskSubsystem{},
	}
}

// GetCPUUsage returns the CPU usage percentage.
func (sm *SystemMonitor) GetCPUUsage() float64 {
	return sm.cpu.GetUsage()
}

// GetMemoryUsage returns the memory usage percentage.
func (sm *SystemMonitor) GetMemoryUsage() float64 {
	return sm.memory.GetUsage()
}

// GetFreeDiskSpace returns the free disk space in GB.
func (sm *SystemMonitor) GetFreeDiskSpace() float64 {
	return sm.disk.GetFreeSpace()
}

/*
现在，我们可以在客户端使用外观模式来获取系统的各项监控信息，而无需直接与底层的子系统交互：
*/
func main() {
	monitor := NewSystemMonitor()

	cpuUsage := monitor.GetCPUUsage()
	memoryUsage := monitor.GetMemoryUsage()
	freeDiskSpace := monitor.GetFreeDiskSpace()

	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Printf("Memory Usage: %.2f%%\n", memoryUsage)
	fmt.Printf("Free Disk Space: %.2f GB\n", freeDiskSpace)
}
