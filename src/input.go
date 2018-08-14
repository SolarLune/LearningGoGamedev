package main

import (
	"github.com/hajimehoshi/ebiten"
)

type KeyInput struct {
	Key ebiten.Key
}

var actionsToInputs = make(map[string][]KeyInput, 0)
var actionStatus = make(map[string]int, 0)

func AddAction(actionName string, keyInput KeyInput) {

	_, exists := actionsToInputs[actionName]

	if !exists {
		actionsToInputs[actionName] = make([]KeyInput, 0)
	}

	actionsToInputs[actionName] = append(actionsToInputs[actionName], keyInput)

	_, exists = actionStatus[actionName]
	if !exists {
		actionStatus[actionName] = 0
	}

}

func UpdateInputMap() {

	for actionName, aSet := range actionsToInputs {

		pushed := false

		for _, input := range aSet {

			if ebiten.IsKeyPressed(input.Key) {

				pushed = true

				break

			}

		}

		if pushed && actionStatus[actionName] != 3 {

			if actionStatus[actionName] == 0 {
				actionStatus[actionName] = 1
			} else if actionStatus[actionName] == 1 {
				actionStatus[actionName] = 2
			}

		} else {

			if actionStatus[actionName] != 3 && actionStatus[actionName] != 0 {
				actionStatus[actionName] = 3
			} else {
				actionStatus[actionName] = 0
			}

		}

	}

}

func IsActionJustPressed(actionName string) bool {
	return actionStatus[actionName] == 1
}

func IsActionJustPressedI(actionName string) int {
	if IsActionJustPressed(actionName) {
		return 1
	}
	return 0
}

func IsActionDown(actionName string) bool {
	return actionStatus[actionName] == 1 || actionStatus[actionName] == 2
}

func IsActionDownI(actionName string) int {
	if IsActionDown(actionName) {
		return 1
	}
	return 0
}

func IsActionJustReleased(actionName string) bool {
	return actionStatus[actionName] == 3
}

func IsActionJustReleasedI(actionName string) int {
	if IsActionJustReleased(actionName) {
		return 1
	}
	return 0
}
