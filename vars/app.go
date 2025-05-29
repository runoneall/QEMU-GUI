package vars

import "path/filepath"

var DATA_PATH = filepath.Join(".", "data")
var CONFIG_PATH = filepath.Join(DATA_PATH, "config")
var CONFIG_FILE = filepath.Join(CONFIG_PATH, "config.json")
var VM_PATH = filepath.Join(DATA_PATH, "vms")
