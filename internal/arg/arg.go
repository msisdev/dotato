package arg

type Args struct {
	Danger	*DangerArgs		`arg:"subcommand:danger|!"`
	Plan		*PlanArgs			`arg:"subcommand:plan|p"`
	Group		*GroupArgs		`arg:"subcommand:group|g"`
	File		*FileArgs			`arg:"subcommand:file|f"`
	Version *VersionArgs	`arg:"subcommand:version|v" help:"Show version"`
	Verbose bool					`arg:"-v,--verbose" help:"Verbose output"`
}

type DangerArgs struct {
	Unlink *DangerUnlinkArgs	`arg:"subcommand:unlink"`
}
type DangerUnlinkArgs struct {
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type PlanArgs struct {
	In *PlanInArgs	`arg:"subcommand:in"`
	Out *PlanOutArgs	`arg:"subcommand:out"`
	Tidy *PlanTidyArgs	`arg:"subcommand:tidy"`
}
type PlanInArgs struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type PlanOutArgs struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type PlanTidyArgs struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type GroupArgs struct {
	In *GroupInArgs	`arg:"subcommand:in"`
	Out *GroupOutArgs	`arg:"subcommand:out"`
	Tidy *GroupTidyArgs	`arg:"subcommand:tidy"`
}
type GroupInArgs struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type GroupOutArgs struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type GroupTidyArgs struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type FileArgs struct {
	In *FileInArgs	`arg:"subcommand:in"`
	Out *FileOutArgs	`arg:"subcommand:out"`
	Move *FileMoveArgs	`arg:"subcommand:move"`
}
type FileInArgs struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type FileOutArgs struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type FileMoveArgs struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type VersionArgs struct {

}