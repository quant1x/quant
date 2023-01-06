package examples

import (
	"fmt"
	tf "github.com/quant1x/quant/utils/tensorflow"
	"github.com/quant1x/quant/utils/tensorflow/op"
	"testing"
)

func TestTensorFlowByGolang(t *testing.T) {
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
}
