package builder

// Builder build the data from vido
type Builder interface {
	Build(video string, params interface{}) ([]interface{}, error)
}

// Implement the build implementation
type Implement string

// Target build target, like cut-frame, calculate flow
type Target string

const (
	// TargetFrame create for frame
	targetFrame Target = "frame"
	// TargetFlow create for flow
	targetFlow Target = "flow"
	// TargetUnknown unknow target
	targetUnknown Target = "unknown"
)

// GetTarget return the target by string
func GetTarget(t string) Target {
	switch t {
	case string(targetFrame):
		return targetFrame
	case string(targetFlow):
		return targetFlow
	default:
		return targetUnknown
	}
}

// IsValid checks target
func (t *Target) IsValid() bool {
	return *t == targetFrame || *t == targetFlow
}

// Pattern generate patterns
type Pattern string

const (
	// PatternRandom cut one frame from the video randomly
	patternRandom Pattern = "random"
	// PatternSample sample small video by interval, then doing random in the
	// small video
	patternSample Pattern = "sample"
	// PatternUnknown unknow target
	patternUnknown Pattern = "unknown"
)

// GetPattern return pattern by string
func GetPattern(p string) Pattern {
	switch p {
	case string(patternRandom):
		return patternRandom
	case string(patternSample):
		return patternSample
	default:
		return patternUnknown
	}
}

// IsValid checks pattern
func (p *Pattern) IsValid() bool {
	return *p == patternRandom || *p == patternSample
}

// NeedInterval checks needness of interval under current the pattern
func (p *Pattern) NeedInterval() bool {
	return *p == patternSample
}

// Params building parameters, it is BAD design
type Params struct {
	Duration int `json:"duration"`
	Interval int `json:"interval"`
}
