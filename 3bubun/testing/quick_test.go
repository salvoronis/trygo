package pad

import (
  "testing"
  "testing/quick"
  "strings"
)

func Pad(s string, max uint) string{
  ln := uint(len(s))
  if ln > max {
    return s[:max] //max - 1 error!!!
  }
  s += strings.Repeat(" ", int(max-ln))
  return s
}

func TestPadGenerative(t *testing.T){
  fn := func(s string, max uint8) bool {
    p := Pad(s, uint(max))
    return len(p) == int(max)
  }

  if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
    t.Error(err)
  }
}
