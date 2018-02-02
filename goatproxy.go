package main

import (
  "context"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/mmcdole/gofeed"
)

// HandleRequest handles request done via AWS Lambda or API Proxy.
func HandleRequest(ctx context.Context, params interface{}) (interface{}, error) {
  fp := gofeed.NewParser()
  feed, _ := fp.ParseURL("https://s3-ap-northeast-1.amazonaws.com/zipang/blog/user/8zm8yP5C/8zm8yPoR.feed")
  return feed, nil
}

func main() {
  lambda.Start(HandleRequest)
}