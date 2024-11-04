package license

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

// GetMotherboardSerial retrieves the motherboard serial number for Windows and Linux
func GetMotherboardSerial() (string, error) {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("wmic", "baseboard", "get", "SerialNumber")
        output, err := cmd.Output()
        if err != nil {
            return "", err
        }
        lines := strings.Split(string(output), "\n")
        if len(lines) > 1 {
            return strings.TrimSpace(lines[1]), nil
        }
    } else if runtime.GOOS == "linux" {
        cmd := exec.Command("sudo", "dmidecode", "-s", "baseboard-serial-number")
        output, err := cmd.Output()
        if err != nil {
            return "", err
        }
        return strings.TrimSpace(string(output)), nil
    }
    return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
}

// GetCPUID retrieves the CPU ID for Windows and Linux
func GetCPUID() (string, error) {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("wmic", "cpu", "get", "ProcessorId")
        output, err := cmd.Output()
        if err != nil {
            return "", err
        }
        lines := strings.Split(string(output), "\n")
        if len(lines) > 1 {
            return strings.TrimSpace(lines[1]), nil
        }
    } else if runtime.GOOS == "linux" {
        data, err := os.ReadFile("/proc/cpuinfo")
        if err != nil {
            return "", err
        }
        lines := strings.Split(string(data), "\n")
        for _, line := range lines {
            if strings.HasPrefix(line, "Serial") {
                parts := strings.Split(line, ":")
                if len(parts) > 1 {
                    return strings.TrimSpace(parts[1]), nil
                }
            }
        }
    }
    return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
}

// GetDiskSerial retrieves the primary disk serial number for Windows and Linux
func GetDiskSerial() (string, error) {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("wmic", "diskdrive", "get", "SerialNumber")
        output, err := cmd.Output()
        if err != nil {
            return "", err
        }
        lines := strings.Split(string(output), "\n")
        if len(lines) > 1 {
            return strings.TrimSpace(lines[1]), nil
        }
    } else if runtime.GOOS == "linux" {
        cmd := exec.Command("udevadm", "info", "--query=all", "--name=sda") // "sda" for primary disk
        output, err := cmd.Output()
        if err != nil {
            return "", err
        }
        lines := strings.Split(string(output), "\n")
        for _, line := range lines {
            if strings.Contains(line, "ID_SERIAL") {
                parts := strings.Split(line, "=")
                if len(parts) > 1 {
                    return strings.TrimSpace(parts[1]), nil
                }
            }
        }
    }
    return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
}

func main() {
    motherboardSerial, err := GetMotherboardSerial()
    if err != nil {
        fmt.Println("Error retrieving motherboard serial:", err)
    } else {
        fmt.Println("Motherboard Serial Number:", motherboardSerial)
    }

    cpuID, err := GetCPUID()
    if err != nil {
        fmt.Println("Error retrieving CPU ID:", err)
    } else {
        fmt.Println("CPU ID:", cpuID)
    }

    diskSerial, err := GetDiskSerial()
    if err != nil {
        fmt.Println("Error retrieving Disk Serial:", err)
    } else {
        fmt.Println("Disk Serial Number:", diskSerial)
    }
}
