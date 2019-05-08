package rule

type baseRule byte

func (x baseRule) StopIfTrue(r *Info) {
	r.rule.StopIfTrue = true
}

func (x baseRule) Validate(r *Info) error {
	return nil
}

