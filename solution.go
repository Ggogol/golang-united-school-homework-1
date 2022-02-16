package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	WorldMap := emoji.Sprint("Hello, :WorldMap:!")
	return WorldMap
}
