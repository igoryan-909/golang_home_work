package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = executeStage(out, done, stage)
	}

	return out
}

func executeStage(in In, done In, stage Stage) Out {
	out := make(Bi)
	stageResult := stage(in)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-stageResult:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()

	return out
}
