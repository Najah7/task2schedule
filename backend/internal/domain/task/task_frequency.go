package task

import "errors"

var (
	ErrTaskFrequencyEmpty   = errors.New("task frequency cannot be empty")
	ErrTaskFrequencyInvalid = errors.New("invalid task frequency")
)

type TaskFrequency struct {
	Value   string
	Label   string
	LabelJp string
}

var taskFrequencies = map[string]TaskFrequency{
	"mon":  {Value: "mon", Label: "Monday", LabelJp: "月曜日"},
	"tue":  {Value: "tue", Label: "Tuesday", LabelJp: "火曜日"},
	"wed":  {Value: "wed", Label: "Wednesday", LabelJp: "水曜日"},
	"thu":  {Value: "thu", Label: "Thursday", LabelJp: "木曜日"},
	"fri":  {Value: "fri", Label: "Friday", LabelJp: "金曜日"},
	"sat":  {Value: "sat", Label: "Saturday", LabelJp: "土曜日"},
	"sun":  {Value: "sun", Label: "Sunday", LabelJp: "日曜日"},
	"once": {Value: "once", Label: "Once", LabelJp: "一度限り"},
}

func NewTaskFrequency(frequency string) (TaskFrequency, error) {
	f, ok := taskFrequencies[frequency]
	if !ok {
		f = TaskFrequency{Value: frequency}
	}
	if err := f.validate(); err != nil {
		return TaskFrequency{}, err
	}
	return f, nil
}

func (f TaskFrequency) validate() error {
	if f.String() == "" {
		return ErrTaskFrequencyEmpty
	}

	if _, ok := taskFrequencies[f.Value]; ok {
		return nil
	}
	return ErrTaskFrequencyInvalid
}

func (f TaskFrequency) String() string {
	return f.Value
}
