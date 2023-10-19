package main

func DebugLog(args ...string) []string {
  out := []string{"[DEBUG]"}
  out = append(out,args...)
  return out
}

func InfoLog(args ...string) []string {
  out := []string{"[INFO]"}
  out = append(out,args...)
  return out
}

func ErrorLog(args ...string) []string {
  out := []string{"[ERROR]"}
  out = append(out,args...)
  return out
}
