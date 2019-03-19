package lvm

// ReportLoopback represents information about a configured loopback device,
// produced by the "losetup --list" command.
type ReportLoopback struct {
	Name      string `json:"name"`
	SizeLimit int64  `json:"sizelimit,string"`
	Offset    int64  `json:"offset,string"`
	AutoClear int64  `json:"autoclear,string"`
	ReadOnly  int64  `json:"ro,string"`
	File      string `json:"back-file"` // may end with " (deleted)"
	DIO       int64  `json:"dio,string"`
}

// ReportPVCommon represents information that's common about multiple report formats.
type ReportPVCommon struct {
	Name       string `json:"pv_name"`
	Attributes string `json:"pv_attr"`
	Format     string `json:"pv_fmt"`
	Size       int64  `json:"pv_size,string"`
	Free       int64  `json:"pv_free,string"`
}

// ReportPV represents the information about a physical volume that is produced
// by the "lvm pvs" command.
type ReportPV struct {
	ReportPVCommon
	VGName string `json:"vg_name"`
}

// ReportPVFull represents the information about a physical volume that is
// produced by the "lvm fullreport" command.
type ReportPVFull struct {
	ReportPVCommon
	Format        string `json:"pv_fmt"`
	UUID          string `json:"pv_uuid"`
	DeviceSize    int64  `json:"dev_size,string"`
	Major         int64  `json:"pv_major,string"`
	Minor         int64  `json:"pv_minor,string"`
	MDAFree       int64  `json:"pv_mda_free,string"`
	MDASize       int64  `json:"pv_mda_size,string"`
	ExtVersion    int64  `json:"pv_ext_vsn,string"`
	ExtStart      int64  `json:"pe_start,string"`
	Used          int64  `json:"pv_used,string"`
	Allocatable   string `json:"pv_allocatable"`
	Exported      string `json:"pv_exported"`
	Missing       string `json:"pv_missing"`
	ExtCount      int64  `json:"pv_pe_count,string"`
	ExtAllocCount int64  `json:"pv_pe_alloc_count,string"`
	Tags          string `json:"pv_tags"`
	MDACount      int64  `json:"pv_mda_count,string"`
	MDAUsedCount  int64  `json:"pv_mda_used_count,string"`
	BAStart       int64  `json:"pv_ba_start,string"`
	BASize        int64  `json:"pv_ba_size,string"`
	InUse         string `json:"pv_in_use"`
	Duplicate     string `json:"pv_duplicate"`
}

// ReportVGCommon represents information that's common about multiple report formats.
type ReportVGCommon struct {
	Name       string `json:"vg_name"`
	PVCount    int64  `json:"pv_count,string"`
	LVCount    int64  `json:"lv_count,string"`
	SnapCount  int64  `json:"snap_count,string"`
	Attributes string `json:"vg_attr"`
	Size       int64  `json:"vg_size,string"`
	Free       int64  `json:"vg_free,string"`
}

// ReportVG represents the information about a volume group that is produced by
// the "lvm vgs" command.
type ReportVG struct {
	ReportVGCommon
}

// ReportVGFull represents the information about a volume group that is
// produced by the "lvm fullreport" command.
type ReportVGFull struct {
	ReportVGCommon
	Format           string `json:"vg_fmt"`
	UUID             string `json:"vg_uuid"`
	Permissions      string `json:"vg_permissions"`
	Extendable       string `json:"vg_extendable"`
	Exported         string `json:"vg_exported"`
	Partial          string `json:"vg_partial"`
	AllocationPolicy string `json:"vg_allocation_policy"`
	Clustered        string `json:"vg_clustered"`
	SysID            string `json:"vg_sysid"`
	SystemID         string `json:"vg_systemid"`
	LockType         string `json:"vg_locktype"`
	LockArgs         string `json:"vg_lockargs"`
	ExtentSize       int64  `json:"vg_extent_size,string"`
	ExtentCount      int64  `json:"vg_extent_count,string"`
	FreeCount        int64  `json:"vg_free_count,string"`
	MaxLV            int64  `json:"max_lv,string"`
	MaxPV            int64  `json:"max_pv,string"`
	MissingPVCount   int64  `json:"vg_missing_pv_count,string"`
	LVCount          int64  `json:"lv_count,string"`
	SnapCount        int64  `json:"snap_count,string"`
	SequenceNumber   int64  `json:"vg_seqno,string"`
	Tags             string `json:"vg_tags"`
	VGProfile        string `json:"vg_profile"`
	MDACount         int64  `json:"vg_mda_count,string"`
	MDAUsedCount     int64  `json:"vg_mda_used_count,string"`
	MDAFree          int64  `json:"vg_mda_free,string"`
	MDASize          int64  `json:"vg_mda_size,string"`
	MDACopies        string `json:"vg_mda_copies"`
}

