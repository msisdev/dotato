package args

type Args struct {
	Danger  *DangerArgs  `arg:"subcommand:danger|!"`
	Import  *ImportArgs  `arg:"subcommand:import|im" help:"Import a plan or group"`
	Init    *InitArgs    `arg:"subcommand:init" help:"Create dotato.yaml file"`
	Export  *ExportArgs  `arg:"subcommand:export|ex" help:"Export a plan or group"`
	Unlink  *UnlinkArgs  `arg:"subcommand:unlink|un" help:"Unlink a plan or group"`
	Version *VersionArgs `arg:"subcommand:version|v" help:"Show version"`
	Where   *WhereArgs   `arg:"subcommand:where|w" help:"Show component location"`
	// Verbose 		bool					`arg:"-v,--verbose" help:"Verbose output"`
	// Interactive	bool					`arg:"-i,--interactive" help:"Interactive mode"`
}

///////////////////////////////////////////////////////////////////////////////

type DangerArgs struct {
	Unlink *DangerUnlinkArgs `arg:"subcommand:unlink" help:"Unlink all links in the current system"`
}
type DangerUnlinkArgs struct {
	Yes      bool `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool `arg:"-n,--no" help:"Exit on confirmation"`
	FilePerm int  `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int  `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ExportArgs struct {
	Plan  *ExportPlanArgs  `arg:"subcommand:plan|p" help:"Export dotfiles in a plan"`
	Group *ExportGroupArgs `arg:"subcommand:group|g" help:"Export dotfiles in a group"`
}
type ExportPlanArgs struct {
	Plan     string `arg:"positional" help:"Plan name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Exit on confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}
type ExportGroupArgs struct {
	Group    string `arg:"positional" help:"Group name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ImportArgs struct {
	Plan  *ImportPlanArgs  `arg:"subcommand:plan|p" help:"Import dotfiles in a plan"`
	Group *ImportGroupArgs `arg:"subcommand:group|g" help:"Import dotfiles in a group"`
}
type ImportPlanArgs struct {
	Plan     string `arg:"positional" help:"Plan name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}
type ImportGroupArgs struct {
	Group    string `arg:"positional" help:"Group name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type InitArgs struct{}

/////////////////////////////////////////////////////////////////////////////////

type UnlinkArgs struct {
	Plan  *UnlinkPlanArgs  `arg:"subcommand:plan|p"`
	Group *UnlinkGroupArgs `arg:"subcommand:group|g"`
}
type UnlinkPlanArgs struct {
	Plan     string `arg:"positional" help:"Plan name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}
type UnlinkGroupArgs struct {
	Group    string `arg:"positional" help:"Group name"`
	Resolver string `arg:"positional" help:"Resolver name"`
	Yes      bool   `arg:"-y,--yes" help:"Skip confirmation"`
	No       bool   `arg:"-n,--no" help:"Skip confirmation"`
	FilePerm int    `arg:"-f,--file-perm" help:"Create file with permission" default:"0644"`
	DirPerm  int    `arg:"-d,--dir-perm" help:"Create directory with permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type VersionArgs struct{}

// /////////////////////////////////////////////////////////////////////////////
type WhereArgs struct {
	State *WhereStateArgs `arg:"subcommand:state|s"`
}
type WhereStateArgs struct{}
