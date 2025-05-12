package args

type Args struct {
	Danger			*DangerArgs		`arg:"subcommand:danger|!"`
	Import			*ImportArgs		`arg:"subcommand:import|im"`
	Init				*InitArgs			`arg:"subcommand:init"`
	Export			*ExportArgs		`arg:"subcommand:export|ex"`
	Unlink 			*UnlinkArgs 	`arg:"subcommand:unlink|un"`
	Version 		*VersionArgs	`arg:"subcommand:version|v" help:"Show version"`
	Where				*WhereArgs		`arg:"subcommand:where|w" help:"Show component location"`
	// Verbose 		bool					`arg:"-v,--verbose" help:"Verbose output"`
	// Interactive	bool					`arg:"-i,--interactive" help:"Interactive mode"`
}

///////////////////////////////////////////////////////////////////////////////

type DangerArgs struct {
	Unlink	*DangerUnlinkArgs	`arg:"subcommand:unlink"`
}
type DangerUnlinkArgs struct {
	Yes				bool	`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm	int 	`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 	`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ExportArgs struct {
	Plan 		*ExportPlanArgs		`arg:"subcommand:plan|p"`
	Group		*ExportGroupArgs	`arg:"subcommand:group|g"`
}
type ExportPlanArgs struct {
	Plan 			string 	`arg:"positional" help:"Plan name"`
	Resolver	string 	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}
type ExportGroupArgs struct {
	Group 		string	`arg:"positional" help:"Group name"`
	Resolver	string 	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type ImportArgs struct {
	Plan 		*ImportPlanArgs		`arg:"subcommand:plan|p"`
	Group		*ImportGroupArgs	`arg:"subcommand:group|g"`
}
type ImportPlanArgs struct {
	Plan 			string 	`arg:"positional" help:"Plan name"`
	Resolver	string 	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}
type ImportGroupArgs struct {
	Group 		string	`arg:"positional" help:"Group name"`
	Resolver 	string	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type InitArgs struct {}

/////////////////////////////////////////////////////////////////////////////////

type UnlinkArgs struct {
	Plan 		*UnlinkPlanArgs		`arg:"subcommand:plan|p"`
	Group		*UnlinkGroupArgs	`arg:"subcommand:group|g"`
}
type UnlinkPlanArgs struct {
	Plan 			string	`arg:"positional" help:"Plan name"`
	Resolver	string 	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}
type UnlinkGroupArgs struct {
	Group 		string 	`arg:"positional" help:"Group name"`
	Resolver	string 	`arg:"positional" help:"Resolver name"`
	Yes 			bool		`arg:"-y,--yes" help:"Skip confirmation"`
	No 				bool	`arg:"-n,--no" help:"Skip confirmation"`
	FilePerm 	int 		`arg:"-f,--file-perm" help:"File permission" default:"0644"`
	DirPerm  	int 		`arg:"-d,--dir-perm" help:"Directory permission" default:"0755"`
}

///////////////////////////////////////////////////////////////////////////////

type VersionArgs struct {}

///////////////////////////////////////////////////////////////////////////////
type WhereArgs struct {
	State *WhereStateArgs `arg:"subcommand:state|s"`
}
type WhereStateArgs struct {}
