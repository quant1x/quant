package linear

import (
	"fmt"
	"math"
)

type LinearRegressionT01 struct {
	W float64
	B float64
}

type MultiVarLinearRegression struct {
	W []float64
	B float64
}

func (model *LinearRegressionT01) Train(features, targets []float64) (float64, float64, []float64, [][]float64) {
	w_init := 0.0
	b_init := 0.0
	alpha := 0.01
	num_iterations := 10000

	w_final, b_final, J_hist, p_hist := gradient_descent_single_var(features, targets, w_init, b_init, alpha, num_iterations)

	model.W = w_final
	model.B = b_final

	return w_final, b_final, J_hist, p_hist
}

func (model *MultiVarLinearRegression) Train(features [][]float64, targets []float64, alpha float64, num_iterations int) ([]float64, float64, []float64) {
	w_init := make([]float64, len(features[0]))
	b_init := 0.0

	w_final, b_final, J_hist := gradient_descent_multi_var(features, targets, w_init, b_init, alpha, num_iterations)

	fmt.Printf("b_final: %f\n", b_final)

	model.W = w_final
	model.B = b_final

	return w_final, b_final, J_hist
}

func (model *LinearRegressionT01) Predict(x float64) float64 {
	return model.W*x + model.B
}

func (model *MultiVarLinearRegression) Predict(x []float64) float64 {
	return dot_product(model.W, x) + model.B
}

func (model *LinearRegressionT01) GetParameters() (float64, float64) {
	return model.W, model.B
}

func (model *MultiVarLinearRegression) GetParameters() ([]float64, float64) {
	return model.W, model.B
}

func compute_cost_single_var(x, y []float64, w, b float64) float64 {
	m := float64(len(x))
	cost_sum := 0.0

	for idx, _ := range x {
		f_wb := w*x[idx] + b
		cost := square(f_wb - y[idx])
		cost_sum = cost_sum + cost
	}
	total_cost := (1 / (2 * m)) * cost_sum
	return total_cost
}

func compute_cost_multi_var(x [][]float64, y []float64, w []float64, b float64) float64 {
	m := float64(len(x))
	cost := 0.0

	for idx, _ := range x {
		f_wb := dot_product(x[idx], w) + b
		cost = square(f_wb-y[idx]) + cost
	}
	total_cost := cost / (2 * m)
	return total_cost
}

func gradient_descent_single_var(x, y []float64, w_init, b_init, alpha float64, num_iterations int) (float64, float64, []float64, [][]float64) {
	J_history := make([]float64, 0)
	parameter_history := make([][]float64, 0)
	b := b_init
	w := w_init

	for idx := 0; idx < num_iterations; idx++ {
		dj_dw, dj_db := compute_gradient_single_var(x, y, w, b)

		b = b - (alpha * dj_db)
		w = w - (alpha * dj_dw)

		if idx < 100000 {
			J_history = append(J_history, compute_cost_single_var(x, y, w, b))
			parameter_history = append(parameter_history, []float64{w, b})
		}
	}
	return w, b, J_history, parameter_history
}

func gradient_descent_multi_var(x [][]float64, y, w_init []float64, b_init float64, alpha float64, num_iterations int) ([]float64, float64, []float64) {
	J_history := make([]float64, 0)
	w := w_init
	b := b_init

	for i := 0; i < num_iterations; i++ {
		dj_dw, dj_db := compute_gradient_multi_var(x, y, w, b)

		alpha_product := multiply_vector(dj_dw, alpha)

		w = subtract_vectors(w, alpha_product)
		b = b - alpha*dj_db

		cost := compute_cost_multi_var(x, y, w, b)
		J_history = append(J_history, cost)
	}

	return w, b, J_history
}

