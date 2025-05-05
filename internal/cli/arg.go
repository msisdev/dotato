package cli

type Args struct {
	Danger			*DangerArgs		`arg:"subcommand:danger|!"`
	Import			*ImportArgs		`arg:"subcommand:import|im"`
	Export			*ExportArgs		`arg:"subcommand:export|ex"`
	Unlink 			*UnlinkArgs 	`arg:"subcommand:unlink|un"`
	Version 		*VersionArgs	`arg:"subcommand:version|v" help:"Show version"`
	Verbose 		bool					`arg:"-v,--verbose" help:"Verbose output"`
	Interactive	bool					`arg:"-i,--interactive" help:"Interactive mode"`
}

///////////////////////////////////////////////////////////////////////////////

type DangerArgs struct {
	Unlink	*DangerUnlinkArgs	`arg:"subcommand:unlink"`
}
type DangerUnlinkArgs struct {
	Yes				bool							`arg:"-y,--yes" help:"Skip confirmation"`
}

///////////////////////////////////////////////////////////////////////////////

type ExportArgs struct {
	Plan 		*ExportPlanArgs		`arg:"subcommand:plan|p"`
	Group		*ExportGroupArgs	`arg:"subcommand:group|g"`
}
type ExportPlanArgs struct {
	Plan 			string 						`arg:"positional" help:"Plan name"`
	Resolver	string 						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}
type ExportGroupArgs struct {
	Group 		string 						`arg:"positional" help:"Group name"`
	Resolver	string 						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}

///////////////////////////////////////////////////////////////////////////////

type ImportArgs struct {
	Plan 		*ImportPlanArgs		`arg:"subcommand:plan|p"`
	Group		*ImportGroupArgs	`arg:"subcommand:group|g"`
}
type ImportPlanArgs struct {
	Plan 			string 						`arg:"positional" help:"Plan name"`
	Resolver	string 						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}
type ImportGroupArgs struct {
	Group 		string 						`arg:"positional" help:"Group name"`
	Resolver 	string						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}

///////////////////////////////////////////////////////////////////////////////

type UnlinkArgs struct {
	Plan 		*UnlinkPlanArgs		`arg:"subcommand:plan|p"`
	Group		*UnlinkGroupArgs	`arg:"subcommand:group|g"`
}
type UnlinkPlanArgs struct {
	Plan 			string 						`arg:"positional" help:"Plan name"`
	Resolver	string 						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}
type UnlinkGroupArgs struct {
	Group 		string 						`arg:"positional" help:"Group name"`
	Resolver	string 						`arg:"positional" help:"Resolver name"`
	Yes 			bool							`arg:"-y,--yes" help:"Skip confirmation"`
}

///////////////////////////////////////////////////////////////////////////////

type VersionArgs struct {}
