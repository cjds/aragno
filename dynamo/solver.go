package dynamo

import "fmt"
import "aragno/zero"

// Constraints are used for specifying that a fixed distance must be maintained between two bodies (or a body and a position)
type Constraint struct {
	bodyA *Body
	bodyB *Body
	pointA Point
	pointB Point
}


// Solve Constraints via projected gauss siedl
func (c *Constraint) solve(dt float64t) {

}


type Composite struct{
	b []*Body

}

func (c *Composite) AddBody(b Body, c Constraint){

}

func (c *Composite) SolveConstraints() {

}

func CollisionResolver() {

}