func compute_gradient_single_var(x []float64, y []float64, w float64, b float64) (float64, float64) {
	// Computes the gradient for linear regression
	// Args:
	//   x: Data, m examples
	//   y: target values
	//   w,b: model parameters
	// Returns
	//   dj_dw: The gradient of the cost w.r.t. the parameters w
	//   dj_db: The gradient of the cost w.r.t. the parameter b

	m := float64(len(x))
	dj_w := 0.0
	dj_b := 0.0

	for idx := range x {
		f_wb := w*x[idx] + b
		dj_w_i := (f_wb - y[idx]) * x[idx]
		dj_b_i := f_wb - y[idx]

		dj_w = dj_w + dj_w_i
		dj_b = dj_b + dj_b_i
	}
	dj_w = dj_w / m
	dj_b = dj_b / m

	return dj_w, dj_b
}

func compute_gradient_multi_var(x [][]float64, y, w []float64, b float64) ([]float64, float64) {
	m := len(x)
	feature_count := len(x[0])

	dj_dw := make([]float64, feature_count)
	dj_db := 0.0

	for i, row := range x {
		err := (dot_product(row, w) + b) - y[i]
		for j := range w {
			dj_dw[j] = dj_dw[j] + err*row[j]
		}
		dj_db = dj_db + err
	}
	dj_dw = mapFunc(dj_dw, func(derivative float64) float64 {
		return derivative / float64(m)
	})

	dj_db = dj_db / float64(m)

	return dj_dw, dj_db
}

func dot_product(v1, v2 []float64) float64 {
	dotProduct := make([]float64, 0)
	for idx, _ := range v1 {
		dotProduct = append(dotProduct, v1[idx]*v2[idx])
	}

	result := reduce(dotProduct, func(acc, current float64) float64 {
		return acc + current
	}, 0)

	return result
}

func multiply_vector(vector []float64, multiplier float64) []float64 {
	return mapFunc(vector, func(vector_value float64) float64 {
		return vector_value * multiplier
	})
}

func subtract_vectors(v1, v2 []float64) []float64 {
	new_vector := make([]float64, 0)
	for i := 0; i < len(v1); i++ {
		new_vector = append(new_vector, v1[i]-v2[i])
	}
	return new_vector
}

func divide_vectors(v1, v2 []float64) []float64 {
	new_vector := make([]float64, 0)
	for i := 0; i < len(v1); i++ {
		new_vector = append(new_vector, v1[i]/v2[i])
	}
	return new_vector
}

func ZScoreNormalize(features []float64, means []float64, sigmas []float64) []float64 {
	return divide_vectors(subtract_vectors(features, means), sigmas)
}

func ZScoreNormalizeDataset(dataset [][]float64) ([][]float64, []float64, []float64) {
	var means, sigmas []float64

	feature_count := len(dataset[0])
	normalized_dataset := make([][]float64, len(dataset))

	// populate empty normalized dataset with empty arrays of length equal to feature count
	for idx, _ := range normalized_dataset {
		normalized_dataset[idx] = make([]float64, feature_count)
	}

	for column_index := 0; column_index < feature_count; column_index++ {
		column := pick_column(dataset, column_index)
		mean := calcMean(column)
		sigma := calcStdDev(column, mean)

		means = append(means, mean)
		sigmas = append(sigmas, sigma)

		for row_index, row := range dataset {
			value := row[column_index]
			normalized_dataset[row_index][column_index] = (value - mean) / sigma
		}
	}
	return normalized_dataset, means, sigmas
}

func pick_column(dataset [][]float64, column_index int) []float64 {
	column := []float64{}
	for _, row := range dataset {
		column = append(column, row[column_index])
	}
	return column
}

func square[T int | float64](num T) T {
	return num * num
}

func sigmoid(z float64) float64 {
	g := 1 / (1 + math.Exp(-z))
	return g
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func mapFunc[T, R any](slice []T, f func(T) R) []R {
	new_slice := make([]R, 0)
	for _, value := range slice {
		new_slice = append(new_slice, f(value))
	}
	return new_slice
}

func calcMean(sample []float64) float64 {
	sum := 0.0
	for _, value := range sample {
		sum = sum + value
	}
	return sum / float64(len(sample))
}

func calcStdDev(sample []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range sample {
		value = square(math.Abs(value - mean))
		sum = sum + value
	}
	variance := sum / float64(len(sample)-1)
	return math.Sqrt(variance)
}
