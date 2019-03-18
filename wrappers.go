package lvm

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	// LVMPath is the path to the "lvm" command.
	LVMPath string
)

func init() {
	if p, err := exec.LookPath("lvm"); err == nil {
		LVMPath = p
	}
}

func runWithoutOutput(cmdPath string, args ...string) error {
	logrus.Debugf("running %v", append([]string{cmdPath}, args...))
	cmd := exec.Command(cmdPath, args...)
	stderr := bytes.Buffer{}
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return errors.Errorf("%s", stderr.String())
	}
	return nil
}

func runWithOutput(cmdPath string, args ...string) (string, error) {
	logrus.Debugf("running %v", append([]string{cmdPath}, args...))
	cmd := exec.Command(cmdPath, args...)
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout
	stderr := bytes.Buffer{}
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return stdout.String(), errors.Errorf("%s", stderr.String())
	}
	return stdout.String(), nil
}

// getPhysicalVolumes returns information about known physical volumes or a
// specific physical volume.
func getPhysicalVolumes(pvname string) (Report, error) {
	report := Report{}
	b := []byte{}
	if pvname != "" {
		raw, err := runWithOutput(LVMPath, "pvs", "--reportformat", "json", "--units", "b", "--nosuffix", pvname)
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvs pvs\" for %q", pvname)
		}
		b = []byte(raw)
	} else {
		raw, err := runWithOutput(LVMPath, "pvs", "--reportformat", "json", "--units", "b", "--nosuffix")
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvs pvs\"")
		}
		b = []byte(raw)
	}
	err := json.Unmarshal(b, &report)
	if err != nil {
		return Report{}, errors.Wrapf(err, "error decoding output from \"lvs pvs\"")
	}
	return report, nil
}

// getVolumeGroups returns information about the known volume groups, or about
// a specific volume group.
func getVolumeGroups(vgname string) (Report, error) {
	report := Report{}
	b := []byte{}
	if vgname != "" {
		raw, err := runWithOutput(LVMPath, "vgs", "--reportformat", "json", "--units", "b", "--nosuffix", vgname)
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvs vgs\" for %q", vgname)
		}
		b = []byte(raw)
	} else {
		raw, err := runWithOutput(LVMPath, "vgs", "--all", "--reportformat", "json", "--units", "b", "--nosuffix")
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvs vgs\"")
		}
		b = []byte(raw)
	}
	err := json.Unmarshal(b, &report)
	if err != nil {
		return Report{}, errors.Wrapf(err, "error decoding output from \"lvs vgs\"")
	}
	return report, nil
}

// getVolumeGroupsFull returns detailed information about all known volume
// groups, or about one specific volume group.
func getVolumeGroupsFull(vgname string) (ReportFull, error) {
	report := ReportFull{}
	b := []byte{}
	if vgname != "" {
		raw, err := runWithOutput(LVMPath, "fullreport", "--reportformat", "json", "--units", "b", "--nosuffix", vgname)
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvm fullreport\" for %q", vgname)
		}
		b = []byte(raw)
	} else {
		raw, err := runWithOutput(LVMPath, "fullreport", "--reportformat", "json", "--units", "b", "--nosuffix")
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvm fullreport\"")
		}
		b = []byte(raw)
	}
	err := json.Unmarshal(b, &report)
	if err != nil {
		return ReportFull{}, errors.Wrapf(err, "error decoding output from \"lvm fullreport\"")
	}
	return report, nil
}

