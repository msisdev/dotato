package arg

type Args struct {
	Danger	*DangerCmd	`arg:"subcommand:!"`
	Plan		*PlanCmd		`arg:"subcommand:plan|p"`
	Group		*GroupCmd		`arg:"subcommand:group|g"`
	File		*FileCmd		`arg:"subcommand:file|f"`
	Version bool				`arg:"-v,--version" help:"Show version"`
}


type DangerCmd struct {
	Unlink *DangerUnlinkCmd	`arg:"subcommand:unlink"`
}
type DangerUnlinkCmd struct {
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type PlanCmd struct {
	In *PlanInCmd	`arg:"subcommand:in"`
	Out *PlanOutCmd	`arg:"subcommand:out"`
	Tidy *PlanTidyCmd	`arg:"subcommand:tidy"`
}
type PlanInCmd struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type PlanOutCmd struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type PlanTidyCmd struct {
	Plan string `arg:"positional" help:"Plan name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type GroupCmd struct {
	In *GroupInCmd	`arg:"subcommand:in"`
	Out *GroupOutCmd	`arg:"subcommand:out"`
	Tidy *GroupTidyCmd	`arg:"subcommand:tidy"`
}
type GroupInCmd struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type GroupOutCmd struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type GroupTidyCmd struct {
	Group string `arg:"positional" help:"Group name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}


type FileCmd struct {
	In *FileInCmd	`arg:"subcommand:in"`
	Out *FileOutCmd	`arg:"subcommand:out"`
	Move *FileMoveCmd	`arg:"subcommand:move"`
}
type FileInCmd struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type FileOutCmd struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
type FileMoveCmd struct {
	File string `arg:"positional" help:"File name"`
	Yes bool	`arg:"-y,--yes" help:"Skip confirmation"`
}