// ReportLVCommon represents information that's common about multiple report formats.
type ReportLVCommon struct {
	Name            string `json:"lv_name"`
	Attributes      string `json:"lv_attr"`
	Size            int64  `json:"lv_size,string"`
	Pool            string `json:"pool_lv"`
	Origin          string `json:"origin"`
	DataPercent     string `json:"data_percent"`
	MetadataPercent string `json:"metadata_percent"`
	MovePV          string `json:"move_pv"`
	MirrorLog       string `json:"mirror_log"`
	CopyPercent     string `json:"copy_percent"`
	ConvertLV       string `json:"convert_lv"`
}

// ReportLV represents the information about a logical volume that is produced
// by the "lvm lvs" command.
type ReportLV struct {
	ReportLVCommon
	VGName string `json:"vg_name"`
}

// ReportLVFull represents the information about a logical volume that is
// produced by the "lvm fullreport" command.
type ReportLVFull struct {
	ReportLVCommon
	UUID                string `json:"lv_uuid"`
	FullName            string `json:"lv_full_name"`
	Path                string `json:"lv_path"`
	DMPath              string `json:"lv_dm_path"`
	Parent              string `json:"lv_parent"`
	Layout              string `json:"lv_layout"`
	Role                string `json:"lv_role"`
	InitialImageSync    string `json:"lv_initial_image_sync"`
	ImageSynced         string `json:"lv_image_synced"`
	Merging             string `json:"lv_merging"`
	Converting          string `json:"lv_converting"`
	AllocationPolicy    string `json:"lv_allocation_policy"`
	AllocationLocked    string `json:"lv_allocation_locked"`
	FixedMinor          string `json:"lv_fixed_minor"`
	MergeFailed         string `json:"lv_merge_failed"`
	SnapshotInvalid     string `json:"lv_snapshot_invalid"`
	SkipActivation      string `json:"lv_skip_activation"`
	WhenFull            string `json:"lv_when_full"`
	Active              string `json:"lv_active"`
	ActiveLocally       string `json:"lv_active_locally"`
	ActiveRemotely      string `json:"lv_active_remotely"`
	ActiveExclusively   string `json:"lv_active_exclusively"`
	Major               int64  `json:"lv_major,string"`
	Minor               int64  `json:"lv_minor,string"`
	ReadAhead           string `json:"lv_read_ahead"`
	MetadataSize        string `json:"lv_metadata_size"`
	SegmentCount        int64  `json:"seg_count,string"`
	Origin              string `json:"origin"`
	OriginUUID          string `json:"origin_uuid"`
	OriginSize          string `json:"origin_size"`
	Ancestors           string `json:"lv_ancestors"`
	FullAncestors       string `json:"lv_full_ancestors"`
	Descendants         string `json:"lv_descendants"`
	FullDescendants     string `json:"lv_full_descendants"`
	DataPercent         string `json:"data_percent"`
	SnapPercent         string `json:"snap_percent"`
	MetadataPercent     string `json:"metadata_percent"`
	CopyPercent         string `json:"copy_percent"`
	SyncPercent         string `json:"sync_percent"`
	RAIDMismatchCount   string `json:"raid_mismatch_count"`
	RAIDSyncAction      string `json:"raid_sync_action"`
	RAIDWriteBehind     string `json:"raid_write_behind"`
	RAIDMinRecoveryRate string `json:"raid_min_recovery_rate"`
	RAIDMaxRecoveryRate string `json:"raid_max_recovery_rate"`
	MovePV              string `json:"move_pv"`
	MovePVUUID          string `json:"move_pv_uuid"`
	ConvertLV           string `json:"convert_lv"`
	ConvertLVUUID       string `json:"convert_lv_uuid"`
	MirrorLog           string `json:"mirror_log"`
	MirrorLogUUID       string `json:"mirror_log_uuid"`
	DataLV              string `json:"data_lv"`
	DataLVUUID          string `json:"data_lv_uuid"`
	MetadataLV          string `json:"metadata_lv"`
	MetadataLVUUID      string `json:"metadata_lv_uuid"`
	PoolLV              string `json:"pool_lv"`
	PoolLVUUID          string `json:"pool_lv_uuid"`
	Tags                string `json:"lv_tags"`
	Profile             string `json:"lv_profile"`
	LockArgs            string `json:"lv_lockargs"`
	Time                string `json:"lv_time"`
	TimeRemoved         string `json:"lv_time_removed"`
	Host                string `json:"lv_host"`
	Modules             string `json:"lv_modules"`
	Historical          string `json:"lv_historical"`
	KernelMajor         int64  `json:"lv_kernel_major,string"`
	KernelMinor         int64  `json:"lv_kernel_minor,string"`
	KernelReadAhead     int64  `json:"lv_kernel_read_ahead,string"`
	Permissions         string `json:"lv_permissions"`
	Suspended           string `json:"lv_suspended"`
	LiveTable           string `json:"lv_live_table"`
	InactiveTable       string `json:"lv_inactive_table"`
	DeviceOpen          string `json:"lv_device_open"`
	CacheTotalBlocks    string `json:"cache_total_blocks"`
	CacheUsedBlocks     string `json:"cache_used_blocks"`
	CacheDirtyBlocks    string `json:"cache_dirty_blocks"`
	CacheReadHits       string `json:"cache_read_hits"`
	CacheReadMisses     string `json:"cache_read_misses"`
	CacheWriteHits      string `json:"cache_write_hits"`
	CacheWriteMisses    string `json:"cache_write_misses"`
	KernelCacheSettings string `json:"kernel_cache_settings"`
	KernelCachePolicy   string `json:"kernel_cache_policy"`
	HealthStatus        string `json:"lv_health_status"`
	KernelDiscards      string `json:"kernel_discards"`
	CheckNeeded         string `json:"lv_check_needed"`
}

