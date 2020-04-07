package t38c

type FenceDetectAction string

const (
	Inside  FenceDetectAction = "inside"
	Outside FenceDetectAction = "outside"
	Enter   FenceDetectAction = "enter"
	Exit    FenceDetectAction = "exit"
	Cross   FenceDetectAction = "cross"
)

func Detect(actions ...FenceDetectAction) []FenceDetectAction {
	return actions
}
