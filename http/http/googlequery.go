package http

// unused so far ...
type query struct {
	search string //required
	Title  string //optional
	tbm    string // =isch is image search ... others??? or is this binary???
	safe   bool   // safe=active,???
	hl     string // language?
	biw    int    // 'b' image width???
	bih    int    // 'b' image height???
	// color ... &tbs=ic:specific,isc:red
	// date range ... tbs=cdr:1,cd_min:4/1/2013,cd_max:5/1/2013
	// size ... tbs=isz:l
}
