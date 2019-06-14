package provider

import (
	"testing"

	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/stretchr/testify/assert"
)

func TestFromClientCapacityEmpty(t *testing.T) {
	// Given I have empy capacity
	// When I convert it to provider capacity
	newCapacity := fromClientCapacity(nil)
	// Then I should have empty
	assert.Nil(t, newCapacity)
}

func TestFromClientCapacityExpression(t *testing.T) {
	// Given I have capacity with expression
	c := &client.Capacity{
		Max:     "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['max'] }",
		Desired: "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['desired'] }",
		Min:     "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['min'] }",
	}
	// When I convert it to provider capacity
	newCapacityArray := *fromClientCapacity(c)
	// Then I should have a capacity array of size one
	assert.Len(t, newCapacityArray, 1)
	newCapacity := newCapacityArray[0]
	// And I should have expression fields filled
	assert.Equal(t, c.Max, newCapacity.MaxExpression)
	assert.Equal(t, c.Desired, newCapacity.DesiredExpression)
	assert.Equal(t, c.Min, newCapacity.MinExpression)
	// And I should have nil in the non expression fields
	assert.Equal(t, 0, newCapacity.Max)
	assert.Equal(t, 0, newCapacity.Desired)
	assert.Equal(t, 0, newCapacity.Min)
}

func TestFromClientCapacityNumerical(t *testing.T) {
	// Given I have capacity with numerical values
	c := &client.Capacity{
		Max:     "3",
		Desired: "2",
		Min:     "1",
	}
	// When I convert it to provider capacity
	newCapacityArray := *fromClientCapacity(c)
	// Then I should have a capacity array of size one
	assert.Len(t, newCapacityArray, 1)
	newCapacity := newCapacityArray[0]
	// And I should have numerical fields filled
	assert.Equal(t, 3, newCapacity.Max)
	assert.Equal(t, 2, newCapacity.Desired)
	assert.Equal(t, 1, newCapacity.Min)
	// And I should have empty in the expression fields
	assert.Empty(t, newCapacity.DesiredExpression)
	assert.Empty(t, newCapacity.MaxExpression)
	assert.Empty(t, newCapacity.MinExpression)
}

func TestToClientCapacityEmpty(t *testing.T) {
	// Given I have empy capacity
	// When I convert it to provider capacity
	newCapacity := toClientCapacity(nil)
	// Then I should have empty
	assert.Nil(t, newCapacity)
}

func TestToClientCapacityExpression(t *testing.T) {
	// Given I have capacity with expression
	c := capacity{
		MaxExpression:     "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['max'] }",
		DesiredExpression: "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['desired'] }",
		MinExpression:     "${ #stage('My webhook')['context']['buildInfo']['scalingPolicies']['min'] }",
	}
	// When I convert it to client capacity
	newCapacity := *toClientCapacity(&[]*capacity{&c})
	// And I should have clients fields filled with expressions
	assert.Equal(t, c.MaxExpression, newCapacity.Max)
	assert.Equal(t, c.DesiredExpression, newCapacity.Desired)
	assert.Equal(t, c.MinExpression, newCapacity.Min)
}

func TestToClientCapacityNumerical(t *testing.T) {
	// Given I have capacity with numerical values
	c := capacity{
		Max:     3,
		Desired: 2,
		Min:     1,
	}
	// When I convert it to client capacity
	newCapacity := *toClientCapacity(&[]*capacity{&c})
	// And I should have clients fields filled with numerals
	assert.Equal(t, "3", newCapacity.Max)
	assert.Equal(t, "2", newCapacity.Desired)
	assert.Equal(t, "1", newCapacity.Min)
}
