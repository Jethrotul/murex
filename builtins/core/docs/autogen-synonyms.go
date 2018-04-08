package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!event`:          `event`,
	`echo`:            `out`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`if`:              `if`,
	`trypipe`:         `trypipe`,
	`>>`:              `>>`,
	`ttyfd`:           `ttyfd`,
	`unset`:           `unset`,
	`get`:             `get`,
	`post`:            `post`,
	`err`:             `err`,
	`g`:               `g`,
	`alter`:           `alter`,
	`swivel-datatype`: `swivel-datatype`,
	`out`:             `out`,
	`f`:               `f`,
	`rx`:              `rx`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`>`:               `>`,
	`catch`:           `catch`,
	`set`:             `set`,
	`append`:          `append`,
	`prepend`:         `prepend`,
	`try`:             `try`,
	`swivel-table`:    `swivel-table`,
	`print`:           `print`,
	`tout`:            `tout`,
	`pt`:              `pt`,
}
