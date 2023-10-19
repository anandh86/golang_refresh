package main

func DescribeNumber(x int) string {

  if x == 0 {
    return "zero"
  } else if x > 0 {
    return "positive"
  } else {
    return "negative"
  }

}
