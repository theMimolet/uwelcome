package system

import (
	"encoding/json"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ImageInfo struct {
	ImageRef string `json:"image-ref"`
	ImageTag string `json:"image-tag"`
}

func GetDesktop() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if desktop == "" {
		return "Unknown desktop"
	}
	return desktop
}

func GetGreenbootInfo() string {
	cmd_grep := exec.Command("grep", "-q", "status is GREEN", "/etc/motd.d/boot-status")
	err := cmd_grep.Run()
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`status is GREEN`)

	isGreen := re.FindString("status is GREEN")
	if isGreen != "" {
		return "healthy"
	} else {
		cmd := exec.Command("cat", "/etc/motd.d/boot-status")
		output, err := cmd.Output()
		if err != nil {
			return ""
		}
		return "`" + string(output) + "`"
	}
}

// Ublue focused command that retrieves the system image reference from their image-info file
func GetImageInfo() ImageInfo {

	infoFile := "/usr/share/ublue-os/image-info.json"

	data, err := os.ReadFile(infoFile)
	if err != nil {
		return ImageInfo{"", ""}
	}

	var info ImageInfo
	json.Unmarshal(data, &info)

	// strip the ostree prefix, same as the sed in bash
	info.ImageRef = strings.TrimPrefix(info.ImageRef, "ostree-image-signed:docker://")

	return info
}

// Gets the OS name from /etc/os-release
func GetOSName() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`NAME="(.*)"`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		return match[1]
	}
	return "Your System"
}

func IsBootcSystem() bool {
	_, err := os.Stat("/run/ostree-booted")
	return err == nil
}
