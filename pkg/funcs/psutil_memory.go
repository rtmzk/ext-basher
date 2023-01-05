package funcs

import (
	"github.com/shirou/gopsutil/v3/mem"
)

var memory, _ = mem.VirtualMemory()

var memoryAvailable = new(Func).
	builder("mem_available", ExportVariables).
	setVariables(unitFormat(memory.Available)).
	build()

var memoryUsage = new(Func).
	builder("mem_usage", ExportVariables).
	setVariables(unitFormat(memory.Used)).
	build()