// getLogicalVolumes returns information about all known logical volumes, about
// the volumes in a specified volume group, or about a specific volume.
func getLogicalVolumes(vgname, volume string) (Report, error) {
	report := Report{}
	b := []byte{}
	if vgname != "" {
		if volume != "" {
			raw, err := runWithOutput(LVMPath, "lvs", "--all", "--reportformat", "json", "--units", "b", "--nosuffix", vgname+"/"+volume)
			if err != nil {
				return report, errors.Wrapf(err, "error running \"lvm lvs\" for %q", vgname+"/"+volume)
			}
			b = []byte(raw)
		} else {
			raw, err := runWithOutput(LVMPath, "lvs", "--all", "--reportformat", "json", "--units", "b", "--nosuffix", vgname)
			if err != nil {
				return report, errors.Wrapf(err, "error running \"lvm lvs\" for %q", vgname)
			}
			b = []byte(raw)
		}
	} else {
		raw, err := runWithOutput(LVMPath, "lvs", "--all", "--reportformat", "json", "--units", "b", "--nosuffix")
		if err != nil {
			return report, errors.Wrapf(err, "error running \"lvm lvs\"")
		}
		b = []byte(raw)
	}
	err := json.Unmarshal(b, &report)
	if err != nil {
		return Report{}, errors.Wrapf(err, "error decoding output from \"lvm lvs\"")
	}
	return report, nil
}

// physicalVolumeIsPresent checks if a physical volume with the specified name
// exists.  Force a rescan of that device for physical volume header data, for
// cases where we've just attached it.
func physicalVolumeIsPresent(pvname string) bool {
	scanned, err := runWithOutput(LVMPath, "pvscan", "--cache", pvname)
	if err != nil {
		logrus.Debugf("lvm pvscan failed: %q", scanned)
		return false
	}
	checked, err := runWithOutput(LVMPath, "pvck", pvname)
	if err != nil {
		logrus.Debugf("lvm pvck failed: %q", checked)
		return false
	}
	return true
}

// volumeNameForID converts an ID into a reasonable volume name.
func volumeNameForID(ID string) string {
    return "layer." + ID
}

// volumePathForID determines the device pathname for a volume with the
// specified ID in a particular volume group, or across all volume groups.
func volumePathForID(vgname, id string) (string, error) {
	lvname := volumeNameForID(id)
	report, err := getVolumeGroupsFull(vgname)
	if err != nil {
		return "", errors.WithStack(err)
	}
	for _, entry := range report.Reports {
		if vgname != "" {
			foundVG := false
			for _, vg := range entry.VGs {
				if vg.Name == vgname {
					foundVG = true
					break
				}
			}
			if !foundVG {
				continue
			}
		}
		for _, lv := range entry.LVs {
			if lv.Name == lvname {
				if lv.DMPath != "" {
					_, err = os.Stat(lv.DMPath)
					if err == nil {
						return lv.DMPath, nil
					}
				}
				if lv.Path != "" {
					_, err = os.Stat(lv.Path)
					if err == nil {
						return lv.Path, nil
					}
				}
				return "", errors.Errorf("found LV %q, but no active path for it", vgname+"/"+lv.Name)
			}
		}
	}
	return "", errors.Errorf("no LV named %q found", vgname+"/"+lvname)
}

// readVolumeGroupForPhysicalVolume will determine the name of the volume group
// to which the specified physical volume belongs.
func readVolumeGroupForPhysicalVolume(pvname string) (string, error) {
	report, err := getPhysicalVolumes(pvname)
	if err != nil {
		return "", errors.WithStack(err)
	}
	for _, entry := range report.Reports {
		for _, pv := range entry.PVs {
			if pv.Name == pvname {
				return pv.VGName, nil
			}
		}
	}
	return "", errors.Errorf("no PV named %q found", pvname)
}

// volumeGroupIsPresent checks if a volume group with the specified name exists.
func volumeGroupIsPresent(vgname string) bool {
	scanned, err := runWithOutput(LVMPath, "vgscan", "--cache")
	if err != nil {
		logrus.Debugf("lvm vgscan failed for %q: %q", vgname, scanned)
		return false
	}
	scanned, err = runWithOutput(LVMPath, "vgs", "--reportformat", "json", "--units", "b", "--nosuffix", vgname)
	if err != nil {
		logrus.Debugf("lvm vgs failed for %q: %q", vgname, scanned)
		return false
	}
	return true
}

