package pipeline

// Starter the handler starter
type Starter struct {
	Handler
}

func (p *Starter) Execute() {

	nextContext := KutkhData{}
	nextContext.param = p.context.param

	for p.nextPipeline != nil {
		if nextContext.err != nil {
			continue
		}
		if p.nextPipeline.method != nil {
			resultContext := p.nextPipeline.method(nextContext)
			nextContext.param = resultContext.result
			nextContext.err = resultContext.err
			p.nextPipeline = p.nextPipeline.nextPipeline
		} else if p.nextPipeline.parallelMethod != nil {

			result := NewResult()
			ch := make(chan Result, len(p.nextPipeline.parallelMethod))
			for i := 0; i < len(p.nextPipeline.parallelMethod); i++ {
				go func(i int) {
					ch <- p.nextPipeline.parallelMethod[i](nextContext)
				}(i)
			}
			tempPipeResult := NewResult()
			for j := len(p.nextPipeline.parallelMethod); j > 0; j-- {
				tempPipeResult = <-ch
				for k, v := range tempPipeResult.result {
					result.result[k] = v
				}
				nextContext.err = tempPipeResult.err
			}

			nextContext.param = result.result
			nextContext.err = result.err
			p.nextPipeline = p.nextPipeline.nextPipeline

		}
	}
}

func NewStarter() *Starter {
	p := &Starter{}
	p.context = KutkhData{}
	p.context.param = make(map[string]interface{})
	p.context.result = make(map[string]interface{})
	return p
}
