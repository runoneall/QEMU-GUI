package vars

var QEMU_SUPPORTED_ARCH = []string{
	"amd64",
	"i386",
	"arm64",
	"arm",
	"riscv64",
}

var QEMU_ARCH = map[string]string{
	"amd64":   "qemu-system-x86_64",
	"i386":    "qemu-system-i386",
	"arm64":   "qemu-system-aarch64",
	"arm":     "qemu-system-arm",
	"riscv64": "qemu-system-riscv64",
}

var QEMU_CPU = []string{
	"host",
	"max",
}

var QEMU_MACHINE = []string{
	"pc",
	"q35",
	"virt",
}

var QEMU_GPU = []string{
	"std",
	"qxl",
	"virtio-gpu-pci",
}

var QEMU_ACCEL = []string{
	"kvm",
	"hvf",
	"whpx",
	"tcg",
}