// getLogicalVolume returns information about the specified logical volume.
func getLogicalVolume(vgname, volume string) (ReportLV, error) {
	report, err := getLogicalVolumes(vgname, volume)
	if err != nil {
		return ReportLV{}, errors.WithStack(err)
	}
	for _, entry := range report.Reports {
		for _, lv := range entry.LVs {
			if lv.Name == volume {
				return lv, nil
			}
		}
	}
	return ReportLV{}, errors.Errorf("no LV named %q found", volume)
}

// logicalVolumeIsPresent checks if a logical volume with the specified name in
// the specified volume group exists.
func logicalVolumeIsPresent(vgname, volume string) bool {
	scanned, err := runWithOutput(LVMPath, "lvscan", "--cache", vgname+"/"+volume)
	if err != nil {
		logrus.Debugf("lvm lvscan failed: %q", scanned)
		return false
	}
	return true
}

// createPhysicalVolume formats a specified device as a physical volume.
func createPhysicalVolume(device string) error {
	err := runWithoutOutput(LVMPath, "pvcreate", device)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm pvcreate\" for %q", device)
	}
	return nil
}

// resizePhysicalVolume tells the kernel that the loopback device may be larger
// now, so the volume group that its in will care about that.
func resizePhysicalVolume(device string) error {
	output, err := runWithOutput(LVMPath, "pvresize", device)
	output = strings.TrimRight(output, "\r\n\t ")
	if err != nil {
		return errors.Wrapf(err, "error checking if device %q has been resized: %q", device, output)
	}
	return nil
}

// createPhysicalVolume formats a specified device as a physical volume.
func createVolumeGroup(vgname string, device ...string) error {
	err := runWithoutOutput(LVMPath, append([]string{"vgcreate", vgname}, device...)...)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm vgcreate\" for %v", device)
	}
	return nil
}

// activateVolumeGroup activates the specified volume group, making all of its
// logical volumes visible.
func activateVolumeGroup(vgname string) error {
	err := runWithoutOutput(LVMPath, "vgchange", "--activate", "y", "--ignoreactivationskip", vgname)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm vgchange --activate y\" for %q", vgname)
	}
	return nil
}

// deactivateVolumeGroup deactivates the specified volume group, making all of
// its logical volumes invisible.
func deactivateVolumeGroup(vgname string) error {
	err := runWithoutOutput(LVMPath, "vgchange", "--activate", "n", vgname)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm vgchange --activate n\" for %q", vgname)
	}
	return nil
}

// activateLogicalVolume activates a single logical volume in the specified
// volume group, making it visible.
func activateLogicalVolume(vgname, volume string) error {
	err := runWithoutOutput(LVMPath, "lvchange", "--activate", "y", "--ignoreactivationskip", vgname+"/"+volume)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm lvchange --activate y\" for %q", vgname+"/"+volume)
	}
	return nil
}

// deactivateLogicalVolume deactivates a single logical volume in the specified
// volume group, making it invisible.
func deactivateLogicalVolume(vgname, volume string) error {
	err := runWithoutOutput(LVMPath, "lvchange", "--activate", "n", vgname+"/"+volume)
	if err != nil {
		return errors.Wrapf(err, "error running \"lvm lvchange --activate n\" for %q", vgname+"/"+volume)
	}
	return nil
}

// read information about the active thin pool
func readPoolInfo(vgname, poolname string) (lvmPoolHistory, error) {
	report, err := getVolumeGroupsFull(vgname)
	if err != nil {
		return lvmPoolHistory{}, errors.Wrapf(err, "error reading information about volume group %q", vgname)
	}
	for _, entry := range report.Reports {
		if vgname != "" {
			foundVG := false
			for _, vg := range entry.VGs {
				if vg.Name == vgname {
					foundVG = true
					break
				}
			}
			if !foundVG {
				continue
			}
		}
		for _, lv := range entry.LVs {
			if lv.Name != poolname {
				continue
			}
			history := lvmPoolHistory{
				VGname:   vgname,
				PoolName: lv.Name,
				PoolUUID: lv.UUID,
			}
			return history, nil
		}
	}
	return lvmPoolHistory{}, errors.Errorf("unable to locate information about pool %q in volume group %q", poolname, vgname)
}
