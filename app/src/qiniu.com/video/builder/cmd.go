package builder

const (
	// Cmd call external command to build data
	Cmd Implement = "cmd"
)

type frameRandom struct{}

func (f frameRandom) Build(video string,
	params interface{},
) ([]interface{}, error) {
	// TODO
	return nil, nil
}
func (f frameRandom) Clean(result interface{}) error {
	return nil
}

type frameSample struct{}

func (f frameSample) Build(video string,
	params interface{},
) ([]interface{}, error) {
	// TODO
	return nil, nil
}
func (f frameSample) Clean(result interface{}) error {
	return nil
}

type flowRandom struct{}

func (f flowRandom) Build(video string,
	params interface{},
) ([]interface{}, error) {
	// TODO
	return nil, nil
}
func (f flowRandom) Clean(result interface{}) error {
	return nil
}

type flowSample struct{}

func (f flowSample) Build(video string,
	params interface{},
) ([]interface{}, error) {
	// TODO
	return nil, nil
}
func (f flowSample) Clean(result interface{}) error {
	return nil
}

func init() {
	Register(Cmd, targetFrame, patternRandom, frameRandom{})
	Register(Cmd, targetFlow, patternRandom, flowRandom{})
	Register(Cmd, targetFrame, patternRandom, frameSample{})
	Register(Cmd, targetFlow, patternRandom, flowSample{})
}
