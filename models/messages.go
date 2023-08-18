package models

// Version data
const Version string = "1.0.0"

// Informational messages
var InfoMessage string = "" +
	"Name:     \033[1;35mTetris\033[0m\n" +
	"Version:  \033[1;36m%s\033[0m\n" +
	"Codename: \033[1;31mYuzu\033[0m\n" +
	"Creator:  \033[1;34mNKTKLN\033[0m\n"

var KeysMessage string = "" +
	"\033[1;34mAll control keys in the game:\033[0m\n\n" +
	"\033[1;37m\"a\"\033[0m or \033[1;37mArrow Left\033[0m  - Moves the figure to the left.\n" +
	"\033[1;37m\"d\"\033[0m or \033[1;37mArrow Right\033[0m - Moves a figure to the right.\n" +
	"\033[1;37m\"s\"\033[0m or \033[1;37mArrow Down\033[0m  - Moves the figure down.\n" +
	"\033[1;37m\"f\"\033[0m or \033[1;37mEnter\033[0m       - Moves the figure to the bottom.\n" +
	"\033[1;37m\"r\"\033[0m or \033[1;37mArrow Up\033[0m    - Rotates the figure.\n" +
	"\033[1;37m\"q\"\033[0m or \033[1;37mESC\033[0m         - Quits the game.\n"

var SizeErrorMessage string = "\033[1;31mTerminal window size is smaller than 20x17!\033[0m"
var StatisticalMessage string = "\033[1;37mScore: \033[1;35m%05d \033[1;37m\nTime:  \033[1;36m%s\033[0m\n"
