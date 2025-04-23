package arg

// import (
// 	"github.com/alexflint/go-arg"
// )

type Args struct {
	Apply		*ApplyCmd	`arg:"subcommand:apply"`
	In			*InCmd		`arg:"subcommand:in"`
	Init		*InitCmd	`arg:"subcommand:init"`
	Move		*MoveCmd	`arg:"subcommand:move|mv"`
	New			*NewCmd		`arg:"subcommand:new"`
	Out			*OutCmd		`arg:"subcommand:out"`
	Plan		*PlanCmd	`arg:"subcommand:plan"`
	Tidy		*TidyCmd	`arg:"subcommand:tidy"`
}

type ApplyCmd struct {
	Plan	string	`arg:"positional"`
}

type InCmd struct {
	
}

type InitCmd struct {

}

type MoveCmd struct {

}

type NewCmd struct {

}

type OutCmd struct {

}

type PlanCmd struct {

}

type TidyCmd struct {

}
