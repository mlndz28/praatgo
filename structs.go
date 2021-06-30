package praatgo

// TextGrid objects (contains tiers)
type TextGrid struct {
	FileType 	string			`json:"File type"`
	Xmin     	float64			`json:"xmin"`
	Xmax     	float64			`json:"xmax"`
	Tiers    	bool			`json:"tiers"`
	Size     	int				`json:"size"`
	Item     	[]interface{}	`json:"item"`
}

// A connected sequence of labeled intervals
type IntervalTier struct {
	Class     string		`json:"class"`
	Name      string		`json:"name"`
	Xmin      float64		`json:"xmin"`
	Xmax      float64		`json:"xmax"`
	Size      int			`json:"size"`
	Intervals []Interval	`json:"intervals"`
}


// Transcription for a given time range
type Interval struct {
	Xmin float64	`json:"xmin"`
	Xmax float64	`json:"xmax"`
	Text string		`json:"text"`
}

// A sequence of labeled points
type TextTier struct {
	Class  string	`json:"class"`
	Name   string	`json:"name"`
	Xmin   float64	`json:"xmin"`
	Xmax   float64	`json:"xmax"`
	Size   int		`json:"size"`
	Points []Point	`json:"points"`
}

// An event in time
type Point struct {
	Number float64	`json:"number"`
	Mark   string	`json:"mark"`
}