// ReportEntry represents part of the information about local storage that is
// produced by any of the "lvm vgs", "lvm pvs", or "lvm lvs" command.
type ReportEntry struct {
	PVs []ReportPV `json:"pv"`
	VGs []ReportVG `json:"vg"`
	LVs []ReportLV `json:"lv"`
}

// ReportEntryFull represents the information specific to a local volume group
// that is produced by the "lvm fullreport" command.
type ReportEntryFull struct {
	PVs []ReportPVFull `json:"pv"`
	VGs []ReportVGFull `json:"vg"`
	LVs []ReportLVFull `json:"lv"`
}

// Report represents the information about local storage that is reported by
// any of the "lvm vgs", "lvm pvs", or "lvm lvs" commands, or by the "losetup"
// command.
type Report struct {
	Reports  []ReportEntry    `json:"report,omitempty"`
	Loopback []ReportLoopback `json:"loopdevices,omitempty"`
}

// ReportFull represents the information about local storage that is produced
// by the "lvm fullreport" command.
type ReportFull struct {
	Reports []ReportEntryFull `json:"report,omitempty"`
}

// this is our record of the last pool that we used, along with its UUID.  We
// need to error out if it's changed, because the set of layers the pool has
// likely no longer matches what higher level APIs think we have.
type LvmPoolHistory struct {
    VGname   string `json:"vgname"`
    PoolName string `json:"poolname"`
    PoolUUID string `json:"uuid"`
}

