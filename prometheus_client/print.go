package prometheus_client

import (
	"fmt"
	"github.com/prometheus/common/model"
)

/**
 * @description: 展示查询结果数据
 * @param {model.Value} value
 * @return {*}
 */
func ShowQryResult(value model.Value) {
	switch value.Type() {
	case model.ValNone:
		fmt.Println("None Type")
	case model.ValScalar:
		fmt.Println("Scalar Type")
		v, _ := value.(*model.Scalar)
		displayScalar(v)
	case model.ValVector:
		fmt.Println("Vector Type")
		v, _ := value.(model.Vector)
		displayVector(v)
	case model.ValMatrix:
		fmt.Println("Matrix Type")
		v, _ := value.(model.Matrix)
		displayMatrix(v)
	case model.ValString:
		fmt.Println("String Type")
		v, _ := value.(*model.String)
		displayString(v)
	default:
		fmt.Printf("Unknow Type")
	}
}

// ================ show different type prometheus data struct. ================ //

/**
 * @description: print Scalar.
 * @param {*model.Scalar} v
 * @return {*}
 */
func displayScalar(v *model.Scalar) {
	fmt.Printf("%s %s\n", v.Timestamp.String(), v.Value.String())
}

/**
 * @description: print Vector.
 * @param {model.Vector} v
 * @return {*}
 */
func displayVector(v model.Vector) {
	for _, i := range v {
		fmt.Println(i.Value.String())
		fmt.Printf("%s %s %s\n", i.Timestamp.String(), i.Metric.String(), i.Value.String())
	}
}

/**
 * @description: print Matric.
 * @param {model.Matrix} v
 * @return {*}
 */
func displayMatrix(v model.Matrix) {
	for _, i := range v {
		fmt.Printf("%s\n", i.Metric.String())
		for _, j := range i.Values {
			fmt.Printf("\t%s %s\n", j.Timestamp.String(), j.Value.String())
		}
	}
}

/**
 * @description: print String.
 * @param {*model.String} v
 * @return {*}
 */
func displayString(v *model.String) {
	fmt.Printf("%s %s\n", v.Timestamp.String(), v.Value)
}
