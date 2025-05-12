package args

type Args struct {
	Danger  *DangerArgs  `arg:"subcommand:danger|!"`
	Import  *ImportArgs  `arg:"subcommand:import|im" help:"import a plan or group"`
	Init    *InitArgs    `arg:"subcommand:init" help:"create dotato.yaml file"`
	Export  *ExportArgs  `arg:"subcommand:export|ex" help:"export a plan or group"`
	Unlink  *UnlinkArgs  `arg:"subcommand:unlink|un" help:"unlink a plan or group"`
	Version *VersionArgs `arg:"subcommand:version|v" help:"show version"`
	Where   *WhereArgs   `arg:"subcommand:where|w" help:"show component location"`
	// Verbose 		bool					`arg:"-v,--verbose" help:"Verbose output"`
	// Interactive	bool					`arg:"-i,--interactive" help:"Interactive mode"`
}

///////////////////////////////////////////////////////////////////////////////

type DangerArgs struct {
	Unlink *DangerUnlinkArgs `arg:"subcommand:unlink" help:"unlink all links in the current system"`
}
type DangerUnlinkArgs struct {
	Yes      bool `arg:"-y,--yes" help:"skip confirmation"`
	No       bool `arg:"-n,--no" help:"exit on confirmation"`
	FilePerm int  `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int  `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ExportArgs struct {
	Plan  *ExportPlanArgs  `arg:"subcommand:plan|p" help:"export dotfiles in a plan"`
	Group *ExportGroupArgs `arg:"subcommand:group|g" help:"export dotfiles in a group"`
}
type ExportPlanArgs struct {
	Plan     string `arg:"positional" help:"plan name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"exit on confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}
type ExportGroupArgs struct {
	Group    string `arg:"positional" help:"group name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"exit on confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ImportArgs struct {
	Plan  *ImportPlanArgs  `arg:"subcommand:plan|p" help:"Import dotfiles in a plan"`
	Group *ImportGroupArgs `arg:"subcommand:group|g" help:"Import dotfiles in a group"`
}
type ImportPlanArgs struct {
	Plan     string `arg:"positional" help:"plan name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"exit on confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}
type ImportGroupArgs struct {
	Group    string `arg:"positional" help:"group name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"exit on confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type InitArgs struct{}

/////////////////////////////////////////////////////////////////////////////////

type UnlinkArgs struct {
	Plan  *UnlinkPlanArgs  `arg:"subcommand:plan|p"`
	Group *UnlinkGroupArgs `arg:"subcommand:group|g"`
}
type UnlinkPlanArgs struct {
	Plan     string `arg:"positional" help:"plan name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}
type UnlinkGroupArgs struct {
	Group    string `arg:"positional" help:"group name"`
	Resolver string `arg:"positional" help:"resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"skip confirmation"`
	No       bool   `arg:"-n,--no" help:"skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type VersionArgs struct{}

// /////////////////////////////////////////////////////////////////////////////
type WhereArgs struct {
	State *WhereStateArgs `arg:"subcommand:state|s"`
}
type WhereStateArgs struct{}
